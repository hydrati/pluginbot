package json

import (
	"bytes"
	"io"
	"strings"

	strip "github.com/hedhyw/jsoncjson"
	. "github.com/hyroge/pluginbot/utils/prelude"
	json "github.com/json-iterator/go"
)

func UnmarshalJsoncString(s string, t Any) error {
	return UnmarshalJsonc(strings.NewReader(s), t)
}

func UnmarshalJsonString(s string, t Any) error {
	return json.UnmarshalFromString(s, t)
}

func UnmarshalJsonBytes(s []byte, t Any) error {
	return json.Unmarshal(s, t)
}

func UnmarshalJsoncBytes(s []byte, t Any) error {
	return UnmarshalJsonc(bytes.NewReader(s), t)
}

func UnmarshalJson(r io.Reader, t Any) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(t)
}

func UnmarshalJsonc(r io.Reader, t Any) error {
	striper := CreateJsonStriper(r)
	return UnmarshalJson(striper, t)
}

func MarshalJson(r Any) ([]byte, error) {
	return json.Marshal(r)
}

func MarshalJsonToString(r Any) (string, error) {
	return json.MarshalToString(r)
}

func MarshalJsonToWriter(r Any, w io.Writer) error {
	return json.NewEncoder(w).Encode(r)
}

func CreateJsonStriper(r io.Reader) io.Reader {
	return strip.NewReader(r)
}
