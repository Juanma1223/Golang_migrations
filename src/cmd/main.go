package main

import (
	sqlexec "go-migrations/src/sqlExec"
)

func main() {
	sqlexec.ApplyAllMigrations("../../migrations")
}
