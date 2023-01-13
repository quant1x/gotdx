package parse

import (
	"math"
)

func GetVolume(ivol int) float64 {
	logPoint := ivol >> (8 * 3)
	//hheax := ivol >> (8*3)  // [4]
	hleax := (ivol >> (8 * 2)) & 0xff // [2]
	lheax := (ivol >> 8) & 0xff       // [1]
	lleax := ivol & 0xff              // [0]

	//dbl_1 := 1.0
	//dbl_2 := 2.0
	//dbl_128 := 128.

	dwEcx := logPoint*2 - 0x7f
	dwEdx := logPoint*2 - 0x86
	dwEsi := logPoint*2 - 0x8e
	dwEax := logPoint*2 - 0x96

	tmpEax := 0
	if dwEcx < 0 {
		tmpEax = -dwEcx
	} else {
		tmpEax = dwEcx
	}

	dblXmm6 := math.Pow(2.0, float64(tmpEax))
	if dwEcx < 0 {
		dblXmm6 = 1.0 / dblXmm6
	}

	dblXmm4 := 0.0
	if hleax > 0x80 {
		tmpdblXmm3 := 0.0
		//tmpdblXmm1 := 0.0
		dwtmpeax := dwEdx + 1
		tmpdblXmm3 = math.Pow(2.0, float64(dwtmpeax))
		dblXmm0 := math.Pow(2.0, float64(dwEdx)) * 128.0
		dblXmm0 += float64(hleax&0x7f) * tmpdblXmm3
		dblXmm4 = dblXmm0
	} else {
		dblXmm0 := 0.0
		if dwEdx >= 0 {
			dblXmm0 = math.Pow(2.0, float64(dwEdx)) * float64(hleax)
		} else {
			dblXmm0 = (1 / math.Pow(2.0, float64(dwEdx))) * float64(hleax)
			dblXmm4 = dblXmm0
		}
	}
	dblXmm3 := math.Pow(2.0, float64(dwEsi)) * float64(lheax)
	dblXmm1 := math.Pow(2.0, float64(dwEax)) * float64(lleax)
	if hleax&0x80 != 0 {
		dblXmm3 *= 2.0
		dblXmm1 *= 2.0
	}
	return dblXmm6 + dblXmm4 + dblXmm3 + dblXmm1
}
