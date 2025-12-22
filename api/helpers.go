package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	// DefaultUserAgent is the default user agent string
	DefaultUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Mobile/15E148 Safari/604.1"
)

// GenerateRandomString generates a string of a specified length
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Validate validates a struct using the validator library
func Validate(v interface{}) error {
	return validator.New().Struct(v)
}

// GetExecutablePath returns the path to the current executable
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(ex)
}

// GetAbsolutePath returns the absolute path for a given path
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

// IsNil checks if a value is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return false
}

// IsString checks if a value is a string
func IsString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

// IsInt checks if a value is an integer
func IsInt(i interface{}) bool {
	_, ok := i.(int)
	return ok
}

// IsInt64 checks if a value is an int64
func IsInt64(i interface{}) bool {
	_, ok := i.(int64)
	return ok
}

// IsFloat64 checks if a value is a float64
func IsFloat64(i interface{}) bool {
	_, ok := i.(float64)
	return ok
}

// IsBool checks if a value is a boolean
func IsBool(i interface{}) bool {
	_, ok := i.(bool)
	return ok
}

// GetBoolValue returns the value of a boolean
func GetBoolValue(i interface{}) (bool, error) {
	b, ok := i.(bool)
	if !ok {
		return false, fmt.Errorf("value is not a boolean")
	}
	return b, nil
}

// GetIntValue returns the value of an integer
func GetIntValue(i interface{}) (int, error) {
	v, ok := i.(int)
	if !ok {
		return 0, fmt.Errorf("value is not an integer")
	}
	return v, nil
}

// GetInt64Value returns the value of an int64
func GetInt64Value(i interface{}) (int64, error) {
	v, ok := i.(int64)
	if !ok {
		return 0, fmt.Errorf("value is not an int64")
	}
	return v, nil
}

// GetFloat64Value returns the value of a float64
func GetFloat64Value(i interface{}) (float64, error) {
	v, ok := i.(float64)
	if !ok {
		return 0, fmt.Errorf("value is not a float64")
	}
	return v, nil
}

// GetStringValue returns the value of a string
func GetStringValue(i interface{}) (string, error) {
	s, ok := i.(string)
	if !ok {
		return "", fmt.Errorf("value is not a string")
	}
	return s, nil
}

// GetStringsFromSlice returns a slice of strings from a given slice
func GetStringsFromSlice(i interface{}) ([]string, error) {
	s, ok := i.([]string)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of strings")
	}
	return s, nil
}

// GetStringsFromSliceOrDefault returns a slice of strings from a given slice or an empty slice if not found
func GetStringsFromSliceOrDefault(i interface{}) ([]string, error) {
	s, err := GetStringsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetIntsFromSlice returns a slice of integers from a given slice
func GetIntsFromSlice(i interface{}) ([]int, error) {
	s, ok := i.([]int)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of integers")
	}
	return s, nil
}

// GetIntsFromSliceOrDefault returns a slice of integers from a given slice or an empty slice if not found
func GetIntsFromSliceOrDefault(i interface{}) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetInt64sFromSlice returns a slice of int64s from a given slice
func GetInt64sFromSlice(i interface{}) ([]int64, error) {
	s, ok := i.([]int64)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of int64s")
	}
	return s, nil
}

// GetInt64sFromSliceOrDefault returns a slice of int64s from a given slice or an empty slice if not found
func GetInt64sFromSliceOrDefault(i interface{}) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetFloat64sFromSlice returns a slice of float64s from a given slice
func GetFloat64sFromSlice(i interface{}) ([]float64, error) {
	s, ok := i.([]float64)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of float64s")
	}
	return s, nil
}

// GetFloat64sFromSliceOrDefault returns a slice of float64s from a given slice or an empty slice if not found
func GetFloat64sFromSliceOrDefault(i interface{}) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBoolsFromSlice returns a slice of booleans from a given slice
func GetBoolsFromSlice(i interface{}) ([]bool, error) {
	s, ok := i.([]bool)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of booleans")
	}
	return s, nil
}

// GetBoolsFromSliceOrDefault returns a slice of booleans from a given slice or an empty slice if not found
func GetBoolsFromSliceOrDefault(i interface{}) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBoolSliceOrDefault returns a slice of booleans from a given slice or an empty slice if not found
func GetBoolSliceOrDefault(i interface{}, defaultValue bool) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return []bool{defaultValue}, nil
	}
	return s, nil
}

