package reqres

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func IsJson(content string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(content), &js) == nil
}

func IsXML(content string) bool {
	// var _xml map[string]interface{}
	var _xml string
	return xml.Unmarshal([]byte(content), &_xml) == nil
}

func IsHTML(content string) bool {
	if strings.Contains(content, "</html>") {
		return !IsJson(content) && !IsXML(content)
	} else {
		return false
	}
}
