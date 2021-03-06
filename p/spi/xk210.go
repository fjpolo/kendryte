// DO NOT EDIT THIS FILE. GENERATED BY xgen.

// +build k210

package spi

import (
	"embedded/mmio"
	"unsafe"

	"github.com/embeddedgo/kendryte/p/bus"
	"github.com/embeddedgo/kendryte/p/mmap"
)

type Periph struct {
	CTRLR0           RCTRLR0
	CTRLR1           RCTRLR1
	SSIENR           RSSIENR
	MWCR             RMWCR
	SER              RSER
	BAUDR            RBAUDR
	TXFTLR           RTXFTLR
	RXFTLR           RRXFTLR
	TXFLR            RTXFLR
	RXFLR            RRXFLR
	SR               RSR
	IMR              RIMR
	ISR              RISR
	RAW_ISR          RRAW_ISR
	TXOICR           RTXOICR
	RXOICR           RRXOICR
	RXUICR           RRXUICR
	MSTICR           RMSTICR
	ICR              RICR
	DMACR            RDMACR
	DMATDLR          RDMATDLR
	DMARDLR          RDMARDLR
	IDR              RIDR
	SSIC_VERSION_ID  RSSIC_VERSION_ID
	DR               [36]RDR
	RX_SAMPLE_DELAY  RRX_SAMPLE_DELAY
	SPI_CTRLR0       RSPI_CTRLR0
	_                uint32
	XIP_MODE_BITS    RXIP_MODE_BITS
	XIP_INCR_INST    RXIP_INCR_INST
	XIP_WRAP_INST    RXIP_WRAP_INST
	XIP_CTRL         RXIP_CTRL
	XIP_SER          RXIP_SER
	XRXOICR          RXRXOICR
	XIP_CNT_TIME_OUT RXIP_CNT_TIME_OUT
	ENDIAN           RENDIAN
}

func SPI0() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI0_BASE))) }
func SPI1() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE))) }
func SPI3() *Periph { return (*Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE))) }

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

func (p *Periph) Bus() bus.Bus {
	return bus.APB2
}

type CTRLR0 uint32

type RCTRLR0 struct{ mmio.U32 }

func (r *RCTRLR0) LoadBits(mask CTRLR0) CTRLR0 { return CTRLR0(r.U32.LoadBits(uint32(mask))) }
func (r *RCTRLR0) StoreBits(mask, b CTRLR0)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCTRLR0) SetBits(mask CTRLR0)         { r.U32.SetBits(uint32(mask)) }
func (r *RCTRLR0) ClearBits(mask CTRLR0)       { r.U32.ClearBits(uint32(mask)) }
func (r *RCTRLR0) Load() CTRLR0                { return CTRLR0(r.U32.Load()) }
func (r *RCTRLR0) Store(b CTRLR0)              { r.U32.Store(uint32(b)) }

type RMCTRLR0 struct{ mmio.UM32 }

func (rm RMCTRLR0) Load() CTRLR0   { return CTRLR0(rm.UM32.Load()) }
func (rm RMCTRLR0) Store(b CTRLR0) { rm.UM32.Store(uint32(b)) }

func (p *Periph) WORK_MODE() RMCTRLR0 {
	return RMCTRLR0{mmio.UM32{&p.CTRLR0.U32, uint32(WORK_MODE)}}
}

func (p *Periph) TMOD() RMCTRLR0 {
	return RMCTRLR0{mmio.UM32{&p.CTRLR0.U32, uint32(TMOD)}}
}

func (p *Periph) DATA_LENGTH() RMCTRLR0 {
	return RMCTRLR0{mmio.UM32{&p.CTRLR0.U32, uint32(DATA_LENGTH)}}
}

func (p *Periph) FRAME_FORMAT() RMCTRLR0 {
	return RMCTRLR0{mmio.UM32{&p.CTRLR0.U32, uint32(FRAME_FORMAT)}}
}

type CTRLR1 uint32

type RCTRLR1 struct{ mmio.U32 }

func (r *RCTRLR1) LoadBits(mask CTRLR1) CTRLR1 { return CTRLR1(r.U32.LoadBits(uint32(mask))) }
func (r *RCTRLR1) StoreBits(mask, b CTRLR1)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCTRLR1) SetBits(mask CTRLR1)         { r.U32.SetBits(uint32(mask)) }
func (r *RCTRLR1) ClearBits(mask CTRLR1)       { r.U32.ClearBits(uint32(mask)) }
func (r *RCTRLR1) Load() CTRLR1                { return CTRLR1(r.U32.Load()) }
func (r *RCTRLR1) Store(b CTRLR1)              { r.U32.Store(uint32(b)) }

