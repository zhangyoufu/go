// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winver

import "syscall"

var (
	VistaOrLater bool
	Win7OrLater  bool
)

func init() {
	v, err := syscall.GetVersion()
	if err != nil {
		// GetVersion may be altered or unavailable in the future.
		VistaOrLater = true
		return
	}
	major := byte(v)
	minor := byte(v >> 8)
	if major >= 6 {
		VistaOrLater = true
	}
	if major > 6 || major == 6 && minor >= 1 {
		Win7OrLater = true
	}
}
