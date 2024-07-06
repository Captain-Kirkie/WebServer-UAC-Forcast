package greetings

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"uacforcast/WebServer/services"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestParseResponse(t *testing.T) {
	t.Logf("parsing response from uac")
    assert := assert.New(t)
	
	
	data, err := os.ReadFile("./test_response.json")
	t.Log(len(data))
	
	assert.Nil(err)
	
	var test services.Advisories

	errorTest := json.Unmarshal(data, &test)

	assert.Nil(errorTest)

	// TODO: deserialize this into an int[]
	split := strings.Split(test.Advisories[0].Advisory.OverallDangerRose, ",")
	assert.Equal(len(split), 24)
}
