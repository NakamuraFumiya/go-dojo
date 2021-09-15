package main

func main() {
	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}

// 真偽値表を作成
// a | b | c | a && b | !c | a && b || !c
// F | F | F |   F    | T  |     T
// F | F | T |   F    | F  |     F
// F | T | T |   F    | F  |     F
// F | T | F |   F    | T  |     T
// T | F | F |   F    | T  |     T
// T | F | T |   F    | F  |     F
// T | T | F |   T    | T  |     T
// T | T | T |   T    | F  |     T
