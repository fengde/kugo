package command

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"unicode/utf8"

	"github.com/fengde/gocommon/filex"
	"github.com/spf13/cast"
)

type NewCommand struct{}

func (self NewCommand) Cmd() string {
	return "new"
}

func (self NewCommand) Exec(args ...string) error {
	num := len(args)
	if num == 0 {
		return fmt.Errorf("args error")
	}

	var templ string
	var projname string

	switch num {
	case 1:
		templ = "normal"
		projname = args[0]
	case 2:
		if _, ok := map[string]int{
			"http": 1,
			"grpc": 1,
			"cli":  1,
		}[args[0]]; !ok {
			return fmt.Errorf("args error")
		}

		templ = args[0]
		projname = args[1]
	}

	if t := cast.ToInt(projname[0]); t < 97 || t > 122 {
		return fmt.Errorf("please change the project name '%s'. it's a bad name, must start with lower char.", projname)
	}

	if err := filex.Copy(fmt.Sprintf("./template/new-%s/{{template}}", templ), projname); err != nil {
		return err
	}

	if err := self.replace(projname, projname); err != nil {
		return err
	}

	return nil
}

func (self NewCommand) Help() string {
	return `kugo new:
	kugo new {project}	-- create go demo project.
	kugo new http {project} -- create go http project which based with gozero.
	kugo new grpc {project} -- create go grpc project which based with gozero.
	kugo new cli {project} -- create go cli project.`
}

func (self NewCommand) replace(dir string, projname string) error {
	var holders = map[string]string{
		"{{template}}": projname,
		"{{Template}}": strings.ToUpper(projname)[:1] + projname[1:],
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		name := entry.Name()
		var newName string
		for holder, pn := range holders {
			if strings.Contains(name, holder) {
				newName = strings.ReplaceAll(name, holder, pn)
			}
		}
		if newName != "" {
			os.Rename(path.Join(dir, name), path.Join(dir, newName))
			name = newName
		}

		filepath := path.Join(dir, name)
		if entry.IsDir() {
			if err := self.replace(filepath, projname); err != nil {
				return err
			}
			continue
		}
		// 替换文本文件占位符内容
		if yes, _ := isTextFile(filepath); !yes {
			fmt.Println(filepath)
			continue
		}

		content, err := filex.ReadFileToString(filepath)
		if err != nil {
			return err
		}

		for holder, pn := range holders {
			content = strings.ReplaceAll(content, holder, pn)
		}

		if err := filex.WriteStringToFile(filepath, content, false); err != nil {
			return err
		}
	}

	return nil
}

func isTextFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 512) // 读取文件的前512字节
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return false, err
	}

	isText := isTextContent(buffer)

	return isText, nil
}

func isTextContent(buffer []byte) bool {
	if isUTF8(buffer) {
		return true
	}

	// 检查常见的非文本文件标志
	nonTextSignatures := [][]byte{
		{0xFF, 0xD8, 0xFF}, // JPEG
		{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG
		{0x47, 0x49, 0x46, 0x38, 0x37, 0x61},             // GIF87a
		{0x47, 0x49, 0x46, 0x38, 0x39, 0x61},             // GIF89a
		{0x25, 0x50, 0x44, 0x46},                         // PDF
	}

	for _, signature := range nonTextSignatures {
		if bytesHasPrefix(buffer, signature) {
			return false
		}
	}

	return true
}

func isUTF8(buffer []byte) bool {
	utf8Bytes := []byte{0xEF, 0xBB, 0xBF}
	return bytesHasPrefix(buffer, utf8Bytes) || utf8.Valid(buffer)
}

func bytesHasPrefix(buffer, prefix []byte) bool {
	return len(buffer) >= len(prefix) && bytesEqual(buffer[:len(prefix)], prefix)
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
