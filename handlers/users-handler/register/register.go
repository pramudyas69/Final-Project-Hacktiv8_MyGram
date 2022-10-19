package register

import (
	"MyGram/controllers/user-controllers/register"
	"MyGram/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.InputRegister
	ctx.ShouldBindJSON(&input)

	resultRegister, errRegister := h.service.RegisterService(&input)

	switch errRegister {
	case "REGISTER_CONFLICT_409":
		utils.APIResponse(ctx, "Email already exist!", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		utils.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		_, errToken := utils.Sign(accessTokenData, utils.GoDotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer log.Fatal(errToken.Error())
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		}

		resultRegister.Password = ""
		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, resultRegister)
	}
}
