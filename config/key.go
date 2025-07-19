//nolint:unused // This file contains helper functions that may be referenced by configuration loading in other modules.
package config

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func mustGetString(key string) string {
	mustHave(key)
	return viper.GetString(key)
}

func mustGetStringArray(key string) []string {
	mustHave(key)
	return optionalGetStringArray(key)
}

func optionalGetStringArray(key string) []string {
	value := viper.GetString(key)
	if value == "" {
		return []string{}
	}
	strs := strings.Split(value, ",")
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	return strs
}

func optionalGetStringMap(key string) map[string]bool {
	result := make(map[string]bool)
	for _, str := range optionalGetStringArray(key) {
		result[str] = true
	}
	return result
}

func mustGetBool(key string) bool {
	mustHave(key)
	return viper.GetBool(key)
}

func mustGetDurationMs(key string) time.Duration {
	return time.Millisecond * time.Duration(mustGetInt(key))
}

func mustGetDurationMinute(key string) time.Duration {
	return time.Minute * time.Duration(mustGetInt(key))
}

func mustGetDurationSeconds(key string) time.Duration {
	return time.Second * time.Duration(mustGetInt(key))
}

func mustGetInt(key string) int {
	mustHave(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}

func mustGetFloat(key string) float64 {
	v, err := strconv.ParseFloat((viper.GetString(key)), 64)
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid float value", key))
	}

	return v
}

func mustHave(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}

func mustGetStringMapInt(key string) map[string]int {
	mustHave(key)
	strs := strings.Split(viper.GetString(key), ",")
	stringMap := make(map[string]int)
	for _, str := range strs {
		kv := strings.Split(str, ":")
		val, err := strconv.ParseInt(kv[1], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse key/value pair: %s, %s", kv[0], kv[1]))
		}
		stringMap[kv[0]] = int(val)
	}
	return stringMap
}

func mustGetBoolMap(key string) map[string]bool {
	mustHave(key)
	strs := strings.Split(viper.GetString(key), ",")
	stringBoolMap := make(map[string]bool)
	for _, str := range strs {
		stringBoolMap[str] = true
	}
	return stringBoolMap
}
