// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winver

import "syscall"

var VistaOrLater bool

func init() {
	v, err := syscall.GetVersion()
	if err != nil {
		// GetVersion may be altered or unavailable in the future.
		VistaOrLater = true
		return
	}
	if byte(v) >= 6 {
		VistaOrLater = true
	}
}
