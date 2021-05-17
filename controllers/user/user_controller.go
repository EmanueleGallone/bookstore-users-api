package user

import (
	"bookstore-users-api/domain/users"
	"bookstore-users-api/services"
	"bookstore-users-api/util/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

func getUser(URL string) (int64, *errors.RestError) {
	userID, userErr := strconv.ParseInt(URL, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("User ID should be a number")
	}
	return userID, nil
}

func GetUser(ctx *gin.Context) {
	userID, userErr := getUser(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	user, err := services.UserServices.GetUser(userID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, user.Marshal(ctx.GetHeader("X-Public") == "true"))
}

func UpdateUser(ctx *gin.Context) {
	var user users.User

	userID, userErr := getUser(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	err := ctx.MustBindWith(&user, binding.JSON)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid JSON")
		ctx.JSON(restError.Status, &restError)
		return
	}

	//From here on, binding completed; 'user' variable now has all the data.
	user.Id = userID
	isPartial := ctx.Request.Method == http.MethodPatch //Checking if partial update is requested

	result, resultErr := services.UserServices.UpdateUser(&isPartial, &user)
	if resultErr != nil {
		ctx.JSON(resultErr.Status, resultErr)
		return
	}

	ctx.JSON(http.StatusOK, result.Marshal(ctx.GetHeader("X-Public") == "true"))
}

func CreateUser(ctx *gin.Context) {
	var user users.User

	err := ctx.MustBindWith(&user, binding.JSON)
	if err != nil {
		restError := errors.NewBadRequestError("Invalid JSON")
		ctx.JSON(restError.Status, &restError)
		return
	}

	result, saveErr := services.UserServices.CreateUser(&user) // passing the user to the service
	if saveErr != nil {
		ctx.JSON(saveErr.Status, &saveErr)
		return
	}

	ctx.JSON(http.StatusCreated, result.Marshal(ctx.GetHeader("X-Public") == "true"))
}

func DeleteUser(ctx *gin.Context) {
	var user users.User

	userID, userErr := getUser(ctx.Param("user_id"))
	if userErr != nil {
		ctx.JSON(userErr.Status, userErr)
		return
	}

	user.Id = userID
	result, delErr := services.UserServices.DeleteUser(&user)
	if delErr != nil {
		ctx.JSON(delErr.Status, delErr)
		return
	}

	ctx.JSON(http.StatusOK, result.Marshal(ctx.GetHeader("X-Public") == "true"))
}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")

	users, err := services.UserServices.SearchUsers(status)
	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, users.Marshal(ctx.GetHeader("X-Public") == "true"))
}
