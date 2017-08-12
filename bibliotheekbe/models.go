package bibliotheekbe

type SearchResponse struct {
	Result []SearchResult `xml:"results>result"`
}

type SearchResult struct {
	ID              string   `xml:"id"`
	Copyright       string   `xml:"copyright"`
	Title           string   `xml:"titles>title"`
	OriginalTitle   string   `xml:"titles>origin-title"`
	Author          string   `xml:"authors>main-author"`
	MainSubject     string   `xml:"subjects>topical-subject"`
	Genres          []string `xml:"genres>genre"`
	TargetAudiences []string `xml:"target-audiences>target-audience"`
	Series          []string `xml:"series>series-title"`
	Covers          []string `xml:"coverimages>coverimage>url"`
}
