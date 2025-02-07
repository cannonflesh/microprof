package main

import (
	"runtime"
	"sync"

	"github.com/cannonflesh/microprof"
	"github.com/sirupsen/logrus"
)

func main() {
	lgr := logrus.New()
	lgr.Level = logrus.DebugLevel
	logEntry := logrus.NewEntry(lgr)
	var (
		mx sync.Mutex
		wg sync.WaitGroup
	)

	numCPU := runtime.NumCPU()
	for i := 1; i <= numCPU; i++ {
		wg.Add(1)
		byCPU := i%2 == 0
		go func(l *logrus.Entry, iter int) {
			mx.Lock()
			defer mx.Unlock()

			wasted := make([]int, 0)
			for ii := 1; ii <= 10000; ii++ {
				wasted = append(wasted, ii)
				if ii%1000 == 0 {
					lgr.Infof("*** goroutine #%d, iteration #%d: ***", iter, ii)
					microprof.PrintProfilingInfo(lgr, microprof.UnitsMb, byCPU)
					lgr.Infof("")
				}
			}

			wg.Done()
		}(logEntry, i)
	}

	wg.Wait()
}
