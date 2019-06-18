package controllers

import (
	"github.com/revel/revel"
)

type AppController struct {
	*revel.Controller
}

func (c AppController) Index() revel.Result {
	return c.Render()
}
