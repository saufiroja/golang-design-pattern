package main

import "fundamental/golang-design-pattern/behavioral-pattern/visitor/employee"

func main() {
	employees := []employee.IEmployee{
		&employee.Manager{Name: "John", Salary: 1000},
		&employee.Engineer{Name: "Doe", Salary: 500},
	}

	calculator := &employee.SalaryCalculator{}
	for _, employee := range employees {
		employee.Accept(calculator)
	}

	println("Total salary is", calculator.TotalSalary)
}
