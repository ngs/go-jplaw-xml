package jplaw

import (
	"encoding/xml"
	"fmt"
)

// Law represents the root element of a Japanese law document
type Law struct {
	XMLName         xml.Name `xml:"Law"`
	Era             Era      `xml:"Era,attr"`
	Year            int      `xml:"Year,attr"`
	Num             string   `xml:"Num,attr"`
	PromulgateMonth int      `xml:"PromulgateMonth,attr,omitempty"`
	PromulgateDay   int      `xml:"PromulgateDay,attr,omitempty"`
	LawType         LawType  `xml:"LawType,attr"`
	Lang            Language `xml:"Lang,attr"`
	LawNum          string   `xml:"LawNum"`
	LawBody         LawBody  `xml:"LawBody"`
}

// LawBody represents the body of a law document
type LawBody struct {
	Subject         string            `xml:"Subject,attr,omitempty"`
	LawTitle        *LawTitle         `xml:"LawTitle,omitempty"`
	EnactStatement  []EnactStatement  `xml:"EnactStatement,omitempty"`
	TOC             *TOC              `xml:"TOC,omitempty"`
	Preamble        *Preamble         `xml:"Preamble,omitempty"`
	MainProvision   MainProvision     `xml:"MainProvision"`
	SupplProvision  []SupplProvision  `xml:"SupplProvision,omitempty"`
	AppdxTable      []AppdxTable      `xml:"AppdxTable,omitempty"`
	AppdxNote       []AppdxNote       `xml:"AppdxNote,omitempty"`
	AppdxStyle      []AppdxStyle      `xml:"AppdxStyle,omitempty"`
	Appdx           []Appdx           `xml:"Appdx,omitempty"`
	AppdxFig        []AppdxFig        `xml:"AppdxFig,omitempty"`
	AppdxFormat     []AppdxFormat     `xml:"AppdxFormat,omitempty"`
}

// LawTitle represents the title of a law
type LawTitle struct {
	Kana       string  `xml:"Kana,attr,omitempty"`
	Abbrev     string  `xml:"Abbrev,attr,omitempty"`
	AbbrevKana string  `xml:"AbbrevKana,attr,omitempty"`
	Content    string  `xml:",chardata"`
	Line       []Line  `xml:"Line,omitempty"`
	Ruby       []Ruby  `xml:"Ruby,omitempty"`
	Sup        []Sup   `xml:"Sup,omitempty"`
	Sub        []Sub   `xml:"Sub,omitempty"`
}

// EnactStatement represents an enactment statement
type EnactStatement struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// TOC represents the table of contents
type TOC struct {
	TOCLabel           *TOCLabel           `xml:"TOCLabel,omitempty"`
	TOCPreambleLabel   *TOCPreambleLabel   `xml:"TOCPreambleLabel,omitempty"`
	TOCPart            []TOCPart           `xml:"TOCPart,omitempty"`
	TOCChapter         []TOCChapter        `xml:"TOCChapter,omitempty"`
	TOCSection         []TOCSection        `xml:"TOCSection,omitempty"`
	TOCArticle         []TOCArticle        `xml:"TOCArticle,omitempty"`
	TOCSupplProvision  *TOCSupplProvision  `xml:"TOCSupplProvision,omitempty"`
	TOCAppdxTableLabel []TOCAppdxTableLabel `xml:"TOCAppdxTableLabel,omitempty"`
}

// TOCLabel represents a TOC label
type TOCLabel struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// TOCPreambleLabel represents a TOC preamble label
type TOCPreambleLabel struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// TOCPart represents a part in the table of contents
type TOCPart struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	PartTitle    PartTitle     `xml:"PartTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange,omitempty"`
	TOCChapter   []TOCChapter  `xml:"TOCChapter,omitempty"`
}

// TOCChapter represents a chapter in the table of contents
type TOCChapter struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	ChapterTitle ChapterTitle  `xml:"ChapterTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange,omitempty"`
	TOCSection   []TOCSection  `xml:"TOCSection,omitempty"`
}

// TOCSection represents a section in the table of contents
type TOCSection struct {
	Num          string          `xml:"Num,attr"`
	Delete       bool            `xml:"Delete,attr,omitempty"`
	SectionTitle SectionTitle    `xml:"SectionTitle"`
	ArticleRange *ArticleRange   `xml:"ArticleRange,omitempty"`
	TOCSubsection []TOCSubsection `xml:"TOCSubsection,omitempty"`
	TOCDivision  []TOCDivision   `xml:"TOCDivision,omitempty"`
}

// TOCSubsection represents a subsection in the table of contents
type TOCSubsection struct {
	Num             string          `xml:"Num,attr"`
	Delete          bool            `xml:"Delete,attr,omitempty"`
	SubsectionTitle SubsectionTitle `xml:"SubsectionTitle"`
	ArticleRange    *ArticleRange   `xml:"ArticleRange,omitempty"`
	TOCDivision     []TOCDivision   `xml:"TOCDivision,omitempty"`
}

