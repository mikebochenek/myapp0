package controllers

import "github.com/robfig/revel"
import "time"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	now := time.Now()
	greeting := ("Today is " + now.Format("Monday"))
	greeting += (" " + now.Format(time.StampMicro))

	return c.Render(greeting)
}
