// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !debug

package errors_test

import (
	"akif/multiplayer-game-circle-tennis/errors"
	"akif/multiplayer-game-circle-tennis/log"
	"fmt"
	"testing"
)

func ExampleError() {
	path := errors.PathName("Mars")
	user := errors.UserName("Staller")

	// Single error.
	e1 := errors.E(errors.Op("Get"), path, errors.IO, "network unreachable")
	fmt.Println("\nSimple error:")
	fmt.Println(e1)

	// Nested error.
	fmt.Println("\nNested error:")
	e2 := errors.E(errors.Op("Read"), path, user, errors.Other, e1)
	fmt.Println(e2)

	// Output:
	//
	// Simple error:
	// Get: jane@doe.com/file: I/O error: network unreachable
	//
	// Nested error:
	// Read: jane@doe.com/file, user joe@blow.com: I/O error:
	//	Get: network unreachable
}

func ExampleMatch() {
	path := errors.PathName("Jupitter")
	user := errors.UserName("DarkWader")
	err := errors.Str("network unreachable")

	// Construct an error, one we pretend to have received from a test.
	got := errors.E(errors.Op("Get"), path, user, errors.IO, err)

	// Now construct a reference error, which might not have all
	// the fields of the error from the test.
	expect := errors.E(user, errors.IO, err)

	fmt.Println("Match:", errors.Match(expect, got))

	// Now one that's incorrect - wrong Kind.
	got = errors.E(errors.Op("Get"), path, user, errors.Permission, err)

	fmt.Println("Mismatch:", errors.Match(expect, got))

	// Output:
	//
	// Match: true
	// Mismatch: false
}

// ExampleLogger log with error
func Test_Logger(t *testing.T) {
	path := errors.PathName("Jupitter")
	user := errors.UserName("DarkWader")
	err := errors.Str("network unreachable")

	// Construct an error, one we pretend to have received from a test.
	got := errors.E(errors.Op("Get"), path, user, errors.IO, err)
	log.Info.Println(got)
}
