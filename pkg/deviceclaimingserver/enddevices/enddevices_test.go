// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package enddevices

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

var (
	supportedJoinEUI   = &types.EUI64{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C}
	unsupportedJoinEUI = &types.EUI64{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0D}
)

func TestUpstream(t *testing.T) {
	a, ctx := test.New(t)

	mock := mockEDCS{}

	// Invalid configs
	conf := &Config{
		Source: "directory",
	}
	upstream, err := NewUpstream(ctx, *conf, mock)
	a.So(err, should.NotBeNil)

	// Upstream test
	conf = &Config{
		NetID:     test.DefaultNetID,
		Source:    "directory",
		Directory: "testdata",
	}

	upstream, err = NewUpstream(ctx, *conf, mock, WithDeviceRegistry(mock))

	test.Must(upstream, err)
	ctx = rights.NewContextWithFetcher(ctx, mock)

	// Invalid JoinEUI.
	ids, err := upstream.Claim(ctx, &ttnpb.ClaimEndDeviceRequest{
		SourceDevice: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers_{
			AuthenticatedIdentifiers: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers{
				JoinEui:            *unsupportedJoinEUI,
				DevEui:             types.EUI64{0x00, 0x04, 0xA3, 0x0B, 0x00, 0x1C, 0x05, 0x30},
				AuthenticationCode: "secret",
			},
		},
		TargetApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "test-app",
		},
	})
	a.So(errors.IsAborted(err), should.BeTrue)
	a.So(ids, should.BeNil)

	_, err = upstream.Unclaim(ctx, &ttnpb.EndDeviceIdentifiers{
		DeviceId: "test-dev",
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "test-app",
		},
		JoinEui: unsupportedJoinEUI,
		DevEui:  &types.EUI64{0x00, 0x04, 0xA3, 0x0B, 0x00, 0x1C, 0x05, 0x30},
	})
	a.So(errors.IsUnauthenticated(err), should.BeTrue)

	resp, err := upstream.GetInfoByJoinEUI(ctx, &ttnpb.GetInfoByJoinEUIRequest{
		JoinEui: unsupportedJoinEUI,
	})
	a.So(err, should.BeNil)
	a.So(resp.SupportsClaiming, should.BeFalse)

	// Valid JoinEUI.
	inf, err := upstream.GetInfoByJoinEUI(ctx, &ttnpb.GetInfoByJoinEUIRequest{
		JoinEui: supportedJoinEUI,
	})
	a.So(err, should.BeNil)
	a.So(inf.JoinEui, should.Resemble, supportedJoinEUI)
	a.So(inf.SupportsClaiming, should.BeTrue)

	ids, err = upstream.Claim(ctx, &ttnpb.ClaimEndDeviceRequest{
		SourceDevice: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers_{
			AuthenticatedIdentifiers: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers{
				JoinEui:            *supportedJoinEUI,
				DevEui:             types.EUI64{0x00, 0x04, 0xA3, 0x0B, 0x00, 0x1C, 0x05, 0x30},
				AuthenticationCode: "secret",
			},
		},
		TargetApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "test-app",
		},
	})
	a.So(!errors.IsUnimplemented(err), should.BeTrue)

	_, err = upstream.Unclaim(ctx, &ttnpb.EndDeviceIdentifiers{
		DeviceId: "test-dev",
		ApplicationIds: &ttnpb.ApplicationIdentifiers{
			ApplicationId: "test-app",
		},
		JoinEui: supportedJoinEUI,
		DevEui:  &types.EUI64{0x00, 0x04, 0xA3, 0x0B, 0x00, 0x1C, 0x05, 0x30},
	})
	a.So(!errors.IsUnavailable(err), should.BeTrue)
}
