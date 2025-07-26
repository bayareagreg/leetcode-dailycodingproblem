package main

import (
	"fmt"
	"strconv"
)

func numberOfWaysToDecodeString(str string) []string {
	if len(str) == 0 {
		return []string{""}
	} else if len(str) == 1 {
		n, err := strconv.Atoi(string(str[0]))
		if err != nil {
			panic("bad input: " + str)
		}
		return []string{string('a' + n - 1)}
	} else {
		ways := numberOfWaysToDecodeString(str[1:])
		n, err := strconv.Atoi(string(str[0]))
		if err != nil {
			panic("bad input: " + str)
		}
		prefix := string('a' + n - 1)
		for i, _ := range ways {
			ways[i] = prefix + ways[i]
		}
		check, err := strconv.Atoi(str[:2])
		if err != nil {
			panic("bad input: " + str)
		}
		if check <= 26 {
			prefix2 := string('a' + check - 1)
			ways2 := numberOfWaysToDecodeString(str[2:])
			for i, _ := range ways2 {
				ways2[i] = prefix2 + ways2[i]
			}
			fmt.Printf("foo: %v\n", ways2)

		}
		return ways
	}
}

func main() {
	fmt.Printf("%v\n", numberOfWaysToDecodeString("111"))
}
