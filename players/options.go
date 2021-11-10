package players

type Options struct {
	Platform string `url:"platform,omitempty"`

	Page  int    `url:"page,omitempty"`
	Sort  string `url:"sort,omitempty"`
	Order string `url:"order,omitempty"`

	Version string `url:"version,omitempty"`

	Position    []Position `url:"position,omitempty"`
	NationID    int        `url:"nation,omitempty"`
	LeagueID    int        `url:"league,omitempty"`
	ClubID      int        `url:"club,omitempty"`
	WeakFoot    int        `url:"wf,omitempty"`
	MinWeakFoot int        `url:"min_wf,omitempty"`
	MaxWeakFoot int        `url:"max_wf,omitempty"`
	Skills      int        `url:"skills,omitempty"`
	MinSkills   int        `url:"min_skills,omitempty"`
	MaxSkills   int        `url:"max_skills,omitempty"`
	Foot        Foot       `url:"foot,omitempty"`
	PsPrice     int        `url:"ps4price,omitempty"`
	XboxPrice   int        `url:"xb1price,omitempty"`
	PcPrice     int        `url:"pcprice,omitempty"`
	Rating      string     `url:"rating,omitempty"`
	Height      string     `url:"height,omitempty"`
	Weight      string     `url:"weight,omitempty"`

	Pace         string `url:"Pace,omitempty"`
	Acceleration string `url:"Acceleration,omitempty"`
	SprintSpeed  string `url:"Sprintspeed,omitempty"`

	Shooting    string `url:"Shooting,omitempty"`
	Positioning string `url:"Positioning,omitempty"`
	Finishing   string `url:"Finishing,omitempty"`
	ShotPower   string `url:"Shotpower,omitempty"`
	LongShots   string `url:"Longshots,omitempty"`
	Volleys     string `url:"Volleys,omitempty"`
	Penalties   string `url:"Penalties,omitempty"`

	Dribbling   string `url:"Dribbling,omitempty"`
	Agility     string `url:"Agility,omitempty"`
	Balance     string `url:"Balance,omitempty"`
	Reactions   string `url:"Reactions,omitempty"`
	BallControl string `url:"Ballcontrol,omitempty"`
	Composure   string `url:"Composure,omitempty"`

	Defending       string `url:"Defending,omitempty"`
	Interceptions   string `url:"Interceptions,omitempty"`
	HeadingAccuracy string `url:"Headingaccuracy,omitempty"`
	Marking         string `url:"Marking,omitempty"`
	StandingTackle  string `url:"Standingtackle,omitempty"`
	SlidingTackle   string `url:"Slidingtackle,omitempty"`

	Passing          string `url:"Passing,omitempty"`
	Vision           string `url:"Vision,omitempty"`
	Crossing         string `url:"Crossing,omitempty"`
	FreeKickAccuracy string `url:"Freekickaccuracy,omitempty"`
	ShortPassing     string `url:"Shortpassing,omitempty"`
	LongPassing      string `url:"Longpassing,omitempty"`
	Curve            string `url:"Curve,omitempty"`

	Physicality string `url:"Physicality,omitempty"`
	Jumping     string `url:"Jumping,omitempty"`
	Stamina     string `url:"Stamina,omitempty"`
	Strength    string `url:"Strength,omitempty"`
	Aggression  string `url:"Aggression,omitempty"`
}
