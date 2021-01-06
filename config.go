package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/hashicorp/go-multierror"
)

func FillEnvTags(confStructPtr interface{}) error {
	value := reflect.ValueOf(confStructPtr)
	if value.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	}
	if value.IsNil() {
		return errors.New("must pass a non nil pointer")
	}
	if value.Type().Elem().Kind() != reflect.Struct {
		return errors.New("must pass a pointer to a struct")
	}

	// use error count and wrapped errors instead of failing early
	// so we get a full list of problems/missing env var in output
	errorCount := 0
	var errorChain error

	numberKindBitSize := map[reflect.Kind]int {
		reflect.Int:    0, // auto
		reflect.Uint:   0, // auto
		reflect.Int8:   8,
		reflect.Uint8:  8,
		reflect.Int16:  16,
		reflect.Uint16: 16,
		reflect.Int32:  32,
		reflect.Uint32: 32,
		reflect.Int64:  64,
		reflect.Uint64: 64,
	}

	isTruthy := func(str string) bool {
		return str != "" && str != "0" && str != "false" && str != "FALSE"
	}

	direct := value.Elem() // already checked its of kind ptr
	directType := direct.Type()

	for i := 0; i < direct.NumField(); i++ {
		tag := directType.Field(i).Tag
		tagEnvName := tag.Get("env")
		if tagEnvName == "" {
			continue
		}

		f := direct.Field(i)
		if !f.CanSet() {
			errorCount++
			errorChain = multierror.Append(errorChain, fmt.Errorf("unable to set %s of type %s", directType.Field(i).Name, f.Kind(),))
			continue
		}

		envRequiredVal := tag.Get("required")
		envVar, envVarSet := os.LookupEnv(tagEnvName)

		if !envVarSet && isTruthy(envRequiredVal) {
			errorCount++
			errorChain = multierror.Append(errorChain, fmt.Errorf("unable to set %s of type %s, required environment variable '%s' missing",  directType.Field(i).Name, f.Kind(), tagEnvName))
			continue
		}

		switch f.Kind() {
		case reflect.String:
			f.SetString(envVar)
		case reflect.Bool:
			f.SetBool(isTruthy(envVar))
		case reflect.Int8: fallthrough
		case reflect.Int16: fallthrough
		case reflect.Int32: fallthrough
		case reflect.Int64: fallthrough
		case reflect.Int:
			asInt, asIntErr := strconv.ParseInt(envVar, 10, numberKindBitSize[f.Kind()])
			if asIntErr != nil {
				errorCount++
				errorChain = multierror.Append(errorChain, fmt.Errorf("unable to set %s of type %s from %s: %w", directType.Field(i).Name, f.Kind(), tagEnvName, asIntErr))
				continue
			}
			f.SetInt(asInt)
		case reflect.Uint8: fallthrough
		case reflect.Uint16: fallthrough
		case reflect.Uint32: fallthrough
		case reflect.Uint64: fallthrough
		case reflect.Uint:
			asUint, asUintErr := strconv.ParseUint(envVar, 10, numberKindBitSize[f.Kind()])
			if asUintErr != nil {
				errorCount++
				errorChain = multierror.Append(errorChain, fmt.Errorf("unable to set %s of type %s from %s: %w", directType.Field(i).Name, f.Kind(), tagEnvName, asUintErr))
				continue
			}
			f.SetUint(asUint)
		default:
			errorCount++
			errorChain = multierror.Append(errorChain, fmt.Errorf("unable to set %s of type %s (supported)",  directType.Field(i).Name, f.Kind()))
			continue
		}
	}

	return errorChain
}


type Config struct {
	Debug           bool   `env:"DEBUG"`
	Environment     string `env:"APP_ENV" required:"true"`
	DbConnectionDSN string `env:"DATABASE_DSN" required:"true"`
}

func main()  {
	config := Config{}
	confErr := FillEnvTags(&config)
	if  confErr != nil {
		log.Fatalf("unable to make config: %v", confErr)
	}

	fmt.Printf("successfuly made config: %+v", config)
}
