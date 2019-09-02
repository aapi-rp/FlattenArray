package apiv1

import (
	"construct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Flatten_standard(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message" : "Bad Request, this endpoint requires json", "status" : 400}`))
		return
	}
	log.Println(string(body))
	var ai standardArrayInput
	err = json.Unmarshal(body, &ai)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message" : "Bad Request, json improperly formatted", "status" : 400}`))
		return
	}

	flatArray := construct.FlattenStandardArrays(ai.ArrayOfArrays...)
	formatArray := strings.Replace(fmt.Sprintf("%v", flatArray), " ", ",", -1)

	w.WriteHeader(200)

	flatResult := fmt.Sprintf(`{"message" : %v, "status" : 200}`, formatArray)
	w.Write([]byte(flatResult))

}

func Flatten_3d(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message" : "Bad Request, this endpoint requires json", "status" : 400}`))
		return
	}
	log.Println(string(body))
	var ta triArrayInput
	err = json.Unmarshal(body, &ta)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message" : "Bad Request, json improperly formatted", "status" : 400}`))
		return
	}

	x := ta.Coordinates.X
	y := ta.Coordinates.Y
	z := ta.Coordinates.Z

	// Create a 3D array
	tryArray := construct.Build3DArray(x, y, z)

	// Flatten the 3D array
	flatten3DArray := construct.Flatten3DArray(tryArray, x, y, z)
	formatArray := strings.Replace(fmt.Sprintf("%v", flatten3DArray), " ", ",", -1)

	w.WriteHeader(200)

	flatResult := fmt.Sprintf(`{"message" : %v, "status" : 200}`, formatArray)
	w.Write([]byte(flatResult))
}
