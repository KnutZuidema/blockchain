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
	OutputIndex   int    `json:"output_index"`
	Signature     string `json:"signature"`
}

func NewGenesisTransaction(receiver string) *Transaction {
	input := &TransactionInput{[]byte{}, -1, "genesis"}
	output := &TransactionOutput{GenesisReward, receiver}
	transaction := &Transaction{nil, []*TransactionInput{input}, []*TransactionOutput{output}}
	transaction.SetId()
	return transaction
}

func (tx *Transaction) SetId() {
	tx.Id = []byte("id")
}

func (output TransactionOutput) Verify(publicKey string) bool {
	return output.PublicKey == publicKey
}

func (input TransactionInput) Verify(signature string) bool {
	return input.Signature == signature
}
