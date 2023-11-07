package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melisource/fury_convertir-divisas/cmd/src/model"
	"github.com/melisource/fury_convertir-divisas/cmd/src/template"
	gokvsclient "github.com/melisource/fury_go-meli-toolkit-kvsclient"
	"log"
)

func main() {

	// Start servidor
	router := gin.Default()

	//Rotas
	template.SetupRoutes(router)

	SaveDBDollar()
	SaveDBEuro()
	SaveDBPeso()

	router.Run(":8000")
}

type ValueDollar struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func SaveDBDollar() {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<
	err := dbc.Save(gokvsclient.MakeItem("usd", ValueDollar{
		Value: 5.17,
		Coin:  "dollar",
	}))

	if err != nil {
		log.Fatal(err)
	}

	usd, err := dbc.Get("usd")
	if err != nil {
		log.Fatal(err)
	}

	resultDollar := ValueDollar{}

	usd.GetValue(&resultDollar)

	fmt.Println(dbc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Valor Atualizado de $ 1", resultDollar.Coin, ": R$", resultDollar.Value)

}

type ValueEuro struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func SaveDBEuro() {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<
	err := dbc.Save(gokvsclient.MakeItem("eur", ValueEuro{
		Value: 5.21,
		Coin:  "euro",
	}))

	if err != nil {
		log.Fatal(err)
	}

	eur, err := dbc.Get("eur")
	if err != nil {
		log.Fatal(err)
	}

	resultEuro := ValueEuro{}

	eur.GetValue(&resultEuro)

	fmt.Println("Valor Atualizado de € 1 ", resultEuro.Coin, ": R$", resultEuro.Value)
}

type ValuePeso struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func SaveDBPeso() {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<
	err := dbc.Save(gokvsclient.MakeItem("pes", ValuePeso{
		Value: 0.014,
		Coin:  "peso",
	}))

	if err != nil {
		log.Fatal(err)
	}

	pes, err := dbc.Get("pes")
	if err != nil {
		log.Fatal(err)
	}

	resultPeso := ValuePeso{}

	pes.GetValue(&resultPeso)

	fmt.Println("Valor Atualizado de $ 1 ", resultPeso.Coin, ": R$", resultPeso.Value)
}
