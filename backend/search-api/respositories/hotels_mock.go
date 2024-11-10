package search

/*
import (
	"context"
	"search-api/dao"
)

type Mock struct {
	data map[int64]hotels.Hotel
}

func NewMock() Mock {
	return Mock{
		data: make(map[int64]hotels.Hotel),
	}
}

func (repo Mock) Search(ctx context.Context, query string, offset int, limit int) ([]hotels.Hotel, error) {
	result := make([]hotels.Hotel, 0)
	for i, hotel := range repo.data {
		if int(i) < offset {
			continue
		}
		if len(result) == limit {
			break
		}
		result = append(result, hotel)
	}
	return result, nil
}
*/
