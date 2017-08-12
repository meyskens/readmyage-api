package main

import "github.com/meyskens/readmyage-api/bibliotheekbe"
import "strings"

func lookUpBibliotheek(isbn string, result chan bool, out chan LookUpResult) {
	res, err := bibliotheekbe.Search(isbn)
	if err != nil || len(res.Result) <= 0 {
		result <- false
		return
	}

	book := res.Result[0] //ISBN has probably one match
	lookup := LookUpResult{
		Source:    "bibliotheek.be",
		Title:     book.Title,
		Author:    book.Author,
		Copyright: book.Copyright,
		Audience:  strings.Join(book.TargetAudiences, ", "),
	}
	if len(book.Covers) > 0 {
		lookup.Cover = book.Covers[0]
	}
	result <- true
	out <- lookup
}
