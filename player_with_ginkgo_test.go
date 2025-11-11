package charlakata_test

import (
	"charlakata"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RpgPlayer", func() {

	Describe("NewRpgPlayer", func() {
		It("should start with the correct initial parameters", func() {
			player := charlakata.NewRpgPlayer(1000, 1)

			Expect(player.GetHealth()).To(Equal(1000), "Health should be 1000")
			Expect(player.GetLevel()).To(Equal(1), "Level should be 1")
			Expect(player.IsAlive()).To(BeTrue(), "Player should be alive")
		})
	})

	Describe("DealDamage", func() {
		var attackerPlayer charlakata.Player

		BeforeEach(func() {
			attackerPlayer = charlakata.NewRpgPlayer(1000, 1)
		})

		DescribeTable("should reduce the target's health correctly",
			func(damageToMake int, expectedHealth int, expectedLevel int, expectedAlive bool) {
				targetPlayer := charlakata.NewRpgPlayer(1000, 1)

				attackerPlayer.DealDamage(targetPlayer, damageToMake)

				Expect(targetPlayer.GetHealth()).To(Equal(expectedHealth))
				Expect(targetPlayer.GetLevel()).To(Equal(expectedLevel))
				Expect(targetPlayer.IsAlive()).To(Equal(expectedAlive))
			},
			Entry("less damage than current health", 200, 800, 1, true),
			Entry("same damage than current health", 1000, 0, 1, false),
			Entry("more damage than current health", 1050, 0, 1, false),
		)
	})
})
