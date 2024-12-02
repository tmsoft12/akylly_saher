package routes

import (
	"awtopark/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func InitRoutes(app *fiber.App) {

	app.Get("/", controllers.GetAllParkdata)
	app.Post("/", controllers.PostParkData)
	app.Get("/dashboard", controllers.CountEmptyExitTimes)
	app.Put("/", controllers.UpdateParkData)
	app.Get("/ws/dashboard", websocket.New(controllers.DashboardWebSocket))
	app.Get("/ws/parkdata", websocket.New(controllers.SendParkDataWebSocket))
	app.Get("/ws/plate", websocket.New(controllers.GetAllParkdataWeb))
	app.Put("/rf", controllers.UpdateRF)

}
