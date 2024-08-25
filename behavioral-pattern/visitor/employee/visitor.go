package employee

type IVisitor interface {
	VisitForEngineer(*Engineer)
	VisitForManager(*Manager)
}

type IEmployee interface {
	Accept(IVisitor)
}

type SalaryCalculator struct {
	TotalSalary int
}

func (s *SalaryCalculator) VisitForEngineer(engineer *Engineer) {
	s.TotalSalary += engineer.Salary
}

func (s *SalaryCalculator) VisitForManager(manager *Manager) {
	s.TotalSalary += manager.Salary
}
