package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main() {

	mat := gocv.IMRead("test1.png", gocv.IMReadColor)
	defer mat.Close()

	bilatera := gocv.NewMat()
	defer bilatera.Close()

	gocv.BilateralFilter(mat, &bilatera, 0, 100, 15)
	gocv.MedianBlur(bilatera, &mat, 5)

	gocv.CvtColor(mat, &mat, gocv.ColorBGRToGray)

	threshold := gocv.NewMat()
	defer threshold.Close()

	gocv.Threshold(mat, &threshold, 200, 255, gocv.ThresholdBinaryInv)

	kernel1 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(5, 4))
	kernel2 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(4, 5))
	kernel3 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(6, 6))
	gocv.MorphologyEx(threshold, &threshold, gocv.MorphOpen, kernel1)
	gocv.MorphologyEx(threshold, &threshold, gocv.MorphOpen, kernel2)
	gocv.MorphologyEx(threshold, &threshold, gocv.MorphClose, kernel3)

	gocv.BitwiseNot(threshold, &threshold)

	gocv.BilateralFilter(threshold, &mat, 0, 100, 15)

	ret := gocv.NewMat()
	defer ret.Close()

	gocv.MedianBlur(mat, &ret, 5)
	gocv.IMWrite("/tmp.png", ret)
}
