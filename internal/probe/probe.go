package probe

import (
	"fmt"
	"net"
	"time"
)

type Counter struct {
	Sent    int
	Loss    int
	Success int
	Failed  int
	Avg     float32
}

func Probe(timeout time.Duration, fullAddr string, counter *Counter) (time.Duration, error) {
	// calling tests functions
	startTimer := time.Now()

	time.Sleep(500 * time.Millisecond)
	conn, err := net.DialTimeout("tcp", fullAddr, timeout)
	if err != nil {
		counter.Failed++
		return time.Duration(counter.Failed), fmt.Errorf("Error for connection : %v", err)
	}
	counter.Sent++
	counter.Success++
	timeElapsed := time.Since(startTimer)

	conn.Close()
	fmt.Printf("Reply from %s , time=%.2fs\n", fullAddr, timeElapsed.Seconds())

	return time.Duration(counter.Success), nil
}
