package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ct_p2pkh  = 0
	ct_p2sh   = 0
	ct_p2wpkh = 0
	ct_p2wsh  = 0
)

func Address() {
	dir := "./mempool"
	files, err := os.ReadDir(dir)
	Handle(err)
	for _, file := range files {
		txData, err := jsonData(dir + "/" + file.Name())
		Handle(err)
		var tx Transaction
		err = json.Unmarshal([]byte(txData), &tx)
		Handle(err)
		for _, vin := range tx.Vin {
			if vin.Prevout.ScriptpubkeyType == "p2pkh" {
				pubkey_asm := vin.Prevout.ScriptpubkeyAsm
				address := p2pkh(pubkey_asm)
				if string(address) == vin.Prevout.ScriptpubkeyAddress {
					// fmt.Println(vin.Prevout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2pkh++
					continue
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vin.Prevout.ScriptpubkeyAddress)
				}
			}

			if vin.Prevout.ScriptpubkeyType == "p2sh" {
				pubkey_asm := vin.Prevout.ScriptpubkeyAsm
				address := p2sh(pubkey_asm)
				if string(address) == vin.Prevout.ScriptpubkeyAddress {
					// fmt.Println(vin.Prevout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2sh++
					continue
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vin.Prevout.ScriptpubkeyAddress)
				}
			}

			if vin.Prevout.ScriptpubkeyType == "v0_p2wpkh" {
				pubkey_asm := vin.Prevout.ScriptpubkeyAsm
				address := p2wpkh(pubkey_asm)
				if string(address) == vin.Prevout.ScriptpubkeyAddress {
					// fmt.Println(vin.Prevout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2wpkh++
					continue
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vin.Prevout.ScriptpubkeyAddress)
				}
			}

			if vin.Prevout.ScriptpubkeyType == "v0_p2wsh" {
				pubkey_asm := vin.Prevout.ScriptpubkeyAsm
				address := p2wsh(pubkey_asm)
				if string(address) == vin.Prevout.ScriptpubkeyAddress {
					// fmt.Println(vin.Prevout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2wsh++
					continue
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vin.Prevout.ScriptpubkeyAddress)
				}
			}
		}
		for _, vout := range tx.Vout {
			if vout.ScriptpubkeyType == "p2pkh" {
				pubkey_asm := vout.ScriptpubkeyAsm
				address := p2pkh(pubkey_asm)
				if string(address) == vout.ScriptpubkeyAddress {
					// fmt.Println(vout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2pkh++
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vout.ScriptpubkeyAddress)
				}
			}

			if vout.ScriptpubkeyType == "p2sh" {
				pubkey_asm := vout.ScriptpubkeyAsm
				address := p2sh(pubkey_asm)
				if string(address) == vout.ScriptpubkeyAddress {
					// fmt.Println(vout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2sh++
					continue
				} else {
					fmt.Println("Address not matched")
					fmt.Println("Address: ", address)
					fmt.Println("Scriptpubkey Address: ", vout.ScriptpubkeyAddress)
				}
			}

			if vout.ScriptpubkeyType == "v0_p2wpkh" {
				pubkey_asm := vout.ScriptpubkeyAsm
				address := p2wpkh(pubkey_asm)
				if string(address) == vout.ScriptpubkeyAddress {
					// fmt.Println(vout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2wpkh++
				} else {
					fmt.Println("Address not matched")
					fmt.Printf("Address: %s\n", address)
					fmt.Println("Scriptpubkey Address: ", vout.ScriptpubkeyAddress)
				}
			}

			if vout.ScriptpubkeyType == "v0_p2wsh" {
				pubkey_asm := vout.ScriptpubkeyAsm
				address := p2wsh(pubkey_asm)
				if string(address) == vout.ScriptpubkeyAddress {
					// fmt.Println(vout.ScriptpubkeyAddress)
					// fmt.Printf(" %s\n", address)
					ct_p2wsh++
				} else {
					fmt.Println("Address not matched")
					fmt.Printf("Address: %s\n", address)
					fmt.Println("Scriptpubkey Address: ", vout.ScriptpubkeyAddress)
				}
			}
		}
	}
	fmt.Println("Count of p2pkh address matched: ", ct_p2pkh)
	fmt.Println("Count of p2sh address matched: ", ct_p2sh)
	fmt.Println("Count of p2wpkh address matched: ", ct_p2wpkh)
	fmt.Println("Count of p2wpkh address matched: ", ct_p2wsh)
}
