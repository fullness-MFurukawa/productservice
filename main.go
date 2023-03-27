package main

import (
	"sample-service/infrastructure/sqlboiler/tests"
	"sample-service/presentation/container"
)

func main() {
	tests.TestDBInit()
	container.Execute()
}
