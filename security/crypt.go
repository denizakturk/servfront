package security

import (
	"encoding/base64"
	"math"
	"github.com/denizakturk/service-frontend/config"
)

type Crypt struct {
	Config       *config.Crypt
	Base64Output bool
}

func NewCrypt(config *config.Crypt) *Crypt {
	return &Crypt{Config: config}
}

func (c *Crypt) Encrypt(rawData string) string {
	key, iv := c.getWorkForwardSecrets()
	crypt := Crypt{Config: &config.Crypt{Key: key, IV: iv}}
	encryptSourceData := crypt.breakForward(rawData)
	encryptData := append([]byte(iv), []byte(encryptSourceData)...)
	var output string
	if c.Base64Output {
		output = base64.StdEncoding.EncodeToString(encryptData)
	} else {
		output = string(encryptData)
	}
	return output
}

func (c *Crypt) Decrypt(encryptData string) string {
	encryptIv := string([]byte(encryptData)[:32])
	encryptSourceData := string([]byte(encryptData)[32:])
	iv := c.breakBackward(encryptIv)
	crypt := Crypt{Config: &config.Crypt{Key: iv}}
	key := crypt.breakForward(string(c.Config.GetKeyBytes()))
	crypt = Crypt{Config: &config.Crypt{Key: key}}

	return crypt.breakBackward(encryptSourceData)
}

func (c *Crypt) getWorkForwardSecrets() (key string, iv string) {
	crypt := Crypt{Config: &config.Crypt{Key: string(c.Config.GetIVBytes())}}
	key = crypt.breakForward(string(c.Config.GetKeyBytes()))
	iv = c.breakForward(string(c.Config.GetIVBytes()))

	return key, iv
}


func (c *Crypt) breakForward(data string) string {
	cryptEngine := &CryptEngine{}
	cryptEngine.Init(data, c.Config.Key)
	for _, cik := range cryptEngine.DataRune {
		cryptEngine.SetCryptItemKeeper(cik)
		cryptEngine.RecountKetRune()
		cryptEngine.ResetDiffer()
		constCountKeyRune := cryptEngine.CountKeyRune
		for _, kr := range cryptEngine.KeyRune {
			cryptEngine.RatioCalc()
			cryptEngine.SumDiffWithDiffer(kr, constCountKeyRune)
			cryptEngine.DecreaseCountKeyRune()
		}
		cryptEngine.DifferRatioCalc()
		cryptEngine.SumWithCryptItem()
		cryptEngine.AddCryptKeeper()
		cryptEngine.DecreaseCountDataRune()
	}

	return cryptEngine.OutputCryptKeeper()
}

func (c *Crypt) breakBackward(data string) string {
	cryptEngine := &CryptEngine{}
	cryptEngine.Init(data, c.Config.Key)
	for _, cik := range cryptEngine.DataRune {
		cryptEngine.SetCryptItemKeeper(cik)
		cryptEngine.RecountKetRune()
		cryptEngine.ResetDiffer()
		constCountKeyRune := cryptEngine.CountKeyRune
		for _, kr := range cryptEngine.KeyRune {
			cryptEngine.RatioCalc()
			cryptEngine.SumDiffWithDiffer(kr, constCountKeyRune)
			cryptEngine.DecreaseCountKeyRune()
		}
		cryptEngine.DifferRatioCalc()
		cryptEngine.ExtractWithCryptItem()
		cryptEngine.AddCryptKeeper()
		cryptEngine.DecreaseCountDataRune()
	}

	return cryptEngine.OutputCryptKeeper()
}

type CryptEngine struct {
	DataRune        []rune
	KeyRune         []rune
	CryptKeeper     []int32
	CryptItemKeeper int32
	CountDataRune   int32
	CountKeyRune    int32
	Differ          int32
	DifferRatio     int32
	Ratio           int32
}

func (ce *CryptEngine) Init(data, key string) {
	ce.DataRune = []rune(data)
	ce.KeyRune = []rune(key)
	ce.CryptKeeper = []int32{}
	ce.RecountDataRune()
}

func (ce *CryptEngine) RecountDataRune() {
	ce.CountDataRune = int32(len(ce.DataRune))
}

func (ce *CryptEngine) RecountKetRune() {
	ce.CountKeyRune = int32(len(ce.KeyRune))
}

func (ce *CryptEngine) RatioCalc() {
	ce.Ratio = (ce.CountKeyRune / ce.CountDataRune) + (ce.CountDataRune / ce.CountKeyRune)
}

func (ce *CryptEngine) DifferRatioCalc() {
	ce.DifferRatio = ce.Differ % int32(math.MaxInt16)
}

func (ce *CryptEngine) SumWithCryptItem() {
	ce.CryptItemKeeper += ce.DifferRatio
}

func (ce *CryptEngine) ExtractWithCryptItem() {
	ce.CryptItemKeeper -= ce.DifferRatio
}

func (ce *CryptEngine) AddCryptKeeper() {
	ce.CryptKeeper = append(ce.CryptKeeper, ce.CryptItemKeeper)
}

func (ce *CryptEngine) SetCryptItemKeeper(item int32) {
	ce.CryptItemKeeper = item
}

func (ce *CryptEngine) ResetDiffer() {
	ce.Differ = int32(0)
}

func (ce *CryptEngine) DecreaseCountDataRune() {
	ce.CountDataRune--
}

func (ce *CryptEngine) DecreaseCountKeyRune() {
	ce.CountKeyRune--
}

func (ce *CryptEngine) SumDiffWithDiffer(keyRuneItem rune, constCountKeyRune int32) {
	ce.Differ += (keyRuneItem + ce.Ratio) + (constCountKeyRune / ce.CountKeyRune)
}

func (ce *CryptEngine) OutputCryptKeeper() string {
	return string(ce.CryptKeeper)
}