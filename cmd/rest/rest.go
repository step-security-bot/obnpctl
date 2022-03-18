/*
Copyright © 2022 Ron Green

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package rest

import (
	"github.com/georgettica/obnpctl/cmd/rest/client"
	"github.com/georgettica/obnpctl/cmd/rest/server"
	"github.com/spf13/cobra"
)

// restCmd represents the rest command.
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func GetCmd() *cobra.Command {
	return restCmd
}

func init() {
	restCmd.AddCommand(client.GetCmd())
	restCmd.AddCommand(server.GetCmd())
}
