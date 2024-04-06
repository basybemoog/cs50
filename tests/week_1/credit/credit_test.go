package credit

import (
	"cs50/problems/week_1/credit"
	"testing"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"identifies 378282246310005 as AMEX", 378282246310005, "AMEX"},
		{"identifies 371449635398431 as AMEX", 371449635398431, "AMEX"},
		{"identifies 5555555555554444 as MASTERCARD", 5555555555554444, "MASTERCARD"},
		{"identifies 5105105105105100 as MASTERCARD", 5105105105105100, "MASTERCARD"},
		{"identifies 4111111111111111 as VISA", 4111111111111111, "VISA"},
		{"identifies 4012888888881881 as VISA", 4012888888881881, "VISA"},
		{"identifies 4222222222222 as VISA", 4222222222222, "VISA"},
		{"identifies 1234567890 as INVALID (invalid length, checksum, identifying digits)", 1234567890, "INVALID"},
		{"identifies 369421438430814 as INVALID (invalid identifying digits)", 369421438430814, "INVALID"},
		{"identifies 4062901840 as INVALID (invalid length)", 4062901840, "INVALID"},
		{"identifies 5673598276138003 as INVALID (invalid identifying digits)", 5673598276138003, "INVALID"},
		{"identifies 4111111111111113 as INVALID (invalid checksum)", 4111111111111113, "INVALID"},
		{"identifies 4222222222223 as INVALID (invalid checksum)", 4222222222223, "INVALID"},
		{"identifies 3400000000000620 as INVALID (AMEX identifying digits, VISA/Mastercard length)", 3400000000000620, "INVALID"},
		{"identifies 430000000000000 as INVALID (VISA identifying digits, AMEX length)", 430000000000000, "INVALID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := credit.Validate((tt.input))
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
