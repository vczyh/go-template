package blog

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model

	PostId  string
	Title   string
	Content string
}

func NewPost(post Post) error {
	post.PostId = fmt.Sprintf("%d", time.Now().Unix())
	return db.Create(&post).Error
}

func GetAllPosts() ([]Post, error) {
	var ps []Post
	return ps, db.Find(&ps).Error
}

func GetPostById(id string) (Post, error) {
	var p Post
	return p, db.Where("post_id = ?", id).First(&p).Error
}
