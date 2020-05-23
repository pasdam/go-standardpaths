package standardpaths

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func TestAppData(t *testing.T) {
	type mocks struct {
		dir string
		err error
	}
	tests := []struct {
		name    string
		mocks   mocks
		want    string
		wantErr bool
	}{
		{
			name: "Should return correct dir",
			mocks: mocks{
				dir: "some-dir",
			},
			want: filepath.Join("some-dir", "standardpaths.test", "data"),
		},
		{
			name: "Should propagate error if os.UserConfigDir raises it",
			mocks: mocks{
				err: errors.New("some-error"),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockit.MockFunc(t, os.UserConfigDir).With().Return(tt.mocks.dir, tt.mocks.err)

			got, err := AppData()

			assert.Equal(t, tt.mocks.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
