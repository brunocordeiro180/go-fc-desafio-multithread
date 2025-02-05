package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const URL_BRASIL_API string = "https://brasilapi.com.br/api/cep/v1/"
const URL_VIA_CEP string = "http://viacep.com.br/ws/"

func main() {

	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça um cep como argumento.")
		return
	}

	cep := os.Args[1]

	go getCep(URL_BRASIL_API+cep, ch1)
	go getCep(URL_VIA_CEP+cep+"/json", ch2)

	select {
	case msg1 := <-ch1:
		fmt.Printf("BRASIL API: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("VIA CEP %s\n", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout!")
	}
}

func getCep(url string, ch chan []byte) {
	resp, err := http.Get(url)
	// time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	ch <- body

}
