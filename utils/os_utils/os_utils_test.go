package os_utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvOrDefault(t *testing.T) {
	key, value, fallback := "FOO", "1", "default"
	os.Setenv(key, value)

	assert.EqualValues(t, GetEnvOrDefault(key, fallback), value,
		"should get the value from the environment")

	key, value, fallback = "BAR", "2", "default2"
	assert.EqualValues(t, GetEnvOrDefault(key, fallback), fallback,
		"should default to the given value, when no variable exists in the environment")
}
