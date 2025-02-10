package jplaw

// Law ...
type Law struct {
	Era             Era      `xml:"Era,attr"`
	Year            int      `xml:"Year,attr"`
	Num             int      `xml:"Num,attr"`
	PromulgateMonth int      `xml:"PromulgateMonth,attr,omitempty"`
	PromulgateDay   int      `xml:"PromulgateDay,attr,omitempty"`
	LawType         LawType  `xml:"LawType,attr"`
	Lang            Language `xml:"Lang,attr"`
	LawNum          LawNum   `xml:"LawNum"`
	Law             string   `xml:"Law"`
	LawBody         *LawBody `xml:"LawBody"`
}

// LawNum ...
type LawNum string

// LawBody ...
type LawBody struct {
	Subject        string            `xml:"Subject,attr,omitempty"`
	LawTitle       *LawTitle         `xml:"LawTitle"`
	EnactStatement []*EnactStatement `xml:"EnactStatement"`
	TOC            *TOC              `xml:"TOC"`
	Preamble       *Preamble         `xml:"Preamble"`
	MainProvision  *MainProvision    `xml:"MainProvision"`
	SupplProvision []string          `xml:"SupplProvision"`
	AppdxTable     []*AppdxTable     `xml:"AppdxTable"`
	AppdxNote      []*AppdxNote      `xml:"AppdxNote"`
	AppdxStyle     []*AppdxStyle     `xml:"AppdxStyle"`
	Appdx          []*Appdx          `xml:"Appdx"`
	AppdxFig       []*AppdxFig       `xml:"AppdxFig"`
	AppdxFormat    []*AppdxFormat    `xml:"AppdxFormat"`
}

// LawTitle ...
type LawTitle struct {
	Kana       string `xml:"Kana,attr,omitempty"`
	Abbrev     string `xml:"Abbrev,attr,omitempty"`
	AbbrevKana string `xml:"AbbrevKana,attr,omitempty"`
	Content    string `xml:",innerxml"`
}

// EnactStatement ...
type EnactStatement struct {
	Content string `xml:",innerxml"`
}

// TOC ...
type TOC struct {
	TOCLabel           *TOCLabel             `xml:"TOCLabel"`
	TOCPreambleLabel   *TOCPreambleLabel     `xml:"TOCPreambleLabel"`
	TOCPart            []*TOCPart            `xml:"TOCPart"`
	TOCChapter         []*TOCChapter         `xml:"TOCChapter"`
	TOCSection         []*TOCSection         `xml:"TOCSection"`
	TOCArticle         []*TOCArticle         `xml:"TOCArticle"`
	TOCSupplProvision  *TOCSupplProvision    `xml:"TOCSupplProvision"`
	TOCAppdxTableLabel []*TOCAppdxTableLabel `xml:"TOCAppdxTableLabel"`
}

// TOCLabel ...
type TOCLabel struct {
	Content string `xml:",innerxml"`
}

// TOCPreambleLabel ...
type TOCPreambleLabel struct {
	Content string `xml:",innerxml"`
}

// TOCPart ...
type TOCPart struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	PartTitle    *PartTitle    `xml:"PartTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange"`
	TOCChapter   []*TOCChapter `xml:"TOCChapter"`
}

// TOCChapter ...
type TOCChapter struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	ChapterTitle *ChapterTitle `xml:"ChapterTitle"`
	ArticleRange *ArticleRange `xml:"ArticleRange"`
	TOCSection   []*TOCSection `xml:"TOCSection"`
}

// TOCSection ...
type TOCSection struct {
	Num           string           `xml:"Num,attr"`
	Delete        bool             `xml:"Delete,attr,omitempty"`
	SectionTitle  *SectionTitle    `xml:"SectionTitle"`
	ArticleRange  *ArticleRange    `xml:"ArticleRange"`
	TOCSubsection []*TOCSubsection `xml:"TOCSubsection"`
	TOCDivision   []*TOCDivision   `xml:"TOCDivision"`
}

