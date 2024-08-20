package main

import (
	"context"
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/setups"
	"golang-postgres/commons/utils"
	"os"
	"os/signal"
)

func main() {
	postgresUtil := utils.NewPostgresConnection()
	defer postgresUtil.Close()

	redisUtil := utils.NewRedisConnection()
	defer redisUtil.Close()

	validate := setups.SetValidator()
	bcryptHelper := helpers.NewBcryptHelper()
	uuidHelper := helpers.NewUuidHelper()

	e := setups.SetEcho(postgresUtil, redisUtil, validate, bcryptHelper, uuidHelper)
	setups.StartEcho(e)
	defer setups.StopEcho(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()
}
