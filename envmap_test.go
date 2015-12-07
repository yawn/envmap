package envmap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEnv(t *testing.T) {

	assert := assert.New(t)

	os.Setenv("foo", "bar")
	assert.Contains(ToMap(os.Environ()).ToEnv(), "foo=bar")

	os.Setenv("wee", "gee=dee")
	assert.Contains(ToMap(os.Environ()).ToEnv(), "wee=gee=dee")

}

func TestToMap(t *testing.T) {

	assert := assert.New(t)

	assert.Regexp(`bin\/(go|godep)$`, ToMap(os.Environ())["_"])

	os.Setenv("foo", "bar")
	assert.Equal("bar", ToMap(os.Environ())["foo"])

	os.Setenv("wee", "gee=dee")
	assert.Equal("gee=dee", ToMap(os.Environ())["wee"])

	os.Clearenv()

	assert.Empty(ToMap(os.Environ()))

}
