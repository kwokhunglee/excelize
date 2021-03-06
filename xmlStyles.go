package excelize

import "encoding/xml"

// xlsxStyleSheet directly maps the stylesheet element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main - currently I have
// not checked it for completeness - it does as much as I need.
type xlsxStyleSheet struct {
	XMLName      xml.Name          `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	NumFmts      *xlsxNumFmts      `xml:"numFmts,omitempty"`
	Fonts        *xlsxFonts        `xml:"fonts,omitempty"`
	Fills        *xlsxFills        `xml:"fills,omitempty"`
	Borders      *xlsxBorders      `xml:"borders,omitempty"`
	CellStyleXfs *xlsxCellStyleXfs `xml:"cellStyleXfs,omitempty"`
	CellXfs      *xlsxCellXfs      `xml:"cellXfs,omitempty"`
	CellStyles   *xlsxCellStyles   `xml:"cellStyles,omitempty"`
	Dxfs         *xlsxDxfs         `xml:"dxfs,omitempty"`
	TableStyles  *xlsxTableStyles  `xml:"tableStyles,omitempty"`
	Colors       *xlsxStyleColors  `xml:"colors,omitempty"`
	ExtLst       *xlsxExtLst       `xml:"extLst"`
}

// xlsxAlignment formatting information pertaining to text alignment in cells.
// There are a variety of choices for how text is aligned both horizontally and
// vertically, as well as indentation settings, and so on.
type xlsxAlignment struct {
	Horizontal      string `xml:"horizontal,attr,omitempty"`
	Indent          int    `xml:"indent,attr,omitempty"`
	JustifyLastLine bool   `xml:"justifyLastLine,attr,omitempty"`
	ReadingOrder    uint64 `xml:"readingOrder,attr,omitempty"`
	RelativeIndent  int    `xml:"relativeIndent,attr,omitempty"`
	ShrinkToFit     bool   `xml:"shrinkToFit,attr,omitempty"`
	TextRotation    int    `xml:"textRotation,attr,omitempty"`
	Vertical        string `xml:"vertical,attr,omitempty"`
	WrapText        bool   `xml:"wrapText,attr,omitempty"`
}

// xlsxLine directly maps the line style element in the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main - currently I have
// not checked it for completeness - it does as much as I need.
type xlsxLine struct {
	Style string     `xml:"style,attr,omitempty"`
	Color *xlsxColor `xml:"color,omitempty"`
}

// xlsxColor is a common mapping used for both the fgColor and bgColor elements
// in the namespace http://schemas.openxmlformats.org/spreadsheetml/2006/main -
// currently I have not checked it for completeness - it does as much as I need.
type xlsxColor struct {
	RGB   string  `xml:"rgb,attr,omitempty"`
	Theme *int    `xml:"theme,attr,omitempty"`
	Tint  float64 `xml:"tint,attr,omitempty"`
}

// xlsxFonts directly maps the fonts element. This element contains all font
// definitions for this workbook.
type xlsxFonts struct {
	Count int         `xml:"count,attr"`
	Font  []*xlsxFont `xml:"font,omitempty"`
}

// xlsxFont directly maps the font element. This element defines the properties
// for one of the fonts used in this workbook.
type xlsxFont struct {
	Font string `xml:",innerxml"`
}

// xlsxFills directly maps the fills element. This element defines the cell
// fills portion of the Styles part, consisting of a sequence of fill records. A
// cell fill consists of a background color, foreground color, and pattern to be
// applied across the cell.
type xlsxFills struct {
	Count int        `xml:"count,attr"`
	Fill  []xlsxFill `xml:"fill,omitempty"`
}

// xlsxFill directly maps the fill element. This element specifies fill
// formatting.
type xlsxFill struct {
	Fill string `xml:",innerxml"`
}

// xlsxBorders directly maps the borders element. This element contains borders
// formatting information, specifying all border definitions for all cells in
// the workbook.
type xlsxBorders struct {
	Count  int           `xml:"count,attr"`
	Border []*xlsxBorder `xml:"border,omitempty"`
}

// xlsxBorder directly maps the border element. Expresses a single set of cell
// border formats (left, right, top, bottom, diagonal). Color is optional. When
// missing, 'automatic' is implied.
type xlsxBorder struct {
	DiagonalDown bool     `xml:"diagonalDown,attr,omitempty"`
	DiagonalUp   bool     `xml:"diagonalUp,attr,omitempty"`
	Outline      bool     `xml:"outline,attr,omitempty"`
	Left         xlsxLine `xml:"left,omitempty"`
	Right        xlsxLine `xml:"right,omitempty"`
	Top          xlsxLine `xml:"top,omitempty"`
	Bottom       xlsxLine `xml:"bottom,omitempty"`
	Diagonal     xlsxLine `xml:"diagonal,omitempty"`
}

// xlsxCellStyles directly maps the cellStyles element. This element contains
// the named cell styles, consisting of a sequence of named style records. A
// named cell style is a collection of direct or themed formatting (e.g., cell
// border, cell fill, and font type/size/style) grouped together into a single
// named style, and can be applied to a cell.
type xlsxCellStyles struct {
	XMLName   xml.Name         `xml:"cellStyles"`
	Count     int              `xml:"count,attr"`
	CellStyle []*xlsxCellStyle `xml:"cellStyle,omitempty"`
}

// xlsxCellStyle directly maps the cellStyle element. This element represents
// the name and related formatting records for a named cell style in this
// workbook.
type xlsxCellStyle struct {
	XMLName       xml.Name `xml:"cellStyle"`
	BuiltInID     *int     `xml:"builtInId,attr,omitempty"`
	CustomBuiltIn *bool    `xml:"customBuiltIn,attr,omitempty"`
	Hidden        *bool    `xml:"hidden,attr,omitempty"`
	ILevel        *bool    `xml:"iLevel,attr,omitempty"`
	Name          string   `xml:"name,attr"`
	XfID          int      `xml:"xfId,attr"`
}

// xlsxCellStyleXfs directly maps the cellStyleXfs element. This element
// contains the master formatting records (xf's) which define the formatting for
// all named cell styles in this workbook. Master formatting records reference
// individual elements of formatting (e.g., number format, font definitions,
// cell fills, etc) by specifying a zero-based index into those collections.
// Master formatting records also specify whether to apply or ignore particular
// aspects of formatting.
type xlsxCellStyleXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxXf directly maps the xf element. A single xf element describes all of the
// formatting for a cell.
type xlsxXf struct {
	ApplyAlignment    bool           `xml:"applyAlignment,attr"`
	ApplyBorder       bool           `xml:"applyBorder,attr"`
	ApplyFill         bool           `xml:"applyFill,attr"`
	ApplyFont         bool           `xml:"applyFont,attr"`
	ApplyNumberFormat bool           `xml:"applyNumberFormat,attr"`
	ApplyProtection   bool           `xml:"applyProtection,attr"`
	BorderID          int            `xml:"borderId,attr"`
	FillID            int            `xml:"fillId,attr"`
	FontID            int            `xml:"fontId,attr"`
	NumFmtID          int            `xml:"numFmtId,attr"`
	PivotButton       bool           `xml:"pivotButton,attr,omitempty"`
	QuotePrefix       bool           `xml:"quotePrefix,attr,omitempty"`
	XfID              *int           `xml:"xfId,attr,omitempty"`
	Alignment         *xlsxAlignment `xml:"alignment"`
}

// xlsxCellXfs directly maps the cellXfs element. This element contains the
// master formatting records (xf) which define the formatting applied to cells
// in this workbook. These records are the starting point for determining the
// formatting for a cell. Cells in the Sheet Part reference the xf records by
// zero-based index.
type xlsxCellXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxDxfs directly maps the dxfs element. This element contains the master
// differential formatting records (dxf's) which define formatting for all non-
// cell formatting in this workbook. Whereas xf records fully specify a
// particular aspect of formatting (e.g., cell borders) by referencing those
// formatting definitions elsewhere in the Styles part, dxf records specify
// incremental (or differential) aspects of formatting directly inline within
// the dxf element. The dxf formatting is to be applied on top of or in addition
// to any formatting already present on the object using the dxf record.
type xlsxDxfs struct {
	Count int        `xml:"count,attr"`
	Dxfs  []*xlsxDxf `xml:"dxf,omitempty"`
}

// xlsxDxf directly maps the dxf element. A single dxf record, expressing
// incremental formatting to be applied.
type xlsxDxf struct {
	Dxf string `xml:",innerxml"`
}

// xlsxTableStyles directly maps the tableStyles element. This element
// represents a collection of Table style definitions for Table styles and
// PivotTable styles used in this workbook. It consists of a sequence of
// tableStyle records, each defining a single Table style.
type xlsxTableStyles struct {
	Count             int               `xml:"count,attr"`
	DefaultPivotStyle string            `xml:"defaultPivotStyle,attr"`
	DefaultTableStyle string            `xml:"defaultTableStyle,attr"`
	TableStyles       []*xlsxTableStyle `xml:"tableStyle,omitempty"`
}

// xlsxTableStyle directly maps the tableStyle element. This element represents
// a single table style definition that indicates how a spreadsheet application
// should format and display a table.
type xlsxTableStyle struct {
	Name              string `xml:"name,attr,omitempty"`
	Pivot             int    `xml:"pivot,attr,omitempty"`
	Count             int    `xml:"count,attr,omitempty"`
	Table             bool   `xml:"table,attr,omitempty"`
	TableStyleElement string `xml:",innerxml"`
}

// xlsxNumFmts directly maps the numFmts element. This element defines the
// number formats in this workbook, consisting of a sequence of numFmt records,
// where each numFmt record defines a particular number format, indicating how
// to format and render the numeric value of a cell.
type xlsxNumFmts struct {
	Count  int           `xml:"count,attr"`
	NumFmt []*xlsxNumFmt `xml:"numFmt,omitempty"`
}

// xlsxNumFmt directly maps the numFmt element. This element specifies number
// format properties which indicate how to format and render the numeric value
// of a cell.
type xlsxNumFmt struct {
	NumFmtID   int    `xml:"numFmtId,attr,omitempty"`
	FormatCode string `xml:"formatCode,attr,omitempty"`
}

// xlsxStyleColors directly maps the colors element. Color information
// associated with this stylesheet. This collection is written whenever the
// legacy color palette has been modified (backwards compatibility settings) or
// a custom color has been selected while using this workbook.
type xlsxStyleColors struct {
	Color string `xml:",innerxml"`
}

// formatBorder directly maps the format settings of the borders.
type formatBorder struct {
	Border []struct {
		Type  string `json:"type"`
		Color string `json:"color"`
		Style int    `json:"style"`
	} `json:"border"`
}
