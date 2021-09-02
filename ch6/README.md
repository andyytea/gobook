# Methods in Go

Object-orieted programming was fundamentally about defining methods to express the properties and operations of each data structure so that a client would never need to access an object's representation directly (risk mutation).

The extra parameter in a method declaration attaches the function to the type of the specified parameter.

Since function calls pass by value, if we want to mutate values or avoid copying large objects, we want to attach a function to the pointer receiver rather than the object itself.

Additionally, we can compose types by struct embedding. Similar to how we used anonymous fields to simplify syntax, we can effectively *promote* methods of embedded fields to the "outer" struct.

```go
type ColoredPoint struct {
    Point
    Color color.RGBA
}

red := color.RGBA{255, 0, 0, 255}
blue := color.RGBA{0, 0, 255, 255}
var p = ColoredPoint{Point{1, 1}, red}
var q = ColoredPoint{Point{5, 4}, red}
fmt.Println(p.Distance(q.Point)) // "5"
p.ScaleBy(2)
q.ScaleBy(2)
fmt.Println(p.Distance(q.Point)) // "10"
```
Another key aspect of OOP is known as **encapsulation** or information hiding. This is accomplished by making objects inaccessible to the clients. In Go, we can encapsulate an object by making it a struct.

However, encapsulation can make it inconvenient to apply regular arithmetic operations to special objects.
