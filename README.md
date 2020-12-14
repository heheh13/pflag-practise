# pflag-practise

plfag is extention over go flag library

go flag library is used for managing cl arguments

## getting started

`https://www.youtube.com/watch?v=v_8584-jm7I`

`https://gobyexample.com/command-line-flags`

`https://github.com/spf13/pflag`

`https://godoc.org/github.com/spf13/pflag#Flag` can view the full doc here

## usage

        wordptr := fmt.int("flagName","flagValue""flagHelp message")
        var next = flag.IntP("next", "n", 23, "a integer pointer")

`go run .` runs with defualt cl args

`go run . --flagname="hello" -n=34` runs with cl args
