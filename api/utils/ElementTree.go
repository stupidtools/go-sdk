package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/globalpayments/go-sdk/api/entities/enums/iflag"
	"github.com/globalpayments/go-sdk/api/entities/enums/imappedconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/istringconstant"
	"github.com/globalpayments/go-sdk/api/entities/enums/target"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/shopspring/decimal"
)

type ElementTree struct {
	doc        *xmlquery.Node
	namespaces map[string]string
}

func NewElementTree(namespaces map[string]string) *ElementTree {
	if namespaces == nil {
		namespaces = make(map[string]string)
	}
	return &ElementTree{
		namespaces: namespaces,
		doc:        &xmlquery.Node{Type: xmlquery.DocumentNode},
	}
}

func (et *ElementTree) Element(tagName string) (*Element, error) {
	var element *xmlquery.Node
	if strings.Contains(tagName, ":") {
		data := strings.Split(tagName, ":")
		namespaceURI, exists := et.namespaces[data[0]]
		if !exists {
			return nil, errors.New("namespace prefix not found")
		}
		element = &xmlquery.Node{
			Parent:       et.doc,
			Data:         data[1],
			Type:         xmlquery.ElementNode,
			Prefix:       data[0],
			NamespaceURI: namespaceURI,
		}
	} else {
		element = &xmlquery.Node{
			Parent: et.doc,
			Data:   tagName,
			Type:   xmlquery.ElementNode,
		}
	}
	return NewElement(et.doc, element, et.namespaces), nil
}

func (et *ElementTree) SubElement(parent *Element, tagName string) (*Element, error) {
	var child *xmlquery.Node
	if strings.Contains(tagName, ":") {
		data := strings.Split(tagName, ":")
		namespaceURI, exists := et.namespaces[data[0]]
		if !exists {
			return nil, errors.New("namespace prefix not found")
		}
		child = &xmlquery.Node{
			Parent:       parent.element,
			Data:         data[1],
			Type:         xmlquery.ElementNode,
			Prefix:       data[0],
			NamespaceURI: namespaceURI,
		}
	} else {
		child = &xmlquery.Node{
			Parent: parent.element,
			Data:   tagName,
			Type:   xmlquery.ElementNode,
		}
	}
	newEl := NewElement(et.doc, child, et.namespaces)
	parent.Append(newEl)
	return newEl, nil
}

// Other subElement overloads would follow a similar pattern

func (et *ElementTree) SubElementWithValue(parent *Element, tagName, value string) (*Element, error) {
	if value == "" {
		return nil, nil
	}
	el, err := et.SubElement(parent, tagName)
	if err != nil {
		return nil, err
	}
	el.Text(value)
	return el, nil
}

func (et *ElementTree) SubElementWithCdataValue(parent *Element, tagName, value string) (*Element, error) {
	if value == "" {
		return nil, nil
	}
	el, err := et.SubElement(parent, tagName)
	if err != nil {
		return nil, err
	}
	el.CData(value)
	return el, nil
}

func (et *ElementTree) ToString(root *Element) (string, error) {
	output := root.GetNode().OutputXML(true)
	return output, nil
}

func (et *ElementTree) SubElementWithInt(parent *Element, tagName string, value int) (*Element, error) {
	if value == 0 {
		return nil, nil
	}
	return et.SubElementWithValue(parent, tagName, fmt.Sprintf("%d", value))
}

func (et *ElementTree) SubElementWithDecimal(parent *Element, tagName string, value *decimal.Decimal) (*Element, error) {
	if value == nil {
		return nil, nil
	}
	return et.SubElementWithValue(parent, tagName, value.String())
}

func (et *ElementTree) SubElementWithStringConstant(parent *Element, tagName string, value istringconstant.IStringConstant) (*Element, error) {
	if value == nil {
		return nil, nil
	}
	return et.SubElementWithValue(parent, tagName, value.GetValue())
}

func (et *ElementTree) SubElementWithMappedConstant(parent *Element, tagName string, value imappedconstant.IMappedConstant) (*Element, error) {
	if value == nil {
		return nil, nil
	}
	return et.SubElementWithValue(parent, tagName, value.GetValue(target.DEFAULT))
}

func (et *ElementTree) SubElementWithFlag(parent *Element, tagName string, value iflag.IFlag) (*Element, error) {
	if value == nil {
		return nil, nil
	}
	return et.SubElementWithValue(parent, tagName, value.GetStringValue())
}

func (et *ElementTree) AddNamespace(prefix, uri string) {
	et.namespaces[prefix] = uri
}

func (et *ElementTree) Get(tagName string) *Element {
	nodes, err := xmlquery.QueryAll(et.doc, fmt.Sprintf("//%s", tagName))
	if err != nil {
		return nil
	}

	if strings.Contains(tagName, ":") {
		data := strings.Split(tagName, ":")
		namespaceURI := et.namespaces[data[0]]
		nodesReplace := make([]*xmlquery.Node, 0)
		for _, n := range nodes {
			if n.NamespaceURI == namespaceURI {
				nodesReplace = append(nodesReplace, n)
			}
		}
		nodes = nodesReplace
	}

	if len(nodes) > 0 {
		return NewElement(et.doc, nodes[0], et.namespaces)
	}
	return nil
}

func ParseXml(buffer []byte) (*ElementTree, error) {
	return ParseWithNamespaces(buffer, nil)
}

func ParseWithNamespaces(buffer []byte, namespaces map[string]string) (*ElementTree, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buffer))
	if err != nil {
		return nil, err
	}

	return &ElementTree{
		doc:        doc,
		namespaces: namespaces,
	}, nil
}
