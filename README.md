# go-multierror
A tiny Go library for bundling together multiple error objects

[![Build Status](https://travis-ci.org/chronos-tachyon/go-multierror.svg)](https://travis-ci.org/chronos-tachyon/go-multierror)
[![godoc](https://chronos-tachyon.net/img/godoc-badge.svg)](http://godoc.org/github.com/chronos-tachyon/go-multierror)

	import (
		"errors"
	
		"github.com/chronos-tachyon/go-multierror"
	)
	
	var (
		A = errors.New("A")
		B = errors.New("B")
		C = errors.New("C")
	)
	
	func Demo(a, b, c bool) error {
		errors := []error(nil)
		if a {
			errors = append(errors, A)
		}
		if b {
			errors = append(errors, B)
		}
		if c {
			errors = append(errors, C)
		}
		return multierror.New(errors)
	}

	// Demo(false, false, false) returns nil
	// Demo(true,  false, false) returns A
	// Demo(true,  true,  true)  returns MultiError{A, B, C}
	//
	// MultiError{A, B, C} stringifies as:
	//	encountered multiple errors:
	//		... A
	//		... B
	//		... C
