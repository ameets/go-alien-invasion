# go-alien-invasion

## Alien Invasion
go-alien-invasion simulates an alien invasion.

The simulation runs until all the aliens have beendestroyed, no cities are remaining, or each alien has moved at least 10,000 times.

## Building from source

* [Go](https://golang.org/doc/install) version 1.8, with $GOPATH set to your
  preferred directory
  
## Installation
```sh
$ git clone https://github.com/ameets/go-alien-invasion.git
$ cd go-alien-invasion
$ go build
```

Start simulation:
```sh
$ ./go-alien-invasion -n <numAliens>
```

Run tests:
```sh
$ go test ./...
```