package animals

import (
	"testing"
)

func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func testMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func testRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
