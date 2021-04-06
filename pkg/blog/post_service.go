package blog

import "fmt"

type PostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostRes struct {
	PostId  string `json:"post_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePostService(req PostReq) error {
	err := NewPost(Post{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		l.Error("create post to mysql error", "error", err)
		return err
	}
	return nil
}

func PostListService() ([]PostRes, error) {
	ps, err := GetAllPosts()
	if err != nil {
		l.Error("get post list error", "error", err)
		return nil, err
	}

	var postRes []PostRes
	for _, p := range ps {
		postRes = append(postRes, PostRes{
			PostId:  p.PostId,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return postRes, nil
}

func GetPostByPostIdService(id string) (PostRes, error) {
	var postRes PostRes
	p, err := GetPostById(id)
	if err != nil {
		l.Error("failed to get post by post_id", "post_id", id)
		return postRes, err
	}
	return PostRes{
		PostId:  p.PostId,
		Title:   p.Title,
		Content: p.Content,
	}, nil
}

func UpdatePostService(postId string, req PostReq) error {
	if postId == "" {
		return fmt.Errorf("param empty")
	}

	var p Post
	if req.Title != "" {
		p.Title = req.Title
	}
	if req.Content != "" {
		p.Content = req.Content
	}
	tx := db.Model(&Post{}).Where("post_id = ?", postId).Updates(&p)
	if err := tx.Error; err != nil {
		l.Error("failed to update post", "post_id", postId, "error", err)
		return err
	}

	return nil
}
