package caesar

import (
	"cs50/problems/week_2/caesar"
	"testing"
)

func TestCaesar(t *testing.T) {
	type args struct {
		plaintText string
		key        int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"encrypts \"a\" as \"b\" using 1 as key", args{"a", 1}, "ciphertext: b"},
		{"encrypts \"barfoo\" as \"yxocll\" using 23 as key", args{"barfoo", 23}, "ciphertext: yxocll"},
		{"encrypts \"BARFOO\" as \"EDUIRR\" using 3 as key", args{"BARFOO", 3}, "ciphertext: EDUIRR"},
		{"encrypts \"BaRFoo\" as \"FeVJss\" using 4 as key", args{"BaRFoo", 4}, "ciphertext: FeVJss"},
		{"encrypts \"barfoo\" as \"onesbb\" using 65 as key", args{"barfoo", 65}, "ciphertext: onesbb"},
		{"encrypts \"world, say hello!\" as \"iadxp, emk tqxxa!\" using 12 as key", args{"world, say hello!", 12}, "ciphertext: iadxp, emk tqxxa!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar.Caesar(tt.args.plaintText, tt.args.key); got != tt.want {
				t.Errorf("Caesar() = %v, want %v", got, tt.want)
			}
		})
	}
}
