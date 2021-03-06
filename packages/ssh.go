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
package packages

import (
	"github.com/megamsys/libgo/cmd"
	"launchpad.net/gnuflag"
)

type SSHCommand struct {
	Host     string
	Username string
	Password string
	Fs       *gnuflag.FlagSet
}

func (cmd *SSHCommand) Flags() *gnuflag.FlagSet {
	if cmd.Fs == nil {
		cmd.Fs = gnuflag.NewFlagSet("", gnuflag.ExitOnError)
		hostMsg := "The host of the server to ssh"
		cmd.Fs.StringVar(&cmd.Host, "host", "localhost", hostMsg)
		cmd.Fs.StringVar(&cmd.Host, "o", "localhost", hostMsg)
		usrMsg := "The username of the server"
		cmd.Fs.StringVar(&cmd.Username, "username", "megdc", usrMsg)
		cmd.Fs.StringVar(&cmd.Username, "u", "megdc", usrMsg)
		pwdMsg := "The password of the server"
		cmd.Fs.StringVar(&cmd.Password, "password", "megdc", pwdMsg)
		cmd.Fs.StringVar(&cmd.Password, "p", "megdc", pwdMsg)

	}
	return cmd.Fs
}

func (c *SSHCommand) Run(context *cmd.Context) error {
	//don't do a thing, ideally invoke urnknall and ssh into the system and exit out.
	return nil
}