// TOCSubsection ...
type TOCSubsection struct {
	Num             string           `xml:"Num,attr"`
	Delete          bool             `xml:"Delete,attr,omitempty"`
	SubsectionTitle *SubsectionTitle `xml:"SubsectionTitle"`
	ArticleRange    *ArticleRange    `xml:"ArticleRange"`
	TOCDivision     []*TOCDivision   `xml:"TOCDivision"`
}

// TOCDivision ...
type TOCDivision struct {
	Num           string         `xml:"Num,attr"`
	Delete        bool           `xml:"Delete,attr,omitempty"`
	DivisionTitle *DivisionTitle `xml:"DivisionTitle"`
	ArticleRange  *ArticleRange  `xml:"ArticleRange"`
}

// TOCArticle ...
type TOCArticle struct {
	Num            string          `xml:"Num,attr"`
	Delete         bool            `xml:"Delete,attr,omitempty"`
	ArticleTitle   *ArticleTitle   `xml:"ArticleTitle"`
	ArticleCaption *ArticleCaption `xml:"ArticleCaption"`
}

// TOCSupplProvision ...
type TOCSupplProvision struct {
	SupplProvisionLabel *SupplProvisionLabel `xml:"SupplProvisionLabel"`
	ArticleRange        *ArticleRange        `xml:"ArticleRange"`
	TOCArticle          []*TOCArticle        `xml:"TOCArticle"`
	TOCChapter          []*TOCChapter        `xml:"TOCChapter"`
}

// TOCAppdxTableLabel ...
type TOCAppdxTableLabel struct {
	Content string `xml:",innerxml"`
}

// ArticleRange ...
type ArticleRange struct {
	Content string `xml:",innerxml"`
}

// Preamble ...
type Preamble struct {
	Paragraph []*Paragraph `xml:"Paragraph"`
}

// MainProvision ...
type MainProvision struct {
	Extract   bool         `xml:"Extract,attr,omitempty"`
	Part      []*Part      `xml:"Part"`
	Chapter   []*Chapter   `xml:"Chapter"`
	Section   []*Section   `xml:"Section"`
	Article   []*Article   `xml:"Article"`
	Paragraph []*Paragraph `xml:"Paragraph"`
}

// Part ...
type Part struct {
	Num       string     `xml:"Num,attr"`
	Delete    bool       `xml:"Delete,attr,omitempty"`
	Hide      bool       `xml:"Hide,attr,omitempty"`
	PartTitle *PartTitle `xml:"PartTitle"`
	Article   []*Article `xml:"Article"`
	Chapter   []*Chapter `xml:"Chapter"`
}

// PartTitle ...
type PartTitle struct {
	Content string `xml:",innerxml"`
}

// Chapter ...
type Chapter struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	Hide         bool          `xml:"Hide,attr,omitempty"`
	ChapterTitle *ChapterTitle `xml:"ChapterTitle"`
	Article      []*Article    `xml:"Article"`
	Section      []*Section    `xml:"Section"`
}

// ChapterTitle ...
type ChapterTitle struct {
	Content string `xml:",innerxml"`
}

// Section ...
type Section struct {
	Num          string        `xml:"Num,attr"`
	Delete       bool          `xml:"Delete,attr,omitempty"`
	Hide         bool          `xml:"Hide,attr,omitempty"`
	SectionTitle *SectionTitle `xml:"SectionTitle"`
	Article      []*Article    `xml:"Article"`
	Subsection   []*Subsection `xml:"Subsection"`
	Division     []*Division   `xml:"Division"`
}

// SectionTitle ...
type SectionTitle struct {
	Content string `xml:",innerxml"`
}

// Subsection ...
type Subsection struct {
	Num             string           `xml:"Num,attr"`
	Delete          bool             `xml:"Delete,attr,omitempty"`
	Hide            bool             `xml:"Hide,attr,omitempty"`
	SubsectionTitle *SubsectionTitle `xml:"SubsectionTitle"`
	Article         []*Article       `xml:"Article"`
	Division        []*Division      `xml:"Division"`
}

// SubsectionTitle ...
type SubsectionTitle struct {
	Content string `xml:",innerxml"`
}

// Division ...
type Division struct {
	Num           string         `xml:"Num,attr"`
	Delete        bool           `xml:"Delete,attr,omitempty"`
	Hide          bool           `xml:"Hide,attr,omitempty"`
	DivisionTitle *DivisionTitle `xml:"DivisionTitle"`
	Article       []*Article     `xml:"Article"`
}

