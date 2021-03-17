package dhconn

import (
	"bufio"
	"fmt"
	"godehashed/parsedh"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const DURI string = "https://api.dehashed.com/search?query="

func DHConn(apikey, email, name, searchterm, uname, outfile, elist string, phone int) {
	username := strings.Split(strings.Trim(apikey, "\n"), ":")[0]
	password := strings.Split(strings.Trim(apikey, "\n"), ":")[1]

	switch searchterm {
	case "email":
		fmt.Println("[*] We are searching for emails.")
		client := http.Client{}
		req, err := http.NewRequest("GET", DURI+email, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
		req.Header.Add("Accept", "application/json")
		req.SetBasicAuth(username, password)

		if err != nil {
			log.Fatal("Dehash connection error: ", err)
		}
		resp, err2 := client.Do(req)

		if err2 != nil {
			log.Fatal("Error in fetching response ", err2)
		}

		defer resp.Body.Close()

		body, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal("Response Error: ", err3)
		}
		parsedh.ParseDH(body, outfile)

	case "name":
		var searchname string

		searchname = strings.ReplaceAll(name, " ", "+")

		fmt.Println("[*] We are searching for names.")
		client := http.Client{}
		req, err := http.NewRequest("GET", DURI+searchname, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
		req.Header.Add("Accept", "application/json")
		req.SetBasicAuth(username, password)

		if err != nil {
			log.Fatal("Dehash connection error: ", err)
		}
		resp, err2 := client.Do(req)

		if err2 != nil {
			log.Fatal("Error in fetching response ", err2)
		}

		defer resp.Body.Close()

		body, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal("Response Error: ", err3)
		}
		parsedh.ParseDH(body, outfile)
	case "phone":
		fmt.Println("[*] We are searching for phone numbers.")
		client := http.Client{}
		req, err := http.NewRequest("GET", DURI+strconv.Itoa(phone), nil)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
		req.Header.Add("Accept", "application/json")
		req.SetBasicAuth(username, password)

		if err != nil {
			log.Fatal("Dehash connection error: ", err)
		}
		resp, err2 := client.Do(req)

		if err2 != nil {
			log.Fatal("Error in fetching response ", err2)
		}

		defer resp.Body.Close()

		body, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal("Response Error: ", err3)
		}
		parsedh.ParseDH(body, outfile)
	case "username":
		fmt.Println("[*] We are searching for Usernames.")
		client := http.Client{}
		req, err := http.NewRequest("GET", DURI+"username:"+uname, nil)
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
		req.Header.Add("Accept", "application/json")
		req.SetBasicAuth(username, password)

		if err != nil {
			log.Fatal("Dehash connection error: ", err)
		}
		resp, err2 := client.Do(req)

		if err2 != nil {
			log.Fatal("Error in fetching response ", err2)
		}

		defer resp.Body.Close()

		body, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Fatal("Response Error: ", err3)
		}
		parsedh.ParseDH(body, outfile)
	case "list":

		fmt.Print("[*] Going into List mode, Will add a 2 second time delay to prevent blacklist.")
		fmt.Print("[*] NOTE: Can take a long time depeding on the size of the list and will use ALOT OF CREDITS.\n")
		time.Sleep(3 * time.Second)
		//opens list of files
		file, err := os.Open(elist)
		if err != nil {
			fmt.Println("Cannot read file", err)
			os.Exit(0)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var dhlist []string

		for scanner.Scan() {
			dhlist = append(dhlist, scanner.Text())
		}

		file.Close()

		for _, line := range dhlist {
			fmt.Println("[*] We are searching for emails.")
			client := http.Client{}
			req, err := http.NewRequest("GET", DURI+line, nil)
			req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
			req.Header.Add("Accept", "application/json")
			req.SetBasicAuth(username, password)

			if err != nil {
				log.Fatal("Dehash connection error: ", err)
			}
			resp, err2 := client.Do(req)

			if err2 != nil {
				log.Fatal("Error in fetching response ", err2)
			}

			defer resp.Body.Close()

			body, err3 := ioutil.ReadAll(resp.Body)
			if err3 != nil {
				log.Fatal("Response Error: ", err3)
			}
			parsedh.ParseDH(body, outfile)

			fmt.Println("[*] Delaying.. Please wait.")
			time.Sleep(2 * time.Second)
		}

	default:
		fmt.Println("Please enter in a valid search term, 'email' or 'name'.")
		os.Exit(0)
	}

}
