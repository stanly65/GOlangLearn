package main

import ("fmt"
"strings"
)
func main() {

	 var n int
	str1 := "*"
	str2 := " "
	//Прочита въведеното в конзолата число
	fmt.Scan(&n)

	// условие за квадратa
	if 2 <= n && n <= 100 {
		         //горен ред на квадрата със стрна "n"*
			fmt.Println(strings.Repeat(str1,n))
		for i := 0; i < (n-2); i++ {  //поредово страните на квадратра
			fmt.Println(str1+ strings.Repeat(str2,n-2)+str1)
				}
		//горен ред на квадрата със стрна "n"*
		fmt.Println(strings.Repeat(str1,n))
	}


}
