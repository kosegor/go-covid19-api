package elastic

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/kosegor/go-covid19-api/app/domain/model"
	"github.com/kosegor/go-covid19-api/app/interface/apierr"
	"github.com/olivere/elastic"
)

const (
	index = "incidents_index"
)

type elasticRepository struct {
	client *elastic.Client
}

func NewElasticRepository() *elasticRepository {
	client, err := NewElasticClient()
	if err == nil {
		return &elasticRepository{
			client: client,
		}
	}
	return nil
}

func (e *elasticRepository) Insert(incident *model.Incident) (*model.Incident, *apierr.ApiError) {
	// Declare a timeout context for the API calls
	ctx, stop := context.WithTimeout(context.Background(), 3*time.Second)
	defer stop()

	// Check if the Elasticsearch index already exists
	exist, err := e.client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v", err)
	} else if exist {
		id, err := e.client.Index().
			Index(index).
			BodyJson(incident).
			Do(ctx)

		if err != nil {
			log.Fatalf("client.Index() ERROR: %v", err)
		} else {
			fmt.Println("\nIndex: ", id.Id)
			fmt.Println("\nElasticsearch document indexed:", incident)
			fmt.Println("doc object TYPE:", reflect.TypeOf(incident))
			incident.ID = id.Id
			_, err = e.client.Flush(index).WaitIfOngoing(true).Do(ctx)
		}
	} else {
		fmt.Println("client.Index() ERROR: the index", index, "does not exist")
	}

	if err != nil {
		return nil, apierr.NewApiError(err.Error(), http.StatusInternalServerError)
	}
	return incident, nil
}

/*
func FindByCountry(country string) ([]*model.Incident, *apierr.ApiError) {

}
*/
