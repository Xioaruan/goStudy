// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ifaceassert_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, ifaceassert.Analyzer, "a", "typeparams")
}
