package main

import (
	"flag"
	"fmt"
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
	//initialize slice with trainees
	dummy := Trainee{1, "Hellman", "Alice", 18, "Diffie", "Bob", "Encryption"}
	trainees := []Trainee{dummy}

	//define and parse flags
	addTrainees := flag.Int("addTrainees", 0, "adds multiple trainees to the database")
	showAll := flag.Bool("showAll", false, "show all trainees in a table")
	showOne := flag.Int("showOne", -1, "show trainee with this ID")
	deleteTrainee := flag.Int("deleteTrainee", -1, "delete trainee with this ID")
	flag.Parse()

	//insert trainees into slice in ascending order according to their IDs
	if *addTrainees > 0{
		for n := 0; n < *addTrainees; n++ {
			j := 0
			newTrainee := AddTrainee()
			for i, trainee := range trainees {
				if trainee.ID > newTrainee.ID {
					var temp []Trainee
					temp = append(temp, newTrainee)
					temp = append(temp, trainees[i:]...)
					trainees = append(trainees[:i], temp...)
					j++
				}
			}
			if j == 0 {
				trainees = append(trainees, newTrainee)
			}
		}
	}

	//delete trainee with this ID from slice
	if *deleteTrainee > -1{
		for i, trainee := range trainees {
			if trainee.ID == *deleteTrainee{
				trainees = append(trainees[:i], trainees[i+1:]...)
			}
		}
	}
	//show all trainees in a table
	if *showAll{
		tableStart()
		for _, trainee := range trainees {
			tableContent(trainee)
		}
		tableEnd()
	}

	//show one trainee in a table
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


