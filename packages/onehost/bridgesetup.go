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
package onehost

import (
	"github.com/megamsys/libgo/cmd"
	"github.com/megamsys/megdc/handler"
	"launchpad.net/gnuflag"
)

type Createnetwork struct {
	Fs     *gnuflag.FlagSet
	PhyDev string
	Bridge string
}

func (g *Createnetwork) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "createnetwork",
		Usage:   `createnetwork [--bridge] name --[phy] name`,
		Desc:   `create network
		Default: bridge name:one, phydev:eth0
	`,
		MinArgs: 0,
	}
}

func (c *Createnetwork) Run(context *cmd.Context) error {
	handler.SunSpin(cmd.Colorfy(handler.Logo, "green", "", "bold"), "", "createnetwork")
	w := handler.NewWrap(c)
	c.createDefaultNetwork(w)
	if h, err := handler.NewHandler(w); err != nil {
		return err
	} else if err := h.Run(); err != nil {
		return err
	}
	return nil
}

func (c *Createnetwork) Flags() *gnuflag.FlagSet {
	if c.Fs == nil {
		c.Fs = gnuflag.NewFlagSet("megdc", gnuflag.ExitOnError)
		c.Fs.StringVar(&c.PhyDev, "phy", "", "Physical device or Network interface (default: eth0)")
		c.Fs.StringVar(&c.Bridge, "bridge", "", "The name of the bridge (default: one)")
	}
	return c.Fs
}
func (c *Createnetwork) createDefaultNetwork(w *handler.WrappedParms) {
	DEFAULT_PACKAGES := []string{"CreateNetwork"}

	if w.Empty() {
		for i := range DEFAULT_PACKAGES {
			w.AddPackage(DEFAULT_PACKAGES[i])
		}
	}
}
