// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package memory

import (
	"golang.org/x/net/context"

	"github.com/luci/gae/impl/dummy"
	"github.com/luci/gae/service/info"
)

type giContextKeyType int

var giContextKey giContextKeyType

func curGID(c context.Context) *globalInfoData {
	return c.Value(giContextKey).(*globalInfoData)
}

// useGI adds a gae.GlobalInfo context, accessible
// by gae.GetGI(c)
func useGI(c context.Context) context.Context {
	return info.SetFactory(c, func(ic context.Context) info.Interface {
		return &giImpl{dummy.Info(), curGID(ic), ic}
	})
}

// globalAppID is the 'AppID' of everythin returned from this memory
// implementation (DSKeys, GlobalInfo, etc.). There's no way to modify this
// value through the API, and there are a couple bits of code where it's hard to
// route this value through to without making the internal APIs really complex.
const globalAppID = "dev~app"

type globalInfoData struct {
	namespace string
}

type giImpl struct {
	info.Interface
	*globalInfoData
	c context.Context
}

var _ = info.Interface((*giImpl)(nil))

func (gi *giImpl) Namespace(ns string) (ret context.Context, err error) {
	return context.WithValue(gi.c, giContextKey, &globalInfoData{ns}), nil
}

func (gi *giImpl) AppID() string {
	return globalAppID
}