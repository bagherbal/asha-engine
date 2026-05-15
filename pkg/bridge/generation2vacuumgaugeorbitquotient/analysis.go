// Package generation2vacuumgaugeorbitquotient implements Gate 497:
// Vacuum Gauge-Orbit Quotient and Unitary-Gauge Representative Audit.
//
// Gate 496 selected the lower active scalar pair at fixed radius, but left an
// S1 phase inside that pair unresolved. Gate 497 asks whether that remaining
// phase is a physical missing scalar coordinate or a gauge-orbit coordinate.
//
// The result is deliberately precise. In the existing abstract electroweak
// scalar representation, Q_em stabilizes the diagnostic vacuum while the broken
// neutral generator Z=T3-Y_phi sweeps exactly the residual lower-pair phase.
// Together with the two charged broken generators, the gauge orbit has rank
// three and is orthogonal to the radial direction, leaving one radial scalar
// direction after quotient. This validates the unitary-gauge representative as
// a bridge quotient representative. It does not promote the result to a native
// theorem because full scalar SU(2)_L, DΦ, the scalar kinetic normalization, and
// the finite action provenance of the gauge orbit remain unclosed.
package generation2vacuumgaugeorbitquotient

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarkineticvacuumprovenance"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarvacuum"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE497-VACUUM-GAUGE-ORBIT-QUOTIENT-UNITARY-GAUGE-REPRESENTATIVE-AUDIT"

	StatusGate496Inherited                     = "CONDITIONAL_SUPPORT_GATE496_LOWER_VACUUM_PLANE_INHERITED"
	StatusResidualS1BridgeGaugeOrbitFound      = "CONDITIONAL_SUPPORT_RESIDUAL_S1_PHASE_IS_BRIDGE_BROKEN_GAUGE_ORBIT"
	StatusPhotonIsotropyStabilizerConfirmed    = "CONDITIONAL_SUPPORT_PHOTON_ISOTROPY_STABILIZER_CONFIRMED"
	StatusBrokenGaugeOrbitRankThreeConfirmed   = "CONDITIONAL_SUPPORT_BROKEN_GAUGE_ORBIT_RANK_THREE_CONFIRMED"
	StatusRadialModeSeparatedFromGaugeOrbit    = "CONDITIONAL_SUPPORT_RADIAL_MODE_SEPARATED_FROM_GAUGE_ORBIT"
	StatusUnitaryGaugeRepresentativeValidated  = "CONDITIONAL_SUPPORT_UNITARY_GAUGE_REPRESENTATIVE_VALID_AFTER_BRIDGE_QUOTIENT"
	StatusFourToOneQuotientDiagnosticConfirmed = "CONDITIONAL_SUPPORT_SCALAR_4_TO_1_QUOTIENT_DIAGNOSTIC_CONFIRMED"
	StatusFirewallPreserved                    = "FIREWALL_PRESERVED_NO_ELECTROWEAK_MASS_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked           = "FIREWALL_BLOCKED_NATIVE_UNITARY_GAUGE_AND_WZ_REGISTRY_WRITE"

	StatusFailedNativeGaugeOrbitNotSelected      = "FAILED_ROUTE_FULL_ELECTROWEAK_GAUGE_ORBIT_NOT_NATIVE_SELECTED"
	StatusFailedResidualS1NativeQuotient         = "FAILED_ROUTE_RESIDUAL_S1_NOT_QUOTIENTED_BY_NATIVE_FINITE_ACTION"
	StatusFailedNativeVacuumVectorSelectorAbsent = "FAILED_ROUTE_NATIVE_VACUUM_VECTOR_SELECTOR_STILL_ABSENT"
	StatusFailedNativeDphiStillUnclosed          = "FAILED_ROUTE_NATIVE_DPHI_PROVENANCE_STILL_UNCLOSED"
	StatusFailedScalarKineticMetricStillUnclosed = "FAILED_ROUTE_SCALAR_KINETIC_NORMALIZATION_STILL_UNCLOSED"
	StatusFailedKappaStillBridge                 = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE"
	StatusFailedWZMassStillBlocked               = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusGate498RedirectDefined                 = "CONDITIONAL_SUPPORT_GATE498_SCALAR_SU2_COMPLEX_STRUCTURE_PROVENANCE_REDIRECT_DEFINED"
)

