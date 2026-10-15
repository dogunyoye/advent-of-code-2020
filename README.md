# Advent of Code 2020

2020's Advent of Code in Go!

Questions can be found [here](https://adventofcode.com/2020)

## Layout

- `day01/` … `day25/` — one `package main` per day
- `data/` — puzzle inputs (`data/*.txt`, gitignored)
- `go.mod` / `go.sum` — single module (`aoc/2020`)

## Run

Requires Go 1.24+.

```sh
go run ./day01
go run ./day11
```

Or build everything:

```sh
go build ./...
go vet ./...
gofmt -l .
```
