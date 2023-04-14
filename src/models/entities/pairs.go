package entities

import basemodel "go-server/src/share/base/base_model"

type PairsEntity struct {
	basemodel.Model
	Name            string `json:"name"`
	FactoryAddress  string `json:"factoryAddress"`
	InteractAddress string `json:"interactAddress"`
	Network         string `json:"network"`
	M5              string `json:"m5"`
	H1              string `json:"h1"`
	H6              string `json:"h6"`
	H24             string `json:"h24"`
	DexLink         string `json:"dexLink"`
}
