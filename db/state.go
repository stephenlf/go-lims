package db

import (
	"time"

	"github.com/stephenlf/go-lims/types"
)

type State []types.Sample

func Init() State {
	totalDissolvedSolids := types.Analysis{
		Name: "Total Dissolved Solids",
		Id:   "A0001",
		Sop:  "S0001",
	}
	totalPlateCount := types.Analysis{
		Name: "Total Plate Count",
		Id:   "A0002",
		Sop:  "S0002",
	}

	gel := types.Sample{
		Name:   "Gooey Gel Caps",
		Matrix: types.Gel,
	}
	powder := types.Sample{
		Name:   "Peter's Powdered Plants",
		Matrix: types.Powder,
	}
	liquid := types.Sample{
		Name:   "Larry's Liquid Luck",
		Matrix: types.Liquid,
	}

	gel.AddTests([]*types.Analysis{
		&totalPlateCount,
	})
	powder.AddTests([]*types.Analysis{
		&totalPlateCount,
	})
	liquid.AddTests([]*types.Analysis{
		&totalPlateCount,
		&totalDissolvedSolids,
	})

	gel.Tests[0].Result = types.Result{
		Value: 1200,
		Unit:  "cfu",
	}
	gel.Tests[0].Resources = []types.Resource{
		&types.Equipment{
			Name: "Incubator",
			LastCalibration: time.Date(
				2024, 1, 1, 1, 1, 1, 1, time.Local,
			),
		},
		&types.Media{
			Name: "Tryptic Soy Agar",
			Result: types.Result{
				Unit:  "cfu",
				Value: 3,
			},
		},
	}

	return []types.Sample{
		gel,
		liquid,
		powder,
	}
}
