package jplaw

import "encoding/xml"

type Law struct {
	XMLName         xml.Name `xml:"Law"`
	Era             string   `xml:"Era,attr"`
	Year            int      `xml:"Year,attr"`
	Num             int      `xml:"Num,attr"`
	PromulgateMonth int      `xml:"PromulgateMonth,attr,omitempty"`
	PromulgateDay   int      `xml:"PromulgateDay,attr,omitempty"`
	LawType         string   `xml:"LawType,attr"`
	Lang            string   `xml:"Lang,attr"`
	LawNum          string   `xml:"LawNum"`
	LawBody         LawBody  `xml:"LawBody"`
}

type LawBody struct {
	XMLName        xml.Name         `xml:"LawBody"`
	LawTitle       *LawTitle        `xml:"LawTitle,omitempty"`
	EnactStatement []EnactStatement `xml:"EnactStatement,omitempty"`
	TOC            *TOC             `xml:"TOC,omitempty"`
	Preamble       *Preamble        `xml:"Preamble,omitempty"`
	MainProvision  MainProvision    `xml:"MainProvision"`
	SupplProvision []SupplProvision `xml:"SupplProvision,omitempty"`
	AppdxTable     []AppdxTable     `xml:"AppdxTable,omitempty"`
	AppdxNote      []AppdxNote      `xml:"AppdxNote,omitempty"`
	AppdxStyle     []AppdxStyle     `xml:"AppdxStyle,omitempty"`
	Appdx          []Appdx          `xml:"Appdx,omitempty"`
	AppdxFig       []AppdxFig       `xml:"AppdxFig,omitempty"`
	AppdxFormat    []AppdxFormat    `xml:"AppdxFormat,omitempty"`
	Subject        string           `xml:"Subject,attr,omitempty"`
}

type LawTitle struct {
	XMLName    xml.Name `xml:"LawTitle"`
	Kana       string   `xml:"Kana,attr,omitempty"`
	Abbrev     string   `xml:"Abbrev,attr,omitempty"`
	AbbrevKana string   `xml:"AbbrevKana,attr,omitempty"`
	Content    string   `xml:",chardata"`
}

type EnactStatement struct {
	XMLName xml.Name `xml:"EnactStatement"`
	Content string   `xml:",chardata"`
}

type TOC struct {
	XMLName            xml.Name             `xml:"TOC"`
	TOCLabel           *TOCLabel            `xml:"TOCLabel,omitempty"`
	TOCPreambleLabel   *TOCPreambleLabel    `xml:"TOCPreambleLabel,omitempty"`
	TOCPart            []TOCPart            `xml:"TOCPart,omitempty"`
	TOCChapter         []TOCChapter         `xml:"TOCChapter,omitempty"`
	TOCSection         []TOCSection         `xml:"TOCSection,omitempty"`
	TOCArticle         []TOCArticle         `xml:"TOCArticle,omitempty"`
	TOCSupplProvision  *TOCSupplProvision   `xml:"TOCSupplProvision,omitempty"`
	TOCAppdxTableLabel []TOCAppdxTableLabel `xml:"TOCAppdxTableLabel,omitempty"`
}

type TOCLabel struct {
	XMLName xml.Name `xml:"TOCLabel"`
	Content string   `xml:",chardata"`
}

type TOCPreambleLabel struct {
	XMLName xml.Name `xml:"TOCPreambleLabel"`
	Content string   `xml:",chardata"`
}

type TOCPart struct {
	XMLName      xml.Name      `xml:"TOCPart"`
	PartTitle    PartTitle     `xml:"PartTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange,omitempty"`
	TOCChapter   []TOCChapter  `xml:"TOCChapter,omitempty"`
	Num          int           `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
}

type TOCChapter struct {
	XMLName      xml.Name      `xml:"TOCChapter"`
	ChapterTitle ChapterTitle  `xml:"ChapterTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange,omitempty"`
	TOCSection   []TOCSection  `xml:"TOCSection,omitempty"`
	Num          int           `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
}

