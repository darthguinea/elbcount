package main

import (
	"flag"

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
	)
	flag.BoolVar(&flagVerbose, "v", false, "verbose")
	flag.Parse()

	if flagVerbose {
		log.SetLevel(log.INFO)
	}

	session := createSession("ap-southeast-2")
	cw.GetMetric(session)
}
