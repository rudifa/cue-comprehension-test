# CUE comprehension timing test

These tests show that the comprehension time of CUE is quadratic in the data size.

## testscript test

Run `testscript -v txtar/test-comprehension.txtar` and observe the time it takes to run vs. data size.

```
WORK=$WORK
PATH=...
...
# exec cat data/schema.cue (2.044s)
> go mod tidy
[stderr]
go: finding module for package github.com/google/uuid
go: downloading github.com/google/uuid v1.3.1
go: found github.com/google/uuid in github.com/google/uuid v1.3.1

> go run main.go -size=200
[stdout]
Generated test data of size 200 and wrote it to file data/data.json

# exec head data/data.json (39.859s)
> exec time cue eval data/schema.cue data/data.json -e '#dict' -f -o data.cue
[stderr]
        0.10 real         0.04 user         0.01 sys
> go run main.go -size=2000
[stdout]
Generated test data of size 2000 and wrote it to file data/data.json

> exec time cue eval data/schema.cue data/data.json -e '#dict' -f -o data.cue
[stderr]
        0.41 real         0.43 user         0.02 sys
> go run main.go -size=20000
[stdout]
Generated test data of size 20000 and wrote it to file data/data.json

> exec time cue eval data/schema.cue data/data.json -e '#dict' -f -o data.cue
[stderr]
       38.19 real        38.43 user         0.26 sys
PASS
```

## manual test

Run `go run main.go gen --size 1000` to create a data file of that size in the `.tmp` directory.

Run `time cue eval schema.cue .tmp/testdata.1000.json -e '#dict' -f -o .tmp/data.1000.cue` to create a CUE file from the data
and observe the time it takes to run vs. data size.

```
cue eval schema.cue .tmp/testdata.1000.json -e '#dict' -f -o   0.17s user 0.03s system 82% cpu 0.236 total
```
