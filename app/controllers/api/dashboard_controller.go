package api

import (
	"GIG/app/controllers"
	"GIG/app/services"
	"github.com/revel/revel"
	"log"
)

type DashboardController struct {
	*revel.Controller
}

func (c DashboardController) GetStats() revel.Result {

	entityStats, err := services.GetGraphStats(false)
	if err != nil {
		log.Println("error reading db stats:", err)
		return c.RenderJSON(controllers.BuildErrorResponse(err, 500))
	}

	return c.RenderJSON(controllers.BuildSuccessResponse(entityStats, 200))
}
