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
	"fmt"

	"github.com/umbrellala/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_H struct {
	PositionAttr    *string
	PolarAttr       *string
	MapAttr         *string
	InvxAttr        sharedTypes.ST_TrueFalse
	InvyAttr        sharedTypes.ST_TrueFalse
	SwitchAttr      sharedTypes.ST_TrueFalseBlank
	XrangeAttr      *string
	YrangeAttr      *string
	RadiusrangeAttr *string
}

func NewCT_H() *CT_H {
	ret := &CT_H{}
	return ret
}

func (m *CT_H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "position"},
			Value: fmt.Sprintf("%v", *m.PositionAttr)})
	}
	if m.PolarAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "polar"},
			Value: fmt.Sprintf("%v", *m.PolarAttr)})
	}
	if m.MapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "map"},
			Value: fmt.Sprintf("%v", *m.MapAttr)})
	}
	if m.InvxAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InvxAttr.MarshalXMLAttr(xml.Name{Local: "invx"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InvyAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InvyAttr.MarshalXMLAttr(xml.Name{Local: "invy"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SwitchAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.SwitchAttr.MarshalXMLAttr(xml.Name{Local: "switch"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.XrangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xrange"},
			Value: fmt.Sprintf("%v", *m.XrangeAttr)})
	}
	if m.YrangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "yrange"},
			Value: fmt.Sprintf("%v", *m.YrangeAttr)})
	}
	if m.RadiusrangeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "radiusrange"},
			Value: fmt.Sprintf("%v", *m.RadiusrangeAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_H) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "position" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PositionAttr = &parsed
			continue
		}
		if attr.Name.Local == "polar" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PolarAttr = &parsed
			continue
		}
		if attr.Name.Local == "map" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MapAttr = &parsed
			continue
		}
		if attr.Name.Local == "invx" {
			m.InvxAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "invy" {
			m.InvyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "switch" {
			m.SwitchAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "xrange" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XrangeAttr = &parsed
			continue
		}
		if attr.Name.Local == "yrange" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.YrangeAttr = &parsed
			continue
		}
		if attr.Name.Local == "radiusrange" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RadiusrangeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_H: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_H and its children
func (m *CT_H) Validate() error {
	return m.ValidateWithPath("CT_H")
}

// ValidateWithPath validates the CT_H and its children, prefixing error messages with path
func (m *CT_H) ValidateWithPath(path string) error {
	if err := m.InvxAttr.ValidateWithPath(path + "/InvxAttr"); err != nil {
		return err
	}
	if err := m.InvyAttr.ValidateWithPath(path + "/InvyAttr"); err != nil {
		return err
	}
	if err := m.SwitchAttr.ValidateWithPath(path + "/SwitchAttr"); err != nil {
		return err
	}
	return nil
}
