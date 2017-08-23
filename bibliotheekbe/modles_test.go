package bibliotheekbe

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInfoReuzenperzikRaw = SearchResponse{
	Result: []SearchResult{
		SearchResult{
			ID:              "|library/marc/vlacc|8385303",
			Copyright:       "\n            Op de metadata (titelbeschrijvingen) en metacontent (covers, samenvattingen, recensies enz.) aangeboden in de API berusten intellectuele eigendomsrechten en databankrechten, zie http://www.cultuurconnect.be/diensten/bibliotheekportalen/open-data. De API mag enkel worden gebruikt in functie van de opdracht van de openbare bibliotheek.\n         ",
			Title:           "De reuzenperzik",
			OriginalTitle:   "James and the giant peach",
			Author:          "Dahl, Roald",
			MainSubject:     "Toveren",
			Genres:          []string{"Humor"},
			TargetAudiences: []string{"age9-11", "ageYouth"},
			Series:          []string{"De fantastische bibliotheek van Roald Dahl"},
			Covers:          []string{"https://webservices.bibliotheek.be/index.php?func=cover&ISBN=9789026119453&VLACCnr=8385303&CDR=&EAN=&ISMN=&coversize=small"},
		},
	},
}

var testInfoReuzenperzik = SearchResponse{
	Result: []SearchResult{
		SearchResult{
			ID:              "|library/marc/vlacc|8385303",
			Copyright:       "\n            Op de metadata (titelbeschrijvingen) en metacontent (covers, samenvattingen, recensies enz.) aangeboden in de API berusten intellectuele eigendomsrechten en databankrechten, zie http://www.cultuurconnect.be/diensten/bibliotheekportalen/open-data. De API mag enkel worden gebruikt in functie van de opdracht van de openbare bibliotheek.\n         ",
			Title:           "De reuzenperzik",
			OriginalTitle:   "James and the giant peach",
			Author:          "Dahl, Roald",
			MainSubject:     "Toveren",
			Genres:          []string{"Humor"},
			TargetAudiences: []string{"ages 9-11", "youth"},
			Series:          []string{"De fantastische bibliotheek van Roald Dahl"},
			Covers:          []string{"https://webservices.bibliotheek.be/index.php?func=cover&ISBN=9789026119453&VLACCnr=8385303&CDR=&EAN=&ISMN=&coversize=small"},
		},
	},
}

func TestSearchResults(t *testing.T) {
	file, _ := os.Open("test/test-search.xml")
	data, _ := ioutil.ReadAll(file)

	res := SearchResponse{}
	xml.Unmarshal(data, &res)

	assert.Equal(t, testInfoReuzenperzikRaw, res)
}
