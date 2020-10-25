# Joni Finding Home Key (Go 1.15)

## Quick Start

### Story

Given the square area floor with available and unavailable route. Given the position of the person in the area. There is a missing object in the area and the person is look forward the object with given clue direction.

### Condition

```json
This is what implemented in `main.go`

########
#......#
#.###..#
#...#.##
#o#....#
########

*details:
# is the wall
. is available route
o is person position
```

### Clue

```json
This is what implemented in `main.go`

1. Going North
2. Going East
3. Going South
```

### Usage

``` bash
$ go run main.go
```
*note: you can modify the input on `main.go` file either the `floor`, `person pos` and the `clue` as well.



## App Info

### Authors

Hirzi Nurfakhrian

### Version

1.0.0