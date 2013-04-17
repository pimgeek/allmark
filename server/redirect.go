// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	pather "github.com/andreaskoch/allmark/path"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request, route string) {

	route = "/" + pather.StripLeadingUrlDirectorySeperator(route)
	http.Redirect(w, r, route, http.StatusMovedPermanently)
}