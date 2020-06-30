// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/umbrellala/unioffice/schema/soo/ofc/custom_properties"
)

func TestPropertiesConstructor(t *testing.T) {
	v := custom_properties.NewProperties()
	if v == nil {
		t.Errorf("custom_properties.NewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.Properties should validate: %s", err)
	}
}

func TestPropertiesMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewProperties()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewProperties()
	xml.Unmarshal(buf, v2)
}
