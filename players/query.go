package players

import "fmt"

type playerQuery struct {
	Platform string `url:"platform,omitempty"`

	Page  int    `url:"page,omitempty"`
	Sort  string `url:"sort,omitempty"`
	Order string `url:"order,omitempty"`

	Version string `url:"version,omitempty"`

	Position    	[]string `url:"position,comma,omitempty"`
	PositionType 	string 	 `url:"pos_type,comma,omitempty"`
	Accelerate  	string   `url:"accelerate,omitempty"`
	NationID    	int      `url:"nation,omitempty"`
	LeagueID    	int      `url:"league,omitempty"`
	ClubID      	int      `url:"club,omitempty"`
	WeakFoot    	int      `url:"wf,omitempty"`
	MinWeakFoot 	int      `url:"min_wf,omitempty"`
	MaxWeakFoot 	int      `url:"max_wf,omitempty"`
	Skills      	int      `url:"skills,omitempty"`
	MinSkills   	int      `url:"min_skills,omitempty"`
	MaxSkills   	int      `url:"max_skills,omitempty"`
	Foot        	Foot     `url:"foot,omitempty"`
	PsPrice     	string   `url:"ps4price,omitempty"`
	XboxPrice   	string   `url:"xb1price,omitempty"`
	PcPrice     	string   `url:"pcprice,omitempty"`
	Rating      	string   `url:"rating,omitempty"`
	Height      	string   `url:"height,omitempty"`
	Weight      	string   `url:"weight,omitempty"`

	Pace             string `url:"Pace,omitempty"`
	Acceleration     string `url:"Acceleration,omitempty"`
	SprintSpeed      string `url:"Sprintspeed,omitempty"`
	Shooting         string `url:"Shooting,omitempty"`
	Positioning      string `url:"Positioning,omitempty"`
	Finishing        string `url:"Finishing,omitempty"`
	ShotPower        string `url:"Shotpower,omitempty"`
	LongShots        string `url:"Longshots,omitempty"`
	Volleys          string `url:"Volleys,omitempty"`
	Penalties        string `url:"Penalties,omitempty"`
	Dribbling        string `url:"Dribbling,omitempty"`
	Agility          string `url:"Agility,omitempty"`
	Balance          string `url:"Balance,omitempty"`
	Reactions        string `url:"Reactions,omitempty"`
	BallControl      string `url:"Ballcontrol,omitempty"`
	Composure        string `url:"Composure,omitempty"`
	Defending        string `url:"Defending,omitempty"`
	Interceptions    string `url:"Interceptions,omitempty"`
	HeadingAccuracy  string `url:"Headingaccuracy,omitempty"`
	Marking          string `url:"Marking,omitempty"`
	StandingTackle   string `url:"Standingtackle,omitempty"`
	SlidingTackle    string `url:"Slidingtackle,omitempty"`
	Passing          string `url:"Passing,omitempty"`
	Vision           string `url:"Vision,omitempty"`
	Crossing         string `url:"Crossing,omitempty"`
	FreeKickAccuracy string `url:"Freekickaccuracy,omitempty"`
	ShortPassing     string `url:"Shortpassing,omitempty"`
	LongPassing      string `url:"Longpassing,omitempty"`
	Curve            string `url:"Curve,omitempty"`
	Physicality      string `url:"Physicality,omitempty"`
	Jumping          string `url:"Jumping,omitempty"`
	Stamina          string `url:"Stamina,omitempty"`
	Strength         string `url:"Strength,omitempty"`
	Aggression       string `url:"Aggression,omitempty"`
}

func newPlayerQuery(opt *Options) *playerQuery {
	price := rangeString(opt.Price, "-", 200, 15000000)
	pq := &playerQuery{
		Platform:    opt.Platform,
		Page:        opt.Page,
		Sort:        opt.Sort,
		Order:       opt.Order,
		Version:     opt.Version,
		Position:    opt.Position,
		PositionType: opt.PositionType,
		Accelerate:  opt.Accelerate,
		NationID:    opt.NationID,
		LeagueID:    opt.LeagueID,
		ClubID:      opt.ClubID,
		MinWeakFoot: opt.WeakFoot.Min,
		MaxWeakFoot: opt.WeakFoot.Max,
		MinSkills:   opt.Skills.Min,
		MaxSkills:   opt.Skills.Max,
		Foot:        opt.Foot,
		Rating:      rangeString(opt.Rating, "-", 30, 99),
		Height:      opt.Height,
		Weight:      opt.Weight,

		Pace:             rangeString(opt.Pace, ",", 0, 99),
		Acceleration:     rangeString(opt.Acceleration, ",", 0, 99),
		SprintSpeed:      rangeString(opt.SprintSpeed, ",", 0, 99),
		Shooting:         rangeString(opt.Shooting, ",", 0, 99),
		Positioning:      rangeString(opt.Positioning, ",", 0, 99),
		Finishing:        rangeString(opt.Finishing, ",", 0, 99),
		ShotPower:        rangeString(opt.ShotPower, ",", 0, 99),
		LongShots:        rangeString(opt.LongShots, ",", 0, 99),
		Volleys:          rangeString(opt.Volleys, ",", 0, 99),
		Penalties:        rangeString(opt.Penalties, ",", 0, 99),
		Dribbling:        rangeString(opt.Dribbling, ",", 0, 99),
		Agility:          rangeString(opt.Agility, ",", 0, 99),
		Balance:          rangeString(opt.Balance, ",", 0, 99),
		Reactions:        rangeString(opt.Reactions, ",", 0, 99),
		BallControl:      rangeString(opt.BallControl, ",", 0, 99),
		Composure:        rangeString(opt.Composure, ",", 0, 99),
		Defending:        rangeString(opt.Defending, ",", 0, 99),
		Interceptions:    rangeString(opt.Interceptions, ",", 0, 99),
		HeadingAccuracy:  rangeString(opt.HeadingAccuracy, ",", 0, 99),
		Marking:          rangeString(opt.Marking, ",", 0, 99),
		StandingTackle:   rangeString(opt.StandingTackle, ",", 0, 99),
		SlidingTackle:    rangeString(opt.SlidingTackle, ",", 0, 99),
		Passing:          rangeString(opt.Passing, ",", 0, 99),
		Vision:           rangeString(opt.Vision, ",", 0, 99),
		Crossing:         rangeString(opt.Crossing, ",", 0, 99),
		FreeKickAccuracy: rangeString(opt.FreeKickAccuracy, ",", 0, 99),
		ShortPassing:     rangeString(opt.ShortPassing, ",", 0, 99),
		LongPassing:      rangeString(opt.LongPassing, ",", 0, 99),
		Curve:            rangeString(opt.Curve, ",", 0, 99),
		Physicality:      rangeString(opt.Physicality, ",", 0, 99),
		Jumping:          rangeString(opt.Jumping, ",", 0, 99),
		Stamina:          rangeString(opt.Stamina, ",", 0, 99),
		Strength:         rangeString(opt.Strength, ",", 0, 99),
		Aggression:       rangeString(opt.Aggression, ",", 0, 99),
	}
	switch opt.Platform {
	case "XB":
		pq.XboxPrice = price
	case "PC":
		pq.PcPrice = price
	default:
		pq.PsPrice = price
	}
	return pq
}

func rangeString(r Range, sep string, minLimit int, maxLimit int) string {
	if r.Min == 0 && r.Max == 0 {
		return ""
	}
	if r.Min < minLimit {
		r.Min = minLimit
	}
	if r.Max > maxLimit {
		r.Max = maxLimit
	}
	return fmt.Sprintf("%d%s%d", r.Min, sep, r.Max)
}
