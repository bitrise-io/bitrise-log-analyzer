package editor

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1475621135, 0),
		Content:     string("<!DOCTYPE html>\n<html>\n<head>\n<!-- css -->\n<style>\nbody {\n    background-color: lightgreen;\n}\n</style>\n<!-- css [end] -->\n</head>\n<body>\n\n<h1>Welcome!</h1>\n\n<p id=\"demo\">Click the \"Send\" button</p>\n<button type=\"button\" onclick=\"sendForTest()\">Send</button>\n\n<!-- javascript -->\n<script>\nfunction sendForTest() {\n  var xhttp;\n  xhttp=new XMLHttpRequest();\n  xhttp.onreadystatechange = function() {\n    if (this.readyState == 4 && this.status == 200) {\n      var respJSON = JSON.parse(this.responseText);\n      console.log(\" -> \", respJSON);\n      document.getElementById(\"demo\").innerHTML = \"Response.message: \" + respJSON.message;\n    }\n  };\n  xhttp.open(\"GET\", \"/api/test-regex\", true);\n  xhttp.send();\n}\n\n</script>\n<!-- javascript [end] -->\n\n</body>\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1475619385, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // index.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`www`, &embedded.EmbeddedBox{
		Name: `www`,
		Time: time.Unix(1475619385, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}
