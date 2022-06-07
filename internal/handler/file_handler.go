package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoUploadController struct {
	Service adapter.AdapterUploadService
}

func (ce *EchoUploadController) UploadImageController(c echo.Context) error {
	// Read form fileds
	img := model.File{}
	c.Bind(&img)

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join("storage/", filepath.Base(file.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, errCopy := io.Copy(dst, src); err != nil {
		return errCopy
	}

	errService := ce.Service.UploadImageService(img, file)
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "upload failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "uploaded",
		"detail":   img,
	})
	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully!</p>", file.Filename))
}
