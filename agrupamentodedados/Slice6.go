package main

import (
	"fmt"
)

var segundoslice = []int{}

func main() {
	primeiroslice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(primeiroslice)

	for i := 0; i < len(primeiroslice); i++ {

		fmt.Println(primeiroslice[i])

		segundoslice = append(primeiroslice[:i], primeiroslice[i:]...)

	}
	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)

	for i, _ := range primeiroslice {

		fmt.Println(primeiroslice[i])

		segundoslice = append(primeiroslice[:i], primeiroslice[i:]...)

	}

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)

	//segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

}
