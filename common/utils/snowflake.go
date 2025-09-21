package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func Setup(nodeId int64) error {
	var err error
	node, err = snowflake.NewNode(nodeId)
	return err
}
func GenInt() int64 {
	return node.Generate().Int64()
}

func GenString() string {
	return node.Generate().String()
}

func Gen() snowflake.ID {
	return node.Generate()
}
