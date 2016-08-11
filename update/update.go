package update

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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
		fmt.Println(strings.Replace(path, buffe.String(), "", -1))
		cmd := exec.Command("go", "get", "-u", "-v", strings.Replace(path, buffe.String(), "", -1))
		ff, e := cmd.CombinedOutput()
		if e != nil {
			fmt.Printf("  operation failure : %s", e.Error())
			//return errors.New("发现错误")
		} else {
			fmt.Println(" Successful operation " + string(ff))
		}
	}
	return nil
}
