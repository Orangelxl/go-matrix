// Copyright (c) 2018-2019 The MATRIX Authors
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64,!386,!ppc64le appengine

package sha3

var (
	xorIn            = xorInGeneric
	copyOut          = copyOutGeneric
	xorInUnaligned   = xorInGeneric
	copyOutUnaligned = copyOutGeneric
)

const xorImplementationUnaligned = "generic"
