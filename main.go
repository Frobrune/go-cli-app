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
	dummy := Trainee{1, "Hellman", "Alice", 18, "Diffie", "Bob", "Encryption"}
	var trainees []Trainee
	trainees = append(trainees, dummy)

	addTrainee := flag.Bool("addTrainee", false, "adds a trainee to the database")
	showAll := flag.Bool("showAll", false, "show all trainees in a table")
	showOne := flag.Int("showOne", -1, "show trainee with this ID")
	deleteTrainee := flag.Int("deleteTrainee", -1, "delete trainee with this ID")
	flag.Parse()
	scanner()
	if *addTrainee{
		fmt.Println("please enter ID, Surname, Name, Age, TA, QPA, Project in this format")
		reader := bufio.NewReader(os.Stdin)
		person, _ := reader.ReadString(',')
		person = strings.Replace(person, "\r\n", "", 1)
		//trainees = append(trainees, newTrainee)
	}
	if *deleteTrainee > -1{

	}
	if *showAll{
		tableStart()
		for _, trainee := range trainees {
			tableContent(trainee)
		}
		tableEnd()
	}
	if *showOne > -1{
		traineeID := Trainee{-1, "", "", -1, "", "", ""}
		for _, trainee := range trainees {
			if trainee.ID == *showOne{
				traineeID = trainee
			}
		}
		if traineeID.ID == -1{
			fmt.Println("no such Trainee in Database")
		} else {
			tableStart()
			tableContent(traineeID)
			tableEnd()
		}
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

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}