package historytransport

import (
	"math"
	"math/cmplx"
)

const (
	pi = math.Pi
)

func BuildDefault() (Bundle, error) { return Build(DefaultInputs()) }

func Build(in InputSet) (Bundle, error) {
	in = withTransportedMasses(in)
	end := buildEndVector(in)
	gb := buildGaugeBoundary(in, end)
	wa := WeakAngleTransport{Sin2ThetaBoundary: in.ASHABoundary.Sin2ThetaBoundary, Sin2ThetaEnd: end.Sin2Theta, DeltaSin2: end.Sin2Theta - in.ASHABoundary.Sin2ThetaBoundary, TransportRequired: math.Abs(end.Sin2Theta-in.ASHABoundary.Sin2ThetaBoundary) > 1e-12, Statuses: []string{StatusBoundaryWeakAngleCertified, StatusWeakAngleTransportResidualVisible}}
	scalar, yBoundary := buildScalarAndYukawaTransport(in, end, gb)
	flavor := buildFlavorTransport(in, end, gb, yBoundary)
	residual := buildResidual(in, gb, wa, scalar, flavor)
	statuses := []string{StatusEndVectorBuilt, StatusBoundaryScaleSolved, StatusBoundaryWeakAngleCertified, StatusWeakAngleTransportResidualVisible, StatusStrongMismatchVisible, StatusScalarTransportComputed, StatusFlavorTransportComputed, StatusCosmologyEndpointQuarantined, StatusNoNativeDerivationClaim, StatusNoPhysicalUnificationClaim, StatusThresholdsNotHidden}
	return Bundle{Inputs: in, EndVector: end, GaugeBoundary: gb, WeakAngleTransport: wa, ScalarTransport: scalar, FlavorTransport: flavor, HistoryResidual: residual, Statuses: statuses}, nil
}

func withTransportedMasses(in InputSet) InputSet {
	alphaS := in.Measured["alpha_s_MZ"].Value
	mu0 := in.Mu0GeV
	for i := range in.Fermions {
		f := &in.Fermions[i]
		f.TargetScaleGeV = mu0
		switch f.Name {
		case "u", "d", "s", "c", "b":
			f.MassAtMZGeV = runMassQCDLO(f.MassGeV, f.InputScaleGeV, mu0, alphaS, mu0, 5)
		case "t":
			f.MassAtMZGeV = runMassQCDLO(f.MassGeV, f.InputScaleGeV, mu0, alphaS, mu0, 6)
		default:
			f.MassAtMZGeV = f.MassGeV
		}
	}
	return in
}

func alphaSOneLoop(mu, alphaRef, muRef float64, nf int) float64 {
	beta0 := 11.0 - 2.0*float64(nf)/3.0
	return 1.0 / (1.0/alphaRef + beta0/(2.0*pi)*math.Log(mu/muRef))
}

func runMassQCDLO(m, muFrom, muTo, alphaRef, muRef float64, nf int) float64 {
	if muFrom <= 0 || muTo <= 0 || m == 0 {
		return m
	}
	aFrom := alphaSOneLoop(muFrom, alphaRef, muRef, nf)
	aTo := alphaSOneLoop(muTo, alphaRef, muRef, nf)
	exponent := 12.0 / (33.0 - 2.0*float64(nf))
	return m * math.Pow(aTo/aFrom, exponent)
}

func measured(in InputSet, key string) float64 { return in.Measured[key].Value }

func buildEndVector(in InputSet) EndVector {
	gf := measured(in, "G_F")
	mW := measured(in, "m_W")
	mZ := measured(in, "m_Z")
	mH := measured(in, "m_H")
	alphaS := measured(in, "alpha_s_MZ")
	v := math.Pow(math.Sqrt2*gf, -0.5)
	g2 := 2.0 * mW / v
	gY2 := math.Pow(2.0*mZ/v, 2) - g2*g2
	gY := math.Sqrt(gY2)
	g1 := math.Sqrt(5.0/3.0) * gY
	g3 := math.Sqrt(4.0 * pi * alphaS)
	lambda := mH * mH / (2.0 * v * v)
	sin2 := 1.0 - (mW*mW)/(mZ*mZ)
	y := yukawaSingularValues(in, v, false)
	ckm := ckmMatrix(in.CKM)
	return EndVector{Mu0GeV: in.Mu0GeV, VGeV: v, GY: gY, G1: g1, G2: g2, G3: g3, AlphaS: alphaS, Sin2Theta: sin2, Lambda: lambda, YukawaSingularValues: y, CKM: complexMatrixJSON(ckm), CKMMagnitudes: absMatrix(ckm), CKMConvention: "PDG standard angles: s12,s13,s23,delta; Vub=s13 exp(-i delta)", QuarkMassTransportNote: "quark input masses transported to M_Z by one-loop QCD v1 approximation; no threshold matching or multi-loop precision claim", Statuses: []string{StatusEndVectorBuilt, StatusNoNativeDerivationClaim, StatusThresholdsNotHidden}}
}

