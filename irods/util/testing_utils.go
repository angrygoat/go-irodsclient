// Package util for unit and functional testing.
package util

import (
	"errors"
	"os"

	"github.com/cyverse/go-irodsclient/irods/common"
	"github.com/magiconair/properties"
)

// LoadPropertiesFromEnvPath - Load testing properties based on the environmental path variable
func LoadPropertiesFromEnvPath() (*properties.Properties, error) {
	LogInfo("LoadProperties()")

	// if I don't have a path provided, look for the env variable

	myPropsPath := os.Getenv(common.TestingPropertiesPath)
	if myPropsPath == "" {
		LogError("no properties path or property path environment variable set")
		return nil, errors.New("no testing properties found")
	}

	LogDebugf("properties loaded from %s", myPropsPath)
	p := properties.MustLoadFile(myPropsPath, properties.UTF8)
	LogDebugf("testing props:%s", p)
	return p, nil

}
