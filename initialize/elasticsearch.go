package initialize

import (
	"context"
	"fmt"
	es "github.com/olivere/elastic/v7"
	"my/global"
	"os"
)

func init() {
	var err error
	addr := []string{"http://127.0.0.1:9200"}
	cli := es.SetURL(addr...)
	global.GVA_ES, err = es.NewClient(es.SetSniff(false), cli)
	checkError(err)
	global.GVA_BulkProcessor, err = global.GVA_ES.BulkProcessor().Name("EsBackgroundWorker-1").Workers(2).BulkSize(100).Do(context.Background())
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
