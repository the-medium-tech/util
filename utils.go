package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"math/rand"
)

func GenesisCards(deckCount int) []string {
	cardSet := [52]string{
		"SA", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "ST", "SJ", "SQ", "SK",
		"DA", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "DT", "DJ", "DQ", "DK",
		"HA", "H2", "H3", "H4", "H5", "H6", "H7", "H8", "H9", "HT", "HJ", "HQ", "HK",
		"CA", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "CT", "CJ", "CQ", "CK",
	}
	totalCards := deckCount * len(cardSet)
	cardOrder := make([]string, totalCards)

	for deck := 0; deck < deckCount; deck++ {
		copy(cardOrder[deck*len(cardSet):], cardSet[:])
	}
	return cardOrder
}

func DeckShuffle(data []string, rand *rand.Rand) []string {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func PopCard(beforeDeck []string, index int) (card string, afterDeck []string) {
	Deck := make([]string, 0) //empty string slice
	card = beforeDeck[index]
	Deck = append(Deck, beforeDeck[:index]...)
	return card, append(Deck, beforeDeck[index+1:]...)
}

func Sha256ToInt64(data string, salt string) int64 {
	dataBytes := []byte(data)
	saltedDataBytes := append(dataBytes, []byte(salt)...)
	hash := sha256.Sum256(saltedDataBytes)
	var sum int64 = 0
	for i := 0; i < 32; i += 8 {
		sum += int64(binary.BigEndian.Uint64(hash[i : i+8]))
	}
	return sum
}

func EncodeDeckToGob(data []string) (string, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), err
}

func DecodeDeckFromGob(data string) ([]string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var dest []string
	decoder := gob.NewDecoder(bytes.NewBuffer(decoded))
	err = decoder.Decode(&dest)
	return dest, err
}

func TranscodeInterfaceToStruct(in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func Gen32byteKey(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