// TOCDivision represents a division in the table of contents
type TOCDivision struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	DivisionTitle DivisionTitle `xml:"DivisionTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange,omitempty"`
}

// TOCArticle represents an article in the table of contents
type TOCArticle struct {
	Num            string         `xml:"Num,attr"`
	Delete         bool           `xml:"Delete,attr,omitempty"`
	ArticleTitle   ArticleTitle   `xml:"ArticleTitle"`
	ArticleCaption ArticleCaption `xml:"ArticleCaption"`
}

// TOCSupplProvision represents supplementary provisions in TOC
type TOCSupplProvision struct {
	SupplProvisionLabel SupplProvisionLabel `xml:"SupplProvisionLabel"`
	ArticleRange        *ArticleRange       `xml:"ArticleRange,omitempty"`
	TOCArticle          []TOCArticle        `xml:"TOCArticle,omitempty"`
	TOCChapter          []TOCChapter        `xml:"TOCChapter,omitempty"`
}

// TOCAppdxTableLabel represents an appendix table label in TOC
type TOCAppdxTableLabel struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ArticleRange represents an article range
type ArticleRange struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Preamble represents the preamble section
type Preamble struct {
	Paragraph []Paragraph `xml:"Paragraph"`
}

// MainProvision represents the main provision
type MainProvision struct {
	Extract   bool        `xml:"Extract,attr,omitempty"`
	Part      []Part      `xml:"Part,omitempty"`
	Chapter   []Chapter   `xml:"Chapter,omitempty"`
	Section   []Section   `xml:"Section,omitempty"`
	Article   []Article   `xml:"Article,omitempty"`
	Paragraph []Paragraph `xml:"Paragraph,omitempty"`
}

// Part represents a part in the law structure
type Part struct {
	Num      string    `xml:"Num,attr"`
	Delete   bool      `xml:"Delete,attr,omitempty"`
	Hide     bool      `xml:"Hide,attr,omitempty"`
	PartTitle PartTitle `xml:"PartTitle"`
	Article  []Article `xml:"Article,omitempty"`
	Chapter  []Chapter `xml:"Chapter,omitempty"`
}

// PartTitle represents a part title
type PartTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Chapter represents a chapter
type Chapter struct {
	Num          string       `xml:"Num,attr"`
	Delete       bool         `xml:"Delete,attr,omitempty"`
	Hide         bool         `xml:"Hide,attr,omitempty"`
	ChapterTitle ChapterTitle `xml:"ChapterTitle"`
	Article      []Article    `xml:"Article,omitempty"`
	Section      []Section    `xml:"Section,omitempty"`
}

// ChapterTitle represents a chapter title
type ChapterTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Section represents a section
type Section struct {
	Num          string       `xml:"Num,attr"`
	Delete       bool         `xml:"Delete,attr,omitempty"`
	Hide         bool         `xml:"Hide,attr,omitempty"`
	SectionTitle SectionTitle `xml:"SectionTitle"`
	Article      []Article    `xml:"Article,omitempty"`
	Subsection   []Subsection `xml:"Subsection,omitempty"`
	Division     []Division   `xml:"Division,omitempty"`
}

// SectionTitle represents a section title
type SectionTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Subsection represents a subsection
type Subsection struct {
	Num             string          `xml:"Num,attr"`
	Delete          bool            `xml:"Delete,attr,omitempty"`
	Hide            bool            `xml:"Hide,attr,omitempty"`
	SubsectionTitle SubsectionTitle `xml:"SubsectionTitle"`
	Article         []Article       `xml:"Article,omitempty"`
	Division        []Division      `xml:"Division,omitempty"`
}

// SubsectionTitle represents a subsection title
type SubsectionTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Division represents a division
type Division struct {
	Num           string        `xml:"Num,attr"`
	Delete        bool          `xml:"Delete,attr,omitempty"`
	Hide          bool          `xml:"Hide,attr,omitempty"`
	DivisionTitle DivisionTitle `xml:"DivisionTitle"`
	Article       []Article     `xml:"Article"`
}

// DivisionTitle represents a division title
type DivisionTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Article represents an article
type Article struct {
	Num            string          `xml:"Num,attr"`
	Delete         bool            `xml:"Delete,attr,omitempty"`
	Hide           bool            `xml:"Hide,attr,omitempty"`
	ArticleCaption *ArticleCaption `xml:"ArticleCaption,omitempty"`
	ArticleTitle   *ArticleTitle   `xml:"ArticleTitle,omitempty"`
	Paragraph      []Paragraph     `xml:"Paragraph"`
	SupplNote      *SupplNote      `xml:"SupplNote,omitempty"`
}

