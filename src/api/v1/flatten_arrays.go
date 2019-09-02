package apiv1

import (
	"construct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	w.WriteHeader(200)

	flatResult := fmt.Sprintf(`{"message" : "%v", "status" : 200}`, flatArray)
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
	var ai triArrayInput
	err = json.Unmarshal(body, &ai)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message" : "Bad Request, json improperly formatted", "status" : 400}`))
		return
	}
}
