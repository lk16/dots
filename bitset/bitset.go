package bitset

import (
    "math/rand"
)

type Bitset uint64

func (bs Bitset) Count() uint {
    const (
        m1  = 0x5555555555555555
        m2  = 0x3333333333333333
        m4  = 0x0f0f0f0f0f0f0f0f
        h01 = 0x0101010101010101
    )
    bs -= (bs >> 1) & m1
    bs = (bs & m2) + ((bs >> 2) & m2)
    bs = (bs + (bs >> 4)) & m4
    return uint((bs * h01) >> 56)
}

func (bs Bitset) TestBit(index uint) bool {
    mask := Bitset(1 << index)
    return bs & mask != 0
}

func (bs Bitset) FirstBit() Bitset {
    return bs & -bs
}

func (bs Bitset) FirstBitIndex() uint {
    
    magictable := [67]uint{
         0,  0,  1, 39,  2, 15, 40, 23,
         3, 12, 16, 59, 41, 19, 24, 54,
         4,  0, 13, 10, 17, 62, 60, 28,
        42, 30, 20, 51, 25, 44, 55, 47,
         5, 32,  0, 38, 14, 22, 11, 58,
        18, 53, 63,  9, 61, 27, 29, 50,
        43, 46, 31, 37, 21, 57, 52,  8,
        26, 49, 45, 36, 56,  7, 48, 35,
         6, 34, 33}

    return magictable[bs.FirstBit() % 67];
}

func (bs Bitset) LastBit() Bitset {
    for mask := Bitset(1 << 63); mask!=0; mask>>=1 {
        if bs & mask != 0 {
            return mask
        }
    }
    return 0
}

func RandomBitset() Bitset {
    return Bitset(rand.Uint64())
}