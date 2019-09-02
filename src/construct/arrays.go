package construct

import "fmt"

func array3D(x, y, z int) [][][]*Construct {
	// Level 1
	tri := make([][][]*Construct, x)
	for i := 0; i < x; i++ {
		// Level 2
		tri[i] = make([][]*Construct, y)
		for j := 0; j < y; j++ {
			// Level 3
			tri[i][j] = make([]*Construct, z)
			for k := 0; k < z; k++ {
				tri[i][j][k] = new(Construct)
				tri[i][j][k].value = i + j + k
			}
		}
	}
	return tri
}

func Build3DArray(X, Y, Z int) [][][]*Construct {
	gen3DArray := array3D(X, Y, Z)
	return gen3DArray
}

func Flatten3DArray(tri [][][]*Construct, X, Y, Z int) []int {

	fmt.Println("| ---      Robert's Lil Demo   ---|")
	fmt.Println("| --- 3D array graphic display ---|")
	fmt.Println("|            welcome =)           |")
	create := make([]int, 0)

	for i := 0; i < X; i++ {
		fmt.Println(" ---------------------------------")
		fmt.Println("x axis", i, "  |z axis")
		fmt.Println(" ---------------------------------")
		create = append(create, i)
		for j := 0; j < Y; j++ {
			fmt.Printf("y axis:    |")
			for k := 0; k < Z; k++ {
				fmt.Printf("%d ", tri[i][j][k].value)
				create = append(create, tri[i][j][k].value)
			}
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println("3D flat array: ", create)
	fmt.Println()
	fmt.Println("3D Flat dedupe array: ", deDupe(create))

	return create
}

func deDupe(s []int) []int {
	bk := make(map[int]bool)
	li := []int{}
	for _, entry := range s {
		if _, value := bk[entry]; !value {
			bk[entry] = true
			li = append(li, entry)
		}
	}
	return li
}

func FlattenStandardArrays(intarrays ...[]int) []int {
	create := make([]int, 0)
	for i := 0; i < len(intarrays); i++ {
		for j := 0; j < len(intarrays[i]); j++ {
			value := intarrays[i][j]
			create = append(create, value)
		}
	}
	return create
}
