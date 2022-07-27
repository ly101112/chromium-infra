// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package handlers

import (
	"go.chromium.org/luci/server/router"
)

// RegisterRoutes registers routes explicitly handled by the handler.
func (h *Handlers) RegisterRoutes(r *router.Router, mw router.MiddlewareChain) {
	// TODO(b/239768873): Migrate to pRPC and apply permission checks.
	r.GET("/api/projects/:project/clusters/:algorithm/:id/failures", mw, h.GetClusterFailures)
	r.GET("/api/projects/:project/config", mw, h.GetConfig)
	r.GET("/api/authState", mw, h.GetAuthState)
	r.GET("/", mw, h.IndexPage)
}
