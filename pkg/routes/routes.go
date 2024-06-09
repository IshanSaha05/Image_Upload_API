package routes

import (
	"github.com/IshanSaha05/File_Upload/pkg/controller"
	"github.com/gin-gonic/gin"
)

func UploadRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/image-upload", controller.UploadImageController())
}
