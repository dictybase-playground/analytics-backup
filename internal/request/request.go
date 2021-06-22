package request

import (
	"strings"

	ga "google.golang.org/api/analyticsreporting/v4"
)

type GoogleAnalyticsRequest struct {
	ViewID     string
	DateStart  string
	DateEnd    string
	Metrics    string
	Dimensions string
	PageLimit  int
	FetchAll   bool
}

func MakeReportRequest(g GoogleAnalyticsRequest) ga.ReportRequest {
	dr := make([]*ga.DateRange, 1)
	dr[0] = &ga.DateRange{StartDate: g.DateStart, EndDate: g.DateEnd}

	metSplit := strings.Split(g.Metrics, ",")
	metrics := make([]*ga.Metric, len(metSplit))
	for i, m := range metSplit {
		metrics[i] = &ga.Metric{Expression: m}
	}

	dimSplit := strings.Split(g.Dimensions, ",")
	dimensions := make([]*ga.Dimension, len(dimSplit))
	for i, d := range dimSplit {
		dimensions[i] = &ga.Dimension{Name: d}
	}

	req := ga.ReportRequest{
		DateRanges: dr,
		Metrics:    metrics,
		Dimensions: dimensions,
		ViewId:     g.ViewID,
		PageSize:   100000, // maximum allowed
	}

	return req
}
