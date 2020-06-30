// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/umbrellala/unioffice/schema/soo/dml/diagram"
)

func TestLayoutDefHdrLstConstructor(t *testing.T) {
	v := diagram.NewLayoutDefHdrLst()
	if v == nil {
		t.Errorf("diagram.NewLayoutDefHdrLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.LayoutDefHdrLst should validate: %s", err)
	}
}

func TestLayoutDefHdrLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewLayoutDefHdrLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewLayoutDefHdrLst()
	xml.Unmarshal(buf, v2)
}