// DivisionTitle ...
type DivisionTitle struct {
	Content string `xml:",innerxml"`
}

// Article ...
type Article struct {
	Num            string          `xml:"Num,attr"`
	Delete         bool            `xml:"Delete,attr,omitempty"`
	Hide           bool            `xml:"Hide,attr,omitempty"`
	ArticleCaption *ArticleCaption `xml:"ArticleCaption"`
	ArticleTitle   *ArticleTitle   `xml:"ArticleTitle"`
	Paragraph      []*Paragraph    `xml:"Paragraph"`
	SupplNote      *SupplNote      `xml:"SupplNote"`
}

// ArticleTitle ...
type ArticleTitle struct {
	Content string `xml:",innerxml"`
}

// ArticleCaption ...
type ArticleCaption struct {
	CommonCaption bool   `xml:"CommonCaption,attr,omitempty"`
	Content       string `xml:",innerxml"`
}

// Paragraph ...
type Paragraph struct {
	Num               int                `xml:"Num,attr"`
	OldStyle          bool               `xml:"OldStyle,attr,omitempty"`
	OldNum            bool               `xml:"OldNum,attr,omitempty"`
	Hide              bool               `xml:"Hide,attr,omitempty"`
	ParagraphCaption  *ParagraphCaption  `xml:"ParagraphCaption"`
	ParagraphNum      *ParagraphNum      `xml:"ParagraphNum"`
	ParagraphSentence *ParagraphSentence `xml:"ParagraphSentence"`
	AmendProvision    []*AmendProvision  `xml:"AmendProvision"`
	Class             []*Class           `xml:"Class"`
	TableStruct       []*TableStruct     `xml:"TableStruct"`
	FigStruct         []*FigStruct       `xml:"FigStruct"`
	StyleStruct       []*StyleStruct     `xml:"StyleStruct"`
	Item              []*Item            `xml:"Item"`
	List              []*List            `xml:"List"`
}

// ParagraphCaption ...
type ParagraphCaption struct {
	CommonCaption bool   `xml:"CommonCaption,attr,omitempty"`
	Content       string `xml:",innerxml"`
}

// ParagraphNum ...
type ParagraphNum struct {
	Content string `xml:",innerxml"`
}

// ParagraphSentence ...
type ParagraphSentence struct {
	Sentence []Sentence `xml:"Sentence"`
}

// SupplNote ...
type SupplNote struct {
	Content string `xml:",innerxml"`
}

// AmendProvision ...
type AmendProvision struct {
	AmendProvisionSentence *AmendProvisionSentence `xml:"AmendProvisionSentence"`
	NewProvision           []*NewProvision         `xml:"NewProvision"`
}

// AmendProvisionSentence ...
type AmendProvisionSentence struct {
	Sentence string `xml:"Sentence"`
}

