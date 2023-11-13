A tool similar to linux command `wc`. It tells you the count of bytes, characters, words and newlines in a text file (or stdin).


This is a solution of [this coding challenge](https://codingchallenges.fyi/challenges/challenge-wc/) by John Crickett.

Self-imposed restriction: only use standard libraries.

**Usage**

Build: `go build -o ccwc *.go`

Run with file(s) as input: `./ccwc -wlcm sample_texts/*.txt`

Run with stdin: `./ccwc -wlcm` or `./ccwc -wlcm -`

Run with piped input: `cat sample.txt | ./ccwc -wlcm`

We can also combine file and stdin: `./ccwc -wlcm sample_texts/test.txt -`

**TODOs**

- [ ] Add tests
- [ ] Make Readers mockable
- [ ] Add benchmarks (specially against WholeReader)
- [ ] Do profiling
- [x] Add optimization for large files; done but not as fast as c version
- [x] Add optimization for files where only byte count is needed
- [ ] Add `--help` option.
- [ ] Concurrent reading of multiple inputs
- [ ] Calculate count for all first, then print