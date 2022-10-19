package user

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/helpers"
	"github.com/rpturbina/final-project-go/pkg/domain/user"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func (u *UserHdlImpl) RegisterUserHdl(ctx *gin.Context) {
	log.Printf("%T - RegisterUserHdl is invoked\n", u)
	defer log.Printf("%T - RegisterUserHdl executed\n", u)

	log.Println("binding body payload from request")

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_FORMAT",
				"error_message": err.Error(),
			},
		})
		return
	}

	// TODO: check where to place the check input logic
	currentAge := helpers.ConvertDOBToCurrentAge(time.Time(user.DOB), time.Now())

	if currentAge < 8 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"message": "Your minimal age should be 8 years old",
			"type":    "BAD_REQUEST",
		})
		return
	}

	log.Println("calling register user service usecase")
	result, errMsg := u.userUsecase.RegisterUserSvc(ctx, user)

	if errMsg.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 96,
			"type": "BAD_REQUEST",
			"invalid_arg": gin.H{
				"error_type":    errMsg.Type,
				"error_message": errMsg.Error.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "user successfully registered",
		"type":    "ACCEPTED",
		"data": gin.H{

			"age":      currentAge,
			"email":    result.Email,
			"id":       result.ID,
			"username": result.UserName,
		},
	})
}

func (u *UserHdlImpl) GetUserByIdHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByIdHdl is invoked\n", u)
	defer log.Printf("%T - GetUserByIdHdl executed\n", u)

	// get query params from url
	userIdParam := ctx.Param("user_id")

	// check user email from query params, if empty -> BAD_REQUEST
	log.Println("check user email from quary params")

	userId, err := strconv.ParseUint(userIdParam, 0, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 96,
			"type": "BAD_REQUEST",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_USER_ID_FORMAT",
				"error_message": err.Error(),
			},
		})
		return
	}

	// calling service/usecase for get user data by id
	log.Println("calling get user by id service usecase")
	result, errMsg := u.userUsecase.GetUserByIdSvc(ctx, userId)
	if errMsg.Error != nil {
		switch errMsg.Type {
		case "INTERNAL_CONNECTION_PROBLEM":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": 99,
				"type": "INTERNAL_SERVER_ERROR",
				"invalid_arg": gin.H{
					"error_type":    errMsg.Type,
					"error_message": errMsg.Error.Error(),
				},
			})
		case "USER_NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 99,
				"type": "BAD_REQUEST",
				"invalid_arg": gin.H{
					"error_type":    errMsg.Type,
					"error_message": errMsg.Error.Error(),
				},
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "user successfully registered",
		"type":    "SUCCESS",
		"data": gin.H{
			"id":            result.ID,
			"username":      result.UserName,
			"social_medias": result.SocialMedias,
		},
	})

}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}
