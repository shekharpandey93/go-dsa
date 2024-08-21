package main

import (
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
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "")
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
