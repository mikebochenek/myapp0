package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/revel"
	"time"
)

func (c App) Index2() revel.Result {
	now := time.Now()
	greeting := ("Today is " + now.Format("Monday"))
	greeting += (" " + now.Format(time.StampMicro))

	con, err := sql.Open("mysql", "test:@/test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(greeting)

	/*
		row := con.QueryRow("select id from Draft where id=?", 490)
		var cb int
		err = row.Scan(&cb)
		fmt.Printf("output was: %d\n", cb)
	*/

	rows, err := con.Query("select id, text from Draft")
	var cb int
	var text string
	for rows.Next() {
		err = rows.Scan(&cb, &text)
		fmt.Printf("draft read: id=%d len=%d\n", cb, len(text))
	}

	var count int
	rows, err = con.Query("select id, sentence from Recommendation")
	for rows.Next() {
		count++
		err = rows.Scan(&cb, &text)
		//fmt.Printf("recommendation read: id=%d len=%d\n", cb, len(text))
	}
	fmt.Printf("count was %d.\n", count)

	defer con.Close()

	expired := (time.Now().Nanosecond() - now.Nanosecond())
	fmt.Printf("microseconds expired: %d (which is about %dms)\n", expired/1000, expired/1000000)

	return c.Render(greeting)
}
