// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/context"
	"github.com/seccijr/quintoweb/util"
	"net/http"
)

func main() {
	util.RouteInstall()
	util.TemplateInstall()
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
