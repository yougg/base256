package base256

import (
	"math"
	"reflect"
	"testing"
)

func TestDecimalToBase256(t *testing.T) {
	tests := []struct {
		name string
		num  uint64
		want []byte
	}{
		{
			name: "zero",
			num:  0,
			want: nil,
		},
		{
			name: "max 1 byte",
			num:  255,
			want: []byte{255},
		},
		{
			name: "max 2 bytes",
			num:  65535,
			want: []byte{255, 255},
		},
		{
			name: "max 3 bytes",
			num:  16777215,
			want: []byte{255, 255, 255},
		},
		{
			name: "max 4 bytes",
			num:  4294967295,
			want: []byte{255, 255, 255, 255},
		},
		{
			name: "max 5 bytes",
			num:  1099511627775,
			want: []byte{255, 255, 255, 255, 255},
		},
		{
			name: "max 6 bytes",
			num:  281474976710655,
			want: []byte{255, 255, 255, 255, 255, 255},
		},
		{
			name: "max 7 bytes",
			num:  72057594037927935,
			want: []byte{255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "max 8 bytes",
			num:  18446744073709551615,
			want: []byte{255, 255, 255, 255, 255, 255, 255, 255},
		},
		{
			name: "max int32",
			num:  math.MaxInt32,
			want: []byte{127, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChars := DecimalToBase256(tt.num); !reflect.DeepEqual(gotChars, tt.want) {
				t.Errorf("%s, DecimalToBase256() = %v, want %v", tt.name, gotChars, tt.want)
			}
		})
	}
}

func TestBase256ToDecimal(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  uint64
	}{
		{
			name:  "zero",
			chars: nil,
			want:  0,
		},
		{
			name:  "max 1 byte",
			chars: []byte{255},
			want:  255,
		},
		{
			name:  "max 2 bytes",
			chars: []byte{255, 255},
			want:  65535,
		},
		{
			name:  "max 3 bytes",
			chars: []byte{255, 255, 255},
			want:  16777215,
		},
		{
			name:  "max 4 bytes",
			chars: []byte{255, 255, 255, 255},
			want:  4294967295,
		},
		{
			name:  "max 5 bytes",
			chars: []byte{255, 255, 255, 255, 255},
			want:  1099511627775,
		},
		{
			name:  "max 6 bytes",
			chars: []byte{255, 255, 255, 255, 255, 255},
			want:  281474976710655,
		},
		{
			name:  "max 7 bytes",
			chars: []byte{255, 255, 255, 255, 255, 255, 255},
			want:  72057594037927935,
		},
		{
			name:  "max 8 bytes",
			chars: []byte{255, 255, 255, 255, 255, 255, 255, 255},
			want:  18446744073709551615,
		},
		{
			name:  "max int32",
			chars: []byte{127, 255, 255, 255},
			want:  math.MaxInt32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base256ToDecimal(tt.chars); got != tt.want {
				t.Errorf("%s, Base256ToDecimal() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkDecimalToBase256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalToBase256(uint64(i))
	}
}

func BenchmarkBase256ToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base256ToDecimal(nil)
		Base256ToDecimal([]byte{1})
		Base256ToDecimal([]byte{2, 3})
		Base256ToDecimal([]byte{4, 5, 6})
		Base256ToDecimal([]byte{7, 8, 9, 10})
		Base256ToDecimal([]byte{11, 12, 13, 14, 15})
		Base256ToDecimal([]byte{16, 17, 18, 19, 20, 21})
		Base256ToDecimal([]byte{22, 23, 24, 25, 26, 27, 28})
		Base256ToDecimal([]byte{29, 30, 31, 32, 33, 34, 35, 36})
	}
}
