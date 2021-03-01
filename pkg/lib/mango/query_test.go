package mango

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/go-playground/assert.v1"
)

func TestQuery_Equal(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key equal to value",
			args: args{
				key:   "key",
				value: "value",
			},
			want: bson.M{
				"key": "value",
			},
		},
	}

	for _, tt := range tests {
		q := NewQuery()

		q.Equal(tt.args.key, tt.args.value)

		assert.Equal(t, q.BSON(), tt.want)
	}
}

func TestQuery_In(t *testing.T) {
	type args struct {
		key    string
		values []interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key in a set of values",
			args: args{
				key:    "key",
				values: []interface{}{"value1", "value2"},
			},
			want: bson.M{
				"key": bson.M{"$in": []interface{}{"value1", "value2"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q.In(tt.args.key, tt.args.values...)
			assert.Equal(t, q.BSON(), tt.want)
		})
	}
}

func TestQuery_GreaterThan(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key greater than value",
			args: args{
				key:   "key",
				value: "value",
			},
			want: bson.M{
				"key": bson.M{"$gt": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q.GreaterThan(tt.args.key, tt.args.value)
			assert.Equal(t, q.BSON(), tt.want)
		})
	}
}

func TestQuery_GreaterThanEqual(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key greater than/equal to value",
			args: args{
				key:   "key",
				value: "value",
			},
			want: bson.M{
				"key": bson.M{"$gte": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q.GreaterThanEqual(tt.args.key, tt.args.value)
			assert.Equal(t, q.BSON(), tt.want)
		})
	}
}

func TestQuery_LessThan(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key smaller than value",
			args: args{
				key:   "key",
				value: "value",
			},
			want: bson.M{
				"key": bson.M{"$lt": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q.LessThan(tt.args.key, tt.args.value)
			assert.Equal(t, q.BSON(), tt.want)
		})
	}
}

func TestQuery_LessThanEqual(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "key smaller than/equal to value",
			args: args{
				key:   "key",
				value: "value",
			},
			want: bson.M{
				"key": bson.M{"$lte": "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQuery()
			q.LessThanEqual(tt.args.key, tt.args.value)
			assert.Equal(t, q.BSON(), tt.want)
		})
	}
}
