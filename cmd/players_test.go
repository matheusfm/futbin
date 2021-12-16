package cmd

import (
	"reflect"
	"testing"

	"github.com/matheusfm/futbin/players"
)

func TestFlagToRange(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want players.Range
	}{
		{name: "exactly", arg: "90", want: players.Range{Min: 90, Max: 90}},
		{name: "range", arg: "80-90", want: players.Range{Min: 80, Max: 90}},
		{name: "invalid", arg: "foo", want: players.Range{}},
		{name: "min invalid", arg: "foo-90", want: players.Range{}},
		{name: "max invalid", arg: "90-bar", want: players.Range{}},
		{name: "min", arg: "90-", want: players.Range{Min: 90}},
		{name: "max", arg: "-99", want: players.Range{Max: 99}},
		{name: "long numbers", arg: "1000000-10000000", want: players.Range{Min: 1000000, Max: 10000000}},
		{name: "comma", arg: "10,000-100,000", want: players.Range{Min: 10000, Max: 100000}},
		{name: "comma min", arg: "10,000-", want: players.Range{Min: 10000}},
		{name: "comma max", arg: "-100,000", want: players.Range{Max: 100000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flagToRange(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flagToRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