// NewProvision ...
type NewProvision struct {
	LawTitle                 *LawTitle                   `xml:"LawTitle"`
	Preamble                 *Preamble                   `xml:"Preamble"`
	TOC                      *TOC                        `xml:"TOC"`
	Part                     []*Part                     `xml:"Part"`
	PartTitle                []*PartTitle                `xml:"PartTitle"`
	Chapter                  []*Chapter                  `xml:"Chapter"`
	ChapterTitle             []*ChapterTitle             `xml:"ChapterTitle"`
	Section                  []*Section                  `xml:"Section"`
	SectionTitle             []*SectionTitle             `xml:"SectionTitle"`
	Subsection               []*Subsection               `xml:"Subsection"`
	SubsectionTitle          []*SubsectionTitle          `xml:"SubsectionTitle"`
	Division                 []*Division                 `xml:"Division"`
	DivisionTitle            []*DivisionTitle            `xml:"DivisionTitle"`
	Article                  []*Article                  `xml:"Article"`
	SupplNote                []*SupplNote                `xml:"SupplNote"`
	Paragraph                []*Paragraph                `xml:"Paragraph"`
	Item                     []*Item                     `xml:"Item"`
	Subitem1                 []*Subitem1                 `xml:"Subitem1"`
	Subitem2                 []*Subitem2                 `xml:"Subitem2"`
	Subitem3                 []*Subitem3                 `xml:"Subitem3"`
	Subitem4                 []*Subitem4                 `xml:"Subitem4"`
	Subitem5                 []*Subitem5                 `xml:"Subitem5"`
	Subitem6                 []*Subitem6                 `xml:"Subitem6"`
	Subitem7                 []*Subitem7                 `xml:"Subitem7"`
	Subitem8                 []*Subitem8                 `xml:"Subitem8"`
	Subitem9                 []*Subitem9                 `xml:"Subitem9"`
	Subitem10                []*Subitem10                `xml:"Subitem10"`
	List                     []*List                     `xml:"List"`
	Sentence                 []string                    `xml:"Sentence"`
	AmendProvision           []*AmendProvision           `xml:"AmendProvision"`
	AppdxTable               []*AppdxTable               `xml:"AppdxTable"`
	AppdxNote                []*AppdxNote                `xml:"AppdxNote"`
	AppdxStyle               []*AppdxStyle               `xml:"AppdxStyle"`
	Appdx                    []*Appdx                    `xml:"Appdx"`
	AppdxFig                 []*AppdxFig                 `xml:"AppdxFig"`
	AppdxFormat              []*AppdxFormat              `xml:"AppdxFormat"`
	SupplProvisionAppdxStyle []*SupplProvisionAppdxStyle `xml:"SupplProvisionAppdxStyle"`
	SupplProvisionAppdxTable []*SupplProvisionAppdxTable `xml:"SupplProvisionAppdxTable"`
	SupplProvisionAppdx      []*SupplProvisionAppdx      `xml:"SupplProvisionAppdx"`
	TableStruct              *TableStruct                `xml:"TableStruct"`
	TableRow                 []*TableRow                 `xml:"TableRow"`
	TableColumn              []string                    `xml:"TableColumn"`
	FigStruct                *FigStruct                  `xml:"FigStruct"`
	NoteStruct               *NoteStruct                 `xml:"NoteStruct"`
	StyleStruct              *StyleStruct                `xml:"StyleStruct"`
	FormatStruct             *FormatStruct               `xml:"FormatStruct"`
	Remarks                  *Remarks                    `xml:"Remarks"`
	LawBody                  *LawBody                    `xml:"LawBody"`
}

// Class ...
type Class struct {
	Num           string         `xml:"Num,attr"`
	ClassTitle    *ClassTitle    `xml:"ClassTitle"`
	ClassSentence *ClassSentence `xml:"ClassSentence"`
	Item          []*Item        `xml:"Item"`
}

// ClassTitle ...
type ClassTitle struct {
	Content string `xml:",innerxml"`
}

// ClassSentence ...
type ClassSentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Item ...
type Item struct {
	Num          string         `xml:"Num,attr"`
	Delete       bool           `xml:"Delete,attr,omitempty"`
	Hide         bool           `xml:"Hide,attr,omitempty"`
	ItemTitle    *ItemTitle     `xml:"ItemTitle"`
	ItemSentence *ItemSentence  `xml:"ItemSentence"`
	Subitem1     []*Subitem1    `xml:"Subitem1"`
	TableStruct  []*TableStruct `xml:"TableStruct"`
	FigStruct    []*FigStruct   `xml:"FigStruct"`
	StyleStruct  []*StyleStruct `xml:"StyleStruct"`
	List         []*List        `xml:"List"`
}

// ItemTitle ...
type ItemTitle struct {
	Content string `xml:",innerxml"`
}

// ItemSentence ...
type ItemSentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem1 ...
type Subitem1 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem1Title    *Subitem1Title    `xml:"Subitem1Title"`
	Subitem1Sentence *Subitem1Sentence `xml:"Subitem1Sentence"`
	Subitem2         []*Subitem2       `xml:"Subitem2"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem1Title ...
type Subitem1Title struct {
	Content string `xml:",innerxml"`
}

// Subitem1Sentence ...
type Subitem1Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem2 ...
type Subitem2 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem2Title    *Subitem2Title    `xml:"Subitem2Title"`
	Subitem2Sentence *Subitem2Sentence `xml:"Subitem2Sentence"`
	Subitem3         []*Subitem3       `xml:"Subitem3"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem2Title ...
