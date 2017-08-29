// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

// ST_TLTimeAnimateValueTime is a union type
type ST_TLTimeAnimateValueTime struct {
	ST_PositiveFixedPercentage *drawingml.ST_PositiveFixedPercentage
	ST_TLTimeIndefinite        ST_TLTimeIndefinite
}

func (m *ST_TLTimeAnimateValueTime) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_TLTimeAnimateValueTime) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_PositiveFixedPercentage != nil {
		if err := m.ST_PositiveFixedPercentage.ValidateWithPath(path + "/ST_PositiveFixedPercentage"); err != nil {
			return err
		}
		mems = append(mems, "ST_PositiveFixedPercentage")
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		mems = append(mems, "ST_TLTimeIndefinite")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_TLTimeAnimateValueTime) String() string {
	if m.ST_PositiveFixedPercentage != nil {
		return m.ST_PositiveFixedPercentage.String()
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		return m.ST_TLTimeIndefinite.String()
	}
	return ""
}