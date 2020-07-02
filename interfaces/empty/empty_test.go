package empty

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	l = List{[]Any{1, "1"}}
)

func TestEmptyInterfaces(t *testing.T) {
	f := func(val Any) {
		switch v := val.(type) {
		case int:
			assert.Equal(t, reflect.Int, reflect.TypeOf(v).Kind())
		case string:
			assert.Equal(t, reflect.String, reflect.TypeOf(v).Kind())
		case bool:
			assert.Equal(t, reflect.Bool, reflect.TypeOf(v).Kind())
		default:
			t.Log("Unexpected type", v)
		}
	}
	f(1)
	f("1")
	f(true)
	f(1.0)
}

func TestList_At(t *testing.T) {
	assert.Equal(t, l.At(0), 1)
	assert.Equal(t, l.At(1), "1")
}

func TestList_Set(t *testing.T) {
	l := List{[]Any{1, "1"}}
	l.Set(0, 0)
	assert.Equal(t, l.At(0), 0)
	l.Set(1, 1)
	assert.Equal(t, l.At(1), 1)
	l.Print(t)
}

func TestList_Append(t *testing.T) {
	l := List{[]Any{1, "1"}}
	l.Append(true)
	assert.Equal(t, l.Len(), 3)
	assert.Equal(t, l.At(-1), true)
}

func TestList_Pop(t *testing.T) {
	l := List{[]Any{1, "1"}}
	assert.Equal(t, l.Pop(), "1")
	assert.Equal(t, l.Pop(), 1)
}

func TestList_Insert(t *testing.T) {
	l := List{[]Any{1, "1"}}
	l.Insert(0, 0)
	l.Insert(10, 10)
	l.Insert(1, 1)
	assert.Equal(t, l.At(0), 0)
	assert.Equal(t, l.At(1), 1)
	assert.Equal(t, l.At(2), 1)
	assert.Equal(t, l.At(3), "1")
	assert.Equal(t, l.At(4), 10)

}
