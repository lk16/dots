package bitset

import "testing"

func genTestBitsets() (ch chan Bitset) {
    ch = make(chan Bitset)
    go func() {
        ch <- Bitset(0)
        ch <- ^Bitset(0)
        for i:=uint(0); i<64; i++ {
            ch <- Bitset(1) << i
        }
        for i:=uint(0); i<64; i++ {
            ch <- ^(Bitset(1) << i)
        }
        for i:=0; i<100000; i++ {
            ch <- RandomBitset()
        }
        close(ch)
    }()
    return
}

func (bs Bitset) count() (count uint) {
    for i := uint(0); i<64; i++ {
        if bs & (1 << i) != 0 {
            count++
        } 
    }
    return
}

func TestBitsetCount(t *testing.T) {
    for bs := range genTestBitsets() {
        expected := bs.count()
        got := bs.Count()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}

func (bs Bitset) firstBit() (Bitset) {
    for mask:= Bitset(1); mask!=0; mask<<=1 {
        if bs & mask != 0 {
            return mask
        } 
    }
    return 0
}

func TestBitsetFirstBit(t *testing.T) {
    for bs := range genTestBitsets() {
        expected := bs.firstBit()
        got := bs.FirstBit()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}

func (bs Bitset) lastBit() Bitset {
    for mask:= Bitset(1 << 63); mask!=0; mask>>=1 {
        if bs & mask != 0 {
            return mask
        } 
    }
    return 0
}

func TestBitsetLastBit(t *testing.T) {
    for bs := range genTestBitsets() {
        expected := bs.lastBit()
        got := bs.LastBit()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}

func (bs Bitset) testBit(index uint) bool {
    return bs & Bitset(1 << index) != 0
}

func TestBitsetTestBit(t *testing.T) {
    for bs := range genTestBitsets() {
        for index:=uint(0); index<64; index++ {
            expected := bs.testBit(index)
            got := bs.TestBit(index)
            if expected != got {
                t.Errorf("Expected %d but got %d\n",expected,got)
            }
        }
    }
}

func (bs Bitset) firstBitIndex() uint {
    for index:=uint(0); index<64; index++ {
        if bs & Bitset(1 << index) != 0 {
            return index
        } 
    }
    return 0
}

func TestBitsetFirstBitIndex(t *testing.T) {
    for bs := range genTestBitsets() {
        expected := bs.firstBitIndex()
        got := bs.firstBitIndex()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}