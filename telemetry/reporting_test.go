package telemetry

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test if the loop is working by asserting whether the log is reset after <interval> seconds
func TestLoop(t *testing.T) {
	t.Parallel()

	// setup
	calledFunctions := NewLog()
	calledFunctions.Enabled = true
	calledFunctions.PeerName = "soggy-whale-bread"
	calledFunctions.Register("REST", "weaviate.something.or.other")

	interval := 1
	url := "http://webhook.site/73641e3c-6d28-4875-aa5e-b0e66abd3b00"

	reporter := NewReporter(calledFunctions, interval, url, true, true)

	go reporter.Start()

	time.Sleep(time.Duration(2) * time.Second)

	logsAfterReporting := calledFunctions.ExtractLoggedRequests()

	// test
	assert.Equal(t, 0, len(*logsAfterReporting))
}

// Register a single request and call the JSON minimization function, then assert whether all fields have been minimized correctly
func TestConvertToMinimizedJSON(t *testing.T) {
	t.Parallel()

	// setup
	calledFunctions := NewLog()
	calledFunctions.Enabled = true
	calledFunctions.PeerName = "tiny-grey-chainsword"
	calledFunctions.Register("REST", "weaviate.something.or.other")

	calledFunctions.Log["[REST]weaviate.something.or.other"].Name = "tiny-grey-chainsword"
	calledFunctions.Log["[REST]weaviate.something.or.other"].When = int64(1550745544)

	transformer := NewOutputTransformer(true)

	jsonLogs := transformer.ConvertToMinimizedJSON(&calledFunctions.Log)

	var miniLog []interface{}

	err := json.Unmarshal([]byte(*jsonLogs), &miniLog)

	if err != nil {
		fmt.Println("error:", err)
	}

	// test
	assert.Equal(t, 1, len(miniLog))

	for _, log := range miniLog {
		mapLog := log.(map[string]interface{})

		name := mapLog["n"].(string)
		logType := mapLog["t"].(string)
		identifier := mapLog["i"].(string)
		amount := mapLog["a"].(float64)
		when := mapLog["w"].(float64)

		assert.Equal(t, "tiny-grey-chainsword", name)
		assert.Equal(t, "REST", logType)
		assert.Equal(t, "weaviate.something.or.other", identifier)
		assert.Equal(t, 1, int(amount))
		assert.Equal(t, int64(1550745544), int64(when))
	}
}

// Create a log containing two logged function types and call the AddTimeStamps function,
// then assert whether both logs now contain the same plausible timestamp.
func TestAddTimestamps(t *testing.T) {
	t.Parallel()

	// setup
	calledFunctions := NewLog()
	calledFunctions.Enabled = true
	calledFunctions.PeerName = "iridiscent-damp-bagel"
	calledFunctions.Register("REST", "weaviate.something.or.other1")
	calledFunctions.Register("REST", "weaviate.something.or.other2")

	past := time.Now().Unix()

	time.Sleep(time.Duration(1) * time.Second)

	reporter := NewReporter(nil, 0, "", true, true)
	reporter.AddTimeStamps(&calledFunctions.Log)

	time.Sleep(time.Duration(1) * time.Second)

	future := time.Now().Unix()

	log1Timestamp := calledFunctions.Log["[REST]weaviate.something.or.other1"].When
	log2Timestamp := calledFunctions.Log["[REST]weaviate.something.or.other2"].When

	notZero := log1Timestamp != 0
	notFromThePast := log1Timestamp > past
	notFromTheFuture := log1Timestamp < future

	// test
	assert.Equal(t, log1Timestamp, log2Timestamp)
	assert.Equal(t, true, notZero)
	assert.Equal(t, true, notFromThePast)
	assert.Equal(t, true, notFromTheFuture)
}

// Create a minimized request log and encode it to CBOR then assert whether the result matches the expected value.
func TestCborEncode(t *testing.T) {
	t.Parallel()

	// setup
	minimizedLog := `[{"n": "upbeat-aquatic-pen", "t": "REST", "i": "weaviate.something.or.other", "a": 1, "w": 1550745544}]`

	outputTransformer := NewOutputTransformer(true)

	encoded, err := outputTransformer.EncodeAsCBOR(&minimizedLog)

	expected := []byte{0x78, 0x67, 0x5b, 0x7b, 0x22, 0x6e, 0x22, 0x3a, 0x20, 0x22, 0x75, 0x70, 0x62, 0x65, 0x61, 0x74, 0x2d, 0x61, 0x71, 0x75, 0x61, 0x74, 0x69, 0x63, 0x2d, 0x70, 0x65, 0x6e, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x52, 0x45, 0x53, 0x54, 0x22, 0x2c, 0x20, 0x22, 0x69, 0x22, 0x3a, 0x20, 0x22, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x2c, 0x20, 0x22, 0x61, 0x22, 0x3a, 0x20, 0x31, 0x2c, 0x20, 0x22, 0x77, 0x22, 0x3a, 0x20, 0x31, 0x35, 0x35, 0x30, 0x37, 0x34, 0x35, 0x35, 0x34, 0x34, 0x7d, 0x5d}

	assert.Equal(t, nil, err)

	assert.Equal(t, expected, *encoded)
}

// func TestReporting(t *testing.T) {
// 	t.Parallel()

// 	URL := "localhost:8087/new"

// 	minimizedLog := `[{"n": "upbeat-aquatic-pen", "t": "REST", "i": "weaviate.something.or.other", "a": 1, "w": 1550745544}]`

// 	outputTransformer := NewOutputTransformer(true)

// 	encoded, _ := outputTransformer.EncodeAsCBOR(&minimizedLog)

// 	poster := NewPoster(URL)
// 	poster.ReportLoggedCalls(encoded)

// 	// setup
// 	// a := []byte{165, 97, 120, 10, 97, 121, 15, 97, 122, 24, 100, 101, 114, 97, 110, 103, 101, 132, 162, 102, 108, 101, 110, 103, 116, 104, 1, 101, 97, 108, 105, 103, 110, 250, 65, 32, 0, 0, 162, 102, 108, 101, 110, 103, 116, 104, 26, 13, 81, 78, 231, 101, 97, 108, 105, 103, 110, 250, 65, 240, 0, 0, 162, 102, 108, 101, 110, 103, 116, 104, 3, 101, 97, 108, 105, 103, 110, 250, 66, 38, 0, 0, 162, 102, 108, 101, 110, 103, 116, 104, 24, 174, 101, 97, 108, 105, 103, 110, 250, 71, 89, 3, 51, 101, 108, 97, 98, 101, 108, 102, 72, 111, 72, 111, 72, 111}
// 	// b := bytes.NewReader(a)
// 	// fmt.Println(b)
// 	// poster := NewPoster("http://webhook.site/73641e3c-6d28-4875-aa5e-b0e66abd3b00")
// 	// poster.ReportLoggedCalls(a)

// 	// test
// 	assert.Equal(t, 1, 1)
// }
