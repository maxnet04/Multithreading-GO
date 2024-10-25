package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Address struct {
	Cep          string `json:"cep"`
	Logradouro   string `json:"logradouro"`
	Complemento  string `json:"complemento"`
	Bairro       string `json:"bairro"`
	Localidade   string `json:"localidade"`
	Uf           string `json:"uf"`
	Ibge         string `json:"ibge"`
	Gia          string `json:"gia"`
	Ddd          string `json:"ddd"`
	Siafi        string `json:"siafi"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type Result struct {
	Adress Address
	Source string
	Error  string
}

func main() {
	cep := "09572320"
	ch := make(chan Result)

	go fetchFromBrasilAPI(cep, ch)
	go fetchFromViaCEP(cep, ch)

	select {
	case result := <-ch:
		if result.Error != "" {
			fmt.Printf("Erro da API %s: %s\n", result.Source, result.Error)
		} else {
			fmt.Printf("Resultado da API %s\n", result.Source)
			fmt.Printf("CEP: %s\n", result.Adress.Cep)
			fmt.Printf("Logradouro: %s\n", result.Adress.Logradouro)
			fmt.Printf("Complemento: %s\n", result.Adress.Complemento)
			fmt.Printf("Bairro: %s\n", result.Adress.Bairro)
			fmt.Printf("Localidade: %s\n", result.Adress.Localidade)
			fmt.Printf("UF: %s\n", result.Adress.Uf)
			fmt.Printf("IBGE: %s\n", result.Adress.Ibge)
			fmt.Printf("GIA: %s\n", result.Adress.Gia)
			fmt.Printf("DDD: %s\n", result.Adress.Ddd)
			fmt.Printf("SIAFI: %s\n", result.Adress.Siafi)

			fmt.Printf("State: %s\n", result.Adress.State)
			fmt.Printf("City: %s\n", result.Adress.City)
			fmt.Printf("Neighborhood: %s\n", result.Adress.Neighborhood)
			fmt.Printf("Street: %s\n", result.Adress.Street)
		}
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}

func fetchFromBrasilAPI(cep string, ch chan Result) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Error: fmt.Sprintf("falha na conexão: %v", err)}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- Result{Source: "BrasilAPI", Error: fmt.Sprintf("status não OK: %d", resp.StatusCode)}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Error: fmt.Sprintf("falha ao ler o corpo da resposta: %v", err)}
		return
	}

	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Error: fmt.Sprintf("falha ao deserializar JSON: %v", err)}
		return
	}

	ch <- Result{Source: "BrasilAPI", Adress: address}
}

func fetchFromViaCEP(cep string, ch chan Result) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Error: fmt.Sprintf("falha na conexão: %v", err)}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- Result{Source: "ViaCEP", Error: fmt.Sprintf("status não OK: %d", resp.StatusCode)}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Error: fmt.Sprintf("falha ao ler o corpo da resposta: %v", err)}
		return
	}

	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Error: fmt.Sprintf("falha ao deserializar JSON: %v", err)}
		return
	}

	ch <- Result{Source: "ViaCEP", Adress: address}
}

func printFields(address Address) {
	if address.Cep != "" {
		fmt.Printf("CEP: %s\n", address.Cep)
	}
	if address.Logradouro != "" {
		fmt.Printf("Logradouro: %s\n", address.Logradouro)
	}
	if address.Complemento != "" {
		fmt.Printf("Complemento: %s\n", address.Complemento)
	}
	if address.Bairro != "" {
		fmt.Printf("Bairro: %s\n", address.Bairro)
	}
	if address.Localidade != "" {
		fmt.Printf("Localidade: %s\n", address.Localidade)
	}
	if address.Uf != "" {
		fmt.Printf("UF: %s\n", address.Uf)
	}
	if address.Ibge != "" {
		fmt.Printf("IBGE: %s\n", address.Ibge)
	}
	if address.Gia != "" {
		fmt.Printf("GIA: %s\n", address.Gia)
	}
	if address.Ddd != "" {
		fmt.Printf("DDD: %s\n", address.Ddd)
	}
	if address.Siafi != "" {
		fmt.Printf("SIAFI: %s\n", address.Siafi)
	}
	if address.State != "" {
		fmt.Printf("State: %s\n", address.State)
	}
	if address.City != "" {
		fmt.Printf("City: %s\n", address.City)
	}
	if address.Neighborhood != "" {
		fmt.Printf("Neighborhood: %s\n", address.Neighborhood)
	}
	if address.Street != "" {
		fmt.Printf("Street: %s\n", address.Street)
	}
}
