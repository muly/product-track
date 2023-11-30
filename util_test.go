package main

import (
	"net/url"
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

func Test_cleanupURL(t *testing.T) {
	amazonUrl, _ := url.Parse("https://www.amazon.in/Pixel-Obsdian-8GB-128GB-Storage/dp/B0CK86X65L/ref=sr_1_1?crid=G92I6EEBL2EC&keywords=pixel+7&qid=1701311235&s=electronics&sprefix=pixel+%2Celectronics%2C140&sr=1-1")
	amazonCleanUrl, _ := url.Parse("https://www.amazon.in/Pixel-Obsdian-8GB-128GB-Storage/dp/B0CK86X65L")

	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name:    "amazon",
			args:    args{u: amazonUrl},
			want:    amazonCleanUrl,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cleanupURL(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("cleanupURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanupURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
