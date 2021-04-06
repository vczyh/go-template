package blog

import (
	"blog-y/pkg/common/route"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	r := router.Group("/blog/v1")

	// posts
	posts := r.Group("/posts")
	posts.POST("", route.Handle(CreatePostAPI))             // 添加文章
	posts.GET("", route.Handle(PostListAPI))                // 获取文章
	posts.GET("/:postId", route.Handle(GetPostByPostIdAPI)) // 根据ID获取文章
	posts.PUT("/:postId", route.Handle(UpdatePostAPI))      // 更新文章
}
