package types

import "time"

type Resource interface {
	Label() string
	IsValid() bool
}

type Equipment struct {
	Name            string
	LastCalibration time.Time
}

func (e *Equipment) Label() string {
	return e.Name
}

func (e *Equipment) IsValid() bool {
	return e.LastCalibration.After(time.Now().AddDate(-1, 0, 0))
}

type Media struct {
	Name   string
	Result Result
}

func (m *Media) Label() string {
	return m.Name
}

func (m *Media) IsValid() bool {
	return m.Result.Value < 10
}
