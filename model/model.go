package model

type Model struct {
	name string
	age  int64
}

func NewModel(name string, age int64) *Model {
	return &Model{
		name: name,
		age:  age,
	}
}

func (m *Model) GetName() string {
	return m.name
}

func (m *Model) GetAge() int64 {
	return m.age
}
