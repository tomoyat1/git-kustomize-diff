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

package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListKustomizeDirs(t *testing.T) {
	wd, _ := os.Getwd()

	dirs, err := ListKustomizeDirs(filepath.Join(wd, "fixtures", "kustomize"), ListKustomizeDirsOpts{})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, []string{
		"a",
		"b",
	}, dirs)

	includeRegexp, _ := regexp.Compile(".*/a$")
	dirs, err = ListKustomizeDirs(filepath.Join(wd, "fixtures", "kustomize"), ListKustomizeDirsOpts{includeRegexp: includeRegexp})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, []string{
		"a",
	}, dirs)

	excludeRegexp, _ := regexp.Compile(".*/a$")
	dirs, err = ListKustomizeDirs(filepath.Join(wd, "fixtures", "kustomize"), ListKustomizeDirsOpts{excludeRegexp: excludeRegexp})
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.Equal(t, []string{
		"b",
	}, dirs)
}
