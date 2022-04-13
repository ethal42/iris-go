/*
 * This Source Code Form is subject to the terms of the Mozilla Public License,
 * v. 2.0. If a copy of the MPL was not distributed with this file, You can
 * obtain one at http://mozilla.org/MPL/2.0/.
 */

package iris

import (
	"fmt"
	"strconv"
	"strings"
)

const escape = "\033"

type Color struct {
	attrs []Attribute
}

type Attribute int

// Effects
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Background
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Creates a new Color object.
func New(value ...Attribute) *Color {
	return &Color{
		attrs: value,
	}
}

func (c *Color) seq() string {
	format := make([]string, len(c.attrs))

	for i, v := range c.attrs {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (c *Color) generate(s string) string {
	return c.format() + s + c.reset()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.seq())
}

func (c *Color) reset() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// Sets terminal's color.
func (c *Color) Set() {
	fmt.Print(c.format())
}

// Resets terminal's color.
func (c *Color) Unset() {
	fmt.Print(c.reset())
}

// Prints text according to c's color.
func (c *Color) Print(text string) {
	c.Set()
	defer c.Unset()
	fmt.Print(text)
}

// Prints formatted text.
func (c *Color) Printf(text string, a ...interface{}) {
	c.Set()
	defer c.Unset()
	fmt.Printf(text, a...)
}

func colorPrint(format string, p Attribute, a ...interface{}) {
	c := New(p)

	if len(a) == 0 {
		c.Print(format)
	} else {
		c.Printf(format, a...)
	}
}

func Black(format string, a ...interface{})   { colorPrint(format, FgBlack, a...) }
func Red(format string, a ...interface{})     { colorPrint(format, FgRed, a...) }
func Green(format string, a ...interface{})   { colorPrint(format, FgGreen, a...) }
func Yellow(format string, a ...interface{})  { colorPrint(format, FgYellow, a...) }
func Blue(format string, a ...interface{})    { colorPrint(format, FgBlue, a...) }
func Magenta(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }
func Cyan(format string, a ...interface{})    { colorPrint(format, FgCyan, a...) }
func White(format string, a ...interface{})   { colorPrint(format, FgWhite, a...) }
