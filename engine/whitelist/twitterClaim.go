package whitelist

import "inverse.so/services"

func ProcessTwitterCallback(token, verifier *string) {

	_, err := services.FetchTwitterAccessToken(token, verifier)
	if err != nil {
		panic(err)
	}

	// do other claim specific stuff here
}
