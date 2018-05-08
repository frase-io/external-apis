package wordpress

import (
	"github.com/sunfjun/go-xmlrpc"
)

type BlogAccount struct {
	Url      string
	UserName string
	PassWord string
	BlogId   string
	PostId	 string
}

type Category struct {
	Id   int
	Name string
}

func GetCategories(ba *BlogAccount, options map[string]interface{}) (categories []*Category) {
	response := xmlrpc.Request(ba.Url, "wp.getCategories", ba.BlogId, ba.UserName, ba.PassWord, options)
	
	for _, params := range response {
		if params == nil {
			return categories
		}
		for _, param := range params.([]interface{}) {
			cat := &Category{}
			categoryId := param.(map[string]interface{})["categoryId"].(int)
			cat.Id = categoryId
			cat.Name = param.(map[string]interface{})["categoryName"].(string)
			categories = append(categories, cat)
		}
	}
	return categories
}

type Author struct {
	Id   string
	Name string
}

func GetAuthors(ba *BlogAccount, options map[string]interface{}) (authors []*Author) {
	response := xmlrpc.Request(ba.Url, "wp.getAuthors", ba.BlogId, ba.UserName, ba.PassWord, options)
	if response == nil {
		return authors
	}
	for _, params := range response {
		if params == nil {
			return authors
		}
		for _, param := range params.([]interface{}) {
			author := &Author{}
			author.Id = param.(map[string]interface{})["user_id"].(string)
			author.Name = param.(map[string]interface{})["display_name"].(string)
			authors = append(authors, author)
		}
	}
	return authors
}

type Tag struct {
	Id   int
	Name string
}

func GetTags(ba *BlogAccount, options map[string]interface{}) (tags []*Tag) {
	response := xmlrpc.Request(ba.Url, "wp.getTags", ba.BlogId, ba.UserName, ba.PassWord, options)
	for _, params := range response {
		if params == nil {
			return tags
		}
		for _, param := range params.([]interface{}) {
			tag := &Tag{}
			tag.Id = param.(map[string]interface{})["tag_id"].(int)
			tag.Name = param.(map[string]interface{})["name"].(string)
			tags = append(tags, tag)
		}
	}
	return tags
}

func NewPost(ba *BlogAccount, options map[string]interface{}) (id string) {
	response := xmlrpc.Request(ba.Url, "wp.newPost", ba.BlogId, ba.UserName, ba.PassWord, options)
	for _, params := range response {
		if params == nil {
			return id
		}
		id = params.(string)
	}
	return id
}

func EditPost(ba *BlogAccount, options map[string]interface{}) (isPostModified bool) {	
	response := xmlrpc.Request(ba.Url, "wp.editPost", ba.BlogId, ba.UserName, ba.PassWord, ba.PostId, options)
	
	for _, params := range response {
		isPostModified = params.(bool)
	}
	return isPostModified
}

func DeletePost(ba *BlogAccount, options map[string]interface{}) (isPostDeleted bool) {	
	response := xmlrpc.Request(ba.Url, "wp.deletePost", ba.BlogId, ba.UserName, ba.PassWord, ba.PostId, options)
	
	for _, params := range response {
		isPostDeleted = params.(bool)
	}
	return isPostDeleted
}

type File struct {
	FileName string
	Url      string
}

func UploadFile(ba *BlogAccount, options map[string]interface{}) (f *File) {
	response := xmlrpc.Request(ba.Url, "wp.uploadFile", ba.BlogId, ba.UserName, ba.PassWord, options)
	for _, params := range response {
		if params == nil {
			return f
		}
	}
	return f
}
