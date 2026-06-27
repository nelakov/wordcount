# wordcount

CLI utility that counts words in a sentence. Whitespace is the only delimiter, so a word is any run of non-whitespace characters.

```
$ ./wordcount "go is awesome"
3
```

## Context

Solution to the "Packages and Modules" exercise (section 3.1) of the [Thank Go!](https://stepik.org/course/96832) Stepik course. It's a separate Go module on purpose: the exercise is about practicing `go mod init` and a standalone module layout.

Main course progress and other exercises live in: [nelakov/thank-go-practice](https://github.com/nelakov/thank-go-practice).

## Build and run

```bash
go build
./wordcount "your sentence here"
```

## Test

```bash
go test
```

Tests print a sha1 hash, the answer code to submit to Stepik.
