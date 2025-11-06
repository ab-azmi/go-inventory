package number

import (
	"fmt"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type CRMNumber struct{}

func (num CRMNumber) BillNumber() string {
	return num.generator("BIL")
}

func (num CRMNumber) InvoiceNumber() string {
	return num.generator("INV")
}

func (num CRMNumber) BillingStatementNumber() string {
	return num.generator("BST")
}

func (num CRMNumber) PaymentNumber() string {
	return num.generator("PMT")
}

func (num CRMNumber) ReceiptNumber() string {
	return num.generator("RCP")
}

/** --- UNEXPORTED FUNCTIONS --- */

func (num CRMNumber) generator(prefix string) string {
	date := time.Now().Format("20060102")
	key := fmt.Sprintf("number_counter:%s:%s", prefix, date)

	conn := xtremepkg.RedisPool.Get()
	defer conn.Close()

	seq, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		log.Panicf("redis increment err: %v", err)
	}

	// Exp: 24 Hours
	conn.Do("EXPIRE", key, 86400)

	return fmt.Sprintf("%s%s%07d", prefix, date, seq)
}
