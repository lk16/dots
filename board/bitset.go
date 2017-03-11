package board

type bitset uint64

func (bs bitset) Count() uint {
    const (
        m1  = 0x5555555555555555 //binary: 0101...
        m2  = 0x3333333333333333 //binary: 00110011..
        m4  = 0x0f0f0f0f0f0f0f0f //binary:  4 zeros,  4 ones ...
        h01 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3...
    )
    bs -= (bs >> 1) & m1             //put count of each 2 bits into those 2 bits
    bs = (bs & m2) + ((bs >> 2) & m2) //put count of each 4 bits into those 4 bits
    bs = (bs + (bs >> 4)) & m4        //put count of each 8 bits into those 8 bits
    return uint((bs * h01) >> 56)    //returns left 8 bits of x + (x<<8) + (x<<16) + (x<<24) + ...
}

func (bs bitset) Test(index uint) bool {
    mask := bitset(1) << index
    return bs & mask != 0
}

func (bs bitset) FirstBit() bitset {
    return bs & -bs
}


func (bs bitset) LastBit() bitset {
    for mask := bitset(1 << 63); mask!=0; mask>>=1 {
        if bs & mask != 0 {
            return mask
        }
    }
    return 0
}