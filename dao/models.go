package dao

import (
	"time"
)

// Book 图书表模型
type Book struct {
	ID             int        `db:"id" json:"id"`
	ISBN           string     `db:"isbn" json:"isbn"`
	Title          string     `db:"title" json:"title"`
	Author         string     `db:"author" json:"author"`
	CoverURL       string     `db:"cover_url" json:"cover_url"`
	Hash           string     `db:"hash" json:"hash"`
	PreHash        string     `db:"pre_hash" json:"pre_hash"`
	Publisher      string     `db:"publisher" json:"publisher"`
	PublishedDate  time.Time  `db:"published_date" json:"published_date"`
	Category       string     `db:"category" json:"category"`
	ContentSummary string     `db:"content_summary" json:"content_summary"`
	Rating         float32    `db:"rating" json:"rating"`
	DownloadCount  int        `db:"download_count" json:"download_count"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at"`
}

// User 用户表模型
type User struct {
	ID              int        `db:"id" json:"id"`
	AvatarURL       string     `db:"avatar_url" json:"avatar_url"`
	AvatarUpdatedAt time.Time  `db:"avatar_updated_at" json:"avatar_updated_at"`
	Username        string     `db:"username" json:"username"`
	PasswordHash    string     `db:"password_hash" json:"-"`
	Salt            string     `db:"salt" json:"-"`
	Email           string     `db:"email" json:"email"`
	Phone           string     `db:"phone" json:"phone"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}

// Download 下载记录表
type Download struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	BookID int `db:"book_id" json:"book_id"`
}

// Borrow 借阅记录表
type Borrow struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	BookID int `db:"book_id" json:"book_id"`
}

// Favorite 收藏表
type Favorite struct {
	ID        int64      `db:"id" json:"id"`
	UserID    int        `db:"user_id" json:"user_id"`
	BookID    int        `db:"book_id" json:"book_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

// Comment 评论表
type Comment struct {
	ID        int64      `db:"id" json:"id"`
	UserID    int        `db:"user_id" json:"user_id"`
	BookID    int        `db:"book_id" json:"book_id"`
	ParentID  *int64     `db:"parent_id" json:"parent_id"`
	Content   string     `db:"content" json:"content"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
