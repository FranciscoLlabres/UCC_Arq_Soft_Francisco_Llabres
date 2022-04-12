package main

import "github.com/FranciscoLlabres/ej-apis/apiCall"

func main() {
	cats, _ := apiCall.ApiCall()
	for _, c := range cats {
		println("URL:", c.URL, "\nID:", c.ID)
	}
}
