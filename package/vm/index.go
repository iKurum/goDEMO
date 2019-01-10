package vm

import (
	"zafkiel/package/model"
)

//IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
	BasePageViewModel
}

//IndexViewModel10p struct
type IndexViewModel10p struct{}

//GetVM func
func (IndexViewModel10p) GetVM(username, flash string, page, limit int) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts, total, _ := u.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("Homepage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

//CreatePost func
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