// GetIntSliceOrDefault returns a slice of integers from a given slice or an empty slice if not found
func GetIntSliceOrDefault(i interface{}, defaultValue int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return []int{defaultValue}, nil
	}
	return s, nil
}

// GetInt64SliceOrDefault returns a slice of int64s from a given slice or an empty slice if not found
func GetInt64SliceOrDefault(i interface{}, defaultValue int64) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return []int64{defaultValue}, nil
	}
	return s, nil
}

// GetFloat64SliceOrDefault returns a slice of float64s from a given slice or an empty slice if not found
func GetFloat64SliceOrDefault(i interface{}, defaultValue float64) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return []float64{defaultValue}, nil
	}
	return s, nil
}

// GetStringsOrDefault returns a string from a given string or a default string if not found
func GetStringsOrDefault(i interface{}, defaultValue string) (string, error) {
	s, err := GetStringValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetStringsOrDefault returns a slice of strings from a given slice or a default slice if not found
func GetStringsOrDefault(i interface{}, defaultValue []string) ([]string, error) {
	s, err := GetStringsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetIntOrDefault returns an integer from a given integer or a default integer if not found
func GetIntOrDefault(i interface{}, defaultValue int) (int, error) {
	v, err := GetIntValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetInt64OrDefault returns an int64 from a given int64 or a default int64 if not found
func GetInt64OrDefault(i interface{}, defaultValue int64) (int64, error) {
	v, err := GetInt64Value(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetFloat64OrDefault returns a float64 from a given float64 or a default float64 if not found
func GetFloat64OrDefault(i interface{}, defaultValue float64) (float64, error) {
	v, err := GetFloat64Value(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetBoolOrDefault returns a boolean from a given boolean or a default boolean if not found
func GetBoolOrDefault(i interface{}, defaultValue bool) (bool, error) {
	b, err := GetBoolValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return b, nil
}

// GetInt64sOrDefault returns a slice of int64s from a given slice or a default slice if not found
func GetInt64sOrDefault(i interface{}, defaultValue []int64) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetFloat64sOrDefault returns a slice of float64s from a given slice or a default slice if not found
func GetFloat64sOrDefault(i interface{}, defaultValue []float64) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetBoolsOrDefault returns a slice of booleans from a given slice or a default slice if not found
func GetBoolsOrDefault(i interface{}, defaultValue []bool) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetIntsOrDefault returns a slice of integers from a given slice or a default slice if not found
func GetIntsOrDefault(i interface{}, defaultValue []int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetMapOrDefault returns a map from a given map or a default map if not found
func GetMapOrDefault(i interface{}, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map")
	}
	return m, nil
}

// GetIntOrDefault returns a slice of integers from a given slice or a default slice if not found
func GetIntsOrDefault(i interface{}, defaultValue []int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetMapOrDefaultString returns a map of strings from a given map or a default map if not found
func GetMapOrDefaultString(i interface{}, defaultValue map[string]string) (map[string]string, error) {
	m, ok := i.(map[string]string)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of strings")
	}
	return m, nil
}

// GetMapOrDefaultInt returns a map of integers from a given map or a default map if not found
func GetMapOrDefaultInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of integers")
	}
	return m, nil
}

// GetMapOrDefaultInt64 returns a map of int64s from a given map or a default map if not found
func GetMapOrDefaultInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of int64s")
	}
	return m, nil
}

// GetMapOrDefaultFloat64 returns a map of float64s from a given map or a default map if not found
func GetMapOrDefaultFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of float64s")
	}
	return m, nil
}

// GetMapOrDefaultBool returns a map of booleans from a given map or a default map if not found
func GetMapOrDefaultBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of booleans")
	}
	return m, nil
}

// GetMapOrDefaultStringSlice returns a map of slices of strings from a given map or a default map if not found
func GetMapOrDefaultStringSlice(i interface{}, defaultValue map[string][]string) (map[string][]string, error) {
	m, ok := i.(map[string][]string)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of strings")
	}
	return m, nil
}

// GetMapOrDefaultIntSlice returns a map of slices of integers from a given map or a default map if not found
func GetMapOrDefaultIntSlice(i interface{}, defaultValue map[string][]int) (map[string][]int, error) {
	m, ok := i.(map[string][]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of integers")
	}
	return m, nil
}