// ArticleTitle represents an article title
type ArticleTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ArticleCaption represents an article caption
type ArticleCaption struct {
	CommonCaption bool   `xml:"CommonCaption,attr,omitempty"`
	Content       string `xml:",chardata"`
	Line          []Line `xml:"Line,omitempty"`
	Ruby          []Ruby `xml:"Ruby,omitempty"`
	Sup           []Sup  `xml:"Sup,omitempty"`
	Sub           []Sub  `xml:"Sub,omitempty"`
}

// Paragraph represents a paragraph
type Paragraph struct {
	Num               int                `xml:"Num,attr"`
	OldStyle          bool               `xml:"OldStyle,attr,omitempty"`
	OldNum            bool               `xml:"OldNum,attr,omitempty"`
	Hide              bool               `xml:"Hide,attr,omitempty"`
	ParagraphCaption  *ParagraphCaption  `xml:"ParagraphCaption,omitempty"`
	ParagraphNum      ParagraphNum       `xml:"ParagraphNum"`
	ParagraphSentence ParagraphSentence  `xml:"ParagraphSentence"`
	AmendProvision    []AmendProvision   `xml:"AmendProvision,omitempty"`
	Class             []Class            `xml:"Class,omitempty"`
	TableStruct       []TableStruct      `xml:"TableStruct,omitempty"`
	FigStruct         []FigStruct        `xml:"FigStruct,omitempty"`
	StyleStruct       []StyleStruct      `xml:"StyleStruct,omitempty"`
	Item              []Item             `xml:"Item,omitempty"`
	List              []List             `xml:"List,omitempty"`
}

// ParagraphCaption represents a paragraph caption
type ParagraphCaption struct {
	CommonCaption bool   `xml:"CommonCaption,attr,omitempty"`
	Content       string `xml:",chardata"`
	Line          []Line `xml:"Line,omitempty"`
	Ruby          []Ruby `xml:"Ruby,omitempty"`
	Sup           []Sup  `xml:"Sup,omitempty"`
	Sub           []Sub  `xml:"Sub,omitempty"`
}

// ParagraphNum represents a paragraph number
type ParagraphNum struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ParagraphSentence represents paragraph sentences
type ParagraphSentence struct {
	Sentence []Sentence `xml:"Sentence"`
}

// SupplNote represents a supplementary note
type SupplNote struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// AmendProvision represents an amendment provision
type AmendProvision struct {
	AmendProvisionSentence *AmendProvisionSentence `xml:"AmendProvisionSentence,omitempty"`
	NewProvision           []NewProvision          `xml:"NewProvision,omitempty"`
}

// AmendProvisionSentence represents an amendment provision sentence
type AmendProvisionSentence struct {
	Sentence Sentence `xml:"Sentence"`
}

