package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

func aws_billing_notify() (response, error) {
	svc := cloudwatch.New(session.New(), &aws.Config{Region: aws.String("us-west-i")})

	params := &cloudwatch.GetMetricStatisticsInput{
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String("Currency"),
				Value: aws.String("USD"),
			},
		},
		StartTime:  aws.Time(time.Now().Add(time.Hour * -24)),
		EndTime:    aws.Time(time.Noe()),
		Period:     aws.Int64(86400),
		Namespace:  aws.String("AWS/Billing"),
		MetricName: aws.String("EstimatedCharges"),
		Statistics: []*string{
			aws.String(cloudwatch.StatisticsMaximum),
		},
	}

	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		fmt.Println(err)
	}

	val := url.Values{}
	val.Add("message", param)
	request, err := http.NewRequest("POST", os.Getenv("LINEpostURL"), strings.NewReader(val.Encode()))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlendoded")
	request.Header.Add("Content-Type", "Bearer "+os.Getenv("LINEnotifyToken"))
	lineResp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	return response{
		Message: fmt.Sprintf("%v", lineResp)
	}, nil
}

func main() {
	lambda.Start(aws_billing_notify)
}
