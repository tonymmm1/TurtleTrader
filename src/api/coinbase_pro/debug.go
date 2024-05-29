package coinbase_pro

import (
	"fmt"

	"turtle/src/config"

	"github.com/go-resty/resty/v2"
)

func rest_debug(resp *resty.Response, err error) { //handles GET requests
	if config.Debug {
		if err != nil {
			fmt.Println("Error occurred during GET request:", err)
		} else {
			fmt.Println("Response Info:")
			fmt.Println("  Error      :", err)
			fmt.Println("  Status Code:", resp.StatusCode())
			fmt.Println("  Status     :", resp.Status())
			fmt.Println("  Proto      :", resp.Proto())
			fmt.Println("  Time       :", resp.Time())
			fmt.Println("  Received At:", resp.ReceivedAt())
			fmt.Println("  Body       :\n", resp)
			fmt.Println()
		}
	}
}
