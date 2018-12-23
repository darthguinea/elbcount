package cw

import (
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/guptarohit/asciigraph"

	"github.com/darthguinea/golib/log"
)

func List(session *session.Session) {
	svc := elbv2.New(session)

	op, err := svc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{})
	if err != nil {
		log.Error("%v", err)
		os.Exit(1)
	}

	log.Print("%35v - %v", "Short Name", "Name")
	for _, elb := range op.LoadBalancers {
		val := strings.Split(*elb.LoadBalancerArn, "/")
		log.Print("%35v - %v", *elb.LoadBalancerName, log.Sprintf("%v/%v/%v", val[1], val[2], val[3]))
	}
}

func GetMetric(session *session.Session,
	name string,
	startTime *time.Time, endTime *time.Time) []*float64 {
	svc := cloudwatch.New(session)

	period := int64(60)
	op, err := svc.GetMetricData(&cloudwatch.GetMetricDataInput{
		StartTime: startTime,
		EndTime:   endTime,
		MetricDataQueries: []*cloudwatch.MetricDataQuery{{
			Id: aws.String("test"),
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Namespace:  aws.String("AWS/ApplicationELB"),
					MetricName: aws.String("RequestCount"),
					Dimensions: []*cloudwatch.Dimension{{
						Name:  aws.String("LoadBalancer"),
						Value: aws.String(name),
					}},
				},
				Period: &period,
				Stat:   aws.String("Sum"),
			},
		}},
	})
	if err != nil {
		log.Error("%v", err)
		return nil
	}
	return op.MetricDataResults[0].Values
}

func DrawGraph(name string, x []*float64) {
	data := []float64{}
	for _, i := range x {
		data = append(data, *i)
	}
	log.Print("%v", name)
	if len(data) > 0 {
		graph := asciigraph.Plot(data, asciigraph.Height(20))
		log.Print("%v", graph)
		log.Print("")
	}
}
