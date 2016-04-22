package main
import "fmt"
import "math"
import "bufio"
import "os"
import "strconv"
import "strings"

func main() {
    in := bufio.NewReader(os.Stdin)
    input, _ := in.ReadString('\n')
    numberOfPlays, _ := strconv.Atoi(strings.TrimSpace(input))
    var plays uint64
    for ; numberOfPlays > 0; numberOfPlays-- {
        input, _ := in.ReadString('\n')
        plays, _ = strconv.ParseUint(strings.TrimSpace(input), 2, 64)
        winner := play(plays)
        fmt.Println(winner)
    }
}

func play(n uint64) string {
    i := 1
    for ;n > 1; i++ {
        if isPowerOf2(n) {
            n /= 2
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
    for ; !isPowerOf2(n); n-- {
        
    }
    return n
}

func isPowerOf2(n uint64) bool {
    var power uint64 = 2
    for ;; power <<= 1 {
        if power > n {
            return false
        }
        if power == n {
            return true
        }
    }
}
