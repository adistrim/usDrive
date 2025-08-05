package models

type UploadRequest struct {
	FileName  string  `json:"fileName" binding:"required"`
	MimeType  string  `json:"mimeType" binding:"required"`
	SizeBytes int64   `json:"sizeBytes" binding:"required"`
	ParentID  *int64  `json:"parentId"`
}

type UploadResponse struct {
	FileID    int64  `json:"fileId"`
	UploadURL string `json:"uploadUrl"`
	ExpiresIn int    `json:"expiresIn"`
}

