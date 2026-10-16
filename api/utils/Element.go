package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/shopspring/decimal"
)

type Element struct {
	doc        *xmlquery.Node
	element    *xmlquery.Node
	namespaces map[string]string
}

func NewElement(doc, element *xmlquery.Node, namespaces map[string]string) *Element {
	return &Element{
		doc:        doc,
		element:    element,
		namespaces: namespaces,
	}
}

func (e *Element) FirstChild() *Element {
	return FromNode(e.doc, e.element.FirstChild, e.namespaces)
}

func (e *Element) Remove(tagName string) *Element {
	child := e.Get(tagName)
	if child != nil {
		xmlquery.RemoveFromTree(e.element)
	}
	return e
}

func (e *Element) GetRootDoc() *xmlquery.Node {
	return e.doc
}

func (e *Element) GetNode() *xmlquery.Node {
	return e.element
}

func (e *Element) CData(text string) *Element {

	node := &xmlquery.Node{
		Type: xmlquery.CharDataNode,
		Data: text,
	}
	xmlquery.AddChild(e.element, node)
	return e
}

func (e *Element) Set(name, value string) *Element {
	e.element.SetAttr(name, value)
	return e
}

func (e *Element) Text(text string) *Element {
	textNode := &xmlquery.Node{
		Type: xmlquery.TextNode,
		Data: text,
	}
	xmlquery.AddChild(e.element, textNode)
	return e
}

func (e *Element) Append(child *Element) *Element {
	xmlquery.AddChild(e.element, child.element)
	return e
}

func (e *Element) Tag() string {
	return e.element.Data
}

func FromNode(doc, node *xmlquery.Node, namespaces map[string]string) *Element {
	return NewElement(doc, node, namespaces)
}

func (e *Element) Has(tagName string) bool {
	return xmlquery.FindOne(e.element, fmt.Sprintf("//%s", tagName)) != nil
}

func (e *Element) Get(tagName string) *Element {
	node := xmlquery.FindOne(e.element, fmt.Sprintf("//%s", tagName))
	if node == nil {
		return nil
	}
	return FromNode(e.doc, node, e.namespaces)
}

func (e *Element) GetAll() []*Element {
	nodes := make([]*xmlquery.Node, 0)
	for child := e.element.FirstChild; child != nil; child = child.NextSibling {
		nodes = append(nodes, child)
	}
	elements := make([]*Element, len(nodes))
	for i, node := range nodes {
		elements[i] = FromNode(e.doc, node, e.namespaces)
	}
	return elements
}

func (e *Element) GetAllByTag(tagName string) []*Element {
	nodes, err := xmlquery.QueryAll(e.element, fmt.Sprintf("//%s", tagName))
	elements := make([]*Element, len(nodes))
	if err != nil {
		return elements
	}
	for i, node := range nodes {
		elements[i] = FromNode(e.doc, node, e.namespaces)
	}
	return elements
}

func (e *Element) GetAttributeString(attributeName string) string {
	return e.element.SelectAttr(attributeName)
}

func (e *Element) GetString(tagNames ...string) string {
	for _, tagName := range tagNames {
		element := e.getElementByTagName(tagName)
		if element != nil {
			return element.InnerText()
		}
	}
	return ""
}

func (e *Element) GetBool(tagName string) (bool, error) {
	element := e.getElementByTagName(tagName)
	if element != nil {
		return strconv.ParseBool(element.InnerText())
	}
	return false, errors.New("element not found")
}

func (e *Element) GetInt(tagName string) *int {
	element := e.getElementByTagName(tagName)
	if element != nil {
		res, err := strconv.Atoi(element.InnerText())
		if err == nil {
			return &res
		}
	}
	return nil
}

func (e *Element) GetDecimal(tagName string) *decimal.Decimal {
	element := e.getElementByTagName(tagName)
	if element != nil {
		res, err := decimal.NewFromString(element.InnerText())
		if err == nil {
			return &res
		}
	}
	return nil
}

func (e *Element) GetDate(tagNames ...string) string {
	layout := "2006-01-02T15:04:05.9999999"
	return e.GetDateWithFormat(layout, tagNames...)
}

func (e *Element) GetDateWithFormat(layout string, tagNames ...string) string {
	for _, tagName := range tagNames {
		element := e.getElementByTagName(tagName)
		if element != nil {
			res, err := time.Parse(layout, element.InnerText())
			if err == nil {
				return res.Format(time.RFC3339)
			}
		}
	}
	return ""
}

func (e *Element) getElementByTagName(tagName string) *xmlquery.Node {
	node := xmlquery.FindOne(e.element, fmt.Sprintf("//%s", tagName))
	if node == nil {
		node = xmlquery.FindOne(e.doc, fmt.Sprintf("//%s", tagName))
	}
	return node
}
