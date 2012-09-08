package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	//"io/ioutil"
	//"os"
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

//// Files

func writeFile() {
}

func readFile() {
}

func removeFile() {
}

func copyFile() {
}

//// Dirs

func createDir() {
}

func removeDir() {
}

func listDir() {
}

func copyDir() {
}

//////// REQUEST HANDLERS

//// File APIs

func fileHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path[filePathLen:]
	path.Clean(p)

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
	path.Clean(p)

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
