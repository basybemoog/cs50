package substitution

import (
	"cs50/problems/week_2/substitution"
	"testing"
)

const CS50Answer = "ciphertext: Cbah ah KH50"
const CS50Question = "This is CS50"
const DUPLICATEANSWER = "duplicate characters"

func TestSubstitution(t *testing.T) {
	type args struct {
		key  string
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"encrypts \"A\" as \"Z\" using ZYXWVUTSRQPONMLKJIHGFEDCBA as key", args{"ZYXWVUTSRQPONMLKJIHGFEDCBA", "A"}, "ciphertext: Z"},
		{"encrypts \"a\" as \"z\" using ZYXWVUTSRQPONMLKJIHGFEDCBA as key", args{"ZYXWVUTSRQPONMLKJIHGFEDCBA", "a"}, "ciphertext: z"},
		{"encrypts \"ABC\" as \"NJQ\" using NJQSUYBRXMOPFTHZVAWCGILKED as key", args{"NJQSUYBRXMOPFTHZVAWCGILKED", "ABC"}, "ciphertext: NJQ"},
		{"encrypts \"XyZ\" as \"KeD\" using NJQSUYBRXMOPFTHZVAWCGILKED as key", args{"NJQSUYBRXMOPFTHZVAWCGILKED", "XyZ"}, "ciphertext: KeD"},
		{"encrypts \"This is CS50\" as \"Cbah ah KH50\" using YUKFRNLBAVMWZTEOGXHCIPJSQD as key", args{"YUKFRNLBAVMWZTEOGXHCIPJSQD", CS50Question}, CS50Answer},
		{"encrypts \"This is CS50\" as \"Cbah ah KH50\" using yukfrnlbavmwzteogxhcipjsqd as key", args{"yukfrnlbavmwzteogxhcipjsqd", CS50Question}, CS50Answer},
		{"encrypts \"This is CS50\" as \"Cbah ah KH50\" using YUKFRNLBAVMWZteogxhcipjsqd as key", args{"YUKFRNLBAVMWZteogxhcipjsqd", CS50Question}, CS50Answer},
		{"encrypts all alphabetic characters using DWUSXNPQKEGCZFJBTLYROHIAVM as key", args{"DWUSXNPQKEGCZFJBTLYROHIAVM", "The quick brown fox jumps over the lazy dog"}, "ciphertext: Rqx tokug wljif nja eozby jhxl rqx cdmv sjp"},
		{"does not encrypt non-alphabetical characters using DWUSXNPQKEGCZFJBTLYROHIAVM as key", args{"DWUSXNPQKEGCZFJBTLYROHIAVM", "Shh... Don't tell!"}, "ciphertext: Yqq... Sjf'r rxcc!"},
		{"handles lack of key", args{"", ""}, "Key must contain 26 characters"},
		{"handles invalid key length", args{"QTXDGMKIPV", ""}, "Key must contain 26 characters"},
		{"handles invalid characters in key", args{"ZWGKPMJ^YISHFEXQON[DLUACVT", ""}, "Invalid characters in key"},
		{"handles duplicate characters in uppercase key", args{"FAZRDTMGQEJPWAXUSKVIYCLONH", ""}, DUPLICATEANSWER},
		{"handles duplicate characters in lowercase key", args{"fazrdtmgqejpwaxuskviyclonh", ""}, DUPLICATEANSWER},
		{"handles multiple duplicate characters in key", args{"MMCcEFGHIJKLMNOPqRqTUVWXeZ", ""}, DUPLICATEANSWER},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substitution.Substitution(tt.args.key, tt.args.word); got != tt.want {
				t.Errorf("Substitution() = %v, want %v", got, tt.want)
			}
		})
	}
}
