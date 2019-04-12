package basics

import (
	"fmt"
)

func Init() {

	// inititalized the variables
	name := "Roxx"
	age := 24

	arrays := []string{"One", "Two", "Three"}

	fmt.Println(name, age)

	for i := 0; i < 5; i++ {
		fmt.Printf("My name is %s, and I am %d years old\n", name, age+i)
	}

	for _, arr := range arrays {

		var input string
		fmt.Scanln(&input)
		fmt.Println(input, arr)
	}

	accounts := make(map[string]float64)

	accounts["Zackman"] = 5000.76
	accounts["Spiderman"] = 455.249

	fmt.Println(accounts)

	fmt.Println(Reverse("Hello"))

}
