package es

import (
	"context"
	"errors"
	"os"

	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
)

// Add object to the index
func Add(index string, data interface{}) error {
	ctx := context.Background()

	url := os.Getenv("ELASTIC_URL")
	if url == "" {
		return errors.New("please set your elastic search url")
	}

	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
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