type TOCSection struct {
	XMLName       xml.Name        `xml:"TOCSection"`
	SectionTitle  SectionTitle    `xml:"SectionTitle"`
	ArticleRange  *ArticleRange   `xml:"ArticleRange,omitempty"`
	TOCSubsection []TOCSubsection `xml:"TOCSubsection,omitempty"`
	TOCDivision   []TOCDivision   `xml:"TOCDivision,omitempty"`
	Num           int             `xml:"Num,attr"`
	Delete        bool            `xml:"Delete,attr,omitempty"`
}

type TOCSubsection struct {
	XMLName         xml.Name        `xml:"TOCSubsection"`
	SubsectionTitle SubsectionTitle `xml:"SubsectionTitle"`
	ArticleRange    *ArticleRange   `xml:"ArticleRange,omitempty"`
	TOCDivision     []TOCDivision   `xml:"TOCDivision,omitempty"`
	Num             int             `xml:"Num,attr"`
	Delete          bool            `xml:"Delete,attr,omitempty"`
}

type TOCDivision struct {
	XMLName       xml.Name      `xml:"TOCDivision"`
	DivisionTitle DivisionTitle `xml:"DivisionTitle"`
	ArticleRange  *ArticleRange `xml:"ArticleRange,omitempty"`
	Num           int           `xml:"Num,attr"`
	Delete        bool          `xml:"Delete,attr,omitempty"`
}

type TOCArticle struct {
	XMLName        xml.Name       `xml:"TOCArticle"`
	ArticleTitle   ArticleTitle   `xml:"ArticleTitle"`
	ArticleCaption ArticleCaption `xml:"ArticleCaption"`
	Num            int            `xml:"Num,attr"`
	Delete         bool           `xml:"Delete,attr,omitempty"`
}

type TOCSupplProvision struct {
	XMLName             xml.Name            `xml:"TOCSupplProvision"`
	SupplProvisionLabel SupplProvisionLabel `xml:"SupplProvisionLabel"`
	ArticleRange        *ArticleRange       `xml:"ArticleRange,omitempty"`
	TOCArticle          []TOCArticle        `xml:"TOCArticle,omitempty"`
	TOCChapter          []TOCChapter        `xml:"TOCChapter,omitempty"`
}

type TOCAppdxTableLabel struct {
	XMLName xml.Name `xml:"TOCAppdxTableLabel"`
	Content string   `xml:",chardata"`
}

type ArticleRange struct {
	XMLName xml.Name `xml:"ArticleRange"`
	Content string   `xml:",chardata"`
}

type Preamble struct {
	XMLName   xml.Name    `xml:"Preamble"`
	Paragraph []Paragraph `xml:"Paragraph"`
}

type MainProvision struct {
	XMLName   xml.Name    `xml:"MainProvision"`
	Part      []Part      `xml:"Part,omitempty"`
	Chapter   []Chapter   `xml:"Chapter,omitempty"`
	Section   []Section   `xml:"Section,omitempty"`
	Article   []Article   `xml:"Article,omitempty"`
	Paragraph []Paragraph `xml:"Paragraph,omitempty"`
	Extract   bool        `xml:"Extract,attr,omitempty"`
}

type Part struct {
	XMLName   xml.Name  `xml:"Part"`
	PartTitle PartTitle `xml:"PartTitle"`
	Article   []Article `xml:"Article,omitempty"`
	Chapter   []Chapter `xml:"Chapter,omitempty"`
	Num       int       `xml:"Num,attr"`
	Delete    bool      `xml:"Delete,attr,omitempty"`
	Hide      bool      `xml:"Hide,attr,omitempty"`
}

type PartTitle struct {
	XMLName xml.Name `xml:"PartTitle"`
	Content string   `xml:",chardata"`
}

type Chapter struct {
	XMLName      xml.Name     `xml:"Chapter"`
	ChapterTitle ChapterTitle `xml:"ChapterTitle"`
	Article      []Article    `xml:"Article,omitempty"`
	Section      []Section    `xml:"Section,omitempty"`
	Num          int          `xml:"Num,attr"`
	Delete       bool         `xml:"Delete,attr,omitempty"`
	Hide         bool         `xml:"Hide,attr,omitempty"`
}

