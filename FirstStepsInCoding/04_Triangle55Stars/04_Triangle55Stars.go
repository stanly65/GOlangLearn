package main

import "fmt"

func main()  {

		str := "*"
		for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", str)
		str = str + "*"
	}
}



