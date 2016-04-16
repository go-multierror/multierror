# go-multierror
A tiny Go library for bundling together multiple error objects

[![godoc](https://chronos-tachyon.net/img/godoc-badge.svg)](https://godoc.org/github.com/go-multierror/multierror)
[![Build Status](https://travis-ci.org/go-multierror/multierror.svg?branch=master)](https://travis-ci.org/go-multierror/multierror)
[![Coverage Status](https://coveralls.io/repos/github/go-multierror/multierror/badge.svg?branch=master)](https://coveralls.io/github/go-multierror/multierror?branch=master)
[![GitHub release](https://img.shields.io/github/release/go-multierror/multierror.svg)]()

	import (
		"errors"
	
		"gopkg.in/multierror.v0"
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
