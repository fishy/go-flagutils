package flagutils_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"go.yhsif.com/flagutils"
)

func TestOneOf(t *testing.T) {
	var a, b, c flagutils.OneOf

	type table struct {
		Label string
		V     *flagutils.OneOf
	}
	cases := []table{
		{
			"a",
			&a,
		},
		{
			"b",
			&b,
		},
		{
			"c",
			&c,
		},
	}
	checkGroupLen := func(test table, expectedLen int) func(t *testing.T) {
		return func(t *testing.T) {
			group := test.V.Group()
			if len(group) != expectedLen {
				t.Errorf(
					"%v.Group() expected length of %d, got %+v",
					test.Label,
					expectedLen,
					group,
				)
			}
		}
	}

	t.Run(
		"Group",
		func(t *testing.T) {
			flagutils.GroupOneOf(&a, &b, &c)
			for _, test := range cases {
				t.Run(
					test.Label,
					checkGroupLen(test, 3),
				)
			}
		},
	)

	t.Run(
		"Set",
		func(t *testing.T) {
			err := a.Set("true")
			if err != nil {
				t.Fatalf("a.Set(\"true\") expected no error, got %v", err)
			}
			if !a.Bool {
				t.Fatalf("a.Bool expected to be true, got %v", a.Bool)
			}

			err = b.Set("true")
			if err != nil {
				t.Fatalf("b.Set(\"true\") expected no error, got %v", err)
			}
			if !b.Bool {
				t.Fatalf("b.Bool expected to be true, got %v", b.Bool)
			}
			if a.Bool {
				t.Errorf("a.Bool expected to be false, got %v", a.Bool)
			}

			a.Bool = true
			b.Bool = true
			err = c.Set("true")
			if err != nil {
				t.Fatalf("c.Set(\"true\") expected no error, got %v", err)
			}
			if !c.Bool {
				t.Fatalf("c.Bool expected to be true, got %v", c.Bool)
			}
			if a.Bool {
				t.Errorf("a.Bool expected to be false, got %v", a.Bool)
			}
			if b.Bool {
				t.Errorf("b.Bool expected to be false, got %v", b.Bool)
			}
		},
	)

	t.Run(
		"GroupAll",
		func(t *testing.T) {
			// Break the grouping in previous tests
			flagutils.GroupOneOf(&a)
			flagutils.GroupOneOf(&b)
			flagutils.GroupOneOf(&c)
			for _, test := range cases {
				t.Run(
					test.Label,
					checkGroupLen(test, 1),
				)
			}

			fs := flag.NewFlagSet("", flag.ExitOnError)
			fs.Var(&a, "a", "")
			fs.Var(&b, "b", "")
			fs.Var(&c, "c", "")
			flagutils.GroupAllOneOf(fs)
			for _, test := range cases {
				t.Run(
					test.Label,
					checkGroupLen(test, 3),
				)
			}
		},
	)
}

// This example demostrates how to use OneOf in your program.
func ExampleOneOf() {
	// Or use flag.CommandLine instead.
	fs := flag.NewFlagSet("", flag.ExitOnError)
	// Define 3 features, can only run one of them at a time.
	var a, b, c flagutils.OneOf
	fs.Var(&a, "featureA", "Run feature A, unsets featureB and featureC")
	fs.Var(&b, "featureB", "Run feature B, unsets featureA and featureC")
	fs.Var(&c, "featureC", "Run feature C, unsets featureA and featureB")
	// You can also use
	//     flagutils.GroupAllOneOf(fs)
	// if you don't have any other OneOf flag values.
	flagutils.GroupOneOf(&a, &b, &c)
	// TODO: Set other flags here.

	fs.Parse(os.Args[1:])
	// Now your program supports -featureA, -featureB, and -featureC args,
	// and the last one in the command line overrides previous one(s).
	//
	// This means you can define a bash (or other shell) alias to specify a
	// default feature, e.g.:
	//     alias my-cmd='my-cmd -featureA'
	// and still be able to override it in command line:
	//     my-cmd -featureB
	// which is actually
	//     my-cmd -featureA -featureB
	// and runs featureB instead of featureA.

	switch {
	default:
		// Please note that OneOf group only guarantees that at most one of the
		// values will be set to true. It's still possible that all of them are
		// false.
		fmt.Fprintln(
			os.Stderr,
			"You have to specify one of -featureA, -featureB, or -featureC flags.",
		)
		os.Exit(-1)
	case a.Bool:
		// TODO: Implement feature A here.
	case b.Bool:
		// TODO: Implement feature B here.
	case c.Bool:
		// TODO: Implement feature C here.
	}
}
