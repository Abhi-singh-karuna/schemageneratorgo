package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FileNameCreation(fileName string) string {
	return fmt.Sprintf("%s %s", fileName[:len(fileName)-len(filepath.Ext(fileName))], ".go")
}

func ExtractFileName(x string) string {

	res2 := strings.Split(x, ".")
	res3 := strings.Split(res2[0], "/")


	return (res3[len(res3)-1])

}
