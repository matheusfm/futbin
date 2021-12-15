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
futbin command line tool

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
      --config string     config file (default is $HOME/.futbin.yaml)
  -h, --help              help for futbin
      --platform string   platform (PS, XB or PC) (default "PS")

Use "futbin [command] --help" for more information about a command.
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