// NewProvision represents a new provision
type NewProvision struct {
	Content                   string                      `xml:",innerxml"`
	LawTitle                  *LawTitle                   `xml:"LawTitle,omitempty"`
	Preamble                  *Preamble                   `xml:"Preamble,omitempty"`
	TOC                       *TOC                        `xml:"TOC,omitempty"`
	Part                      []Part                      `xml:"Part,omitempty"`
	PartTitle                 []PartTitle                 `xml:"PartTitle,omitempty"`
	Chapter                   []Chapter                   `xml:"Chapter,omitempty"`
	ChapterTitle              []ChapterTitle              `xml:"ChapterTitle,omitempty"`
	Section                   []Section                   `xml:"Section,omitempty"`
	SectionTitle              []SectionTitle              `xml:"SectionTitle,omitempty"`
	Subsection                []Subsection                `xml:"Subsection,omitempty"`
	SubsectionTitle           []SubsectionTitle           `xml:"SubsectionTitle,omitempty"`
	Division                  []Division                  `xml:"Division,omitempty"`
	DivisionTitle             []DivisionTitle             `xml:"DivisionTitle,omitempty"`
	Article                   []Article                   `xml:"Article,omitempty"`
	SupplNote                 []SupplNote                 `xml:"SupplNote,omitempty"`
	Paragraph                 []Paragraph                 `xml:"Paragraph,omitempty"`
	Item                      []Item                      `xml:"Item,omitempty"`
	Subitem1                  []Subitem1                  `xml:"Subitem1,omitempty"`
	Subitem2                  []Subitem2                  `xml:"Subitem2,omitempty"`
	Subitem3                  []Subitem3                  `xml:"Subitem3,omitempty"`
	Subitem4                  []Subitem4                  `xml:"Subitem4,omitempty"`
	Subitem5                  []Subitem5                  `xml:"Subitem5,omitempty"`
	Subitem6                  []Subitem6                  `xml:"Subitem6,omitempty"`
	Subitem7                  []Subitem7                  `xml:"Subitem7,omitempty"`
	Subitem8                  []Subitem8                  `xml:"Subitem8,omitempty"`
	Subitem9                  []Subitem9                  `xml:"Subitem9,omitempty"`
	Subitem10                 []Subitem10                 `xml:"Subitem10,omitempty"`
	List                      []List                      `xml:"List,omitempty"`
	Sentence                  []Sentence                  `xml:"Sentence,omitempty"`
	AmendProvision            []AmendProvision            `xml:"AmendProvision,omitempty"`
	AppdxTable                []AppdxTable                `xml:"AppdxTable,omitempty"`
	AppdxNote                 []AppdxNote                 `xml:"AppdxNote,omitempty"`
	AppdxStyle                []AppdxStyle                `xml:"AppdxStyle,omitempty"`
	Appdx                     []Appdx                     `xml:"Appdx,omitempty"`
	AppdxFig                  []AppdxFig                  `xml:"AppdxFig,omitempty"`
	AppdxFormat               []AppdxFormat               `xml:"AppdxFormat,omitempty"`
	SupplProvisionAppdxStyle  []SupplProvisionAppdxStyle  `xml:"SupplProvisionAppdxStyle,omitempty"`
	SupplProvisionAppdxTable  []SupplProvisionAppdxTable  `xml:"SupplProvisionAppdxTable,omitempty"`
	SupplProvisionAppdx       []SupplProvisionAppdx       `xml:"SupplProvisionAppdx,omitempty"`
	TableStruct               *TableStruct                `xml:"TableStruct,omitempty"`
	TableRow                  []TableRow                  `xml:"TableRow,omitempty"`
	TableColumn               []TableColumn               `xml:"TableColumn,omitempty"`
	FigStruct                 *FigStruct                  `xml:"FigStruct,omitempty"`
	NoteStruct                *NoteStruct                 `xml:"NoteStruct,omitempty"`
	StyleStruct               *StyleStruct                `xml:"StyleStruct,omitempty"`
	FormatStruct              *FormatStruct               `xml:"FormatStruct,omitempty"`
	Remarks                   *Remarks                    `xml:"Remarks,omitempty"`
	LawBody                   *LawBody                    `xml:"LawBody,omitempty"`
}

// Class represents a class
type Class struct {
	Num           string         `xml:"Num,attr"`
	ClassTitle    *ClassTitle    `xml:"ClassTitle,omitempty"`
	ClassSentence ClassSentence  `xml:"ClassSentence"`
	Item          []Item         `xml:"Item,omitempty"`
}

// ClassTitle represents a class title
type ClassTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ClassSentence represents a class sentence
type ClassSentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

// Item represents an item
type Item struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	Hide         bool          `xml:"Hide,attr,omitempty"`
	ItemTitle    *ItemTitle    `xml:"ItemTitle,omitempty"`
	ItemSentence ItemSentence  `xml:"ItemSentence"`
	Subitem1     []Subitem1    `xml:"Subitem1,omitempty"`
	TableStruct  []TableStruct `xml:"TableStruct,omitempty"`
	FigStruct    []FigStruct   `xml:"FigStruct,omitempty"`
	StyleStruct  []StyleStruct `xml:"StyleStruct,omitempty"`
	List         []List        `xml:"List,omitempty"`
}

// ItemTitle represents an item title
type ItemTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ItemSentence represents an item sentence
type ItemSentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

// Subitem1 through Subitem10 represent nested subitems
type Subitem1 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem1Title    *Subitem1Title    `xml:"Subitem1Title,omitempty"`
	Subitem1Sentence Subitem1Sentence  `xml:"Subitem1Sentence"`
	Subitem2         []Subitem2        `xml:"Subitem2,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem1Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem1Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem2 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem2Title    *Subitem2Title    `xml:"Subitem2Title,omitempty"`
	Subitem2Sentence Subitem2Sentence  `xml:"Subitem2Sentence"`
	Subitem3         []Subitem3        `xml:"Subitem3,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem2Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem2Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem3 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem3Title    *Subitem3Title    `xml:"Subitem3Title,omitempty"`
	Subitem3Sentence Subitem3Sentence  `xml:"Subitem3Sentence"`
	Subitem4         []Subitem4        `xml:"Subitem4,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem3Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem3Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem4 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem4Title    *Subitem4Title    `xml:"Subitem4Title,omitempty"`
	Subitem4Sentence Subitem4Sentence  `xml:"Subitem4Sentence"`
	Subitem5         []Subitem5        `xml:"Subitem5,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem4Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem4Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem5 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem5Title    *Subitem5Title    `xml:"Subitem5Title,omitempty"`
	Subitem5Sentence Subitem5Sentence  `xml:"Subitem5Sentence"`
	Subitem6         []Subitem6        `xml:"Subitem6,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem5Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem5Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem6 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem6Title    *Subitem6Title    `xml:"Subitem6Title,omitempty"`
	Subitem6Sentence Subitem6Sentence  `xml:"Subitem6Sentence"`
	Subitem7         []Subitem7        `xml:"Subitem7,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem6Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem6Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem7 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem7Title    *Subitem7Title    `xml:"Subitem7Title,omitempty"`
	Subitem7Sentence Subitem7Sentence  `xml:"Subitem7Sentence"`
	Subitem8         []Subitem8        `xml:"Subitem8,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem7Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem7Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem8 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem8Title    *Subitem8Title    `xml:"Subitem8Title,omitempty"`
	Subitem8Sentence Subitem8Sentence  `xml:"Subitem8Sentence"`
	Subitem9         []Subitem9        `xml:"Subitem9,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem8Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem8Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem9 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem9Title    *Subitem9Title    `xml:"Subitem9Title,omitempty"`
	Subitem9Sentence Subitem9Sentence  `xml:"Subitem9Sentence"`
	Subitem10        []Subitem10       `xml:"Subitem10,omitempty"`
	TableStruct      []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct     `xml:"StyleStruct,omitempty"`
	List             []List            `xml:"List,omitempty"`
}

