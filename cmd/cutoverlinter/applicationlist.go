package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//	"audrahub.com/cutover/cmd/cutoverapi
	"audrahub.com/cutover/cmd/datatypes"
)

// For all Normal templates that have been approved and in the workspace
// Wells Fargo technology
// App ID
// RTO Start - Task Type
// RTO End - Task Type

func getApplications() {
	request := "/runbooks?is_template=true&template_status=approved"
	url := urlcore + request + "&workspace_id=" + sandbox
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error 1")
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error 2")
		}

		var response = datatypes.COTemplate{}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			fmt.Println("Error:  ", err)
		}
		for _, runbook := range response.Data {
			fmt.Println("*********************************")
			fmt.Printf("* RUNBOOKID: %+v\n", runbook.ID)
			fmt.Printf("* RUNBOOK TYPE: %+v\n", runbook.Type)
			fmt.Println("*********************************")

			cfArray, ok := runbook.Attributes.CustomFields.([]interface{})
			if ok {
				for _, record := range cfArray {
					rec, ok := record.(map[string]interface{})
					if ok {
						//fmt.Printf("Record is %+v\n", rec)
						fmt.Printf("name %v\n ", rec["name"])

						str, ok := rec["value"].(string)
						if ok {
							fmt.Printf("value %v\n ", str)
						} else {
							//val, ok = rec["value"].(string)
							for _, val := range rec["value"].([]interface{}) {
								fmt.Printf("value %v\n ", val)

							}
						}
					}
				}
			}

		}

	}
}
