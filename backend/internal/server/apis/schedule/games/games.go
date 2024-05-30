package games

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/algorithms/round_robin"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

var numberOfGamesPerTeam int

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/games", auth.Authenticated, handleGenerate)
}

func handleGenerate(c *fiber.Ctx) error {
	numberOfGamesPerTeam = 10
	log.Info("Reading body\n")

	// TODO read from the body
	// seasonID, csvFile, algorithm
	// TODO get ice times from csvFile

	// TODO read leagues from db

	var leagues []structures.League

	var iceTimes []string

	// TODO Call the selected algorithm
	games, err := round_robin.RoundRobin(leagues, iceTimes, numberOfGamesPerTeam)
	if err != nil {
		return responder.InternalServerError(c, fiber.StatusInternalServerError, err)
	}

	// TODO save to db

	// TODO generate analysis

	return responder.Ok(c, games)
}
