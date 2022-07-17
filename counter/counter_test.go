package counter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	const example = "Tell the audience what you're going to say. Say it. Then tell them what you've said"
	m := MapFrequencies(example)
	exp := map[string]int{
		"tell":     2,
		"the":      1,
		"audience": 1,
		"what":     2,
		"you're":   1,
		"going":    1,
		"to":       1,
		"say":      2,
		"it":       1,
		"then":     1,
		"them":     1,
		"you've":   1,
		"said":     1,
	}

	assert.Equal(t, exp, m)
}

func TestHTMLSimple(t *testing.T) {
	const example = `
		<html>
			<h1>Hello world</h1>

			<p>hello, this is my page</p>

			<p>this is a test.</p>
		</html>
	`
	m := MapFrequencies(example)
	fmt.Println(m)
	exp := map[string]int{
		"hello": 2,
		"world": 1,
		"this":  2,
		"is":    2,
		"a":     1,
		"test":  1,
		"my":    1,
		"page":  1,
	}

	assert.Equal(t, exp, m)

}

func TestEmpty(t *testing.T) {
	assert.Equal(t, MapFrequencies(""), map[string]int{})
}
