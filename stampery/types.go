package stampery

type anchor struct {
	chain int
	tx    string
}

type Proof struct {
	version  int
	siblings []string
	root     string
	anchor   anchor
}

type Event struct {
	Type string
	Data interface{}
}
