A tool similar to linux command `wc`. It tells you the count of bytes, characters, words and lines in a text file (or stdin).


This is a solution of [this coding challenge](https://codingchallenges.fyi/challenges/challenge-wc/) by John Crickett.

Self-imposed restriction: only use standard libraries.

**Usage**

Build: `go build -o ccwc *.go`

Run with a file: `./ccwc -wlcm sample.txt`

Run with stdin: `./ccwc -wlcm` or `./ccwc -wlcm -`

Run with pipe input: `cat sample.txt | ./ccwc -wlcm`

**TODOs**

- [ ] Add tests
- [ ] Make Readers mockable
- [ ] Add benchmarks
- [ ] Add optimization for large files
- [ ] Add `--help` option.


**Known Bugs**

- [ ] stdin gives correct result, but piping a text file gives wrong numbers (except word count)