package api

import (
	"GIG/app/controllers"
	"GIG/app/repositories"
	"GIG/app/services"
	"github.com/revel/revel"
)

type GraphController struct {
	*revel.Controller
}

func (c GraphController) GetGraph() revel.Result {

	graph, err := repositories.EntityRepository{}.GetGraph()

	if err != nil {
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	return c.RenderJSON(controllers.BuildSuccessResponse(services.GetGraph(graph), 200))
}
