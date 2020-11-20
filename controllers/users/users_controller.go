package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moz5691/bookstore_users-api/domain/users"
	"github.com/moz5691/bookstore_users-api/services"
	"github.com/moz5691/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	//fmt.Println("get here", user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(&user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "fix me")
//}
