// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties

import "baliance.com/gooxml"

func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_CoreProperties", NewCT_CoreProperties)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keywords", NewCT_Keywords)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keyword", NewCT_Keyword)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "coreProperties", NewCoreProperties)
}