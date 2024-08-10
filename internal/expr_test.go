package internal

import (
	"testing"
	"time"
)

func TestWithin(t *testing.T) {
	tests := []struct {
		dt      time.Time
		durstr  string
		want    bool
		wantErr bool
	}{
		{
			dt:      time.Now(),
			durstr:  "-1h",
			want:    true,
			wantErr: false,
		},
		{
			dt:      time.Now().Add(-2 * time.Hour),
			durstr:  "-1h",
			want:    false,
			wantErr: false,
		},
		{
			dt:      time.Now().Add(-2 * time.Hour),
			durstr:  "-3h",
			want:    true,
			wantErr: false,
		},
		{
			dt:      time.Now(),
			durstr:  "invalid_duration",
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := Within(tt.dt, tt.durstr)
		if (err != nil) != tt.wantErr {
			t.Errorf("Within(%v, %v) error = %v, wantErr %v", tt.dt, tt.durstr, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("Within(%v, %v) = %v, want %v", tt.dt, tt.durstr, got, tt.want)
		}
	}
}
