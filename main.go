package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/justincampbell/anybar"
)

type Build struct {
	ActorEmail string `json:"actor_email"`
	AppSetup   struct {
		ID string `json:"id"`
	} `json:"app_setup"`
	ClearCache    bool   `json:"clear_cache"`
	CommitBranch  string `json:"commit_branch"`
	CommitMessage string `json:"commit_message"`
	CommitSha     string `json:"commit_sha"`
	CreatedAt     string `json:"created_at"`
	Debug         bool   `json:"debug"`
	Dyno          struct {
		Size string `json:"size"`
	} `json:"dyno"`
	ID           string      `json:"id"`
	Message      interface{} `json:"message"`
	Number       int64       `json:"number"`
	Organization struct {
		Name string `json:"name"`
	} `json:"organization"`
	Pipeline struct {
		ID string `json:"id"`
	} `json:"pipeline"`
	SourceBlobURL string `json:"source_blob_url"`
	Status        string `json:"status"`
	UpdatedAt     string `json:"updated_at"`
	User          struct {
		ID string `json:"id"`
	} `json:"user"`
	WarningMessage interface{} `json:"warning_message"`
}

func checkBuild() {
	anybar.Orange()

	fmt.Println("checking")

	output, err := exec.Command("heroku", "ci", "--json", "--pipeline", os.Getenv("PIPELINE")).Output()

	if err != nil {
		log.Fatal(err)
	}

	var builds []Build
	json.Unmarshal(output, &builds)

	fmt.Println(builds[0].Status)

	switch builds[0].Status {
	case "building":
		anybar.Orange()
	case "running":
		anybar.Orange()
	case "failed":
		anybar.Red()
	case "errored":
		anybar.Exclamation()
	case "succeeded":
		anybar.Green()
	default:
		anybar.Question()
	}

	fmt.Println("waiting")
}

func main() {
	timerCh := time.Tick(time.Duration(1) * time.Minute)

	checkBuild()

	for range timerCh {
		checkBuild()
	}
}
