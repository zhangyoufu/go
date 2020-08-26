// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "syscall"

func isWin7OrLater() bool {
	v, err := syscall.GetVersion()
	if err != nil {
		// GetVersion may be altered or unavailable in the future.
		return true
	}
	major := byte(v)
	minor := byte(v >> 8)
	return major > 6 || major == 6 && minor >= 1
}
