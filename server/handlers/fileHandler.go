package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"usdrive/db"
)

type File struct {
	ID         int64      `json:"id"`
	ParentID   *int64     `json:"parentId"`
	Name       string     `json:"name"`
	IsFolder   bool       `json:"isFolder"`
	MimeType   *string    `json:"mimeType"`
	SizeBytes  *int64     `json:"sizeBytes"`
	Status     string     `json:"status"`
	StoragePath *string   `json:"storagePath"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}


func ListActiveFiles(c *gin.Context) {
	pgClient := db.GetDBInstance()
	rows, err := pgClient.Query(
		context.Background(),
		`SELECT id, parent_id, name, is_folder, mime_type, size_bytes, status, storage_path, created_at, updated_at
		 FROM files WHERE status = 'active'`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var f File
		err := rows.Scan(
			&f.ID, &f.ParentID, &f.Name, &f.IsFolder, &f.MimeType, &f.SizeBytes,
			&f.Status, &f.StoragePath, &f.CreatedAt, &f.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse file data"})
			return
		}
		files = append(files, f)
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
