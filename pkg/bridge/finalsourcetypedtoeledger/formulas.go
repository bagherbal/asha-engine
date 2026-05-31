package finalsourcetypedtoeledger

import "math"

func ScaleBridgeVOverMP() float64 {
	S := SBoundarySplit
	return math.Exp(-12*Pi + math.Sqrt(3)/2 + 2*S + 148*S*S)
}

func HiggsQuartic() float64 {
	S := SBoundarySplit
	L := 1 / (8 * Pi)
	return (3.0 / 8.0) * (1 + L) * (1.0/3.0 - S)
}

func TauAction() float64 {
	S := SBoundarySplit
	return 4*Pi/3 + 3.0/10.0 + 7.0/72.0 - S + 0.5*(72+27)*S*S
}

func TopAction() float64 {
	S := SBoundarySplit
	L := 1 / (8 * Pi)
	return L - 5*S + 138*S*S
}

func BottomAction() float64 {
	S := SBoundarySplit
	return 4*Pi/3 - 56*S + 106*S*S
}

func CKMTheta12() float64 {
	S := SBoundarySplit
	return 0.25 - 18*S + 158*S*S
}

func CKMTheta23() float64 {
	S := SBoundarySplit
	L := 1 / (8 * Pi)
	 return L + (5.0/3.0)*S - (8-2*L)*S*S
}

func CKMTheta13() float64 {
	S := SBoundarySplit
	L := 1 / (8 * Pi)
	return 72*L*S - 1.5*S*S
}

func CKMDelta() float64 {
	S := SBoundarySplit
	return Pi/3 + 71*S + (93.0/4.0)*S*S
}

func MajoranaScaleCoefficient() float64 {
	S := SBoundarySplit
	return math.Sqrt(2*Pi) + 49*S + 90*S*S
}

func NeutrinoM2OverM3() float64 {
	S := SBoundarySplit
	L := 1 / (8 * Pi)
	return 4*L + 10*S
}
