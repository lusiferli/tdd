package mars

import (
	"testing"
)

func TestCar_TurnLeft(t *testing.T) {
	startPosition := Position{0, 0, NORTH}
	car := CreateCar(startPosition)
	car.StartPosition = startPosition

	positionNow, _ := car.TurnLeft()
	if positionNow.X == 0 && positionNow.Y == 0 && positionNow.Direction == WEST {
		t.Log("car turn left success")

	} else {
		t.Fatal("car turn left fail")
	}
}

func TestCar_TurnRight(t *testing.T) {
	startPosition := Position{0, 0, NORTH}
	car := CreateCar(startPosition)
	car.StartPosition = startPosition

	positionNow, _ := car.TurnRight()
	if positionNow.X == 0 && positionNow.Y == 0 && positionNow.Direction == EAST {
		t.Log("car turn right success")
	} else {
		t.Fatal("car turn right fail")
	}
}

func TestCar_TurnRight_Two(t *testing.T) {
	startPosition := Position{0, 0, EAST}
	car := CreateCar(startPosition)
	car.StartPosition = startPosition

	positionNow, _ := car.TurnRight()
	if positionNow.X == 0 && positionNow.Y == 0 && positionNow.Direction == SOUTH {
		t.Log("car turn right success")
	} else {
		t.Fatal("car turn right fail")
	}
}

func TestCar_Move(t *testing.T) {
	startPosition := Position{0, 0, NORTH}
	car := CreateCar(startPosition)
	car.StartPosition = startPosition

	positionNow, _ := car.Move()
	if positionNow.X == 0 && positionNow.Y == 1 && positionNow.Direction == NORTH {
		t.Log("car move success")
	} else {
		t.Fatal("car move fail")
	}
}

func TestCar_RunCmd(t *testing.T) {
	startPosition := Position{0, 0, NORTH}
	car := CreateCar(startPosition)
	car.StartPosition = startPosition

	positionNow, _ := car.RunCmd("M")
	if positionNow.X == 0 && positionNow.Y == 1 && positionNow.Direction == NORTH {
		t.Log("car turn right success")
	} else {
		t.Fatal("")
	}
}
