package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

// 此处定义了一个 Swagger的对象，专门用于 Swagger 接口文档的展示，为了下面这个注释可以正常使用
// @Success 200 {object} model.TagSwagger "成功"
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}
