package enums

import (
	"testing"
)

func TestOutcome_IsValid(t *testing.T) {
	tests := []struct {
		name string
		o    Outcome
		want bool
	}{
		{
			name: "success: valid enum",
			o:    Win,
			want: true,
		},
		{
			name: "fail: invalid enum",
			o:    Outcome("invalid"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsValid(); got != tt.want {
				t.Errorf("Outcome.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutcome_String(t *testing.T) {
	tests := []struct {
		name string
		o    Outcome
		want string
	}{
		{
			name: "success: output string",
			o:    Win,
			want: "win",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("Outcome.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
