package jplaw

type Law struct {
	Era             Era      `xml:"Era,attr"`
	Year            int      `xml:"Year,attr"`
	Num             string   `xml:"Num,attr"`
	PromulgateMonth int      `xml:"PromulgateMonth,attr"`
	PromulgateDay   int      `xml:"PromulgateDay,attr"`
	LawType         LawType  `xml:"LawType,attr"`
	Lang            Language `xml:"Lang,attr"`
	LawBody         struct {
		AppdxTable struct {
			Num             string `xml:"Num,attr"`
			AppdxTableTitle struct {
				WritingMode string `xml:"WritingMode,attr"`
				CharData    string `xml:",chardata"`
			} `xml:"AppdxTableTitle"`
			RelatedArticleNum string `xml:"RelatedArticleNum"`
			TableStruct       struct {
				Table struct {
					WritingMode string `xml:"WritingMode,attr"`
					TableRow    []struct {
						TableColumn []struct {
							BorderBottom string `xml:"BorderBottom,attr"`
							BorderLeft   string `xml:"BorderLeft,attr"`
							BorderRight  string `xml:"BorderRight,attr"`
							BorderTop    string `xml:"BorderTop,attr"`
							Sentence     []struct {
								Num         string `xml:"Num,attr"`
								WritingMode string `xml:"WritingMode,attr"`
								CharData    string `xml:",chardata"`
							} `xml:"Sentence"`
						} `xml:"TableColumn"`
					} `xml:"TableRow"`
				} `xml:"Table"`
			} `xml:"TableStruct"`
		} `xml:"AppdxTable"`
		LawTitle struct {
			Abbrev     string `xml:"Abbrev,attr"`
			AbbrevKana string `xml:"AbbrevKana,attr"`
			Kana       string `xml:"Kana,attr"`
			CharData   string `xml:",chardata"`
		} `xml:"LawTitle"`
		MainProvision struct {
			Chapter []struct {
				Num     string `xml:"Num,attr"`
				Article []struct {
					Num            string `xml:"Num,attr"`
					ArticleCaption *struct {
						CharData string `xml:",chardata"`
						Ruby     []struct {
							CharData string `xml:",chardata"`
							Rt       string `xml:"Rt"`
						} `xml:"Ruby"`
					} `xml:"ArticleCaption"`
					ArticleTitle string `xml:"ArticleTitle"`
					Paragraph    []struct {
						Num  string `xml:"Num,attr"`
						Item []struct {
							Num          float64 `xml:"Num,attr"`
							ItemSentence struct {
								Column []struct {
									Num      string `xml:"Num,attr"`
									Sentence struct {
										Num         string `xml:"Num,attr"`
										WritingMode string `xml:"WritingMode,attr"`
										CharData    string `xml:",chardata"`
									} `xml:"Sentence"`
								} `xml:"Column"`
								Sentence *struct {
									Num         string `xml:"Num,attr"`
									WritingMode string `xml:"WritingMode,attr"`
									CharData    string `xml:",chardata"`
								} `xml:"Sentence"`
							} `xml:"ItemSentence"`
							ItemTitle string `xml:"ItemTitle"`
							Subitem1  []struct {
								Num              string `xml:"Num,attr"`
								Subitem1Sentence struct {
									Sentence struct {
										Num         string `xml:"Num,attr"`
										WritingMode string `xml:"WritingMode,attr"`
										CharData    string `xml:",chardata"`
										Ruby        *struct {
											CharData string `xml:",chardata"`
											Rt       string `xml:"Rt"`
										} `xml:"Ruby"`
									} `xml:"Sentence"`
								} `xml:"Subitem1Sentence"`
								Subitem1Title string `xml:"Subitem1Title"`
							} `xml:"Subitem1"`
						} `xml:"Item"`
						List []struct {
							ListSentence struct {
								Sentence struct {
									Num         string `xml:"Num,attr"`
									WritingMode string `xml:"WritingMode,attr"`
									CharData    string `xml:",chardata"`
								} `xml:"Sentence"`
							} `xml:"ListSentence"`
						} `xml:"List"`
						ParagraphNum      string `xml:"ParagraphNum"`
						ParagraphSentence struct {
							Sentence []struct {
								Function    *string `xml:"Function,attr"`
								Num         string  `xml:"Num,attr"`
								WritingMode string  `xml:"WritingMode,attr"`
								CharData    string  `xml:",chardata"`
								Ruby        []struct {
									CharData string `xml:",chardata"`
									Rt       string `xml:"Rt"`
								} `xml:"Ruby"`
							} `xml:"Sentence"`
						} `xml:"ParagraphSentence"`
						TableStruct *struct {
							Table struct {
								WritingMode string `xml:"WritingMode,attr"`
								TableRow    []struct {
									TableColumn []struct {
										BorderBottom string `xml:"BorderBottom,attr"`
										BorderLeft   string `xml:"BorderLeft,attr"`
										BorderRight  string `xml:"BorderRight,attr"`
										BorderTop    string `xml:"BorderTop,attr"`
										Sentence     []struct {
											Num         string `xml:"Num,attr"`
											WritingMode string `xml:"WritingMode,attr"`
											CharData    string `xml:",chardata"`
										} `xml:"Sentence"`
									} `xml:"TableColumn"`
								} `xml:"TableRow"`
							} `xml:"Table"`
						} `xml:"TableStruct"`
					} `xml:"Paragraph"`
				} `xml:"Article"`
				ChapterTitle string `xml:"ChapterTitle"`
				Section      []struct {
					Num     string `xml:"Num,attr"`
					Article []struct {
						Num            float64 `xml:"Num,attr"`
						ArticleCaption *string `xml:"ArticleCaption"`
						ArticleTitle   string  `xml:"ArticleTitle"`
						Paragraph      []struct {
							Num  string `xml:"Num,attr"`
							Item []struct {
								Num          string `xml:"Num,attr"`
								ItemSentence struct {
									Sentence struct {
										Num         string `xml:"Num,attr"`
										WritingMode string `xml:"WritingMode,attr"`
										CharData    string `xml:",chardata"`
									} `xml:"Sentence"`
								} `xml:"ItemSentence"`
								ItemTitle string `xml:"ItemTitle"`
							} `xml:"Item"`
							ParagraphNum      string `xml:"ParagraphNum"`
							ParagraphSentence struct {
								Sentence []struct {
									Function    *string `xml:"Function,attr"`
									Num         string  `xml:"Num,attr"`
									WritingMode string  `xml:"WritingMode,attr"`
									CharData    string  `xml:",chardata"`
								} `xml:"Sentence"`
							} `xml:"ParagraphSentence"`
						} `xml:"Paragraph"`
					} `xml:"Article"`
					SectionTitle string `xml:"SectionTitle"`
					Subsection   []struct {
						Num     string `xml:"Num,attr"`
						Article []struct {
							Num            float64 `xml:"Num,attr"`
							ArticleCaption string  `xml:"ArticleCaption"`
							ArticleTitle   string  `xml:"ArticleTitle"`
							Paragraph      []struct {
								Num  string `xml:"Num,attr"`
								Item []struct {
									Num          string `xml:"Num,attr"`
									ItemSentence struct {
										Column []struct {
											Num      string `xml:"Num,attr"`
											Sentence struct {
												Num         string `xml:"Num,attr"`
												WritingMode string `xml:"WritingMode,attr"`
												CharData    string `xml:",chardata"`
											} `xml:"Sentence"`
										} `xml:"Column"`
										Sentence *struct {
											Num         string `xml:"Num,attr"`
											WritingMode string `xml:"WritingMode,attr"`
											CharData    string `xml:",chardata"`
										} `xml:"Sentence"`
									} `xml:"ItemSentence"`
									ItemTitle string `xml:"ItemTitle"`
									Subitem1  []struct {
										Num              string `xml:"Num,attr"`
										Subitem1Sentence struct {
											Sentence struct {
												Num         string `xml:"Num,attr"`
												WritingMode string `xml:"WritingMode,attr"`
												CharData    string `xml:",chardata"`
											} `xml:"Sentence"`
										} `xml:"Subitem1Sentence"`
										Subitem1Title string `xml:"Subitem1Title"`
									} `xml:"Subitem1"`
								} `xml:"Item"`
								ParagraphNum      string `xml:"ParagraphNum"`
								ParagraphSentence struct {
									Sentence []struct {
										Function    *string `xml:"Function,attr"`
										Num         string  `xml:"Num,attr"`
										WritingMode string  `xml:"WritingMode,attr"`
										CharData    string  `xml:",chardata"`
									} `xml:"Sentence"`
								} `xml:"ParagraphSentence"`
								TableStruct *struct {
									Table struct {
										WritingMode string `xml:"WritingMode,attr"`
										TableRow    []struct {
											TableColumn []struct {
												BorderBottom string `xml:"BorderBottom,attr"`
												BorderLeft   string `xml:"BorderLeft,attr"`
												BorderRight  string `xml:"BorderRight,attr"`
												BorderTop    string `xml:"BorderTop,attr"`
												Sentence     []struct {
													Num         string `xml:"Num,attr"`
													WritingMode string `xml:"WritingMode,attr"`
													CharData    string `xml:",chardata"`
												} `xml:"Sentence"`
											} `xml:"TableColumn"`
										} `xml:"TableRow"`
									} `xml:"Table"`
								} `xml:"TableStruct"`
							} `xml:"Paragraph"`
						} `xml:"Article"`
						SubsectionTitle string `xml:"SubsectionTitle"`
					} `xml:"Subsection"`
				} `xml:"Section"`
			} `xml:"Chapter"`
		} `xml:"MainProvision"`
		SupplProvision []struct {
			AmendLawNum string `xml:"AmendLawNum,attr"`
			Extract     *bool  `xml:"Extract,attr"`
			Article     []struct {
				Num            string  `xml:"Num,attr"`
				ArticleCaption *string `xml:"ArticleCaption"`
				ArticleTitle   string  `xml:"ArticleTitle"`
				Paragraph      []struct {
					Num  string `xml:"Num,attr"`
					Item []struct {
						Num          string `xml:"Num,attr"`
						ItemSentence struct {
							Column []struct {
								Num      string `xml:"Num,attr"`
								Sentence struct {
									Num         string `xml:"Num,attr"`
									WritingMode string `xml:"WritingMode,attr"`
									CharData    string `xml:",chardata"`
								} `xml:"Sentence"`
							} `xml:"Column"`
							Sentence *struct {
								Num         string `xml:"Num,attr"`
								WritingMode string `xml:"WritingMode,attr"`
								CharData    string `xml:",chardata"`
							} `xml:"Sentence"`
						} `xml:"ItemSentence"`
						ItemTitle string `xml:"ItemTitle"`
					} `xml:"Item"`
					ParagraphNum      string `xml:"ParagraphNum"`
					ParagraphSentence struct {
						Sentence []struct {
							Function    *string `xml:"Function,attr"`
							Num         string  `xml:"Num,attr"`
							WritingMode string  `xml:"WritingMode,attr"`
							CharData    string  `xml:",chardata"`
						} `xml:"Sentence"`
					} `xml:"ParagraphSentence"`
					TableStruct *struct {
						Table struct {
							WritingMode string `xml:"WritingMode,attr"`
							TableRow    []struct {
								TableColumn []struct {
									BorderBottom string `xml:"BorderBottom,attr"`
									BorderLeft   string `xml:"BorderLeft,attr"`
									BorderRight  string `xml:"BorderRight,attr"`
									BorderTop    string `xml:"BorderTop,attr"`
									Sentence     []struct {
										Num         string `xml:"Num,attr"`
										WritingMode string `xml:"WritingMode,attr"`
										CharData    string `xml:",chardata"`
									} `xml:"Sentence"`
								} `xml:"TableColumn"`
							} `xml:"TableRow"`
						} `xml:"Table"`
					} `xml:"TableStruct"`
				} `xml:"Paragraph"`
			} `xml:"Article"`
			Paragraph []struct {
				Num  string `xml:"Num,attr"`
				Item []struct {
					Num          string `xml:"Num,attr"`
					ItemSentence struct {
						Column []struct {
							Num      string `xml:"Num,attr"`
							Sentence struct {
								Num         string `xml:"Num,attr"`
								WritingMode string `xml:"WritingMode,attr"`
								CharData    string `xml:",chardata"`
							} `xml:"Sentence"`
						} `xml:"Column"`
						Sentence *struct {
							Num         string `xml:"Num,attr"`
							WritingMode string `xml:"WritingMode,attr"`
							CharData    string `xml:",chardata"`
						} `xml:"Sentence"`
					} `xml:"ItemSentence"`
					ItemTitle string `xml:"ItemTitle"`
				} `xml:"Item"`
				ParagraphCaption  *string `xml:"ParagraphCaption"`
				ParagraphNum      string  `xml:"ParagraphNum"`
				ParagraphSentence struct {
					Sentence []struct {
						Function    *string `xml:"Function,attr"`
						Num         string  `xml:"Num,attr"`
						WritingMode string  `xml:"WritingMode,attr"`
						CharData    string  `xml:",chardata"`
						Ruby        []struct {
							CharData string `xml:",chardata"`
							Rt       string `xml:"Rt"`
						} `xml:"Ruby"`
					} `xml:"Sentence"`
				} `xml:"ParagraphSentence"`
			} `xml:"Paragraph"`
			SupplProvisionLabel string `xml:"SupplProvisionLabel"`
		} `xml:"SupplProvision"`
		TOC struct {
			TOCChapter []struct {
				Num          string  `xml:"Num,attr"`
				ArticleRange *string `xml:"ArticleRange"`
				ChapterTitle string  `xml:"ChapterTitle"`
				TOCSection   []struct {
					Num           string  `xml:"Num,attr"`
					ArticleRange  *string `xml:"ArticleRange"`
					SectionTitle  string  `xml:"SectionTitle"`
					TOCSubsection []struct {
						Num             string `xml:"Num,attr"`
						ArticleRange    string `xml:"ArticleRange"`
						SubsectionTitle string `xml:"SubsectionTitle"`
					} `xml:"TOCSubsection"`
				} `xml:"TOCSection"`
			} `xml:"TOCChapter"`
			TOCLabel          string `xml:"TOCLabel"`
			TOCSupplProvision struct {
				SupplProvisionLabel string `xml:"SupplProvisionLabel"`
			} `xml:"TOCSupplProvision"`
		} `xml:"TOC"`
	} `xml:"LawBody"`
	LawNum string `xml:"LawNum"`
}
