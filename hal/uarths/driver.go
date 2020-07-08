// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uarths

import (
	"embedded/rtos"
	"sync/atomic"
	"time"
)

type DriverError uint8

const (
	// ErrBufOverflow is returned if one or more received bytes has been dropped
	// because of the lack of free space in the driver's receive buffer.
	ErrBufOverflow DriverError = iota + 1

	// ErrTimeout is returned if timeout occured. It means that the read/write
	// operation has been interrupted. In case of write you can not determine
	// the exact number of bytes sent to the remote party.
	ErrTimeout
)

// Error implements error interface.
func (e DriverError) Error() string {
	switch e {
	case ErrBufOverflow:
		return "uart: buffer overflow"
	case ErrTimeout:
		return "uart: timeout"
	}
	return ""
}

type Driver struct {
	p *Periph

	// rx state
	rxbuf    []byte
	nextr    uint32
	nextw    uint32
	rxcmd    uint32
	rxready  rtos.Note
	overflow bool

	// tx state
	txdata string
	txn    int
	txdone rtos.Note

	isr       uint32
	timeoutRx time.Duration
	timeoutTx time.Duration
}

const (
	cmdNone = iota
	cmdWakeup
)

const (
	isrNone = iota
	isrRx
	isrTx
)

// NewDriver returns a new driver for p.
func NewDriver(p *Periph) *Driver {
	p.DisableIRQ(TxMin | RxMax)
	p.SetTxConf(0, 0)
	p.SetRxConf(0, 0)
	return &Driver{p: p, timeoutRx: -1, timeoutTx: -1}
}

func (d *Driver) Periph() *Periph {
	return d.p
}

func (d *Driver) SetStopBits(n int) {
	cfg, _ := d.p.TxConf()
	switch n {
	case 1:
		cfg &^= TxStop2b
	case 2:
		cfg |= TxStop2b
	default:
		panic("uarths: support only 1 or 2 stop bits")
	}
	d.p.SetTxConf(cfg, 0)
}

func (d *Driver) SetBaudrate(br int) {
	d.p.SetBaudrate(br)
}

// SetReadTimeout sets the read timeout used by Read* functions.
func (d *Driver) SetReadTimeout(timeout time.Duration) {
	d.timeoutRx = timeout
}

// SetWriteTimeout sets the write timeout used by Write* functions.
func (d *Driver) SetWriteTimeout(timeout time.Duration) {
	d.timeoutTx = timeout
}

// ISR handles UARTHS interrupts.
func (d *Driver) ISR() {
	if b, ok := d.p.Load(); ok {
		atomic.StoreUint32(&d.isr, isrRx)
		for {
			nextw := d.nextw
			d.rxbuf[nextw] = byte(b)
			if nextw++; int(nextw) == len(d.rxbuf) {
				nextw = 0
			}
			if nextw != atomic.LoadUint32(&d.nextr) {
				atomic.StoreUint32(&d.nextw, nextw)
				if atomic.CompareAndSwapUint32(&d.rxcmd, cmdWakeup, cmdNone) {
					d.rxready.Wakeup()
				}
			} else {
				d.overflow = true
			}
			if b, ok = d.p.Load(); !ok {
				break
			}
		}
		atomic.StoreUint32(&d.isr, isrNone)
	}
	if d.p.Events()&TxMin != 0 {
		atomic.StoreUint32(&d.isr, isrTx)
		if d.txn >= len(d.txdata) {
			d.p.SetTxMinCnt(0)
			d.txdone.Wakeup()
		} else {
			for {
				if d.p.TxFull() {
					if m := 9 - d.txn; m > txMin {
						if m > 7 {
							m = 7
						}
						d.p.SetTxMinCnt(m)
					}
					break
				}
				d.p.Store(int(d.txdata[d.txn]))
				if d.txn++; d.txn == len(d.txdata) {
					break
				}
			}
		}
		atomic.StoreUint32(&d.isr, isrNone)
	}
}
