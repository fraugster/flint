package imports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultImport(t *testing.T) {
	imp := `import (
	"os"
	"another"

	xx "external.com/ooo"
	// comment
	"external.com/gsggs" // comment
	. "dot.com/import"
)`

	assert.NoError(t, multiLiner(imp))

	imp = `import (
	"os"

	"another"

	"external.com/ooo"
	uu "external.com/gsggs"
	)`

	assert.Error(t, multiLiner(imp))

	imp = `import (
	"external.com/ooo"
	"external.com/gsggs"

	"os"
	"another"
)`

	assert.Error(t, multiLiner(imp))

	imp = `import (

	"external.com/ooo"
	"external.com/gsggs"

	"os"
	"another"
)`

	assert.Error(t, multiLiner(imp))

	imp = `import (
	"os"
	"another"

	"external.com/ooo"

	"external.com/gsggs"
)`
	assert.Error(t, multiLiner(imp))

	imp = `import (
	"os"
	"another"
	
	"external.com/ooo"
	    "external.com/gsggs" //Extra space
)`
	assert.Error(t, multiLiner(imp))

	imp = `import (
	"os"
	"another" // comment 
	"external.com/ooo"
	"external.com/gsggs" // Comment
)`

	assert.Error(t, multiLiner(imp))
}

func TestOneLiner(t *testing.T) {
	imp := `import correct "os"`
	assert.NoError(t, oneLiner(imp))

	imp = `import "os"`
	assert.NoError(t, oneLiner(imp))

	imp = `import ""`
	assert.Error(t, oneLiner(imp))

	imp = `import   "os"`
	assert.Error(t, oneLiner(imp))
}
