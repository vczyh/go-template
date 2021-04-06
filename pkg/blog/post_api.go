package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreatePostAPI(c *gin.Context) (interface{}, error) {
	var req PostReq
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return nil, CreatePostService(req)
}

func PostListAPI(c *gin.Context) (interface{}, error) {
	return PostListService()
}

func GetPostByPostIdAPI(c *gin.Context) (interface{}, error) {
	postId := c.Param("postId")
	return GetPostByPostIdService(postId)
}

func UpdatePostAPI(c *gin.Context) (interface{}, error) {
	postId := c.Param("postId")
	var req PostReq
	if err := c.Bind(&req); err != nil {
		return nil, fmt.Errorf("param error")
	}
	return nil, UpdatePostService(postId, req)
}
