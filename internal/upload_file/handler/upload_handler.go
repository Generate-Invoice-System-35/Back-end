package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"Back-end/internal/upload_file/adapter"
	"Back-end/internal/upload_file/model"

	"github.com/labstack/echo/v4"
)

type EchoUploadImageController struct {
	Service adapter.AdapterUploadImageService
}

// UploadImageController godoc
// @Summary      Upload Image
// @Description  User can upload image
// @Tags         File
// @accept       json
// @Produce      json
// @Router       /upload-image [post]
// @param        data  body      model.File  true  "required"
// @Success      201   {object}  model.File
// @Failure      500   {object}  model.File
func (ce *EchoUploadImageController) UploadImageController(c echo.Context) error {
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
	if _, errCopy := io.Copy(dst, src); err != errCopy {
		return errCopy
	}

	errService := ce.Service.CreateImageService(img, file)
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "upload failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "uploaded",
		"detail":  img,
	})
	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully!</p>", file.Filename))
}

// GetImagesController godoc
// @Summary      Get All Images Information
// @Description  User can get all images information
// @Tags         File
// @accept       json
// @Produce      json
// @Router       /upload-image [get]
// @Success      200   {object}  model.File
func (ce *EchoUploadImageController) GetImagesController(c echo.Context) error {
	images := ce.Service.GetAllImagesService()

	return c.JSONPretty(http.StatusOK, images, " ")
}

// GetImageController godoc
// @Summary      Get Image Information by Id
// @Description  User can get image information by id
// @Tags         File
// @accept       json
// @Produce      json
// @Router       /upload-image/{id} [get]
// @param        id    path      int         true  "id"
// @Success      200  {object}  model.File
// @Failure      404  {object}  model.File
func (ce *EchoUploadImageController) GetImageController(c echo.Context) error {
	// Get ID Param
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	image, err := ce.Service.GetImageByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "upload failed",
		})
	}

	return c.JSONPretty(http.StatusCreated, image, " ")
}

// UpdateImageController godoc
// @Summary      Update Image Information
// @Description  User can update image information
// @Tags         File
// @accept       json
// @Produce      json
// @Router       /upload-image/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.File  true  "required"
// @Success      200  {object}  model.File
// @Failure      500   {object}  model.File
func (ce *EchoUploadImageController) UpdateImageController(c echo.Context) error {
	// Get ID Param
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

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

	errService := ce.Service.UpdateImageService(intID, img, file)
	if errService != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "upload failed",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "uploaded",
		"detail":  img,
	})
}

// DeleteImageController godoc
// @Summary      Delete Image Information
// @Description  User can delete image information if they want it
// @Tags         File
// @accept       json
// @Produce      json
// @Router       /upload-image/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.File
// @Failure      500  {object}  model.File
func (ce *EchoUploadImageController) DeleteImageController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteImageService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
		"id":      intID,
	})
}
