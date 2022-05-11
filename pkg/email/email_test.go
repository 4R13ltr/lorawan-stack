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

package email_test

import (
	"os"
	"testing"

	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/v3/pkg/email"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

// welcomeEmailData is the data specifically for the welcome email.
type welcomeEmailData struct {
	email.TemplateData
	ActivationToken string
}

func TestEmail(t *testing.T) {
	a, ctx := test.New(t)

	registry := email.NewTemplateRegistry()

	welcomeEmailTemplate, err := email.NewTemplateFS(
		os.DirFS("testdata"), "welcome",
		email.FSTemplate{
			SubjectTemplate:      "Welcome to {{ .Network.Name }}",
			HTMLTemplateBaseFile: "base.html",
			HTMLTemplateFile:     "welcome.html",
			TextTemplateBaseFile: "base.txt",
			TextTemplateFile:     "welcome.txt",
			IncludePatterns:      []string{"header.html", "footer.html", "header.txt", "footer.txt"},
		},
	)
	a.So(err, should.BeNil)

	registry.RegisterTemplate(welcomeEmailTemplate)

	a.So(registry.RegisteredTemplates(), should.Contain, "welcome")
	a.So(registry.GetTemplate(ctx, "welcome"), should.Resemble, welcomeEmailTemplate)

	message, err := welcomeEmailTemplate.Execute(&welcomeEmailData{
		TemplateData: email.NewTemplateData(&email.NetworkConfig{
			Name:              "The Things Network",
			IdentityServerURL: "https://eu1.cloud.thethings.network/oauth",
			ConsoleURL:        "https://console.cloud.thethings.network",
		}, &ttnpb.User{
			Name:                "John Doe",
			PrimaryEmailAddress: "john.doe@example.com",
		}),
	})

	if a.So(err, should.BeNil) && a.So(message, should.NotBeNil) {
		a.So(message.Subject, should.Equal, "Welcome to The Things Network")
		a.So(message.HTMLBody, should.ContainSubstring, `<div class="header">`)
		a.So(message.HTMLBody, should.ContainSubstring, "Welcome to The Things Network, John Doe!")
		a.So(message.HTMLBody, should.ContainSubstring, `<div class="footer">`)
		a.So(message.TextBody, should.ContainSubstring, "==================")
		a.So(message.TextBody, should.ContainSubstring, "Welcome to The Things Network, John Doe!")
	}
}