type Subitem2Title struct {
	Content string `xml:",innerxml"`
}

// Subitem2Sentence ...
type Subitem2Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem3 ...
type Subitem3 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem3Title    *Subitem3Title    `xml:"Subitem3Title"`
	Subitem3Sentence *Subitem3Sentence `xml:"Subitem3Sentence"`
	Subitem4         []*Subitem4       `xml:"Subitem4"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem3Title ...
type Subitem3Title struct {
	Content string `xml:",innerxml"`
}

// Subitem3Sentence ...
type Subitem3Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem4 ...
type Subitem4 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem4Title    *Subitem4Title    `xml:"Subitem4Title"`
	Subitem4Sentence *Subitem4Sentence `xml:"Subitem4Sentence"`
	Subitem5         []*Subitem5       `xml:"Subitem5"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem4Title ...
type Subitem4Title struct {
	Content string `xml:",innerxml"`
}

// Subitem4Sentence ...
type Subitem4Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem5 ...
type Subitem5 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem5Title    *Subitem5Title    `xml:"Subitem5Title"`
	Subitem5Sentence *Subitem5Sentence `xml:"Subitem5Sentence"`
	Subitem6         []*Subitem6       `xml:"Subitem6"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem5Title ...
type Subitem5Title struct {
	Content string `xml:",innerxml"`
}

// Subitem5Sentence ...
type Subitem5Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem6 ...
type Subitem6 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem6Title    *Subitem6Title    `xml:"Subitem6Title"`
	Subitem6Sentence *Subitem6Sentence `xml:"Subitem6Sentence"`
	Subitem7         []*Subitem7       `xml:"Subitem7"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem6Title ...
type Subitem6Title struct {
	Content string `xml:",innerxml"`
}

// Subitem6Sentence ...
type Subitem6Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem7 ...
type Subitem7 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem7Title    *Subitem7Title    `xml:"Subitem7Title"`
	Subitem7Sentence *Subitem7Sentence `xml:"Subitem7Sentence"`
	Subitem8         []*Subitem8       `xml:"Subitem8"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem7Title ...
type Subitem7Title struct {
	Content string `xml:",innerxml"`
}

// Subitem7Sentence ...
type Subitem7Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem8 ...
type Subitem8 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem8Title    *Subitem8Title    `xml:"Subitem8Title"`
	Subitem8Sentence *Subitem8Sentence `xml:"Subitem8Sentence"`
	Subitem9         []*Subitem9       `xml:"Subitem9"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem8Title ...
type Subitem8Title struct {
	Content string `xml:",innerxml"`
}

// Subitem8Sentence ...
type Subitem8Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem9 ...
type Subitem9 struct {
	Num              string            `xml:"Num,attr"`
	Delete           bool              `xml:"Delete,attr,omitempty"`
	Hide             bool              `xml:"Hide,attr,omitempty"`
	Subitem9Title    *Subitem9Title    `xml:"Subitem9Title"`
	Subitem9Sentence *Subitem9Sentence `xml:"Subitem9Sentence"`
	Subitem10        []*Subitem10      `xml:"Subitem10"`
	TableStruct      []*TableStruct    `xml:"TableStruct"`
	FigStruct        []*FigStruct      `xml:"FigStruct"`
	StyleStruct      []*StyleStruct    `xml:"StyleStruct"`
	List             []*List           `xml:"List"`
}

// Subitem9Title ...
type Subitem9Title struct {
	Content string `xml:",innerxml"`
}

// Subitem9Sentence ...
type Subitem9Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Subitem10 ...
type Subitem10 struct {
	Num               string             `xml:"Num,attr"`
	Delete            bool               `xml:"Delete,attr,omitempty"`
	Hide              bool               `xml:"Hide,attr,omitempty"`
	Subitem10Title    *Subitem10Title    `xml:"Subitem10Title"`
	Subitem10Sentence *Subitem10Sentence `xml:"Subitem10Sentence"`
	TableStruct       []*TableStruct     `xml:"TableStruct"`
	FigStruct         []*FigStruct       `xml:"FigStruct"`
	StyleStruct       []*StyleStruct     `xml:"StyleStruct"`
	List              []*List            `xml:"List"`
}

