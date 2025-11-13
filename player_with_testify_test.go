package charlakata_test

import (
	"charlakata"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerStartsWithCorrectParametersTestify(t *testing.T) {
	player := charlakata.NewRpgPlayer(1000, 1)

	assert.Equal(t, 1000, player.GetHealth(), "Health should be 1000")
	assert.Equal(t, 1, player.GetLevel(), "Level should be 1")
	assert.True(t, player.IsAlive(), "Player should be alive")
}

func TestDealDamageReducesHealthOfTargetTestify(t *testing.T) {
	cases := []struct {
		description    string
		damageToMake   int
		expectedHealth int
		expectedLevel  int
		expectedAlive  bool
	}{
		{"less damage than current health", 200, 800, 1, true},
		{"same damage than current health", 1000, 0, 1, false},
		{"more damage than current health", 1050, 0, 1, false},
	}
	attackerPlayer := charlakata.NewRpgPlayer(1000, 1)

	for i, expectedCase := range cases {
		t.Run(fmt.Sprintf("Case %d: %s", i, expectedCase.description), func(t *testing.T) {
			targetPlayer := charlakata.NewRpgPlayer(1000, 1)

			attackerPlayer.DealDamage(targetPlayer, expectedCase.damageToMake)

			assert.Equal(t, expectedCase.expectedHealth, targetPlayer.GetHealth())
			assert.Equal(t, expectedCase.expectedLevel, targetPlayer.GetLevel())
			assert.Equal(t, expectedCase.expectedAlive, targetPlayer.IsAlive())
		})
	}
}
