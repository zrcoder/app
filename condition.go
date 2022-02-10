package app

import (
	"context"
	"io"
	"net/url"

	"github.com/zrcoder/app/errors"
)

// Condition represents a control structure that displays nodes depending on a
// given expression.
type Condition interface {
	UI

	// ElseIf sets the condition with the given nodes if previous expressions
	// were not met and given expression is true.
	ElseIf(expr bool, elems ...UI) Condition

	// Else sets the condition with the given UI elements if previous
	// expressions were not met.
	Else(elems ...UI) Condition
}

// If returns a condition that filters the given elements according to the given
// expression.
func If(expr bool, elems ...UI) Condition {
	if !expr {
		elems = nil
	}

	return condition{
		body:      FilterUIElems(elems...),
		satisfied: expr,
	}
}

type condition struct {
	body      []UI
	satisfied bool
}

func (c condition) ElseIf(expr bool, elems ...UI) Condition {
	if c.satisfied {
		return c
	}

	if expr {
		c.body = FilterUIElems(elems...)
		c.satisfied = expr
	}

	return c
}

func (c condition) Else(elems ...UI) Condition {
	return c.ElseIf(true, elems...)
}

func (c condition) Kind() Kind {
	return Selector
}

func (c condition) JsValue() Value {
	return nil
}

func (c condition) Mounted() bool {
	return false
}

func (c condition) name() string {
	return "if.else"
}

func (c condition) self() UI {
	return c
}

func (c condition) setSelf(UI) {
}

func (c condition) context() context.Context {
	return nil
}

func (c condition) dispatcher() Dispatcher {
	return nil
}

func (c condition) attributes() map[string]string {
	return nil
}

func (c condition) eventHandlers() map[string]eventHandler {
	return nil
}

func (c condition) parent() UI {
	return nil
}

func (c condition) setParent(UI) {
}

func (c condition) children() []UI {
	return c.body
}

func (c condition) mount(Dispatcher) error {
	return errors.New("condition is not mountable").
		Tag("name", c.name()).
		Tag("kind", c.Kind())
}

func (c condition) dismount() {
}

func (c condition) update(UI) error {
	return errors.New("condition cannot be updated").
		Tag("name", c.name()).
		Tag("kind", c.Kind())
}

func (c condition) onNav(*url.URL) {
}

func (c condition) onAppUpdate() {
}

func (c condition) onAppInstallChange() {
}

func (c condition) onResize() {
}

func (c condition) preRender(Page) {
}

func (c condition) html(w io.Writer) {
	panic("shoulnd not be called")
}

func (c condition) htmlWithIndent(w io.Writer, indent int) {
	panic("shoulnd not be called")
}
