package commands

import "github.com/Devver-Inc/cli/internal/types"

func List(m types.Model) []types.Node {
	return m.Nodes
}
