package flagutils

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"strconv"
)

// RGB is a flag type that can be used for RGB colors.
//
// It converts between string format "rrggbb" (no leading "#") and RGB color.
//
// It implements color.Color and flag.Getter interfaces.
type RGB color.RGBA

var (
	_ flag.Getter = (*RGB)(nil)
	_ color.Color = (*RGB)(nil)
)

func (c *RGB) String() string {
	return fmt.Sprintf("%.2x%.2x%.2x", c.R, c.G, c.B)
}

// Set implements flag.Value.
func (c *RGB) Set(s string) error {
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return err
	}
	if v > 0xffffff {
		return errors.New("value out of range")
	}
	c.R = uint8((v & 0xff0000) >> 16)
	c.G = uint8((v & 0xff00) >> 8)
	c.B = uint8(v & 0xff)
	c.A = 0
	return nil
}

// Get implements flag.Getter.
func (c *RGB) Get() interface{} {
	return (color.RGBA)(*c)
}

// RGBA implements color.Color.
func (c *RGB) RGBA() (uint32, uint32, uint32, uint32) {
	return (*color.RGBA)(c).RGBA()
}
