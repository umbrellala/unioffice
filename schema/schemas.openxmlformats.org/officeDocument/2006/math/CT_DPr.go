// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type CT_DPr struct {
	BegChr *CT_Char
	SepChr *CT_Char
	EndChr *CT_Char
	Grow   *CT_OnOff
	Shp    *CT_Shp
	CtrlPr *CT_CtrlPr
}

func NewCT_DPr() *CT_DPr {
	ret := &CT_DPr{}
	return ret
}
func (m *CT_DPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.BegChr != nil {
		sebegChr := xml.StartElement{Name: xml.Name{Local: "m:begChr"}}
		e.EncodeElement(m.BegChr, sebegChr)
	}
	if m.SepChr != nil {
		sesepChr := xml.StartElement{Name: xml.Name{Local: "m:sepChr"}}
		e.EncodeElement(m.SepChr, sesepChr)
	}
	if m.EndChr != nil {
		seendChr := xml.StartElement{Name: xml.Name{Local: "m:endChr"}}
		e.EncodeElement(m.EndChr, seendChr)
	}
	if m.Grow != nil {
		segrow := xml.StartElement{Name: xml.Name{Local: "m:grow"}}
		e.EncodeElement(m.Grow, segrow)
	}
	if m.Shp != nil {
		seshp := xml.StartElement{Name: xml.Name{Local: "m:shp"}}
		e.EncodeElement(m.Shp, seshp)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "begChr":
				m.BegChr = NewCT_Char()
				if err := d.DecodeElement(m.BegChr, &el); err != nil {
					return err
				}
			case "sepChr":
				m.SepChr = NewCT_Char()
				if err := d.DecodeElement(m.SepChr, &el); err != nil {
					return err
				}
			case "endChr":
				m.EndChr = NewCT_Char()
				if err := d.DecodeElement(m.EndChr, &el); err != nil {
					return err
				}
			case "grow":
				m.Grow = NewCT_OnOff()
				if err := d.DecodeElement(m.Grow, &el); err != nil {
					return err
				}
			case "shp":
				m.Shp = NewCT_Shp()
				if err := d.DecodeElement(m.Shp, &el); err != nil {
					return err
				}
			case "ctrlPr":
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DPr
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DPr) Validate() error {
	return m.ValidateWithPath("CT_DPr")
}
func (m *CT_DPr) ValidateWithPath(path string) error {
	if m.BegChr != nil {
		if err := m.BegChr.ValidateWithPath(path + "/BegChr"); err != nil {
			return err
		}
	}
	if m.SepChr != nil {
		if err := m.SepChr.ValidateWithPath(path + "/SepChr"); err != nil {
			return err
		}
	}
	if m.EndChr != nil {
		if err := m.EndChr.ValidateWithPath(path + "/EndChr"); err != nil {
			return err
		}
	}
	if m.Grow != nil {
		if err := m.Grow.ValidateWithPath(path + "/Grow"); err != nil {
			return err
		}
	}
	if m.Shp != nil {
		if err := m.Shp.ValidateWithPath(path + "/Shp"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}