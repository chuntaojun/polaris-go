//@Author: springliao
// @Description:
//@Time: 2021/10/06 17:57

package main

import (
	"flag"
	"log"

	"github.com/polarismesh/polaris-go/api"
)

var (
	namespace string
	service   string
	token     string
)

func initArgs() {
	flag.StringVar(&namespace, "namespace", "default", "namespace")
	flag.StringVar(&service, "service", "", "service")
	flag.StringVar(&token, "token", "", "token")
}

const (
	host          = "127.0.0.1"
	startPort     = 2001
	instanceCount = 5
)

func main() {
	initArgs()
	flag.Parse()
	if len(namespace) == 0 || len(service) == 0 {
		log.Print("namespace and service are required")
		return
	}
	limiter, err := api.NewLimitAPI()
	if nil != err {
		log.Fatalf("fail to create LimitAPI, err is %v", err)
	}
	defer limiter.Destroy()

	req := api.NewQuotaRequest()

	limiter.GetQuota(req)

}
