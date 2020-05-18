package database

import (
	"fmt"
	"os"
)
import (
	"github.com/elastic/go-elasticsearch/v7"
)

var es *elasticsearch.Client

func init() {
	var err error
	config := elasticsearch.Config{}
	config.Addresses = []string{"http://127.0.0.1:9200"}
	es, err = elasticsearch.NewClient(config)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
