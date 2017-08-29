// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_Scenarios struct {
	// Current Scenario
	CurrentAttr *uint32
	// Last Shown Scenario
	ShowAttr *uint32
	// Sequence of References
	SqrefAttr *ST_Sqref
	// Scenario
	Scenario []*CT_Scenario
}

func NewCT_Scenarios() *CT_Scenarios {
	ret := &CT_Scenarios{}
	return ret
}
func (m *CT_Scenarios) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.CurrentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "current"},
			Value: fmt.Sprintf("%v", *m.CurrentAttr)})
	}
	if m.ShowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "show"},
			Value: fmt.Sprintf("%v", *m.ShowAttr)})
	}
	if m.SqrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
			Value: fmt.Sprintf("%v", *m.SqrefAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sescenario := xml.StartElement{Name: xml.Name{Local: "x:scenario"}}
	e.EncodeElement(m.Scenario, sescenario)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Scenarios) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "current" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CurrentAttr = &pt
		}
		if attr.Name.Local == "show" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.ShowAttr = &pt
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = &parsed
		}
	}
lCT_Scenarios:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "scenario":
				tmp := NewCT_Scenario()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Scenario = append(m.Scenario, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scenarios
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Scenarios) Validate() error {
	return m.ValidateWithPath("CT_Scenarios")
}
func (m *CT_Scenarios) ValidateWithPath(path string) error {
	for i, v := range m.Scenario {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Scenario[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}