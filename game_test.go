package npbbis

import "testing"

func TestKishiNoHitterGame(t *testing.T) {
	game, err := GetGame("20140502", "2014050200471")
	if err != nil || game == nil {
		t.Error("failed to fetch game info")
	}
	if len(game.Homeruns) != 0 {
		t.Error("failed to parse homeruns")
	}
}