// GetMapOrDefaultInt64Slice returns a map of slices of int64s from a given map or a default map if not found
func GetMapOrDefaultInt64Slice(i interface{}, defaultValue map[string][]int64) (map[string][]int64, error) {
	m, ok := i.(map[string][]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of int64s")
	}
	return m, nil
}

// GetMapOrDefaultFloat64Slice returns a map of slices of float64s from a given map or a default map if not found
func GetMapOrDefaultFloat64Slice(i interface{}, defaultValue map[string][]float64) (map[string][]float64, error) {
	m, ok := i.(map[string][]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of float64s")
	}
	return m, nil
}

// GetMapOrDefaultBoolSlice returns a map of slices of booleans from a given map or a default map if not found
func GetMapOrDefaultBoolSlice(i interface{}, defaultValue map[string][]bool) (map[string][]bool, error) {
	m, ok := i.(map[string][]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of booleans")
	}
	return m, nil
}

// GetMapOrDefaultStringInt returns a map of pairs of strings and integers from a given map or a default map if not found
func GetMapOrDefaultStringInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and integers")
	}
	return m, nil
}

// GetMapOrDefaultStringInt64 returns a map of pairs of strings and int64s from a given map or a default map if not found
func GetMapOrDefaultStringInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and int64s")
	}
	return m, nil
}

// GetMapOrDefaultStringFloat64 returns a map of pairs of strings and float64s from a given map or a default map if not found
func GetMapOrDefaultStringFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and float64s")
	}
	return m, nil
}

// GetMapOrDefaultStringBool returns a map of pairs of strings and booleans from a given map or a default map if not found
func GetMapOrDefaultStringBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and booleans")
	}
	return m, nil
}

// GetMapOrDefaultIntInt64 returns a map of pairs of integers and int64s from a given map or a default map if not found
func GetMapOrDefaultIntInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and int64s")
	}
	return m, nil
}

// GetMapOrDefaultIntFloat64 returns a map of pairs of integers and float64s from a given map or a default map if not found
func GetMapOrDefaultIntFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and float64s")
	}
	return m, nil
}

// GetMapOrDefaultIntBool returns a map of pairs of integers and booleans from a given map or a default map if not found
func GetMapOrDefaultIntBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and booleans")
	}
	return m, nil
}

// GetMapOrDefaultInt64Float64 returns a map of pairs of int64s and float64s from a given map or a default map if not found
func GetMapOrDefaultInt64Float64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of int64s and float64s")
	}
	return m, nil
}

// GetMapOrDefaultInt64Bool returns a map of pairs of int64s and booleans from a given map or a default map if not found
func GetMapOrDefaultInt64Bool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of int64s and booleans")
	}
	return m, nil
}

// GetMapOrDefaultFloat64Bool returns a map of pairs of float64s and booleans from a given map or a default map if not found
func GetMapOrDefaultFloat64Bool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of float64s and booleans")
	}
	return m, nil
}

// GetMapOrDefaultBoolInt returns a map of pairs of booleans and integers from a given map or a default map if not found
func GetMapOrDefaultBoolInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and integers")
	}
	return m, nil
}

// GetMapOrDefaultBoolInt64 returns a map of pairs of booleans and int64s from a given map or a default map if not found
func GetMapOrDefaultBoolInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and int64s")
	}
	return m, nil
}

// GetMapOrDefaultBoolFloat64 returns a map of pairs of booleans and float64s from a given map or a default map if not found
func GetMapOrDefaultBoolFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and float64s")
	}
	return m, nil
}

// IsNil checks if a value is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return false
}

// IsString checks if a value is a string
func IsString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

// IsInt checks if a value is an integer
func IsInt(i interface{}) bool {
	_, ok := i.(int)
	return ok
}

// IsInt64 checks if a value is an int64
func IsInt64(i interface{}) bool {
	_, ok := i.(int64)
	return ok
}

// IsFloat64 checks if a value is a float64
func IsFloat64(i interface{}) bool {
	_, ok := i.(float64)
	return ok
}

// IsBool checks if a value is a boolean
func IsBool(i interface{}) bool {
	_, ok := i.(bool)
	return ok
}

// GetBoolValue returns the value of a boolean
func GetBoolValue(i interface{}) (bool, error) {
	b, ok := i.(bool)
	if !ok {
		return false, fmt.Errorf("value is not a boolean")
	}
	return b, nil
}

