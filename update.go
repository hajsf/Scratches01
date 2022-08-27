package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/check", Check)
	http.ListenAndServe(":1337", nil) // Logger(os.Stderr, http.DefaultServeMux)
}

func Check(w http.ResponseWriter, r *http.Request) {
	if !HasHeader(r, "X-App-Name") || !HasHeader(r, "X-App-Version") {
		return
	}

	name, version, path := GetApp(r.Header.Get("X-App-Name"))
	if name == "" {
		return
	}

	if r.Header.Get("X-App-Version") != version {
		fd, err := os.Open(path)
		defer fd.Close()
		if err != nil {
			return
		}
		FileHeader := make([]byte, 512)
		fd.Read(FileHeader)
		FileContentType := http.DetectContentType(FileHeader)

		FileStat, _ := fd.Stat()
		FileSize := strconv.FormatInt(FileStat.Size(), 10)

		w.Header().Set("Content-Type", FileContentType)
		w.Header().Set("Content-Length", FileSize)

		fd.Seek(0, 0)
		io.Copy(w, fd)
	}
}

func HasHeader(r *http.Request, h string) bool {
	return r.Header.Get(h) != ""
}

func GetApp(appname string) (name, version, path string) {
	db, _ := sql.Open("sqlite3", "./ota.db")
	defer db.Close()

	rows, err := db.Query("select name, version, path from app where name=?", appname)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&name, &version, &path)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	return "", "", ""
}
