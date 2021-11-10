package players

type Player struct {
	ID         int `json:"ID"`
	PlayerID   int `json:"playerid"`
	ResourceID int `json:"resource_id"`
	Image      int `json:"playerimage"`

	Name       string `json:"playername"`
	CommonName string `json:"common_name"`

	Position string `json:"position"`
	Rating   int    `json:"rating"`
	RareType int    `json:"raretype"`

	Pace        int `json:"pac"`
	Shooting    int `json:"sho"`
	Passing     int `json:"pas"`
	Dribbling   int `json:"dri"`
	Defending   int `json:"def"`
	Physicality int `json:"phy"`

	ClubID   int    `json:"club"`
	ClubName string `json:"club_name"`
	NationID int    `json:"nation"`
	LeagueID int    `json:"league"`
	League   string `json:"league_name"`

	PsPrice      int `json:"ps_LCPrice"`
	PsPRP        int `json:"ps_PRP"`
	PsMinPrice   int `json:"ps_MinPrice"`
	PsMaxPrice   int `json:"ps_MaxPrice"`
	PcPrice      int `json:"pc_LCPrice"`
	PcPRP        int `json:"pc_PRP"`
	PcMinPrice   int `json:"pc_MinPrice"`
	PcMaxPrice   int `json:"pc_MaxPrice"`
	XboxPrice    int `json:"xbox_LCPrice"`
	XboxPRP      int `json:"xbox_PRP"`
	XboxMinPrice int `json:"xbox_MinPrice"`
	XboxMaxPrice int `json:"xbox_MaxPrice"`
}

func (p Player) Price(platform string) int {
	switch platform {
	case "PC":
		return p.PcPrice
	case "XB":
		return p.XboxPrice
	default:
		return p.PsPrice
	}
}

func (p Player) MinPrice(platform string) int {
	switch platform {
	case "PC":
		return p.PcMinPrice
	case "XB":
		return p.XboxMinPrice
	default:
		return p.PsMinPrice
	}
}

func (p Player) MaxPrice(platform string) int {
	switch platform {
	case "PC":
		return p.PcMaxPrice
	case "XB":
		return p.XboxMaxPrice
	default:
		return p.PsMaxPrice
	}
}

func (p Player) PRP(platform string) int {
	switch platform {
	case "PC":
		return p.PcPRP
	case "XB":
		return p.XboxPRP
	default:
		return p.PsPRP
	}
}

type playersData struct {
	Players []Player `json:"data"`
}
