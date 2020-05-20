package database

import (
	"context"
	"fmt"
	es "github.com/olivere/elastic/v7"
	"os"
)

var Es *es.Client
var BulkService *es.BulkProcessor

func init() {
	var err error
	addr := []string{"http://127.0.0.1:9200"}
	cli := es.SetURL(addr...)
	Es, err = es.NewClient(es.SetSniff(false), cli)
	checkError(err)
	BulkService, err = Es.BulkProcessor().Name("EsBackgroundWorker-1").Workers(2).BulkSize(100).Do(context.Background())
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
