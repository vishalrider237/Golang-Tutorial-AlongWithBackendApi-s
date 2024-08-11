package url

import (
	"fmt"
	"net/url"
)

func URL() {
	myurl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	paredUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Can not parse url!!,", err)
		return
	}
	fmt.Println("Scheme:", paredUrl.Scheme)
	fmt.Println("Host:", paredUrl.Host)
	fmt.Println("Path:", paredUrl.Path)
	fmt.Println("Raw Query:", paredUrl.RawQuery)
}
