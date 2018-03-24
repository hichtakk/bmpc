package main

import (
	"math/big"
	"testing"
)

func TestConvertUnit(t *testing.T) {
	btc := 10000000000.0
	expected := 100.0
	actual := convertUnit(btc)
	if actual != expected {
		t.Error(actual)
	}
}

func TestGetHashrate(t *testing.T) {
	unit := "MH"
	hash_value := 20
	expected := new(big.Float).SetFloat64(20000000)
	actual := getHashrate(float64(hash_value), &unit)
	if actual != expected {
		t.Error(actual)
	}
}

func TestCalcProfit(t *testing.T) {
	difficulty := new(big.Float).SetFloat64(47427554950.648)
	hashrate := new(big.Float).SetFloat64(21990232555520)
	block_reward := 25.0
	mining_days := 30
	expected := 6.9954270327711034
	actual, err := calcProfit(difficulty, hashrate, block_reward, mining_days)
	if err != nil {
		t.Error("test failed")
	}
	if actual != expected {
		t.Error(actual)
	}
}
