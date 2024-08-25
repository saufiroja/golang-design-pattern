package employee

type Manager struct {
	Name   string
	Salary int
}

func (m *Manager) Accept(v IVisitor) {
	v.VisitForManager(m)
}
