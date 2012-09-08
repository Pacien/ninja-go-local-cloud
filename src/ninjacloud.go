package main

import (
	"encoding/json"
	//"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"path"
	"io"
	"path/filepath"
)

const APP_NAME = "Ninja Go Local Cloud"
const APP_VERSION = "0.1 Draft"

var versionFlag bool
var interfaceFlag string
var portFlag string
var rootFlag string

const filePath = "/file/"
const dirPath = "/directory/"
const webPath = "/web?url="
const statusPath = "/cloudstatus"

const filePathLen = len(filePath)
const dirPathLen = len(dirPath)
const webPathLen = len(webPath)
const statusPathLen = len(statusPath)

//////// FILESYSTEM

func exist(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

//// Files

func writeFile(path string, content []byte, overwrite bool) (err error) {
	if !overwrite {
		if exist(path) {
			err = os.ErrExist
			return
		}
	}
	err = ioutil.WriteFile(path, content, 0600)
	return
}

func readFile(path string) (content []byte, err error) {
	content, err = ioutil.ReadFile(path)
	return
}

func removeFile(path string) (err error) {
	err = os.Remove(path)
	return
}

func moveFile(source string, dest string) (err error) {
	err = os.Rename(source, dest)
	return
}

func copyFile(source string, dest string) (err error) {
	// from https://gist.github.com/2876519
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}
	return
}

//// Dirs

func createDir(path string) (err error) {
	err = os.MkdirAll(path, 0600)
	return
}

func removeDir(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func listDir(path string) (list []os.FileInfo, err error) {
	list, err = ioutil.ReadDir(path)
	return
}

func moveDir(source string, dest string) (err error) {
	err = os.Rename(source, dest)
	return
}

func copyDir(source string, dest string) (err error) {
	// from https://gist.github.com/2876519
	fi, err := os.Stat(source)
	if err != nil {
		return
	}
	if !fi.IsDir() {
		return os.ErrInvalid
	}
	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return os.ErrExist
	}
	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return
	}
	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				return
			}
		} else {
			err = copyFile(sfp, dfp)
			if err != nil {
				return
			}
		}
	}
	return
}

//////// REQUEST HANDLERS

//// File APIs

func fileHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[filePathLen:]
	filepath.Clean(p)

	switch r.Method {
	case "POST":
		// Create a new file
	case "PUT":
		if r.Header.Get("sourceURI") == "" {
			// Update an existing file (save over existing file)
		} else {
			// Copy, Move of an existing file 
		}
	case "DELETE":
		// Delete an existing file
	case "GET":
		// Read an existing file
	}

}

//// Directory APIs

func dirHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[dirPathLen:]
	filepath.Clean(p)

	switch r.Method {
	case "POST":
		// Create a new directory
	case "DELETE":
		// Delete an existing directory
	case "GET":
		// List the contents of an existing directory
	case "PUT":
		// Copy, Move of an existing directory
	}

}

//// Web API

// Get text or binary data from a URL
func getDataHandler(w http.ResponseWriter, r *http.Request) {
}

//// Cloud Status API

// Get the cloud status JSON
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	cloudStatus := map[string]string{
		"name":        APP_NAME,
		"version":     APP_VERSION,
		"server-root": rootFlag,
		"status":      "running",
	}
	j, err := json.Marshal(cloudStatus)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

//////// INIT and MAIN

func init() {
	flag.BoolVar(&versionFlag, "v", false, "Print the version number.")
	flag.StringVar(&interfaceFlag, "i", "localhost", "Listening interface.")
	flag.StringVar(&portFlag, "p", "58080", "Listening port.")
	flag.StringVar(&rootFlag, "r", ".", "Root directory.")
}

func main() {
	flag.Parse()

	if versionFlag {
		log.Println("Version:", APP_VERSION)
		return
	}

	log.Println("Starting " + APP_NAME + " " + APP_VERSION + " on " + interfaceFlag + ":" + portFlag + " in " + rootFlag)

	http.HandleFunc(filePath, fileHandler)
	http.HandleFunc(dirPath, dirHandler)
	http.HandleFunc(webPath, getDataHandler)
	http.HandleFunc(statusPath, getStatusHandler)

	err := http.ListenAndServe(interfaceFlag+":"+portFlag, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
