package plantuml_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/candy12t/go-plantuml"
)

func TestEncode(t *testing.T) {
	tests := map[string]struct {
		in   io.Reader
		want string
	}{
		"sample1": {
			in: func() io.Reader {
				f, err := os.Open("./misc/sample1.pu")
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				return f
			}(),
			want: "SYWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vt98pKi1IG82003__m==",
		},
		"sample2": {
			in: func() io.Reader {
				f, err := os.Open("./misc/sample2.pu")
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				return f
			}(),
			want: "Z9JFJy8m5CVl_IiQF7FmPEHWs3f8J8Wn4kFviJwnsTQErpSH_tfJ8DADmJZ2FjzVRyccde6ugKhX2sDh8AZa2l9YJQwnMhdIaRoRpMPffYBY2wpUac56AvaQ5D4pZvi6ROu9aTiU33B4UbdiqhB1FZ1dHwaZZNGBlZ2Vk30MOyLg0EqEIXupbSrx5Az0R79JW-NR6yMYJbBcz1ffM3Ttbb-WGlbSrP3pCBmqloZl7uR1eSM7wtTgmwQ1I-p9z8RNUhtgFxAVda2DFq90-5E-UoHHdwR8qToGPwbAy7uamSjePz8cbvWxg_lHj8qku9Ad4Y9qaEEdLj94Pkx3KH5gcczWRFSGyQ-EDfr8HHWaRs6_vrwFjEMHTZpKrNmLIvKSd-K7PUOi2esUoUqwWzl1PykFjOyhjD0uF23P-uHBXkRxvFxBJT7gy1dw2m00__y=",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := plantuml.Encode(&buf, tt.in); err != nil {
				t.Errorf("Encode() error = %v", err)
				return
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := map[string]struct {
		in   io.Reader
		want string
	}{
		"sample1": {
			in: bytes.NewBufferString("SYWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vt98pKi1IG82003__m=="),
			want: func() string {
				f, err := os.Open("./misc/sample1.pu")
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, f); err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
				return buf.String()
			}(),
		},
		"sample2": {
			in: bytes.NewBufferString("Z9JFJy8m5CVl_IiQF7FmPEHWs3f8J8Wn4kFviJwnsTQErpSH_tfJ8DADmJZ2FjzVRyccde6ugKhX2sDh8AZa2l9YJQwnMhdIaRoRpMPffYBY2wpUac56AvaQ5D4pZvi6ROu9aTiU33B4UbdiqhB1FZ1dHwaZZNGBlZ2Vk30MOyLg0EqEIXupbSrx5Az0R79JW-NR6yMYJbBcz1ffM3Ttbb-WGlbSrP3pCBmqloZl7uR1eSM7wtTgmwQ1I-p9z8RNUhtgFxAVda2DFq90-5E-UoHHdwR8qToGPwbAy7uamSjePz8cbvWxg_lHj8qku9Ad4Y9qaEEdLj94Pkx3KH5gcczWRFSGyQ-EDfr8HHWaRs6_vrwFjEMHTZpKrNmLIvKSd-K7PUOi2esUoUqwWzl1PykFjOyhjD0uF23P-uHBXkRxvFxBJT7gy1dw2m00__y="),
			want: func() string {
				f, err := os.Open("./misc/sample2.pu")
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, f); err != nil {
					t.Fatalf("failed to read file: %v", err)
				}
				return buf.String()
			}(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := plantuml.Decode(&buf, tt.in); err != nil {
				t.Errorf("Decode() error = %v", err)
				return
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
