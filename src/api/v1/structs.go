package apiv1

type standardArrayInput struct {
	ArrayOfArrays [][]int `json:"array_of_arrays"`
}

type triArrayInput struct {
	Coordinates struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	} `json:"coordinates"`
}
