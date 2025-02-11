package dbmodel

import "time"

// type Post struct {
// 	Title   string `gorm:"not null"`
// 	Content string `gorm:"not null"`
// 	Author  string `gorm:"not null;uniqueIndex"`
// }

type Comment struct {
	ID        int32      `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID    int32      `json:"PostId" gorm:"not null;index"`
	Author    string     `json:"Author" gorm:"not null"`
	Content   string     `json:"Content" gorm:"not null"`
	ParentID  *int32     `json:"ParentId" gorm:"index"`
	CreatedAt time.Time  `json:"Created_At" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"Updated_At" gorm:"type:timestamp;default:null"`
}

type Post struct {
	ID              int32      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string     `json:"Title" gorm:"not null"`
	Content         string     `json:"Content" gorm:"not null"`
	Author          string     `json:"Author" gorm:"not null"`
	PublishedAt     time.Time  `json:"Published_At" gorm:"type:timestamp;not null"`
	UpdatedAt       *time.Time `json:"Updated_At" gorm:"type:timestamp;default:null"`
	Comments        []Comment  `json:"Comments" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
	CommentsAllowed bool       `json:"Comments_Allowed" gorm:"not null;default:true"`
}
