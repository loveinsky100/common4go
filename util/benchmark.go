/**
 * @author leo
 * @date 2020/8/18 11:24 上午
 */
package util

import (
	"fmt"
	"time"
)

type BenchmarkFunc func()

func Run(runTimes int, benchmarkFunc BenchmarkFunc) {
	start := time.Now().UnixNano() / 1e6
	for index := 0; index < runTimes; index ++ {
		benchmarkFunc()
	}

	end := time.Now().UnixNano() / 1e6
	fmt.Printf("[Benchmark] run method in %d times, cost: %d ms\n", runTimes, end - start)
}