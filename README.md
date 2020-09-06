# linq-go
simple query data with Go LINQ

**Go LINQ** is a query package for Go. Essentially it has ability to apply queries to slices and collections using SQL-like methods.

## Installation
```
go get github.com/ahmetb/go-linq/v3
```

## Simple Implementing
※ import the package

#### First example
※ create a `Company` struct and array of it
```
type Company struct {
	Name    string
	Country string
	City    string
}

companies := []Company{
	Company{Name: "Microsoft", Country: "USA", City: "Redmond"},
	Company{Name: "Google", Country: "USA", City: "Palo Alto"},
	Company{Name: "Facebook", Country: "USA", City: "Palo Alto"},
	Company{Name: "Uber", Country: "USA", City: "San Francisco"},
	Company{Name: "Tweeter", Country: "USA", City: "San Francisco"},
	Company{Name: "SoundCloud", Country: "Germany", City: "Berlin"},
}
```

※ now let's find all companies's name that located in USA
```
var allUSCompanies []string
From(companies).Where(func(c interface{}) bool {
	return c.(Company).Country == "USA"
}).Select(func(c interface{}) interface{} {
	return c.(Company).Name
}).ToSlice(&allUSCompanies)

fmt.Println("Companies located in USA")
fmt.Println(allUSCompanies)
```

#### Let's try another example

※ create a `Employee` struct and array of it
```
type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

employees := []Employee{
	Employee{Id: 1, Name: "John", Department: "HR", Salary: 33290},
	Employee{Id: 2, Name: "Max", Department: "HR", Salary: 35900},
	Employee{Id: 3, Name: "Josh", Department: "HR", Salary: 40000},
	Employee{Id: 4, Name: "Jane", Department: "Marketing", Salary: 19300},
	Employee{Id: 5, Name: "Emily", Department: "Marketing", Salary: 25000},
	Employee{Id: 6, Name: "Bob", Department: "HR", Salary: 32100},
	Employee{Id: 7, Name: "Phil", Department: "Management", Salary: 93500},
}
```

※ you can find with specific value, in this case let's find employee's name that has salary = 40000
```
salary := 40000.0
var staff []string
From(employees).Where(func(c interface{}) bool {
	return c.(Employee).Salary == salary
}).Select(func(c interface{}) interface{} {
	return c.(Employee).Name
}).ToSlice(&staff)

fmt.Println("Employee with Salary:", salary)
fmt.Println(staff)
```

※ or you can `return` String format like this, to make it looks neat. Here we gonna find employees work in HR department and return `Id`, `Name` and `Salary` of them
```
var hrStaffs []string
From(employees).Where(func(c interface{}) bool {
	return c.(Employee).Department == "HR"
}).Select(func(c interface{}) interface{} {
	return fmt.Sprintf("Id: %d, Name: %s, Salary: %.2f\n", c.(Employee).Id, c.(Employee).Name, c.(Employee).Salary)
}).ToSlice(&hrStaffs)

fmt.Println("Employees in HR Department")
fmt.Println(hrStaffs)
```
##

> **Referrences**
>
> https://github.com/ahmetb/go-linq
>
> http://blog.ralch.com/tutorial/golang-query-data-with-linq/
