package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melisource/fury_convertir-divisas/cmd/src/model"
	"log"
	"net/http"
)

// ValueDollar GET
type ValueDollar struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func GetDollar(c *gin.Context) {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<

	usd, err := dbc.Get("usd")

	if err != nil {
		log.Fatal(http.StatusNotFound)
	}

	resultDollar := ValueDollar{}

	usd.GetValue(&resultDollar)

	total := model.Total * resultDollar.Value

	totalReal := model.Total
	s := fmt.Sprintf("Valor Atualizado de R$ %v Reais: $ %v Dollares", totalReal, total)

	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}

// ValueEuro GET
type ValueEuro struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func GetEuro(c *gin.Context) {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<

	eur, err := dbc.Get("eur")
	if err != nil {
		log.Fatal(http.StatusNotFound)
	}

	resultEuro := ValueEuro{}

	eur.GetValue(&resultEuro)

	total := model.Total * resultEuro.Value

	totalEuro := model.Total

	s := fmt.Sprintf("Valor Atualizado de R$ %v Reais: € %v Euros", totalEuro, total)

	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}

// ValuePeso GET
type ValuePeso struct {
	Value float64 `json:"value"`
	Coin  string  `json:"coin"`
}

func GetPeso(c *gin.Context) {
	dbc := model.ConnectKvs() /// <<<<<<<<<<<<<<<<<

	eur, err := dbc.Get("eur")
	if err != nil {
		log.Fatal(http.StatusNotFound)
	}

	resultPeso := ValuePeso{}

	eur.GetValue(&resultPeso)

	total := model.Total * resultPeso.Value

	totalPeso := model.Total

	s := fmt.Sprintf("Valor Atualizado de R$ %v Reais: $ %v Pesos", totalPeso, total)

	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}

func PongFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
