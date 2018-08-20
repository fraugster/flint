package imports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	src := `
package xx

import "os" // One liner and comment 

	`

	assert.NoError(t, checker{}.Process(context.Background(), []byte(src)))

	src = `
package xx
	
import (
	"os" //  comment 
	"other"

	"other.com/xxx"
)`
	assert.NoError(t, checker{}.Process(context.Background(), []byte(src)))

	src = `
	package xx
		
import (
	"os" //  comment 
	"other" 
	"other.com/xxx" // invalid position
)`
	assert.Error(t, checker{}.Process(context.Background(), []byte(src)))

	src = `
package xx

import "multi"

import (
	"os" //  comment 
	"other" 

	"other.com/xxx" // invalid position
)`
	assert.Error(t, checker{}.Process(context.Background(), []byte(src)))

	src = `
invalid syntax 


import (
	"os" //  comment 
	"other"

	"other.com/xxx"
)`
	assert.Error(t, checker{}.Process(context.Background(), []byte(src)))

}
