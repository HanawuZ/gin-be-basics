package auth

import "github.com/gin-gonic/gin"

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninReponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// TODO : implement login function
func Login(c *gin.Context) {

}
