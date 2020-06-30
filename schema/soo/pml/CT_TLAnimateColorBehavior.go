// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/umbrellala/unioffice"
	"github.com/umbrellala/unioffice/schema/soo/dml"
)

type CT_TLAnimateColorBehavior struct {
	// Color Space
	ClrSpcAttr ST_TLAnimateColorSpace
	// Direction
	DirAttr ST_TLAnimateColorDirection
	CBhvr   *CT_TLCommonBehaviorData
	// By
	By *CT_TLByAnimateColorTransform
	// From
	From *dml.CT_Color
	// To
	To *dml.CT_Color
}

func NewCT_TLAnimateColorBehavior() *CT_TLAnimateColorBehavior {
	ret := &CT_TLAnimateColorBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLAnimateColorBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ClrSpcAttr != ST_TLAnimateColorSpaceUnset {
		attr, err := m.ClrSpcAttr.MarshalXMLAttr(xml.Name{Local: "clrSpc"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DirAttr != ST_TLAnimateColorDirectionUnset {
		attr, err := m.DirAttr.MarshalXMLAttr(xml.Name{Local: "dir"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.By != nil {
		seby := xml.StartElement{Name: xml.Name{Local: "p:by"}}
		e.EncodeElement(m.By, seby)
	}
	if m.From != nil {
		sefrom := xml.StartElement{Name: xml.Name{Local: "p:from"}}
		e.EncodeElement(m.From, sefrom)
	}
	if m.To != nil {
		seto := xml.StartElement{Name: xml.Name{Local: "p:to"}}
		e.EncodeElement(m.To, seto)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimateColorBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "clrSpc" {
			m.ClrSpcAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_TLAnimateColorBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cBhvr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cBhvr"}:
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "by"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "by"}:
				m.By = NewCT_TLByAnimateColorTransform()
				if err := d.DecodeElement(m.By, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "from"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "from"}:
				m.From = dml.NewCT_Color()
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "to"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "to"}:
				m.To = dml.NewCT_Color()
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLAnimateColorBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateColorBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimateColorBehavior and its children
func (m *CT_TLAnimateColorBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateColorBehavior")
}

// ValidateWithPath validates the CT_TLAnimateColorBehavior and its children, prefixing error messages with path
func (m *CT_TLAnimateColorBehavior) ValidateWithPath(path string) error {
	if err := m.ClrSpcAttr.ValidateWithPath(path + "/ClrSpcAttr"); err != nil {
		return err
	}
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.By != nil {
		if err := m.By.ValidateWithPath(path + "/By"); err != nil {
			return err
		}
	}
	if m.From != nil {
		if err := m.From.ValidateWithPath(path + "/From"); err != nil {
			return err
		}
	}
	if m.To != nil {
		if err := m.To.ValidateWithPath(path + "/To"); err != nil {
			return err
		}
	}
	return nil
}
