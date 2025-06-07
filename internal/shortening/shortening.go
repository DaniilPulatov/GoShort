package shortening

import (
	"log"
	"net/url"
	"strings"
)

const data = "Wn4B5Kqvfx3uGJytMdL8RmgpTXYZsPzaEN72AcUQCDbFhkHeSr193Vwj"

var dataSize = uint32(len(data))

func Shorten(id uint32) string {
	var (
		nums    []uint32
		num     = id
		builder strings.Builder
	)

	for num > 0 {
		nums = append(nums, num%dataSize)
		num = num / dataSize
	}
	nums = reverse(nums)
	for _, num := range nums {
		builder.WriteString(string(data[num]))
	}

	return builder.String()
}

func reverse(nums []uint32) []uint32 {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func AddBaseUrl(baseUrl, identifier string) (string, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Println("Error parsing base url:", err)
		return "", err
	}
	parsedUrl.Path = identifier
	return parsedUrl.String(), nil
}
