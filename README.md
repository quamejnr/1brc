## Description

Trying my hands on the 1 billion row challenge to see how best I can improve it

## How to run code

There are multiple binaries 1brc\* being named in ascending order as the code is being optimized.
The repo currently contains 1million rows csv which is used to run the benchmarks.
To generate the 1billion rows of data run the command

```sh
python3 assets/create_measurements.py 1_000_000_000
```

You can run any of the binaries as

```sh
./1brc4 -file=<filename>
```

NB: `<filename>` should be replaced by the path to the 1billion rows file generated above
For example;

```sh
./1brc4 -file=measurements.txt
```

## Running Benchmarks

You can run go benchmarks by running

```sh
go test ./... -bench=.
```

These benchmarks uses a copy of the 1 million rows of data.
You can also use the time command on linux to test how fast it runs.

```sh
time ./1brc4 -file=<filename>
```

## Optimizations
These benchmarks were taken from a 2021 M1 Pro 16GB

### Naive Implementation - 2m15s

First naive implementation took `2:15s`

### First Optimization - 1m35s

- Changed from `strings.Split` to `strings.Cut`.
  `strings.Split` will walk through the whole line looking for the separator but `strings.Cut` returns immediately it finds one which is more appropriate here.
  This reduced the time by circa `45s`. New time is `1:35s`

### Second optimization - 1m06s

- Since the temperatures are more deterministic, moved away from `strconv.ParseFloat` and wrote my own parser.
- Also, moved from `scanner.Text` to `scanner.Bytes` so I can work with bytes.
  Reduced our time by circa `30s` putting as around `1:06s` now

### Third optimizaton - 56s

- Moved from `scanner` to `ReadSlice` since scanner was doing some extra checks which I don't need here.
- Moved from `bytes.Cut` to a custom function `cut` that starts reading the lines from the end so we reach our delimiter faster

### Fourth optimization - 49s

- Moved from ReadSlice and called Read directly on our file using a buffer size of 1mb.
