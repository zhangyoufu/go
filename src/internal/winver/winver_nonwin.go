// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package winver

// dummy constants that should not be referenced
const (
	VistaOrLater = false
	Win7OrLater = false
)
