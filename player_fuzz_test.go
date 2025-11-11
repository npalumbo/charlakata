package charlakata_test

import (
	"charlakata"
	"testing"
)

func TestDealDamageFromFuzz(t *testing.T) {
	attackerPlayer := charlakata.NewRpgPlayer(1000, 1)

	targetPlayer := charlakata.NewRpgPlayer(1, 1)

	attackerPlayer.DealDamage(targetPlayer, 86)

	assertPlayerHealth(t, targetPlayer, 0)
	assertPlayerLevel(t, targetPlayer, 1)
	assertAlive(t, targetPlayer, false)
}

func Fuzz_Damage_Boundaries(f *testing.F) {
	f.Add(1000, 1)  // Full health, small damage
	f.Add(1000, 0)  // Full health, zero damage
	f.Add(100, 100) // Exact lethal damage

	attackerPlayer := charlakata.NewRpgPlayer(1000, 1)

	f.Fuzz(func(t *testing.T, initialHealth, damageAmount int) {
		h := getAlwaysPositiveFromOneToThousand(initialHealth)

		targetPlayer := charlakata.NewRpgPlayer(h, 1)

		attackerPlayer.DealDamage(targetPlayer, damageAmount)

		if targetPlayer.GetHealth() < 0 {
			t.Errorf("Health must not be negative, got %d after %d damage", targetPlayer.GetHealth(), damageAmount)
		}

		if damageAmount > 0 && targetPlayer.GetHealth() > h {
			t.Errorf("Damage increased health. Initial: %d, Damage: %d, Final: %d", h, damageAmount, targetPlayer.GetHealth())
		}

		if targetPlayer.GetHealth() == 0 && targetPlayer.IsAlive() {
			t.Errorf("Character is Dead (Health 0) but IsAlive is true.")
		}
		if targetPlayer.GetHealth() > 0 && !targetPlayer.IsAlive() {
			t.Errorf("Character is Alive (Health %d) but IsAlive is false.", targetPlayer.GetHealth())
		}
	})
}

func getAlwaysPositiveFromOneToThousand(initialHealth int) int {
	h := initialHealth%1000 + 1
	if h < 0 {
		h = -h
	}
	return h
}
