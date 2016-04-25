package main
import "fmt"
import "math"
import "bufio"
import "os"
import "strconv"
import "strings"
import "bytes"

func main() {
    in := bufio.NewReader(os.Stdin)
    input, _ := in.ReadString('\n')
    numberOfPlays, _ := strconv.Atoi(strings.TrimSpace(input))
    var plays uint64
    for ; numberOfPlays > 0; numberOfPlays-- {
        input, _ := in.ReadString('\n')
        plays, _ = strconv.ParseUint(strings.TrimSpace(input), 10, 64)
        winner := play(plays)
        fmt.Println(winner)
    }
}

func play(n uint64) string {
    i := 1
    for ;n > 1; i++ {
        if isPowerOf2(n) {
            n -= n >> 1
        } else {
            n -= getNearestN(n)
        }
    }
    if math.Pow(-1.0, float64(i)) > 0 {
        return "Louise"
    } else {
        return "Richard"
    }
}

func getNearestN(n uint64) uint64 {
    binary := strconv.FormatUint(n, 2)
    var buffer bytes.Buffer
    firstBitSet := false
    for _, bit := range binary {
        if !firstBitSet && bit == 49 {
            buffer.WriteByte('1')
            firstBitSet = true
        } else {
            buffer.WriteByte('0')
        }
    }
    result, _ := strconv.ParseUint(buffer.String(), 2, 64)
    return result
}

func isPowerOf2(n uint64) bool {
    binary := strconv.FormatUint(n, 2)
    singleBitSet := false
    for _, bit := range binary {
        if bit == 49 {
            if singleBitSet {
                return false
            } else {
                singleBitSet = true
            }
        }
    }
    return singleBitSet
}
