package client

import (
	"encoding/json"
	"github.com/verdade/go-expert-multithreading/dto"
	"io"
	"log"
	"net/http"
	"strings"
)

const URI_BRASIL_API = "https://brasilapi.com.br/api/cep/v1/"
const URI_VIACEP_API = "http://viacep.com.br/ws/CEP/json/"

func GetCepBrasilAPI(cep string) dto.BrasilApiDTO {
	var response dto.BrasilApiDTO

	req, err := http.NewRequest("GET", URI_BRASIL_API+cep, nil)
	if err != nil {
		log.Println(">>> Erro New Request  %v\n <<<", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(">>> Erro DefaultClient 2%v\n <<<", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(">>> Erro ReadAll %v\n <<<", err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(">>> Erro Unmarshal  4 %v\n <<<", err)
	}

	return response
}

func GetViaCepAPI(cep string) dto.ViaCepApiDTO {
	var response dto.ViaCepApiDTO
	req, err := http.NewRequest("GET", strings.Replace(URI_VIACEP_API, "CEP", cep, 1), nil)
	if err != nil {
		log.Println(">>> Erro New Request  %v\n <<<", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(">>> Erro DefaultClient 2%v\n <<<", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(">>> Erro ReadAll %v\n <<<", err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(">>> Erro Unmarshal  4 %v\n <<<", err)
	}

	return response
}
