package npbbis

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Game struct {
	Date     string
	Id       string
	Homeruns []*Homerun
}

func NewGame(date string, id string, homeruns []*Homerun) *Game {
	return &Game{
		Date:     date,
		Id:       id,
		Homeruns: homeruns,
	}
}

func GetGame(date string, id string) (*Game, error) {
	year := date[0:4]
	uri := "http://bis.npb.or.jp/" + year + "/games/s" + id + ".html"
	doc, err := newGoqueryDocument(uri)
	if err != nil {
		return nil, err
	}

	hrs := []*Homerun{}
	doc.Find("#gmdivhr .gmresults").Each(func(i int, s *goquery.Selection) {
		hrs = append(hrs, parseHomerun(s.Text())...)
	})
	return NewGame(date, id, hrs), nil
}

func GetGames(date string) ([]*Game, error) {
	uri := "http://bis.npb.or.jp/" + date[0:4] + "/games/gm" + date + ".html"
	doc, err := newGoqueryDocument(uri)
	if err != nil {
		return nil, err
	}
	games := []*Game{}
	doc.Find(".contentsgame .contentsinfo a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			id := href[1:14]
			var game *Game
			game, err = GetGame(date, id)
			games = append(games, game)
		}
	})
	return games, err
}

// parsers

func parseHomerun(line string) []*Homerun {
	hrs := []*Homerun{}
	for _, chunk := range strings.SplitAfter(line, "］ ")[1:] {
		lastBatter := ""
		for _, subchunk := range strings.Split(chunk, ",") {
			tokens := strings.Split(subchunk, " ")
			batter := tokens[0]
			number := tokens[1]
			scenario := tokens[3]
			pitcher := tokens[4]
			if len(batter) == 0 {
				batter = lastBatter
			}
			hrs = append(hrs, NewHomerun(batter, number, scenario, pitcher))
			lastBatter = batter
		}
	}
	return hrs
}
