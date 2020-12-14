package main

import (
	"fmt"
	"net"
	"time"

	flag "github.com/spf13/pflag"
)

// package main

// import (
// 	"fmt"

//
// )

// func main() {
// 	// wordPtr := flag.String("word", "foo", "a string")
// 	// numPtr := flag.Int("numb", 42, "an int")
// 	// boolPtr := flag.Bool("fork", false, "a bool")

// 	// var svar string

// 	// flag.StringVar(&svar, "svar", "bar", "a string var")
// 	// helpptr := flag.Int("helper", 45, "helptr")

// 	// var ip *int = flag.Int("myflag","f" 1234, "help message for flagname")
// 	// flag.Parse()
// 	// fmt.Println(*helpptr)
// 	// fmt.Println("word :", *wordPtr)
// 	// fmt.Println("numb: ", *numPtr)
// 	// fmt.Println("fork: ", *boolPtr)
// 	// fmt.Println("svar: ", svar)
// 	// fmt.Println("tail :", flag.Args())

// 	// flag.VarP(&)

// 	var ip *int = flag.Int("ip", 1234, "myflag is a int flag")
// 	var svar string
// 	flag.StringVar(&svar, "svar", "string value", "string help messege")
// 	fmt.Println("all: ", flag.Args())

// 	var next = flag.IntP("next", "n", 23, "a integer pointer")
// 	var native = flag.IntP("native", "a", 32, "a native flag")
// 	flag.Lookup("next").NoOptDefVal = "32"

// 	flag.Parse()
// 	fmt.Println(flag.Args())

// 	fmt.Println("ip :", *ip)
// 	fmt.Println("svar: ", svar)
// 	fmt.Println("all: ", flag.Args())
// 	fmt.Println("next", *next)
// 	fmt.Println("natvie", *native)

// }

func main() {
	valP := flag.BoolSlice("valp", []bool{true, false}, "help")
	ip := flag.IP("ip", net.IPv4(8, 8, 8, 8), "simple ip adress")
	dur := flag.Duration("dur", 1000000000, "idk")
	hello := flag.StringToInt("helo", map[string]int{
		"a": 1,
		"b": 2,
	}, "maps")
	flag.Parse()
	fmt.Println("valp", *valP)
	fmt.Println("dur", *dur)
	fmt.Println("ip", ip)
	fmt.Println("maps", *hello)
	flag.PrintDefaults()
	time.Sleep(*dur)

}
