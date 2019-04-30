package main

import "github.com/sivsivsree/go-now/concurrency"

func main() {

	concurrency.UsingRange()
	concurrency.BufferedChannel()
	concurrency.WorkerPulls()
	concurrency.UsingSelect()
	concurrency.Simple()
	concurrency.UsingRange()
	concurrency.BufferedChannel()
}
