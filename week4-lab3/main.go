package main

import (
	"fmt"
	"errors"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:year`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}

func (s *Student) Validate() error  {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year most be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main()  {
	//var st Student = Student{ID:"1", Name:"Pannathorn", Email:"jansaytong_p@silpakorn.edu", Year:3, GPA:4.00}

	//st := Student({ID:"1", Name:"Pannathorn", Email:"jansaytong_p@silpakorn.edu", Year:3, GPA:4.00})

	students := []Student{
		{ID:"1", Name:"Pannathorn", Email:"jansaytong_p@silpakorn.edu", Year:3, GPA:4.00},
		{ID:"2", Name:"Jubjib", Email:"jubjib@gmail.com", Year:3, GPA:3.99},
	}

	newStudent := Student{ID:"3", Name:"eiei", Email:"eiei@gmail.com", Year:3, GPA:2.00}
	students = append(students, newStudent)

	for i , student:= range students {
		fmt.Printf("%d Honor = %v\n", i,  student.IsHonor())
		fmt.Printf("%d Validation = %v\n", i, student.Validate())
	}
	

}