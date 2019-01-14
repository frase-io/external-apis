package wordpress

import (
	"github.com/carthics/go-xmlrpc"
	"fmt"
	"reflect"
	"strconv"
	"time"
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
			v := param.(map[string]interface{})["categoryId"]
			var categoryId int
			if reflect.ValueOf(v).Kind() == reflect.String{
				abc := param.(map[string]interface{})["categoryId"].(string)
				categoryId,_ = strconv.Atoi(abc)
			}else{
				categoryId = param.(map[string]interface{})["categoryId"].(int)
			}			
			cat.Id = categoryId
			cat.Name = param.(map[string]interface{})["categoryName"].(string)
			categories = append(categories, cat)
		}
	}
	return categories
}

type Author struct {
	Id   	  string
	Name 	  string
	UserLogin string
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
			author.UserLogin = param.(map[string]interface{})["user_login"].(string)
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

func UploadFile(ba *BlogAccount, options map[string]interface{}) (id string) {
	response := xmlrpc.Request(ba.Url, "wp.uploadFile", ba.BlogId, ba.UserName, ba.PassWord, options["data"])
	
	for _, params := range response {
		if params == nil {
			return id
		}

		fmt.Println(params)
		id = params.(map[string]interface{})["id"].(string)
		
	}
	return id
}

type Post struct {
	PostTitle 	string 		`json: post_title, omit_empty`
	PostDate 	time.Time 	`json: post_title, omit_empty`
	PostStatus  string		`json: post_status, omit_empty`
	PostAuthor  string		`json: post_author, omit_empty`
}

func GetPosts(ba *BlogAccount, options map[string]interface{}) (listOfPosts []Post){
	response := xmlrpc.Request(ba.Url, "wp.getPosts", ba.BlogId, ba.UserName, ba.PassWord, options)
	
	for _, params := range response {
		
		if params == nil {
			return listOfPosts
		}

		for _, param := range params.([]interface{}) {
			eachPost := Post{}
			eachPost.PostTitle = param.(map[string]interface{})["post_title"].(string)
			eachPost.PostDate = param.(map[string]interface{})["post_date"].(time.Time)
			eachPost.PostStatus = param.(map[string]interface{})["post_status"].(string)
			eachPost.PostAuthor = param.(map[string]interface{})["post_author"].(string)
			listOfPosts = append(listOfPosts, eachPost)
		}
	}

	return listOfPosts
}
