[![GoDoc](https://godoc.org/github.com/fishy/go-flagutils?status.svg)](https://godoc.org/github.com/fishy/go-flagutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/fishy/go-flagutils)](https://goreportcard.com/report/github.com/fishy/go-flagutils)

# Go Flag Utils

This is a Go library that provides a few types you can use with
[flag.Var()](https://godoc.org/flag#Var).

## Sample Code

There are detailed examples for each type
[on GoDoc](https://godoc.org/github.com/fishy/go-flagutils#pkg-examples),
but here's a quick example:

```go
// Default to red
var color flagutils.RGB{
	R: 0xff,
}
flag.Var(
	&color,
	"color",
	"The color to use",
)
var featureA, featureB flagutils.OneOf
flagutils.GroupOneOf(&featureA, &featureB)
flag.Var(
	&featureA,
	"featureA",
	"Run feature A, unsets featureB",
)
flag.Var(
	&featureB,
	"featureB",
	"Run feature B, unsets featureA",
)
flag.Parse()
switch {
case featureA.Bool:
	// Run feature A
case featureB.Bool:
	// Run feature B
}
```

## License

[BSD License](https://github.com/fishy/go-flagutils/blob/master/LICENSE).
