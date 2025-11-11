package charlakata_test

import (
	"charlakata"
	"fmt"
	"testing"
)

func TestPlayerStartsWithCorrectParameters(t *testing.T) {
	player := charlakata.NewRpgPlayer(1000, 1)

	assertPlayerHealth(t, player, 1000)
	assertPlayerLevel(t, player, 1)
	assertAlive(t, player, true)
}

func TestDealDamageReducesHealthOfTarget(t *testing.T) {
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

			assertPlayerHealth(t, targetPlayer, expectedCase.expectedHealth)
			assertPlayerLevel(t, targetPlayer, expectedCase.expectedLevel)
			assertAlive(t, targetPlayer, expectedCase.expectedAlive)
		})
	}
}

func assertAlive(t *testing.T, player charlakata.Player, expectedAlive bool) {
	t.Helper()
	if player.IsAlive() != expectedAlive {
		t.Errorf("player.isAlive should be %v", expectedAlive)
	}
}

func assertPlayerLevel(t *testing.T, player charlakata.Player, expectedLevel int) {
	t.Helper()
	if player.GetLevel() != expectedLevel {
		t.Errorf("player.level should be %d", expectedLevel)
	}
}

func assertPlayerHealth(t *testing.T, player charlakata.Player, expectedHealth int) {
	t.Helper()
	actualHeath := player.GetHealth()
	if actualHeath != expectedHealth {
		t.Errorf("player.health should be %d, but was: %d", expectedHealth, actualHeath)
	}
}
