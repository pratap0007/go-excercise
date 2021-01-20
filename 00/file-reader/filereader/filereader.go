package filereader

import (
	"io/ioutil"
)

// removes blank spaces from string with '\n'
func Filereader(path string) (string, error) {
	filedata, err := ioutil.ReadFile(path)

	return string(filedata), err

}
