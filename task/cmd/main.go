package main

import (
	"fmt"
	"os"
)

func GetENVVariables() {

	fmt.Println("Local env is working...")
	os.Setenv("SERVER_RUN_PORT", "127.0.0.1:9091")
	os.Setenv("JWT_SECRET", "2JxyMBGNMR8p4Bdaiua5QY8SjkTXhWkEci3rR3An")
	os.Setenv("JWT_EXP_TIME", "600")
	os.Setenv("JWT_SECRET_CLIENT", "GP6ziyBQLJWAVjhrFC2WZv2SaVah0mDBCL64cg6n")
	os.Setenv("JWT_EXP_TIME_CLIENT", "1")

	os.Setenv("ABS_FILIAL_CODE", "09012")
	os.Setenv("ABS_USER_LOGIN", "pdp_hr")
	os.Setenv("ABS_USER_PASSWORD", "pdp_hr")

	os.Setenv("POSTGRES_HOST", "127.0.0.1")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_DB", "auth_service")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "root")

	// os.Setenv("AD_HOST", "172.25.102.200:636")
	os.Setenv("AD_HOST", "hamkor.local:389")
	os.Setenv("AD_MAIL", "@hamkor.local")
	os.Setenv("VERIFY", "true")

	os.Setenv("ABSAPI_URL", "http://172.25.105.20:2030/api")
	os.Setenv("RPC_CLIENT", "127.0.0.1:7099")

	// os.Setenv("RABBIT_MQ_URL", "amqp://oneid:oneid@172.25.105.20:5672/oneid")
	// os.Setenv("RABBIT_MQ_QUEUE", "oneid_hr_queue")
	os.Setenv("RABBIT_MQ_URL", "amqp://rabbit:rabbit@172.25.105.20:5672")
	os.Setenv("RABBIT_MQ_QUEUE", "oneid_queue")
}

func main() {
	//app.Run()
	GetENVVariables()
	fmt.Println(os.Getenv("SERVER_RUN_PORT"))
	fmt.Println(os.Getenv("test"))
	fmt.Println(os.Getenv("t"))
	s, _ := os.LookupEnv("RABBIT_MQ_URL")
	fmt.Println(s)
}
