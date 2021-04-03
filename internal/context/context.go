package context

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/vkbotserver/domain"
	"golang.org/x/text/language"
)

type (
	Context struct {
		Lang      language.Tag
		Container di.Container
	}
)

func WithContext(req *domain.Request, context *Context) *domain.Request {
	var (
		clonedReq = new(domain.Request)
	)

	*clonedReq = *req
	clonedReq.Context = context

	return clonedReq
}

func GetContext(req *domain.Request) *Context {
	return req.Context.(*Context)
}
