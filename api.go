package shh

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

type API struct {
	engine *gin.Engine
	baseURL string
	port string
}

type BasicCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ToUnfold struct {
	Credentials BasicCredentials  `json:"credentials"`
	FoldedMessage string  `json:"message"`
	PublicKey string `json:"public_key"`
}

type ToFold struct {
	Credentials BasicCredentials  `json:"credentials"`
	Message string `json:"message"`
	PublicKey string `json:"public_key"`
}

func (api *API) run(wizard *Wizard) error {
	genURL := path.Join("/", api.baseURL, "generate")
	foldURL := path.Join("/", api.baseURL, "fold")
	unfoldURL := path.Join("/", api.baseURL, "unfold")

	api.implementGen(genURL, wizard)
	api.implementFold(foldURL, wizard)
	api.implementUnfold(unfoldURL, wizard)

	p := api.port
	if !strings.HasPrefix(api.port, ":") {
		p = ":" + api.port
	}

	return api.engine.Run(p)
}