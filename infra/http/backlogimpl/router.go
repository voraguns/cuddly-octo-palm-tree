package backlogimpl

import bi "cuddly-octo-palm-tree/service/backlogimpl"

type Router struct {
	biService bi.BiService
}

func NewRouter(svc bi.BiService) *Router {
	return &Router{
		biService: svc,
	}
}
