package main

import (
	"reflect"
	"testing"

	scrape "github.com/muly/product-scrape"
)

func Test_shouldNotify(t *testing.T) {
	type args struct {
		i trackInput
		p scrape.Product
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "checking product availability",
			args: args{i: trackInput{TypeOfRequest: "AVAILABILITY"},
				p: scrape.Product{Availability: true}},
			want: true,
		},
		{
			name: "checking product unavailability",
			args: args{i: trackInput{TypeOfRequest: "AVAILABILITY"},
				p: scrape.Product{Availability: false}},
			want: false,
		},
		{
			name: "less than minimum threshold limit ",
			args: args{i: trackInput{TypeOfRequest: "PRICE",
				MinThreshold: 700.0000},
				p: scrape.Product{Price: 699.99999},
			},
			want: true,
		},
		{
			name: "greater than minimum threshold limit",
			args: args{i: trackInput{TypeOfRequest: "PRICE",
				MinThreshold: 800.000},
				p: scrape.Product{Price: 19000.087786},
			},
			want: false,
		},
		{
			name: "empty product",
			args: args{i: trackInput{TypeOfRequest: "PRICE",
				MinThreshold: 800.000},
				p: scrape.Product{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldNotify(tt.args.i, tt.args.p); got != tt.want {
				t.Errorf("shouldNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readSupportedWebsites(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "happy case",
			want: map[string]bool{
				"scrapeme.live":    true,
				"www.flipkart.com": true,
				"www.amazon.in":    true,
				"mkp.gem.gov.in":   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readSupportedWebsites()
			if (err != nil) != tt.wantErr {
				t.Errorf("readSupportedWebsites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readSupportedWebsites() = %v, want %v", got, tt.want)
			}
		})
	}
}
