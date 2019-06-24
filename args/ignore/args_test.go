package ignore

import "testing"

func Test_bool(t *testing.T) {

	args := Args{Schema: "l:bool", Command: "-l"}
	if args.get("l") != false {
		t.Fail()
	}

	args = Args{Schema: "l:bool", Command: "-l FaLse"}
	if args.get("l") != false {
		t.Fail()
	}

	args = Args{Schema: "l:bool", Command: "-l true"}
	if args.get("l") != true {
		t.Fail()
	}
}

func Test_dir(t *testing.T) {

	args := Args{Schema: "f:string", Command: "-f /usr/local"}
	if args.get("f") != "/usr/local" {
		t.Fail()
	}

	args = Args{Schema: "f:string", Command: "-f"}
	if args.get("f") != "." {
		t.Fail()
	}
}

func Test_int(t *testing.T) {

	args := Args{Schema: "d:int", Command: "-d"}
	if args.get("d") != 0 {
		t.Fail()
	}

	args = Args{Schema: "d:int", Command: "-d 001"}
	if args.get("d") != 1 {
		t.Fail()
	}
}

func testInt(t *testing.T) {
	args := Args{Schema: "l:bool,d:int,f:string", Command: "-l -d 8010 -f /usr/local"}
	if args.get("l") != true {
		t.Fail()
	}

	if args.get("d") != 8080 {
		t.Fail()
	}

	if args.get("f") != "/usr/local" {
		t.Fail()
	}
}

func test2(t *testing.T) {
	args := Args{Schema: "l:bool,d:int,f:string,s:string", Command: "-l true -d -9 -f /usr/local -s"}
	if args.get("l") != true {
		t.Fail()
	}
	if args.get("d") != -9 {
		t.Fail()
	}
	if args.get("f") != "/usr/local" {
		t.Fail()
	}
	if args.get("s") != nil {
		t.Fail()
	}
}
