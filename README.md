# futbin

An unofficial command line utility for [futbin](https://www.futbin.com/). Also, it can be used as a library.

# Install

## Pre-compiled executables

Get them [here](https://github.com/matheusfm/futbin/releases).

## Source

You need `go` installed and `GOBIN` in your `PATH`. Once that is done, run the command:

```shell
go install github.com/matheusfm/futbin@latest
```

# Usage

```
Usage:
  futbin [command]

Available Commands:
  cardversions card versions
  clubs        clubs
  completion   generate the autocompletion script for the specified shell
  help         Help about any command
  latest       latest added players
  leagues      leagues
  nations      nations
  players      players
  popular      popular players
  totw         current TOTW players

Flags:
      --config string     Config file (default is $HOME/.futbin.yaml)
  -h, --help              help for futbin
      --platform string   Platform (PS, XB or PC) (default "PS")

Use "futbin [command] --help" for more information about a command.
```

### `players` command examples:

```
Usage:
  futbin players [flags]

Examples:
1.  # Brazilian players in LaLiga:
    futbin players --nation 54 --league 53
2.  # OTW players (see options in 'futbin cardversions' command):
    futbin players --version otw
3.  # Brazilian players with more than 90 of passing:
    futbin players --nation 54 --passing 90-
4.  # Icons with a maximum price of 300K:
    futbin players --league 2118 --price -300000
5.  # Players with 5 weak foot and 5 skills:
    futbin players --wf 5 --skills 5

Flags:
      --acceleration string         Acceleration
      --aggression string           Aggression
      --agility string              Agility
      --balance string              Balance
      --ball-control string         Ball Control
      --club int                    Club ID
      --composure string            Composure
      --crossing string             Crossing
      --curve string                Curve
      --defending string            Defending
      --dribbling string            Dribbling
      --finishing string            Finishing
      --free-kick-accuracy string   Free Kick Accuracy
      --heading-accuracy string     Heading Accuracy
  -h, --help                        help for players
      --interceptions string        Interceptions
      --jumping string              Jumping
      --league int                  League ID
      --long-passing string         Long Passing
      --long-shots string           Long Shots
      --marking string              Marking
      --nation int                  Nation ID
      --ovr string                  Rating
      --pace string                 Pace
      --page int                    Page (default 1)
      --passing string              Passing
      --penalties string            Penalties
      --physicality string          Physicality
      --positioning string          Positioning
      --price string                Price
      --reactions string            Reactions
      --shooting string             Shooting
      --short-passing string        Short Passing
      --shot-power string           Shot Power
      --skills string               Skills
      --sliding-tackle string       Sliding Tackle
      --sprint-speed string         Sprint Speed
      --stamina string              Stamina
      --standing-tackle string      Standing Tackle
      --strength string             Strength
      --version string              Card version
      --vision string               Vision
      --volleys string              Volleys
      --wf string                   Weak Foot

Global Flags:
      --config string     Config file (default is $HOME/.futbin.yaml)
      --platform string   Platform (PS, XB or PC) (default "PS")
```

## Usage (Library)

```go
package main

import (
	"fmt"
	"os"

	"github.com/matheusfm/futbin/players"
)

func main() {
	playersClient := players.DefaultClient()
	p, err := playersClient.Get(&players.Options{
		Platform: "PS",
		NationID: 54, // Brazil
		LeagueID: 53, // LaLiga
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	for _, player := range p {
		fmt.Println(player.CommonName)
	}
}
```

# License

See [LICENSE](https://github.com/matheusfm/futbin/blob/main/LICENSE).
