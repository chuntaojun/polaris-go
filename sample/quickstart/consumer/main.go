/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/polarismesh/polaris-go/api"
)

var (
	namespace string
	service   string
)

func initArgs() {
	flag.StringVar(&namespace, "namespace", "default", "namespace")
	flag.StringVar(&service, "service", "tdsql-ops-server", "service")
}

func main() {
	initArgs()
	flag.Parse()
	if len(namespace) == 0 || len(service) == 0 {
		log.Print("namespace and service are required")
		return
	}
	consumer, err := api.NewConsumerAPI()
	if nil != err {
		log.Fatalf("fail to create consumerAPI, err is %v", err)
	}
	defer consumer.Destroy()

	log.Printf("start to invoke getAllInstances operation")
	getAllRequest := &api.GetAllInstancesRequest{}
	getAllRequest.Namespace = namespace
	getAllRequest.Service = service
	allInstResp, err := consumer.GetAllInstances(getAllRequest)
	if nil != err {
		log.Fatalf("fail to getAllInstances, err is %v", err)
	}
	instances := allInstResp.GetInstances()
	if len(instances) > 0 {
		for i, instance := range instances {
			log.Printf("instance getAllInstances %d is %s:%d", i, instance.GetHost(), instance.GetPort())
		}
	}

	log.Printf("start to invoke getOneInstance operation")
	for i := 0; i < 1; i++ {
		getAllRequest := &api.GetAllInstancesRequest{}
		getAllRequest.Namespace = namespace
		getAllRequest.Service = service
		allInstResp, err := consumer.GetAllInstances(getAllRequest)
		if nil != err {
			log.Fatalf("fail to getOneInstance, err is %v", err)
		}

		instances := allInstResp.GetInstances()
		log.Printf("%d get", i)
		for _, instance := range instances {
			log.Printf("instance is %s:%d", instance.GetHost(), instance.GetPort())
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
