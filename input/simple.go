package main

import "fmt"

var (
	afirstName, asecondName            string
	bfirstName, bsecondName            string
	cfirstName, csecondName, clastName string
)

func main() {
	fmt.Println("Please enter the firstName and secondName: ")
	fmt.Scan(&afirstName, &asecondName)
	fmt.Printf("firstName is %s, secondName is %s\n", afirstName, asecondName)

	fmt.Println("Please enter the firstName and secondName: ")
	fmt.Scanln(&bfirstName, &bsecondName)
	fmt.Printf("firstName is %s, secondName is %s\n", bfirstName, bsecondName)

	fmt.Println("Please enter the firstName and secondName: ")
	fmt.Scanf("%s, %s", &cfirstName, &csecondName)
	fmt.Printf("firstName is %s, secondName is %s", cfirstName, csecondName)
}
