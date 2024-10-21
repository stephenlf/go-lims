package types

type matrix int

const (
	Liquid matrix = iota
	Gel    matrix = iota
	Powder matrix = iota
)
