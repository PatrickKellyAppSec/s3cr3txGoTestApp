package main

import (
	//"crypto/tls"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, Thank you for testing s3cr3tx with Go")
	fmt.Println("Please enter your input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Printf("client: could not read the input: %s\n", err)
	}
	var strInput = scanner.Text()
	fmt.Println("Please enter 'e' to encrypt or 'd' to decrypt: ")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	err2 := scanner2.Err()
	if err != nil {
		fmt.Printf("client: could not read the input: %s\n", err2)
	}
	var strDirection = scanner2.Text()
	if strInput != "" && strDirection != "" {
		var strOutput = io(strDirection, strInput)
		fmt.Println(strOutput)
	}

	fmt.Println("Done with s3cr3tx Go Test!")
}
func io(strDirection string, strInput string) string {
	var email string = "pk@gratitech.com"                                                                                                                            //"you@yourdomain.com"
	var APIToken string = "yourS3cr3txAPIToken"
	var AuthCode string =  "yourS3cr3txAuthCode"
	var baseurl string = "https://s3cr3tx.com/Values"
	req, err := http.NewRequest(http.MethodGet, baseurl, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Email", email)
	req.Header.Set("APIToken", APIToken)
	req.Header.Set("AuthCode", AuthCode)
	req.Header.Set("EorD", strDirection)
	req.Header.Set("Input", strInput)
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	if strDirection == "e" {
		var strOutput = "Your encrypted text is : " + string(resBody)
		return strOutput
	}
	if strDirection == "d" {
		var strOutput = "Your decrypted text is : " + string(resBody)
		return strOutput
	}
	return ""
}
