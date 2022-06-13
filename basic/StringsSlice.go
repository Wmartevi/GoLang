package main

import "fmt"

func main() {
	//slice bytes
	s := "ascii, éôÂ^S~lpl*/*89/**, `kskklssm"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v-, %T - %#U - %#X\n", v, v, v, v)

	}

	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v-, %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}

}
