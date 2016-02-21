package npbbis

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func newGoqueryDocument(uri string) (*goquery.Document, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(res.Body)
}
