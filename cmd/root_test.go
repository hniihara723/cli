/*
Package cmd includes all of the Create Go App CLI commands.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

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
	"testing"
)

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully init",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConfig()
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
		})
	}
}