package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Trainee struct{
	ID int
	Surname string
	Name string
	Age int
	TA string
	QPA string
	Project string
}

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
	var dummy Trainee
	dummy.ID = 1
	dummy.Surname = "Hellman"
	dummy.Name = "Alice"
	dummy.Age = 18
	dummy.TA = "Diffie"
	dummy.QPA = "Bob"
	dummy.Project = "Encryption"


	var trainees []Trainee
	trainees = append(trainees, dummy)
	addTrainee := flag.Bool("addTrainee", false, "adds a trainee to the database")
	showAll := flag.Bool("showAll", false, "show all trainees in a table")
	//showOne := flag.Int("showOne", -1, "show trainee with this ID")
	//delete
	flag.Parse()
	if *addTrainee{
		fmt.Println("please provide a name")
		reader := bufio.NewReader(os.Stdin)
		person,_ := reader.ReadString('\n')
		person = strings.Replace(person, "\r\n", "", 1)
		people = append(people, person)
	}
	if *showAll{
		tableStart()
		for _, trainee := range trainees {
			tableContent(trainee)
		}
		tableEnd()
	}

}

func tableStart(){
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
	fmt.Println("| ID  | Nachname      | Vorname       | Alter    | Lernbegleiter    | QPA              | Projekt          |")
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
}

func tableEnd(){
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
}

func tableContent(trainee Trainee){
		id := strconv.Itoa(trainee.ID)
		age := strconv.Itoa(trainee.Age)
		dataPoint := "| "
		dataPoint += id
		for i := len(dataPoint); i < 6; i++{
			dataPoint += " "
		}
		dataPoint += "| " + trainee.Surname
		for i := len(dataPoint); i < 22; i++{
			dataPoint += " "
		}
		dataPoint += "| " + trainee.Name
		for i := len(dataPoint); i < 38; i++{
			dataPoint += " "
		}
		dataPoint += "| " + age
		for i := len(dataPoint); i < 49; i++{
			dataPoint += " "
		}
		dataPoint += "| " + trainee.TA
		for i := len(dataPoint); i < 68; i++{
			dataPoint += " "
		}
		dataPoint += "| " + trainee.QPA
		for i := len(dataPoint); i < 87; i++{
			dataPoint += " "
		}
		dataPoint += "| " + trainee.Project
		for i := len(dataPoint); i < 106; i++{
			dataPoint += " "
		}
		dataPoint += "|"
		fmt.Println(dataPoint)


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