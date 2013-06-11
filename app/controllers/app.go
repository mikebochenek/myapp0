package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/revel"
	"log"
	"myapp0/app/models"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	now := time.Now()
	greeting := ("Today is " + now.Format("Monday"))
	greeting += (" " + now.Format(time.StampMicro) + "\n")

	fmt.Println("-----------------------------------")

	con, err := sql.Open("mysql", "test:@/test")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(greeting)

	rows, err := con.Query("select id, owner, donetext, donedate from done where deleted = false order by createdate")
	d := models.Done{}
	for rows.Next() {
		err = rows.Scan(&d.Id, &d.Owner, &d.Donetext, &d.Donedate)
		log.Printf("done read: id=%d len=%d\n", d.Id, len(d.Donetext))
	}

	defer con.Close()

	expired := (time.Now().Nanosecond() - now.Nanosecond())
	log.Printf("microseconds expired: %d (which is about %dms)\n", expired/1000, expired/1000000)

	return c.Render(greeting)
}
