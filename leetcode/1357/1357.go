package leetcode

type Cashier struct {
	catalog            *catalog
	discount           int
	discountCheckpoint int
	processedCustomer  int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	return Cashier{
		catalog:            NewCatalog(products, prices),
		discount:           discount,
		discountCheckpoint: n,
		processedCustomer:  0,
	}
}

// https://leetcode.com/problems/apply-discount-every-n-orders/

func (c *Cashier) GetBill(products []int, amount []int) float64 {
	c.processedCustomer++
	total := float64(0)
	for i, productID := range products {
		total += float64(c.catalog.GetPrice(productID) * amount[i])
	}
	if c.timeForDiscount() {
		total = total * ((100 - float64(c.discount)) / 100)
	}
	return total
}

func (c *Cashier) timeForDiscount() bool {
	return (c.processedCustomer % c.discountCheckpoint) == 0
}

type catalog struct {
	catalogContainer map[int]int
}

func (c *catalog) GetPrice(productID int) int {
	return c.catalogContainer[productID]
}

func NewCatalog(products []int, prices []int) *catalog {
	catalogContainer := map[int]int{}
	for i, product := range products {
		catalogContainer[product] = prices[i]
	}
	return &catalog{
		catalogContainer: catalogContainer,
	}
}
