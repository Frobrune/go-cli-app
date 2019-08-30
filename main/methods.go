package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//asks user for input and returns a trainee as output
func AddTrainee() Trainee{
		fmt.Println("please enter ID, Surname, Name, Age, TA, QPA, Project Information for this Trainee. press ENTER after each value")
		var newTrainee Trainee
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newTrainee.ID, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		newTrainee.Surname = scanner.Text()
		scanner.Scan()
		newTrainee.Name = scanner.Text()
		scanner.Scan()
		newTrainee.Age, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		newTrainee.TA = scanner.Text()
		scanner.Scan()
		newTrainee.QPA = scanner.Text()
		scanner.Scan()
		newTrainee.Project = scanner.Text()
		return newTrainee
}

//Start of the Table
func tableStart(){
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
	fmt.Println("| ID  | Nachname      | Vorname       | Alter    | Lernbegleiter    | QPA              | Projekt          |")
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
}

//End of the Table
func tableEnd(){
	fmt.Println("+-----+---------------+---------------+----------+------------------+------------------+------------------+")
}

//input a trainee into a table
func tableContent(trainee Trainee){
	id := strconv.Itoa(trainee.ID)
	age := strconv.Itoa(trainee.Age)
	dataPoint := "| " + fillString(id, 4)
	dataPoint += "| " + fillString(trainee.Surname, 14)
	dataPoint += "| " + fillString(trainee.Name, 14)
	dataPoint += "| " + fillString(age, 9)
	dataPoint += "| " + fillString(trainee.TA, 17)
	dataPoint += "| " + fillString(trainee.QPA, 17)
	dataPoint += "| " + fillString(trainee.Project, 17) + "|"
	fmt.Println(dataPoint)
}

//fillString fills a string with spaces up to fill length.
func fillString(str string, fill int) string{
	for  i := len(str); i < fill; i++{
		str +=" "
	}
	return str
}
