package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	SetLogLevel(9)
}

func shutdown() {
}

func TestPropsFromEnv(t *testing.T) {
	setup()
	testingProps, err := LoadPropertiesFromEnvPath()

	assert.Nil(t, err, "error message %s", "formatted")
	assert.NotNil(t, testingProps, "testingProps was not returned")

}
