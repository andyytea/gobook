# Interfaces in Go

Interface types express generalizations or abstractions about the behaviours of other types. Interfaces allow us to write functions that are more flexible and adaptable because they are not tied to the specific details of one particular implementation.

In Go, any object that possesses all of the necessary methods of an interface will implicitly implement an interface. This allows us to write new interfaces that are satisfied by existing concrete types without changing the existing types - which is useful for types defined in packages out of our own control.

## Interfaces as Contracts

A concrete type specifies the exact representation of its values and exposes the intrinsic operations of that representation (arithmetic, indexing, etc.) and any additional behaviours through its methods. In other words, you know exactly what a concrete type is and what you can do with it.

On the other hand, Go has a type known as an interface type - an abstract type. It doesn't expose the representation or internal structure of its values nor does it expose the set of basic operations (it doesn't really care at all).

Whne you have the value of an interface type, you only know what behaviours are provided by its methods. In other words, you know nothing about what it is; you know only what it can do.

A great example is the `fmt` package and its use of the `Stringer` interface.

```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```
Whereas we know that bytes.Buffer implements the io.Writer interface because it has a `Write` method, observe:
```go
package io

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
## Interface Satisfaction

A type satisfies an interface for as long as it possesses all the methods the interface requires. Thus, a type can implement multiple interfaces.

The empty interface `interface{}` specifies no demands on the types that satisfy it, so we immediately have that every type implements the empty interface, allowing us to assign any value to the empty interface.

It is occasionally useful to document and assert the relationship that a type has with the interface that it satisfies.
```go
var w io.Writer = new(bytes.Buffer)

/* OR (to avoid allocating a new variable for asserting relationship) */

var _ io.Writer = (*bytes.Buffer)(nil)
```
We can make intentional choices to name methods with the same name such that they implement convenient interfaces.

A nil interface has both `(type, value) = (nil, nil)`

## Sorting in Go

Go offers a convenient `sort` package with an interface for any type that has methods `Len(), Less(i, j int) bool, Swap(i, j int)`. Additionally, you can conveniently specify more complex conditions for sorting to override the `Less(...)` method.

```go
type customSort struct {
    t []*Track
    less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

sort.Sort(customSort{ tracks, func (x, y *Track) bool {
    if x.Title != y.Title {
        return x.Title < y.Title
    }
    if x.Year != y.Year {
        return x.Year < y.Year
    } 
    if x.Length != y.Length {
        return x.Length < y.Length
    }
    return false
}})
```
