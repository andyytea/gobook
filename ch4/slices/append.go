package slices

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // sufficient room to extend slice
        z = x[:zlen] 
    } else {
        // insufficient space to extend, realloc new array
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z,x)
    }
    z[len(x)] = y
    return z
}
