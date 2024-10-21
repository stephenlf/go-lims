package types

type Sample struct {
	Name   string
	Matrix matrix
	Tests  []Test
}

func (s *Sample) AddTests(analyses []*Analysis) {
	for i := 0; i < len(analyses); i++ {
		s.Tests = append(s.Tests, FromAnalysis(analyses[i]))
	}
}
