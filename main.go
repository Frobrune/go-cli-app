package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main(){
	/*
	fmt.Println("Hello World")
	sortMe := []int{2, 15, 4, 6, 9, 2, 8, 34, 12, 23, 27}
	fmt.Println(bubbleSort(sortMe))
	 */
	/*
	age := flag.Int("Age", 20, "set Age")
	name := flag.String("Name", "Alice", "set Name")
	 */
	people := []string{"Alice", "Bob"}
	addPerson := flag.Bool("addPerson", false, "adds a person to the database")
	showAll := flag.Bool("showAll", false, "display all elements")
	flag.Parse()
	if *addPerson{
		fmt.Println("please provide a name")
		reader := bufio.NewReader(os.Stdin)
		person,_ := reader.ReadString('\n')
		person = strings.Replace(person, "\r\n", "", 1)
		people = append(people, person)
	}
	if *showAll{
		fmt.Println(people)
	}

}


func bubbleSort(a []int) []int{
	var i int = 1
	for i != 0{
		i = 0
		for j := 0; j < (len(a)-1); j++{
			if a[j] > a[j+1] {
				var t int = a[j]
				a[j] = a[j+1]
				a[j+1] = t
				i++
			}
		}
	}
	return a
}