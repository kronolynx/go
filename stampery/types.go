package stampery

//Anchor type that represents a stampery proof
type Anchor struct {
	Chain int
	Tx    string
}

//Proof for stampery
type Proof struct {
	Version  int
	Siblings []string
	Root     string
	Anchor   Anchor
}

//temporary placeholder when the proof doesn't have siblings
type temp struct {
	Version  int
	Siblings interface{}
	Root     string
	Anchor   Anchor
}

//Event has 3 types: ready, proof, error
type Event struct {
	Type string
	Data interface{}
}
