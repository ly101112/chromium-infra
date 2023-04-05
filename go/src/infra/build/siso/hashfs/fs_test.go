// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package hashfs_test

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"infra/build/siso/hashfs"
	"infra/build/siso/reapi/digest"
	"infra/build/siso/reapi/merkletree"
)

func TestStamp(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	execRoot := t.TempDir()
	opt := hashfs.Option{}
	hfs, err := hashfs.New(ctx, opt)
	if err != nil {
		t.Fatalf("New=%v", err)
	}
	defer func() {
		err := hfs.Close(ctx)
		if err != nil {
			t.Fatalf("hfs.Close=%v", err)
		}
	}()
	// `go test -count=1000` would easily catch the race.
	for i := 0; i < 100; i++ {
		fname := filepath.Join("obj/components", fmt.Sprintf("%d.stamp", i))
		t.Run(fname, func(t *testing.T) {
			t.Parallel()
			var cmdhash []byte
			now := time.Now()
			_, err := hfs.Stat(ctx, execRoot, fname)
			if err == nil {
				t.Fatalf("Stat(%s)=_, %v; want nil error", fname, err)
			}
			t.Logf("Write(%q, %v)", fname, now)
			err = hfs.WriteFile(ctx, execRoot, fname, nil, false, now, cmdhash)
			if err != nil {
				t.Errorf("Write(%s)=%v; want nil error", fname, err)
			}
			fi, err := hfs.Stat(ctx, execRoot, fname)
			if err != nil {
				t.Fatalf("Stat(%s)=_, %v; want nil error", fname, err)
			}
			if got, want := fi.Name(), filepath.Base(fname); got != want {
				t.Errorf("fi.Name()=%q; want=%q", got, want)
			}
			if fi.Size() != 0 {
				t.Errorf("fi.Size()=%d; want=0", fi.Size())
			}
			if got, want := fi.Mode(), fs.FileMode(0644); got != want {
				t.Errorf("fi.Mode()=%v; want=%v", got, want)
			}
			if got, want := fi.ModTime(), now; !got.Equal(want) {
				t.Errorf("fi.ModTime()=%v; want=%v", got, want)
			}
			if fi.IsDir() {
				t.Errorf("fi.IsDir()=true; want=false")
			}
			fullname := filepath.ToSlash(filepath.Join(execRoot, fname))
			got, ok := fi.Sys().(merkletree.Entry)
			if !ok {
				t.Fatalf("fi.Sys()=%T, want merkletree.Entry", fi.Sys())
			}
			if got.Name != fullname {
				t.Errorf("entry.Name=%q, want=%q", got.Name, fullname)
			}
			if got.Data.Digest() != digest.Empty {
				t.Errorf("entry.Data.Digest=%v, want=%v", got.Data.Digest(), digest.Empty)
			}
		})
	}
}

func setupFiles(t *testing.T, dir string, files map[string]string) {
	t.Helper()
	for k, v := range files {
		fname := filepath.Join(dir, k)
		err := os.MkdirAll(filepath.Dir(fname), 0755)
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile(fname, []byte(v), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadDir(t *testing.T) {
	dir := t.TempDir()
	setupFiles(t, dir, map[string]string{
		"base/base.h":        "",
		"base/debug/debug.h": "",
		"base/version.h":     "",
	})
	ctx := context.Background()
	hashFS, err := hashfs.New(ctx, hashfs.Option{})
	if err != nil {
		t.Fatalf("hashfs.New(...)=_, %v; want nil err", err)
	}
	defer hashFS.Close(ctx)
	t.Logf("base/debug")
	dents, err := hashFS.ReadDir(ctx, dir, "base/debug")
	if err != nil {
		t.Errorf("hashfs.ReadDir(ctx, %q, %q)=%v, %v; want nil err", dir, "base/debug", dents, err)
	}
	if got, want := len(dents), 1; got != want {
		t.Errorf("len(dents)=%d; want=%d", got, want)
	} else {
		if got, want := dents[0].Name(), "debug.h"; got != want {
			t.Errorf("Name=%q; want=%q", got, want)
		}
		if got, want := dents[0].IsDir(), false; got != want {
			t.Errorf("IsDir=%t; want=%t", got, want)
		}
	}

	t.Logf("base")
	dents, err = hashFS.ReadDir(ctx, dir, "base")
	if err != nil {
		t.Errorf("hashfs.ReadDir(ctx, %q, %q)=%v, %v; want nil err", dir, "base", dents, err)
	}
	if got, want := len(dents), 3; got != want {
		t.Errorf("len(dents)=%d; want=%d", got, want)
	}
	got := make([]string, 0, len(dents))
	for _, dent := range dents {
		got = append(got, dent.Name())
		// debug is dir, others are not.
		isDir := dent.Name() == "debug"
		if got, want := dent.IsDir(), isDir; got != want {
			t.Errorf("%s IsDir=%t; want=%t", dent.Name(), got, want)
		}
	}
	sort.Strings(got)
	want := []string{"base.h", "debug", "version.h"}
	sort.Strings(want)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("names; -want +got:\n%s", diff)
	}
}