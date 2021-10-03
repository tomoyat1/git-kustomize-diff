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
	"io"
	"os"
	"path/filepath"
	"strings"
)

type GitDir struct {
	WorkDir WorkDir
}

func (gd *GitDir) CommitHash(target string) (string, error) {
	stdout, _, err := gd.WorkDir.RunCommand("git", "rev-parse", "-q", "--short", target)
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}

func (gd *GitDir) CurrentBranch() (string, error) {
	stdout, _, err := gd.WorkDir.RunCommand("git", "branch", "--show-current")
	if err != nil {
		return "", err
	}
	return strings.Trim(stdout, "\n"), nil
}

func (gd *GitDir) Clone(dstDirPath string) (*GitDir, error) {
	rootDir, err := gd.GetRootDir()
	if err != nil {
		return nil, err
	}
	_, _, err = gd.WorkDir.RunCommand("git", "clone", rootDir, dstDirPath)
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(gd.WorkDir.Dir)
	if err != nil {
		return nil, err
	}
	relPath, err := filepath.Rel(rootDir, absPath)
	if err != nil {
		return nil, err
	}
	return &GitDir{
		WorkDir: WorkDir{Dir: filepath.Join(dstDirPath, relPath)},
	}, nil
}

func (gd *GitDir) GetRootDir() (string, error) {
	baseDirPath, _, err := gd.WorkDir.RunCommand("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", err
	}
	return strings.Trim(baseDirPath, "\n"), nil
}

func (gd *GitDir) CopyConfig(targetGitDir *GitDir) error {
	baseDirPath, err := gd.GetRootDir()
	if err != nil {
		return err
	}
	src, err := os.Open(filepath.Join(baseDirPath, ".git", "config"))
	if err != nil {
		return err
	}
	defer src.Close()
	targetDirPath, err := targetGitDir.GetRootDir()
	if err != nil {
		return err
	}
	dst, err := os.Create(filepath.Join(targetDirPath, ".git", "config"))
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(src, dst)
	if err != nil {
		return err
	}
	return nil
}

func (gd *GitDir) Fetch() error {
	_, _, err := gd.WorkDir.RunCommand("git", "fetch", "--all")
	if err != nil {
		return err
	}
	return nil
}

func (gd *GitDir) Checkout(target string) error {
	_, _, err := gd.WorkDir.RunCommand("git", "checkout", target)
	if err != nil {
		return err
	}
	return nil
}

func (gd *GitDir) Merge(target string) error {
	_, _, err := gd.WorkDir.RunCommand("git", "merge", "--no-ff", target)
	if err != nil {
		return err
	}
	return nil
}

func (gd *GitDir) SetUser() error {
	email := "anonymous@example.com"
	name := "anonymous"
	_, _, err := gd.WorkDir.RunCommand("git", "config", "user.email", email)
	if err != nil {
		return err
	}
	_, _, err = gd.WorkDir.RunCommand("git", "config", "user.name", name)
	if err != nil {
		return err
	}
	return nil
}
