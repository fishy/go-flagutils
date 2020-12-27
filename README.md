[![PkgGoDev](https://pkg.go.dev/badge/go.yhsif.com/flagutils)](https://pkg.go.dev/go.yhsif.com/flagutils)
[![Go Report Card](https://goreportcard.com/badge/go.yhsif.com/flagutils)](https://goreportcard.com/report/go.yhsif.com/flagutils)

# Go Flag Utils

This is a Go library that provides a few types you can use with
[flag.Var()](https://pkg.go.dev/flag#Var).

## Sample Code

There are detailed examples for each type
[on pkg.go.dev](https://pkg.go.dev/go.yhsif.com/flagutils#pkg-examples),
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

[BSD License](LICENSE).
