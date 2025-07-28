package config

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setupViperForTest() {
	viper.Reset()
	viper.Set("TEST_STRING", "test_value")
	viper.Set("TEST_INT", "123")
	viper.Set("TEST_BOOL", "true")
	viper.Set("TEST_FLOAT", "12.34")
	viper.Set("TEST_ARRAY", "item1,item2,item3")
	viper.Set("TEST_MAP", "key1:10,key2:20")
	viper.Set("TEST_BOOL_MAP", "option1,option2")
}

func TestMustGetString(t *testing.T) {
	setupViperForTest()

	result := mustGetString("TEST_STRING")
	assert.Equal(t, "test_value", result)
}

func TestMustGetString_Panic(t *testing.T) {
	setupViperForTest()

	assert.Panics(t, func() {
		mustGetString("NON_EXISTENT_KEY")
	})
}

func TestMustGetInt(t *testing.T) {
	setupViperForTest()

	result := mustGetInt("TEST_INT")
	assert.Equal(t, 123, result)
}

func TestMustGetInt_InvalidValue(t *testing.T) {
	setupViperForTest()

	assert.Panics(t, func() {
		mustGetInt("TEST_STRING") // "test_value" is not a valid integer
	})
}

func TestMustGetBool(t *testing.T) {
	setupViperForTest()

	result := mustGetBool("TEST_BOOL")
	assert.True(t, result)
}

func TestMustGetFloat(t *testing.T) {
	setupViperForTest()

	result := mustGetFloat("TEST_FLOAT")
	assert.Equal(t, 12.34, result)
}

func TestMustGetDurationMs(t *testing.T) {
	setupViperForTest()

	result := mustGetDurationMs("TEST_INT")
	expected := time.Duration(123) * time.Millisecond
	assert.Equal(t, expected, result)
}

func TestMustGetDurationMinute(t *testing.T) {
	setupViperForTest()

	result := mustGetDurationMinute("TEST_INT")
	expected := time.Duration(123) * time.Minute
	assert.Equal(t, expected, result)
}

func TestMustGetDurationSeconds(t *testing.T) {
	setupViperForTest()

	result := mustGetDurationSeconds("TEST_INT")
	expected := time.Duration(123) * time.Second
	assert.Equal(t, expected, result)
}

func TestMustGetStringArray(t *testing.T) {
	setupViperForTest()

	result := mustGetStringArray("TEST_ARRAY")
	expected := []string{"item1", "item2", "item3"}
	assert.Equal(t, expected, result)
}

func TestOptionalGetStringArray(t *testing.T) {
	setupViperForTest()

	result := optionalGetStringArray("TEST_ARRAY")
	expected := []string{"item1", "item2", "item3"}
	assert.Equal(t, expected, result)

	// Test with empty key
	emptyResult := optionalGetStringArray("NON_EXISTENT_KEY")
	assert.Equal(t, []string{}, emptyResult)
}

func TestOptionalGetStringMap(t *testing.T) {
	setupViperForTest()

	result := optionalGetStringMap("TEST_ARRAY")
	expected := map[string]bool{
		"item1": true,
		"item2": true,
		"item3": true,
	}
	assert.Equal(t, expected, result)
}

func TestMustGetStringMapInt(t *testing.T) {
	setupViperForTest()

	result := mustGetStringMapInt("TEST_MAP")
	expected := map[string]int{
		"key1": 10,
		"key2": 20,
	}
	assert.Equal(t, expected, result)
}

func TestMustGetBoolMap(t *testing.T) {
	setupViperForTest()

	result := mustGetBoolMap("TEST_BOOL_MAP")
	expected := map[string]bool{
		"option1": true,
		"option2": true,
	}
	assert.Equal(t, expected, result)
}

func TestMustHave(t *testing.T) {
	setupViperForTest()

	// Should not panic for existing key
	assert.NotPanics(t, func() {
		mustHave("TEST_STRING")
	})

	// Should panic for non-existent key
	assert.Panics(t, func() {
		mustHave("NON_EXISTENT_KEY")
	})
}
