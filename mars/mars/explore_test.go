package mars

import "testing"

func TestTravel_CreateTravel(t *testing.T) {
	travel := CreateTravel()

	if travel != nil {
		t.Log("crete travel success")
	} else {
		t.Fatal("create travel fail")
	}
}

func TestTravel_MakeMarsSize(t *testing.T) {
	travel := CreateTravel()
	travel.MakeMarsSize(5, 5)
	if travel.MarsMap.X == 5 && travel.MarsMap.Y == 5 {
		t.Log("make mars size success")
	} else {
		t.Fatal("make mars size fail")
	}
}

func TestTravel_DeployCar(t *testing.T) {
	travel := CreateTravel()
	travel.MakeMarsSize(5, 5)
	startPosition := Position{1, 1, NORTH}
	travel.DeployCar(startPosition)

	if travel.CarActive != nil && travel.CarActive.CurrentPosition.X == 1 {
		t.Log("deploy car success")
	} else {
		t.Fatal("deploy car fail")
	}
}

func TestTravel_SendCmd(t *testing.T) {
	travel := CreateTravel()
	travel.MakeMarsSize(5, 5)
	startPosition := Position{1, 1, NORTH}
	travel.DeployCar(startPosition)
	cmd := "LLRRM"

	travel.SendCmd(cmd)

	if travel.CarActive != nil && travel.CarActive.CurrentPosition.X == 1 && travel.CarActive.CurrentPosition.Y == 2 {
		t.Log("send cmd success")
	} else {
		t.Fatal("")
	}
}

func TestTravel_CarDrop(t *testing.T) {
	travel := CreateTravel()
	travel.MakeMarsSize(1, 1)
	startPosition := Position{1, 1, NORTH}
	travel.DeployCar(startPosition)

	carId := travel.CarActive.id
	cmd := "M"
	travel.SendCmd(cmd)
	if travel.CarActive == nil && travel.CarDrop != nil && travel.CarDrop[carId].X == 1 && travel.CarDrop[carId].Y == 1 {
		t.Log("car drop success")
	} else {
		t.Fatal("")
	}
}

func TestTravel_IgnoreCmd(t *testing.T) {
	travel := CreateTravel()
	travel.MakeMarsSize(1, 1)
	startPosition := Position{1, 1, NORTH}
	travel.DeployCar(startPosition)

	carId := travel.CarActive.id
	cmd := "M"
	// 第一辆车执行之后掉下去了
	travel.SendCmd(cmd)

	// 部署第二辆车
	travel.DeployCar(startPosition)
	cmdNew := "ML"
	travel.SendCmd(cmdNew)

	if travel.CarDrop != nil && travel.CarDrop[carId].X == 1 && travel.CarDrop[carId].Y == 1 {
		// 车 1 掉下去了, 所以忽略了M,只执行了L
		if travel.CarActive != nil && travel.CarActive.CurrentPosition.Direction == WEST {
			t.Log("car two finish job success")
		} else {
			t.Log("car two do the job fail")

		}
	} else {
		t.Log("car one drop fail")
	}
}
