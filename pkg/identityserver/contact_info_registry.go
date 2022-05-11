// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identityserver

import (
	"context"
	"time"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/email/templates"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	errContactNoCollaborator  = errors.DefineInvalidArgument("contact_no_collaborator", "contact `{contact}` is not a collaborator")
	errNoValidationNeeded     = errors.DefineInvalidArgument("no_validation_needed", "no validation needed for this contact info")
	errValidationsAlreadySent = errors.DefineAlreadyExists("validations_already_sent", "validations for this contact info already sent")
)

func validateCollaboratorEqualsContact(collaborator, contact *ttnpb.OrganizationOrUserIdentifiers) error {
	if contact == nil {
		return nil
	}
	if collaborator.EntityType() != contact.EntityType() || collaborator.IDString() != contact.IDString() {
		return errContactNoCollaborator.WithAttributes("contact", contact.IDString())
	}
	return nil
}

func validateContactIsCollaborator(ctx context.Context, st store.Store, contact *ttnpb.OrganizationOrUserIdentifiers, entity *ttnpb.EntityIdentifiers) error {
	if contact == nil {
		return nil
	}
	_, err := st.GetMember(ctx, contact, entity)
	if err != nil {
		if errors.IsNotFound(err) {
			return errContactNoCollaborator.WithAttributes("contact", contact.IDString())
		}
		return err
	}
	return nil
}

func (is *IdentityServer) requestContactInfoValidation(ctx context.Context, ids *ttnpb.EntityIdentifiers) (*ttnpb.ContactInfoValidation, error) {
	// NOTE: This does NOT check auth. Internal use only.
	id, err := auth.GenerateID(ctx)
	if err != nil {
		return nil, err
	}
	var contactInfo []*ttnpb.ContactInfo
	err = is.store.Transact(ctx, func(ctx context.Context, st store.Store) (err error) {
		contactInfo, err = st.GetContactInfo(ctx, ids)
		return err
	})
	if err != nil {
		return nil, err
	}
	now := time.Now()
	ttl := is.configFromContext(ctx).UserRegistration.ContactInfoValidation.TokenTTL
	expires := now.Add(ttl)
	emailValidations := make(map[string]*ttnpb.ContactInfoValidation)
	for _, info := range contactInfo {
		if info.ContactMethod == ttnpb.ContactMethod_CONTACT_METHOD_EMAIL && info.ValidatedAt == nil {
			validation, ok := emailValidations[info.Value]
			if !ok {
				key, err := auth.GenerateKey(ctx)
				if err != nil {
					return nil, err
				}
				validation = &ttnpb.ContactInfoValidation{
					Id:        id,
					Token:     key,
					Entity:    ids,
					CreatedAt: ttnpb.ProtoTimePtr(now),
					ExpiresAt: ttnpb.ProtoTimePtr(expires),
				}
				emailValidations[info.Value] = validation
			}
			validation.ContactInfo = append(validation.ContactInfo, info)
		}
	}
	if len(emailValidations) == 0 {
		return nil, errNoValidationNeeded.New()
	}

	err = is.store.Transact(ctx, func(ctx context.Context, st store.Store) (err error) {
		for email, validation := range emailValidations {
			validation, err = st.CreateValidation(ctx, validation)
			if err != nil {
				if errors.IsAlreadyExists(err) {
					delete(emailValidations, email)
					continue
				}
				return err
			}
			log.FromContext(ctx).WithFields(log.Fields(
				"email", email,
				"token", validation.Token,
			)).Info("Created email validation token")
			emailValidations[email] = validation
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var pendingContactInfo []*ttnpb.ContactInfo
	for address, validation := range emailValidations {
		// Prepare validateData outside of goroutine to avoid issues with range variable or races with unsetting the Token.
		validateData := &templates.ValidateData{
			EntityIdentifiers: validation.Entity,
			ID:                validation.Id,
			Token:             validation.Token,
			TTL:               ttl,
		}
		go is.SendTemplateEmailToUsers(is.FromRequestContext(ctx), "validate", func(_ context.Context, data email.TemplateData) (email.TemplateData, error) {
			validateData.TemplateData = data
			return validateData, nil
		}, &ttnpb.User{PrimaryEmailAddress: address})
		pendingContactInfo = append(pendingContactInfo, validation.ContactInfo...)
		validation.Token = "" // Unset tokens after sending emails
	}
	if len(pendingContactInfo) == 0 {
		return nil, errValidationsAlreadySent.New()
	}

	return &ttnpb.ContactInfoValidation{
		Id:          id,
		Entity:      ids,
		ContactInfo: pendingContactInfo,
		CreatedAt:   ttnpb.ProtoTimePtr(now),
	}, nil
}

func (is *IdentityServer) validateContactInfo(ctx context.Context, req *ttnpb.ContactInfoValidation) (*pbtypes.Empty, error) {
	err := is.store.Transact(ctx, func(ctx context.Context, st store.Store) error {
		return st.Validate(ctx, req)
	})
	if err != nil {
		return nil, err
	}
	return ttnpb.Empty, nil
}

type contactInfoRegistry struct {
	*IdentityServer
}

var errNoContactInfoForEntity = errors.DefineInvalidArgument("no_contact_info", "no contact info for this entity type")

func (cir *contactInfoRegistry) RequestValidation(ctx context.Context, ids *ttnpb.EntityIdentifiers) (*ttnpb.ContactInfoValidation, error) {
	var err error
	switch id := ids.GetIds().(type) {
	case *ttnpb.EntityIdentifiers_ApplicationIds:
		err = rights.RequireApplication(ctx, *id.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_SETTINGS_BASIC)
	case *ttnpb.EntityIdentifiers_ClientIds:
		err = rights.RequireClient(ctx, *id.ClientIds, ttnpb.Right_RIGHT_CLIENT_ALL)
	case *ttnpb.EntityIdentifiers_GatewayIds:
		err = rights.RequireGateway(ctx, *id.GatewayIds, ttnpb.Right_RIGHT_GATEWAY_SETTINGS_BASIC)
	case *ttnpb.EntityIdentifiers_OrganizationIds:
		err = rights.RequireOrganization(ctx, *id.OrganizationIds, ttnpb.Right_RIGHT_ORGANIZATION_SETTINGS_BASIC)
	case *ttnpb.EntityIdentifiers_UserIds:
		err = rights.RequireUser(ctx, *id.UserIds, ttnpb.Right_RIGHT_USER_SETTINGS_BASIC)
	default:
		return nil, errNoContactInfoForEntity.New()
	}
	if err != nil {
		return nil, err
	}
	return cir.requestContactInfoValidation(ctx, ids)
}

func (cir *contactInfoRegistry) Validate(ctx context.Context, req *ttnpb.ContactInfoValidation) (*pbtypes.Empty, error) {
	return cir.validateContactInfo(ctx, req)
}
