package board

import "testing"

func TestBitsetCount(t *testing.T) {
    lame_count := func (bs bitset) (count uint) {
        for i := uint(0); i<64; i++ {
            if bs & (1 << i) != 0 {
                count++
            } 
        }
        return count
    }

    test_values := []bitset{
        0,
        0xFFFFFFFFFFFFFFFF,
        0x1,
        0xF,
        0x8000000000000000,
        0x123FFFEBCD837215}

    for _,v := range test_values {
        expected := lame_count(v)
        got := v.Count()
        if expected != got {
            t.Errorf("Expected",expected," but got",got)
        }
    }

}