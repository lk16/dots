package board

import "testing"

func genTestBitsets() (ch chan bitset) {
    ch = make(chan bitset)
    go func() {
        ch <- bitset(0)
        ch <- ^bitset(0)
        for i:=uint(0); i<64; i++ {
            ch <- bitset(1) << i
        }
        for i:=uint(0); i<64; i++ {
            ch <- ^(bitset(1) << i)
        }
        for i:=0; i<100000; i++ {
            ch <- RandomBitset()
        }
        close(ch)
    }()
    return
}

func (bs bitset) count() (count uint) {
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

func (bs bitset) firstBit() (bitset) {
    for mask:= bitset(1); mask!=0; mask<<=1 {
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

func (bs bitset) lastBit() (bitset) {
    for mask:= bitset(1 << 63); mask!=0; mask>>=1 {
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

func (bs bitset) testBit(index uint) bool {
    return bs & bitset(1 << index) != 0
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

func (bs bitset) firstBitIndex() uint {
    for index:=uint(0); index<64; index++ {
        if bs & bitset(1 << index) != 0 {
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