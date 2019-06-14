package lmsr

import (
	"math"
)

// Fund ...
func Fund(liquidity float64, n int) float64 {
	return liquidity * math.Log(float64(n))
}

// Liquidity ...
func Liquidity(fund float64, n int) float64 {
	return fund / math.Log(float64(n))
}

func exp(liquidity, share float64) float64 {
	return math.Exp(share / liquidity)
}

func coefficient(liquidity float64, shares []float64) ([]float64, float64) {
	tmp := make([]float64, len(shares))
	max := 0.0
	for i, x := range shares {
		tmp[i] = x / liquidity
		if max < tmp[i] {
			max = tmp[i]
		}
	}
	return tmp, max
}

func shiftExpSum(max float64, data []float64) float64 {
	sum := 0.0

	for _, x := range data {
		sum += math.Exp(x - max)
	}

	return sum
}

func lnsum(liquidity float64, shares []float64) float64 {
	tmp, max := coefficient(liquidity, shares)

	sum := shiftExpSum(max, tmp)

	return math.Log(sum) + max

	/*sum := 0.0
	fmt.Println("shares:", len(shares))
	for _, x := range shares {
		sum += exp(liquidity, x)
	}

	return sum*/
}

// Sum LMSR sum
/*func Sum(liquidity float64, shares []float64) float64 {

	/*sum := 0.0

	for _, x := range shares {
		sum += exp(liquidity, x)
	}

	return sum

}*/

// Price share price
//func Price(liquidity, share, sum float64) float64 {
//	return exp(liquidity, share) / sum
//}

// Cost LMSR cost
func Cost(liquidity float64, shares []float64) float64 {
	return liquidity * lnsum(liquidity, shares)
}

// CmpPrice compute price for all shares
func CmpPrice(liquidity float64, shares []float64) []float64 {
	tmp, max := coefficient(liquidity, shares)
	sum := shiftExpSum(max, tmp)

	for i, x := range tmp {
		tmp[i] = math.Exp(x - max - math.Log(sum))
	}

	return tmp
	/*sum := lnsum(liquidity, shares)

	tmp := make([]float64, len(shares))

	for i, x := range shares {
		tmp[i] = exp(liquidity, x) / sum
	}

	return tmp*/
}

// Estimate estimate cost when buying or selling some shares
func Estimate(liquidity float64, org []float64, outcome int, amount float64) float64 {
	after := make([]float64, len(org))

	copy(after, org)

	after[outcome] += amount

	return Cost(liquidity, after) - Cost(liquidity, org)
}

// Volume return amount can buy how many volume
func Volume(liquidity float64, org []float64, outcome int, amount float64) float64 {
	a := math.Exp(amount/liquidity) - 1

	tmp, max := coefficient(liquidity, org)
	sum := shiftExpSum(max, tmp)

	p := math.Exp(tmp[outcome] - max - math.Log(sum))

	return liquidity * math.Log(a/p+1)

}
