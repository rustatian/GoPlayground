package main

import (
	"bitbucket.org/inturnco/go-sdk/types"
	"reflect"
	"testing"
)

var mmm map[string][]string
var k = "keeey"

func init() {
	mmm = make(map[string][]string, 10000)
}

func Benchmark_Bad(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i <= b.N; i++ {
		res := Bad(mmm, k)
		_ = res
	}
}

func Benchmark_Good(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i <= b.N; i++ {
		res := Good(mmm, k)
		_ = res
	}
}

//go:noinline
func Bad(m map[string][]string, key string) *string {
	val := m[key]
	if len(val) == 0 {
		return nil
	}

	str := val[0]
	if str == "" {
		return nil

	}
	return types.String(str)
}

//go:noinline
func Good(m map[string][]string, key string) *string {
	if val, ok := m[key]; ok {
		if len(val) > 0 && val[0] != "" {
			return types.String(val[0])

		}
	}

	return nil
}

func TestFirstStringOrNilFromMap(t *testing.T) {

	type args struct {

		m   map[string][]string

		key string

	}

	tests := []struct {

		name string

		args args

		want *string

	}{

		{

			name: "returns pointer to value",

			args: args{

				m:   map[string][]string{"companyid": []string{"0000-000-000"}},

				key: "companyid",

			},

			want: types.String("0000-000-000"),

		},

		{

			name: "returns nil if key doesn't exist",

			args: args{

				m:   map[string][]string{"companyid": []string{"0000-000-000"}},

				key: "locked",

			},

			want: nil,

		},

		{

			name: "returns nil if value doesn't exist",

			args: args{

				m:   map[string][]string{"companyid": []string{""}},

				key: "companyid",

			},

			want: nil,

		},

	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := Good(tt.args.m, tt.args.key); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("StringOrNilFromMap() = %v, want %v", got, tt.want)

			}

		})

	}

}