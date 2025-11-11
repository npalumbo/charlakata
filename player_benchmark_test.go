package charlakata_test

import (
	"charlakata"
	"testing"
)

func BenchmarkPlayerRPG_Attack(b *testing.B) {
	attackerPlayer := charlakata.NewRpgPlayer(1000, 1)
	for b.Loop() {
		targetPlayer := charlakata.NewRpgPlayer(1000, 1)
		attackerPlayer.DealDamage(targetPlayer, 10)
	}
}
