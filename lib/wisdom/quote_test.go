package wisdom_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/arv28/wisdom-service/lib/wisdom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	q := NewQuote("Learn Go!", "Ajai")
	assert.Equal(t, "Learn Go!", q.Quote)
	assert.Equal(t, "Ajai", q.Author)
}

func TestQuoteAsString(t *testing.T) {
	q := NewQuote("Learn Go!", "Ajai")
	assert.Equal(t, "Learn Go! - Ajai\n", fmt.Sprintf("%s", q))
}

func TestQuoteAsJson(t *testing.T) {
	q := NewQuote("Learn Go!", "Ajai")
	expected := `{ "quote": "Learn Go!", "author": "Ajai" }`
	bytes, err := json.Marshal(q)
	require.Nil(t, err)
	assert.JSONEq(t, expected, string(bytes))
}
