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
	"strconv"

	"github.com/umbrellala/unioffice/schema/soo/ofc/sharedTypes"
)

type AG_AllCoreAttributes struct {
	HrefAttr              *string
	TargetAttr            *string
	ClassAttr             *string
	TitleAttr             *string
	AltAttr               *string
	CoordsizeAttr         *string
	CoordoriginAttr       *string
	WrapcoordsAttr        *string
	PrintAttr             sharedTypes.ST_TrueFalse
	IdAttr                *string
	StyleAttr             *string
	SpidAttr              *string
	OnedAttr              sharedTypes.ST_TrueFalse
	RegroupidAttr         *int64
	DoubleclicknotifyAttr sharedTypes.ST_TrueFalse
	ButtonAttr            sharedTypes.ST_TrueFalse
	UserhiddenAttr        sharedTypes.ST_TrueFalse
	BulletAttr            sharedTypes.ST_TrueFalse
	HrAttr                sharedTypes.ST_TrueFalse
	HrstdAttr             sharedTypes.ST_TrueFalse
	HrnoshadeAttr         sharedTypes.ST_TrueFalse
	HrpctAttr             *float32
	HralignAttr           OfcST_HrAlign
	AllowincellAttr       sharedTypes.ST_TrueFalse
	AllowoverlapAttr      sharedTypes.ST_TrueFalse
	UserdrawnAttr         sharedTypes.ST_TrueFalse
	BordertopcolorAttr    *string
	BorderleftcolorAttr   *string
	BorderbottomcolorAttr *string
	BorderrightcolorAttr  *string
	DgmlayoutAttr         OfcST_DiagramLayout
	DgmnodekindAttr       *int64
	DgmlayoutmruAttr      OfcST_DiagramLayout
	InsetmodeAttr         OfcST_InsetMode
}

func NewAG_AllCoreAttributes() *AG_AllCoreAttributes {
	ret := &AG_AllCoreAttributes{}
	return ret
}

