package parsedh

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DHAPI struct {
	Balance int `json:"balance"`
	Entries []struct {
		ID             string `json:"id"`
		Email          string `json:"email"`
		IPAddress      string `json:"ip_address"`
		Username       string `json:"username"`
		Password       string `json:"password"`
		HashedPassword string `json:"hashed_password"`
		Name           string `json:"name"`
		Vin            string `json:"vin"`
		Address        string `json:"address"`
		Phone          string `json:"phone"`
		DatabaseName   string `json:"database_name"`
	} `json:"entries"`
	Success bool   `json:"success"`
	Took    string `json:"took"`
	Total   int    `json:"total"`
}

func ParseDH(body []byte, outfile string) {
	var jsonAPI DHAPI

	err := json.Unmarshal([]byte(body), &jsonAPI)

	if err != nil {
		log.Fatal("There was an error unmarshaling body", err)
	}

	total := jsonAPI.Total
	balance := jsonAPI.Balance
	dhdata := jsonAPI.Entries

	var user []string
	for _, value := range dhdata {
		user = append(user, "Database: "+value.DatabaseName+" ", "Username: "+value.Username+" ", "Email: "+value.Email+" ", "Password: "+value.Password+" ", "Hash: "+value.HashedPassword+" ", "Phone: "+value.Phone+" ", "Name: "+value.Name+" ", "Address: "+value.Address+" ", "\n")
	}

	// If outfile is not empty will export data to outfile.
	if outfile != "" {
		csvFile, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		for _, entry := range jsonAPI.Entries {
			var row []string
			row = append(row, entry.DatabaseName)
			row = append(row, entry.Username)
			row = append(row, entry.Email)
			row = append(row, entry.Password)
			row = append(row, entry.HashedPassword)
			row = append(row, entry.Name)
			row = append(row, entry.Address)
			row = append(row, entry.Phone)
			writer.Write(row)
		}
		writer.Flush()
	}

	// By default will be displayed to console
	for x := 0; x < len(user); x++ {
		if user[x] == "" {
			continue
		}
		fmt.Println(user[x])
	}

	fmt.Println("[*] Total leaks found: ", total)
	fmt.Println("[*] Your API balance remaining: ", strconv.Itoa(balance))

}

func SetHeader(outfile string) {

	if outfile != "" {
		csvFile, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		header := []string{"Database Name", "Username", "Email", "Password", "Hashed Password", "Name", "Address", "Phone"}
		writer.Write(header)
		writer.Flush()
	}
}
