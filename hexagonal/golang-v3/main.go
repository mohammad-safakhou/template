package main

//go:generate sqlboiler --wipe psql -o adapters/repository/models

import (
	"template/cmd"
)

func main() {
	cmd.Execute()
}
