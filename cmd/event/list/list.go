// Copyright © 2019 VMware, INC
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

package list

import (
	"fmt"
	"github.com/edgexfoundry-holding/edgex-cli/config"
	"github.com/edgexfoundry-holding/edgex-cli/pkg/utils"
	"io"
	"strconv"
	"text/tabwriter"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var limit int32

// NewCommand returns the list device command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "A list of all device services",
		Long:  `Return all device services sorted by id.`,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error){

			var url string
			if len(args) > 0 {
				var limitUrl string
				device := args[0]
				if limit > 0 {
					limitUrl = strconv.FormatInt(int64(limit), 10)
				} else {
					limitUrl = strconv.FormatInt(int64(50), 10)
				}
				url = config.Conf.Clients["CoreData"].Url() + clients.ApiEventRoute + "/device/" + device + "/" + limitUrl
			} else {
				url = config.Conf.Clients["CoreData"].Url() + clients.ApiEventRoute
			}

			var events []models.Event
			err = client.ListHelper(url, &events)
			if err != nil {
				return
			}
			pw := viper.Get("writer").(io.WriteCloser)
			w := new(tabwriter.Writer)
			w.Init(pw, 0, 8, 1, '\t', 0)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", "Event ID", "Device", "Origin", "Created", "Modified")

			for _, event := range eventList.rd {
				fmt.Fprintf(w, "%s\t%s\t%v\t%v\t%s\t\n",
					event.ID,
					event.Device,
					event.Origin,
					utils.DisplayDuration(event.Created),
					utils.DisplayDuration(event.Modified),
				)
			}
			w.Flush()
			return
		},
	}
	cmd.Flags().Int32VarP(&limit, "limit", "l", 0, "Limit number of results")
	return cmd
}
