package datastructures

import "fmt"

// person is..
type person struct {
	age    int
	name   string
	weight float64
	next   *person
}

func TestLikedList() {

	personList := &person{} // creating a pointer to struct, same as &person

	var optn int
	for {
		fmt.Println("\n Select option :- \n\t1.Print List\n\t2.Add to list\n\t3.Exit")
		_, _ = fmt.Scanf("%d", &optn)
		switch optn {
		case 1:
			PrintList(personList)
		case 2:
			AddToList(personList)
		case 3:
			return
		default:
			panic("Exiting..")
		}
	}

}

func PrintList(personList *person) {
	// print("\033[H\033[2J")
	for p := personList; p != nil; p = p.next {
		fmt.Println(p)
	}
	_, _ = fmt.Scanf("%d")
}

func AddToList(personList *person) {
	fmt.Println("Adding..")
	newPerson := &person{}
	fmt.Println("Enter the Name :")
	_, _ = fmt.Scanf("%s", &newPerson.name)
	fmt.Println("Enter the age :")
	_, _ = fmt.Scanf("%d", &newPerson.age)
	fmt.Println("Enter the weight :")
	_, _ = fmt.Scanf("%f", &newPerson.weight)

	AddNode(newPerson, personList)
	fmt.Println(personList)
	fmt.Printf("person %s added to list\n\n", newPerson.name)
}

func AddNode(newPerson, personList *person) *person {
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
