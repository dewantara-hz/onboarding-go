package reqres

import "testing"

func TestAccess(t *testing.T) {
	cases := []struct {
		url       string
		checkType string
		validity  bool
	}{
		{"https://jsonplaceholder.typicode.com/todos/1", "json", true},
		{"https://www.w3schools.com/xml/note.xml", "xml", true},
		{"http://google.com/", "html", true},
	}

	for _, _case := range cases {
		content := Access(_case.url)
		checkType := _case.checkType
		validity := _case.validity

		if checkType == "json" {
			if IsJson(content) != validity {
				t.Errorf("error")
			}
		} else if checkType == "xml" {
			if IsXML(content) != validity {
				t.Errorf("error")
			}
		} else if checkType == "html" {
			if IsHTML(content) != validity {
				t.Errorf("error")
			}
		}

	}
}
