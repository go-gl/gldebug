package gpuinfo

import (
	"log"
	"testing"
	"time"
)

func TestNvidia(t *testing.T) {
	ch := PollGPUMemory()

	go func() {
		for x := range ch {
			log.Printf("GPU status: %v", x)
		}
	}()

	time.Sleep(2 * time.Second)
	close(ch)
}
