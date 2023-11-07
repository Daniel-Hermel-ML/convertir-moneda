package model

import (
	gokvsclient "github.com/melisource/fury_go-meli-toolkit-kvsclient"
)

var Total = 100.0

func ConnectKvs() gokvsclient.Client {
	kvsConfig := gokvsclient.MakeKvsConfig()

	Client := gokvsclient.MakeKvsClient("convertirdivisaskvs", kvsConfig)

	if Client == nil {
		panic("Não foi possível conectar com KVS")
	}
	return Client

}
