package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rodrigoherera/game-jobs/src/job/config"
	"github.com/rodrigoherera/game-jobs/src/job/datastore/mysql"
	"github.com/rodrigoherera/game-jobs/src/job/repository"
	"github.com/rodrigoherera/game-jobs/src/job/service"
)

func main() {
	config.LoadEnvVars()
	jobs := config.JOBS

	s := gocron.NewScheduler(time.UTC)

	db := mysql.NewMysql(config.DBUSER, config.DBPASSWORD, config.DBNAME, config.DBHOST, config.DBPORT)

	characterRepo := repository.NewCharacterRepo(db)
	characterService := service.NewCharacterHandler(characterRepo)

	for _, task := range jobs.Tasks {
		switch task.Name {
		case "RestoreLife":
			s.Cron(task.Cron).Do(characterService.RestoreLife)
		case "RestoreStamina":
			s.Cron(task.Cron).Do(characterService.RestoreStamina)
		case "RenewShops":
			fmt.Println("Not implemented yet")
		default:
			fmt.Println("Default value")
		}
	}

	s.StartBlocking()
}
