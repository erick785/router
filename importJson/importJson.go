package importjson

import "encoding/json"

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
	Tokens map[string]bool
	Pairs  map[string][]string
}

func NewImportData() *ImportData {
	return &ImportData{
		Tokens: make(map[string]bool),
		Pairs:  make(map[string][]string),
	}
}

func (i *ImportData) Import(context []byte) error {
	var response Response
	if err := json.Unmarshal(context, &response); err != nil {
		return err
	}

	for _, pair := range response.Data.Pairs {
		i.Tokens[pair.Token0.ID] = true
		i.Tokens[pair.Token1.ID] = true
		i.Pairs[pair.Token0.ID+"_"+pair.Token1.ID] = append(i.Pairs[pair.Token0.ID+pair.Token1.ID], pair.ID)
	}

	return nil
}
