// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/umbrellala/unioffice/schema/soo/pml"
)

func TestCmLstConstructor(t *testing.T) {
	v := pml.NewCmLst()
	if v == nil {
		t.Errorf("pml.NewCmLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CmLst should validate: %s", err)
	}
}

func TestCmLstMarshalUnmarshal(t *testing.T) {
	v := pml.NewCmLst()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCmLst()
	xml.Unmarshal(buf, v2)
}