// GetIntValue returns the value of an integer
func GetIntValue(i interface{}) (int, error) {
	v, ok := i.(int)
	if !ok {
		return 0, fmt.Errorf("value is not an integer")
	}
	return v, nil
}

// GetInt64Value returns the value of an int64
func GetInt64Value(i interface{}) (int64, error) {
	v, ok := i.(int64)
	if !ok {
		return 0, fmt.Errorf("value is not an int64")
	}
	return v, nil
}

// GetFloat64Value returns the value of a float64
func GetFloat64Value(i interface{}) (float64, error) {
	v, ok := i.(float64)
	if !ok {
		return 0, fmt.Errorf("value is not a float64")
	}
	return v, nil
}

// GetStringValue returns the value of a string
func GetStringValue(i interface{}) (string, error) {
	s, ok := i.(string)
	if !ok {
		return "", fmt.Errorf("value is not a string")
	}
	return s, nil
}

// GetStringsFromSlice returns a slice of strings from a given slice
func GetStringsFromSlice(i interface{}) ([]string, error) {
	s, ok := i.([]string)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of strings")
	}
	return s, nil
}

// GetStringsFromSliceOrDefault returns a slice of strings from a given slice or an empty slice if not found
func GetStringsFromSliceOrDefault(i interface{}) ([]string, error) {
	s, err := GetStringsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetIntsFromSlice returns a slice of integers from a given slice
func GetIntsFromSlice(i interface{}) ([]int, error) {
	s, ok := i.([]int)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of integers")
	}
	return s, nil
}

// GetIntsFromSliceOrDefault returns a slice of integers from a given slice or an empty slice if not found
func GetIntsFromSliceOrDefault(i interface{}) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetInt64sFromSlice returns a slice of int64s from a given slice
func GetInt64sFromSlice(i interface{}) ([]int64, error) {
	s, ok := i.([]int64)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of int64s")
	}
	return s, nil
}

// GetInt64sFromSliceOrDefault returns a slice of int64s from a given slice or an empty slice if not found
func GetInt64sFromSliceOrDefault(i interface{}) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetFloat64sFromSlice returns a slice of float64s from a given slice
func GetFloat64sFromSlice(i interface{}) ([]float64, error) {
	s, ok := i.([]float64)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of float64s")
	}
	return s, nil
}

// GetFloat64sFromSliceOrDefault returns a slice of float64s from a given slice or an empty slice if not found
func GetFloat64sFromSliceOrDefault(i interface{}) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBoolsFromSlice returns a slice of booleans from a given slice
func GetBoolsFromSlice(i interface{}) ([]bool, error) {
	s, ok := i.([]bool)
	if !ok {
		return nil, fmt.Errorf("value is not a slice of booleans")
	}
	return s, nil
}

// GetBoolsFromSliceOrDefault returns a slice of booleans from a given slice or an empty slice if not found
func GetBoolsFromSliceOrDefault(i interface{}) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetBoolSliceOrDefault returns a slice of booleans from a given slice or an empty slice if not found
func GetBoolSliceOrDefault(i interface{}, defaultValue bool) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return []bool{defaultValue}, nil
	}
	return s, nil
}

// GetIntSliceOrDefault returns a slice of integers from a given slice or an empty slice if not found
func GetIntSliceOrDefault(i interface{}, defaultValue int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return []int{defaultValue}, nil
	}
	return s, nil
}

// GetInt64SliceOrDefault returns a slice of int64s from a given slice or an empty slice if not found
func GetInt64SliceOrDefault(i interface{}, defaultValue int64) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return []int64{defaultValue}, nil
	}
	return s, nil
}

// GetFloat64SliceOrDefault returns a slice of float64s from a given slice or an empty slice if not found
func GetFloat64SliceOrDefault(i interface{}, defaultValue float64) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return []float64{defaultValue}, nil
	}
	return s, nil
}

