// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
	"errors"

	"github.com/pulumi/lumi/pkg/util/cmdutil"
	"github.com/spf13/cobra"
)

func newPackVerifyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "verify [package]",
		Short: "Check that a package's metadata and IL are correct",
		Long: "Check that a package's metadata and IL are correct\n" +
			"\n" +
			"A package contains intermediate language (IL) that encodes symbols, definitions,\n" +
			"and executable code.  This IL must obey a set of specific rules for it to be considered\n" +
			"legal and valid.  Otherwise, evaluating it will fail.\n" +
			"\n" +
			"The verify command thoroughly checks the package's IL against these rules, and issues\n" +
			"errors anywhere it doesn't obey them.  This is generally useful for tools developers\n" +
			"and can ensure that code does not fail at runtime, when such invariants are checked.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// Create a compiler object and perform the verification.
			if !verify(cmd, args) {
				return errors.New("verification failed")
			}
			return nil
		}),
	}

	return cmd
}
