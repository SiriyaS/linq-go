package main

import (
	"fmt"

	. "github.com/ahmetb/go-linq"
)

type Company struct {
	Name    string
	Country string
	City    string
}

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	companies := []Company{
		Company{Name: "Microsoft", Country: "USA", City: "Redmond"},
		Company{Name: "Google", Country: "USA", City: "Palo Alto"},
		Company{Name: "Facebook", Country: "USA", City: "Palo Alto"},
		Company{Name: "Uber", Country: "USA", City: "San Francisco"},
		Company{Name: "Tweeter", Country: "USA", City: "San Francisco"},
		Company{Name: "SoundCloud", Country: "Germany", City: "Berlin"},
	}
	// fmt.Println(companies)

	// find all companies located in USA
	var allUSCompanies []string
	From(companies).Where(func(c interface{}) bool {
		return c.(Company).Country == "USA"
	}).Select(func(c interface{}) interface{} {
		return c.(Company).Name
	}).ToSlice(&allUSCompanies)

	fmt.Println("Companies located in USA")
	fmt.Println(allUSCompanies)


	employees := []Employee{
		Employee{Id: 1, Name: "John", Department: "HR", Salary: 33290},
		Employee{Id: 2, Name: "Max", Department: "HR", Salary: 35900},
		Employee{Id: 3, Name: "Josh", Department: "HR", Salary: 40000},
		Employee{Id: 4, Name: "Jane", Department: "Marketing", Salary: 19300},
		Employee{Id: 5, Name: "Emily", Department: "Marketing", Salary: 25000},
		Employee{Id: 6, Name: "Bob", Department: "HR", Salary: 32100},
		Employee{Id: 7, Name: "Phil", Department: "Management", Salary: 93500},
	}

	// find employee with Salary
	salary := 40000.0
	var staff []string
	From(employees).Where(func(c interface{}) bool {
		return c.(Employee).Salary == salary
	}).Select(func(c interface{}) interface{} {
		return c.(Employee).Name
	}).ToSlice(&staff)

	fmt.Println("Employee with Salary:", salary)
	fmt.Println(staff)

	// Find all employees in HR
	var hrStaffs []string
	From(employees).Where(func(c interface{}) bool {
		return c.(Employee).Department == "HR"
	}).Select(func(c interface{}) interface{} {
		return fmt.Sprintf("Id: %d, Name: %s, Salary: %.2f\n", c.(Employee).Id, c.(Employee).Name, c.(Employee).Salary)
	}).ToSlice(&hrStaffs)

	fmt.Println("Employees in HR Department")
	fmt.Println(hrStaffs)
}
