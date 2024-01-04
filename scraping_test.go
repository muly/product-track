package main

import (
	"fmt"
	"testing"
)

func Test_callScraping(t *testing.T) {
	t.Skip() // skipping because the output is not deterministic (note: we are calling with real products URLs)
	var err error
	supportedWebsites, err = readSupportedWebsites()
	if err != nil {
		t.Errorf("readSupportedWebsites() failed %v", err)
		return
	}

	tests := []struct {
		name    string
		rawURL  string
		wantErr bool
	}{
		{
			name:    "flipkart",
			rawURL:  "https://www.flipkart.com/vivo-t2-pro-5g-new-moon-black-256-gb/p/itm1230688cdef18?pid=MOBGT4RZMZFEWDY7&lid=LSTMOBGT4RZMZFEWDY7EEVQUQ&marketplace=FLIPKART&store=tyy%2F4io&srno=b_1_1&otracker=browse&fm=organic&iid=d4fd9eb7-9cb7-48a6-82db-134e1077255b.MOBGT4RZMZFEWDY7.SEARCH&ppt=hp&ppn=homepage&ssid=oj2jvrt9ls0000001699124880241",
			wantErr: false,
		},
		// {
		// 	name:    "amazon",
		// 	rawURL:  "https://www.amazon.in/Fastrack-Limitless-Biggest-SingleSync-Watchfaces/dp/B0BZ8T21V4?ref_=Oct_DLandingS_D_8dbdc968_0",
		// 	wantErr: false,
		// },
		// {
		// 	name:    "scrapeme",
		// 	rawURL:  "https://scrapeme.live/shop/Bulbasaur/",
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := callScraping(tt.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
