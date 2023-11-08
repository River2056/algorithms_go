package determineifacellisreachable

import "math"

func IsReadhableAtTime(sx, sy, fx, fy, t int) bool {
    minDist := math.Max(math.Abs(float64(sx) - float64(fx)), math.Abs(float64(sy) - float64(fy)))
    if minDist == 0 {
        if t == 1 {
            return false
        } else {
            return true
        }
    }

    return t >= int(minDist)
}
