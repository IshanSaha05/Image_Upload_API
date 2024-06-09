package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/IshanSaha05/File_Upload/pkg/config"
	"github.com/gin-gonic/gin"
)

func UploadImageController() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := os.Stat(config.UploadLocation); os.IsNotExist(err) {
			err := os.Mkdir(config.UploadLocation, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, config.MaxUploadSize)

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file received."})
			return
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !config.AllowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File type not allowed."})
			return
		}

		filename := filepath.Base(file.Filename)
		filepath := filepath.Join(config.UploadLocation, filename)

		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file."})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "File Uploaded Successfully",
			"filepath": filepath})
	}
}
