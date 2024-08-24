package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	//url := "https://amazon-products1.p.rapidapi.com/search?country=US&query=MacBook%2BPro&page=1&categoryId=165793011"
	url := "https://stgobain-dev2-fra-core-api.multicloud-ibm.com/api/v1/aiops/inventory/resources/list/view"
	payload := strings.NewReader(`{
"page": 1,
"size": 10,
"sortBy": "id",
"orderBy": "asc",
"searchBy": "",
"provider": {},
"sourceTypes": [],
"applications": [],
"teams": [],
"notificationEvent": "GLOBAL",
"appCategories": [],
"localFilter": {
"topInsight": "",
"breakdown": {
"key": "",
"value": ""
},
"geoMap": {
"country": {
"key": "",
"value": ""
},
"city": {
"key": "",
"value": ""
}
}
},
"application": {
"orderBy": "",
"sortBy": "",
"searchBy": ""
},
"resource": {
"orderBy": "asc",
"page": 1,
"size": 10,
"sortBy": "id",
"searchBy": ""
},
"associated": {
"orderBy": "asc",
"page": 1,
"size": 10,
"sortBy": "id",
"searchBy": "",
"insights": ""
}
} `)
	req, _ := http.NewRequest(context.TODO(), "POST", url, payload)

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImtkTUdzVlNXbnVoNkRUMGloMEg5SUtqZ0hic3JjOHZ5IiwidHlwIjoiSldUIiwieDV0IjoiRnRaaElsYkJNWnhFNlN0NExwRml3MEI1X3VJIiwieDV0I1MyNTYiOiIxT002RDJxUkYyYW1DdTBkQzdPYWVueTBCNGdGbXU0NEZPVGJCa1UtYlV3In0.eyJpZCI6IjYyZjEwM2Y5M2IyZDA2NjlkYTBlMGFjYiIsInVzZXJpZCI6IjYyZjEwM2Y5M2IyZDA2NjlkYTBlMGFjYiIsInRva2VuaWQiOiI2MzA1Zjk5MDZhMjE5ZGIwNjI0OWU0YTEiLCJuYW1lIjoiU0hBRklZQSBMT01BREEiLCJlbWFpbCI6IlNoYWZpeWEuTG9tYWRhQGt5bmRyeWwuY29tIiwiZ2l2ZW5fbmFtZSI6IlNIQUZJWUEiLCJmYW1pbHlfbmFtZSI6IkxPTUFEQSIsInByZWZlcnJlZF91c2VybmFtZSI6IlNoYWZpeWEuTG9tYWRhQGt5bmRyeWwuY29tIiwiZ3JvdXBzIjpbIjVmOTAxNjYzNzNjMjFiNjdkMDFmNTBmNCJdLCJyb2xlcyI6W3siciI6Im1lbWJlciIsIm8iOiI1ZjkwMTY2NTczYzIxYjY3ZDAxZjUwZjUifV0sInRlbmFudCI6IjVmOTAxNjYzNzNjMjFiNjdkMDFmNTBmNCIsIm1hc3RlciI6ZmFsc2UsInRlbmFudF90eXBlIjoiY2xpZW50Iiwic291cmNlX3RlbmFudCI6IiIsImludGVybmFsIjpmYWxzZSwicmVmcmVzaCI6IjVmOTAxNjYzNzNjMjFiNjdkMDFmNTBmNC5NZU5aMWo3b3RHMzI0Z0RzZE1sd3RDZmEyYVdZaXFmb2JpZWcxWTJ1RTF3eEFwV0VjSXlrZVJSeXprMk1NT1R5IiwiYWNjZXNzIjoiNWY5MDE2NjM3M2MyMWI2N2QwMWY1MGY0LkZEa3FkNUhWQnJHVUhJWmIxQUVOWVFnUm11cEFOQldlT1gyUWsyUml2dTFXUmNuY003dnhxVEZLRVlnV0JZNWQiLCJ0eXBlIjoiQmVhcmVyIiwib3JnIjoiNWY5MDE2NjU3M2MyMWI2N2QwMWY1MGY1IiwiaXNzIjoiaHR0cHM6Ly9zdGdvYmFpbi1kZXYyLWZyYS1jb3JlLm11bHRpY2xvdWQtaWJtLmNvbSIsInN1YiI6IjYyZjEwM2Y5M2IyZDA2NjlkYTBlMGFjYiIsImV4cCI6MTY2MTM0MzA5MiwibmJmIjoxNjYxMzM1OTUyLCJpYXQiOjE2NjEzMzU5NTIsImp0aSI6ImVkYjVhZDg3LTVkZDItNDNmZi04MmZhLTY2ZTBkY2M3Nzg2YiIsImlkcCI6IjVmMzY4ZGQ0MGFlZDMxMWExMjdlZGUxZCIsIm92ZXJzaXplIjpmYWxzZSwiaWRwYXV0aCI6ZmFsc2UsImlkcG5hbWUiOiJJQk1JRCIsInN1Yl90eXBlIjoiaWFtX2lkIn0.LUxRxFRtjzPsg2_UiVCRgNBOP8c6bG5M_tHT4LyU0YGcHk1C5kEwjYSlpYE1q9xjDG7j146o0E5HwtxzC3x8fG1PduvThp-r0WIpb4sZ-lKV0Nds-u6itVsy1ya_V61fZ0R8ABbe5k5I9DkS7IGR-QcabuImqNXFdqZ1h563xX3NBQfa5vOZLmfR21MPYkxJ-pKeUhirIx1qUnFNZg_a5HeQy5APPq83UqMEWhyO6yKwfZ9Qbs867HgcvDfUfAvtxFhWh163UCeSz5kH59GCobnR-nchIG7z7I0o63UgBK7gp8czluA-pcG7WooTR1xevefzj0j48Rsn7jPidtb8jQ")
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func formPayload(size, page int) []byte {
	return []byte(fmt.Sprintf(`{
"page": %v,
"size": %v,
"sortBy": "id",
"orderBy": "asc",
"searchBy": "",
"provider": {},
"sourceTypes": [],
"applications": [],
"teams": [],
"notificationEvent": "GLOBAL",
"appCategories": [],
"localFilter": {
"topInsight": "",
"breakdown": {
"key": "",
"value": ""
},
"geoMap": {
"country": {
"key": "",
"value": ""
},
"city": {
"key": "",
"value": ""
}
}
},
"application": {
"orderBy": "",
"sortBy": "",
"searchBy": ""
},
"resource": {
"orderBy": "asc",
"page": %v,
"size": %v,
"sortBy": "id",
"searchBy": ""
},
"associated": {
"orderBy": "asc",
"page": %v,
"size": %v,
"sortBy": "id",
"searchBy": "",
"insights": ""
}
}`, page, size, page, size, page, size))
}