type RMCTRLR1 struct{ mmio.UM32 }

func (rm RMCTRLR1) Load() CTRLR1   { return CTRLR1(rm.UM32.Load()) }
func (rm RMCTRLR1) Store(b CTRLR1) { rm.UM32.Store(uint32(b)) }

type SSIENR uint32

type RSSIENR struct{ mmio.U32 }

func (r *RSSIENR) LoadBits(mask SSIENR) SSIENR { return SSIENR(r.U32.LoadBits(uint32(mask))) }
func (r *RSSIENR) StoreBits(mask, b SSIENR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSSIENR) SetBits(mask SSIENR)         { r.U32.SetBits(uint32(mask)) }
func (r *RSSIENR) ClearBits(mask SSIENR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RSSIENR) Load() SSIENR                { return SSIENR(r.U32.Load()) }
func (r *RSSIENR) Store(b SSIENR)              { r.U32.Store(uint32(b)) }

type RMSSIENR struct{ mmio.UM32 }

func (rm RMSSIENR) Load() SSIENR   { return SSIENR(rm.UM32.Load()) }
func (rm RMSSIENR) Store(b SSIENR) { rm.UM32.Store(uint32(b)) }

type MWCR uint32

type RMWCR struct{ mmio.U32 }

func (r *RMWCR) LoadBits(mask MWCR) MWCR { return MWCR(r.U32.LoadBits(uint32(mask))) }
func (r *RMWCR) StoreBits(mask, b MWCR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMWCR) SetBits(mask MWCR)       { r.U32.SetBits(uint32(mask)) }
func (r *RMWCR) ClearBits(mask MWCR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RMWCR) Load() MWCR              { return MWCR(r.U32.Load()) }
func (r *RMWCR) Store(b MWCR)            { r.U32.Store(uint32(b)) }

type RMMWCR struct{ mmio.UM32 }

func (rm RMMWCR) Load() MWCR   { return MWCR(rm.UM32.Load()) }
func (rm RMMWCR) Store(b MWCR) { rm.UM32.Store(uint32(b)) }

type SER uint32

type RSER struct{ mmio.U32 }

func (r *RSER) LoadBits(mask SER) SER { return SER(r.U32.LoadBits(uint32(mask))) }
func (r *RSER) StoreBits(mask, b SER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSER) SetBits(mask SER)      { r.U32.SetBits(uint32(mask)) }
func (r *RSER) ClearBits(mask SER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSER) Load() SER             { return SER(r.U32.Load()) }
func (r *RSER) Store(b SER)           { r.U32.Store(uint32(b)) }

type RMSER struct{ mmio.UM32 }

func (rm RMSER) Load() SER   { return SER(rm.UM32.Load()) }
func (rm RMSER) Store(b SER) { rm.UM32.Store(uint32(b)) }

type BAUDR uint32

type RBAUDR struct{ mmio.U32 }

func (r *RBAUDR) LoadBits(mask BAUDR) BAUDR { return BAUDR(r.U32.LoadBits(uint32(mask))) }
func (r *RBAUDR) StoreBits(mask, b BAUDR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBAUDR) SetBits(mask BAUDR)        { r.U32.SetBits(uint32(mask)) }
func (r *RBAUDR) ClearBits(mask BAUDR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RBAUDR) Load() BAUDR               { return BAUDR(r.U32.Load()) }
func (r *RBAUDR) Store(b BAUDR)             { r.U32.Store(uint32(b)) }

type RMBAUDR struct{ mmio.UM32 }

func (rm RMBAUDR) Load() BAUDR   { return BAUDR(rm.UM32.Load()) }
func (rm RMBAUDR) Store(b BAUDR) { rm.UM32.Store(uint32(b)) }

type TXFTLR uint32

type RTXFTLR struct{ mmio.U32 }

func (r *RTXFTLR) LoadBits(mask TXFTLR) TXFTLR { return TXFTLR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXFTLR) StoreBits(mask, b TXFTLR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXFTLR) SetBits(mask TXFTLR)         { r.U32.SetBits(uint32(mask)) }
func (r *RTXFTLR) ClearBits(mask TXFTLR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RTXFTLR) Load() TXFTLR                { return TXFTLR(r.U32.Load()) }
func (r *RTXFTLR) Store(b TXFTLR)              { r.U32.Store(uint32(b)) }

