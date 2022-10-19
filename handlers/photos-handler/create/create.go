package handlerCreate

import (
	"MyGram/controllers/photos-controllers/create"
	"MyGram/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service create.Service
}

func NewHandlerCreatePhotos(service create.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateStudentHandler(ctx *gin.Context) {

	var input create.InputCreatePhoto
	ctx.ShouldBindJSON(&input)

	resultCreatePhoto, errCreatePhoto := h.service.CreatePhotoService(&input)

	switch errCreatePhoto {

	case "CREATE_STUDENT_FAILED_403":
		utils.APIResponse(ctx, "Create new photo failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		utils.APIResponse(ctx, "Create new photo successfully", http.StatusCreated, http.MethodPost, resultCreatePhoto)
	}
}
