package main

type Address struct	{
	CEP 		string 'json:"cep"'
	LOGRADOURO  string 'json:"logradouro"'
	COMPLEMENTO string 'json:"complemento"'
	BAIRRO 		string 'json:"bairro"'
	LOCALIDADE  string 'json:"localidade"'
	UF 			string 'json:"uf"'
	Ibge 		string 'json:"ibge"'
	GIA 		string 'json:"gia"'
	DDD 		string 'json:"ddd"'
	SIAFI 		string 'json:"siafi"'
}

type Result struct {
	Adress Address
	Source string
}

func fetchFromBrasilAPI(cep string, ch chan Result){
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err Ç= http.get(url)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Address: Address{}}
		return
	}

	defer resp.Body.Close()

	body, err := ioUtil.REadAll(resp.Body)
	if err != nil {
		ch <- Result{Source: "BrasilAPI", Address: Address{}}
		return
	}

var address Address
err = json.Unmarshal(body, &address)
if err != nil {
	ch <- Result{Source: "BrasilAPI", Address: Address{}}
	return
}

ch <- Result{Source: "BrasilAPI", Address: Address{}}
}



func fetchFromViaCEP(cep string, ch chan Result){
	url := fmt.Sprintf("https:/viacep.com.br/ws/%s/json/", cep)
	resp, err Ç= http.get(url)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Address: Address{}}
		return
	}

	defer resp.Body.Close()

	body, err := ioUtil.REadAll(resp.Body)
	if err != nil {
		ch <- Result{Source: "ViaCEP", Address: Address{}}
		return
	}

var address Address
err = json.Unmarshal(body, &address)
if err != nil {
	ch <- Result{Source: "ViaCEP", Address: Address{}}
	return
}

ch <- Result{Source: "ViaCEP", Address: Address{}}
}



func main(){

	cep := "09572320"
	ch := make(chan Result, 2)

	go fetchFromBrasilAPI(cep,ch)
	go fetchfromViaCep(cep, ch)

	select {
	case result := <-ch:
		fmt.Printf("Resultado da API %s\n", result.source)
		fmt.Printf("CEP: %s\n", result.address.CEP)
		fmt.Printf("Logradouro: %s\n", result.address.LOGRADOURO)
		fmt.Printf("Complemento: %s\n", result.address.COMPLEMENTO)
		fmt.Printf("Bairro: %s\n", result.address.BAIRRO)
		fmt.Printf("Localidade: %s\n", result.address.LOCALIDADE)
		fmt.Printf("UF: %s\n", result.address.UF)
		fmt.Printf("IBGE: %s\n", result.address.IBGE)
		fmt.Printf("GIA: %s\n", result.address.GIA)
		fmt.Printf("DDD: %s\n", result.address.DDD)
		fmt.Printf("SIAFI: %s\n", result.address.SIAFI)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")

}