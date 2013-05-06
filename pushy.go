package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/CSUA/pushy/model"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var config model.Configuration
var configPath *string = flag.String("config", "/etc/pushy.json", "path to configuration JSON file")
var logPath *string = flag.String("log", "/var/log/pushy.log", "path to log file")

func main() {
	// Parse command-line options
	flag.Parse()

	// Load configuration file
	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded configuration: %+v", config)
	if err := configFile.Close(); err != nil {
		log.Fatal(err)
	}

	// Create a log file and redirect output there
	logFile, err := os.OpenFile(*logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	// Make sure we have git on this computer
	if _, err := exec.LookPath("git"); err != nil {
		log.Fatalf("No git binary found. error: %v", err)
	}

	// Get uid of target user
	uid, err := GetUidByName(config.User)
	if err != nil {
		log.Fatal(err)
	}

	// Get gid of target group
	gid, err := GetGidByName(config.Group)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("UID: %d, GID: %d", uid, gid)

	// Start serving
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	payloadData := []byte(r.FormValue("payload"))
	var payload model.Payload
	json.Unmarshal(payloadData, &payload)

	if repoConfig := config.FindRepositoryConfig(payload.Repository); repoConfig != nil {
		log.Printf("Repository configuration found: %+v", repoConfig)
		pull(repoConfig)
	} else {
		log.Printf("ERROR: Repository configuration not found for: %+v", payload.Repository)
		http.Error(w, "Not configured for that repository", 400)
	}
}

func pull(repositoryConfig *model.RepositoryConfig) {
	// save the present working directory; cd back to it once this method is done
	log.Printf("Changing working directory to %s", repositoryConfig.Path)
	if oldWorkingDirectory, err := os.Getwd(); err != nil {
		log.Printf("WARNING: could not get working directory. error: %v", err)
	} else {
		defer os.Chdir(oldWorkingDirectory)
	}

	// cd to the repository directory
	if err := os.Chdir(repositoryConfig.Path); err != nil {
		log.Printf("WARNING: could not change directory; skipping repository update. error: %v", err)
		return
	}

	// git fetch origin <branch>
	fetch := exec.Command("git", "fetch", "origin", repositoryConfig.Branch)
	log.Printf("Running fetch command: %+v", fetch.Args)
	if err := fetch.Run(); err != nil {
		log.Printf("ERROR: problem running fetch command. error: %+v", err)
		return
	}

	// git checkout -f origin/<branch>
	checkout := exec.Command("git", "checkout", "-f", fmt.Sprint("origin/", repositoryConfig.Branch))
	log.Printf("Running checkout command: %+v", checkout.Args)
	if err := checkout.Run(); err != nil {
		log.Printf("ERROR: problem running checkout command. error: %+v", err)
		return
	}
	log.Printf("Repository updated!")
}
