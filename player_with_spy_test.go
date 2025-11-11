package charlakata_test

import (
	"charlakata"
	"fmt"
	"testing"
)

type SpyPlayer struct {
	level               int
	health              int
	alive               bool
	GetHealthCallCount  int
	GetLevelCallCount   int
	IsAliveCallCount    int
	DealDamageCallCount int
	TakeDamageCallCount int
}

func (s *SpyPlayer) GetHealth() int {
	s.GetHealthCallCount++
	return s.health
}

func NewSpyPlayer(level int, health int, alive bool) *SpyPlayer {
	return &SpyPlayer{level: level, health: health, alive: alive}
}

func (s *SpyPlayer) GetLevel() int {
	s.GetLevelCallCount++
	return s.level
}

func (s *SpyPlayer) IsAlive() bool {
	s.IsAliveCallCount++
	return s.alive
}

func (s *SpyPlayer) DealDamage(p charlakata.Player, damageToDo int) {
	s.DealDamageCallCount++
	s.TakeDamage(damageToDo)
}

func (s *SpyPlayer) TakeDamage(health int) {
	s.TakeDamageCallCount++
}

func TestDamageWithSpies(t *testing.T) {
	attackerPlayer := charlakata.NewRpgPlayer(1000, 1)

	cases := []struct {
		description           string
		damageToDo            int
		targetHealthRemaining int
		alive                 bool
	}{
		{"takes less than 1000", 999, 1, true},
		{"takes exactly 1000", 1000, 0, false},
		{"takes more than a 1000", 1001, 0, false},
	}

	for i, expected := range cases {

		t.Run(fmt.Sprintf("Case %d: %s", i, expected.description), func(t *testing.T) {
			targetPlayer := NewSpyPlayer(1, expected.targetHealthRemaining, expected.alive)
			attackerPlayer.DealDamage(targetPlayer, expected.damageToDo)

			assertPlayerHealth(t, targetPlayer, expected.targetHealthRemaining)
			assertAlive(t, targetPlayer, expected.alive)

			if targetPlayer.DealDamageCallCount != 0 {
				t.Errorf("Expected 0 call to DealDamage, but was: %d", targetPlayer.DealDamageCallCount)
			}

			if targetPlayer.TakeDamageCallCount != 1 {
				t.Errorf("Expected 1 call to TakeDamage, but was: %d", targetPlayer.TakeDamageCallCount)
			}
		})
	}

}
