package main

import (
	"flag"
	"strings"
	"os"
	"log"
	"fmt"
)

func recursiveScanFolder(folder string) [] string {
	return scanGitFolders(make([]string, 0), folder)
}
/*
Recursively finds and stores all the folders ending with .git
Does not explore .git folders, and skips vendor and node_modules folders
*/
func scanGitFolders(folders []string, folder string) [] string {
	folder = strings.TrimSuffix(folder, "/")	

	f,err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			
		}
	}

	return folders;
}

func scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	print(repositories)
}

func stats(email string){
	print("stats")
}

func main(){
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repos")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if (folder != ""){
		scan(folder)
		return
	}

	stats(email)
}