type ChapterTitle struct {
	XMLName xml.Name `xml:"ChapterTitle"`
	Content string   `xml:",chardata"`
}

type Section struct {
	XMLName      xml.Name     `xml:"Section"`
	SectionTitle SectionTitle `xml:"SectionTitle"`
	Article      []Article    `xml:"Article,omitempty"`
	Subsection   []Subsection `xml:"Subsection,omitempty"`
	Division     []Division   `xml:"Division,omitempty"`
	Num          int          `xml:"Num,attr"`
	Delete       bool         `xml:"Delete,attr,omitempty"`
	Hide         bool         `xml:"Hide,attr,omitempty"`
}

type SectionTitle struct {
	XMLName xml.Name `xml:"SectionTitle"`
	Content string   `xml:",chardata"`
}

type Subsection struct {
	XMLName         xml.Name        `xml:"Subsection"`
	SubsectionTitle SubsectionTitle `xml:"SubsectionTitle"`
	Article         []Article       `xml:"Article,omitempty"`
	Division        []Division      `xml:"Division,omitempty"`
	Num             int             `xml:"Num,attr"`
	Delete          bool            `xml:"Delete,attr,omitempty"`
	Hide            bool            `xml:"Hide,attr,omitempty"`
}

type SubsectionTitle struct {
	XMLName xml.Name `xml:"SubsectionTitle"`
	Content string   `xml:",chardata"`
}

type Division struct {
	XMLName       xml.Name      `xml:"Division"`
	DivisionTitle DivisionTitle `xml:"DivisionTitle"`
	Article       []Article     `xml:"Article,omitempty"`
	Num           int           `xml:"Num,attr"`
	Delete        bool          `xml:"Delete,attr,omitempty"`
	Hide          bool          `xml:"Hide,attr,omitempty"`
}

type DivisionTitle struct {
	XMLName xml.Name `xml:"DivisionTitle"`
	Content string   `xml:",chardata"`
}

type Article struct {
	XMLName        xml.Name        `xml:"Article"`
	ArticleCaption *ArticleCaption `xml:"ArticleCaption,omitempty"`
	ArticleTitle   *ArticleTitle   `xml:"ArticleTitle,omitempty"`
	Paragraph      []Paragraph     `xml:"Paragraph"`
	SupplNote      *SupplNote      `xml:"SupplNote,omitempty"`
	Num            string          `xml:"Num,attr"`
	Delete         bool            `xml:"Delete,attr,omitempty"`
	Hide           bool            `xml:"Hide,attr,omitempty"`
}

type ArticleTitle struct {
	XMLName xml.Name `xml:"ArticleTitle"`
	Content string   `xml:",chardata"`
}

type ArticleCaption struct {
	XMLName       xml.Name `xml:"ArticleCaption"`
	Content       string   `xml:",chardata"`
	CommonCaption bool     `xml:"CommonCaption,attr,omitempty"`
}

type Paragraph struct {
	XMLName           xml.Name          `xml:"Paragraph"`
	ParagraphCaption  *ParagraphCaption `xml:"ParagraphCaption,omitempty"`
	ParagraphNum      ParagraphNum      `xml:"ParagraphNum"`
	ParagraphSentence ParagraphSentence `xml:"ParagraphSentence"`
	AmendProvision    []AmendProvision  `xml:"AmendProvision,omitempty"`
	Class             []Class           `xml:"Class,omitempty"`
	TableStruct       []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct         []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct       []StyleStruct     `xml:"StyleStruct,omitempty"`
	Item              []Item            `xml:"Item,omitempty"`
	List              []List            `xml:"List,omitempty"`
	Num               int               `xml:"Num,attr"`
	OldStyle          bool              `xml:"OldStyle,attr,omitempty"`
	OldNum            bool              `xml:"OldNum,attr,omitempty"`
	Hide              bool              `xml:"Hide,attr,omitempty"`
}

type ParagraphCaption struct {
	XMLName       xml.Name `xml:"ParagraphCaption"`
	Content       string   `xml:",chardata"`
	CommonCaption bool     `xml:"CommonCaption,attr,omitempty"`
}

type ParagraphNum struct {
	XMLName xml.Name `xml:"ParagraphNum"`
	Content string   `xml:",chardata"`
}

