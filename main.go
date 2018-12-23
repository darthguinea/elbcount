package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"./src/cw"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/darthguinea/golib/log"
)

func createSession(r string) *session.Session {
	return session.New(&aws.Config{
		Region: aws.String(r),
	})
}

func main() {
	var (
		flagVerbose bool
		flagList    bool
		flagWatch   bool
		flagRegion  string
		flagNames   string
	)
	flag.BoolVar(&flagVerbose, "v", false, "verbose")
	flag.StringVar(&flagRegion, "r", "ap-southeast-2", "-r <region>")
	flag.StringVar(&flagNames, "n", "", "-n <elb_name,elb_name,elb_name>")
	flag.BoolVar(&flagList, "l", false, "-l list elbs in the region")
	flag.BoolVar(&flagWatch, "w", false, "-w watch the request count")
	flag.Parse()

	if flagVerbose {
		log.SetLevel(log.INFO)
	}

	session := createSession(flagRegion)

	if flagList {
		cw.List(session)
		os.Exit(0)
	}

	elb := strings.Split(flagNames, ",")

	startTime := time.Now().Add(time.Duration(-60) * time.Minute)
	endTime := time.Now()

	log.Print("Between [%v] -> [%v]",
		startTime.Format("2006-1-2 15:04"),
		endTime.Format("2006-1-2 15:04"))

	for {
		for _, elbName := range elb {
			data := cw.GetMetric(session, elbName, &startTime, &endTime)
			cw.DrawGraph(elbName, data)
		}
		if !flagWatch {
			os.Exit(0)
		}
		time.Sleep(10 * time.Second)
	}
}
