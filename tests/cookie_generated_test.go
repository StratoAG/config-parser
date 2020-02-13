// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v2/parsers"
)

func TestCookie(t *testing.T) {
	tests := map[string]bool{
		"cookie test": true,
		"cookie myCookie domain dom1 indirect postonly": true,
		"cookie myCookie domain dom1 domain dom2 indirect postonly": true,
		"cookie myCookie indirect maxidle 10 maxlife 5 postonly": true,
		"cookie myCookie indirect maxidle 10": true,
		"cookie myCookie indirect maxlife 10": true,
		"cookie myCookie domain dom1 domain dom2 httponly indirect maxidle 10 maxlife 5 nocache postonly preserve rewrite secure": true,
		"cookie": false,
		"cookie myCookie maxidle something": false,
		"cookie myCookie maxlife something": false,
		"---": false,
		"--- ---": false,
	}
	parser := &parsers.Cookie{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			err := ProcessLine(line, parser)
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if line != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