type ParagraphSentence struct {
	XMLName  xml.Name   `xml:"ParagraphSentence"`
	Sentence []Sentence `xml:"Sentence"`
}

type SupplNote struct {
	XMLName xml.Name `xml:"SupplNote"`
	Content string   `xml:",chardata"`
}

type AmendProvision struct {
	XMLName                xml.Name                `xml:"AmendProvision"`
	AmendProvisionSentence *AmendProvisionSentence `xml:"AmendProvisionSentence,omitempty"`
	NewProvision           []NewProvision          `xml:"NewProvision,omitempty"`
}

type AmendProvisionSentence struct {
	XMLName  xml.Name `xml:"AmendProvisionSentence"`
	Sentence Sentence `xml:"Sentence"`
}

type NewProvision struct {
	XMLName                  xml.Name                   `xml:"NewProvision"`
	LawTitle                 *LawTitle                  `xml:"LawTitle,omitempty"`
	Preamble                 *Preamble                  `xml:"Preamble,omitempty"`
	TOC                      *TOC                       `xml:"TOC,omitempty"`
	Part                     []Part                     `xml:"Part,omitempty"`
	PartTitle                []PartTitle                `xml:"PartTitle,omitempty"`
	Chapter                  []Chapter                  `xml:"Chapter,omitempty"`
	ChapterTitle             []ChapterTitle             `xml:"ChapterTitle,omitempty"`
	Section                  []Section                  `xml:"Section,omitempty"`
	SectionTitle             []SectionTitle             `xml:"SectionTitle,omitempty"`
	Subsection               []Subsection               `xml:"Subsection,omitempty"`
	SubsectionTitle          []SubsectionTitle          `xml:"SubsectionTitle,omitempty"`
	Division                 []Division                 `xml:"Division,omitempty"`
	DivisionTitle            []DivisionTitle            `xml:"DivisionTitle,omitempty"`
	Article                  []Article                  `xml:"Article,omitempty"`
	SupplNote                []SupplNote                `xml:"SupplNote,omitempty"`
	Paragraph                []Paragraph                `xml:"Paragraph,omitempty"`
	Item                     []Item                     `xml:"Item,omitempty"`
	Subitem1                 []Subitem1                 `xml:"Subitem1,omitempty"`
	Subitem2                 []Subitem2                 `xml:"Subitem2,omitempty"`
	Subitem3                 []Subitem3                 `xml:"Subitem3,omitempty"`
	Subitem4                 []Subitem4                 `xml:"Subitem4,omitempty"`
	Subitem5                 []Subitem5                 `xml:"Subitem5,omitempty"`
	Subitem6                 []Subitem6                 `xml:"Subitem6,omitempty"`
	Subitem7                 []Subitem7                 `xml:"Subitem7,omitempty"`
	Subitem8                 []Subitem8                 `xml:"Subitem8,omitempty"`
	Subitem9                 []Subitem9                 `xml:"Subitem9,omitempty"`
	Subitem10                []Subitem10                `xml:"Subitem10,omitempty"`
	List                     []List                     `xml:"List,omitempty"`
	Sentence                 []Sentence                 `xml:"Sentence,omitempty"`
	AmendProvision           []AmendProvision           `xml:"AmendProvision,omitempty"`
	AppdxTable               []AppdxTable               `xml:"AppdxTable,omitempty"`
	AppdxNote                []AppdxNote                `xml:"AppdxNote,omitempty"`
	AppdxStyle               []AppdxStyle               `xml:"AppdxStyle,omitempty"`
	Appdx                    []Appdx                    `xml:"Appdx,omitempty"`
	AppdxFig                 []AppdxFig                 `xml:"AppdxFig,omitempty"`
	AppdxFormat              []AppdxFormat              `xml:"AppdxFormat,omitempty"`
	SupplProvisionAppdxStyle []SupplProvisionAppdxStyle `xml:"SupplProvisionAppdxStyle,omitempty"`
	SupplProvisionAppdxTable []SupplProvisionAppdxTable `xml:"SupplProvisionAppdxTable,omitempty"`
	SupplProvisionAppdx      []SupplProvisionAppdx      `xml:"SupplProvisionAppdx,omitempty"`
	TableStruct              []TableStruct              `xml:"TableStruct,omitempty"`
	TableRow                 []TableRow                 `xml:"TableRow,omitempty"`
	TableColumn              []TableColumn              `xml:"TableColumn,omitempty"`
	FigStruct                []FigStruct                `xml:"FigStruct,omitempty"`
	NoteStruct               []NoteStruct               `xml:"NoteStruct,omitempty"`
	StyleStruct              []StyleStruct              `xml:"StyleStruct,omitempty"`
	FormatStruct             []FormatStruct             `xml:"FormatStruct,omitempty"`
	Remarks                  []Remarks                  `xml:"Remarks,omitempty"`
	LawBody                  *LawBody                   `xml:"LawBody,omitempty"`
}

