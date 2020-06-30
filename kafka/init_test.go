package kafka

import "testing"

func TestNewKafka(t *testing.T)  {
	NewKafka("0.0.0.0", "80", "mysql")
}
