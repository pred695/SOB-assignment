package Blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pred695/code-challenge-2024-pred695/Structs"
	"github.com/pred695/code-challenge-2024-pred695/Utils"
)

var Bh Structs.BlockHeader = Structs.BlockHeader{
	Version:       7,
	PrevBlockHash: "0000000000000000000000000000000000000000000000000000000000000000",
	MerkleRoot:    "",
	Time:          time.Now().Unix(),
	Bits:          0x1f00ffff,
	Nonce:         0,
}

func GetAllTxid() []string {
	var permittedTxIDs []string
	dir := "./mempool"
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		txData, err := Utils.JsonData(dir + "/" + file.Name())
		Utils.Handle(err)
		var tx Structs.Transaction
		err = json.Unmarshal([]byte(txData), &tx)
		serialized, _ := Utils.SerializeTransaction(&tx)
		txID := Utils.ReverseBytes(Utils.To_sha(Utils.To_sha(serialized)))
		permittedTxIDs = append(permittedTxIDs, hex.EncodeToString(txID))
	}
	return permittedTxIDs

}
func MineBlock() {
	netReward, TxIDs, _ := Utils.Prioritize()

	cbTx := Utils.CreateCoinbase(netReward)
	serializedcbTx, _ := Utils.SerializeTransaction(cbTx)
	fmt.Printf("CBTX: %x\n", serializedcbTx)
	txidsnew := GetAllTxid()
	fmt.Println("Length of txids: ", len(txidsnew))
	TxIDs = append([]string{hex.EncodeToString(Utils.ReverseBytes(Utils.To_sha(Utils.To_sha(serializedcbTx))))}, GetAllTxid()...)
	mkr := Utils.NewMerkleTree(TxIDs)
	Bh.MerkleRoot = hex.EncodeToString(mkr.Data)
	cbtxbase := Utils.CalculateBaseSize(cbTx)
	cbtxwitness := Utils.CalculateWitnessSize(cbTx)
	fmt.Println("Cbtx wt: ", cbtxwitness+(cbtxbase*4))
	if ProofOfWork(&Bh) {
		file, _ := os.Create("output.txt")
		defer file.Close()
		serializedBh := Utils.SerializeBlockHeader(&Bh)
		segserialized, _ := Utils.SegWitSerialize(cbTx)
		file.WriteString(hex.EncodeToString(serializedBh) + "\n")
		file.WriteString(hex.EncodeToString(segserialized) + "\n")
		for _, tx := range TxIDs {
			file.WriteString(tx + "\n")
		}
	}
}
