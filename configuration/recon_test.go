package recon

import (
	"testing"
	"time"
)

type ConfigTestSpec struct {
	StringVar string
	IntVar    int
	BoolVar   bool
	DoubleVar float64
	FloatVar  float32
	Template  string
	Timeout   time.Duration
	StrList   []string
	StrMap    map[string]int
}

func TestLoadConfig(t *testing.T) {
	var configSpec ConfigTestSpec
	err := LoadConfig("test.env", &configSpec, "test")
	if err != nil {
		t.Errorf("Failed to load test.env", err)
	}

}
