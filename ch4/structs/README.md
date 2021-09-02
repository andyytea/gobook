# Structs in Go

A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity. Each value is called a `field`.

```go
type Employee struct {
    ID          int
    Name        string
    Address     string
    DoB         time.Time
    Position    string
    Salary      int
    ManagerID   int
}

var dilbert Employee

// Mutation
dilbert.Salary -= 5000
position := &dilbert.Position
*position = "Senior " + *position

// Returning pointers to a struct

func EmployeeByID(id int) *Employee { /* ... */ }

id := dilbert.ID
EmployeeByID(id).Salary = 0
```

## Struct Literals

For readability, it's good practice to specify the field names that are being set so that all else can be considered omissions (zero values).

Literals can be useful in making comparisons:

```go
type address struct {
    hostname    string
    port        int
}

hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```

## Struct Embeddings

For the sake of simplification, we can define anonymous fields - a field with a type, but no name. Compare the following:

```go
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}

var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

```go
type Point struct {
    X, Y int
}

type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}

var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```

Both struct definitions are the same at compile time. However, note that there is no convenient syntax for defining the a struct literal

```go
w = Wheel{8, 8, 5, 20}                      // compilation error: unknown fields
w = Wheel{X:8, Y:8, Radius:5, Spokes:20}    // compilation error: unknown fields
w = Wheel {
    Circle: Circle {
        Point:  Point{X:8, Y:8},
        Radius: 5,
    },
    Spokes: 20,
}
```

## JSON in Go

Go has a great package for encoding and decoding data in common formats (JSON, ASN1, XML, etc.). Additionally, struct types are key in working with JSON data.
```go

type Movie struct {
    Title   string
    Year    int     `json:"released"`
    Color   bool    `json:"color,omitempty"`
    Actors  []string
}
```
The process of converting a Go data structure to JSON is known as *marshaling*. 
```go
data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSON marshaling failes: %s", err)
}
fmt.Printf("%s\n", data)
```
Marshal produces a byte slice containing a very long string with no extraneous white space.
```
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true," Actors":["Steve McQueen","Jacqueline Bisset"]}]
```
The field tags are a string of meta data associated at compile time with the field of a struct. In addition to marshaling data, we can also unmarshal data into a Go struct. (since the meta data complicates things)
