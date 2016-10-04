package editor

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1475618618, 0),
		Content:     string("<html>\n<head>\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"main.css\">\n</head>\n<body>\n    Welcome!\n</body>\n</html>"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `main.css`,
		FileModTime: time.Unix(1475618678, 0),
		Content:     string("body {\n    background-color: lightgreen;\n}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1475618522, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // index.html
			file3, // main.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`www`, &embedded.EmbeddedBox{
		Name: `www`,
		Time: time.Unix(1475618522, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
			"main.css":   file3,
		},
	})
}
