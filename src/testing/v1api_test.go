package testing_

import (
	"base"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_v1_flat_standard(t *testing.T) {
	json := []byte(arraySeries)

	postUrl := fmt.Sprintf("%s://%s/%s", base.GetUrlScheme(), base.GetApiHost(), "v1/flatten_standard")
	fmt.Println()
	fmt.Println("-----API Endpoint flatten_standard test------")
	fmt.Println()
	fmt.Println("Post URL is: ", postUrl)

	resp, request_err := http.Post(postUrl, "application/json", bytes.NewBuffer(json))

	if request_err != nil {
		t.Errorf("Request Failed")
	}

	if resp.StatusCode == 400 {
		t.Errorf("Improperly formatted request")
	}

	body, body_err := ioutil.ReadAll(resp.Body)

	if body_err != nil {
		t.Errorf("Could not read response body from standard api")
	}

	fmt.Println("Response from Endpoint: ", string(body))
	fmt.Println()
	fmt.Println("-----END flatten_standard test------")
	fmt.Println()
}

func Test_v1_flat_3d(t *testing.T) {
	json := []byte(triArray)

	postUrl := fmt.Sprintf("%s://%s/%s", base.GetUrlScheme(), base.GetApiHost(), "v1/flatten_3d")
	fmt.Println()
	fmt.Println("-----API Endpoint flatten_3d test------")
	fmt.Println()
	fmt.Println("Post URL is: ", postUrl)

	resp, request_err := http.Post(postUrl, "application/json", bytes.NewBuffer(json))

	if request_err != nil {
		t.Errorf("Request Failed")
	}

	if resp.StatusCode == 400 {
		t.Errorf("Improperly formatted request")
	}

	body, body_err := ioutil.ReadAll(resp.Body)

	if body_err != nil {
		t.Errorf("Could not read response body from standard api")
	}

	fmt.Println("Response from Endpoint: ", string(body))
	fmt.Println()
	fmt.Println("-----END flatten_3d test------")
	fmt.Println()
}

var arraySeries = `

{

  "array_of_arrays" : [
      [1,2,3,4],
      [5,6,7,8,9],
      [10,11,12,13]
  ]
}
`

var triArray = `
{
  "coordinates" : {
    "x" : 3,
    "y" : 2,
    "z" : 6
  }
}
`
