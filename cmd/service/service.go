package main

import (
	"fmt"
	"os"
	"sftasks/v2/pkg/storage"
	"sftasks/v2/pkg/storage/postgres"
)

var db storage.Interface

func main() {

	pss := os.Getenv("pss")

	connstr := "postgres://postgres:" + pss + "@localhost/taskssf"

	db, err := postgres.New(connstr)
	if err != nil {
		fmt.Println(err)
	}

	//rowstasks, err := db.TasksByLabel(4) //TasksByAuthor(1) //db.Tasks(0)
	//fmt.Println(rowstasks)

	//rowtask, err := db.Tasks(4)
	//fmt.Println(rowtask[0])
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//rowtask[0].Title = "testupdate"

	//db.TasksUpdateByID(rowtask[0].ID, rowtask[0])

	db.TaskDelete(4)
}