type RMTXFTLR struct{ mmio.UM32 }

func (rm RMTXFTLR) Load() TXFTLR   { return TXFTLR(rm.UM32.Load()) }
func (rm RMTXFTLR) Store(b TXFTLR) { rm.UM32.Store(uint32(b)) }

type RXFTLR uint32

type RRXFTLR struct{ mmio.U32 }

func (r *RRXFTLR) LoadBits(mask RXFTLR) RXFTLR { return RXFTLR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXFTLR) StoreBits(mask, b RXFTLR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXFTLR) SetBits(mask RXFTLR)         { r.U32.SetBits(uint32(mask)) }
func (r *RRXFTLR) ClearBits(mask RXFTLR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRXFTLR) Load() RXFTLR                { return RXFTLR(r.U32.Load()) }
func (r *RRXFTLR) Store(b RXFTLR)              { r.U32.Store(uint32(b)) }

type RMRXFTLR struct{ mmio.UM32 }

func (rm RMRXFTLR) Load() RXFTLR   { return RXFTLR(rm.UM32.Load()) }
func (rm RMRXFTLR) Store(b RXFTLR) { rm.UM32.Store(uint32(b)) }

type TXFLR uint32

type RTXFLR struct{ mmio.U32 }

func (r *RTXFLR) LoadBits(mask TXFLR) TXFLR { return TXFLR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXFLR) StoreBits(mask, b TXFLR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXFLR) SetBits(mask TXFLR)        { r.U32.SetBits(uint32(mask)) }
func (r *RTXFLR) ClearBits(mask TXFLR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RTXFLR) Load() TXFLR               { return TXFLR(r.U32.Load()) }
func (r *RTXFLR) Store(b TXFLR)             { r.U32.Store(uint32(b)) }

type RMTXFLR struct{ mmio.UM32 }

func (rm RMTXFLR) Load() TXFLR   { return TXFLR(rm.UM32.Load()) }
func (rm RMTXFLR) Store(b TXFLR) { rm.UM32.Store(uint32(b)) }

type RXFLR uint32

type RRXFLR struct{ mmio.U32 }

func (r *RRXFLR) LoadBits(mask RXFLR) RXFLR { return RXFLR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXFLR) StoreBits(mask, b RXFLR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXFLR) SetBits(mask RXFLR)        { r.U32.SetBits(uint32(mask)) }
func (r *RRXFLR) ClearBits(mask RXFLR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRXFLR) Load() RXFLR               { return RXFLR(r.U32.Load()) }
func (r *RRXFLR) Store(b RXFLR)             { r.U32.Store(uint32(b)) }

type RMRXFLR struct{ mmio.UM32 }

func (rm RMRXFLR) Load() RXFLR   { return RXFLR(rm.UM32.Load()) }
func (rm RMRXFLR) Store(b RXFLR) { rm.UM32.Store(uint32(b)) }

type SR uint32

type RSR struct{ mmio.U32 }

