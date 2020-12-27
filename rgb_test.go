package flagutils_test

import (
	"flag"
	"fmt"
	"image/color"
	"reflect"
	"strings"
	"testing"

	"go.yhsif.com/flagutils"
)

func TestRGB(t *testing.T) {
	t.Run(
		"Error",
		func(t *testing.T) {
			errors := []string{
				"#ffffff",  // with leading #
				"ffffffff", // out of range
				"Loripsum", // not hex string
			}

			for _, s := range errors {
				t.Run(
					s,
					func(t *testing.T) {
						var c flagutils.RGB
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
					"ff0000",
					color.RGBA{
						R: 0xff,
					},
				},
				{
					"000001",
					color.RGBA{
						B: 1,
					},
				},
				{
					"FfFfFf",
					color.RGBA{
						R: 0xff,
						G: 0xff,
						B: 0xff,
					},
				},
			}

			for _, test := range cases {
				t.Run(
					test.S,
					func(t *testing.T) {
						c := flagutils.RGB{
							A: 0xff,
						}
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

// This example demostrates how to use RGB in your program.
func ExampleRGB() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	// Default value
	color := flagutils.RGB{
		R: 0xff,
	}
	fs.Var(&color, "color", "the color you want")
	fs.Parse([]string{"-color", "0000ff"})
	fmt.Println(&color)
	fmt.Println(color.RGBA())
	// Output:
	// 0000ff
	// 0 0 65535 0
}
