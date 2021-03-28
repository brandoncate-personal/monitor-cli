package domain

type Every struct {
	Seconds *int
	Minutes *int
	Hours   *int
}

type Schedule struct {
	Every Every
}

type Test struct {
	Dir  string
	File string
	Name string
}

type TestCase struct {
	Name     string
	Test     Test
	Schedule Schedule
}

type Export struct {
	Http string
}

type Config struct {
	Runner string
	Tests  []TestCase
	Export Export
}