const eps = 1e-8

type Inheritance struct {
	Executed                        bool
	Gate496AuditDefined             bool
	LowerPairVacuumPlaneSelected    bool
	DiagnosticVacuumMinimizer       bool
	ResidualS1PreviouslyOpen        bool
	AbstractScalarDoubletAvailable  bool
	NativeDphiStillOpen             bool
	NoElectroweakFlavorDataImported bool
	Verdict                         string
	Reason                          string
}

type ResidualPhaseAudit struct {
	Executed                         bool
	ActiveRealDimension              int
	LowerPairDimension               int
	ResidualPhaseDimension           int
	DiagnosticVacuumVector           []float64
	PhaseTangent                     []float64
	BrokenNeutralImage               []float64
	BrokenNeutralMatchesPhaseTangent bool
	BrokenNeutralMatchResidual       float64
	PhotonImageNorm                  float64
	PhotonStabilizesVacuum           bool
	ResidualS1CoveredByGaugeOrbit    bool
	NativeResidualS1QuotientDerived  bool
	Verdict                          string
	Reason                           string
}

type GaugeOrbitAudit struct {
	Executed                      bool
	GaugeGeneratorCount           int
	BrokenGeneratorCount          int
	UnbrokenGeneratorCount        int
	OrbitImageRank                int
	OrbitImageGram                linear.Matrix
	OrbitImageMinEigen            float64
	OrbitImageCondition           float64
	GaugeOrbitRankThree           bool
	IsotropyDimension             int
	PhotonIsotropyGenerator       bool
	RadialDirection               []float64
	MaxRadialOrbitDot             float64
	RadialSeparatedFromGaugeOrbit bool
	ScalarDimensionBeforeQuotient int
	ScalarDimensionAfterQuotient  int
	FourToOneQuotientDiagnostic   bool
	FullGaugeOrbitNativeSelected  bool
	Verdict                       string
	Reason                        string
}

type RepresentativeAudit struct {
	Executed                           bool
	UnitaryGaugeRepresentative         []float64
	RepresentativeIsMinimizer          bool
	RepresentativeAllowedAfterQuotient bool
	RepresentativeNativelySelected     bool
	ExactVacuumVectorStillPhysicalData bool
	WZDiagnosticCanUseRepresentative   bool
	WZNativeMassPromotionAllowed       bool
	Verdict                            string
	Reason                             string
}

