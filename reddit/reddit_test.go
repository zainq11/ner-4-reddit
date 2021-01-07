package reddit

import (
	"fmt"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

// Call parseContent with a file reader and checks for a valid return value
func TestParseContent(t *testing.T) {
	f, err := os.Open("resources/reddit_json_test.txt")
	check(err)

	response, respErr := parseContent(f)
	if respErr != nil {
		t.Fatalf("Error during parsing content: %v", respErr)
	}

	if response.Data.Children == nil || len(response.Data.Children) == 0 {
		t.Fatalf("No reddit content: %v", respErr)
	}
}

func TestMustBuildReadUrl(t *testing.T) {
	url := mustBuildReadUrl("r/postrock")
	//fmt.Println(url)
	assertEqual(t, url, "https://oauth.reddit.com/r/postrock/new.json", "Did not match")
}
