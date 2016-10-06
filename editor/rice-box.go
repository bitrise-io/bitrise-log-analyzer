package editor

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1475694875, 0),
		Content:     string("<!DOCTYPE html>\n<html>\n<head>\n<!-- css -->\n<style>\nbody {\n    background-color: lightgreen;\n}\n</style>\n<!-- css [end] -->\n</head>\n<body>\n\n<h1>Welcome!</h1>\n\n<textarea id=\"log-input\" placeholder=\"drop your log in there\" style=\"width: 100%\"></textarea>\n<input id=\"pattern\" type=\"text\" placeholder=\"pattern\"></input>\n<button id=\"submit\" type=\"button\" onclick=\"sendForTest()\">Send</button>\n\n<h2>Results:</h2>\n<ul id=\"results-list\">\n</ul>\n\n<!-- javascript -->\n<script>\ndocument.getElementById(\"pattern\").addEventListener(\"keyup\", function(event) {\n    event.preventDefault();\n    if (event.keyCode == 13) {\n        document.getElementById(\"submit\").click();\n    }\n});\n\nfunction sendForTest() {\n  var xhttp;\n  xhttp = new XMLHttpRequest();\n  xhttp.onreadystatechange = function() {\n    if (this.readyState == 4 && this.status == 200) {\n      var respJSON = JSON.parse(this.responseText);\n      console.log(\" -> respJSON:\", respJSON);\n\n      // results\n      var resultsListElem = document.getElementById(\"results-list\")\n      // clear out results\n      while (resultsListElem.hasChildNodes()) {\n        resultsListElem.removeChild(resultsListElem.firstChild);\n      }\n      // add new ones\n      var matches = respJSON.matches;\n      if (!matches) {\n        matches = [\"NO MATCH!\"]\n      }\n      for (var i = 0; i < matches.length; i++) {\n        var aMatch = matches[i]\n        var liElem = document.createElement(\"LI\");\n        liElem.innerHTML = aMatch\n        resultsListElem.appendChild(liElem);\n      }\n    }\n  };\n  xhttp.open(\"POST\", \"/api/test-regex\", true);\n  xhttp.send(JSON.stringify({\n    log: document.getElementById(\"log-input\").value,\n    pattern: document.getElementById(\"pattern\").value\n  }));\n}\n\n</script>\n<!-- javascript [end] -->\n\n</body>\n</html>"),
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
