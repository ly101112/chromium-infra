// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gerrit

import (
	"context"

	"go.chromium.org/luci/common/errors"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
)

// Client defines a subset of Gerrit API used by rubber-stamper.
type Client interface {
}

// ClientFactory creates Client tied to Gerrit host and LUCI project.
type ClientFactory func(ctx context.Context, gerritHost string) (Client, error)

// Client must be a subset of gerritpb.Client
var _ Client = (gerritpb.GerritClient)(nil)

var clientCtxKey = "infra/appengine/rubber-stamper/internal/client/gerrit.Client"

// setClientFactory puts a given ClientFactory into in the context.
func setClientFactory(ctx context.Context, f ClientFactory) context.Context {
	return context.WithValue(ctx, &clientCtxKey, f)
}

// Setup puts a production ClientFactory into the context.
func Setup(ctx context.Context) context.Context {
	return setClientFactory(ctx, newFactory().makeClient)
}

// GetCurrentClient returns the Client in the context or an error.
func GetCurrentClient(ctx context.Context, gerritHost string) (Client, error) {
	f, _ := ctx.Value(&clientCtxKey).(ClientFactory)
	if f == nil {
		return nil, errors.New("not a valid Gerrit context, no ClientFactory available")
	}
	return f(ctx, gerritHost)
}
