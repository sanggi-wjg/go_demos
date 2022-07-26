package builds_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sanggi-wjg/go_demos/sample/builds"
)

func TestAdd(t *testing.T) {
	// given
	var a, b int = 1, 2

	// when
	result := builds.Add(1, 2)

	// then
	assert.Equal(t, result, a+b)
}

func TestJsonDecode(t *testing.T) {
	// given
	text := []byte(`{"name":"snow","age":22,"weight":60.1,"location":"41.40338, 2.17403","hobbies":["workout","watching youtube"]}`)

	// when
	result, err := builds.JsonDecode(text)
	t.Log(result)

	// then
	assert.Equal(t, err, nil)
}

func TestJsonDecodeWithFailCase(t *testing.T) {
	// given : 컴마 삭제
	text := []byte(`{"name":"snow","age":22,"weight":60.1"location":"41.40338, 2.17403","hobbies":["workout","watching youtube"]}`)

	// when
	_, err := builds.JsonDecode(text) // invalid character '"' after object key:value pair
	t.Log(err)

	// then
	assert.Equal(t, err.Error(), `invalid character '"' after object key:value pair`)
}

func TestJsonDecodeWithFailCase2(t *testing.T) {
	// given : N개 공백
	text := []byte(`{"name":"snow","age":22,"weight":60.1",  location":"41.40338, 2.17403","hobbies":["workout","watching youtube"]}`)

	// when
	_, err := builds.JsonDecode(text)

	// then
	assert.Equal(t, err.Error(), `invalid character '"' after object key:value pair`)
}

func TestJsonDecodeWithFailCase3(t *testing.T) {
	// given : " 삭제
	text := []byte(`{"name":"snow","age":22,"weight":60.1",location":"41.40338, 2.17403","hobbies":["workout","watching youtube"]}`)

	// when
	_, err := builds.JsonDecode(text)

	// then
	assert.Equal(t, err.Error(), `invalid character '"' after object key:value pair`)
}

func TestJsonDecodeWithFailCase4(t *testing.T) {
	// given : " 끝자리 삭제
	text := []byte(`{"name":"snow","age":22,"weight":60.1",location":"41.40338, 2.17403"`)

	// when
	_, err := builds.JsonDecode(text)

	// then
	assert.Equal(t, err.Error(), `invalid character '"' after object key:value pair`)
}
