package tokens

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Token struct {
	gorm.Model
	TokenName       string
	TokenSymbol     string
	ContractAddress string `gorm:"type:varchar(100);unique_index`
	OwnerAddress    string
	TotallSupply    int64
	Abi             string
}

func (self *Token) IDString() string {
	return strconv.Itoa(int(self.ID))
}
