from itertools import chain, combinations
from operator import ior
from functools import reduce

def powerset(iterable):
    s = list(iterable)
    return chain.from_iterable(combinations(s, r) for r in range(len(s)+1))

def ffs(x):
    return (x&-x).bit_length()-1

def bitset_get_set_bits(x):
    bits = []
    while x != 0:
        bit = 1 << ffs(x)
        bits.append(bit)
        x &= ~bit
    return bits

def bitset_powerset(x):
    bits = bitset_get_set_bits(x)
    pset = powerset(bits)
    result = []
    for bitset in pset:
        item = 0
        for bit in bitset:
            item |= bit
        result.append(item)
    return result

down_opp_mask = 0x0000010101010101

down = list(bitset_powerset(down_opp_mask))

assert len(set(down)) == 64

def down_hash(x):
    return x % 127

assert len(set(down_hash(x)&0x3F for x in down)) == 64

down_flip_table = [0] * 64

for item in down:
    index = down_hash(item) & 0x3F
    flipped = ((1 << ffs(item + 1)) - 1) & down_opp_mask