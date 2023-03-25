package controller

import (
	"TopTips/dto"
	"TopTips/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	AuthService service.AuthService
	JWTService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) authController {
	return authController{
		AuthService: authService,
		JWTService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RequestRegisterDTO

	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}
	ctx.JSON(http.StatusAccepted, "sms code sent")
}

func (c *authController) Login(ctx *gin.Context) {
	var login dto.LoginDTO

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	ctx.JSON(200, gin.H{"message": "wait"})
}

func (c *authController) ConfirmSMS(ctx *gin.Context) {
	var confirm dto.CheckCodeRequest
	var responseDTO dto.ResponseDTO
	if err := ctx.ShouldBindJSON(&confirm); err != nil {
		ctx.JSON(
			http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("bad request: %v\n", err),
			},
		)
		// exit process
		return
	}

	if confirm.Code != "1111" {
		ctx.JSON(http.StatusConflict, gin.H{"message": "wrong sms code."})
		return
	}
	token_, err := c.JWTService.GenerateToken(confirm.PhoneNumber, "")
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "can not generate token"})
		return
	}
	responseDTO.Token = token_
	responseDTO.Profile = dto.Profile{
		FirtsName: "",
		LastName:  "",
		Phone:     "",
		Status:    "FILL_DATA",
		Password:  "",
		Role:      "",
		UserId:    0,
		Token:     token_,
	}
	ctx.JSON(201, responseDTO)
}

func (c *authController) Forgot(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "wait"})
}
