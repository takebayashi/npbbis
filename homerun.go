package npbbis

import ()

type Homerun struct {
	Batter   string
	Number   string
	Scenario string
	Pitcher  string
}

func NewHomerun(batter string, number string, scenario string, pitcher string) *Homerun {
	return &Homerun{
		Batter:   batter,
		Number:   number,
		Scenario: scenario,
		Pitcher:  pitcher,
	}
}
