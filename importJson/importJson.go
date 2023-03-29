package importjson

import (
	"encoding/json"
	"math/big"
	"sync"
)

type Token struct {
	ID string `json:"id"`
}

type Pair struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	TrackedReserveBnb string `json:"trackedReserveBNB"`
	Token0            Token  `json:"token0"`
	Token1            Token  `json:"token1"`
}

type Data struct {
	Pairs []Pair `json:"pairs"`
}

type Response struct {
	Data Data `json:"data"`
}

type ImportData struct {
	Tokens sync.Map // map[string]bool
	Pairs  sync.Map // map[string][]string
}

func NewImportData() *ImportData {
	return &ImportData{
		Tokens: sync.Map{},
		Pairs:  sync.Map{},
	}
}

func (i *ImportData) Import(context []byte) error {
	var response Response
	if err := json.Unmarshal(context, &response); err != nil {
		return err
	}

	for _, pair := range response.Data.Pairs {
		i.Tokens.Store(pair.Token0.ID, true)
		i.Tokens.Store(pair.Token1.ID, true)

		// i.Pairs[pair.Token0.ID+"_"+pair.Token1.ID] = append(i.Pairs[pair.Token0.ID+pair.Token1.ID], pair.ID)
		res, ok := i.Pairs.Load(pair.Token0.ID + "_" + pair.Token1.ID)
		if !ok {
			res = []string{}
		}
		i.Pairs.Store(pair.Token0.ID+"_"+pair.Token1.ID, append(res.([]string), pair.ID))
	}

	return nil
}

func (i *ImportData) TokensForEach(f func(key, value interface{}) bool) {
	i.Tokens.Range(f)
}

func (i *ImportData) PairsForEach(f func(key, value interface{}) bool) {
	i.Pairs.Range(f)
}

func (i *ImportData) GetTokensLen() int {
	var len int
	i.TokensForEach(func(key, value interface{}) bool {
		len++
		return true
	})
	return len
}

func (i *ImportData) GetPairsLen() int {
	var len int
	i.PairsForEach(func(key, value interface{}) bool {
		len++
		return true
	})
	return len
}

func (i *ImportData) GetPair(token0, token1 string) []string {
	intAddress1 := new(big.Int)
	intAddress1.SetString(token0, 0)
	intAddress2 := new(big.Int)
	intAddress2.SetString(token1, 0)

	// Compare addresses
	result := intAddress1.Cmp(intAddress2)
	if result == 0 {
		return []string{}
	} else if result > 0 {
		token0, token1 = token1, token0
	}

	res, ok := i.Pairs.Load(token0 + "_" + token1)
	if !ok {
		return []string{}
	}
	return res.([]string)
}
