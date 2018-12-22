package cw

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"github.com/darthguinea/golib/log"
)

func GetMetric(session *session.Session) {
	svc := cloudwatch.New(session)

	startTime := time.Now().Add(time.Duration(-10) * time.Minute)
	endTime := time.Now()

	log.Print("Between [%v] -> [%v]", startTime.Format("2006-1-2 15:04"), endTime.Format("2006-1-2 15:04"))

	period := int64(300)
	filter := &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []*cloudwatch.MetricDataQuery{{
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Namespace:  aws.String("AWS/ApplicationELB"),
					MetricName: aws.String("RequestCount"),
					Dimensions: []*cloudwatch.Dimension{{
						Name:  aws.String("LoadBalancer"),
						Value: aws.String("<load balancer name>"),
					}},
				},
				Period: &period,
				Stat:   aws.String("Average"),
			},
		}},
	}
	op, _ := svc.GetMetricData(filter)

	log.Print("%v", filter)
	log.Print("%v", op)
}
