package leetcode

// 1009. Complement of Base 10 Integer

func bitwiseComplement(N int) int {
    if N == 0 {
        return 1
    }
    t := N
    c := 0
    for t > 0 {
        t /= 2
        c++
    }
    
    b := ^((-1) << c)
    return N ^ b
}
