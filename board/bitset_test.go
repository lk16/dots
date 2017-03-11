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

func (bs bitset) testCount() (count uint) {
    for i := uint(0); i<64; i++ {
        if bs & (1 << i) != 0 {
            count++
        } 
    }
    return
}

func TestBitsetCount(t *testing.T) {
    ch := genTestBitsets()
    for v := range ch {
        expected := v.testCount()
        got := v.Count()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}

func (bs bitset) testFirstBit() (bitset) {
    for mask:= bitset(1); mask!=0; mask<<=1 {
        if bs & mask != 0 {
            return mask
        } 
    }
    return 0
}

func TestBitsetFirstBit(t *testing.T) {
    ch := genTestBitsets()
    for v := range ch {
        expected := v.testFirstBit()
        got := v.FirstBit()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}

func (bs bitset) testLastBit() (bitset) {
    for mask:= bitset(1) << 63; mask!=0; mask>>=1 {
        if bs & mask != 0 {
            return mask
        } 
    }
    return 0
}

func TestBitsetLastBit(t *testing.T) {
    ch := genTestBitsets()
    for v := range ch {
        expected := v.testLastBit()
        got := v.LastBit()
        if expected != got {
            t.Errorf("Expected %d but got %d\n",expected,got)
        }
    }
}