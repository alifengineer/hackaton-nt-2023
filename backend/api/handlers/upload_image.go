package handlers

import (
	"context"
	"crypto/tls"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/dilmurodov/xakaton_nt/api/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	nethttp "net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type UploadResponse struct {
	Filename string `json:"filename"`
}

type File struct {
	File       *multipart.FileHeader `form:"file" binding:"required"`
	MerchantId string                `form:"merchant_id"`
}

type Path struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

// @Security ApiKeyAuth
// @Router /api/v1/image/upload [post]
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} Path
func (h *Handler) UploadImage(c *gin.Context) {
	var (
		file File
	)

	err := c.ShouldBind(&file)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	fName := uuid.NewString()
	file.File.Filename = strings.ReplaceAll(file.File.Filename, " ", "_")
	file.File.Filename = fName + "_" + file.File.Filename

	dst, err := os.Getwd()
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	minioClient, err := minio.New(h.cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccessKey, ""),
		Secure: true,
		Transport: &nethttp.Transport{
			DisableCompression: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}
	err = c.SaveUploadedFile(file.File, dst+"/"+file.File.Filename)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	_, err = minioClient.FPutObject(
		context.Background(),
		"docs",
		file.File.Filename,
		dst+"/"+file.File.Filename,
		minio.PutObjectOptions{ContentType: ""},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		os.Remove(dst + "/" + file.File.Filename)
		return
	}

	os.Remove(dst + "/" + file.File.Filename)
	url := fmt.Sprintf("https://%v/docs/%v", h.cfg.MinioEndpoint, file.File.Filename)
	h.handleResponse(c, http.OK, Path{
		Url:      url,
		Filename: file.File.Filename,
	})
}