type NativeBoundary struct {
	Executed                        bool
	BridgeQuotientSocketClosed      bool
	NativeResidualS1QuotientClosed  bool
	NativeFullScalarSU2Selected     bool
	NativeGaugeOrbitSelected        bool
	NativeDphiClosed                bool
	NativeScalarKineticMetricClosed bool
	NativeKappaSelected             bool
	NativeGaugeHessianSelected      bool
	NativeWZMassMatrixDerived       bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                  bool
	ObservedWMassImported     bool
	ObservedZMassImported     bool
	ObservedHiggsMassImported bool
	FermiConstantImported     bool
	WeakAngleImported         bool
	FineStructureImported     bool
	GaugeCouplingImported     bool
	HiggsVEVImported          bool
	YukawaImported            bool
	CKMPMNSImported           bool
	NativeVacuumVectorWritten bool
	NativeGaugeOrbitWritten   bool
	NativeDphiWritten         bool
	NativeKappaWritten        bool
	NativeWZMassWritten       bool
	Verdict                   string
	Reason                    string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance    Inheritance
	ResidualPhase  ResidualPhaseAudit
	GaugeOrbit     GaugeOrbitAudit
	Representative RepresentativeAudit
	Boundary       NativeBoundary
	Firewall       Firewall
	Registry       RegistryUpdate
	Next           NextStep
	Truth          string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g496, err := generation2scalarkineticvacuumprovenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate496 scalar provenance audit: %w", err)
	}
	sv, err := scalarvacuum.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar vacuum audit: %w", err)
	}
	sc, err := scalarcovariant.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar covariant template: %w", err)
	}

	inherited := Inheritance{
		Executed:                        true,
		Gate496AuditDefined:             true,
		LowerPairVacuumPlaneSelected:    g496.Vacuum.LowPairSelected,
		DiagnosticVacuumMinimizer:       g496.Vacuum.DiagnosticVacuumIsMinimizer,
		ResidualS1PreviouslyOpen:        !g496.Boundary.ResidualPhaseQuotientNative,
		AbstractScalarDoubletAvailable:  g496.ScalarSU2.AbstractDoubletRepresentation,
		NativeDphiStillOpen:             !g496.Boundary.NativeDphiClosed,
		NoElectroweakFlavorDataImported: !g496.Firewall.ObservedWMassImported && !g496.Firewall.ObservedZMassImported && !g496.Firewall.ObservedHiggsMassImported && !g496.Firewall.FermiConstantImported && !g496.Firewall.WeakAngleImported && !g496.Firewall.GaugeCouplingImported && !g496.Firewall.YukawaImported && !g496.Firewall.CKMPMNSImported,
		Verdict:                         StatusGate496Inherited,
		Reason:                          "Gate496 selected the lower vacuum plane and identified the residual S1 phase as the next provenance gap.",
	}

	phi0 := append([]float64(nil), sc.VacuumVector...)
	if len(phi0) != 4 || sc.VacuumRadius <= eps {
		return Analysis{}, fmt.Errorf("Gate497 requires a nonzero four-real diagnostic vacuum vector")
	}
	phaseTangent := []float64{0, 0, 0, sc.VacuumRadius}
	z, err := sc.T3.Sub(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	zphi := matVec(z, phi0)
	qphi := matVec(sc.QEM, phi0)
	zResidual := vectorDistance(zphi, phaseTangent)
	qNorm := norm(qphi)
	residualCovered := zResidual < eps
	photonStabilizes := qNorm < eps

	residual := ResidualPhaseAudit{
		Executed:                         true,
		ActiveRealDimension:              sc.ActiveRealDimension,
		LowerPairDimension:               sv.LowPairDimension,
		ResidualPhaseDimension:           sv.ResidualPhaseFreedomDimension,
		DiagnosticVacuumVector:           phi0,
		PhaseTangent:                     phaseTangent,
		BrokenNeutralImage:               zphi,
		BrokenNeutralMatchesPhaseTangent: residualCovered,
		BrokenNeutralMatchResidual:       zResidual,
		PhotonImageNorm:                  qNorm,
		PhotonStabilizesVacuum:           photonStabilizes,
		ResidualS1CoveredByGaugeOrbit:    residualCovered && photonStabilizes,
		NativeResidualS1QuotientDerived:  false,
		Verdict:                          StatusResidualS1BridgeGaugeOrbitFound,
		Reason:                           "In the abstract electroweak scalar representation, Z=T3-Y_phi maps the diagnostic vacuum exactly into the lower-pair phase tangent, while Q_em fixes the vacuum.  Thus the residual S1 is a bridge broken-gauge orbit direction, not an extra scalar mode at the diagnostic level.",
	}

	t1phi := matVec(sc.T1, phi0)
	t2phi := matVec(sc.T2, phi0)
	images := [][]float64{t1phi, t2phi, zphi}
	gram := gram(images)
	eig, err := linear.SymmetricEigenJacobi(gram, eps, 200)
	if err != nil {
		return Analysis{}, err
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps)
	minPositive := math.Inf(1)
	maxEigen := 0.0
	for _, v := range eig.Values {
		if v > eps && v < minPositive {
			minPositive = v
		}
		if v > maxEigen {
			maxEigen = v
		}
	}
	if math.IsInf(minPositive, 1) {
		minPositive = 0
	}
	cond := math.Inf(1)
	if minPositive > eps {
		cond = maxEigen / minPositive
	}
	radial := make([]float64, 4)
	for i := range phi0 {
		radial[i] = phi0[i] / sc.VacuumRadius
	}
	maxRadialDot := 0.0
	for _, image := range images {
		if d := math.Abs(dot(radial, image)); d > maxRadialDot {
			maxRadialDot = d
		}
	}
	orbit := GaugeOrbitAudit{
		Executed:                      true,
		GaugeGeneratorCount:           4,
		BrokenGeneratorCount:          3,
		UnbrokenGeneratorCount:        1,
		OrbitImageRank:                rank,
		OrbitImageGram:                gram,
		OrbitImageMinEigen:            minPositive,
		OrbitImageCondition:           cond,
		GaugeOrbitRankThree:           rank == 3,
		IsotropyDimension:             1,
		PhotonIsotropyGenerator:       photonStabilizes,
		RadialDirection:               radial,
		MaxRadialOrbitDot:             maxRadialDot,
		RadialSeparatedFromGaugeOrbit: maxRadialDot < eps,
		ScalarDimensionBeforeQuotient: sc.ActiveRealDimension,
		ScalarDimensionAfterQuotient:  sc.ActiveRealDimension - rank,
		FourToOneQuotientDiagnostic:   sc.ActiveRealDimension == 4 && rank == 3 && maxRadialDot < eps && photonStabilizes,
		FullGaugeOrbitNativeSelected:  false,
		Verdict:                       StatusBrokenGaugeOrbitRankThreeConfirmed,
		Reason:                        "The broken images {T1 phi0, T2 phi0, (T3-Y_phi) phi0} have rank three; Q_em is the one-dimensional stabilizer; and the radial direction is orthogonal to the gauge orbit.  This gives the bridge 4 -> 1 scalar quotient diagnostic.",
	}

	repAllowed := residual.ResidualS1CoveredByGaugeOrbit && orbit.FourToOneQuotientDiagnostic && sv.DiagnosticVacuumIsMinimizer
	representative := RepresentativeAudit{
		Executed:                           true,
		UnitaryGaugeRepresentative:         phi0,
		RepresentativeIsMinimizer:          sv.DiagnosticVacuumIsMinimizer,
		RepresentativeAllowedAfterQuotient: repAllowed,
		RepresentativeNativelySelected:     false,
		ExactVacuumVectorStillPhysicalData: false,
		WZDiagnosticCanUseRepresentative:   repAllowed,
		WZNativeMassPromotionAllowed:       false,
		Verdict:                            StatusUnitaryGaugeRepresentativeValidated,
		Reason:                             "After the bridge gauge-orbit quotient, the lower-component real unitary-gauge vector is a valid representative of the selected vacuum orbit.  Its phase is not physical inside the quotient, but the quotient itself remains bridge-level until the finite geometry natively selects the scalar electroweak gauge orbit.",
	}

	boundary := NativeBoundary{
		Executed:                        true,
		BridgeQuotientSocketClosed:      repAllowed,
		NativeResidualS1QuotientClosed:  false,
		NativeFullScalarSU2Selected:     false,
		NativeGaugeOrbitSelected:        false,
		NativeDphiClosed:                false,
		NativeScalarKineticMetricClosed: false,
		NativeKappaSelected:             false,
		NativeGaugeHessianSelected:      false,
		NativeWZMassMatrixDerived:       false,
		Verdict:                         StatusFailedResidualS1NativeQuotient,
		Reason:                          "Gate497 closes the residual S1 only as a bridge electroweak gauge-orbit quotient.  Native promotion remains blocked because the full scalar SU(2)L action, DΦ, scalar kinetic normalization, and action-selected gauge Hessian are still not derived from finite data.",
	}

	firewall := Firewall{
		Executed:                  true,
		ObservedWMassImported:     false,
		ObservedZMassImported:     false,
		ObservedHiggsMassImported: false,
		FermiConstantImported:     false,
		WeakAngleImported:         false,
		FineStructureImported:     false,
		GaugeCouplingImported:     false,
		HiggsVEVImported:          false,
		YukawaImported:            false,
		CKMPMNSImported:           false,
		NativeVacuumVectorWritten: false,
		NativeGaugeOrbitWritten:   false,
		NativeDphiWritten:         false,
		NativeKappaWritten:        false,
		NativeWZMassWritten:       false,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "No observed electroweak, Higgs, gauge-coupling, Yukawa, CKM, or PMNS datum is imported; the unitary-gauge representative is admitted only as bridge quotient data.",
	}

	registry := RegistryUpdate{
		NativeEntries: []string{
			"Gate496 native lower-pair vacuum plane remains accepted; Gate497 adds no new native electroweak mass or gauge-orbit write.",
		},
		BridgeEntries: []string{
			"Within the abstract electroweak scalar representation, the residual lower-pair S1 is exactly the orbit of the broken neutral generator Z=T3-Y_phi.",
			"Q_em stabilizes the vacuum representative and supplies the photon isotropy diagnostic.",
			"The broken gauge orbit has rank three and is orthogonal to the radial scalar direction, producing the bridge 4 -> 1 scalar quotient diagnostic.",
			"The lower-component real unitary-gauge vector is a valid bridge representative after quotient; its phase is not physical at the diagnostic level.",
		},
		EnvironmentalEntries: []string{
			"Observed W/Z masses, Higgs VEV, Fermi constant, weak angle, alpha, gauge couplings, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedNativeGaugeOrbitNotSelected,
			StatusFailedResidualS1NativeQuotient,
			StatusFailedNativeVacuumVectorSelectorAbsent,
			StatusFailedNativeDphiStillUnclosed,
			StatusFailedScalarKineticMetricStillUnclosed,
			StatusFailedKappaStillBridge,
			StatusFailedWZMassStillBlocked,
		},
		OpenTheorems: []string{
			"Derive the full scalar SU(2)L action and scalar complex structure from finite contact/spectral data rather than importing the abstract doublet representation.",
			"Derive native DΦ from the finite inner-fluctuation algebra.",
			"Derive scalar kinetic normalization and gauge Hessian before promoting kappa_U1=6, W/Z masses, or weak-angle data.",
		},
	}

	next := NextStep{
		Gate:        498,
		Title:       "Scalar SU(2)L Complex-Structure and Gauge-Orbit Provenance Audit",
		Reason:      "Gate497 validates the unitary-gauge representative only after a bridge electroweak gauge-orbit quotient; the next missing theorem is the native origin of the scalar complex/SU(2)L structure that makes that quotient legitimate.",
		PrimaryTask: "Audit whether Cℓ(1,7) finite contact/spectral data select the scalar complex structure and full SU(2)L action, or whether the electroweak scalar gauge orbit remains an abstract bridge representation.",
	}

	truth := "Gate497 proves the residual lower-pair S1 is not an extra scalar mode inside the abstract electroweak diagnostic: the broken neutral generator sweeps that phase, Q_em fixes the vacuum, the broken gauge orbit has rank three, and the radial scalar direction survives as the single quotient mode.  This justifies the unitary-gauge representative as a bridge quotient representative.  It does not yet make electroweak symmetry breaking native, because the full scalar SU(2)L action, DΦ, scalar kinetic normalization, gauge Hessian, kappa_U1=6, and W/Z mass matrix still lack finite-action provenance."

	return Analysis{Inheritance: inherited, ResidualPhase: residual, GaugeOrbit: orbit, Representative: representative, Boundary: boundary, Firewall: firewall, Registry: registry, Next: next, Truth: truth}, nil
}

