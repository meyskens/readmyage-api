package main

// LookUpResponse is the response of a lookup
type LookUpResponse struct {
	Results []LookUpResult `json:"results"`
}

// LookUpResult is the result from a single source
type LookUpResult struct {
	Source    string `json:"source"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Cover     string `json:"cover"`
	Audience  string `json:"audience"`
	Copyright string `json:"copyright"`
}
