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
	response, err := http.Get(address)
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