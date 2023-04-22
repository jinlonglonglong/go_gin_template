package util

import "testing"

func TestGetDayOfMonth(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "default",
			want: "13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayOfMonthString(); got != tt.want {
				t.Errorf("GetDayOfMonthString() = %v, want %v", got, tt.want)
			}
		})
	}
}