func (m *AG_AllCoreAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.ClassAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "class"},
			Value: fmt.Sprintf("%v", *m.ClassAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.AltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alt"},
			Value: fmt.Sprintf("%v", *m.AltAttr)})
	}
	if m.CoordsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordsize"},
			Value: fmt.Sprintf("%v", *m.CoordsizeAttr)})
	}
	if m.CoordoriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordorigin"},
			Value: fmt.Sprintf("%v", *m.CoordoriginAttr)})
	}
	if m.WrapcoordsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wrapcoords"},
			Value: fmt.Sprintf("%v", *m.WrapcoordsAttr)})
	}
	if m.PrintAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PrintAttr.MarshalXMLAttr(xml.Name{Local: "print"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.OnedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnedAttr.MarshalXMLAttr(xml.Name{Local: "oned"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RegroupidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:regroupid"},
			Value: fmt.Sprintf("%v", *m.RegroupidAttr)})
	}
	if m.DoubleclicknotifyAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.DoubleclicknotifyAttr.MarshalXMLAttr(xml.Name{Local: "doubleclicknotify"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ButtonAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ButtonAttr.MarshalXMLAttr(xml.Name{Local: "button"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserhiddenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserhiddenAttr.MarshalXMLAttr(xml.Name{Local: "userhidden"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BulletAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.BulletAttr.MarshalXMLAttr(xml.Name{Local: "bullet"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrAttr.MarshalXMLAttr(xml.Name{Local: "hr"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrstdAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrstdAttr.MarshalXMLAttr(xml.Name{Local: "hrstd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrnoshadeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrnoshadeAttr.MarshalXMLAttr(xml.Name{Local: "hrnoshade"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrpctAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:hrpct"},
			Value: fmt.Sprintf("%v", *m.HrpctAttr)})
	}
	if m.HralignAttr != OfcST_HrAlignUnset {
		attr, err := m.HralignAttr.MarshalXMLAttr(xml.Name{Local: "hralign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowincellAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowincellAttr.MarshalXMLAttr(xml.Name{Local: "allowincell"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowoverlapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowoverlapAttr.MarshalXMLAttr(xml.Name{Local: "allowoverlap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserdrawnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserdrawnAttr.MarshalXMLAttr(xml.Name{Local: "userdrawn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BordertopcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:bordertopcolor"},
			Value: fmt.Sprintf("%v", *m.BordertopcolorAttr)})
	}
	if m.BorderleftcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderleftcolor"},
			Value: fmt.Sprintf("%v", *m.BorderleftcolorAttr)})
	}
	if m.BorderbottomcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderbottomcolor"},
			Value: fmt.Sprintf("%v", *m.BorderbottomcolorAttr)})
	}
	if m.BorderrightcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderrightcolor"},
			Value: fmt.Sprintf("%v", *m.BorderrightcolorAttr)})
	}
	if m.DgmlayoutAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayout"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DgmnodekindAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:dgmnodekind"},
			Value: fmt.Sprintf("%v", *m.DgmnodekindAttr)})
	}
	if m.DgmlayoutmruAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutmruAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayoutmru"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InsetmodeAttr != OfcST_InsetModeUnset {
		attr, err := m.InsetmodeAttr.MarshalXMLAttr(xml.Name{Local: "insetmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	return nil
}

func (m *AG_AllCoreAttributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bullet" {
			m.BulletAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oned" {
			m.OnedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "regroupid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RegroupidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmnodekind" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmnodekindAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "doubleclicknotify" {
			m.DoubleclicknotifyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderrightcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderrightcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "button" {
			m.ButtonAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderleftcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderleftcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userhidden" {
			m.UserhiddenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userdrawn" {
			m.UserdrawnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayoutmru" {
			m.DgmlayoutmruAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayout" {
			m.DgmlayoutAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderbottomcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderbottomcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bordertopcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BordertopcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowoverlap" {
			m.AllowoverlapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hr" {
			m.HrAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrstd" {
			m.HrstdAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrnoshade" {
			m.HrnoshadeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrpct" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.HrpctAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hralign" {
			m.HralignAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			m.PrintAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "wrapcoords" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WrapcoordsAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordoriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordsize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "alt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltAttr = &parsed
			continue
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Local == "class" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ClassAttr = &parsed
			continue
		}
		if attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_AllCoreAttributes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_AllCoreAttributes and its children
func (m *AG_AllCoreAttributes) Validate() error {
	return m.ValidateWithPath("AG_AllCoreAttributes")
}

// ValidateWithPath validates the AG_AllCoreAttributes and its children, prefixing error messages with path
func (m *AG_AllCoreAttributes) ValidateWithPath(path string) error {
	if err := m.PrintAttr.ValidateWithPath(path + "/PrintAttr"); err != nil {
		return err
	}
	if err := m.OnedAttr.ValidateWithPath(path + "/OnedAttr"); err != nil {
		return err
	}
	if err := m.DoubleclicknotifyAttr.ValidateWithPath(path + "/DoubleclicknotifyAttr"); err != nil {
		return err
	}
	if err := m.ButtonAttr.ValidateWithPath(path + "/ButtonAttr"); err != nil {
		return err
	}
	if err := m.UserhiddenAttr.ValidateWithPath(path + "/UserhiddenAttr"); err != nil {
		return err
	}
	if err := m.BulletAttr.ValidateWithPath(path + "/BulletAttr"); err != nil {
		return err
	}
	if err := m.HrAttr.ValidateWithPath(path + "/HrAttr"); err != nil {
		return err
	}
	if err := m.HrstdAttr.ValidateWithPath(path + "/HrstdAttr"); err != nil {
		return err
	}
	if err := m.HrnoshadeAttr.ValidateWithPath(path + "/HrnoshadeAttr"); err != nil {
		return err
	}
	if err := m.HralignAttr.ValidateWithPath(path + "/HralignAttr"); err != nil {
		return err
	}
	if err := m.AllowincellAttr.ValidateWithPath(path + "/AllowincellAttr"); err != nil {
		return err
	}
	if err := m.AllowoverlapAttr.ValidateWithPath(path + "/AllowoverlapAttr"); err != nil {
		return err
	}
	if err := m.UserdrawnAttr.ValidateWithPath(path + "/UserdrawnAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutAttr.ValidateWithPath(path + "/DgmlayoutAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutmruAttr.ValidateWithPath(path + "/DgmlayoutmruAttr"); err != nil {
		return err
	}
	if err := m.InsetmodeAttr.ValidateWithPath(path + "/InsetmodeAttr"); err != nil {
		return err
	}
	return nil
}
