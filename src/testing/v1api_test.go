package testing_

import (
	"base"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func test_v1_flat_standard(t *testing.T) {
	json := []byte(arraySeries)

	postUrl := fmt.Sprintf("%s://%s/%s", base.GetUrlScheme(), base.GetApiHost(), "/v1/")
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

	log.Println(string(body))
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
