package main

import (
	"context"
	"os"
	"os/signal"
	"project-user/configurations"
	"project-user/helpers"
	"project-user/setups"
	"project-user/utils"
)

func main() {
	configuration := configurations.NewConfiguration()

	mysqlUtil := utils.NewMysqlConnection(configuration.ProjectUserMysqlUsername, configuration.ProjectUserMysqlPassword, configuration.ProjectUserMysqlHost, configuration.ProjectUserMysqlDatabase, configuration.ProjectUserMysqlMaxOpenConnection, configuration.ProjectUserMysqlMaxIdleConnection, configuration.ProjectUserMysqlConnectionMaxLifetime, configuration.ProjectUserMysqlConnectionMaxIdletime)
	defer mysqlUtil.Close()

	redisUtil := utils.NewRedisConnection(configuration.ProjectUserRedisHost, configuration.ProjectUserRedisPort, configuration.ProjectUserRedisDatabase)
	defer redisUtil.Close()

	validate := setups.SetValidator()
	bcryptHelper := helpers.NewBcryptHelper()
	uuidHelper := helpers.NewUuidHelper()

	e := setups.SetEcho(mysqlUtil, redisUtil, validate, bcryptHelper, uuidHelper)
	setups.StartEcho(e, configuration.ProjectUserApplicationHost)
	defer setups.StopEcho(e, configuration.ProjectUserApplicationHost)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
}
