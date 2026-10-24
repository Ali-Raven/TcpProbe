package tests

import (
	"net"
	"testing"
	"time"

	"github.com/TcpProb/internal/probe"
)

func TestProbeSuccess(t *testing.T) {
	// create temp listener
	lisnr, err := net.Listen("tcp", "10.60.40.55:0")
	if err != nil {
		t.Errorf("Failed to create listener: %v", err)
	}

	fullAddr := lisnr.Addr().String()

	var counter probe.Counter
	_, err = probe.Probe(1*time.Second, fullAddr, &counter)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}

	if counter.Success != 1 {
		t.Errorf("Expected 1 success, got %d", counter.Success)
	}

	if counter.Failed != 0 {
		t.Errorf("Expected 0 failed, got %d", counter.Failed)
	}

	if counter.Sent != 1 {
		t.Errorf("Expected 1 sent, got %d", counter.Sent)
	}

	if counter.Loss != 0 {
		t.Errorf("Expected 0 loss, got %d", counter.Loss)
	}
}

func TestTimeout(t *testing.T) {
}
