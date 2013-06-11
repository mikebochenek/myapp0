package models

type Done struct {
	Id         int
	Owner      int
	Donetext   string
	Donedate   int
	Createdate int
	Deleted    bool
}
