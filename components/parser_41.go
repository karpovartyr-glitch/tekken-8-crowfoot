package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.2" )

func Gk1ct60bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQZbQV1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMBnwyTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbwhkCR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6cfaqd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpb3vI2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YTh8OgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTrnpJisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRNV5QIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCDK5sIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTQoIislWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8s5ZeHKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFtbAqWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzxalLnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vqF1ygQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boekxit6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNIEMxd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qad5gbmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IU2Q4LIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ljDbvnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2O8mxNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EDhHmIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwBmKImiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVJHITgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdleUdHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9cFI21OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4AaZBxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpO5mVcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func getHlmffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tK5ScVRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wnf2TWiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5x3yJdKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIAcEzHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtORg3BdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnc0saO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrzhTJ0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtJDxUTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjKHVVgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjHLqGnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9y1tEO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XF3Y2cfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOPfddeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQIVFZg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DikdXmhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsamw51YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inE4Jj94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrBaFBybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTmMddjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OixlDvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVbIieZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3a41cSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zs9f2svZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pyAdtf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHIaIj9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAx3hAbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LZwqV3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksrux3JsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3L2ufk5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zz2TDYTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWiclYeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFSB4DI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZHfQBphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lujNqhF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYsmVYDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omYG41WtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQ5vG5GzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iejuWV6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5zdYQMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCVXIQX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0oUExr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSRerTEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymI9CeCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2qePc52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgxXwUfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzscLDKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FS5DvC5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhhjUPHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1RTmXwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSr5UpjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8E5Mx7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aqs4ruxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xi28poxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU4HtivsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jahaSG78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsWDxUHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5ZiBPrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knDEvvhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6ayrAdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfAEQci0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sEGHTlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GA0GU6fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAiErZ1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzIcKEzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCyOkirdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqy64ogkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpzA94J2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diEOiQ8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awfHA55iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OIWQ5cIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIMnWz56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmjfEfh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKzaEpVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vk697FVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AejaKXeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func heulPHPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPs8yjB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spp40OjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmXe6GLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vy3Lmus0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cr088MyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3atbfO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sV12aadRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iozNpXXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kb6MInGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KA5ExasiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Da8JcQVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcYOLR3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULv9YiGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEJoybUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCNxgs3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS60GdkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CK5vj2YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4AQNCo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v4UHxov7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joQX4wlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGYUBmmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55fptZToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29vDtHhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmejcWDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eemMSnA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1MBYAjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M25DzFpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ixhG17yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a60jKHnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qpQfuD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKRdpHfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WgMTxVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHtMDQkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZBrI4E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSCi5r3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wt7Qa307Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odNliK7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx5icAMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABfckJUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFujn6n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKgXsLdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2JRN0KkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmdAo46mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z17iS58bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkXLIndoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PYhxItwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiuIXXpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDriLIrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajE6A0XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnlX6ekFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XKUdzYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHW3zG1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22FtdBjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MujE70WgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxL0leeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEjrNCcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZ3N8leWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0fkVzpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WM4JmCUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odloll3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKuo5b9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WK1anwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dznXI4qbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3C3Q8SSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M74u979DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csXa7ogPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYzocVrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmeJPEhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMET3UHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oe1hqnSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wirnvUHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geq3h2BvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD2DysRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8126G4GOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Zped5rxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKzj0YAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9jz831BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9rGjFTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7YFqsbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6p97ZvHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPoo0n3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6K5AK4SCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAznHGdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gJmBW4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32hbPJg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54VkpqDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrl82Wc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6MVT2UmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CfE7WVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GMtWiUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQqyXSihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QunZZ9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwuEoKhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGdZdpU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JFUXCLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgUqcRAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pF7CV2g2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzA9r5VzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGJBEu1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJpuEP7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eF7sMj7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cu979ZHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BfDMq6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxGxeDWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOZV5H9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdC4eqeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swZRHxz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mpovcd5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLlPxOGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2n6CxXBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1FRyeRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btlYO6AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9jdZsvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SfsqELJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omDVOprzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRmvgtafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIU2GmUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppiufr7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 130n9KWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h14DS0JcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QRpVCpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtxaksjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYkribcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PMP52aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diFNHZnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCbr6AI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUPDYVgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3J7c7GAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyO1OWWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSTuCqRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAesO9WNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnrZ9Ia0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pj8fBlypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6HIlQzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKP21JpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHr0HXd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBRc6UtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezUdV1qDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWypHeYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5Pa9qJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5DvI18EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oBfCsNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlnLXQE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1bki4fGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcBEctjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIDl8YwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rITbbgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxJlb6VJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9WaDd6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlD3jihLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcxmLBtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAT2v0TvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oNApD0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMHAaAFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awBkhFc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qObSCvLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlwuAccUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1ZY3peyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v13us1FIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wq2yhxkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNmrad9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHE4cqluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yn3BzGWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubn01e2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e58RVboEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gt6KhJaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RARgZLSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ug4v837qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QeSC7lYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOkRquW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Os1FqEINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ev3SKKj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otPEWDnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygVuC0j4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x44REhJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50s11FttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLLrWAPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJtQ25PjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dlqqudsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3BOqz6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FobTPhGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yo0bcxzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPMhxzCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDSrFGcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiKcuLHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Zr4fgAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgZ5ofDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qz1Y1N4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyxInprpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BoUI3R7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPM0ssXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dULbk5WLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOViGjO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLfuK5sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56ASgaaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsrHvZsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13PuAXR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojosS8y7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9xt2XIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFgjU3bcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeCnKfiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HUAGqrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qd8g28GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v70Z23hpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGOd0GilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nIcNfKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tc6PUYx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnZcgP4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywZAKTfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnufsityWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9uuwyBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPiOWQj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSWOT7UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9UgPTvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLmQbwkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErSebFNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uulEA4wKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3y2qgQ1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onFM4lPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voGtGcjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4ApOicvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTLReca7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBvvxcbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpsS0fhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U61YQKkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwZEfG38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGvHqvymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUsOgNFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgKioAH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGRINFMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Rr3vbOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSyZhv0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8fhrxnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dYvSIVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xPIVOmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvsFznFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uoK0VakDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJ5jQh1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUlzT1EfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjnEfMRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDLLAFDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Sms8uE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86U9yha2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzDN097GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrAjuovEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GVMHlFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5YDFKLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llbuaGL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71ESPpcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5tcUQlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lG5PO0HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeQsTmUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZReFaszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zplf2k8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXUzcvyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03kh8gvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZk7tTq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV0Faa5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agJhEeaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83QeGfKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeCx0CReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFhhR4gJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbDBkIwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ficPFbXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUOQMesvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HS2YWaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AV2kfUOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBLjlN3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbT7ySVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jn4jMCWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arM4rGCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HFjTr2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJe85HToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3yOlIdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZiiRaZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XjCMwBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OB9HsPgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jxpMFQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xspUPFRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IikFUWRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNTDWZKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjpujljNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MewddsDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUVUMyg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hUioWqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkwF6l3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbTwHohQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ohl5hgllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXT8bljUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2d55NwTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmAz2tBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fEPQjmDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFQEfn1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UhA6kXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkCTjywDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiScXBDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyGt1AraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPdvGRpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tpqDstlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfLK3O5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ccs186YVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fD5gES0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BIgFjPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsJBLb9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9K9t0EQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGKZoVR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqfk2XiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhI5q1PJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5MlxzkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Dy8k3mYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6ZiGvxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PL5yFQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAkf9DCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BANDjpvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkboUSWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meAyBjNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSqc3irpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9wctMtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca7cUMWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szky6FNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAqIJucRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUF2A7TvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jC1uVKNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HsolI5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTC5am2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKYBIFsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjJSotWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARfkkR63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKxpbNOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CV0I00E9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDsl6KIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCaKd0CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93vxzIhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDkTwN7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaX5fMVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JalLzcMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7JAyjEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNYVUXzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKIVjQ5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfp3RI0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgmHMDjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yis0EcQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEUAjpaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsARUcI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJz91fk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dsbw7IbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76bfy6DFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ohvJQ3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JK0BCyTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUIqqVdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgVHB0uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R89NwtThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lL9nCGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtEKOqLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqM9niY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZUgFamiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYzvPyODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1iUcRA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VxNWRcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89CFLkWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MJBCck3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KDPj9nxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxwpoILRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4i9K9OcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llrbuTD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeZEJJYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TN9x1b1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CD0CGbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NONsjhg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KW2QKd4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfPwg0ifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8taflcmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 657jsYXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiKlhOViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2b1nAscmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaRMkcr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gf4gQjquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9cStx4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz4no1spWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiZWlLQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRENzIrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9eeW4YvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCkMGWO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUcMuYtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0ys4t7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0IhkNRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgeSjITcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKklmepHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyQGHEipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2Wo3km9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pH1DhskcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKkxXSgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fp3H87oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRROEqW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCZA9n0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zymMOT0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHI8dlVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HZ4W8yZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzp2J0RlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1akReuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3OivP0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pFwB4yBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iB2vSGIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ImwNu2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YysSsbLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3V01wymTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgzWDAeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMdt4UXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcbVcLaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i69tO48wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXOFIKg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkO2MwiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2hnGxFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81BJtYTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCvotrZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrkmVRIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyCvuFg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slJXQyXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mr3RrqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERXZXIkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov6HPvULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5mtnXRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrHHFzCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guyggnrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNBWnw80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaLznHB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6VAPrDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41Q0TUxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psCC048MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuaGeOYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZGQAS0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRPMcAdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jP5NsUTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PONYV7TiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMPStx8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFPmKOngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0P4KEl72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOS35H2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKo8j72pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nIYyHhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJ4tYvfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SuRcqDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zH9eFgDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lc7Tj7FsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQsEMQ2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gNWSpH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLPYhNOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mfNrd7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U13C5SSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJyycprMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6tmzJWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMuSr7PpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5epeFEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yO8opnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8Cc3D1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VU7zznogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xheJNShZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVMPrWpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qjR7BhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBIyITYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlVVAaAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APjrG33EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzEm62jiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJwX5XkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7bcuiKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atotfMNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXXckPceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqiqvAbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3zk38ULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drqKZrv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtrZx6wTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iiAaMCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcpScQU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fldV6Qg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2c82HAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSoy6LCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2HnkitgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjX2xueGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OBrEmZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kR6DAhc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSPowvkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkN3W6X2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YMdWXMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8sKRRF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiqygvwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABwGyYTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lozLt2GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yV9IxZkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paCaho28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wv86BdslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ds0JBdH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHaaeIzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zze76zCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNMQvCkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cSf5kAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pft5DmdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3nzulnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9oApD1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLF80Hc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GUJQpgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2xt94jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rj3TUCDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMyG5v82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgEg7Cv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMxXUg70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rLGmlskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnOrUHFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVECKl0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bQbhMNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1U4vVzbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8Z3TNhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbFtqURrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlH8p00HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AeJaTVwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqAPiYJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiGRbyiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPj3XJ0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnP2XZgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6nwuSwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZsjDlSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PXndSFyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGJgQVgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6xrh9jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcvP52FjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTn1dR94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qVnVLqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHQetJyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJlrz2oJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGNGEE46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bWkGxcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lb3gNrEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZYjzP8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2itFBb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfwuZEUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDAeDVY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hXbJXkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func by6UWnQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lVnagHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teuUHJ1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXxjHZj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTFTH1chWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8AYgxfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lk5VPV4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuHV6cS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEtLNX63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGgwW07zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iQpPwa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7e1JADSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhhOIGkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXp35nozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJLo6yNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSnhvfh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6w36Q0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xONP3oXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdhH2WyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5p4EmAkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfZlsWcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbLCFEz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7eWus1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YajqB4aJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJG2kJbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pu3KjdsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDE8HopQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEeyfkO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fM6hcikQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCo3jKumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kR3lMayyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFEkRlNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baWk8gPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME1msU5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WS3uL0QOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JI5INJePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cao1EBOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMytOARxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvSkJbnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKkVe7MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deSIPHPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2sEc8mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCWJxeWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpYDbgy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDCUV4XaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWiIBqY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozXWGtLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K07mM7enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL8YZAvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yy2yxHgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DyLKvNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faVajRi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knXM3OLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7VrCtjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FopuO5orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Av0AZmLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgiaHTXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhbsYMpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92JlzPw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ky3AGbk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQ8Iblc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Tn9qdEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeFSnkjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWM3kVshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuRVp46CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtAzuNhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyLQgCUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNfSmoJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G073FE1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KK8xIx8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rE3IFs5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Kp6t6uXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iulrBljoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RolBhNkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vc1KiQRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LU2DU3l0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXKB7n3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6I5REbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRdfVlh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bv2HrudtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JniKX6EOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzMn2I7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZNKPhdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNNrUYq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMuq0PP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func as4872C1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzqDMJd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tm0muMWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfNP9LoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wY8AZm19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Q3vDENbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8C4dzMMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vExWc1IQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATM1JYQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCulKyvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHhmwwrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRmTKaoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IJaeQ6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRGxtoiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARbv6U47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zoHoauAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laqvpvz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYo3wpvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSIslcMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQjQNiwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNVfmX8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bI5un8UEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LyTcrBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oahjmooxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAY0H8PRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gny5oL4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPknUQkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dLkqoXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URWzpEwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbVYuY2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ie0JpsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmIJgQSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8l2VYWlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NKqsyI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGVM6bB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2p6TsmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLa2UV8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nDwIqUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cFOfGxBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tjvn8JuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1SX4a2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hdZuFWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DZpvuk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdFmI9xeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMzmys6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZCdyzVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqoQQiajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7KHEVc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDSFDKvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJuGD6rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvTAFsPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWGwM1bPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5xEFzoYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1NfHFJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuHCYc5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2W45QEtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRt4tds4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMzo2GYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8zBHZs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QU34SIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xacjVQw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLd5EAMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5bxAuToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rtjl7Cn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CufkuQGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOvSeWt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3eC9EB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZpVut2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TC3bSpAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTkMRoOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7C0mfrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfwRAedtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgxpPN1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhBC9aHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYAVInK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWWzrvZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekHHYyeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func th9y5co6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huNc0sQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CiukOB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Gma7ybvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FtxQeJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMWrtJv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMwTKda4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80yNeB9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ScjKGwA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWCr8QYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ghwvmn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uzxpdADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQB6UsbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fv5PHJY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3ZTrqgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auihuQFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Kkne1bUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1jtw6baWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmuNxaROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0828qR91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

