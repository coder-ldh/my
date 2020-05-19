package database

import (
	"fmt"
	es "github.com/olivere/elastic/v7"
	"os"
)

var Es *es.Client

func init() {
	var err error
	addr := []string{"http://127.0.0.1:9200"}
	cli := es.SetURL(addr...)
	Es, err = es.NewClient(es.SetSniff(false), cli)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
