package main

import (
	"reflect"
	"testing"
)

func Test_scrape(t *testing.T) {
	t.Skip() // WARNING: enable this unit test only for manual trigger of testing after updating the expected price values right before running the tests. Note: the price and availability is not deterministic.
	type args struct {
		url string
		t   scrapeTags
	}
	tests := []struct {
		name    string
		args    args
		want    product
		wantErr bool
	}{
		{
			name: "debug amazon.in",
			args: args{
				url: "https://www.amazon.in/Vaseline-Intensive-Restore-Lasting-Moisturisation/dp/B08M4QF7V9",
				t: scrapeTags{
					availability: "#availability",
					price:        "div.a-section.a-spacing-none.aok-align-center",
					priceChild:   "span.a-price-whole",
				},
			},
			want: product{
				Url:          "https://www.amazon.in/Vaseline-Intensive-Restore-Lasting-Moisturisation/dp/B08M4QF7V9",
				Price:        245,
				Availability: true,
			},
			wantErr: false,
		},
		{
			name: "debug flipkart.com",
			args: args{
				url: "https://www.flipkart.com/sony-alpha-ilce-6100y-aps-c-mirrorless-camera-dual-lens-16-50-mm-55-210-zoom-featuring-eye-af-4k-movie-recording/p/itmf4d10371a5240",
				t: scrapeTags{
					availability: "div._3XINqE",
					price:        "div._30jeq3._16Jk6d",
					priceChild:   "",
				},
			},
			want: product{
				Url:          "https://www.flipkart.com/sony-alpha-ilce-6100y-aps-c-mirrorless-camera-dual-lens-16-50-mm-55-210-zoom-featuring-eye-af-4k-movie-recording/p/itmf4d10371a5240",
				Price:        78990,
				Availability: true,
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scrape(tt.args.url, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("scrape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrape() = %v, want %v", got, tt.want)
			}
		})
	}
}
