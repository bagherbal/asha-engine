package historytransport

import (
	"math"
	"math/cmplx"
)

func buildFlavorTransport(in InputSet, end EndVector, gb GaugeBoundary, yb yukawaBoundaryState) FlavorTransport {
	ckm := ckmMatrix(in.CKM)
	yMZ := end.YukawaSingularValues
	yL := YukawaSingularValues{UpQuarks: yb.Up, DownQuarks: yb.Down, ChargedLeptons: yb.Leptons}
	matMZ := buildYukawaMatrices(yMZ, ckm)
	matL := buildYukawaMatrices(yL, ckm)
	invMZ := buildInvariants(yMZ)
	invL := buildInvariants(yL)
	j := imag(ckm[0][1] * ckm[1][2] * cmplx.Conj(ckm[0][2]) * cmplx.Conj(ckm[1][1]))
	qe := koideQe(in)
	patterns := []string{
		"Yukawa spectra remain sharply hierarchical; hierarchy is observed history data, not ASHA-native derivation.",
		"CKM phase and Jarlskog invariant are imported endpoint data and held fixed in v1 transport.",
		"Neutrino/PMNS sector skipped in v1 by explicit convention.",
		"Full matrix RGE with thresholds is required before any precision flavor-boundary claim.",
	}
	return FlavorTransport{Mu0GeV: in.Mu0GeV, Lambda12GeV: gb.Lambda12GeV, YukawaSingularValuesMZ: yMZ, YukawaSingularValuesLambda12: yL, YukawaMatricesMZ: matMZ, YukawaMatricesLambda12: matL, InvariantsMZ: invMZ, InvariantsLambda12: invL, CKM: complexMatrixJSON(ckm), JCKM: j, KoideQe: qe, Convention: "Y_u=diag(y_u,y_c,y_t); Y_d=V_CKM diag(y_d,y_s,y_b); Y_e=diag(y_e,y_mu,y_tau); PMNS skipped in v1", ResidualPatterns: patterns, Statuses: []string{StatusFlavorTransportComputed, StatusNoNativeDerivationClaim, StatusThresholdsNotHidden}}
}

func buildYukawaMatrices(y YukawaSingularValues, ckm [][]complex128) map[string][][]ComplexValue {
	yu := diagComplex([]float64{y.UpQuarks["u"], y.UpQuarks["c"], y.UpQuarks["t"]})
	ydDiag := []float64{y.DownQuarks["d"], y.DownQuarks["s"], y.DownQuarks["b"]}
	yd := make([][]complex128, 3)
	for i := 0; i < 3; i++ {
		yd[i] = make([]complex128, 3)
		for j := 0; j < 3; j++ {
			yd[i][j] = ckm[i][j] * complex(ydDiag[j], 0)
		}
	}
	ye := diagComplex([]float64{y.ChargedLeptons["e"], y.ChargedLeptons["mu"], y.ChargedLeptons["tau"]})
	return map[string][][]ComplexValue{"Y_u": complexMatrixJSON(yu), "Y_d": complexMatrixJSON(yd), "Y_e": complexMatrixJSON(ye)}
}

func diagComplex(vals []float64) [][]complex128 {
	m := make([][]complex128, len(vals))
	for i := range vals {
		m[i] = make([]complex128, len(vals))
		m[i][i] = complex(vals[i], 0)
	}
	return m
}

func buildInvariants(y YukawaSingularValues) map[string]YukawaInvariants {
	return map[string]YukawaInvariants{
		"Y_u": invariants([]float64{y.UpQuarks["u"], y.UpQuarks["c"], y.UpQuarks["t"]}),
		"Y_d": invariants([]float64{y.DownQuarks["d"], y.DownQuarks["s"], y.DownQuarks["b"]}),
		"Y_e": invariants([]float64{y.ChargedLeptons["e"], y.ChargedLeptons["mu"], y.ChargedLeptons["tau"]}),
	}
}

func invariants(singular []float64) YukawaInvariants {
	spec := make([]float64, len(singular))
	det, tr := 1.0, 0.0
	for i, x := range singular {
		s := x * x
		spec[i] = s
		det *= s
		tr += s
	}
	return YukawaInvariants{SpecYdagY: spec, DetYdagY: det, TraceYdagY: tr}
}

func koideQe(in InputSet) float64 {
	m := map[string]float64{}
	for _, f := range in.Fermions {
		m[f.Name] = f.MassGeV
	}
	sum := m["e"] + m["mu"] + m["tau"]
	root := math.Sqrt(m["e"]) + math.Sqrt(m["mu"]) + math.Sqrt(m["tau"])
	return sum / (root * root)
}

func buildResidual(in InputSet, gb GaugeBoundary, wa WeakAngleTransport, st ScalarTransport, ft FlavorTransport) HistoryResidual {
	var r HistoryResidual
	r.Gauge.Delta3 = gb.Delta3
	r.Gauge.R3 = gb.R3
	r.WeakAngle.DeltaSin2 = wa.DeltaSin2
	r.Scalar.LambdaLambda12 = st.LambdaLambda12
	r.Scalar.ZeroCrossingScaleGeV = st.ZeroCrossingScaleGeV
	r.Scalar.VacuumStabilityStatus = st.VacuumStabilityStatus
	r.Flavor.YukawaHierarchyInvariants = ft.InvariantsLambda12
	r.Flavor.JCKM = ft.JCKM
	r.Flavor.KoideQe = ft.KoideQe
	r.Flavor.ResidualPatterns = ft.ResidualPatterns
	r.Cosmology = in.Cosmology
	r.Statuses = []string{StatusStrongMismatchVisible, StatusWeakAngleTransportResidualVisible, StatusScalarTransportComputed, StatusFlavorTransportComputed, StatusCosmologyEndpointQuarantined, StatusNoNativeDerivationClaim}
	r.Interpretation = "R_hist is the bridge-only fingerprint left after ASHA boundary normalization is transported to the measured endpoint; nonzero entries are history seals, not failed algebra."
	return r
}
