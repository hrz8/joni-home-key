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

*NOTE:
- you can modify the input on `main.go` file either the `floor`, `person pos` and the `clue` as well.
- the program has 2 method for the solution and implemented 1 of them. the first method is `conventional nested loop` and the other one is `recursive method` which implemented by default. both method working properly. the different is the **recursive one is more flexible for the direction input.**



## App Info

### Authors

Hirzi Nurfakhrian

### Version

1.0.0