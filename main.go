package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	//scanner()

	if *addTrainee{
		fmt.Println("please enter ID, Surname, Name, Age, TA, QPA, Project. press ENTER after each value")
		var newTrainee Trainee
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		newTrainee.ID, _ =	strconv.Atoi(scanner.Text())
		scanner.Scan()
		newTrainee.Surname = scanner.Text()
		scanner.Scan()
		newTrainee.Name = scanner.Text()
		scanner.Scan()
		newTrainee.Age, _ =	strconv.Atoi(scanner.Text())
		scanner.Scan()
		newTrainee.TA = scanner.Text()
		scanner.Scan()
		newTrainee.QPA = scanner.Text()
		scanner.Scan()
		newTrainee.Project = scanner.Text()

		trainees = append(trainees, newTrainee)
	}

	if *deleteTrainee > -1{
		for i, trainee := range trainees {
			if trainee.ID == *deleteTrainee{
				trainees = append(trainees[:i], trainees[i+1:]...)
			}
		}

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
	i := 0
	for  i < 8{
		scanner.Scan()
		fmt.Println(scanner.Text())
		i++
	}
}