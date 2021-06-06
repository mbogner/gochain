package main

import (
	"crypto/sha512"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// https://tutorialedge.net/golang/the-go-init-function/
func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})
}

// <editor-fold desc="Data Structure">

type Tx struct {
	source string
	target string
	amount float64
}

func (t *Tx) StringToHash() string {
	return "Tx(" + t.source + "," + t.target + "," + fmt.Sprintf("%g", t.amount) + ")"
}
func (t *Tx) ToString() string {
	return "Tx(source=" + t.source + ",target=" + t.target + ",amount=" + fmt.Sprintf("%g", t.amount) + ")"
}

type Block struct {
	number  int64
	tx      Tx
	comment string
	ts      int64
	last    string
	current string
}

func (d *Block) StringToHash() string {
	return "Block(" +
		strconv.FormatInt(d.number, 10) + "," +
		d.tx.StringToHash() + "," +
		strconv.FormatInt(d.ts, 10) + "," +
		d.comment + "," +
		d.last +
		")"
}
func (d *Block) ToString() string {
	return "Block(number=" + strconv.FormatInt(d.number, 10) +
		",tx=" + d.tx.ToString() +
		",comment=" + d.comment +
		",ts=" + strconv.FormatInt(d.ts, 10) +
		",last=" + d.last +
		",current=" + d.current +
		")"
}
func (d *Block) Seal() {
	if "" == d.current {
		d.current = fmt.Sprintf("%x", sha512.Sum512([]byte(d.StringToHash())))
	} else {
		log.Error("can not change sealed object")
	}
}

//</editor-fold>

/**
 * This method creates a sealed Block entry and increments cnt.
 */
func createBlock(source string, target string, amount float64, cnt *int64, comment string, lastHash string) Block {
	localCnt := *cnt // create new variable from pointer value
	*cnt++           // increment cnt

	block := &Block{
		number: localCnt,
		tx: Tx{
			source: source,
			target: target,
			amount: amount,
		},
		comment: comment,
		ts:      time.Now().UnixNano(),
		last:    lastHash,
	}
	block.Seal()
	return *block
}

// https://golang.org/doc/code#Command
func main() {
	log.Info("starting application")
	var cnt int64 = 0

	block0 := createBlock("noSource", "noTarget", 0.0, &cnt, "genesis", "0")
	log.Debug("block0: " + block0.ToString())

	block1 := createBlock("random1", "mbo", 1.5, &cnt, "first tx", block0.current)
	log.Debug("block1: " + block1.ToString())

	block2 := createBlock("random2", "mbo", 2.5, &cnt, "second tx", block1.current)
	log.Debug("block2: " + block2.ToString())

	block3 := createBlock("random3", "mbo", 3.5, &cnt, "third tx", block2.current)
	log.Debug("block3: " + block3.ToString())

	log.Info("done")
}
