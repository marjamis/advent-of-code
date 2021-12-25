package advent2019

// Day1 function
func Day1(requireSubFuel bool, modules ...int) (totalFuel int) {
	for _, module := range modules {
		moduleFuel := (module / 3) - 2

		additionalFuel := moduleFuel
		if requireSubFuel {
			needMoreFuel := true
			for needMoreFuel {
				additionalFuel = (additionalFuel / 3) - 2
				if additionalFuel <= 0 {
					needMoreFuel = false
				} else {
					totalFuel += additionalFuel
				}
			}
		}

		totalFuel += moduleFuel
	}
	return
}
