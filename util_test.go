package main

import (
	"testing"
)

func Test_shouldNotify(t *testing.T) {
	type args struct {
		i trackInput
		p product
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "checking product availability",
			args: args{i: trackInput{TypeOfRequest: "AVAILABILITY"},
				p: product{Availability: true}},
			want: true,
		},
		{
			name: "checking product unavailability",
			args: args{i: trackInput{TypeOfRequest: "AVAILABILITY"},
				p: product{Availability: false}},
			want: false,
		},
		{
			name: "less than minimum threshold limit ",
			args: args{i: trackInput{TypeOfRequest: "PRICE",
				MinThreshold: 700.0000},
				p: product{Price: 699.99999},
			},
			want: true,
		},
		{
			name: "greater than minimum threshold limit",
			args: args{i: trackInput{TypeOfRequest: "PRICE",
				MinThreshold: 800.000},
				p: product{Price: 19000.087786},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notifyConditions(tt.args.i, tt.args.p); got != tt.want {
				t.Errorf("shouldNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPrice(t *testing.T) {
	type args struct {
		price string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "non number",
			args:    args{"rohith"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "number with commas",
			args:    args{"1,899,99"},
			want:    189999,
			wantErr: false,
		},
		{
			name:    "number with ₹",
			args:    args{"₹1,899,99"},
			want:    189999,
			wantErr: false,
		},
		{
			name:    "number with $",
			args:    args{"$1,899,99"},
			want:    189999,
			wantErr: false,
		},
		{
			name:    "number",
			args:    args{"8000"},
			want:    8000,
			wantErr: false,
		},
		{
			name:    "number with decimals",
			args:    args{"1899.999"},
			want:    1899.999,
			wantErr: false,
		},
		{
			name:    "number with decimals and commas",
			args:    args{"1,899.999"},
			want:    1899.999,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := priceConvertor(tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkAvailability(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "product available",
			args: args{"In stock"},
			want: true,
		},
		{
			name: "less number of products available",
			args: args{"Hurry, only 4 items left!"},
			want: true,
		},
		{
			name: "less number of products available: regex without comma",
			args: args{"Hurry only 4 items left!"},
			want: true,
		},
		{
			name: "product unavailable",
			args: args{"Currently unavailable"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAvailability(tt.args.s); got != tt.want {
				t.Errorf("checkAvailability() = %v, want %v", got, tt.want)
			}
		})
	}
}