type Subitem9Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem9Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem10 struct {
	Num               string             `xml:"Num,attr"`
	Delete            bool               `xml:"Delete,attr,omitempty"`
	Hide              bool               `xml:"Hide,attr,omitempty"`
	Subitem10Title    *Subitem10Title    `xml:"Subitem10Title,omitempty"`
	Subitem10Sentence Subitem10Sentence  `xml:"Subitem10Sentence"`
	TableStruct       []TableStruct      `xml:"TableStruct,omitempty"`
	FigStruct         []FigStruct        `xml:"FigStruct,omitempty"`
	StyleStruct       []StyleStruct      `xml:"StyleStruct,omitempty"`
	List              []List             `xml:"List,omitempty"`
}

type Subitem10Title struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

type Subitem10Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

// Sentence represents a sentence
// NOTE: This type now uses custom XML unmarshaling to properly handle mixed content with inline Ruby elements
type Sentence struct {
	Num          int             `xml:"Num,attr,omitempty"`
	Function     string          `xml:"Function,attr,omitempty"` // main or proviso
	Indent       string          `xml:"Indent,attr,omitempty"`
	WritingMode  WritingMode     `xml:"WritingMode,attr,omitempty"`
	Content      string          // Plain text content (backward compatibility)
	Line         []Line          `xml:"Line,omitempty"`
	QuoteStruct  []QuoteStruct   `xml:"QuoteStruct,omitempty"`
	ArithFormula []ArithFormula  `xml:"ArithFormula,omitempty"`
	Ruby         []Ruby          // Ruby elements (backward compatibility)
	Sup          []Sup           `xml:"Sup,omitempty"`
	Sub          []Sub           `xml:"Sub,omitempty"`
	MixedContent MixedContent    // The properly ordered mixed content
}

// UnmarshalXML implements custom unmarshaling to handle mixed content
func (s *Sentence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "Num":
			fmt.Sscanf(attr.Value, "%d", &s.Num)
		case "Function":
			s.Function = attr.Value
		case "Indent":
			s.Indent = attr.Value
		case "WritingMode":
			s.WritingMode = WritingMode(attr.Value)
		}
	}
	
	// Parse mixed content
	if err := s.MixedContent.UnmarshalXML(d, start); err != nil {
		return err
	}
	
	// Set backward compatibility fields
	s.Content = s.MixedContent.String()
	
	// Extract Ruby elements for backward compatibility
	s.Ruby = []Ruby{}
	for _, node := range s.MixedContent.Nodes {
		if rubyNode, ok := node.(RubyNode); ok && rubyNode.IsRuby() {
			ruby := Ruby{
				Content: rubyNode.Base,
				Rt:      make([]Rt, len(rubyNode.Rt)),
			}
			for i, rt := range rubyNode.Rt {
				ruby.Rt[i] = Rt{Content: rt}
			}
			s.Ruby = append(s.Ruby, ruby)
		}
	}
	
	return nil
}

// HTML returns the HTML representation with properly positioned Ruby tags
func (s *Sentence) HTML() string {
	return s.MixedContent.HTML()
}

// Column represents a column
type Column struct {
	Num        int       `xml:"Num,attr,omitempty"`
	LineBreak  bool      `xml:"LineBreak,attr,omitempty"`
	Align      string    `xml:"Align,attr,omitempty"` // left, center, right, justify
	Sentence   []Sentence `xml:"Sentence"`
}

