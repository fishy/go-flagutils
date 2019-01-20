package flagutils

import (
	"flag"
	"strconv"
)

// OneOf is a flag type that works like bool, but can be grouped together.
//
// When you group several OneOf values together,
// (via GroupOneOf or GroupAllOneOf)
// whenever one of them is set to true,
// the rest of the group will be set to false automatically.
//
// It implements flag.Getter interface.
type OneOf struct {
	Bool bool

	group []*OneOf
}

var _ flag.Getter = (*OneOf)(nil)

func (b *OneOf) String() string {
	return strconv.FormatBool(b.Bool)
}

// Set implements flag.Value.
func (b *OneOf) Set(s string) (err error) {
	b.Bool, err = strconv.ParseBool(s)
	if err != nil || !b.Bool {
		return
	}
	// This value was set to true, set all other values in the group to false.
	for _, v := range b.group {
		if v != b {
			v.Bool = false
		}
	}
	return
}

// IsBoolFlag implements flag.Value.
func (b *OneOf) IsBoolFlag() bool {
	return true
}

// Get implements flag.Getter.
func (b *OneOf) Get() interface{} {
	return b.Bool
}

// Group returns a copy of the group b is in.
func (b *OneOf) Group() []*OneOf {
	ret := make([]*OneOf, len(b.group))
	copy(ret, b.group)
	return ret
}

// GroupOneOf groups OneOf values together.
//
// Please note that in case of regrouping,
// it only affects the values in the new group.
// For example say you have a previous group:
//
//     var a, b, c OneOf
//     GroupOneOf(&a, &b, &c)
//
// Then you regroup a and b:
//
//     GroupOneOf(&a, &b)
//
// After the second GroupOneOf, a and b have the group info of each other,
// but c still records the group info of all of them.
// You will need to also regroup c with itself to complete the regrouping:
//
//     GroupOneOf(&c)
func GroupOneOf(values ...*OneOf) {
	for _, v := range values {
		v.group = make([]*OneOf, len(values))
		copy(v.group, values)
	}
}

// GroupAllOneOf group all OneOf flag values inside the FlagSet together.
//
// It finds all OneOf values in the FlagSet, then calls GroupOneOf.
func GroupAllOneOf(fs *flag.FlagSet) {
	values := make([]*OneOf, 0)
	fs.VisitAll(func(f *flag.Flag) {
		if oneof, ok := f.Value.(*OneOf); ok {
			values = append(values, oneof)
		}
	})
	GroupOneOf(values...)
}
