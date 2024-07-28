# Simple quiz written in Golang

The program takes questions from a CSV file and the user answers them one by one.

You can specify the path to another CSV file and set a timer using flags.

```
-csv string
        a csv file in the format of 'question,answer'. (default "problems.csv")
-limit int
        the time limit for the quiz in seconds (default 30)
```
Topics covered in this project:

- [encoding/csv](https://pkg.go.dev/encoding/csv) library
- goroutines
- channels
- command-line flags
