/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package ubuntu

import (
	"github.com/megamsys/megdc/templates"
	"github.com/megamsys/urknall"
)

var ubuntugatewayremove *UbuntuGatewayRemove

func init() {
	ubuntugatewayremove = &UbuntuGatewayRemove{}
	templates.Register("UbuntuGatewayRemove", ubuntugatewayremove)
}

type UbuntuGatewayRemove struct{}

func (tpl *UbuntuGatewayRemove) Render(p urknall.Package) {
	p.AddTemplate("gateway", &UbuntuGatewayRemoveTemplate{})
}

func (tpl *UbuntuGatewayRemove) Options(t *templates.Template) {
}

func (tpl *UbuntuGatewayRemove) Run(target urknall.Target) error {
	return urknall.Run(target, &UbuntuGatewayRemove{})
}

type UbuntuGatewayRemoveTemplate struct{}

func (m *UbuntuGatewayRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("megamgateway",
		RemovePackage("megamgateway"),
		RemovePackages(""),
		PurgePackages("megamgateway"),
		Shell("dpkg --get-selections megam*"),
	)
	pkg.AddCommands("megamgateway-clean",
		Shell("rm -r /var/lib/urknall/gateway*"),
	)
}
