package flagutils_test

import (
	"flag"
	"fmt"
	"image/color"
	"reflect"
	"strings"
	"testing"

	"github.com/fishy/go-flagutils"
)

func TestRGBA(t *testing.T) {
	t.Run(
		"Error",
		func(t *testing.T) {
			errors := []string{
				"#ffffffff",  // with leading #
				"ffffffffff", // out of range
				"Loripsum",   // not hex string
			}

			for _, s := range errors {
				t.Run(
					s,
					func(t *testing.T) {
						var c flagutils.RGBA
						err := c.Set(s)
						if err == nil {
							t.Errorf(
								"Set(%q) expected error, got: %v",
								s,
								err,
							)
						}
					},
				)
			}
		},
	)

	t.Run(
		"NoError",
		func(t *testing.T) {
			cases := []struct {
				S     string
				Color color.RGBA
			}{
				{
					"ff000080",
					color.RGBA{
						R: 0xff,
						A: 0x80,
					},
				},
				{
					"aBcDeF01",
					color.RGBA{
						R: 0xab,
						G: 0xcd,
						B: 0xef,
						A: 1,
					},
				},
				{
					"ffffffff",
					color.RGBA{
						R: 0xff,
						G: 0xff,
						B: 0xff,
						A: 0xff,
					},
				},
			}

			for _, test := range cases {
				t.Run(
					test.S,
					func(t *testing.T) {
						var c flagutils.RGBA
						err := c.Set(test.S)
						if err != nil {
							t.Fatalf(
								"Set(%q) expected no error, got: %v",
								test.S,
								err,
							)
						}

						if c.String() != strings.ToLower(test.S) {
							t.Errorf("%v.String() expected %s", &c, strings.ToLower(test.S))
						}

						if !reflect.DeepEqual(test.Color, color.RGBA(c)) {
							t.Errorf(
								"Set(%q) expected %+v, got %+v",
								test.S,
								test.Color,
								c,
							)
						}
					},
				)
			}
		},
	)
}

// This example demostrates how to use RGBA in your program.
func ExampleRGBA() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	// Default value
	color := flagutils.RGBA{
		R: 0xff,
	}
	fs.Var(&color, "color", "the color you want")
	fs.Parse([]string{"-color", "ff000080"})
	fmt.Println(&color)
	fmt.Println(color.RGBA())
	// Output:
	// ff000080
	// 65535 0 0 32896
}
