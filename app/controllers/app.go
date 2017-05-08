package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var q [3]int = [3]int{1, 2, 3}
	return c.RenderJSON(q)
}

/*func (c App) Show() revel.Result {
	parm := c.Params.Get("id")
	return c.RenderJSON(parm)
}*/

func (c App) Show() revel.Result {
	var  id  int
	c.Params.Bind(&id,"id")
	//parm := c.Params.Get("id")
	return c.RenderJSON(id)
}