// Subitem10Title ...
type Subitem10Title struct {
	Content string `xml:",innerxml"`
}

// Subitem10Sentence ...
type Subitem10Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
	Table    string   `xml:"Table"`
}

// Sentence ...
type Sentence struct {
	Num         int    `xml:"Num,attr,omitempty"`
	Function    string `xml:"Function,attr,omitempty"`
	Indent      string `xml:"Indent,attr,omitempty"`
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// Column ...
type Column struct {
	Num       int    `xml:"Num,attr,omitempty"`
	LineBreak bool   `xml:"LineBreak,attr,omitempty"`
	Align     string `xml:"Align,attr,omitempty"`
	Column    string `xml:"Column"`
}

// SupplProvision ...
type SupplProvision struct {
	Type                     string                      `xml:"Type,attr,omitempty"`
	AmendLawNum              string                      `xml:"AmendLawNum,attr,omitempty"`
	Extract                  bool                        `xml:"Extract,attr,omitempty"`
	SupplProvisionLabel      *SupplProvisionLabel        `xml:"SupplProvisionLabel"`
	Chapter                  []*Chapter                  `xml:"Chapter"`
	Article                  []*Article                  `xml:"Article"`
	Paragraph                []*Paragraph                `xml:"Paragraph"`
	SupplProvisionAppdxTable []*SupplProvisionAppdxTable `xml:"SupplProvisionAppdxTable"`
	SupplProvisionAppdxStyle []*SupplProvisionAppdxStyle `xml:"SupplProvisionAppdxStyle"`
	SupplProvision           string                      `xml:"SupplProvision"`
}

// SupplProvisionLabel ...
type SupplProvisionLabel struct {
	Content string `xml:",innerxml"`
}

// SupplProvisionAppdxTable ...
type SupplProvisionAppdxTable struct {
	Num                           int                           `xml:"Num,attr,omitempty"`
	SupplProvisionAppdxTableTitle SupplProvisionAppdxTableTitle `xml:"SupplProvisionAppdxTableTitle"`
	RelatedArticleNum             *RelatedArticleNum            `xml:"RelatedArticleNum"`
	TableStruct                   []*TableStruct                `xml:"TableStruct"`
}

// SupplProvisionAppdxTableTitle ...
type SupplProvisionAppdxTableTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// SupplProvisionAppdxStyle ...
type SupplProvisionAppdxStyle struct {
	Num                           int                           `xml:"Num,attr,omitempty"`
	SupplProvisionAppdxStyleTitle SupplProvisionAppdxStyleTitle `xml:"SupplProvisionAppdxStyleTitle"`
	RelatedArticleNum             *RelatedArticleNum            `xml:"RelatedArticleNum"`
	StyleStruct                   []*StyleStruct                `xml:"StyleStruct"`
}

// SupplProvisionAppdxStyleTitle ...
type SupplProvisionAppdxStyleTitle struct {
	WritingMode                   string `xml:"WritingMode,attr,omitempty"`
	Content                       string `xml:",innerxml"`
	SupplProvisionAppdxStyleTitle string `xml:"SupplProvisionAppdxStyleTitle"`
}

// SupplProvisionAppdx ...
type SupplProvisionAppdx struct {
	Num               int                `xml:"Num,attr,omitempty"`
	ArithFormulaNum   *ArithFormulaNum   `xml:"ArithFormulaNum"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	ArithFormula      []*ArithFormula    `xml:"ArithFormula"`
}

// AppdxTable ...
type AppdxTable struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxTableTitle   AppdxTableTitle    `xml:"AppdxTableTitle"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	TableStruct       []*TableStruct     `xml:"TableStruct"`
	Item              []*Item            `xml:"Item"`
	Remarks           *Remarks           `xml:"Remarks"`
}

// AppdxTableTitle ...
type AppdxTableTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// AppdxNote ...
type AppdxNote struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxNoteTitle    AppdxNoteTitle     `xml:"AppdxNoteTitle"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	NoteStruct        []*NoteStruct      `xml:"NoteStruct"`
	FigStruct         []*FigStruct       `xml:"FigStruct"`
	TableStruct       []*TableStruct     `xml:"TableStruct"`
	Remarks           *Remarks           `xml:"Remarks"`
}

// AppdxNoteTitle ...
type AppdxNoteTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// AppdxStyle ...
type AppdxStyle struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxStyleTitle   AppdxStyleTitle    `xml:"AppdxStyleTitle"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	StyleStruct       []*StyleStruct     `xml:"StyleStruct"`
	Remarks           *Remarks           `xml:"Remarks"`
}

