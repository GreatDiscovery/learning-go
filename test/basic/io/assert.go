// Copyright 2016 CodisLabs. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package io

func Must(b bool) {
	if b {
		return
	}
	panic("assertion failed")
}

func MustNoError(err error) {
	if err == nil {
		return
	}
	panic("error happens, assertion failed")
}
