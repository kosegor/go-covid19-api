package elastic

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic"
)

func NewElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second), // quit trying after 5 seconds
	)

	// Check and see if olivere's NewClient() method returned an error
	if err != nil {
		fmt.Println("elastic.NewClient() ERROR: %v", err)
		log.Fatalf("quiting connection..")
	} else {
		// Print client information
		fmt.Println("client:", client)
		fmt.Println("client TYPE:", reflect.TypeOf(client))
	}

	return client, err
}
