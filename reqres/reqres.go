package reqres

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Access the given url and return the body content
func Access(url string) string {

	// Access the url
	resp, err := http.Get(url)
	bodyString := ""
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error url : ", url)
		return ""
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}

	return bodyString
}

// IsJSON check whether the given content is JSON
func IsJSON(content string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(content), &js) == nil
}

// IsXML check whether the given content is XML
func IsXML(content string) bool {
	// var _xml map[string]interface{}
	var _xml string
	return xml.Unmarshal([]byte(content), &_xml) == nil
}

// IsHTML check whether the given content is HTML
func IsHTML(content string) bool {

	validity := false

	if strings.Contains(content, "</html>") {
		validity = !IsJSON(content) && !IsXML(content)
	}

	return validity
}
