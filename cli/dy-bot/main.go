// Copyright 2018 The Dongyue members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/spf13/cobra"

	"github.com/dyweb/dy-bot/cli/dy-bot/server"
	"github.com/dyweb/dy-bot/cli/dy-bot/server/config"
	"github.com/dyweb/dy-bot/pkg/gh"
	"github.com/dyweb/dy-bot/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

func main() {
	var cfg config.Config
	var cmdServe = &cobra.Command{
		Use:  "Dongyue Bot",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			gh.InitGitHubClient(cfg.Owner, cfg.Repo, cfg.AccessToken)
			s := server.NewServer(cfg)
			log.Fatal(s.Run())
		},
	}

	flagSet := cmdServe.Flags()
	flagSet.StringVarP(&cfg.Owner, "owner", "o", "", "GitHub ID to which connect in GitHub")
	flagSet.StringVarP(&cfg.Repo, "repo", "r", "", "GitHub repo to which connect in GitHub")
	flagSet.StringVarP(&cfg.HTTPListen, "listen", "l", "", "where does automan listened on")
	flagSet.StringVarP(&cfg.AccessToken, "token", "t", "", "access token to have some control on resources")

	cmdServe.Execute()
}
