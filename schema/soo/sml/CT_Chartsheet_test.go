// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/umbrellala/unioffice/schema/soo/sml"
)

func TestCT_ChartsheetConstructor(t *testing.T) {
	v := sml.NewCT_Chartsheet()
	if v == nil {
		t.Errorf("sml.NewCT_Chartsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Chartsheet should validate: %s", err)
	}
}

func TestCT_ChartsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Chartsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Chartsheet()
	xml.Unmarshal(buf, v2)
}
