package standardpaths

import (
	"os"
	"testing"
)

func Test_appName(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{
			want: "standardpaths.test",
		},
		{
			args: []string{
				"some-arg-0",
				"some-arg-1",
			},
			want: "some-arg-0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if tt.args != nil {
				oldArgs := os.Args
				os.Args = tt.args
				defer func() { os.Args = oldArgs }()
			}

			if got := appName(); got != tt.want {
				t.Errorf("appName() = %v, want %v", got, tt.want)
			}
		})
	}
}
