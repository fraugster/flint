# flint

[![Build Status](https://secure.travis-ci.org/fraugster/flint.png?branch=master)](http://travis-ci.org/fraugster/flint)

Slightly stricter linter for the Go programming language.

This code is considered to be work in progress.

### Purpose

The current functionality of flint is focused on imports formatting.  
Irregularities such as grouping or empty spaces are caught:

```
import (
    "os"
    "another"

    "external.com/ooo" // fragmented external imports section

        "external.com/gsggs // empty space
```

### Installation:

You need to have Golang installed, otherwise follow the guide at [https://golang.org/doc/install](https://golang.org/doc/install).

```
go get github.com/fraugster/flint
```

### Usage:

Flint can be applied to three target types - files, directories and packages.

Apply to package in current directory:
```
flint [flags]
```

Apply to packages:
```
flint [flags] [packages]
```

Apply to directories. The `/...` suffix includes all sub-directories.
```
flint [flags] [directories]
```

Available flags:

```
-generated 
        # also check generated files (by the standard generated flag)
```
```
-ignore string
        # file pattern to ignore. based on https://golang.org/pkg/path/#Match
```
```
-set_exit_status
        # set exit status to 1 if any issues are found (default true)
```

### License

© Fraugster 2018, LICENSE pending.
