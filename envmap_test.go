package envmap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEnv(t *testing.T) {

	assert := assert.New(t)

	os.Setenv("foo", "bar")
	assert.Contains(Import().ToEnv(), "foo=bar")

	os.Setenv("wee", "gee=dee")
	assert.Contains(Import().ToEnv(), "wee=gee=dee")

}

func TestToMap(t *testing.T) {

	assert := assert.New(t)

	assert.Regexp(`bin\/(go|godep)$`, Import()["_"])

	os.Setenv("foo", "bar")
	assert.Equal("bar", Import()["foo"])

	os.Setenv("wee", "gee=dee")
	assert.Equal("gee=dee", Import()["wee"])

	os.Clearenv()

	assert.Empty(Import())

}