// GetStringsOrDefault returns a string from a given string or a default string if not found
func GetStringsOrDefault(i interface{}, defaultValue string) (string, error) {
	s, err := GetStringValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetStringsOrDefault returns a slice of strings from a given slice or a default slice if not found
func GetStringsOrDefault(i interface{}, defaultValue []string) ([]string, error) {
	s, err := GetStringsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetIntOrDefault returns an integer from a given integer or a default integer if not found
func GetIntOrDefault(i interface{}, defaultValue int) (int, error) {
	v, err := GetIntValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetInt64OrDefault returns an int64 from a given int64 or a default int64 if not found
func GetInt64OrDefault(i interface{}, defaultValue int64) (int64, error) {
	v, err := GetInt64Value(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetFloat64OrDefault returns a float64 from a given float64 or a default float64 if not found
func GetFloat64OrDefault(i interface{}, defaultValue float64) (float64, error) {
	v, err := GetFloat64Value(i)
	if err != nil {
		return defaultValue, nil
	}
	return v, nil
}

// GetBoolOrDefault returns a boolean from a given boolean or a default boolean if not found
func GetBoolOrDefault(i interface{}, defaultValue bool) (bool, error) {
	b, err := GetBoolValue(i)
	if err != nil {
		return defaultValue, nil
	}
	return b, nil
}

// GetInt64sOrDefault returns a slice of int64s from a given slice or a default slice if not found
func GetInt64sOrDefault(i interface{}, defaultValue []int64) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetFloat64sOrDefault returns a slice of float64s from a given slice or a default slice if not found
func GetFloat64sOrDefault(i interface{}, defaultValue []float64) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetBoolsOrDefault returns a slice of booleans from a given slice or a default slice if not found
func GetBoolsOrDefault(i interface{}, defaultValue []bool) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetIntsOrDefault returns a slice of integers from a given slice or a default slice if not found
func GetIntsOrDefault(i interface{}, defaultValue []int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return defaultValue, nil
	}
	return s, nil
}

// GetMapOrDefault returns a map from a given map or a default map if not found
func GetMapOrDefault(i interface{}, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	m, ok := i.(map[string]interface{})
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map")
	}
	return m, nil
}

// GetIntSliceOrDefault returns a slice of integers from a given slice or a default slice if not found
func GetIntSliceOrDefault(i interface{}, defaultValue []int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return []int{defaultValue}, nil
	}
	return s, nil
}

// GetInt64SliceOrDefault returns a slice of int64s from a given slice or a default slice if not found
func GetInt64SliceOrDefault(i interface{}, defaultValue []int64) ([]int64, error) {
	s, err := GetInt64sFromSlice(i)
	if err != nil {
		return []int64{defaultValue}, nil
	}
	return s, nil
}

// GetFloat64SliceOrDefault returns a slice of float64s from a given slice or a default slice if not found
func GetFloat64SliceOrDefault(i interface{}, defaultValue []float64) ([]float64, error) {
	s, err := GetFloat64sFromSlice(i)
	if err != nil {
		return []float64{defaultValue}, nil
	}
	return s, nil
}

// GetBoolsOrDefault returns a slice of booleans from a given slice or a default slice if not found
func GetBoolsOrDefault(i interface{}, defaultValue []bool) ([]bool, error) {
	s, err := GetBoolsFromSlice(i)
	if err != nil {
		return []bool{defaultValue}, nil
	}
	return s, nil
}

// GetIntsOrDefault returns a slice of integers from a given slice or a default slice if not found
func GetIntsOrDefault(i interface{}, defaultValue []int) ([]int, error) {
	s, err := GetIntsFromSlice(i)
	if err != nil {
		return []int{defaultValue}, nil
	}
	return s, nil
}

// GetMapOrDefaultString returns a map of strings from a given map or a default map if not found
func GetMapOrDefaultString(i interface{}, defaultValue map[string]string) (map[string]string, error) {
	m, ok := i.(map[string]string)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of strings")
	}
	return m, nil
}

// GetMapOrDefaultInt returns a map of integers from a given map or a default map if not found
func GetMapOrDefaultInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of integers")
	}
	return m, nil
}

// GetMapOrDefaultInt64 returns a map of int64s from a given map or a default map if not found
func GetMapOrDefaultInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of int64s")
	}
	return m, nil
}

// GetMapOrDefaultFloat64 returns a map of float64s from a given map or a default map if not found
func GetMapOrDefaultFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of float64s")
	}
	return m, nil
}

// GetMapOrDefaultBool returns a map of booleans from a given map or a default map if not found
func GetMapOrDefaultBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of booleans")
	}
	return m, nil
}

// GetMapOrDefaultStringSlice returns a map of slices of strings from a given map or a default map if not found
func GetMapOrDefaultStringSlice(i interface{}, defaultValue map[string][]string) (map[string][]string, error) {
	m, ok := i.(map[string][]string)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of strings")
	}
	return m, nil
}

