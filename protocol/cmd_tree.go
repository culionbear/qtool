package protocol

import "github.com/culionbear/qtool/ds/queue"

type CmdTree struct {
	cmd		*queue.Manager[any]
	node	*queue.Node[any]
	bro		*CmdTree
}