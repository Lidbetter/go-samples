package main_test

import (
	"os"
	"testing"

	main "github.com/Lidbetter/go-samples"
)

func TestFillEnvTagsSimple(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117"`
		FooBool   bool   `env:"TEST_ENV_FOO_117"`
		FooInt    int    `env:"TEST_ENV_FOO_117"`
		FooInt8   int8   `env:"TEST_ENV_FOO_117"`
		FooInt16  int16  `env:"TEST_ENV_FOO_117"`
		FooInt32  int32  `env:"TEST_ENV_FOO_117"`
		FooInt64  int64  `env:"TEST_ENV_FOO_117"`
		FooUint   uint   `env:"TEST_ENV_FOO_117"`
		FooUint8  uint8  `env:"TEST_ENV_FOO_117"`
		FooUint16 uint16 `env:"TEST_ENV_FOO_117"`
		FooUint32 uint32 `env:"TEST_ENV_FOO_117"`
		FooUint64 uint64 `env:"TEST_ENV_FOO_117"`
	}{}

	err := main.FillEnvTags(&conf)
	if err != nil {
		t.Fatalf("error filling conf struct: %v", err)
	}

	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
	if conf.FooBool != true {
		t.Fatalf("FooBool expected 'true', actual '%t'", conf.FooBool)
	}
	if conf.FooInt != int(117) {
		t.Fatalf("FooInt expected '117', actual '%d'", conf.FooInt)
	}
	if conf.FooInt8 != int8(117) {
		t.Fatalf("FooInt8 expected '117', actual '%d'", conf.FooInt8)
	}
	if conf.FooInt16 != int16(117) {
		t.Fatalf("FooInt16 expected '117', actual '%d'", conf.FooInt16)
	}
	if conf.FooInt32 != int32(117) {
		t.Fatalf("FooInt32 expected '117', actual '%d'", conf.FooInt32)
	}
	if conf.FooInt64 != int64(117) {
		t.Fatalf("FooInt64 expected '117', actual '%d'", conf.FooInt64)
	}
	if conf.FooUint != uint(117) {
		t.Fatalf("FooUint expected '117', actual '%d'", conf.FooUint)
	}
	if conf.FooUint8 != uint8(117) {
		t.Fatalf("FooUint8 expected '117', actual '%d'", conf.FooUint8)
	}
	if conf.FooUint16 != uint16(117) {
		t.Fatalf("FooUint16 expected '117', actual '%d'", conf.FooUint16)
	}
	if conf.FooUint32 != uint32(117) {
		t.Fatalf("FooUint32 expected '117', actual '%d'", conf.FooUint32)
	}
	if conf.FooUint64 != uint64(117) {
		t.Fatalf("FooUint64 expected '117', actual '%d'", conf.FooUint64)
	}
}

func TestFillEnvTagsRequired(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117" required:"true"`
		FooStringRequired string `env:"TEST_ENV_FOO_117______MISSING" required:"true"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when an environment variable is marked as required, but is not set")
	}

	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
}

func TestFillEnvTagsRequiredFalse(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117"`
		FooStringNotRequired1 string `env:"TEST_ENV_FOO_117______MISSING"`
		FooStringNotRequired2 string `env:"TEST_ENV_FOO_117______MISSING" required:""`
		FooStringNotRequired3 string `env:"TEST_ENV_FOO_117______MISSING" required:"0"`
		FooStringNotRequired4 string `env:"TEST_ENV_FOO_117______MISSING" required:"false"`
		FooStringNotRequired5 string `env:"TEST_ENV_FOO_117______MISSING" required:"FALSE"`
	}{}

	err := main.FillEnvTags(&conf)
	if err != nil {
		t.Fatalf("error filling conf struct: %v", err)
	}
	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
}

func TestFillEnvTagsOverridesExistingValue(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117"`
	}{
		FooString: "existing value",
	}

	err := main.FillEnvTags(&conf)
	if err != nil {
		t.Fatalf("error filling conf struct: %v", err)
	}

	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
}

func TestFillEnvTagsBool(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	_ = os.Setenv("TEST_ENV_FOO_0", "0")
	_ = os.Setenv("TEST_ENV_FOO_EMPTY", "")
	_ = os.Setenv("TEST_ENV_FOO_FALSE", "false")
	_ = os.Setenv("TEST_ENV_FOO_FALSE2", "FALSE")
	conf := struct {
		FooBoolTrue bool `env:"TEST_ENV_FOO_117"`
		FooBoolFalse bool `env:"TEST_ENV_FOO_117______MISSING"`
		FooBoolFalse1 bool `env:"TEST_ENV_FOO_0"`
		FooBoolFalse2 bool `env:"TEST_ENV_FOO_EMPTY"`
		FooBoolFalse3 bool `env:"TEST_ENV_FOO_FALSE"`
		FooBoolFalse4 bool `env:"TEST_ENV_FOO_FALSE2"`
	}{}

	err := main.FillEnvTags(&conf)
	if err != nil {
		t.Fatalf("error filling conf struct: %v", err)
	}

	if conf.FooBoolTrue != true {
		t.Fatalf("FooBoolTrue expected 'true', actual '%t'", conf.FooBoolTrue)
	}
	if conf.FooBoolFalse != false {
		t.Fatalf("FooBoolFalse expected 'false', actual '%t'", conf.FooBoolFalse)
	}
	if conf.FooBoolFalse1 != false {
		t.Fatalf("FooBoolFalse1 expected 'false', actual '%t'", conf.FooBoolFalse1)
	}
	if conf.FooBoolFalse2 != false {
		t.Fatalf("FooBoolFalse2 expected 'false', actual '%t'", conf.FooBoolFalse2)
	}
	if conf.FooBoolFalse3 != false {
		t.Fatalf("FooBoolFalse3 expected 'false', actual '%t'", conf.FooBoolFalse3)
	}
	if conf.FooBoolFalse4 != false {
		t.Fatalf("FooBoolFalse4 expected 'false', actual '%t'", conf.FooBoolFalse4)
	}
}

