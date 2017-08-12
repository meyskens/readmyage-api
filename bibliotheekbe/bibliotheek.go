package bibliotheekbe

import (
	"encoding/xml"

	resty "gopkg.in/resty.v0"
)

const authkey = "ac135e89f84460a251a6283a14180a22" // Found on https://docs.google.com/spreadsheets/d/1PnyOzCkxSnEWKqmTk6bsnuorAIk27-8yvD2gKf6zq_4/edit#gid=9

var r = resty.New().SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).SetHostURL("http://zoeken.bibliotheek.be/api/v0")

// Search does a search and sends back the results
func Search(term string) (SearchResponse, error) {
	res := SearchResponse{}

	resp, err := r.R().SetQueryParams(map[string]string{
		"q":             term,
		"authorization": authkey,
	}).Get("/search")
	if err != nil {
		return res, err
	}

	err = xml.Unmarshal(resp.Body(), &res)

	return res, err
}