type Class struct {
	XMLName       xml.Name      `xml:"Class"`
	ClassTitle    *ClassTitle   `xml:"ClassTitle,omitempty"`
	ClassSentence ClassSentence `xml:"ClassSentence"`
	Item          []Item        `xml:"Item,omitempty"`
	Num           int           `xml:"Num,attr"`
}

type ClassTitle struct {
	XMLName xml.Name `xml:"ClassTitle"`
	Content string   `xml:",chardata"`
}

type ClassSentence struct {
	XMLName  xml.Name   `xml:"ClassSentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Item struct {
	XMLName      xml.Name      `xml:"Item"`
	ItemTitle    *ItemTitle    `xml:"ItemTitle,omitempty"`
	ItemSentence ItemSentence  `xml:"ItemSentence"`
	Subitem1     []Subitem1    `xml:"Subitem1,omitempty"`
	TableStruct  []TableStruct `xml:"TableStruct,omitempty"`
	FigStruct    []FigStruct   `xml:"FigStruct,omitempty"`
	StyleStruct  []StyleStruct `xml:"StyleStruct,omitempty"`
	List         []List        `xml:"List,omitempty"`
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	Hide         bool          `xml:"Hide,attr,omitempty"`
}

type ItemTitle struct {
	XMLName xml.Name `xml:"ItemTitle"`
	Content string   `xml:",chardata"`
}

type ItemSentence struct {
	XMLName  xml.Name   `xml:"ItemSentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem1 struct {
	XMLName          xml.Name         `xml:"Subitem1"`
	Subitem1Title    *Subitem1Title   `xml:"Subitem1Title,omitempty"`
	Subitem1Sentence Subitem1Sentence `xml:"Subitem1Sentence"`
	Subitem2         []Subitem2       `xml:"Subitem2,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem1Title struct {
	XMLName xml.Name `xml:"Subitem1Title"`
	Content string   `xml:",chardata"`
}

type Subitem1Sentence struct {
	XMLName  xml.Name   `xml:"Subitem1Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem2 struct {
	XMLName          xml.Name         `xml:"Subitem2"`
	Subitem2Title    *Subitem2Title   `xml:"Subitem2Title,omitempty"`
	Subitem2Sentence Subitem2Sentence `xml:"Subitem2Sentence"`
	Subitem3         []Subitem3       `xml:"Subitem3,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem2Title struct {
	XMLName xml.Name `xml:"Subitem2Title"`
	Content string   `xml:",chardata"`
}

type Subitem2Sentence struct {
	XMLName  xml.Name   `xml:"Subitem2Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem3 struct {
	XMLName          xml.Name         `xml:"Subitem3"`
	Subitem3Title    *Subitem3Title   `xml:"Subitem3Title,omitempty"`
	Subitem3Sentence Subitem3Sentence `xml:"Subitem3Sentence"`
	Subitem4         []Subitem4       `xml:"Subitem4,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem3Title struct {
	XMLName xml.Name `xml:"Subitem3Title"`
	Content string   `xml:",chardata"`
}

type Subitem3Sentence struct {
	XMLName  xml.Name   `xml:"Subitem3Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem4 struct {
	XMLName          xml.Name         `xml:"Subitem4"`
	Subitem4Title    *Subitem4Title   `xml:"Subitem4Title,omitempty"`
	Subitem4Sentence Subitem4Sentence `xml:"Subitem4Sentence"`
	Subitem5         []Subitem5       `xml:"Subitem5,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem4Title struct {
	XMLName xml.Name `xml:"Subitem4Title"`
	Content string   `xml:",chardata"`
}

type Subitem4Sentence struct {
	XMLName  xml.Name   `xml:"Subitem4Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem5 struct {
	XMLName          xml.Name         `xml:"Subitem5"`
	Subitem5Title    *Subitem5Title   `xml:"Subitem5Title,omitempty"`
	Subitem5Sentence Subitem5Sentence `xml:"Subitem5Sentence"`
	Subitem6         []Subitem6       `xml:"Subitem6,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem5Title struct {
	XMLName xml.Name `xml:"Subitem5Title"`
	Content string   `xml:",chardata"`
}

type Subitem5Sentence struct {
	XMLName  xml.Name   `xml:"Subitem5Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem6 struct {
	XMLName          xml.Name         `xml:"Subitem6"`
	Subitem6Title    *Subitem6Title   `xml:"Subitem6Title,omitempty"`
	Subitem6Sentence Subitem6Sentence `xml:"Subitem6Sentence"`
	Subitem7         []Subitem7       `xml:"Subitem7,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem6Title struct {
	XMLName xml.Name `xml:"Subitem6Title"`
	Content string   `xml:",chardata"`
}

type Subitem6Sentence struct {
	XMLName  xml.Name   `xml:"Subitem6Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem7 struct {
	XMLName          xml.Name         `xml:"Subitem7"`
	Subitem7Title    *Subitem7Title   `xml:"Subitem7Title,omitempty"`
	Subitem7Sentence Subitem7Sentence `xml:"Subitem7Sentence"`
	Subitem8         []Subitem8       `xml:"Subitem8,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem7Title struct {
	XMLName xml.Name `xml:"Subitem7Title"`
	Content string   `xml:",chardata"`
}

type Subitem7Sentence struct {
	XMLName  xml.Name   `xml:"Subitem7Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem8 struct {
	XMLName          xml.Name         `xml:"Subitem8"`
	Subitem8Title    *Subitem8Title   `xml:"Subitem8Title,omitempty"`
	Subitem8Sentence Subitem8Sentence `xml:"Subitem8Sentence"`
	Subitem9         []Subitem9       `xml:"Subitem9,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem8Title struct {
	XMLName xml.Name `xml:"Subitem8Title"`
	Content string   `xml:",chardata"`
}

type Subitem8Sentence struct {
	XMLName  xml.Name   `xml:"Subitem8Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem9 struct {
	XMLName          xml.Name         `xml:"Subitem9"`
	Subitem9Title    *Subitem9Title   `xml:"Subitem9Title,omitempty"`
	Subitem9Sentence Subitem9Sentence `xml:"Subitem9Sentence"`
	Subitem10        []Subitem10      `xml:"Subitem10,omitempty"`
	TableStruct      []TableStruct    `xml:"TableStruct,omitempty"`
	FigStruct        []FigStruct      `xml:"FigStruct,omitempty"`
	StyleStruct      []StyleStruct    `xml:"StyleStruct,omitempty"`
	List             []List           `xml:"List,omitempty"`
	Num              int              `xml:"Num,attr"`
	Delete           bool             `xml:"Delete,attr,omitempty"`
	Hide             bool             `xml:"Hide,attr,omitempty"`
}

