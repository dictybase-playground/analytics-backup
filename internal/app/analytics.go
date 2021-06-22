package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dictybase-playground/analytics-backup/internal/request"
	"github.com/urfave/cli"
	ga "google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

func GetAnalyticReports(c *cli.Context) error {
	ctx := context.Background()
	s, err := ga.NewService(ctx, option.WithCredentialsFile(c.String("credentials")))
	if err != nil {
		return err
	}
	g := request.GoogleAnalyticsRequest{
		ViewID:     c.String("view-id"),
		DateStart:  c.String("start-date"),
		DateEnd:    c.String("end-date"),
		Metrics:    c.String("metrics"),
		Dimensions: c.String("dimensions"),
	}
	// convert to ReportRequest format
	req := request.MakeReportRequest(g)
	// create ReportRequest slice and fetch analytics data
	rl := make([]*ga.ReportRequest, 1)
	rl[0] = &req
	report, err := s.Reports.BatchGet(&ga.GetReportsRequest{ReportRequests: rl}).Do()
	if err != nil {
		return err
	}

	if report == nil {
		return fmt.Errorf("got nil report %v", report)
	}

	// get json version of reports
	file, err := json.Marshal(report)
	if err != nil {
		return err
	}
	// write to desired output file
	err = ioutil.WriteFile(c.String("output-file"), file, 0644)
	if err != nil {
		return err
	}

	return nil
}
