package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	Debug           bool   `env:"DEBUG"`
	Environment     string `env:"APP_ENV" required:"true"`
	DbConnectionDSN string `env:"DATABASE_DSN" required:"true"`
}

func MakeConfigFromEnv() (Config, error) {
	config := Config{}
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(&config).Elem()

	// use error count instead of failing early so we get a full list of problems/missing env var in output
	errorCount := 0
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		if !f.CanSet() {
			errorCount++
			log.Printf("config: unable to set %s", t.Field(i).Name)
			continue
		}

		tag := t.Field(i).Tag.Get("env")
		required := t.Field(i).Tag.Get("required")
		envVar, envVarSet := os.LookupEnv(tag)

		if (required == "true" || required == "1") && !envVarSet {
			errorCount++
			log.Printf("config: missing environment variable %s", tag)
			continue
		}

		switch f.Kind() {
		case reflect.String:
			f.SetString(envVar)
		case reflect.Bool:
			f.SetBool(envVar == "true" || envVar == "1")
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			asInt, asIntErr := strconv.ParseInt(envVar, 10, 0)
			if asIntErr != nil {
				errorCount++
				log.Printf("config: unable to convert %s to int64", envVar)
				continue
			}
			f.SetInt(asInt)
		default:
			errorCount++
			log.Printf("config: unable to set %s of type %s", t.Field(i).Name, f.Kind())
			continue
		}
	}

	if errorCount > 0 {
		return Config{}, fmt.Errorf("missing or unable to set %d config values from environment variables", errorCount)
	}

	return config, nil
}


func main()  {
	config, configErr := MakeConfigFromEnv()
	if  configErr != nil {
		log.Fatalf("unable to make config: %v", configErr)
	}

	fmt.Printf("successfuly made config: %+v", config)
}
