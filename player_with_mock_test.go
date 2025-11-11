package charlakata_test

import (
	"charlakata"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestDamageWithMocks(t *testing.T) {
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
			targetPlayer := charlakata.NewMockPlayer(t)
			// Given
			targetPlayer.EXPECT().TakeDamage(mock.AnythingOfType("int")).Return()
			targetPlayer.EXPECT().GetHealth().Return(expected.targetHealthRemaining)
			targetPlayer.EXPECT().IsAlive().Return(expected.alive)

			// When
			attackerPlayer.DealDamage(targetPlayer, expected.damageToDo)

			// Then
			assertPlayerHealth(t, targetPlayer, expected.targetHealthRemaining)
			assertAlive(t, targetPlayer, expected.alive)

			targetPlayer.AssertExpectations(t)

			targetPlayer.AssertNumberOfCalls(t, "GetHealth", 2)
			targetPlayer.AssertNumberOfCalls(t, "DealDamage", 0)
			targetPlayer.AssertNumberOfCalls(t, "TakeDamage", 1)

		})
	}

}
