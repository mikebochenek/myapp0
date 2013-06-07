package controllers

import "github.com/robfig/revel"
import "time"
import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	now := time.Now()
	greeting := ("Today is " + now.Format("Monday"))
	greeting += (" " + now.Format(time.StampMicro))

	con, err := sql.Open("mysql", "test:@/test")
	if err != nil {
		fmt.Print(err)
	}

	/*
		row := con.QueryRow("select createDate from Draft where id=?", 490)
			cb := nil
			err := row.Scan(&cb)
			fmt.Print(cb)
	*/

	defer con.Close()

	return c.Render(greeting)
}