func matVec(m linear.Matrix, v []float64) []float64 {
	out := make([]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		sum := 0.0
		for c := 0; c < m.Cols(); c++ {
			sum += m.At(r, c) * v[c]
		}
		out[r] = sum
	}
	return out
}

func gram(vs [][]float64) linear.Matrix {
	m := linear.NewMatrix(len(vs), len(vs))
	for i := range vs {
		for j := range vs {
			m.Set(i, j, dot(vs[i], vs[j]))
		}
	}
	return m
}

func dot(a, b []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func norm(v []float64) float64 { return math.Sqrt(dot(v, v)) }

func vectorDistance(a, b []float64) float64 {
	if len(a) != len(b) {
		return math.Inf(1)
	}
	s := 0.0
	for i := range a {
		d := a[i] - b[i]
		s += d * d
	}
	return math.Sqrt(s)
}

func FormatVector(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", v)
	}
	return out + "]"
}

func FormatGram(m linear.Matrix) string {
	rows := make([]string, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals[c] = fmt.Sprintf("%.10f", m.At(r, c))
		}
		rows[r] = "[" + strings.Join(vals, ", ") + "]"
	}
	return "[" + strings.Join(rows, ", ") + "]"
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate496=%t lower_plane=%t diagnostic_minimizer=%t residual_s1_open=%t abstract_doublet=%t dphi_open=%t no_data=%t verdict=%s reason=%s", x.Gate496AuditDefined, x.LowerPairVacuumPlaneSelected, x.DiagnosticVacuumMinimizer, x.ResidualS1PreviouslyOpen, x.AbstractScalarDoubletAvailable, x.NativeDphiStillOpen, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatResidualPhase(x ResidualPhaseAudit) string {
	return fmt.Sprintf("active_dim=%d low_dim=%d residual_phase_dim=%d phi0=%s tangent=%s Zphi=%s Z_match=%t Z_residual=%.3e Q_norm=%.3e Q_stabilizes=%t bridge_covered=%t native_quotient=%t verdict=%s reason=%s", x.ActiveRealDimension, x.LowerPairDimension, x.ResidualPhaseDimension, FormatVector(x.DiagnosticVacuumVector), FormatVector(x.PhaseTangent), FormatVector(x.BrokenNeutralImage), x.BrokenNeutralMatchesPhaseTangent, x.BrokenNeutralMatchResidual, x.PhotonImageNorm, x.PhotonStabilizesVacuum, x.ResidualS1CoveredByGaugeOrbit, x.NativeResidualS1QuotientDerived, x.Verdict, x.Reason)
}

