package main

import (
	"fmt"

	"github.com/tecnologer/dicegame/server/models"
	"github.com/tecnologer/dicegame/src/utils"
)

func main() {
	models.InitGames()
	for i := 0; i < 10; i++ {
		fmt.Println(models.CreateGame("", 0))
	}

	fmt.Println(utils.GetPublicIP())
}
