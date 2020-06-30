// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"

	"github.com/umbrellala/unioffice"
)

type OfcEquationxml struct {
	OfcCT_EquationXml
}

func NewOfcEquationxml() *OfcEquationxml {
	ret := &OfcEquationxml{}
	ret.OfcCT_EquationXml = *NewOfcCT_EquationXml()
	return ret
}

func (m *OfcEquationxml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:equationxml"
	return m.OfcCT_EquationXml.MarshalXML(e, start)
}

func (m *OfcEquationxml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_EquationXml = *NewOfcCT_EquationXml()
	for _, attr := range start.Attr {
		if attr.Name.Local == "contentType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ContentTypeAttr = &parsed
			continue
		}
	}
lOfcEquationxml:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				if anyEl, err := unioffice.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lOfcEquationxml
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcEquationxml and its children
func (m *OfcEquationxml) Validate() error {
	return m.ValidateWithPath("OfcEquationxml")
}

// ValidateWithPath validates the OfcEquationxml and its children, prefixing error messages with path
func (m *OfcEquationxml) ValidateWithPath(path string) error {
	if err := m.OfcCT_EquationXml.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
