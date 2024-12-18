// 练习 1.5: 修改前面的Lissajous程序里的调色板，由黑色改为绿色。
// 我们可以用color.RGBA{0xRR, 0xGG, 0xBB, 0xff}来得到#RRGGBB这个色值，
// 三个十六进制的字符串分别代表红、绿、蓝像素。

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"             // 导入image包（图像处理）
    "image/color"       // 导入image下的子包color（颜色）
    "image/gif"         // 导入gif包（动图）
    "io"
    "math"              // 导入math包（数学）
    "math/rand"         // 导入rand包（随机数）
    "os"                
    "time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}  // 颜色数组

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
	greenIndex = 2 // 绿色
)

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    // Thanks to Randall McPherson for pointing out the omission.
    rand.Seed(time.Now().UTC().UnixNano())      // 使用当前时间生成随机数
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                greenIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