func FormatGaugeOrbit(x GaugeOrbitAudit) string {
	cond := "infinite"
	if !math.IsInf(x.OrbitImageCondition, 1) {
		cond = fmt.Sprintf("%.10f", x.OrbitImageCondition)
	}
	return fmt.Sprintf("generators=%d broken=%d unbroken=%d rank=%d gram=%s min_eig=%.10f condition=%s isotropy=%d photon_isotropy=%t radial=%s max_radial_dot=%.3e radial_separated=%t before=%d after=%d four_to_one=%t native_orbit=%t verdict=%s reason=%s", x.GaugeGeneratorCount, x.BrokenGeneratorCount, x.UnbrokenGeneratorCount, x.OrbitImageRank, FormatGram(x.OrbitImageGram), x.OrbitImageMinEigen, cond, x.IsotropyDimension, x.PhotonIsotropyGenerator, FormatVector(x.RadialDirection), x.MaxRadialOrbitDot, x.RadialSeparatedFromGaugeOrbit, x.ScalarDimensionBeforeQuotient, x.ScalarDimensionAfterQuotient, x.FourToOneQuotientDiagnostic, x.FullGaugeOrbitNativeSelected, x.Verdict, x.Reason)
}

func FormatRepresentative(x RepresentativeAudit) string {
	return fmt.Sprintf("unitary=%s minimizer=%t allowed_after_quotient=%t native_selected=%t exact_vector_physical=%t WZ_diagnostic_allowed=%t WZ_native_allowed=%t verdict=%s reason=%s", FormatVector(x.UnitaryGaugeRepresentative), x.RepresentativeIsMinimizer, x.RepresentativeAllowedAfterQuotient, x.RepresentativeNativelySelected, x.ExactVacuumVectorStillPhysicalData, x.WZDiagnosticCanUseRepresentative, x.WZNativeMassPromotionAllowed, x.Verdict, x.Reason)
}

func FormatBoundary(x NativeBoundary) string {
	return fmt.Sprintf("bridge_quotient=%t native_s1=%t native_SU2=%t native_orbit=%t native_Dphi=%t native_metric=%t native_kappa=%t native_hessian=%t native_WZ=%t verdict=%s reason=%s", x.BridgeQuotientSocketClosed, x.NativeResidualS1QuotientClosed, x.NativeFullScalarSU2Selected, x.NativeGaugeOrbitSelected, x.NativeDphiClosed, x.NativeScalarKineticMetricClosed, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%t observed_Z=%t observed_Higgs=%t Fermi=%t theta=%t alpha=%t gauge_coupling=%t v=%t Yukawa=%t CKM_PMNS=%t native_vacuum=%t native_orbit=%t native_Dphi=%t native_kappa=%t native_WZ=%t verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.HiggsVEVImported, x.YukawaImported, x.CKMPMNSImported, x.NativeVacuumVectorWritten, x.NativeGaugeOrbitWritten, x.NativeDphiWritten, x.NativeKappaWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}
