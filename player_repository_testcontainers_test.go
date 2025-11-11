package charlakata_test

import (
	"charlakata"
	"charlakata/testhelpers"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PostgresPlayerRepositoryTestSuite struct {
	suite.Suite
	pgContainer      *testhelpers.PostgresContainer
	playerRepository charlakata.PlayerRepository
	ctx              context.Context
}

func (suite *PostgresPlayerRepositoryTestSuite) TestLoadAllPlayers() {
	t := suite.T()

	// When
	loadedPlayers, err := suite.playerRepository.LoadAllPlayers(suite.ctx)

	// Then
	assert.NoError(t, err)
	assert.Len(t, loadedPlayers, 2)

	firstPlayer := loadedPlayers[0]
	assert.Equal(t, 1, firstPlayer.ID)
	assert.Equal(t, 100, firstPlayer.Health)

	secondPlayer := loadedPlayers[1]
	assert.Equal(t, 2, secondPlayer.ID)
	assert.Equal(t, 150, secondPlayer.Health)

}

func (suite *PostgresPlayerRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.T(), suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	playerRepository, err := charlakata.NewPostgresPlayerRepository(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.playerRepository = playerRepository
}

func (suite *PostgresPlayerRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestPostgresPlayerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresPlayerRepositoryTestSuite))
}