// SupplProvision represents supplementary provisions
type SupplProvision struct {
	Type                     string                      `xml:"Type,attr,omitempty"` // New or Amend
	AmendLawNum              string                      `xml:"AmendLawNum,attr,omitempty"`
	Extract                  bool                        `xml:"Extract,attr,omitempty"`
	SupplProvisionLabel      SupplProvisionLabel         `xml:"SupplProvisionLabel"`
	Chapter                  []Chapter                   `xml:"Chapter,omitempty"`
	Article                  []Article                   `xml:"Article,omitempty"`
	Paragraph                []Paragraph                 `xml:"Paragraph,omitempty"`
	SupplProvisionAppdxTable []SupplProvisionAppdxTable  `xml:"SupplProvisionAppdxTable,omitempty"`
	SupplProvisionAppdxStyle []SupplProvisionAppdxStyle  `xml:"SupplProvisionAppdxStyle,omitempty"`
	SupplProvisionAppdx      []SupplProvisionAppdx       `xml:"SupplProvisionAppdx,omitempty"`
}

// SupplProvisionLabel represents a supplementary provision label
type SupplProvisionLabel struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// SupplProvisionAppdxTable represents supplementary provision appendix table
type SupplProvisionAppdxTable struct {
	Num                          int                           `xml:"Num,attr,omitempty"`
	SupplProvisionAppdxTableTitle SupplProvisionAppdxTableTitle `xml:"SupplProvisionAppdxTableTitle"`
	RelatedArticleNum            *RelatedArticleNum            `xml:"RelatedArticleNum,omitempty"`
	TableStruct                  []TableStruct                 `xml:"TableStruct,omitempty"`
}

// SupplProvisionAppdxTableTitle represents title for supplementary provision appendix table
type SupplProvisionAppdxTableTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// SupplProvisionAppdxStyle represents supplementary provision appendix style
type SupplProvisionAppdxStyle struct {
	Num                          int                           `xml:"Num,attr,omitempty"`
	SupplProvisionAppdxStyleTitle SupplProvisionAppdxStyleTitle `xml:"SupplProvisionAppdxStyleTitle"`
	RelatedArticleNum            *RelatedArticleNum            `xml:"RelatedArticleNum,omitempty"`
	StyleStruct                  []StyleStruct                 `xml:"StyleStruct,omitempty"`
}

// SupplProvisionAppdxStyleTitle represents title for supplementary provision appendix style
type SupplProvisionAppdxStyleTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// SupplProvisionAppdx represents supplementary provision appendix
type SupplProvisionAppdx struct {
	Num               int                `xml:"Num,attr,omitempty"`
	ArithFormulaNum   *ArithFormulaNum   `xml:"ArithFormulaNum,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	ArithFormula      []ArithFormula     `xml:"ArithFormula"`
}

// AppdxTable represents an appendix table
type AppdxTable struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxTableTitle   *AppdxTableTitle   `xml:"AppdxTableTitle,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	TableStruct       []TableStruct      `xml:"TableStruct,omitempty"`
	Item              []Item             `xml:"Item,omitempty"`
	Remarks           *Remarks           `xml:"Remarks,omitempty"`
}

// AppdxTableTitle represents an appendix table title
type AppdxTableTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// AppdxNote represents an appendix note
type AppdxNote struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxNoteTitle    *AppdxNoteTitle    `xml:"AppdxNoteTitle,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	NoteStruct        []NoteStruct       `xml:"NoteStruct,omitempty"`
	FigStruct         []FigStruct        `xml:"FigStruct,omitempty"`
	TableStruct       []TableStruct      `xml:"TableStruct,omitempty"`
	Remarks           *Remarks           `xml:"Remarks,omitempty"`
}

// AppdxNoteTitle represents an appendix note title
type AppdxNoteTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// AppdxStyle represents an appendix style
type AppdxStyle struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxStyleTitle   *AppdxStyleTitle   `xml:"AppdxStyleTitle,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	StyleStruct       []StyleStruct      `xml:"StyleStruct,omitempty"`
	Remarks           *Remarks           `xml:"Remarks,omitempty"`
}

// AppdxStyleTitle represents an appendix style title
type AppdxStyleTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// AppdxFormat represents an appendix format
type AppdxFormat struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxFormatTitle  *AppdxFormatTitle  `xml:"AppdxFormatTitle,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	FormatStruct      []FormatStruct     `xml:"FormatStruct,omitempty"`
	Remarks           *Remarks           `xml:"Remarks,omitempty"`
}

// AppdxFormatTitle represents an appendix format title
type AppdxFormatTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// Appdx represents an appendix
type Appdx struct {
	ArithFormulaNum   *ArithFormulaNum   `xml:"ArithFormulaNum,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	ArithFormula      []ArithFormula     `xml:"ArithFormula"`
	Remarks           *Remarks           `xml:"Remarks,omitempty"`
}

