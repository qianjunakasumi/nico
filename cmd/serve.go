// Copyright (c) 2023 qianjunakasumi (https://github.com/qianjunakasumi)
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/qianjunakasumi/nico/internal/service"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Enable nico service",
	Long:  `Enable nico server service, listening network address for external services`,
	Run: func(cmd *cobra.Command, args []string) {

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM /* docker stop */)

		errCh := make(chan error)
		nico := service.NewNicoService(nil) // todo manual di
		go func() {
			errCh <- nico.Run(context.Background())
		}()

		select {
		case <-sigCh:
			log.Info().Msg("Goodbye! Greeting from nico")
			//todo close grpc/graphQL service
			//todo close database connection

			return
		case err := <-errCh:
			if err != nil {
				log.Error().Err(err).Msg("Enable nico service failed")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
