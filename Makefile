# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=hello_test
GOFILES=\
	dummy.go\
	hello_test.go\

include $(GOROOT)/src/Make.pkg
