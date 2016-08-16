package update

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ()

var (
	gopath string
	buffe  bytes.Buffer
	depth  int
)

func Update() {
	gopath := os.Getenv("gopath")
	gopath = strings.Replace(gopath, "\\", "/", -1)
	buffe.WriteString(gopath)
	buffe.WriteString("/src/")
	depth = len(strings.Split(buffe.String(), "/"))
	filepath.Walk(buffe.String(), Iterator)
	buffe.Reset()
}
func Iterator(path string, info os.FileInfo, err error) error {
	if len(strings.Split(path, "/")) == depth {
		return nil
	}
	path = strings.Replace(path, "\\", "/", -1)
	if len(strings.Split(path, "/")) > depth && info.IsDir() && !strings.Contains(path, ".git") {
		url := strings.Replace(path, buffe.String(), "", -1)
		//		go func(l string) {
		cmd := exec.Command("go", "get", "-u", "-v", url)
		_, e := cmd.CombinedOutput()
		if e != nil {
			fmt.Printf("#######  operation failure : %s", e.Error())
			//return errors.New("发现错误")
		} else {
			fmt.Println("not prefix  ", url)
		}

		//		}(url)
	}
	return nil
}
