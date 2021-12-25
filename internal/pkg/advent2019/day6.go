package advent2019

import "strings"

type orbitPoint struct {
	Name       string
	Orbits     *orbitPoint
	OrbittedBy orbitPoints
}

type orbitPoints []*orbitPoint

func findNode(op *orbitPoint, find string) *orbitPoint {
	if op.Name == find {
		return op
	}

	for _, p := range op.OrbittedBy {
		if res := findNode(p, find); res != nil && res.Name == find {
			return res
		}
	}

	return nil
}

//As I used findNode multiple times over multiple iterations this is a separate function.
func traverse(op *orbitPoint, count int) (*orbitPoint, int) {
	tmp := 0
	for _, p := range op.OrbittedBy {
		_, val := traverse(p, count+1)
		tmp += val
	}

	return op, count + tmp
}

func iteration(uom *orbitPoint, mapItems []string) (missing []string) {
	for _, mapItem := range mapItems {
		mapPoint := strings.Split(mapItem, ")")
		// fmt.Printf("Initial - Given Parent: %s Orbitter: %s / ", mapPoint[0], mapPoint[1])
		parent := findNode(uom, mapPoint[0])
		if parent != nil {
			new := &orbitPoint{
				mapPoint[1],
				parent,
				nil,
			}
			parent.OrbittedBy = append(parent.OrbittedBy, new)
			// fmt.Printf("Result - Returned Parent: %s Orbiter %s\n", parent.Name, new.Name)
		} else {
			missing = append(missing, mapItem)
		}
	}
	return
}

func getPathFromEndNode(op *orbitPoint) []string {
	pointer := op.Orbits
	var opa []string
	for pointer != nil {
		opa = append(opa, pointer.Name)
		pointer = pointer.Orbits
	}

	last := len(opa) - 1
	for i := 0; i < len(opa)/2; i++ {
		opa[i], opa[last-i] = opa[last-i], opa[i]
	}

	return opa
}

// Day6Part1 function
func Day6Part1(mapItems []string) int {
	//Universal Orbit Map
	uom := &orbitPoint{
		"COM",
		nil,
		nil,
	}

	//Rather than devising some sort of pre-sorting I basically just go over it multiple times to ensure any misses on the first round get picked up on later rounds. I'm sure there is a more efficient way but my sort and other tests didn't really work so...
	for len(mapItems) > 0 {
		mapItems = iteration(uom, mapItems)
	}

	_, count := traverse(uom, 0)
	return count
}

// Day6Part2 function
func Day6Part2(mapItems []string) int {
	//Universal Orbit Map
	uom := &orbitPoint{
		"COM",
		nil,
		nil,
	}

	//Rather than devising some sort of pre-sorting I basically just go over it multiple times to ensure any misses on the first round get picked up on later rounds. I'm sure there is a more efficient way but my sort and other tests didn't really work so...
	for len(mapItems) > 0 {
		mapItems = iteration(uom, mapItems)
	}

	san := getPathFromEndNode(findNode(uom, "SAN"))
	you := getPathFromEndNode(findNode(uom, "YOU"))

	i := 0
	countOfSame := 0
	for true {
		if i < len(you) && i < len(san) {
			if you[i] == san[i] {
				countOfSame++
			}
		} else {
			break
		}
		i++
	}

	return (len(you) - countOfSame) + (len(san) - countOfSame)
}
