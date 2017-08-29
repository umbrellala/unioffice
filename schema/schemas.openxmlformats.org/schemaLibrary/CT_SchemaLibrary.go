// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package schemaLibrary

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_SchemaLibrary struct {
	Schema []*CT_Schema
}

func NewCT_SchemaLibrary() *CT_SchemaLibrary {
	ret := &CT_SchemaLibrary{}
	return ret
}
func (m *CT_SchemaLibrary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Schema != nil {
		seschema := xml.StartElement{Name: xml.Name{Local: "ma:schema"}}
		e.EncodeElement(m.Schema, seschema)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SchemaLibrary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SchemaLibrary:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "schema":
				tmp := NewCT_Schema()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Schema = append(m.Schema, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SchemaLibrary
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SchemaLibrary) Validate() error {
	return m.ValidateWithPath("CT_SchemaLibrary")
}
func (m *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for i, v := range m.Schema {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Schema[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}