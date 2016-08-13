package envmap

import (
	"os"
	"regexp"
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

func TestPushPop(t *testing.T) {

	assert := assert.New(t)

	filter := func(s string) bool {
		return s == "FOO_BAR_WEE"
	}

	m1 := Envmap{
		"BAR":         "bar1",
		"FOO_BAR_WEE": "foo1",
	}

	m2 := m1.Push("_", filter)

	m2["FOO_BAR_WEE"] = "foo2"

	assert.NotEqual(m1, m2)
	assert.Equal("bar1", m2["BAR"])
	assert.Equal("foo2", m2["FOO_BAR_WEE"])
	assert.Equal("foo1", m2["_FOO_BAR_WEE"])

	m3 := m2.Push("_", filter)

	m3["FOO_BAR_WEE"] = "foo3"

	assert.NotEqual(m2, m3)
	assert.Equal("bar1", m3["BAR"])
	assert.Equal("foo3", m3["FOO_BAR_WEE"])
	assert.Equal("foo2", m3["_FOO_BAR_WEE"])
	assert.Equal("foo1", m3["__FOO_BAR_WEE"])

	m4 := m3.Pop("_", filter)

	assert.NotEqual(m3, m4)
	assert.Equal("bar1", m4["BAR"])
	assert.Equal("foo2", m4["FOO_BAR_WEE"])
	assert.Equal("foo1", m4["_FOO_BAR_WEE"])

	m5 := m4.Pop("_", filter)

	assert.NotEqual(m4, m5)
	assert.Equal("bar1", m5["BAR"])
	assert.Equal("foo1", m5["FOO_BAR_WEE"])

	m6 := m5.Pop("_", filter)

	assert.NotEqual(m5, m6)
	assert.Equal("bar1", m5["BAR"])

	m7 := m6.Pop("_", filter)
	assert.Equal(m6, m7)

}

func TestSubset(t *testing.T) {

	assert := assert.New(t)

	filter := regexp.MustCompile("^(?:FOO|BOO|GOO)$").MatchString

	m1 := Envmap{
		"BAR": "bar",
		"BOO": "boo",
		"FOO": "foo",
	}

	m2 := m1.Subset(filter)

	assert.NotEqual(m1, m2)
	assert.Len(m2, 2)
	assert.Contains(m2, "FOO")
	assert.Contains(m2, "BOO")

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
