// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js
package teminal

import "errors"

// IsTerminal returns true if the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return false
}

// ReadPassword reads a line of input from a terminal without local echo.  This
// is commonly used for inputting passwords and other sensitive data. The slice
// returned does not include the \n.
func ReadPassword(fd int) ([]byte, error) {
	return nil, errors.New("unsupported")
}