// GetMapOrDefaultIntSlice returns a map of slices of integers from a given map or a default map if not found
func GetMapOrDefaultIntSlice(i interface{}, defaultValue map[string][]int) (map[string][]int, error) {
	m, ok := i.(map[string][]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of integers")
	}
	return m, nil
}

// GetMapOrDefaultInt64Slice returns a map of slices of int64s from a given map or a default map if not found
func GetMapOrDefaultInt64Slice(i interface{}, defaultValue map[string][]int64) (map[string][]int64, error) {
	m, ok := i.(map[string][]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of int64s")
	}
	return m, nil
}

// GetMapOrDefaultFloat64Slice returns a map of slices of float64s from a given map or a default map if not found
func GetMapOrDefaultFloat64Slice(i interface{}, defaultValue map[string][]float64) (map[string][]float64, error) {
	m, ok := i.(map[string][]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of float64s")
	}
	return m, nil
}

// GetMapOrDefaultBoolSlice returns a map of slices of booleans from a given map or a default map if not found
func GetMapOrDefaultBoolSlice(i interface{}, defaultValue map[string][]bool) (map[string][]bool, error) {
	m, ok := i.(map[string][]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of slices of booleans")
	}
	return m, nil
}

// GetMapOrDefaultStringInt returns a map of pairs of strings and integers from a given map or a default map if not found
func GetMapOrDefaultStringInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and integers")
	}
	return m, nil
}

// GetMapOrDefaultStringInt64 returns a map of pairs of strings and int64s from a given map or a default map if not found
func GetMapOrDefaultStringInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and int64s")
	}
	return m, nil
}

// GetMapOrDefaultStringFloat64 returns a map of pairs of strings and float64s from a given map or a default map if not found
func GetMapOrDefaultStringFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and float64s")
	}
	return m, nil
}

// GetMapOrDefaultStringBool returns a map of pairs of strings and booleans from a given map or a default map if not found
func GetMapOrDefaultStringBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of strings and booleans")
	}
	return m, nil
}

// GetMapOrDefaultIntInt64 returns a map of pairs of integers and int64s from a given map or a default map if not found
func GetMapOrDefaultIntInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and int64s")
	}
	return m, nil
}

// GetMapOrDefaultIntFloat64 returns a map of pairs of integers and float64s from a given map or a default map if not found
func GetMapOrDefaultIntFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and float64s")
	}
	return m, nil
}

// GetMapOrDefaultIntBool returns a map of pairs of integers and booleans from a given map or a default map if not found
func GetMapOrDefaultIntBool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of integers and booleans")
	}
	return m, nil
}

// GetMapOrDefaultInt64Float64 returns a map of pairs of int64s and float64s from a given map or a default map if not found
func GetMapOrDefaultInt64Float64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of int64s and float64s")
	}
	return m, nil
}

// GetMapOrDefaultInt64Bool returns a map of pairs of int64s and booleans from a given map or a default map if not found
func GetMapOrDefaultInt64Bool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of int64s and booleans")
	}
	return m, nil
}

// GetMapOrDefaultFloat64Bool returns a map of pairs of float64s and booleans from a given map or a default map if not found
func GetMapOrDefaultFloat64Bool(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, ok := i.(map[string]bool)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of float64s and booleans")
	}
	return m, nil
}

// GetMapOrDefaultBoolInt returns a map of pairs of booleans and integers from a given map or a default map if not found
func GetMapOrDefaultBoolInt(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, ok := i.(map[string]int)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and integers")
	}
	return m, nil
}

// GetMapOrDefaultBoolInt64 returns a map of pairs of booleans and int64s from a given map or a default map if not found
func GetMapOrDefaultBoolInt64(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, ok := i.(map[string]int64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and int64s")
	}
	return m, nil
}

// GetMapOrDefaultBoolFloat64 returns a map of pairs of booleans and float64s from a given map or a default map if not found
func GetMapOrDefaultBoolFloat64(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, ok := i.(map[string]float64)
	if !ok {
		return defaultValue, fmt.Errorf("value is not a map of pairs of booleans and float64s")
	}
	return m, nil
}

