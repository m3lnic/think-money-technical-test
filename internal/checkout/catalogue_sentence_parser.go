package checkout

import (
	"errors"
	"regexp"
	"strconv"
)

var ErrInvalidInput error = errors.New("invalid input")

// > Returns
//
//		> Item name
//		> Cost
//		> Quantity
//	 	> If quantity > 0 then it's a discount
func ParseCatalogueItemOrDiscountFromSentence(sentence string) (string, int, int, error) {
	re := regexp.MustCompile(`^(?:(\d+)\s+)?([\w\s]+)\s+cost\s+(\d+)$`)
	matches := re.FindStringSubmatch(sentence)

	if len(matches) == 0 {
		return "", 0, 0, ErrInvalidInput
	}

	quantity := 0
	if matches[1] != "" {
		val, _ := strconv.ParseInt(matches[1], 10, 64)
		quantity = int(val)
	}

	cost, _ := strconv.ParseInt(matches[3], 10, 64)

	return matches[2], int(cost), quantity, nil
}