// ArithFormulaNum represents an arithmetic formula number
type ArithFormulaNum struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// ArithFormula represents an arithmetic formula
type ArithFormula struct {
	Num     int    `xml:"Num,attr,omitempty"`
	Content string `xml:",innerxml"`
}

// AppdxFig represents an appendix figure
type AppdxFig struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxFigTitle     *AppdxFigTitle     `xml:"AppdxFigTitle,omitempty"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum,omitempty"`
	FigStruct         []FigStruct        `xml:"FigStruct,omitempty"`
	TableStruct       []TableStruct      `xml:"TableStruct,omitempty"`
}

// AppdxFigTitle represents an appendix figure title
type AppdxFigTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// TableStruct represents a table structure
type TableStruct struct {
	TableStructTitle *TableStructTitle `xml:"TableStructTitle,omitempty"`
	Table            Table             `xml:"Table"`
	Remarks          []Remarks         `xml:"Remarks,omitempty"`
}

// TableStructTitle represents a table structure title
type TableStructTitle struct {
	WritingMode WritingMode `xml:"WritingMode,attr,omitempty"`
	Content     string      `xml:",chardata"`
	Line        []Line      `xml:"Line,omitempty"`
	Ruby        []Ruby      `xml:"Ruby,omitempty"`
	Sup         []Sup       `xml:"Sup,omitempty"`
	Sub         []Sub       `xml:"Sub,omitempty"`
}

// Table represents a table
type Table struct {
	WritingMode     WritingMode       `xml:"WritingMode,attr,omitempty"`
	TableHeaderRow  []TableHeaderRow  `xml:"TableHeaderRow,omitempty"`
	TableRow        []TableRow        `xml:"TableRow"`
}

// TableRow represents a table row
type TableRow struct {
	TableColumn []TableColumn `xml:"TableColumn"`
}

// TableHeaderRow represents a table header row
type TableHeaderRow struct {
	TableHeaderColumn []TableHeaderColumn `xml:"TableHeaderColumn"`
}

// TableHeaderColumn represents a table header column
type TableHeaderColumn struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// TableColumn represents a table column
type TableColumn struct {
	BorderTop    string       `xml:"BorderTop,attr,omitempty"`    // solid, none, dotted, double
	BorderBottom string       `xml:"BorderBottom,attr,omitempty"` // solid, none, dotted, double
	BorderLeft   string       `xml:"BorderLeft,attr,omitempty"`   // solid, none, dotted, double
	BorderRight  string       `xml:"BorderRight,attr,omitempty"`  // solid, none, dotted, double
	Rowspan      string       `xml:"rowspan,attr,omitempty"`
	Colspan      string       `xml:"colspan,attr,omitempty"`
	Align        string       `xml:"Align,attr,omitempty"`        // left, center, right, justify
	Valign       string       `xml:"Valign,attr,omitempty"`       // top, middle, bottom
	Part         []Part       `xml:"Part,omitempty"`
	Chapter      []Chapter    `xml:"Chapter,omitempty"`
	Section      []Section    `xml:"Section,omitempty"`
	Subsection   []Subsection `xml:"Subsection,omitempty"`
	Division     []Division   `xml:"Division,omitempty"`
	Article      []Article    `xml:"Article,omitempty"`
	Paragraph    []Paragraph  `xml:"Paragraph,omitempty"`
	Item         []Item       `xml:"Item,omitempty"`
	Subitem1     []Subitem1   `xml:"Subitem1,omitempty"`
	Subitem2     []Subitem2   `xml:"Subitem2,omitempty"`
	Subitem3     []Subitem3   `xml:"Subitem3,omitempty"`
	Subitem4     []Subitem4   `xml:"Subitem4,omitempty"`
	Subitem5     []Subitem5   `xml:"Subitem5,omitempty"`
	Subitem6     []Subitem6   `xml:"Subitem6,omitempty"`
	Subitem7     []Subitem7   `xml:"Subitem7,omitempty"`
	Subitem8     []Subitem8   `xml:"Subitem8,omitempty"`
	Subitem9     []Subitem9   `xml:"Subitem9,omitempty"`
	Subitem10    []Subitem10  `xml:"Subitem10,omitempty"`
	FigStruct    []FigStruct  `xml:"FigStruct,omitempty"`
	Remarks      *Remarks     `xml:"Remarks,omitempty"`
	Sentence     []Sentence   `xml:"Sentence,omitempty"`
	Column       []Column     `xml:"Column,omitempty"`
}

// FigStruct represents a figure structure
type FigStruct struct {
	FigStructTitle *FigStructTitle `xml:"FigStructTitle,omitempty"`
	Fig            Fig             `xml:"Fig"`
	Remarks        []Remarks       `xml:"Remarks,omitempty"`
}

// FigStructTitle represents a figure structure title
type FigStructTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Fig represents a figure
type Fig struct {
	Src string `xml:"src,attr"`
}

