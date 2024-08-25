package employee

type Engineer struct {
	Name   string
	Salary int
}

func (e *Engineer) Accept(v IVisitor) {
	v.VisitForEngineer(e)
}
