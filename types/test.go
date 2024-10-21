package types

type Test struct {
	Name      string
	Analysis  *Analysis
	Result    Result
	Resources []Resource
}

func FromAnalysis(analysis *Analysis) Test {
	return Test{
		Name:     analysis.Name,
		Analysis: analysis,
	}
}
