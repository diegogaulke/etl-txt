package es

import (
	"context"

	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
)

// Add object to the index
func Add(index string, data interface{}) error {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL("https://elastic:3umksqI0sSUWGFYpnMz3gvb4@27f4df016cb74b68990e96f17a0980b4.sa-east-1.aws.found.io:9243"), elastic.SetSniff(false))
	if err != nil {
		return err
	}

	uuid := uuid.NewV4()
	_, err = client.Index().Index(index).Id(uuid.String()).BodyJson(data).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}