// GetMapOrDefaultStringSliceOrDefault returns a map of slices of strings from a given map or a default map if not found
func GetMapOrDefaultStringSliceOrDefault(i interface{}, defaultValue map[string][]string) (map[string][]string, error) {
	m, err := GetMapOrDefaultStringSlice(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultIntSliceOrDefault returns a map of slices of integers from a given map or a default map if not found
func GetMapOrDefaultIntSliceOrDefault(i interface{}, defaultValue map[string][]int) (map[string][]int, error) {
	m, err := GetMapOrDefaultIntSlice(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultInt64SliceOrDefault returns a map of slices of int64s from a given map or a default map if not found
func GetMapOrDefaultInt64SliceOrDefault(i interface{}, defaultValue map[string][]int64) (map[string][]int64, error) {
	m, err := GetMapOrDefaultInt64Slice(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultFloat64SliceOrDefault returns a map of slices of float64s from a given map or a default map if not found
func GetMapOrDefaultFloat64SliceOrDefault(i interface{}, defaultValue map[string][]float64) (map[string][]float64, error) {
	m, err := GetMapOrDefaultFloat64Slice(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultBoolSliceOrDefault returns a map of slices of booleans from a given map or a default map if not found
func GetMapOrDefaultBoolSliceOrDefault(i interface{}, defaultValue map[string][]bool) (map[string][]bool, error) {
	m, err := GetMapOrDefaultBoolSlice(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultStringIntOrDefault returns a map of pairs of strings and integers from a given map or a default map if not found
func GetMapOrDefaultStringIntOrDefault(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, err := GetMapOrDefaultStringInt(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultStringInt64OrDefault returns a map of pairs of strings and int64s from a given map or a default map if not found
func GetMapOrDefaultStringInt64OrDefault(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, err := GetMapOrDefaultStringInt64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultStringFloat64OrDefault returns a map of pairs of strings and float64s from a given map or a default map if not found
func GetMapOrDefaultStringFloat64OrDefault(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, err := GetMapOrDefaultStringFloat64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultStringBoolOrDefault returns a map of pairs of strings and booleans from a given map or a default map if not found
func GetMapOrDefaultStringBoolOrDefault(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, err := GetMapOrDefaultStringBool(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultIntInt64OrDefault returns a map of pairs of integers and int64s from a given map or a default map if not found
func GetMapOrDefaultIntInt64OrDefault(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, err := GetMapOrDefaultIntInt64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultIntFloat64OrDefault returns a map of pairs of integers and float64s from a given map or a default map if not found
func GetMapOrDefaultIntFloat64OrDefault(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, err := GetMapOrDefaultIntFloat64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultIntBoolOrDefault returns a map of pairs of integers and booleans from a given map or a default map if not found
func GetMapOrDefaultIntBoolOrDefault(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, err := GetMapOrDefaultIntBool(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultInt64Float64OrDefault returns a map of pairs of int64s and float64s from a given map or a default map if not found
func GetMapOrDefaultInt64Float64OrDefault(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, err := GetMapOrDefaultInt64Float64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultInt64BoolOrDefault returns a map of pairs of int64s and booleans from a given map or a default map if not found
func GetMapOrDefaultInt64BoolOrDefault(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, err := GetMapOrDefaultInt64Bool(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultFloat64BoolOrDefault returns a map of pairs of float64s and booleans from a given map or a default map if not found
func GetMapOrDefaultFloat64BoolOrDefault(i interface{}, defaultValue map[string]bool) (map[string]bool, error) {
	m, err := GetMapOrDefaultFloat64Bool(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultBoolIntOrDefault returns a map of pairs of booleans and integers from a given map or a default map if not found
func GetMapOrDefaultBoolIntOrDefault(i interface{}, defaultValue map[string]int) (map[string]int, error) {
	m, err := GetMapOrDefaultBoolInt(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultBoolInt64OrDefault returns a map of pairs of booleans and int64s from a given map or a default map if not found
func GetMapOrDefaultBoolInt64OrDefault(i interface{}, defaultValue map[string]int64) (map[string]int64, error) {
	m, err := GetMapOrDefaultBoolInt64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}

// GetMapOrDefaultBoolFloat64OrDefault returns a map of pairs of booleans and float64s from a given map or a default map if not found
func GetMapOrDefaultBoolFloat64OrDefault(i interface{}, defaultValue map[string]float64) (map[string]float64, error) {
	m, err := GetMapOrDefaultBoolFloat64(i, defaultValue)
	if err != nil {
		return defaultValue, err
	}
	return m, nil
}