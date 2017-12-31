package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	name       string
	attributes map[string]interface{}
	constant   node.Node
}

func NewConstFetch(constant node.Node) node.Node {
	return ConstFetch{
		"ConstFetch",
		map[string]interface{}{},
		constant,
	}
}

func (n ConstFetch) Name() string {
	return "ConstFetch"
}

func (n ConstFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ConstFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ConstFetch) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n ConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constant != nil {
		vv := v.GetChildrenVisitor("constant")
		n.constant.Walk(vv)
	}

	v.LeaveNode(n)
}
