package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//1
	myHobbies := [3]string{"Coding", "Listening music", "Walking"}
	fmt.Println(myHobbies)

	//2
	fmt.Println(myHobbies[0])
	newList := myHobbies[1:3]
	fmt.Println(newList)

	//3
	mainHobbies := myHobbies[:2]
	fmt.Println(mainHobbies)

	//4
	mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)

	//5
	myGoals := []string{"Learning GO.", "Learning all the basics."}
	fmt.Println(myGoals)

	//6
	myGoals[1] = "Learn all the details."
	myGoals = append(myGoals, "Learn all the basics.")
	fmt.Println(myGoals)

	//7
	dynamicProductList := []Product{Product{title: "a", id: "1", price: 35}, Product{title: "b", id: "2", price: 50.25}}
	dynamicProductList = append(dynamicProductList, Product{title: "c", id: "3", price: 7.75})
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
