package main

import (
	grpcc "mail-service/grpc"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpcc.GRPCListen()
	}()

	wg.Wait()
}
