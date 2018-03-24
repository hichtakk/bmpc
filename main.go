package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func get(resource string) (string, error) {
	response, err := http.Get(resource)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getDifficulty() (*big.Float, error) {
	url := "https://blockchain.info/q/getdifficulty"
	difficulty, err := get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var difficultyValue *big.Float
	difficultyValue = new(big.Float)
	fmt.Sscan(difficulty, difficultyValue)

	return difficultyValue, nil
}

func getBlockReward() (float64, error) {
	url := "https://blockchain.info/q/bcperblock"
	rewardStr, err := get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	reward, err := strconv.ParseFloat(rewardStr, 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return reward, nil
}

func convertUnit(btc float64) float64 {
	satoshi := 100000000.0
	value := btc / satoshi

	return value
}

func calcProfit(d *big.Float, h *big.Float, r float64, day int) (float64, error) {
	reward := big.NewFloat(r)
	days := big.NewFloat(float64(day))
	seconds := big.NewFloat(86400)
	power := new(big.Float).SetFloat64(math.Pow(2, 32))
	value := new(big.Float)
	value.Mul(reward, days)
	value.Mul(value, seconds)
	value.Quo(value, d)
	value.Quo(value, power)
	value.Mul(value, h)
	fvalue, _ := value.Float64()

	return fvalue, nil
}

func getHashrate(h float64, u *string) *big.Float {
	// TH/s
	var unit int
	switch *u {
	case "TH":
		unit = 1000000000000
	case "GH":
		unit = 1000000000
	case "MH":
		unit = 1000000
	case "KH":
		unit = 1000
	case "H":
		unit = 1
	default:
		unit = 1000000000000
	}
	hashrate := new(big.Float).SetFloat64(h * float64(unit))

	return hashrate
}

func main() {
	// flags
	var (
		d = flag.Int("d", 1, "days")
		u = flag.String("u", "TH", "hashrate unit")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("too few argument")
		os.Exit(1)
	}

	// hashrate
	hash, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	hashrate := getHashrate(hash, u)

	// difficulty
	difficulty, err := getDifficulty()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// reward per block (satoshi)
	reward_satoshi, err := getBlockReward()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reward := convertUnit(reward_satoshi)

	profit, err := calcProfit(difficulty, hashrate, reward, *d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(profit)
}
