package main

import "fmt"

// Person is..
type Person struct {
	age    int
	name   string
	weight float64
	next   *Person
}

func main() {

	personList := &Person{} // creating a pointer to struct, same as &Person

	var optn int
	for {
		fmt.Println("\n Select option :- \n\t1.Print List\n\t2.Add to list\n\t3.Exit")
		fmt.Scanf("%d", &optn)
		switch optn {
		case 1:
			printList(personList)
		case 2:
			addToList(personList)
		case 3:
		default:
			panic("Exiting..")
		}
	}

}

func printList(personList *Person) {
	// print("\033[H\033[2J")
	for p := personList; p != nil; p = p.next {
		fmt.Println(p)
	}
	fmt.Scanf("%d")
}

func addToList(personList *Person) {
	fmt.Println("Adding..")
	newPerson := &Person{}
	fmt.Println("Enter the Name :")
	fmt.Scanf("%s", &newPerson.name)
	fmt.Println("Enter the age :")
	fmt.Scanf("%d", &newPerson.age)
	fmt.Println("Enter the weight :")
	fmt.Scanf("%f", &newPerson.weight)

	addNode(newPerson, personList)
	fmt.Println(personList)
	fmt.Printf("Person %s added to list\n\n", newPerson.name)
}

func addNode(newPerson, personList *Person) *Person {
	if personList == nil {
		return personList
	}
	for p := personList; p != nil; p = p.next {
		if p.next == nil {
			p.next = newPerson
			return personList
		}
	}
	return personList
}
