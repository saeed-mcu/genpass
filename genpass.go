package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func usage() {
	fmt.Println(" Usage : ")
	fmt.Println(string(os.Args[0]) + "  \"cccdddcccddd\"")
	fmt.Println(" ")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := func() string {
		chars := "abcdefhjkmnprstuvwxyz"
		return string(chars[rand.Intn(len(chars))])
	}
	d := func() string {
		digits := "23456789"
		return string(digits[rand.Intn(len(digits))])
	}
	cc := func(s string) string {
		re := regexp.MustCompile("c")
		s = re.ReplaceAllStringFunc(s, func(_ string) string {
			return c()
		})
		re = regexp.MustCompile("d")
		s = re.ReplaceAllStringFunc(s, func(_ string) string {
			return d()
		})
		return s
	}

	var inputString = ""
	if len(os.Args) < 2 {
		inputString = "dddcccdddcccddd"
	} else if len(os.Args) == 2 {
		inputString = os.Args[1]
	} else {
		usage()
		return
	}

	fmt.Println(cc(inputString))
}
