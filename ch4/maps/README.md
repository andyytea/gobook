# Maps in Go

In Go, maps are implemented as hash tables. More specifically, a `map` type in go is a reference to a hash table written `map[Key]Value`. All of the keys in a given map are of the same type and all of the values are of the same type, but the keys need not be the same type as the values.

The key type K must be comparable using the standard `==` operator, so that the map can test whether a given key is equal to one already within it. We don't want to compare equality of floating point values and we also don't want to allow NaN as a possible value.

It is an unordered collection of key/value pairs in which all the keys are distinct, and the value associated with a given key can be retrieved , updated, or removed using a constant number of key comparisons in the average case.

```go
ages := map[string]int{
    "alice":    31,
    "charlie":  34,
}
```
Additionally, we can `make` maps
```go
ages := make(map[string]int)
/* OR */
ages := map[string]int{}
```
Note: we can mutate values within a map, but we cannot take the address of a value in the map (since maps can grow in size and lead to invalidated memory)
```go
ages["bob"] = ages["bob"] + 1
ages["bob"] += 1
ages["bob"]++
```
For convenience, we can enumerate all the KV pairs in a map using range:
```go
for k, v := range ages {
    /* ... */
}
```
The book notes that a common pattern for enumerating the map in order is to use the `sort.Strings` function from the `sort` pacakge in the case that keys are strings.
```go
import "sort"

/* [fragment] */
names := make([]string, 0, len(ages))
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range ages {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```
IMPORTANT: a `nil` map is not an empty map - meaning that an uninitialized variable with a map type is a nil map reference by default. This is important only to prevent attempts to store values in a nil map. A common pattern for checking the validity of a key is as follows:
```go
if age, ok := ages["bob"]; !ok { /* ... */ }
```
Similar to slices, we cannot check for equality of two maps without writing a "deep" equality function.

In a language like Python, you get a built in `set` type, but with Go, we can use a map as a set type.

The interesting use of maps in this chapter is in the form of a directed graph:
```go
var graph = make(map[string]map[string]bool)

func addEdge (from, to string) {
    edges := graph[from]
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
}

func hasEdge(from, to string) bool {
    return graph[from][to]
}
```
