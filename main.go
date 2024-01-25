package main

import (
	"fmt"
	"github.com/verdade/go-expert-multithreading/client"
	"github.com/verdade/go-expert-multithreading/dto"
	"os"
	"time"
)

func main() {

	for _, cep := range os.Args[1:] {
		c1 := make(chan dto.BrasilApiDTO)
		c2 := make(chan dto.ViaCepApiDTO)

		go func() {
			res := client.GetCepBrasilAPI(cep)
			c1 <- res
		}()
		go func() {
			res := client.GetViaCepAPI(cep)
			c2 <- res
		}()

		select {
		case result := <-c1:
			fmt.Printf("API BRASIL >> CEP: %s, Cidade: %s, UF: %s, Logradouro: %s\n", result.Cep, result.City, result.State, result.Street)
		case result := <-c2:
			fmt.Printf("API VIACEP >> CEP: %s, Cidade: %s, UF: %s, Logradouro: %s\n", result.Cep, result.Localidade, result.Uf, result.Logradouro)
		case <-time.After(time.Second):
			fmt.Printf("Timeout")
		}
	}

}
