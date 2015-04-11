package npbbis

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func newGoqueryDocument(uri string) (*goquery.Document, error) {
	reader, err := createShiftJISReader(uri)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(reader)
}

func createShiftJISReader(uri string) (io.Reader, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return transform.NewReader(res.Body, japanese.ShiftJIS.NewDecoder()), nil
}
