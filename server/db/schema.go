package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type FileStatus string

const (
	StatusPendingUpload FileStatus = "pending_upload"
	StatusActive        FileStatus = "active"
	StatusError         FileStatus = "error"
)

func (fs *FileStatus) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to unmarshal FileStatus value: %v", value)
	}
	*fs = FileStatus(str)
	return nil
}

func (fs FileStatus) Value() (driver.Value, error) {
	return string(fs), nil
}

func (FileStatus) GormDataType() string {
	return "file_status"
}


type User struct {
	ID         uint      `gorm:"primaryKey"`
	GoogleID   string    `gorm:"unique;not null"`
	Email      string    `gorm:"unique;not null"`
	FullName   string
	AvatarURL  string
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

type File struct {
	ID          uint64     `gorm:"primaryKey"`
	OwnerID     uint       `gorm:"not null;index"`
	ParentID    *uint64    `gorm:"index:idx_files_parent_id"`
	Name        string     `gorm:"not null"`
	IsFolder    bool       `gorm:"not null;default:false"`
	MimeType    string
	SizeBytes   int64
	Status      FileStatus `gorm:"type:file_status;not null;default:pending_upload"`
	StoragePath string
	CreatedAt   time.Time  `gorm:"not null;default:now()"`
	UpdatedAt   time.Time  `gorm:"not null;default:now()"`
	
	Parent      *File      `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:CASCADE"`
	Owner       *User      `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE"`
}
