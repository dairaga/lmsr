package lmsr

import (
	"math"
	"testing"
)

func TestFundLiquidity(t *testing.T) {
	t.Log("fund:", Fund(100, 2))
	t.Log("liquid:", Liquidity(69.31471805599453, 2))

	var shares0 [2]float64
	var shares1 [2]float64

	// init
	prices := CmpPrice(100.0, shares0[:])
	cost0 := Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost0)

	// 1st
	added := 100.0
	shares0[0] = shares0[0] + added
	prices = CmpPrice(100.0, shares0[:])
	cost1 := Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume := Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 2nd
	added = 40.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 3rd
	added = 20.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 4th
	added = 50.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 5th
	added = 100.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 6th
	added = 50.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 7th
	added = -40.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 8th
	added = 30.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 9th
	added = 40.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 10th
	added = 300.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 11th
	added = -10.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 12th
	added = 150.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 13th
	added = -40.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 14th
	added = 20.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 15th
	added = 40.0
	shares0[0] = shares0[0] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 0, cost1-cost0)
	t.Log("volume:", volume)
	shares1[0] = shares1[0] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 16th
	added = 200.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}

	// 17th
	added = -100.0
	shares0[1] = shares0[1] + added
	cost0 = cost1
	prices = CmpPrice(100.0, shares0[:])
	cost1 = Cost(100.0, shares0[:])
	t.Log("prices:", prices)
	t.Log("cost:", cost1)
	t.Log("amount:", cost1-cost0)

	volume = Volume(100.0, shares1[:], 1, cost1-cost0)
	t.Log("volume:", volume)
	shares1[1] = shares1[1] + added
	t.Log("shares:", shares0, shares1)

	if math.Abs(volume-added) >= 0.00000001 {
		t.Error("not match:", volume, added)
	}
}
