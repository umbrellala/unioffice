// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
)

type CT_RelIds struct {
	DmAttr string
	LoAttr string
	QsAttr string
	CsAttr string
}

func NewCT_RelIds() *CT_RelIds {
	ret := &CT_RelIds{}
	return ret
}
func (m *CT_RelIds) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:dm"},
		Value: fmt.Sprintf("%v", m.DmAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:lo"},
		Value: fmt.Sprintf("%v", m.LoAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:qs"},
		Value: fmt.Sprintf("%v", m.QsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:cs"},
		Value: fmt.Sprintf("%v", m.CsAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_RelIds) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "dm" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DmAttr = parsed
		}
		if attr.Name.Local == "lo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LoAttr = parsed
		}
		if attr.Name.Local == "qs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.QsAttr = parsed
		}
		if attr.Name.Local == "cs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CsAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_RelIds: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_RelIds) Validate() error {
	return m.ValidateWithPath("CT_RelIds")
}
func (m *CT_RelIds) ValidateWithPath(path string) error {
	return nil
}