//go:build !wasm

package app

import (
	"net/url"
	"runtime"

	"github.com/zrcoder/app/errors"
)

var (
	window    = &browserWindow{}
	errNoWasm = errors.New("unsupported instruction").
			Tag("required-architecture", "wasm").
			Tag("current-architecture", runtime.GOARCH)
)

type value struct{}

func (v value) Bool() bool {
	return false
}

func (v value) Call(m string, args ...interface{}) Value {
	return value{}
}

func (v value) Float() float64 {
	return 0
}

func (v value) Get(p string) Value {
	return value{}
}

func (v value) Index(i int) Value {
	return value{}
}

func (v value) InstanceOf(t Value) bool {
	return false
}

func (v value) Int() int {
	return 0
}

func (v value) Invoke(args ...interface{}) Value {
	return value{}
}

func (v value) IsNaN() bool {
	return false
}

func (v value) IsNull() bool {
	return true
}

func (v value) IsUndefined() bool {
	return true
}

func (v value) JsValue() Value {
	return v
}

func (v value) Length() int {
	return 0
}

func (v value) New(args ...interface{}) Value {
	return value{}
}

func (v value) Set(p string, x interface{}) {
}

func (v value) SetIndex(i int, x interface{}) {
}

func (v value) String() string {
	return ""
}

func (v value) Truthy() bool {
	return false
}

func (v value) Type() Type {
	panic(errNoWasm)
}

func (v value) getAttr(k string) string {
	return ""
}

func (v value) setAttr(k, val string) {
}

func (v value) delAttr(k string) {
}

func (v value) firstChild() Value {
	return value{}
}

func (v value) appendChild(c Wrapper) {
}

func (v value) replaceChild(new, old Wrapper) {
}

func (v value) removeChild(c Wrapper) {
}

func (v value) firstElementChild() Value {
	return value{}
}

func (v value) addEventListener(event string, fn Func) {
}

func (v value) removeEventListener(event string, fn Func) {
}

func (v value) setNodeValue(val string) {
}

func (v value) setInnerHTML(val string) {
}

func (v value) setInnerText(val string) {
}

func null() Value {
	return value{}
}

func undefined() Value {
	return value{}
}

func valueOf(x interface{}) Value {
	return value{}
}

type function struct {
	value
}

func (f function) Release() {
}

func funcOf(fn func(this Value, args []Value) interface{}) Func {
	return function{value: value{}}
}

type browserWindow struct {
	value
}

func (w browserWindow) URL() *url.URL {
	return &url.URL{}
}

func (w browserWindow) Size() (width, height int) {
	return 0, 0
}

func (w browserWindow) CursorPosition() (x, y int) {
	return 0, 0
}

func (w browserWindow) setCursorPosition(x, y int) {
}

func (w *browserWindow) GetElementByID(id string) Value {
	return value{}
}

func (w *browserWindow) ScrollToID(id string) {
}

func (w *browserWindow) AddEventListener(event string, h EventHandler) func() {
	return func() {}
}

func (w *browserWindow) setBody(body UI) {
}

func (w *browserWindow) createElement(tag string) (Value, error) {
	return value{}, nil
}

func (w *browserWindow) createTextNode(v string) Value {
	return value{}
}

func (w *browserWindow) addHistory(u *url.URL) {
}

func (w *browserWindow) replaceHistory(u *url.URL) {
}

func copyBytesToGo(dst []byte, src Value) int {
	return 0
}

func copyBytesToJS(dst Value, src []byte) int {
	return 0
}
