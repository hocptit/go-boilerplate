package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	basemodel "go-server/src/share/base/base_model"
)

type DexEntity struct {
	// M5              datatypes.JSON `json:"m5" gorm:"type:jsonb;default:'{}'"`
	// H1              datatypes.JSON `json:"h1" gorm:"type:jsonb;default:'{}'"`
	// H6              datatypes.JSON `json:"h6" gorm:"type:jsonb;default:'{}'"`
	// H24             datatypes.JSON `json:"h24" gorm:"type:jsonb;default:'{}'"`

	basemodel.Model
	Name            string  `json:"name"`
	FactoryAddress  string  `json:"factoryAddress" gorm:"column:factory_address"`
	InteractAddress string  `json:"interactAddress" gorm:"column:interact_address"`
	Network         string  `json:"network"`
	M5              DexStat `json:"m5" gorm:"type:jsonb;default:'{}'"`
	H1              DexStat `json:"h1" gorm:"type:jsonb;default:'{}'"`
	H6              DexStat `json:"h6" gorm:"type:jsonb;default:'{}'"`
	H24             DexStat `json:"h24" gorm:"type:jsonb;default:'{}'"`
	DexLink         string  `json:"dexLink"`
}

func (DexEntity) TableName() string {
	return "dexs"
}

type JSON interface {
	Value() (driver.Value, error)
	Scan(value interface{}) error
}

type DexStat struct {
	TotalTXCount string `json:"totalTXCount"`
	TotalVolume  string `json:"totalVolume"`
}

func (dexStat DexStat) Value() (driver.Value, error) {
	return json.Marshal(dexStat)
}

func (dexStat *DexStat) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &dexStat)
}
