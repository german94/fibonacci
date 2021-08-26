package fibcalc

import "math/big"

var cache map[string]*big.Int

func init() {
	cache = make(map[string]*big.Int)
}

type FibonacciCalculator struct {
	cache map[string]*big.Int
}

// Creates a new FibonacciCalculator, ready for getting new fibonacci
// terms.
func NewFibonacciCalculator() *FibonacciCalculator {
	return &FibonacciCalculator{cache: make(map[string]*big.Int)}
}

// GetTerm performs the fibonacci calculation of the n term.
// Already calculated terms are cached in memory.
func (c *FibonacciCalculator) GetTerm(n *big.Int) *big.Int {
	strNumber := n.String()
	if c.cache[strNumber] != nil {
		return c.cache[strNumber]
	}
	if n.Cmp(big.NewInt(2)) < 0 {
		c.cache[strNumber] = big.NewInt(0).Set(n)
		return c.cache[strNumber]
	}
	term1 := c.GetTerm(big.NewInt(0).Sub(n, big.NewInt(1)))
	term2 := c.GetTerm(big.NewInt(0).Sub(n, big.NewInt(2)))

	c.cache[strNumber] = big.NewInt(0).Add(term1, term2)

	return c.cache[strNumber]
}
