// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Enable nico service",
	Long:  `Enable nico server service, listening network address for external services`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		// todo pull up main server and init
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
