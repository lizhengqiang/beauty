package beauty

import "gopkg.in/macaron.v1"

type Handler interface {
}

type Service struct {
	Ctx *macaron.Context
}

func (this *Service) Handler(pattern string, h ...macaron.Handler) {
	this.Ctx.Get(pattern, h)

}
