## Description
Trying my hands on the 1 billion row challenge to see how best I can improve it

### Naive Implementation - 2m15s
First naive implementation took `2:15s`

### First Optimization - 1m35s
- Changed from `strings.Split` to `strings.Cut`.
`strings.Split` will walk through the whole line looking for the separator but `strings.Cut` returns immediately it finds one which is more appropriate here.
This reduced the time by circa `45s`. New time is `1:35s`

### Second optimization - 1m6s
- Since the temperatures are more deterministic, moved away from `strconv.ParseFloat` and wrote my own parser.
- Also, moved from `scanner.Text` to `scanner.Bytes` so I can work with bytes.
Reduced our time by circa `30s` putting as around `1:06s` now

### Third optimizaton - 56s
- Moved from `scanner` to `ReadSlice` since scanner was doing some extra checks which I don't need here.
- Moved from `bytes.Cut` to a custom function `cut` that starts reading the lines from the end so we reach our delimiter faster
