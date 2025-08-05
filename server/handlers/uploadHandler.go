package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"

	"usdrive/config"
	"usdrive/db"
	"usdrive/models"
)

func RequestUpload(c *gin.Context) {
	var uploadRequest models.UploadRequest
	if err := c.ShouldBindJSON(&uploadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	pgClient := db.GetDBInstance()
	var fileID int64
	err := pgClient.QueryRow(
		context.Background(),
		`INSERT INTO files (parent_id, name, is_folder, mime_type, size_bytes, status) 
		 VALUES ($1, $2, $3, $4, $5, 'pending_upload') 
		 RETURNING id`,
		uploadRequest.ParentID, uploadRequest.FileName, false, uploadRequest.MimeType, uploadRequest.SizeBytes,
	).Scan(&fileID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file record"})
		return
	}

	storagePath := generateStoragePath(fileID, uploadRequest.FileName)

	_, err = pgClient.Exec(
		context.Background(),
		"UPDATE files SET storage_path = $1 WHERE id = $2",
		storagePath, fileID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update storage path"})
		return
	}

	r2Client := db.GetR2Client()
	presignClient := s3.NewPresignClient(r2Client)
	presignedReq, err := presignClient.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(config.ENV.R2Bucket),
		Key: aws.String(storagePath),
		ContentType: aws.String(uploadRequest.MimeType),
	}, s3.WithPresignExpires(time.Minute*15))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate upload URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileId":     fileID,
		"uploadUrl":  presignedReq.URL,
		"expiresIn":  15 * 60,
	})
}

func generateStoragePath(fileID int64, fileName string) string {
	return "files/" + strconv.FormatInt(fileID, 10) + "/" + fileName
}

func CompleteUpload(c *gin.Context) {
	fileIDStr := c.Param("fileId")
	fileID, err := strconv.ParseInt(fileIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID format"})
		return
	}

	pgClient := db.GetDBInstance()

	var exists bool
	err = pgClient.QueryRow(
		context.Background(),
		"SELECT EXISTS(SELECT 1 FROM files WHERE id = $1 AND status = 'pending_upload')",
		fileID,
	).Scan(&exists)

	if err != nil || !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found or already processed"})
		return
	}

	_, err = pgClient.Exec(
		context.Background(),
		"UPDATE files SET status = 'active', updated_at = NOW() WHERE id = $1",
		fileID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update file status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File upload completed successfully",
		"fileId": fileID,
	})
}