type Subitem9Title struct {
	XMLName xml.Name `xml:"Subitem9Title"`
	Content string   `xml:",chardata"`
}

type Subitem9Sentence struct {
	XMLName  xml.Name   `xml:"Subitem9Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Subitem10 struct {
	XMLName           xml.Name          `xml:"Subitem10"`
	Subitem10Title    *Subitem10Title   `xml:"Subitem10Title,omitempty"`
	Subitem10Sentence Subitem10Sentence `xml:"Subitem10Sentence"`
	TableStruct       []TableStruct     `xml:"TableStruct,omitempty"`
	FigStruct         []FigStruct       `xml:"FigStruct,omitempty"`
	StyleStruct       []StyleStruct     `xml:"StyleStruct,omitempty"`
	List              []List            `xml:"List,omitempty"`
	Num               int               `xml:"Num,attr"`
	Delete            bool              `xml:"Delete,attr,omitempty"`
	Hide              bool              `xml:"Hide,attr,omitempty"`
}

type Subitem10Title struct {
	XMLName xml.Name `xml:"Subitem10Title"`
	Content string   `xml:",chardata"`
}

type Subitem10Sentence struct {
	XMLName  xml.Name   `xml:"Subitem10Sentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type List struct {
	XMLName      xml.Name     `xml:"List"`
	ListSentence ListSentence `xml:"ListSentence"`
	Num          int          `xml:"Num,attr"`
	Delete       bool         `xml:"Delete,attr,omitempty"`
	Hide         bool         `xml:"Hide,attr,omitempty"`
}

