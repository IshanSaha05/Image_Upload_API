package main

import (
	"github.com/IshanSaha05/File_Upload/pkg/config"
	"github.com/IshanSaha05/File_Upload/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UploadRoute(router)

	router.Run(":" + config.Port)
}
