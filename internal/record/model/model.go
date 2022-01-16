package recordHandlerType

import (
	"errors"
	"time"
)

type GetRecordsRequest struct {
	StartDate string `json:"startDate" `
	EndDate   string `json:"endDate" `
	MinCount  int    `json:"minCount" `
	MaxCount  int    `json:"maxCount" `
}

type GetRecordsResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"records"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type GetRecordsDBModel struct {
	StartDate time.Time
	EndDate   time.Time
	MinCount  int
	MaxCount  int
}

func (g *GetRecordsRequest) ToGetRecordsDBModel() (*GetRecordsDBModel, error) {

	startDate, err := time.Parse("2006-01-02", g.StartDate)
	if err != nil {
		return nil, errors.New(`Invalid time format. Example startDate:"yyyy-mm-dd"`)
	}
	endDate, err := time.Parse("2006-01-02", g.EndDate)
	if err != nil {
		return nil, errors.New(`Invalid time format. Example endDate:"yyyy-mm-dd"`)
	}

	return &GetRecordsDBModel{
		StartDate: startDate,
		EndDate:   endDate,
		MinCount:  g.MinCount,
		MaxCount:  g.MaxCount,
	}, nil
}
