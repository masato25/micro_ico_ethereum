package tokens

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Token struct {
	gorm.Model
	TokenName       string
	tokenSymbol     string
	ContractAddress string `gorm:"type:varchar(100);unique_index`
	OwnerAddress    string
}
