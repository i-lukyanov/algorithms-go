package common

import (
	"testing"
)

func TestStringToFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				s: "98.765432",
			},
			want:    98.765432,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				s: "qwer",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	type args struct {
		num   float64
		point int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Case_2_point",
			args{
				num:   12.34567890,
				point: 2,
			},
			"12.35",
		},
		{
			"Case_6_point",
			args{
				num:   12.34567890,
				point: 6,
			},
			"12.345679",
		},
		{
			"Case_7_point",
			args{
				num:   12.345678912,
				point: 7,
			},
			"12.3456789",
		},
		{
			"Case_10_point",
			args{
				num:   12.345678912354,
				point: 10,
			},
			"12.3456789124",
		},
		{
			"Case_default",
			args{
				num:   12.34567890,
				point: 4,
			},
			"12.35",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.num, tt.args.point); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
