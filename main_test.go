package main

import (
	"fmt"
	"testing"
)

func Test_process(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		wantErr bool
	}{
		{
			name:    "flipkart",
			rawURL:  "https://www.flipkart.com/hp-deskjet-1112-single-function-color-inkjet-printer/p/itme9wgtatxfzggg?pid=PRNE9WGTZCQGJ6PZ&lid=LSTPRNE9WGTZCQGJ6PZND3GV1&marketplace=FLIPKART&store=6bo%2Ftia%2Fffn%2Ft64&srno=b_3_98&otracker=hp_omu_Best%2Bof%2BElectronics_2_3.dealCard.OMU_D54DFY00C5JD_3&otracker1=hp_rich_navigation_PINNED_neo%2Fmerchandising_NA_NAV_EXPANDABLE_navigationCard_cc_2_L2_view-all%2Chp_omu_PINNED_neo%2Fmerchandising_Best%2Bof%2BElectronics_NA_dealCard_cc_2_NA_view-all_3&fm=neo%2Fmerchandising",
			wantErr: false,
		},
		{
			name:    "amazon",
			rawURL:  "https://www.amazon.in/Fastrack-Limitless-Biggest-SingleSync-Watchfaces/dp/B0BZ8T21V4?ref_=Oct_DLandingS_D_8dbdc968_0",
			wantErr: false,
		},
		{
			name:    "scrareme",
			rawURL:  "https://scrapeme.live/shop/Bulbasaur/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := process(tt.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
