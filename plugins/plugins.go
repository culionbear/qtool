package plugins

import (
	"plugin"

	"github.com/culionbear/qtool/classes"
	"github.com/culionbear/qtool/nodes"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

func Add(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return err
	}
	if err = addClass(plug); err != nil {
		return err
	}
	return addNode(plug)
}

func addClass(plug *plugin.Plugin) error {
	class, err := plug.Lookup("ClassTable")
	if err != nil {
		return err
	}
	classList, ok := class.(*[]template.Class)
	if !ok {
		return qerror.NewString("ClassTable is not the list of template.Class")
	}
	for _, v := range *classList {
		if classes.Exists(v.Name()) {
			return qerror.NewString(string(v.Name()) + " class is exist")
		}
	}
	for _, v := range *classList {
		if err := classes.Set(v.Name(), v); err != nil {
			return err
		}
	}
	return nil
}

func addNode(plug *plugin.Plugin) error {
	node, err := plug.Lookup("NodeTable")
	if err != nil {
		return nil
	}
	nodeTable, ok := node.(*map[string]template.NewNode)
	if !ok {
		return qerror.NewString("ClassTable is not the list of template.Class")
	}
	for k := range *nodeTable {
		if nodes.Exists([]byte(k)) {
			return qerror.NewString(k + " node is exist")
		}
	}
	for k, v := range *nodeTable {
		if err := nodes.Set([]byte(k), v); err != nil {
			return err
		}
	}
	return nil
}
