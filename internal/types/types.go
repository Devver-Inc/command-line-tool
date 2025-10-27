package types

type Model struct {
	Nodes    []Node
	Selected map[int]struct{}
	Cursor   int
}

type Node struct {
	ID   string
	Name string
	URL  string
}

type Config struct {
	Interactive bool
	Command     string
	Args        []string
}
