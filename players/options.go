package players

type Options struct {
	Platform string

	Page  int
	Sort  string
	Order string

	Version string

	Position   []string
	Accelerate string
	NationID   int
	LeagueID   int
	ClubID     int
	Height     string
	Weight     string
	WeakFoot   Range
	Skills     Range
	Foot       Foot
	Rating     Range
	Price      Range

	Pace             Range
	Acceleration     Range
	SprintSpeed      Range
	Shooting         Range
	Positioning      Range
	Finishing        Range
	ShotPower        Range
	LongShots        Range
	Volleys          Range
	Penalties        Range
	Dribbling        Range
	Agility          Range
	Balance          Range
	Reactions        Range
	BallControl      Range
	Composure        Range
	Defending        Range
	Interceptions    Range
	HeadingAccuracy  Range
	Marking          Range
	StandingTackle   Range
	SlidingTackle    Range
	Passing          Range
	Vision           Range
	Crossing         Range
	FreeKickAccuracy Range
	ShortPassing     Range
	LongPassing      Range
	Curve            Range
	Physicality      Range
	Jumping          Range
	Stamina          Range
	Strength         Range
	Aggression       Range
}

type Range struct {
	Min int
	Max int
}