// AppdxStyleTitle ...
type AppdxStyleTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// AppdxFormat ...
type AppdxFormat struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxFormatTitle  AppdxFormatTitle   `xml:"AppdxFormatTitle"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	FormatStruct      []*FormatStruct    `xml:"FormatStruct"`
	Remarks           *Remarks           `xml:"Remarks"`
}

// AppdxFormatTitle ...
type AppdxFormatTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// Appdx ...
type Appdx struct {
	ArithFormulaNum   *ArithFormulaNum   `xml:"ArithFormulaNum"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	ArithFormula      []*ArithFormula    `xml:"ArithFormula"`
	Remarks           *Remarks           `xml:"Remarks"`
}

// ArithFormulaNum ...
type ArithFormulaNum struct {
	Content string `xml:",innerxml"`
}

// ArithFormula ...
type ArithFormula struct {
	Num     int    `xml:"Num,attr,omitempty"`
	Content string `xml:",innerxml"`
}

// AppdxFig ...
type AppdxFig struct {
	Num               int                `xml:"Num,attr,omitempty"`
	AppdxFigTitle     AppdxFigTitle      `xml:"AppdxFigTitle"`
	RelatedArticleNum *RelatedArticleNum `xml:"RelatedArticleNum"`
	FigStruct         []*FigStruct       `xml:"FigStruct"`
	TableStruct       []*TableStruct     `xml:"TableStruct"`
}

// AppdxFigTitle ...
type AppdxFigTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// TableStruct ...
type TableStruct struct {
	TableStructTitle TableStructTitle `xml:"TableStructTitle"`
	Remarks          []*Remarks       `xml:"Remarks"`
	Table            Table            `xml:"Table"`
}

// TableStructTitle ...
type TableStructTitle struct {
	WritingMode string `xml:"WritingMode,attr,omitempty"`
	Content     string `xml:",innerxml"`
}

// Table ...
type Table struct {
	WritingMode    string            `xml:"WritingMode,attr,omitempty"`
	TableHeaderRow []*TableHeaderRow `xml:"TableHeaderRow"`
}

// TableRow ...
type TableRow struct {
	TableColumn []*TableColumn `xml:"TableColumn"`
}

// TableHeaderRow ...
type TableHeaderRow struct {
	TableHeaderColumn []*TableHeaderColumn `xml:"TableHeaderColumn"`
}

// TableHeaderColumn ...
type TableHeaderColumn struct {
	Content string `xml:",innerxml"`
}

// TableColumn ...
type TableColumn struct {
	BorderTop    string        `xml:"BorderTop,attr,omitempty"`
	BorderBottom string        `xml:"BorderBottom,attr,omitempty"`
	BorderLeft   string        `xml:"BorderLeft,attr,omitempty"`
	BorderRight  string        `xml:"BorderRight,attr,omitempty"`
	Rowspan      int           `xml:"rowspan,attr,omitempty"`
	Colspan      int           `xml:"colspan,attr,omitempty"`
	Align        string        `xml:"Align,attr,omitempty"`
	Valign       string        `xml:"Valign,attr,omitempty"`
	Part         []*Part       `xml:"Part"`
	Chapter      []*Chapter    `xml:"Chapter"`
	Section      []*Section    `xml:"Section"`
	Subsection   []*Subsection `xml:"Subsection"`
	Division     []*Division   `xml:"Division"`
	Article      []*Article    `xml:"Article"`
	Paragraph    []*Paragraph  `xml:"Paragraph"`
	Item         []*Item       `xml:"Item"`
	Subitem1     []*Subitem1   `xml:"Subitem1"`
	Subitem2     []*Subitem2   `xml:"Subitem2"`
	Subitem3     []*Subitem3   `xml:"Subitem3"`
	Subitem4     []*Subitem4   `xml:"Subitem4"`
	Subitem5     []*Subitem5   `xml:"Subitem5"`
	Subitem6     []*Subitem6   `xml:"Subitem6"`
	Subitem7     []*Subitem7   `xml:"Subitem7"`
	Subitem8     []*Subitem8   `xml:"Subitem8"`
	Subitem9     []*Subitem9   `xml:"Subitem9"`
	Subitem10    []*Subitem10  `xml:"Subitem10"`
	FigStruct    []*FigStruct  `xml:"FigStruct"`
	Remarks      *Remarks      `xml:"Remarks"`
	Sentence     []string      `xml:"Sentence"`
	TableColumn  string        `xml:"TableColumn"`
}

