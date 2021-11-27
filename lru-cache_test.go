package lru_cache

import (
	"container/list"
	"testing"
)

func TestCache_Add(t *testing.T) {
	type fields struct {
		size  int
		elems *list.List
		items map[interface{}]*list.Element
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		addOk  bool
		hasOk  bool
	}{
		// TODO: Add test cases.
		{
			name: "add one",
			fields: fields{
				size:  10,
				elems: list.New(),
				items: make(map[interface{}]*list.Element),
			},
			args:  args{key: 1, value: "one"},
			addOk: true,
			hasOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				size:  tt.fields.size,
				elems: tt.fields.elems,
				items: tt.fields.items,
			}
			if hasOk := c.Exists(tt.args.key); hasOk == tt.hasOk {
				t.Errorf("before Exists() = %v, want %v", hasOk, !tt.hasOk)
			}

			if addOk := c.Add(tt.args.key, tt.args.value); addOk != tt.addOk {
				t.Errorf("Add() = %v, want %v", addOk, tt.addOk)
			}

			if hasOk := c.Exists(tt.args.key); hasOk != tt.hasOk {
				t.Errorf("after Exists() = %v, want %v", hasOk, tt.hasOk)
			}
		})
	}
}
