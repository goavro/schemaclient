package schemaclient

import (
	"reflect"
	"testing"

	avro "github.com/elodina/go-avro"
)

func TestSchemaRegistry(t *testing.T) {
	client := NewCachedSchemaRegistryClient("http://localhost:8081")
	rawSchema := "{\"namespace\": \"metrics\",\"type\": \"record\",\"name\": \"Timings\",\"fields\": [{\"name\": \"id\", \"type\": \"long\"},{\"name\": \"timings\",  \"type\": {\"type\":\"array\", \"items\": \"long\"} }]}"
	schema, err := avro.ParseSchema(rawSchema)
	assert(t, err, nil)
	id, err := client.Register("test1", schema)
	assert(t, err, nil)
	assertNot(t, id, 0)

	schema, err = client.GetByID(id)
	assert(t, err, nil)
	assertNot(t, schema, nil)

	metadata, err := client.GetLatestSchemaMetadata("test1")
	assert(t, err, nil)
	assertNot(t, metadata, nil)

	version, err := client.GetVersion("test1", schema)
	assert(t, err, nil)
	assertNot(t, version, 0)
}

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func assertNot(t *testing.T, actual interface{}, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("%v should not be %v", actual, expected)
	}
}
