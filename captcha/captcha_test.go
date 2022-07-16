/**
 * @Au.charshor Awen
 * @Description
 * @Date 2021/7/20
 **/

package captcha

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"golang.org/x/image/font"
)

/**
 * @Description: 获取当前目录
 * @return string
 */
func getPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path + "/.."
}

// go test -race base.go captcha_test.go
func TestGetCaptchaGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	n := 30
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			t.Logf(">>> %p\n", GetCaptcha())
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestImageSize(t *testing.T) {
	capt := GetCaptcha()
	fmt.Println(capt)
	capt.SetImageSize(Size{Width: 300, Height: 300})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestSetThumbSize(t *testing.T) {
	capt := GetCaptcha()

	capt.SetThumbSize(Size{Width: 300, Height: 300})

	chars := []string{"HE", "CA", "WO", "NE", "HT", "IE", "PG", "GI", "CH", "CO", "DA"}
	_ = capt.SetRangChars(chars)
	// capt.SetImageFontDistort(0)
	// capt.SetImageFontDistort(0)
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	file := getPWD() + "/tests/.cache/" + fmt.Sprintf("%v", RandInt(1, 200)) + "Img.png"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	defer logFile.Close()
	i := strings.Index(b64, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64[i+1:]))
	io.Copy(logFile, dec)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestChars(t *testing.T) {
	capt := GetCaptcha()
	// chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// capt.SetRangChars(strings.Split(chars, ""))
	chars := []string{"HE", "CA", "WO", "NE", "HT", "IE", "PG", "GI", "CH", "CO", "DA"}
	_ = capt.SetRangChars(chars)
	// chars := []string{"你","好","呀","这","是","点","击","验","证","码","哟"}
	// capt.SetRangChars(chars)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestColors(t *testing.T) {
	capt := GetCaptcha()
	capt.SetTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
		"#712217",
		"#885500",
		"#392585",
	})

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestAlpha(t *testing.T) {
	capt := GetCaptcha()

	capt.SetImageFontAlpha(0.5)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestImageFontDistort(t *testing.T) {
	capt := GetCaptcha()

	capt.SetImageFontDistort(DistortLevel2)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestRangAnglePos(t *testing.T) {
	capt := GetCaptcha()

	rang := []RangeVal{
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
	}
	capt.SetTextRangAnglePos(rang)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBackground(t *testing.T) {
	capt := GetCaptcha()

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestThumbBgCircles(t *testing.T) {
	capt := GetCaptcha()

	capt.SetThumbBgCirclesNum(200)

	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(len(b64))
	fmt.Println(len(tb64))
	fmt.Println(key)
	fmt.Println(dots)
}

func TestDemo(t *testing.T) {
	capt := GetCaptcha()

	min, max := 2, 4
	capt.SetFontHinting(font.HintingNone)
	capt.SetThumbBackground(nil)
	capt.SetTextShadow(false)
	capt.SetRangCheckTextLen(RangeVal{Min: min, Max: max})
	capt.SetTextRangLen(RangeVal{Min: max, Max: max + 3})
	capt.SetRangFontSize(RangeVal{Min: 26, Max: 32})
	capt.SetRangCheckFontSize(RangeVal{Min: 20, Max: 28})
	capt.SetTextRangAnglePos([]RangeVal{
		{Min: 15, Max: 30},
		{Min: 45, Max: 60},
		//{75, 90},
		{Min: 105, Max: 115},
		//{125, 130},
		{Min: 145, Max: 160},
		//{175, 190},
		{Min: 205, Max: 215},
		//{225, 230},
		{Min: 245, Max: 260},
		//{275, 290},
		{Min: 305, Max: 315},
		//{325, 330},
		{Min: 345, Max: 360},
	})
	capt.SetImageQuality(QualityCompressNone)
	capt.SetImageFontDistort(DistortLevel4)
	// capt.SetImageFontAlpha(0.8)
	capt.SetThumbFontDistort(DistortNone)
	capt.SetThumbBgDistort(DistortNone)
	capt.SetThumbBgCirclesNum(32)
	capt.SetThumbBgSlimLineNum(2)

	for i := 0; i < 1; i++ {
		dots, b64, tb64, key, err := capt.Generate()
		if err != nil {
			panic(err)
			return
		}
		b64 = b64[strings.Index(b64, ",")+1:]
		ddd, _ := base64.StdEncoding.DecodeString(b64) // 成图片文件并把文件写入到buffer
		_ = ioutil.WriteFile("./b64.png", ddd, 0o666)  // buffer输出到jpg文件中（不做处理，直接写到文件）

		tb64 = tb64[strings.Index(tb64, ",")+1:]
		ddd, _ = base64.StdEncoding.DecodeString(tb64) // 成图片文件并把文件写入到buffer
		_ = ioutil.WriteFile("./bt64.png", ddd, 0o666) // buffer输出到jpg文件中（不做处理，直接写到文件）
		_ = key
		for _, dot := range dots {
			fmt.Print(dot.Text)
		}
		fmt.Println()
	}
}
