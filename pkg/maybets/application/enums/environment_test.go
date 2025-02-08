package enums

import (
	"testing"
)

func TestEnvironment_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    Environment
		want bool
	}{
		{
			name: "success: valid enum",
			e:    Local,
			want: true,
		},
		{
			name: "fail: invalid enum",
			e:    Environment("invalid"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("Environment.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironment_String(t *testing.T) {
	tests := []struct {
		name string
		e    Environment
		want string
	}{
		{
			name: "success: convert to string",
			e:    Local,
			want: "LOCAL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Environment.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
