// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Ratio struct {
	NAttr int64
	DAttr int64
}

func NewCT_Ratio() *CT_Ratio {
	ret := &CT_Ratio{}
	return ret
}
func (m *CT_Ratio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "d"},
		Value: fmt.Sprintf("%v", m.DAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Ratio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.NAttr = parsed
		}
		if attr.Name.Local == "d" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DAttr = parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Ratio: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_Ratio) Validate() error {
	return m.ValidateWithPath("CT_Ratio")
}
func (m *CT_Ratio) ValidateWithPath(path string) error {
	return nil
}