func (r *RSR) LoadBits(mask SR) SR  { return SR(r.U32.LoadBits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

type IMR uint32

type RIMR struct{ mmio.U32 }

func (r *RIMR) LoadBits(mask IMR) IMR { return IMR(r.U32.LoadBits(uint32(mask))) }
func (r *RIMR) StoreBits(mask, b IMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) SetBits(mask IMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR) ClearBits(mask IMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR) Load() IMR             { return IMR(r.U32.Load()) }
func (r *RIMR) Store(b IMR)           { r.U32.Store(uint32(b)) }

type RMIMR struct{ mmio.UM32 }

func (rm RMIMR) Load() IMR   { return IMR(rm.UM32.Load()) }
func (rm RMIMR) Store(b IMR) { rm.UM32.Store(uint32(b)) }

type ISR uint32

type RISR struct{ mmio.U32 }

func (r *RISR) LoadBits(mask ISR) ISR { return ISR(r.U32.LoadBits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

type RAW_ISR uint32

type RRAW_ISR struct{ mmio.U32 }

func (r *RRAW_ISR) LoadBits(mask RAW_ISR) RAW_ISR { return RAW_ISR(r.U32.LoadBits(uint32(mask))) }
func (r *RRAW_ISR) StoreBits(mask, b RAW_ISR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRAW_ISR) SetBits(mask RAW_ISR)          { r.U32.SetBits(uint32(mask)) }
func (r *RRAW_ISR) ClearBits(mask RAW_ISR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RRAW_ISR) Load() RAW_ISR                 { return RAW_ISR(r.U32.Load()) }
func (r *RRAW_ISR) Store(b RAW_ISR)               { r.U32.Store(uint32(b)) }

type RMRAW_ISR struct{ mmio.UM32 }

func (rm RMRAW_ISR) Load() RAW_ISR   { return RAW_ISR(rm.UM32.Load()) }
func (rm RMRAW_ISR) Store(b RAW_ISR) { rm.UM32.Store(uint32(b)) }

type TXOICR uint32

type RTXOICR struct{ mmio.U32 }

func (r *RTXOICR) LoadBits(mask TXOICR) TXOICR { return TXOICR(r.U32.LoadBits(uint32(mask))) }
func (r *RTXOICR) StoreBits(mask, b TXOICR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXOICR) SetBits(mask TXOICR)         { r.U32.SetBits(uint32(mask)) }
func (r *RTXOICR) ClearBits(mask TXOICR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RTXOICR) Load() TXOICR                { return TXOICR(r.U32.Load()) }
func (r *RTXOICR) Store(b TXOICR)              { r.U32.Store(uint32(b)) }

type RMTXOICR struct{ mmio.UM32 }

func (rm RMTXOICR) Load() TXOICR   { return TXOICR(rm.UM32.Load()) }
func (rm RMTXOICR) Store(b TXOICR) { rm.UM32.Store(uint32(b)) }

type RXOICR uint32

type RRXOICR struct{ mmio.U32 }

func (r *RRXOICR) LoadBits(mask RXOICR) RXOICR { return RXOICR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXOICR) StoreBits(mask, b RXOICR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXOICR) SetBits(mask RXOICR)         { r.U32.SetBits(uint32(mask)) }
func (r *RRXOICR) ClearBits(mask RXOICR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRXOICR) Load() RXOICR                { return RXOICR(r.U32.Load()) }
func (r *RRXOICR) Store(b RXOICR)              { r.U32.Store(uint32(b)) }

type RMRXOICR struct{ mmio.UM32 }

func (rm RMRXOICR) Load() RXOICR   { return RXOICR(rm.UM32.Load()) }
func (rm RMRXOICR) Store(b RXOICR) { rm.UM32.Store(uint32(b)) }

type RXUICR uint32

type RRXUICR struct{ mmio.U32 }

func (r *RRXUICR) LoadBits(mask RXUICR) RXUICR { return RXUICR(r.U32.LoadBits(uint32(mask))) }
func (r *RRXUICR) StoreBits(mask, b RXUICR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXUICR) SetBits(mask RXUICR)         { r.U32.SetBits(uint32(mask)) }
func (r *RRXUICR) ClearBits(mask RXUICR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RRXUICR) Load() RXUICR                { return RXUICR(r.U32.Load()) }
func (r *RRXUICR) Store(b RXUICR)              { r.U32.Store(uint32(b)) }

type RMRXUICR struct{ mmio.UM32 }

func (rm RMRXUICR) Load() RXUICR   { return RXUICR(rm.UM32.Load()) }
func (rm RMRXUICR) Store(b RXUICR) { rm.UM32.Store(uint32(b)) }

type MSTICR uint32

type RMSTICR struct{ mmio.U32 }

func (r *RMSTICR) LoadBits(mask MSTICR) MSTICR { return MSTICR(r.U32.LoadBits(uint32(mask))) }
func (r *RMSTICR) StoreBits(mask, b MSTICR)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMSTICR) SetBits(mask MSTICR)         { r.U32.SetBits(uint32(mask)) }
func (r *RMSTICR) ClearBits(mask MSTICR)       { r.U32.ClearBits(uint32(mask)) }
func (r *RMSTICR) Load() MSTICR                { return MSTICR(r.U32.Load()) }
func (r *RMSTICR) Store(b MSTICR)              { r.U32.Store(uint32(b)) }

type RMMSTICR struct{ mmio.UM32 }

func (rm RMMSTICR) Load() MSTICR   { return MSTICR(rm.UM32.Load()) }
func (rm RMMSTICR) Store(b MSTICR) { rm.UM32.Store(uint32(b)) }

type ICR uint32

type RICR struct{ mmio.U32 }

func (r *RICR) LoadBits(mask ICR) ICR { return ICR(r.U32.LoadBits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

type DMACR uint32

type RDMACR struct{ mmio.U32 }

func (r *RDMACR) LoadBits(mask DMACR) DMACR { return DMACR(r.U32.LoadBits(uint32(mask))) }
func (r *RDMACR) StoreBits(mask, b DMACR)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDMACR) SetBits(mask DMACR)        { r.U32.SetBits(uint32(mask)) }
func (r *RDMACR) ClearBits(mask DMACR)      { r.U32.ClearBits(uint32(mask)) }
func (r *RDMACR) Load() DMACR               { return DMACR(r.U32.Load()) }
func (r *RDMACR) Store(b DMACR)             { r.U32.Store(uint32(b)) }

type RMDMACR struct{ mmio.UM32 }

func (rm RMDMACR) Load() DMACR   { return DMACR(rm.UM32.Load()) }
func (rm RMDMACR) Store(b DMACR) { rm.UM32.Store(uint32(b)) }

type DMATDLR uint32

type RDMATDLR struct{ mmio.U32 }

func (r *RDMATDLR) LoadBits(mask DMATDLR) DMATDLR { return DMATDLR(r.U32.LoadBits(uint32(mask))) }
func (r *RDMATDLR) StoreBits(mask, b DMATDLR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDMATDLR) SetBits(mask DMATDLR)          { r.U32.SetBits(uint32(mask)) }
func (r *RDMATDLR) ClearBits(mask DMATDLR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RDMATDLR) Load() DMATDLR                 { return DMATDLR(r.U32.Load()) }
func (r *RDMATDLR) Store(b DMATDLR)               { r.U32.Store(uint32(b)) }

type RMDMATDLR struct{ mmio.UM32 }

func (rm RMDMATDLR) Load() DMATDLR   { return DMATDLR(rm.UM32.Load()) }
func (rm RMDMATDLR) Store(b DMATDLR) { rm.UM32.Store(uint32(b)) }

type DMARDLR uint32

type RDMARDLR struct{ mmio.U32 }

func (r *RDMARDLR) LoadBits(mask DMARDLR) DMARDLR { return DMARDLR(r.U32.LoadBits(uint32(mask))) }
func (r *RDMARDLR) StoreBits(mask, b DMARDLR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDMARDLR) SetBits(mask DMARDLR)          { r.U32.SetBits(uint32(mask)) }
func (r *RDMARDLR) ClearBits(mask DMARDLR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RDMARDLR) Load() DMARDLR                 { return DMARDLR(r.U32.Load()) }
func (r *RDMARDLR) Store(b DMARDLR)               { r.U32.Store(uint32(b)) }

type RMDMARDLR struct{ mmio.UM32 }

func (rm RMDMARDLR) Load() DMARDLR   { return DMARDLR(rm.UM32.Load()) }
func (rm RMDMARDLR) Store(b DMARDLR) { rm.UM32.Store(uint32(b)) }

type IDR uint32

type RIDR struct{ mmio.U32 }

func (r *RIDR) LoadBits(mask IDR) IDR { return IDR(r.U32.LoadBits(uint32(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U32.Load()) }
func (r *RIDR) Store(b IDR)           { r.U32.Store(uint32(b)) }

type RMIDR struct{ mmio.UM32 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM32.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM32.Store(uint32(b)) }

type SSIC_VERSION_ID uint32

type RSSIC_VERSION_ID struct{ mmio.U32 }

func (r *RSSIC_VERSION_ID) LoadBits(mask SSIC_VERSION_ID) SSIC_VERSION_ID {
	return SSIC_VERSION_ID(r.U32.LoadBits(uint32(mask)))
}
func (r *RSSIC_VERSION_ID) StoreBits(mask, b SSIC_VERSION_ID) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *RSSIC_VERSION_ID) SetBits(mask SSIC_VERSION_ID)   { r.U32.SetBits(uint32(mask)) }
func (r *RSSIC_VERSION_ID) ClearBits(mask SSIC_VERSION_ID) { r.U32.ClearBits(uint32(mask)) }
func (r *RSSIC_VERSION_ID) Load() SSIC_VERSION_ID          { return SSIC_VERSION_ID(r.U32.Load()) }
func (r *RSSIC_VERSION_ID) Store(b SSIC_VERSION_ID)        { r.U32.Store(uint32(b)) }

type RMSSIC_VERSION_ID struct{ mmio.UM32 }

func (rm RMSSIC_VERSION_ID) Load() SSIC_VERSION_ID   { return SSIC_VERSION_ID(rm.UM32.Load()) }
func (rm RMSSIC_VERSION_ID) Store(b SSIC_VERSION_ID) { rm.UM32.Store(uint32(b)) }

type DR uint32

type RDR struct{ mmio.U32 }

func (r *RDR) LoadBits(mask DR) DR  { return DR(r.U32.LoadBits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

type RX_SAMPLE_DELAY uint32

type RRX_SAMPLE_DELAY struct{ mmio.U32 }

func (r *RRX_SAMPLE_DELAY) LoadBits(mask RX_SAMPLE_DELAY) RX_SAMPLE_DELAY {
	return RX_SAMPLE_DELAY(r.U32.LoadBits(uint32(mask)))
}
func (r *RRX_SAMPLE_DELAY) StoreBits(mask, b RX_SAMPLE_DELAY) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *RRX_SAMPLE_DELAY) SetBits(mask RX_SAMPLE_DELAY)   { r.U32.SetBits(uint32(mask)) }
func (r *RRX_SAMPLE_DELAY) ClearBits(mask RX_SAMPLE_DELAY) { r.U32.ClearBits(uint32(mask)) }
func (r *RRX_SAMPLE_DELAY) Load() RX_SAMPLE_DELAY          { return RX_SAMPLE_DELAY(r.U32.Load()) }
func (r *RRX_SAMPLE_DELAY) Store(b RX_SAMPLE_DELAY)        { r.U32.Store(uint32(b)) }

type RMRX_SAMPLE_DELAY struct{ mmio.UM32 }

func (rm RMRX_SAMPLE_DELAY) Load() RX_SAMPLE_DELAY   { return RX_SAMPLE_DELAY(rm.UM32.Load()) }
func (rm RMRX_SAMPLE_DELAY) Store(b RX_SAMPLE_DELAY) { rm.UM32.Store(uint32(b)) }

type SPI_CTRLR0 uint32

type RSPI_CTRLR0 struct{ mmio.U32 }

func (r *RSPI_CTRLR0) LoadBits(mask SPI_CTRLR0) SPI_CTRLR0 {
	return SPI_CTRLR0(r.U32.LoadBits(uint32(mask)))
}
func (r *RSPI_CTRLR0) StoreBits(mask, b SPI_CTRLR0) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSPI_CTRLR0) SetBits(mask SPI_CTRLR0)      { r.U32.SetBits(uint32(mask)) }
func (r *RSPI_CTRLR0) ClearBits(mask SPI_CTRLR0)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSPI_CTRLR0) Load() SPI_CTRLR0             { return SPI_CTRLR0(r.U32.Load()) }
func (r *RSPI_CTRLR0) Store(b SPI_CTRLR0)           { r.U32.Store(uint32(b)) }

type RMSPI_CTRLR0 struct{ mmio.UM32 }

func (rm RMSPI_CTRLR0) Load() SPI_CTRLR0   { return SPI_CTRLR0(rm.UM32.Load()) }
func (rm RMSPI_CTRLR0) Store(b SPI_CTRLR0) { rm.UM32.Store(uint32(b)) }

func (p *Periph) AITM() RMSPI_CTRLR0 {
	return RMSPI_CTRLR0{mmio.UM32{&p.SPI_CTRLR0.U32, uint32(AITM)}}
}

func (p *Periph) ADDR_LENGTH() RMSPI_CTRLR0 {
	return RMSPI_CTRLR0{mmio.UM32{&p.SPI_CTRLR0.U32, uint32(ADDR_LENGTH)}}
}

func (p *Periph) INST_LENGTH() RMSPI_CTRLR0 {
	return RMSPI_CTRLR0{mmio.UM32{&p.SPI_CTRLR0.U32, uint32(INST_LENGTH)}}
}

func (p *Periph) WAIT_CYCLES() RMSPI_CTRLR0 {
	return RMSPI_CTRLR0{mmio.UM32{&p.SPI_CTRLR0.U32, uint32(WAIT_CYCLES)}}
}

type XIP_MODE_BITS uint32

type RXIP_MODE_BITS struct{ mmio.U32 }

func (r *RXIP_MODE_BITS) LoadBits(mask XIP_MODE_BITS) XIP_MODE_BITS {
	return XIP_MODE_BITS(r.U32.LoadBits(uint32(mask)))
}
func (r *RXIP_MODE_BITS) StoreBits(mask, b XIP_MODE_BITS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXIP_MODE_BITS) SetBits(mask XIP_MODE_BITS)      { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_MODE_BITS) ClearBits(mask XIP_MODE_BITS)    { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_MODE_BITS) Load() XIP_MODE_BITS             { return XIP_MODE_BITS(r.U32.Load()) }
func (r *RXIP_MODE_BITS) Store(b XIP_MODE_BITS)           { r.U32.Store(uint32(b)) }

type RMXIP_MODE_BITS struct{ mmio.UM32 }

func (rm RMXIP_MODE_BITS) Load() XIP_MODE_BITS   { return XIP_MODE_BITS(rm.UM32.Load()) }
func (rm RMXIP_MODE_BITS) Store(b XIP_MODE_BITS) { rm.UM32.Store(uint32(b)) }

type XIP_INCR_INST uint32

type RXIP_INCR_INST struct{ mmio.U32 }

func (r *RXIP_INCR_INST) LoadBits(mask XIP_INCR_INST) XIP_INCR_INST {
	return XIP_INCR_INST(r.U32.LoadBits(uint32(mask)))
}
func (r *RXIP_INCR_INST) StoreBits(mask, b XIP_INCR_INST) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXIP_INCR_INST) SetBits(mask XIP_INCR_INST)      { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_INCR_INST) ClearBits(mask XIP_INCR_INST)    { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_INCR_INST) Load() XIP_INCR_INST             { return XIP_INCR_INST(r.U32.Load()) }
func (r *RXIP_INCR_INST) Store(b XIP_INCR_INST)           { r.U32.Store(uint32(b)) }

type RMXIP_INCR_INST struct{ mmio.UM32 }

func (rm RMXIP_INCR_INST) Load() XIP_INCR_INST   { return XIP_INCR_INST(rm.UM32.Load()) }
func (rm RMXIP_INCR_INST) Store(b XIP_INCR_INST) { rm.UM32.Store(uint32(b)) }

type XIP_WRAP_INST uint32

type RXIP_WRAP_INST struct{ mmio.U32 }

func (r *RXIP_WRAP_INST) LoadBits(mask XIP_WRAP_INST) XIP_WRAP_INST {
	return XIP_WRAP_INST(r.U32.LoadBits(uint32(mask)))
}
func (r *RXIP_WRAP_INST) StoreBits(mask, b XIP_WRAP_INST) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXIP_WRAP_INST) SetBits(mask XIP_WRAP_INST)      { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_WRAP_INST) ClearBits(mask XIP_WRAP_INST)    { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_WRAP_INST) Load() XIP_WRAP_INST             { return XIP_WRAP_INST(r.U32.Load()) }
func (r *RXIP_WRAP_INST) Store(b XIP_WRAP_INST)           { r.U32.Store(uint32(b)) }

type RMXIP_WRAP_INST struct{ mmio.UM32 }

func (rm RMXIP_WRAP_INST) Load() XIP_WRAP_INST   { return XIP_WRAP_INST(rm.UM32.Load()) }
func (rm RMXIP_WRAP_INST) Store(b XIP_WRAP_INST) { rm.UM32.Store(uint32(b)) }

type XIP_CTRL uint32

type RXIP_CTRL struct{ mmio.U32 }

func (r *RXIP_CTRL) LoadBits(mask XIP_CTRL) XIP_CTRL { return XIP_CTRL(r.U32.LoadBits(uint32(mask))) }
func (r *RXIP_CTRL) StoreBits(mask, b XIP_CTRL)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXIP_CTRL) SetBits(mask XIP_CTRL)           { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_CTRL) ClearBits(mask XIP_CTRL)         { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_CTRL) Load() XIP_CTRL                  { return XIP_CTRL(r.U32.Load()) }
func (r *RXIP_CTRL) Store(b XIP_CTRL)                { r.U32.Store(uint32(b)) }

type RMXIP_CTRL struct{ mmio.UM32 }

func (rm RMXIP_CTRL) Load() XIP_CTRL   { return XIP_CTRL(rm.UM32.Load()) }
func (rm RMXIP_CTRL) Store(b XIP_CTRL) { rm.UM32.Store(uint32(b)) }

type XIP_SER uint32

type RXIP_SER struct{ mmio.U32 }

func (r *RXIP_SER) LoadBits(mask XIP_SER) XIP_SER { return XIP_SER(r.U32.LoadBits(uint32(mask))) }
func (r *RXIP_SER) StoreBits(mask, b XIP_SER)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXIP_SER) SetBits(mask XIP_SER)          { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_SER) ClearBits(mask XIP_SER)        { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_SER) Load() XIP_SER                 { return XIP_SER(r.U32.Load()) }
func (r *RXIP_SER) Store(b XIP_SER)               { r.U32.Store(uint32(b)) }

