package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Throw struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewThrow(expr node.Node) node.Node {
	return Throw{
		"Throw",
		map[string]interface{}{},
		expr,
	}
}

func (n Throw) Name() string {
	return "Throw"
}

func (n Throw) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Throw) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Throw) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Throw) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
