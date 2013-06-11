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
	greeting := (now.Format("Monday") + " " + now.Format(time.StampMilli))

	fmt.Println("-----------------------------------")

	con, err := sql.Open("mysql", "test:@/test")
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := con.Query("select id, owner, donetext, donedate, createdate, deleted from done " +
		"where deleted = false order by createdate")
	dones := make([]*models.Done, 0, 10)
	var id, owner, donedate, createdate int
	var donetext string
	var deleted bool

	for rows.Next() {
		err = rows.Scan(&id, &owner, &donetext, &donedate, &createdate, &deleted)
		dones = append(dones, &models.Done{id, owner, donetext, donedate, createdate, deleted})
		log.Printf("read: id=%d text=%s\n", dones[len(dones)-1].Id, dones[len(dones)-1].Donetext)
	}

	insertSQL := "insert into done (owner, donetext, donedate, createdate, deleted) values (?, ?, ?, ?, false)"
	_, err = con.Exec(insertSQL, 1, greeting, time.Now().Unix(), time.Now().Unix())

	defer con.Close()

	expired := (time.Now().Nanosecond() - now.Nanosecond())
	log.Printf("microseconds expired: %d (which is about %dms)\n", expired/1000, expired/1000000)

	return c.Render(greeting)
}
