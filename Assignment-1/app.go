package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id      string
	Name    string
	Address string
	Job     string
	Reason  string
}

var students []Student = []Student{
	{
		Id:      "A1",
		Name:    "Aqilla",
		Address: "Semarang",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding",
	},
	{
		Id:      "A2",
		Name:    "Noor",
		Address: "Surabaya",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding react",
	},
	{
		Id:      "A3",
		Name:    "Lazuardi",
		Address: "Bali",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding golang",
	},
}

func main() {
	var inputs = os.Args
	
	if !(len(inputs) >= 2) {
        log.Fatalln("Untuk menjalankan program menggunakan perintah go run main.go [id]")
    }

	result, err := FindStudent(inputs[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("\nID :" + result.Id)
	fmt.Printf("\nNama :"+ result.Name)
	fmt.Printf("\nAddress :"+ result.Address)
	fmt.Printf("\nJob :"+ result.Job)
	fmt.Printf("\nReason :"+ result.Reason)
}

func FindStudent(studentId string) (Student, error) {
	for _, value := range students {
		if value.Id == studentId {
			return value, nil
		}
	}

	return Student{}, errors.New("Student not found")
}

