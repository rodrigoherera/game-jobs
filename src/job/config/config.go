package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/rodrigoherera/game-jobs/src/job/domain"
)

var (
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
	DBHOST     string
	DBPORT     string
	HTTPPORT   string
	JOBS       domain.Jobs
)

func LoadEnvVars() {
	e := godotenv.Load("dev.env")
	if e != nil {
		fmt.Println("couldn't load env file")
		panic(e)
	}
	loadAppVars()
	loadJobsFile()
	loadDBVars()
}

func loadAppVars() {
	HTTPPORT = fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
}

func loadJobsFile() {
	file, err := os.Open("jobs.json")
	if err != nil {
		fmt.Println("couldn't load json file")
		panic(err)
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("couldn't read json file")
		panic(err)
	}

	err = json.Unmarshal(byteValue, &JOBS)

	if err != nil {
		fmt.Println("couldn't unmarshall json file")
		panic(err)
	}
}

func loadDBVars() {
	DBUSER = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBNAME = os.Getenv("DB_NAME")
	DBHOST = os.Getenv("DB_HOST")
	DBPORT = os.Getenv("DB_PORT")
}