// FigStruct ...
type FigStruct struct {
	FigStructTitle *FigStructTitle `xml:"FigStructTitle"`
	Remarks        []*Remarks      `xml:"Remarks"`
	Fig            *Fig            `xml:"Fig"`
}

// FigStructTitle ...
type FigStructTitle struct {
	Content string `xml:",innerxml"`
}

// Fig ...
type Fig struct {
	Src string `xml:"src,attr"`
}

// NoteStruct ...
type NoteStruct struct {
	NoteStructTitle *NoteStructTitle `xml:"NoteStructTitle"`
	Remarks         []*Remarks       `xml:"Remarks"`
	Note            *Note            `xml:"Note"`
}

// NoteStructTitle ...
type NoteStructTitle struct {
	Content string `xml:",innerxml"`
}

// Note ...
type Note struct {
	Content string `xml:",innerxml"`
}

// StyleStruct ...
type StyleStruct struct {
	StyleStructTitle *StyleStructTitle `xml:"StyleStructTitle"`
	Remarks          []*Remarks        `xml:"Remarks"`
	Style            *Style            `xml:"Style"`
}

// StyleStructTitle ...
type StyleStructTitle struct {
	Content string `xml:",innerxml"`
}

// Style ...
type Style struct {
	Content string `xml:",innerxml"`
}

// FormatStruct ...
type FormatStruct struct {
	FormatStructTitle *FormatStructTitle `xml:"FormatStructTitle"`
	Remarks           []*Remarks         `xml:"Remarks"`
	Format            *Format            `xml:"Format"`
}

// FormatStructTitle ...
type FormatStructTitle struct {
	Content string `xml:",innerxml"`
}

// Format ...
type Format struct {
	Content string `xml:",innerxml"`
}

// RelatedArticleNum ...
type RelatedArticleNum struct {
	Content string `xml:",innerxml"`
}

// Remarks ...
type Remarks struct {
	RemarksLabel *RemarksLabel `xml:"RemarksLabel"`
	Item         []*Item       `xml:"Item"`
	Sentence     []string      `xml:"Sentence"`
}

// RemarksLabel ...
type RemarksLabel struct {
	LineBreak bool   `xml:"LineBreak,attr,omitempty"`
	Content   string `xml:",innerxml"`
}

// List ...
type List struct {
	ListSentence *ListSentence `xml:"ListSentence"`
	Sublist1     []*Sublist1   `xml:"Sublist1"`
}

// ListSentence ...
type ListSentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
}

// Sublist1 ...
type Sublist1 struct {
	Sublist1Sentence *Sublist1Sentence `xml:"Sublist1Sentence"`
	Sublist2         []*Sublist2       `xml:"Sublist2"`
}

// Sublist1Sentence ...
type Sublist1Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
}

// Sublist2 ...
type Sublist2 struct {
	Sublist2Sentence *Sublist2Sentence   `xml:"Sublist2Sentence"`
	Sublist3         []*Sublist3Sentence `xml:"Sublist3"`
}

// Sublist2Sentence ...
type Sublist2Sentence struct {
	Sentence []string `xml:"Sentence"`
	Column   []string `xml:"Column"`
}

// Sublist3 ...
type Sublist3 *Sublist3Sentence

// Sublist3Sentence ...
type Sublist3Sentence struct {
	Sublist3Sentence *Sublist3Sentence `xml:"Sublist3Sentence"`
}
