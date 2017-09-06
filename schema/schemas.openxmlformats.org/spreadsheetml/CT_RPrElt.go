// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_RPrElt struct {
	// Font
	RFont *CT_FontName
	// Character Set
	Charset *CT_IntProperty
	// Font Family
	Family *CT_IntProperty
	// Bold
	B *CT_BooleanProperty
	// Italic
	I *CT_BooleanProperty
	// Strike Through
	Strike *CT_BooleanProperty
	// Outline
	Outline *CT_BooleanProperty
	// Shadow
	Shadow *CT_BooleanProperty
	// Condense
	Condense *CT_BooleanProperty
	// Extend
	Extend *CT_BooleanProperty
	// Text Color
	Color *CT_Color
	// Font Size
	Sz *CT_FontSize
	// Underline
	U *CT_UnderlineProperty
	// Vertical Alignment
	VertAlign *CT_VerticalAlignFontProperty
	// Font Scheme
	Scheme *CT_FontScheme
}

func NewCT_RPrElt() *CT_RPrElt {
	ret := &CT_RPrElt{}
	return ret
}

func (m *CT_RPrElt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RFont != nil {
		serFont := xml.StartElement{Name: xml.Name{Local: "x:rFont"}}
		e.EncodeElement(m.RFont, serFont)
	}
	if m.Charset != nil {
		secharset := xml.StartElement{Name: xml.Name{Local: "x:charset"}}
		e.EncodeElement(m.Charset, secharset)
	}
	if m.Family != nil {
		sefamily := xml.StartElement{Name: xml.Name{Local: "x:family"}}
		e.EncodeElement(m.Family, sefamily)
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "x:b"}}
		e.EncodeElement(m.B, seb)
	}
	if m.I != nil {
		sei := xml.StartElement{Name: xml.Name{Local: "x:i"}}
		e.EncodeElement(m.I, sei)
	}
	if m.Strike != nil {
		sestrike := xml.StartElement{Name: xml.Name{Local: "x:strike"}}
		e.EncodeElement(m.Strike, sestrike)
	}
	if m.Outline != nil {
		seoutline := xml.StartElement{Name: xml.Name{Local: "x:outline"}}
		e.EncodeElement(m.Outline, seoutline)
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "x:shadow"}}
		e.EncodeElement(m.Shadow, seshadow)
	}
	if m.Condense != nil {
		secondense := xml.StartElement{Name: xml.Name{Local: "x:condense"}}
		e.EncodeElement(m.Condense, secondense)
	}
	if m.Extend != nil {
		seextend := xml.StartElement{Name: xml.Name{Local: "x:extend"}}
		e.EncodeElement(m.Extend, seextend)
	}
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "x:color"}}
		e.EncodeElement(m.Color, secolor)
	}
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "x:sz"}}
		e.EncodeElement(m.Sz, sesz)
	}
	if m.U != nil {
		seu := xml.StartElement{Name: xml.Name{Local: "x:u"}}
		e.EncodeElement(m.U, seu)
	}
	if m.VertAlign != nil {
		severtAlign := xml.StartElement{Name: xml.Name{Local: "x:vertAlign"}}
		e.EncodeElement(m.VertAlign, severtAlign)
	}
	if m.Scheme != nil {
		sescheme := xml.StartElement{Name: xml.Name{Local: "x:scheme"}}
		e.EncodeElement(m.Scheme, sescheme)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RPrElt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_RPrElt:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rFont":
				m.RFont = NewCT_FontName()
				if err := d.DecodeElement(m.RFont, &el); err != nil {
					return err
				}
			case "charset":
				m.Charset = NewCT_IntProperty()
				if err := d.DecodeElement(m.Charset, &el); err != nil {
					return err
				}
			case "family":
				m.Family = NewCT_IntProperty()
				if err := d.DecodeElement(m.Family, &el); err != nil {
					return err
				}
			case "b":
				m.B = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.B, &el); err != nil {
					return err
				}
			case "i":
				m.I = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.I, &el); err != nil {
					return err
				}
			case "strike":
				m.Strike = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.Strike, &el); err != nil {
					return err
				}
			case "outline":
				m.Outline = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.Outline, &el); err != nil {
					return err
				}
			case "shadow":
				m.Shadow = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.Shadow, &el); err != nil {
					return err
				}
			case "condense":
				m.Condense = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.Condense, &el); err != nil {
					return err
				}
			case "extend":
				m.Extend = NewCT_BooleanProperty()
				if err := d.DecodeElement(m.Extend, &el); err != nil {
					return err
				}
			case "color":
				m.Color = NewCT_Color()
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			case "sz":
				m.Sz = NewCT_FontSize()
				if err := d.DecodeElement(m.Sz, &el); err != nil {
					return err
				}
			case "u":
				m.U = NewCT_UnderlineProperty()
				if err := d.DecodeElement(m.U, &el); err != nil {
					return err
				}
			case "vertAlign":
				m.VertAlign = NewCT_VerticalAlignFontProperty()
				if err := d.DecodeElement(m.VertAlign, &el); err != nil {
					return err
				}
			case "scheme":
				m.Scheme = NewCT_FontScheme()
				if err := d.DecodeElement(m.Scheme, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_RPrElt %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RPrElt
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RPrElt and its children
func (m *CT_RPrElt) Validate() error {
	return m.ValidateWithPath("CT_RPrElt")
}

// ValidateWithPath validates the CT_RPrElt and its children, prefixing error messages with path
func (m *CT_RPrElt) ValidateWithPath(path string) error {
	if m.RFont != nil {
		if err := m.RFont.ValidateWithPath(path + "/RFont"); err != nil {
			return err
		}
	}
	if m.Charset != nil {
		if err := m.Charset.ValidateWithPath(path + "/Charset"); err != nil {
			return err
		}
	}
	if m.Family != nil {
		if err := m.Family.ValidateWithPath(path + "/Family"); err != nil {
			return err
		}
	}
	if m.B != nil {
		if err := m.B.ValidateWithPath(path + "/B"); err != nil {
			return err
		}
	}
	if m.I != nil {
		if err := m.I.ValidateWithPath(path + "/I"); err != nil {
			return err
		}
	}
	if m.Strike != nil {
		if err := m.Strike.ValidateWithPath(path + "/Strike"); err != nil {
			return err
		}
	}
	if m.Outline != nil {
		if err := m.Outline.ValidateWithPath(path + "/Outline"); err != nil {
			return err
		}
	}
	if m.Shadow != nil {
		if err := m.Shadow.ValidateWithPath(path + "/Shadow"); err != nil {
			return err
		}
	}
	if m.Condense != nil {
		if err := m.Condense.ValidateWithPath(path + "/Condense"); err != nil {
			return err
		}
	}
	if m.Extend != nil {
		if err := m.Extend.ValidateWithPath(path + "/Extend"); err != nil {
			return err
		}
	}
	if m.Color != nil {
		if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
			return err
		}
	}
	if m.Sz != nil {
		if err := m.Sz.ValidateWithPath(path + "/Sz"); err != nil {
			return err
		}
	}
	if m.U != nil {
		if err := m.U.ValidateWithPath(path + "/U"); err != nil {
			return err
		}
	}
	if m.VertAlign != nil {
		if err := m.VertAlign.ValidateWithPath(path + "/VertAlign"); err != nil {
			return err
		}
	}
	if m.Scheme != nil {
		if err := m.Scheme.ValidateWithPath(path + "/Scheme"); err != nil {
			return err
		}
	}
	return nil
}