type ListSentence struct {
	XMLName  xml.Name   `xml:"ListSentence"`
	Sentence []Sentence `xml:"Sentence,omitempty"`
	Column   []Column   `xml:"Column,omitempty"`
	Table    *Table     `xml:"Table,omitempty"`
}

type Sentence struct {
	XMLName     xml.Name `xml:"Sentence"`
	WritingMode string   `xml:"WritingMode,attr,omitempty"`
	Num         int      `xml:"Num,attr"`
	CharData    string   `xml:",chardata"`
}

type TableStruct struct {
	XMLName xml.Name `xml:"TableStruct"`
	Table   Table    `xml:"Table"`
}

type Table struct {
	XMLName     xml.Name   `xml:"Table"`
	WritingMode string     `xml:"WritingMode,attr,omitempty"`
	TableRow    []TableRow `xml:"TableRow"`
}

type TableRow struct {
	XMLName     xml.Name      `xml:"TableRow"`
	TableColumn []TableColumn `xml:"TableColumn"`
}

type TableColumn struct {
	XMLName      xml.Name   `xml:"TableColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type FigStruct struct {
	XMLName xml.Name `xml:"FigStruct"`
	Fig     Fig      `xml:"Fig"`
}

type Fig struct {
	XMLName     xml.Name `xml:"Fig"`
	WritingMode string   `xml:"WritingMode,attr,omitempty"`
	FigRow      []FigRow `xml:"FigRow"`
}

type FigRow struct {
	XMLName   xml.Name    `xml:"FigRow"`
	FigColumn []FigColumn `xml:"FigColumn"`
}

type FigColumn struct {
	XMLName      xml.Name   `xml:"FigColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type StyleStruct struct {
	XMLName xml.Name `xml:"StyleStruct"`
	Style   Style    `xml:"Style"`
}

type Style struct {
	XMLName     xml.Name   `xml:"Style"`
	WritingMode string     `xml:"WritingMode,attr,omitempty"`
	StyleRow    []StyleRow `xml:"StyleRow"`
}

type StyleRow struct {
	XMLName     xml.Name      `xml:"StyleRow"`
	StyleColumn []StyleColumn `xml:"StyleColumn"`
}

type StyleColumn struct {
	XMLName      xml.Name   `xml:"StyleColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type FormatStruct struct {
	XMLName xml.Name `xml:"FormatStruct"`
	Format  Format   `xml:"Format"`
}

type Format struct {
	XMLName     xml.Name    `xml:"Format"`
	WritingMode string      `xml:"WritingMode,attr,omitempty"`
	FormatRow   []FormatRow `xml:"FormatRow"`
}

type FormatRow struct {
	XMLName      xml.Name       `xml:"FormatRow"`
	FormatColumn []FormatColumn `xml:"FormatColumn"`
}

type FormatColumn struct {
	XMLName      xml.Name   `xml:"FormatColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type Remarks struct {
	XMLName xml.Name `xml:"Remarks"`
	Remark  Remark   `xml:"Remark"`
}

type Remark struct {
	XMLName     xml.Name    `xml:"Remark"`
	WritingMode string      `xml:"WritingMode,attr,omitempty"`
	RemarkRow   []RemarkRow `xml:"RemarkRow"`
}