func yukawaSingularValues(in InputSet, v float64, boundary bool) YukawaSingularValues {
	up := map[string]float64{}
	down := map[string]float64{}
	lep := map[string]float64{}
	for _, f := range in.Fermions {
		mass := f.MassAtMZGeV
		if boundary {
			mass = f.MassGeV
		}
		y := math.Sqrt2 * mass / v
		switch f.Name {
		case "u", "c", "t":
			up[f.Name] = y
		case "d", "s", "b":
			down[f.Name] = y
		case "e", "mu", "tau":
			lep[f.Name] = y
		}
	}
	return YukawaSingularValues{UpQuarks: up, DownQuarks: down, ChargedLeptons: lep}
}

func buildGaugeBoundary(in InputSet, end EndVector) GaugeBoundary {
	b1, b2, b3 := 41.0/10.0, -19.0/6.0, -7.0
	logL := 8.0 * pi * pi * (1.0/(end.G1*end.G1) - 1.0/(end.G2*end.G2)) / (b1 - b2)
	lambda12 := in.Mu0GeV * math.Exp(logL)
	g1L := runGaugeOneLoop(end.G1, b1, logL)
	g2L := runGaugeOneLoop(end.G2, b2, logL)
	g3L := runGaugeOneLoop(end.G3, b3, logL)
	gStar := 0.5 * (g1L + g2L)
	delta3 := 1.0/(g3L*g3L) - 1.0/(gStar*gStar)
	interp := "threshold_needed"
	if math.Abs(delta3) < 1e-6 {
		interp = "no_threshold_needed"
	}
	return GaugeBoundary{Mu0GeV: in.Mu0GeV, Lambda12GeV: lambda12, LogLambda12Mu0: logL, GStar: gStar, G1Lambda: g1L, G2Lambda: g2L, G3Lambda: g3L, Delta3: delta3, R3: g3L / gStar, Interpretation: interp, RGEConvention: "dg_i/dln(mu)=b_i g_i^3/(16*pi^2); therefore d(g_i^-2)/dln(mu)=-b_i/(8*pi^2)", Statuses: []string{StatusBoundaryScaleSolved, StatusStrongMismatchVisible, StatusNoPhysicalUnificationClaim}}
}

func runGaugeOneLoop(g0, b, logMuMu0 float64) float64 {
	inv := 1.0/(g0*g0) - b/(8.0*pi*pi)*logMuMu0
	return 1.0 / math.Sqrt(inv)
}

func ckmMatrix(in CKMInput) [][]complex128 {
	s12, s13, s23, delta := in.S12, in.S13, in.S23, in.Delta
	c12 := math.Sqrt(1.0 - s12*s12)
	c13 := math.Sqrt(1.0 - s13*s13)
	c23 := math.Sqrt(1.0 - s23*s23)
	eip := cmplx.Exp(complex(0, delta))
	eim := cmplx.Exp(complex(0, -delta))
	return [][]complex128{
		{complex(c12*c13, 0), complex(s12*c13, 0), complex(s13, 0) * eim},
		{complex(-s12*c23, 0) - complex(c12*s23*s13, 0)*eip, complex(c12*c23, 0) - complex(s12*s23*s13, 0)*eip, complex(s23*c13, 0)},
		{complex(s12*s23, 0) - complex(c12*c23*s13, 0)*eip, complex(-c12*s23, 0) - complex(s12*c23*s13, 0)*eip, complex(c23*c13, 0)},
	}
}

func complexMatrixJSON(m [][]complex128) [][]ComplexValue {
	out := make([][]ComplexValue, len(m))
	for i := range m {
		out[i] = make([]ComplexValue, len(m[i]))
		for j, z := range m[i] {
			out[i][j] = ComplexValue{Re: real(z), Im: imag(z), Abs: cmplx.Abs(z)}
		}
	}
	return out
}

func absMatrix(m [][]complex128) [][]float64 {
	out := make([][]float64, len(m))
	for i := range m {
		out[i] = make([]float64, len(m[i]))
		for j, z := range m[i] {
			out[i][j] = cmplx.Abs(z)
		}
	}
	return out
}
