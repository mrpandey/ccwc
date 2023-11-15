A tool similar to linux command `wc`. It tells you the count of bytes, characters, words and newlines in a text file (or stdin).


This is a solution of [this coding challenge](https://codingchallenges.fyi/challenges/challenge-wc/) by John Crickett.

Self-imposed restriction: only use standard libraries.

**Usage**

Build: `go build -o ccwc *.go`

Run with file(s) as input: `./ccwc -wlcm sample_texts/*.txt`

Run with stdin: `./ccwc -wlcm` or `./ccwc -wlcm -`

Run with piped input: `cat sample.txt | ./ccwc -wlcm`

We can also combine file and stdin: `./ccwc -wlcm sample_texts/test.txt -`

A couple test files are available in sample_texts directory. Larger text files can be downloaded from [here](https://zenodo.org/records/8196385) (thanks to [Logpai](https://github.com/logpai)).

**TODOs**

- [ ] Add tests
- [ ] Make stuff mockable
- [ ] Add benchmarks comparing with `wc` and different counting/reading approaches
- [ ] Profiling
- [x] Add optimization for large files; done but not as fast as `wc`
- [x] Add optimization for files where only byte count is needed
- [ ] Add `--help` option
- [x] Concurrent reading of multiple inputs
- [ ] Limit the number of goroutines?
- [ ] Print happens concurrently. How to use the same numWidth for all of them?
- [ ] Is counting words, lines, etc. concurrently faster than calculating them all in one function call?

**Known Bugs**

- [ ] `./ccwc -wlcm - -` is only reading once from input. Process stdin sequentially, not concurrently.