package storm

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	dir, _ := ioutil.TempDir(os.TempDir(), "storm")
	defer os.RemoveAll(dir)
	db, _ := Open(filepath.Join(dir, "storm.db"))

	u := UniqueNameUser{Name: "John", ID: 10}
	err := db.Save(&u)
	assert.NoError(t, err)

	v := UniqueNameUser{}
	err = db.One("Name", "John", &v)
	assert.NoError(t, err)
	assert.Equal(t, u, v)

	for i := 0; i < 10; i++ {
		w := IndexedNameUser{Name: "John", ID: i + 1}
		err := db.Save(&w)
		assert.NoError(t, err)
	}

	x := IndexedNameUser{}
	err = db.One("Name", "John", &x)
	assert.NoError(t, err)
	assert.Equal(t, "John", x.Name)
	assert.Equal(t, 1, x.ID)
	assert.Zero(t, x.age)
	assert.True(t, x.DateOfBirth.IsZero())

	err = db.One("Name", "Mike", &x)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)

	err = db.One("", nil, &x)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)

	err = db.One("", "Mike", nil)
	assert.Error(t, err)
	assert.Equal(t, ErrStructPtrNeeded, err)

	err = db.One("", nil, nil)
	assert.Error(t, err)
	assert.Equal(t, ErrStructPtrNeeded, err)

	y := UniqueNameUser{Name: "Jake", ID: 200}
	err = db.Save(&y)
	assert.NoError(t, err)

	var y2 UniqueNameUser
	err = db.One("ID", 200, &y2)
	assert.NoError(t, err)
	assert.Equal(t, y, y2)

	n := NestedID{}
	n.ID = "100"
	n.Name = "John"

	err = db.Save(&n)
	assert.NoError(t, err)

	var n2 NestedID
	err = db.One("ID", "100", &n2)
	assert.NoError(t, err)
	assert.Equal(t, n, n2)
}
