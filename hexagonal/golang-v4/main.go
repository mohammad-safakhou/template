package main

import "git.siz-tel.com/charging/template/cmd"

//go:generate sqlboiler --wipe --no-tests --add-soft-deletes psql -o internal/repository/boiler

func main() {
	cmd.Execute()
}
