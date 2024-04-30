package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapSum(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	assert.Equal(t, 3, mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	assert.Equal(t, 5, mapSum.Sum("ap"))
	mapSum.Insert("apple", 10)
	assert.Equal(t, 12, mapSum.Sum("ap"))
}
