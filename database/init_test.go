package database

import "testing"

func TestNewDatabase(t *testing.T)  {
	NewDatabase("0.0.0.0", "80", "mysql")
}
