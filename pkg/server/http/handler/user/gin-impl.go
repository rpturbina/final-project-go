package user

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/final-project-go/helpers"
	"github.com/rpturbina/final-project-go/pkg/domain/message"
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
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	currentAge := helpers.ConvertDOBToCurrentAge(time.Time(user.DOB), time.Now())

	log.Println("checking user age")
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
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    01,
		"message": "user has successfully registered",
		"type":    "ACCEPTED",
		"data": gin.H{
			"age":      currentAge,
			"email":    result.Email,
			"id":       result.ID,
			"username": result.Username,
		},
	})
}

func (u *UserHdlImpl) GetUserByIdHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByIdHdl is invoked\n", u)
	defer log.Printf("%T - GetUserByIdHdl executed\n", u)

	log.Println("check user_id from path parameter")
	userIdParam := ctx.Param("user_id")

	userId, err := strconv.ParseUint(userIdParam, 0, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 96,
			"type": "BAD_REQUEST",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_FORMAT",
				"error_message": err.Error(),
			},
		})
		return
	}

	log.Println("calling get user by id service usecase")
	result, errMsg := u.userUsecase.GetUserByIdSvc(ctx, userId)
	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "user is found",
		"type":    "SUCCESS",
		"data": gin.H{
			"id":            result.ID,
			"username":      result.Username,
			"social_medias": result.SocialMedias,
		},
	})
}

func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	log.Printf("%T - UpdateUserHdl is invoked\n", u)
	defer log.Printf("%T - UpdateUserHdl executed\n", u)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	var updatedUser user.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed to bind payload",
			"invalid_arg": gin.H{
				"error_type":    "INVALID_PAYLOAD",
				"error_message": err.Error(),
			},
		})
		return
	}

	idToken, errMsg := u.userUsecase.UpdateUserSvc(ctx, userId, updatedUser.Email, updatedUser.Username)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "user has been successfully updated",
		"type":    "SUCCESS",
		"data": gin.H{
			"id_token": idToken,
		},
	})
}

func (u *UserHdlImpl) DeleteUserHdl(ctx *gin.Context) {
	log.Printf("%T - DeleteUserHdl is invoked\n", u)
	defer log.Printf("%T - DeleteUserHdl executed\n", u)

	stringUserId := ctx.Value("user").(string)

	userId, _ := strconv.ParseUint(stringUserId, 0, 64)

	errMsg := u.userUsecase.DeleteUserSvc(ctx, userId)

	if errMsg.Error != nil {
		message.ErrorResponseSwitcher(ctx, errMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"message": "Your account has been successfully deleted",
		"type":    "SUCCESS",
	})
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}
