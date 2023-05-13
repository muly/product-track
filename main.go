package main

import (
	"fmt"

	"net/url"
)

func main() {
	rawURL := "https://www.flipkart.com/hp-deskjet-1112-single-function-color-inkjet-printer/p/itme9wgtatxfzggg?pid=PRNE9WGTZCQGJ6PZ&lid=LSTPRNE9WGTZCQGJ6PZND3GV1&marketplace=FLIPKART&store=6bo%2Ftia%2Fffn%2Ft64&srno=b_3_98&otracker=hp_omu_Best%2Bof%2BElectronics_2_3.dealCard.OMU_D54DFY00C5JD_3&otracker1=hp_rich_navigation_PINNED_neo%2Fmerchandising_NA_NAV_EXPANDABLE_navigationCard_cc_2_L2_view-all%2Chp_omu_PINNED_neo%2Fmerchandising_Best%2Bof%2BElectronics_NA_dealCard_cc_2_NA_view-all_3&fm=neo%2Fmerchandising"

	p, err := process(rawURL)
	if err != nil {
		fmt.Println("error occurred while scraping", err)
		return
	}

	fmt.Printf("%+v", p)
}

func process(rawURL string) (product, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return product{}, err
	}

	switch u.Hostname() {
	case "scrapeme.live":
		return scrapeme(rawURL)
	case "www.flipkart.com":
		return flipkart(rawURL)
	case "www.amazon.in":
		return amazon(rawURL)
	default:
		fmt.Printf("%s is not supported\n", u.Hostname())
		return product{}, err
	}
}
