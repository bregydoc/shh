package shh

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) implementGen(path string, wizard *Wizard) {
	api.engine.POST(path, func(c *gin.Context) {
		pair, err := wizard.craftNewDefaultPair()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		cred  := new(BasicCredentials)

		if err = c.ShouldBindJSON(cred); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		token := cred.Username + ":" + cred.Password
		if err = wizard.verifyToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err = wizard.store.RegisterNewPair(token, pair); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"public_key":  pair.PublicKeyPem,
		})
	})
}

func (api *API) implementFold(path string, wizard *Wizard) {
	api.engine.POST(path, func(c *gin.Context) {
		if !wizard.fullAvailableAPI {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "resource not usable",
			})
			return
		}

		req  := new(ToFold)

		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		token := req.Credentials.Username + ":" + req.Credentials.Password
		if err := wizard.verifyToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}


		message, err := wizard.encrypt(req.PublicKey, req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}


		c.JSON(http.StatusOK, gin.H{
			"message":  message,
		})
	})
}

func (api *API) implementUnfold(path string, wizard *Wizard) {
	api.engine.POST(path, func(c *gin.Context) {
		if !wizard.fullAvailableAPI {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "resource not usable",
			})
			return
		}
		
		req  := new(ToUnfold)

		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		token := req.Credentials.Username + ":" + req.Credentials.Password
		if err := wizard.verifyToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		pair, err := wizard.store.GetPair(token, req.PublicKey)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		log.Println(pair)

		encryptedMessage, err := wizard.decryptCipher(pair.PrivateKeyPem, req.FoldedMessage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  encryptedMessage,
		})
	})
}