package stampery

type Anchor struct {
	Chain int
	Tx    string
}
type Proof struct {
	Version  int
	Siblings []string
	Root     string
	Anchor   Anchor
}

type Proof2 struct {
	Version  uint64
	Siblings interface{}
	Root     string
	Anchor   Anchor
}

type Event struct {
	Type string
	Data interface{}
}