type RMXIP_SER struct{ mmio.UM32 }

func (rm RMXIP_SER) Load() XIP_SER   { return XIP_SER(rm.UM32.Load()) }
func (rm RMXIP_SER) Store(b XIP_SER) { rm.UM32.Store(uint32(b)) }

type XRXOICR uint32

type RXRXOICR struct{ mmio.U32 }

func (r *RXRXOICR) LoadBits(mask XRXOICR) XRXOICR { return XRXOICR(r.U32.LoadBits(uint32(mask))) }
func (r *RXRXOICR) StoreBits(mask, b XRXOICR)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXRXOICR) SetBits(mask XRXOICR)          { r.U32.SetBits(uint32(mask)) }
func (r *RXRXOICR) ClearBits(mask XRXOICR)        { r.U32.ClearBits(uint32(mask)) }
func (r *RXRXOICR) Load() XRXOICR                 { return XRXOICR(r.U32.Load()) }
func (r *RXRXOICR) Store(b XRXOICR)               { r.U32.Store(uint32(b)) }

type RMXRXOICR struct{ mmio.UM32 }

func (rm RMXRXOICR) Load() XRXOICR   { return XRXOICR(rm.UM32.Load()) }
func (rm RMXRXOICR) Store(b XRXOICR) { rm.UM32.Store(uint32(b)) }

