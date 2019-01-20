package flagutils

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"strconv"
)

// RGBA is a flag type that can be used for RGBA colors.
//
// It converts between string format "rrggbbaa" (no leading "#") and RGBA color.
//
// It implements color.Color and flag.Getter interfaces.
type RGBA color.RGBA

var (
	_ flag.Getter = (*RGBA)(nil)
	_ color.Color = (*RGBA)(nil)
)

func (c *RGBA) String() string {
	return fmt.Sprintf("%.2x%.2x%.2x%.2x", c.R, c.G, c.B, c.A)
}

// Set implements flag.Value.
func (c *RGBA) Set(s string) error {
	v, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return err
	}
	if v > 0xffffffff {
		return errors.New("value out of range")
	}
	c.R = uint8((v & 0xff000000) >> 24)
	c.G = uint8((v & 0xff0000) >> 16)
	c.B = uint8((v & 0xff00) >> 8)
	c.A = uint8(v & 0xff)
	return nil
}

// Get implements flag.Getter.
func (c *RGBA) Get() interface{} {
	return (color.RGBA)(*c)
}

// RGBA implements color.Color.
func (c *RGBA) RGBA() (uint32, uint32, uint32, uint32) {
	return (*color.RGBA)(c).RGBA()
}
