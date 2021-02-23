package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	assert.Equal(t, "SH600030", _getSymbolCode("600030"))
	assert.Equal(t, "SZ300750", _getSymbolCode("300750"))
	assert.Equal(t, "00005", _getSymbolCode("00005"))
}

func TestGetPrice(t *testing.T) {
	cookies := _getCookies()
	fmt.Println(GetQuoteData("600036", cookies))
	fmt.Println(GetQuoteData("03690", cookies))
	fmt.Println(GetQuoteData("01810", cookies))
}
