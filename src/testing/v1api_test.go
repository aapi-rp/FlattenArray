package testing_

import (
	"base"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func test_v1_flat_standard() {
	json := []byte(arraySeries)

	postUrl := fmt.Sprintf("%s://%s/%s", base.GetUrlScheme(), base.GetApiHost(), "/v1/")
	resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(json))
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
