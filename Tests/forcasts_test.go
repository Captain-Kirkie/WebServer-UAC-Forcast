package greetings

import (
	"encoding/json"
	"io/ioutil"
	"os"
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


	for i := 0; i < len(test.Advisories); i++ {
    	t.Logf("%v", test.Advisories[i].Advisory.DangerRose1)
	}
}

// just for testing
type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
}

func TestParsingPractice(t *testing.T) {
    assert := assert.New(t)

	// Open our jsonFile
	jsonFile, err := os.Open("usertest.json")
	// if we os.Open returns an error then handle it
	assert.Nil(err)
	
	t.Log("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	byteValue, _ := ioutil.ReadAll(jsonFile)

// we initialize our Users array
var users Users

// we unmarshal our byteArray which contains our
// jsonFile's content into 'users' which we defined above
json.Unmarshal(byteValue, &users)

// we iterate through every user within our users array and
// print out the user Type, their name, and their facebook url
// as just an example
for i := 0; i < len(users.Users); i++ {
    t.Logf("User Type: " + users.Users[i].Type)
    // t.Logf("User Age: " + strconv.Itoa(users.Users[i].Age))
    t.Logf("User Name: " + users.Users[i].Name)
}
		
}