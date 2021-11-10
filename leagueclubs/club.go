package leagueclubs

type Club struct {
	ID           int    `json:"club_id"`
	Name         string `json:"club_name"`
	LeagueID     int    `json:"league_id"`
	PlayerCount  int    `json:"player_count"`
	SpecialCount int    `json:"special_count"`
	GoldCount    int    `json:"gold_count"`
	SilverCount  int    `json:"silver_count"`
	BronzeCount  int    `json:"bronze_count"`
}

type League struct {
	ID    int    `json:"league_id"`
	Name  string `json:"league_name"`
	Clubs []Club `json:"clubs"`
}

type leaguesAndClubs struct {
	Leagues []League `json:"data"`
}