type RemarkRow struct {
	XMLName      xml.Name       `xml:"RemarkRow"`
	RemarkColumn []RemarkColumn `xml:"RemarkColumn"`
}

type RemarkColumn struct {
	XMLName      xml.Name   `xml:"RemarkColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type AppdxTable struct {
	XMLName           xml.Name        `xml:"AppdxTable"`
	Num               int             `xml:"Num,attr"`
	AppdxTableTitle   AppdxTableTitle `xml:"AppdxTableTitle"`
	RelatedArticleNum string          `xml:"RelatedArticleNum"`
	TableStruct       TableStruct     `xml:"TableStruct"`
}

type AppdxTableTitle struct {
	XMLName     xml.Name `xml:"AppdxTableTitle"`
	WritingMode string   `xml:"WritingMode,attr,omitempty"`
	CharData    string   `xml:",chardata"`
}

type AppdxNote struct {
	XMLName xml.Name `xml:"AppdxNote"`
	Note    Note     `xml:"Note"`
}

type Note struct {
	XMLName     xml.Name  `xml:"Note"`
	WritingMode string    `xml:"WritingMode,attr,omitempty"`
	NoteRow     []NoteRow `xml:"NoteRow"`
}

type NoteRow struct {
	XMLName    xml.Name     `xml:"NoteRow"`
	NoteColumn []NoteColumn `xml:"NoteColumn"`
}

type NoteColumn struct {
	XMLName      xml.Name   `xml:"NoteColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type AppdxStyle struct {
	XMLName xml.Name `xml:"AppdxStyle"`
	Style   Style    `xml:"Style"`
}

type Appdx struct {
	XMLName  xml.Name `xml:"Appdx"`
	Appendix Appendix `xml:"Appendix"`
}

type Appendix struct {
	XMLName     xml.Name      `xml:"Appendix"`
	WritingMode string        `xml:"WritingMode,attr,omitempty"`
	AppendixRow []AppendixRow `xml:"AppendixRow"`
}

type AppendixRow struct {
	XMLName        xml.Name         `xml:"AppendixRow"`
	AppendixColumn []AppendixColumn `xml:"AppendixColumn"`
}

type AppendixColumn struct {
	XMLName      xml.Name   `xml:"AppendixColumn"`
	BorderBottom string     `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string     `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string     `xml:"BorderRight,attr,omitempty"`
	BorderTop    string     `xml:"BorderTop,attr,omitempty"`
	Sentence     []Sentence `xml:"Sentence"`
}

type AppdxFig struct {
	XMLName xml.Name `xml:"AppdxFig"`
	Fig     Fig      `xml:"Fig"`
}

type AppdxFormat struct {
	XMLName xml.Name `xml:"AppdxFormat"`
	Format  Format   `xml:"Format"`
}

type SupplProvision struct {
	XMLName             xml.Name    `xml:"SupplProvision"`
	AmendLawNum         string      `xml:"AmendLawNum,attr"`
	Extract             *bool       `xml:"Extract,attr,omitempty"`
	Article             []Article   `xml:"Article,omitempty"`
	Paragraph           []Paragraph `xml:"Paragraph,omitempty"`
	SupplProvisionLabel string      `xml:"SupplProvisionLabel"`
}

type SupplProvisionAppdxStyle struct {
	XMLName xml.Name `xml:"SupplProvisionAppdxStyle"`
	Style   Style    `xml:"Style"`
}

type SupplProvisionAppdxTable struct {
	XMLName xml.Name `xml:"SupplProvisionAppdxTable"`
	Table   Table    `xml:"Table"`
}

type SupplProvisionAppdx struct {
	XMLName  xml.Name `xml:"SupplProvisionAppdx"`
	Appendix Appendix `xml:"Appendix"`
}

type SupplProvisionLabel struct {
	XMLName xml.Name `xml:"SupplProvisionLabel"`
	Content string   `xml:",chardata"`
}

type NoteStruct struct {
	XMLName         xml.Name  `xml:"NoteStruct"`
	NoteStructTitle *string   `xml:"NoteStructTitle,omitempty"`
	Remarks         []Remarks `xml:"Remarks,omitempty"`
	Note            Note      `xml:"Note"`
}
