package users

import (
	"net/http"

	"github.com/TestardR/bookstore_users-api/domain/users"
	"github.com/TestardR/bookstore_users-api/services"
	"github.com/TestardR/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	/*
		Above code is equivalent
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// TODO: Handle error
			return
		}

		if err := json.Unmarshal(bytes, &user); err != nil {
			// TODO: Handle json error
			return
		}
	*/

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "get-user")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "search-user")
}
