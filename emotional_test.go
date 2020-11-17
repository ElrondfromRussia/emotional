package emotional

import (
	"testing"
)

func TestHappy(t *testing.T) {
	Happy()
	HappyCat()
	Love()
}

func TestSadness(t *testing.T) {
	Cry()
	Sad()
	Dead()
	DeadLine()
}

func TestNeutral(t *testing.T) {
	PockerFace()
	Shrugging()
	Iiiiiigor()
	SenseyReaction()
	WowReally()
	Gopher()
	DruggedGopher()
	CoolGopher()
	Cat()
}

func TestEvil(t *testing.T) {
	Fuck()
	Disgust()
	Grrr()
	GoAway()
}

func TestTeam(t *testing.T) {
	Backend()
	Frontend()
	NewGuys()
}

// do not test "madness"...just be crazy in the reality...
func TestMadness1(t *testing.T) {
	Depression()
}

func TestMadness2(t *testing.T) {
	Apatic()
}

func TestMadness3(t *testing.T) {
	Isteric()
}
