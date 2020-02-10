// Copyright © 2020 Banzai Cloud
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

package main

import (
	"time"

	"github.com/banzaicloud/bank-vaults/pkg/sdk/vault"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Bank-Vaults on the target Kubernetes cluster",
	Long: `This configuration is an extension to what is available through the Vault configuration:
			https://www.vaultproject.io/docs/configuration/index.html. With this it is possible to
			configure secret engines, auth methods, etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		appConfig.BindPFlag(cfgOnce, cmd.PersistentFlags().Lookup(cfgOnce))
		appConfig.BindPFlag(cfgFatal, cmd.PersistentFlags().Lookup(cfgFatal))
		appConfig.BindPFlag(cfgUnsealPeriod, cmd.PersistentFlags().Lookup(cfgUnsealPeriod))
		appConfig.BindPFlag(cfgVaultConfigFile, cmd.PersistentFlags().Lookup(cfgVaultConfigFile))

	},
}

func init() {
	installCmd.PersistentFlags().Bool(cfgOnce, false, "Run configure only once")
	installCmd.PersistentFlags().Bool(cfgFatal, false, "Make configuration errors fatal to the configurator")
	installCmd.PersistentFlags().Duration(cfgUnsealPeriod, time.Second*5, "How often to attempt to unseal the Vault instance")
	installCmd.PersistentFlags().StringSlice(cfgVaultConfigFile, []string{vault.DefaultConfigFile}, "The filename of the YAML/JSON Vault configuration")

	rootCmd.AddCommand(installCmd)
}
