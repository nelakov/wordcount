# wordcount

CLI utility that counts words in a sentence. Whitespace is the only delimiter — a word is any sequence of non-whitespace characters.

```
$ ./wordcount "go is awesome"
3
```

## Context

Solution to the "Packages and Modules" exercise (section 3.1) of the [Thank Go!](https://stepik.org/course/96832) Stepik course. Scaffolded as a separate Go module per course instructions — the point of the exercise is to practice `go mod init` and a standalone module layout.

Main course progress and other exercises live in: [elakovnick24/thank-go-practice](https://github.com/elakovnick24/thank-go-practice).

## Build and run

```bash
go build
./wordcount "your sentence here"
```

## Test

```bash
go test
```

Tests print a sha1 hash — the answer code to submit to Stepik.
