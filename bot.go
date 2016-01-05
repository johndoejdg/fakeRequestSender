package main

import "fmt"
import "net/http"
import "os"
import "io/ioutil"
import "math/rand"
import "time"

func main() {
	counter := 0
	counter2 := 0
	for counter < 1 {
		if counter2 < 2 {
			counter2++
			go sendFakeRequest()
		} else {
			counter2 = 0
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func sendFakeRequest() {
	address := generateIp();
	fmt.Println(address)
	doResponse(address)
}

func generateIp() (string)  {
	part1 := rand.Intn(250)
	part2 := rand.Intn(250)
	part3 := rand.Intn(250)
	part4 := rand.Intn(250)
	ip := fmt.Sprint(part1,".",part2,".",part3,".",part4)
	schema := "http"
	if part1 > 150 {
		schema += "s"
	}
	address := schema+"://"+ip

	return address
}

func doResponse(address string)  {
	userAgent := getUserAgent()
	client := &http.Client{}
	req, err := http.NewRequest("GET", address, nil)
	req.Header.Set("User-Agent", userAgent)
	if err != nil {
		fmt.Printf("%s", err)
		go sendFakeRequest()
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		go sendFakeRequest()
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}

func getUserAgent() (string) {
	a := rand.Intn(100)
	var userAgent string
	if (a > 5 && a < 10) {
		userAgent = "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; WOW64; Trident/4.0; SLCC1)"
	} else if (a > 10 && a < 20) {
		userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25"
	} else if (a > 20 && a < 30) {
		userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:33.0) Gecko/20120101 Firefox/33.0"
	} else if (a > 30 && a < 40) {
		userAgent = "Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14"
	} else if (a > 40) {
		userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/7046A194A"
	}

	return userAgent
}