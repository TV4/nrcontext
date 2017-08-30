# nrcontext


[![Build Status](https://travis-ci.org/TV4/nrcontext.svg?branch=master)](https://travis-ci.org/TV4/nrcontext)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/nrcontext)](https://goreportcard.com/report/github.com/TV4/nrcontext)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/nrcontext)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/nrcontext#license-mit)

The [go-agent](https://github.com/newrelic/go-agent) from [New Relic](https://newrelic.com/golang)
currently does not have support for storing `newrelic.Application` and `newrelic.Transaction`
in [context.Context](https://golang.org/pkg/context/#Context).

This package is an experiment in doing just that.

## Status

Still under active development. Expect some additions but few if any changes to the public API.

## Installation

    go get -u github.com/TV4/nrcontext

## Background

Initially based on ideas in the article [Using Go’s ‘context’ library for performance monitoring](https://medium.com/@gosamv/using-gos-context-library-for-performance-monitoring-aaf25dacb0fe) by Sam Vilain.

An alternative to this package is [newrelic-context](https://github.com/smacker/newrelic-context) by Maxim Sukharev.

## License (MIT)

Copyright (c) 2017 TV4

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="230" height="230">

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