type XIP_CNT_TIME_OUT uint32

type RXIP_CNT_TIME_OUT struct{ mmio.U32 }

func (r *RXIP_CNT_TIME_OUT) LoadBits(mask XIP_CNT_TIME_OUT) XIP_CNT_TIME_OUT {
	return XIP_CNT_TIME_OUT(r.U32.LoadBits(uint32(mask)))
}
func (r *RXIP_CNT_TIME_OUT) StoreBits(mask, b XIP_CNT_TIME_OUT) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *RXIP_CNT_TIME_OUT) SetBits(mask XIP_CNT_TIME_OUT)   { r.U32.SetBits(uint32(mask)) }
func (r *RXIP_CNT_TIME_OUT) ClearBits(mask XIP_CNT_TIME_OUT) { r.U32.ClearBits(uint32(mask)) }
func (r *RXIP_CNT_TIME_OUT) Load() XIP_CNT_TIME_OUT          { return XIP_CNT_TIME_OUT(r.U32.Load()) }
func (r *RXIP_CNT_TIME_OUT) Store(b XIP_CNT_TIME_OUT)        { r.U32.Store(uint32(b)) }

type RMXIP_CNT_TIME_OUT struct{ mmio.UM32 }

func (rm RMXIP_CNT_TIME_OUT) Load() XIP_CNT_TIME_OUT   { return XIP_CNT_TIME_OUT(rm.UM32.Load()) }
func (rm RMXIP_CNT_TIME_OUT) Store(b XIP_CNT_TIME_OUT) { rm.UM32.Store(uint32(b)) }

type ENDIAN uint32

type RENDIAN struct{ mmio.U32 }

func (r *RENDIAN) LoadBits(mask ENDIAN) ENDIAN { return ENDIAN(r.U32.LoadBits(uint32(mask))) }
func (r *RENDIAN) StoreBits(mask, b ENDIAN)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RENDIAN) SetBits(mask ENDIAN)         { r.U32.SetBits(uint32(mask)) }
func (r *RENDIAN) ClearBits(mask ENDIAN)       { r.U32.ClearBits(uint32(mask)) }
func (r *RENDIAN) Load() ENDIAN                { return ENDIAN(r.U32.Load()) }
func (r *RENDIAN) Store(b ENDIAN)              { r.U32.Store(uint32(b)) }

type RMENDIAN struct{ mmio.UM32 }

func (rm RMENDIAN) Load() ENDIAN   { return ENDIAN(rm.UM32.Load()) }
func (rm RMENDIAN) Store(b ENDIAN) { rm.UM32.Store(uint32(b)) }
