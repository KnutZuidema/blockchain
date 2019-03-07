package models

var GenesisReward = 10

type Transaction struct {
	Id      []byte               `json:"id"`
	Inputs  []*TransactionInput  `json:"inputs"`
	Outputs []*TransactionOutput `json:"outputs"`
}

type TransactionOutput struct {
	Value     int    `json:"value"`
	PublicKey string `json:"public_key"`
}

type TransactionInput struct {
	TransactionId []byte `json:"transaction_id"`
	Value         int    `json:"value"`
	Signature     string `json:"signature"`
}

func NewGenesisTransaction(receiver, data string) *Transaction {
	input := &TransactionInput{[]byte{}, -1, data}
	output := &TransactionOutput{GenesisReward, receiver}
	transaction := &Transaction{nil, []*TransactionInput{input}, []*TransactionOutput{output}}
	transaction.SetId()
	return transaction
}

func (tx *Transaction) SetId() {
	tx.Id = []byte("id")
}
