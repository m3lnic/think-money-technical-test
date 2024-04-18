package checkout

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
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

type ICatalogueSentenceParser interface {
	Parse(string) error
}

func NewCatalogueSentenceParser(catalogue ICatalogueRepository, discountCatalogue IDiscountCatalogueRepository) ICatalogueSentenceParser {
	return &catalogueSentenceParser{
		catalogue:         catalogue,
		discountCatalogue: discountCatalogue,
	}
}

type catalogueSentenceParser struct {
	catalogue         ICatalogueRepository
	discountCatalogue IDiscountCatalogueRepository
}

type TempParsedData struct {
	Name     string
	Price    int
	Quantity int
}

func (csp *catalogueSentenceParser) processParsedData(data TempParsedData) error {
	sku, _, err := csp.catalogue.GetByItemName(data.Name)
	if err != nil {
		return err
	}

	// > Item
	if data.Quantity == 0 {
		csp.catalogue.Update(sku, NewItem(data.Name, data.Price))
		return nil
	}

	// > Discount
	newDiscount := NewDiscount(data.Quantity, data.Price)
	if _, err := csp.discountCatalogue.Read(sku); err == nil {
		// > We could update the fetched discount struct and pass that
		// > But this is faster
		// > Bad for GC though
		csp.discountCatalogue.Update(sku, newDiscount)
		return nil
	}

	csp.discountCatalogue.Create(sku, newDiscount)

	return nil
}

// Parse implements ICatalogueSentenceParser.
func (csp *catalogueSentenceParser) Parse(sentence string) error {
	re := regexp.MustCompile(`[,.]+`)
	sentences := re.Split(sentence, -1)

	toProcess := make([]TempParsedData, 0)
	for _, current := range sentences {
		cleaned := strings.TrimSpace(current)
		if cleaned == "" {
			continue
		}

		name, price, quantity, err := ParseCatalogueItemOrDiscountFromSentence(cleaned)
		if err != nil {
			return err
		}

		toProcess = append(toProcess, TempParsedData{
			Name:     name,
			Price:    price,
			Quantity: quantity,
		})
	}

	for _, current := range toProcess {
		if err := csp.processParsedData(current); err != nil {
			return err
		}
	}

	return nil
}
