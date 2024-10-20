package search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stevenferrer/solr-go"
	hotels "search-api/dao"
)

type SolrConfig struct {
	BaseURL    string
	Collection string
}

type Solr struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolr(config SolrConfig) Solr {
	return Solr{
		Client:     solr.NewJSONClient(config.BaseURL),
		Collection: config.Collection,
	}
}

func (repo Solr) Search(ctx context.Context, query string, offset int, limit int) ([]hotels.Hotel, error) {
	queryParser := solr.NewChildrenQueryParser().Query(query).BuildParser()
	solrQuery := solr.NewQuery(queryParser).Sort("rating").Offset(offset).Limit(limit)
	response, err := repo.Client.Query(ctx, repo.Collection, solrQuery)
	if err != nil {
		return nil, fmt.Errorf("error running query against Solr: %w", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("error running query against Solr: %w", response.Error)
	}
	bytes, err := json.Marshal(response.Response.Documents)
	if err != nil {
		return nil, fmt.Errorf("error matching Solr results: %w", err)
	}
	result := make([]hotels.Hotel, 0)
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling Solr results: %w", err)
	}
	return result, nil
}
