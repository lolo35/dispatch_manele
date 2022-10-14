package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const layoutISO = "2006-01-02 15:04"
const dateLayout = "2006-01-02"

func Err(err string) {
	now := time.Now()

	dateTime := now.Format(layoutISO)
	date := now.Format(dateLayout)
	cwd, _ := os.Getwd()
	mkerr := os.MkdirAll("storage/logs/", os.ModePerm)

	if mkerr != nil {
		log.Fatal(mkerr)
	}

	path := filepath.Join(cwd, "storage/logs/", fmt.Sprintf("%s.log", date))
	newFilePath := filepath.FromSlash(path)

	f, ferr := os.OpenFile(newFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if ferr != nil {
		log.Fatal(ferr)
	}

	w := bufio.NewWriter(f)

	_, werr := w.WriteString(fmt.Sprintf("Error %s : %s\n", dateTime, err))

	if werr != nil {
		log.Fatal(werr)
	}
	w.Flush()
}