func TestFillEnvTagsOverflow8(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_1177", "1177")
	conf := struct {
		FooInt8   int8   `env:"TEST_ENV_FOO_1177"`
	}{}
	conf2 := struct {
		FooUint8  uint8  `env:"TEST_ENV_FOO_1177"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when setting int8 to 1177 (overflow)")
	}
	err2 := main.FillEnvTags(&conf2)
	if err2 == nil {
		t.Fatalf("conf struct should error when setting uint8 to 1177 (overflow)")
	}
}

func TestFillEnvTagsOverflow16(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117766", "117766")
	conf := struct {
		FooInt16   int16   `env:"TEST_ENV_FOO_117766"`
	}{}
	conf2 := struct {
		FooUint16  uint16 `env:"TEST_ENV_FOO_117766"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when setting int16 to 117766 (overflow)")
	}
	err2 := main.FillEnvTags(&conf2)
	if err2 == nil {
		t.Fatalf("conf struct should error when setting uint16 to 117766 (overflow)")
	}
}

func TestFillEnvTagsOverflow32(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117766554433", "117766554433")
	conf := struct {
		FooInt32   int32   `env:"TEST_ENV_FOO_117766554433"`
	}{}
	conf2 := struct {
		FooUint32  uint32 `env:"TEST_ENV_FOO_117766554433"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when setting int32 to 117766554433 (overflow)")
	}
	err2 := main.FillEnvTags(&conf2)
	if err2 == nil {
		t.Fatalf("conf struct should error when setting uint32 to 117766554433 (overflow)")
	}
}

func TestFillEnvTagsOverflow64(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_1177665544332211001122", "1177665544332211001122")
	conf := struct {
		FooInt64   int64   `env:"TEST_ENV_FOO_1177665544332211001122"`
	}{}
	conf2 := struct {
		FooUint64  uint64 `env:"TEST_ENV_FOO_1177665544332211001122"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when setting int64 to 1177665544332211001122 (overflow)")
	}
	err2 := main.FillEnvTags(&conf2)
	if err2 == nil {
		t.Fatalf("conf struct should error when setting uint32 to 1177665544332211001122 (overflow)")
	}
}

func TestFillEnvTagsIgnoresOtherFields(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117"`
		Bar       string
		baz       string
	}{}

	err := main.FillEnvTags(&conf)
	if err != nil {
		t.Fatalf("error filling conf struct: %v", err)
	}

	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
	if conf.Bar != "" {
		t.Fatalf("FooString expected '', actual '%s'", conf.Bar)
	}
}

func TestFillEnvTagsErrorsOnUnexportedFields(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string `env:"TEST_ENV_FOO_117"`
		foo       string `env:"TEST_ENV_FOO_117"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when unexpored field has env tag")
	}
	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
}


func TestFillEnvTagsErrorsOnUnsupportedTypes(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct {
		FooString string   `env:"TEST_ENV_FOO_117"`
		FooStruct struct{} `env:"TEST_ENV_FOO_117"`
	}{}

	err := main.FillEnvTags(&conf)
	if err == nil {
		t.Fatalf("conf struct should error when unsupported field has env tag")
	}
	if conf.FooString != "117" {
		t.Fatalf("FooString expected '117', actual '%s'", conf.FooString)
	}
}


func TestFillEnvTagsHandlesBadInput(t *testing.T) {
	_ = os.Setenv("TEST_ENV_FOO_117", "117")
	conf := struct{}{}
	err := main.FillEnvTags(conf)
	if err == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}

	// nil ptr
	var conf1 *[]byte = nil
	err1 := main.FillEnvTags(conf1)
	if err1 == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}

	conf2 := int(1)
	err2 := main.FillEnvTags(conf2)
	if err2 == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}
	err3 := main.FillEnvTags(&conf2)
	if err3 == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}

	var conf4 error = nil
	err4 := main.FillEnvTags(conf4)
	if err4 == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}
	err5 := main.FillEnvTags(&conf4)
	if err5 == nil {
		t.Fatalf("FillEnvTags should error when passed in non struct ptr")
	}
}
