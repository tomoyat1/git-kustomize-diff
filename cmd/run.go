/*
Copyright 2021 Daisuke Taniwaki.

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

package cmd

import (
	"fmt"

	"github.com/dtaniwaki/kustomize-diff/pkg/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run kustomize-diff",
	Long:  `Run kustomize-diff`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stdout, err := utils.GetCommitHash(args[0])
		if err != nil {
			return err
		}
		fmt.Println(stdout)
		return nil
	},
}
