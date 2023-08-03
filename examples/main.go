package main

import (
	"fmt"
	. "github.com/qwenode/gwda"
	extOpenCV "github.com/qwenode/gwda-ext-opencv"
	"log"
)

func main() {
	driver, err := NewUSBDriver(nil)
	checkErr(err)

	driverExt, err := extOpenCV.Extend(driver, 0.95)
	checkErr(err, "扩展 driver ，指定匹配阀值为 95%（在不修改或者使用 `OnlyOnceThreshold` 的情况下）")

	pathZero := "/Users/hero/Documents/temp/2020-05/opencv/flag0.png"
	err = driverExt.Tap(pathZero)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")

	err = driverExt.TapOffset(pathZero, 0.1, 0.1)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 10%）")

	pathSeven := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"
	err = driverExt.TapOffset(pathSeven, 0.2, 0.8)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 20%， y 向下👇偏移 80%）")

	err = driverExt.DoubleTap(pathSeven)
	checkErr(err, "找到图片(匹配度 >= 95%)后双击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")

	err = driverExt.DoubleTapOffset(pathSeven, 0.1, 0.25)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 25%）")

	pathSlash := "/Users/hero/Documents/temp/2020-05/opencv/flag.png"
	err = driverExt.SwipeDown(pathSlash, 0.5)
	checkErr(err, "向下👇滑动，滑动距离为图片高度的 50%（默认从图片的正中间顶部向底部滑动，默认滑动距离为 1个 图片高度）")

	err = driverExt.SwipeDownOffset(pathSlash, 0.25, 1)
	checkErr(err, "向下👇滑动（ x 向右👉偏移 25%， y 向下👇偏移 100% ）")

	err = driverExt.SwipeDownOffset(pathSlash, -0.25, -0.8)
	checkErr(err, "向下👇滑动（ x 向左👈偏移 25%， y 向上👆偏移 80% ）")

	// WDADebug = true

	// 撤销 3次 操作
	undo(driverExt, 3)

	err = driverExt.SwipeUp(pathSlash, 0.5)
	checkErr(err, "向上👆滑动，滑动距离为图片高度的 50%（默认从图片的正中间底部向顶部滑动，默认滑动距离为 1个 图片高度）")

	err = driverExt.SwipeUpOffset(pathSlash, 0.9, 0.6)
	checkErr(err, "向上👆滑动（起始滑动点 x 向右👉偏移 90%， y 向下👇偏移 60% ）")

	err = driverExt.OnlyOnceThreshold(0.92).SwipeUpOffset(pathSlash, -0.1, -0.05, 0.3)
	checkErr(err, "向上👆滑动，临时指定匹配阀值为 92%，滑动距离为图片高度的 30%（起始滑动点 x 向左👈偏移 10%， y 向上👆偏移 5% ）")

	// 撤销 3次 操作
	undo(driverExt, 3)

	err = driverExt.SwipeLeft(pathSlash, 0.5)
	checkErr(err, "向左👈滑动，滑动距离为图片宽度的 50%（默认从图片的正中间右侧向左侧滑动，默认滑动距离为 1个 图片宽度）")

	err = driverExt.SwipeLeftOffset(pathSlash, 0.5, 0.55)
	checkErr(err, "向左👈滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 55% ）")

	err = driverExt.OnlyOnceThreshold(0.92).SwipeLeftOffset(pathSlash, -0.15, -0.25)
	checkErr(err, "向左👈滑动，临时指定匹配阀值为 92%（起始滑动点 x 向左👈偏移 15%， y 向上👆偏移 25% ）")

	// driverExt.Debug(extOpenCV.DmNotMatch)

	// 撤销 3次 操作
	undo(driverExt, 3)

	err = driverExt.SwipeRight(pathSlash, 0.5)
	checkErr(err, "向右👉滑动，滑动距离为图片宽度的 50%（默认从图片的正中间左侧向右侧滑动，默认滑动距离为 1个 图片宽度）")

	err = driverExt.SwipeRightOffset(pathSlash, 0.5, 0.6)
	checkErr(err, "向右👉滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 60% ）")

	err = driverExt.OnlyOnceThreshold(0.90).SwipeRightOffset(pathSlash, -0.25, -0.05)
	checkErr(err, "向右👉滑动（起始滑动点 x 向左👈偏移 25%， y 向上👆偏移 5% ）")

	// 撤销 10次 操作
	undo(driverExt, 10)
}

func undo(dExt *extOpenCV.DriverExt, n int) {
	pathUndo := "/Users/hero/Documents/temp/2020-05/opencv/undo.png"
	err := dExt.TapWithNumber(pathUndo, n)
	checkErr(err, fmt.Sprintf("撤销 %d次 操作\n", n))
}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
