// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpuinfo

import (
	"log"
	"os/exec"
)

func haveExe(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// Returns a GPU{} struct every ~1s, or blocks indefinitely if there is no known
// way to get the spare GPU memory information.
// Polling is terminated by calling close() on the channel
func PollGPUMemory() chan GPU {
	switch {
	case haveExe("nvidia-smi"):
		return PollNvidiaGPUMemory()
	}
	log.Print("PollGPUMemory() doesn't know any way to measure GPU memory")
	return make(chan GPU)
}
