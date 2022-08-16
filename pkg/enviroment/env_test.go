package enviroment_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/fsvxavier/unico/pkg/enviroment"
)

func TestSetFileConfig(t *testing.T) {

	os.Setenv("ENV", "teste")

	var env enviroment.ConfigEnviroment
	env.Env = "teste"
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_INT")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvConfig(t *testing.T) {

	os.Setenv("ENV", "teste")

	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_INT")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvString(t *testing.T) {

	os.Setenv("ENV", "teste")

	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_STRING")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvEmpty(t *testing.T) {

	os.Setenv("ENV", "teste")

	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_EMPTY")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvBool(t *testing.T) {

	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_BOOL")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}

func TestEnvInt(t *testing.T) {
	os.Setenv("ENV", "production")

	var env enviroment.ConfigEnviroment
	env.SetFileConfig("./mock/env.json")
	env.GetTag("ENV")

	ret, err := env.GetTag("TEST_ENV_INT")
	assert.NoError(t, err)
	assert.NotNil(t, ret)
}
