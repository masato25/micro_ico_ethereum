package views

import (
	"github.com/gin-contrib/multitemplate"
)

var prefixViewFolder = "app/views/"

// http://blog.questionable.services/article/approximating-html-template-inheritance/
func GetMultiRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("common_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"common/index.html")
	r.AddFromFiles("user_register", prefixViewFolder+"layouts/base.html", prefixViewFolder+"users/register.html")
	r.AddFromFiles("user_login", prefixViewFolder+"layouts/base.html", prefixViewFolder+"users/login.html")
	r.AddFromFiles("user_logout", prefixViewFolder+"layouts/base.html", prefixViewFolder+"users/logout.html")
	r.AddFromFiles("token_list", prefixViewFolder+"layouts/base.html", prefixViewFolder+"tokens/index.html")
	r.AddFromFiles("token_dapp", prefixViewFolder+"layouts/base.html", prefixViewFolder+"tokens/dapp.html")
	r.AddFromFiles("create_token", prefixViewFolder+"layouts/base.html", prefixViewFolder+"tokens/create_token.html")
	r.AddFromFiles("wallet_index", prefixViewFolder+"layouts/base.html", prefixViewFolder+"wallets/index.html")
	return r
}
