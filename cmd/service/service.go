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
		//return nil, err
	}

	rowstasks, err := db.Tasks(0, 0, 0)
	fmt.Println(rowstasks)
	//rows, err := db.Query(context.Background(), `
	//	SELECT * FROM labels;`)

	//for rowstasks.Next() {
	//	var st int
	//	var st2 string
	//	err = rowstasks.Scan(
	//		&st,
	//		&st2,
	//	)

	//	fmt.Println(st, st2)

	//}
	//defer db.Close()
}
