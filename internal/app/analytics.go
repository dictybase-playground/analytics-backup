package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/urfave/cli"
	ga "google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

type GoogleAnalyticsRequest struct {
	// Service    *ga.Service
	ViewID     string
	DateStart  string
	DateEnd    string
	Metrics    string
	Dimensions string
	PageLimit  int
	FetchAll   bool
}

func GetAnalyticReports(c *cli.Context) error {
	ctx := context.Background()
	s, err := ga.NewService(ctx, option.WithCredentialsFile(c.String("credentials")))
	if err != nil {
		return err
	}
	g := &GoogleAnalyticsRequest{
		ViewID:     c.String("view-id"),
		DateStart:  c.String("start-date"),
		DateEnd:    c.String("end-date"),
		Metrics:    c.String("metrics"),
		Dimensions: c.String("dimensions"),
		PageLimit:  10000,
		FetchAll:   true,
	}

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
	}

	rl := make([]*ga.ReportRequest, 1)
	rl[0] = &req
	report, err := s.Reports.BatchGet(&ga.GetReportsRequest{ReportRequests: rl}).Do()
	if err != nil {
		return err
	}

	if report == nil {
		return fmt.Errorf("got nil report %v", report)
	}

	file, _ := json.Marshal(report)

	_ = ioutil.WriteFile("output.json", file, 0644)

	log.Println("\n## fetched ", string(file))

	return nil
}