// NoteStruct represents a note structure
type NoteStruct struct {
	NoteStructTitle *NoteStructTitle `xml:"NoteStructTitle,omitempty"`
	Note            Note             `xml:"Note"`
	Remarks         []Remarks        `xml:"Remarks,omitempty"`
}

// NoteStructTitle represents a note structure title
type NoteStructTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Note represents a note
type Note struct {
	Content string `xml:",innerxml"`
}

// StyleStruct represents a style structure
type StyleStruct struct {
	StyleStructTitle *StyleStructTitle `xml:"StyleStructTitle,omitempty"`
	Style            Style             `xml:"Style"`
	Remarks          []Remarks         `xml:"Remarks,omitempty"`
}

// StyleStructTitle represents a style structure title
type StyleStructTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Style represents a style
type Style struct {
	Content string `xml:",innerxml"`
}

// FormatStruct represents a format structure
type FormatStruct struct {
	FormatStructTitle *FormatStructTitle `xml:"FormatStructTitle,omitempty"`
	Format            Format             `xml:"Format"`
	Remarks           []Remarks          `xml:"Remarks,omitempty"`
}

// FormatStructTitle represents a format structure title
type FormatStructTitle struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Format represents a format
type Format struct {
	Content string `xml:",innerxml"`
}

// RelatedArticleNum represents a related article number
type RelatedArticleNum struct {
	Content string `xml:",chardata"`
	Line    []Line `xml:"Line,omitempty"`
	Ruby    []Ruby `xml:"Ruby,omitempty"`
	Sup     []Sup  `xml:"Sup,omitempty"`
	Sub     []Sub  `xml:"Sub,omitempty"`
}

// Remarks represents remarks
type Remarks struct {
	RemarksLabel RemarksLabel `xml:"RemarksLabel"`
	Item         []Item       `xml:"Item,omitempty"`
	Sentence     []Sentence   `xml:"Sentence,omitempty"`
}

// RemarksLabel represents a remarks label
type RemarksLabel struct {
	LineBreak bool   `xml:"LineBreak,attr,omitempty"`
	Content   string `xml:",chardata"`
	Line      []Line `xml:"Line,omitempty"`
	Ruby      []Ruby `xml:"Ruby,omitempty"`
	Sup       []Sup  `xml:"Sup,omitempty"`
	Sub       []Sub  `xml:"Sub,omitempty"`
}

// List represents a list
type List struct {
	ListSentence ListSentence `xml:"ListSentence"`
	Sublist1     []Sublist1   `xml:"Sublist1,omitempty"`
}

// ListSentence represents a list sentence
type ListSentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
}

// Sublist1 represents a first-level sublist
type Sublist1 struct {
	Sublist1Sentence Sublist1Sentence `xml:"Sublist1Sentence"`
	Sublist2         []Sublist2       `xml:"Sublist2,omitempty"`
}

// Sublist1Sentence represents a first-level sublist sentence
type Sublist1Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
}

// Sublist2 represents a second-level sublist
type Sublist2 struct {
	Sublist2Sentence Sublist2Sentence `xml:"Sublist2Sentence"`
	Sublist3         []Sublist3       `xml:"Sublist3,omitempty"`
}

// Sublist2Sentence represents a second-level sublist sentence
type Sublist2Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
}

// Sublist3 represents a third-level sublist
type Sublist3 struct {
	Sublist3Sentence Sublist3Sentence `xml:"Sublist3Sentence"`
}

// Sublist3Sentence represents a third-level sublist sentence
type Sublist3Sentence struct {
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
}

// QuoteStruct represents a quote structure
type QuoteStruct struct {
	Content string `xml:",innerxml"`
}

// Ruby represents ruby text (phonetic guides)
type Ruby struct {
	Content string `xml:",chardata"`
	Rt      []Rt   `xml:"Rt,omitempty"`
}

// Rt represents ruby text content
type Rt struct {
	Content string `xml:",chardata"`
}

// Line represents a line
type Line struct {
	Style        string         `xml:"Style,attr,omitempty"` // dotted, double, none, solid
	Content      string         `xml:",chardata"`
	QuoteStruct  []QuoteStruct  `xml:"QuoteStruct,omitempty"`
	ArithFormula []ArithFormula `xml:"ArithFormula,omitempty"`
	Ruby         []Ruby         `xml:"Ruby,omitempty"`
	Sup          []Sup          `xml:"Sup,omitempty"`
	Sub          []Sub          `xml:"Sub,omitempty"`
}

// Sup represents superscript text
type Sup struct {
	Content string `xml:",chardata"`
}

// Sub represents subscript text
type Sub struct {
	Content string `xml:",chardata"`
}

// WritingMode represents text direction
type WritingMode string

const (
	WritingModeVertical   WritingMode = "vertical"
	WritingModeHorizontal WritingMode = "horizontal"
)