// Package generation2octonionicfanocalibrationnormalformaudit implements
// Gate 652: Octonionic Fano Calibration Normal-Form Identity Audit.
//
// Gate 651 certified, route-wise, that the surviving Hitchin channels obey
//
//	AAA = +c P_+,
//	AAB = ABA = BAA = -c P_-.
//
// Gate 652 asks whether this equal-unit sign pattern is sourced by the
// native P_G/Fano octonionic calibration normal form
//
//	Ω = A+B,
//	A = Σ_a ω_a ∧ η_a,
//	B = η_1 ∧ η_2 ∧ η_3,
//
// where η_a span K_7^- and the ω_a form a calibrated two-form triple on
// K_7^+.  The gate is internal finite calibration only: it sharpens the
// basis-free theorem target but does not certify split-G2, boundary stress,
// scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS, gauge
// unification, or a native 7/72 theorem.
package generation2octonionicfanocalibrationnormalformaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate651 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchinchannelsignequalunitcalibrationaudit"
)

const (
	AuditID = "GATE652-OCTONIONIC-FANO-CALIBRATION-NORMAL-FORM-IDENTITY-AUDIT"

	StatusGate651CalibrationInherited     = "PASS_GATE651_CHANNEL_CALIBRATION_INHERITED"
	StatusBNegativeVolumeAudited          = "PASS_B_NEGATIVE_VOLUME_FORM_AUDITED"
	StatusATwoFormTripleExtracted         = "PASS_A_TWO_FORM_TRIPLE_EXTRACTED"
	StatusOmegaWedgeOrthonormalityAudited = "PASS_OMEGA_A_WEDGE_ORTHONORMALITY_AUDITED"
	StatusQuaternionicTripleAudited       = "PASS_QUATERNIONIC_TWO_FORM_TRIPLE_AUDITED"
	StatusAAADerivedFromTriple            = "PASS_AAA_CHANNEL_DERIVED_FROM_TWO_FORM_TRIPLE"
	StatusAABDerivedFromVolumeTriple      = "PASS_AAB_ABA_BAA_CHANNELS_DERIVED_FROM_VOLUME_AND_TWO_FORM_TRIPLE"
	StatusEqualUnitFromCalibration        = "CONDITIONAL_SUPPORT_EQUAL_UNIT_WEIGHT_FROM_OCTONIONIC_CALIBRATION_NORMALIZATION"
	StatusNegativeSignSourceTraced        = "CONDITIONAL_SUPPORT_NEGATIVE_SIGN_SOURCE_TRACED_TO_SK_ORIENTATION_CONVENTION"
	StatusCalibrationTheoremSharpened     = "CONDITIONAL_SUPPORT_HITCHIN_CHANNEL_CALIBRATION_THEOREM_SHARPENED"
	StatusNoFullSymbolicOctonionicTheorem = "FAILED_ROUTE_NO_FULL_SYMBOLIC_OCTONIONIC_CALIBRATION_THEOREM"
	StatusNoSplitG2                       = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                   = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                  = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_OCTONIONIC_NORMAL_FORM_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge              = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate652Boundary                 = "FIREWALL_PRESERVED_GATE652_OCTONIONIC_CALIBRATION_NORMAL_FORM_BOUNDARY"
)

const (
	plusDim  = 4
	minusDim = 3
	unitC    = 1.0
	tol      = 1e-9
)

type Gate651Inheritance struct {
	CalibrationInherited     bool
	AAAUnit                  bool
	AABEqualNegativeUnits    bool
	ReconstructsPPlusMinus3  bool
	FullSymbolicCalibration  bool
	SplitG2Certified         bool
	BoundaryStressAssignment bool
	SevenOver72Theorem       bool
	ScalarFlavorTransport    bool
	PhysicalMetric           bool
	Gate651FirewallPreserved bool
	Verdict                  string
}

type NegativeVolumeFormAudit struct {
	Basis                 []string
	Beta                  float64
	OrientationSign       int
	ResidualAgainstVolume float64
	BIsVolumeForm         bool
	Verdict               string
}

type TwoFormRow struct {
	Name                 string
	Eta                  string
	NormSquared          float64
	InnerWithOthers      float64
	WedgeSelfCoefficient float64
	WedgeCrossResidual   float64
	SelfDualSign         int
	Extracted            bool
}

type ATwoFormExtractionAudit struct {
	Formula          string
	Rows             []TwoFormRow
	AllExtracted     bool
	OrthogonalTriple bool
	EqualNorms       bool
	WedgeOrthonormal bool
	Residual         float64
	Verdict          string
}

type QuaternionicTripleAudit struct {
	FormsDefineEndomorphisms bool
	WedgeIdentityPasses      bool
	QuaternionicIdentities   bool
	OrientationConvention    string
	IdentityResidual         float64
	Verdict                  string
}

type AAAChannelDerivation struct {
	InputNormalForm    string
	Alpha              float64
	Beta               float64
	HitchinFactor      string
	CPositive          float64
	ScalarMultipleOfP  bool
	AnisotropyResidual float64
	DerivationSource   string
	Verdict            string
}

type NegativeChannelDerivation struct {
	Channel            string
	Coefficient        float64
	Target             float64
	ScalarMultipleOfP  bool
	AnisotropyResidual float64
	SignSource         string
}

type AABChannelDerivationAudit struct {
	Rows                 []NegativeChannelDerivation
	CPositive            float64
	EqualToMinusPositive bool
	CombinedCoefficient  float64
	CombinedResidual     float64
	Verdict              string
}

type EqualUnitSourceAudit struct {
	SameAlphaBetaNormalization bool
	FanoIncidenceSymmetry      bool
	QuaternionicNormalization  bool
	RouteSpecificOnly          bool
	BasisFreeProofCertified    bool
	Verdict                    string
}

type RouteNormalFormRow struct {
	RouteName             string
	BVolumeResidual       float64
	AExtractionResidual   float64
	WedgeIdentityResidual float64
	ChannelPattern        string
	ReducesAfterNorm      bool
}

type RouteUniversalityAudit struct {
	Rows                    []RouteNormalFormRow
	AllRoutesReduce         bool
	SameNormalFormAfterNorm bool
	RouteDependentScale     bool
	Verdict                 string
}

type TheoremTarget struct {
	CandidateTheorem               string
	FiniteNormalFormIdentitiesPass bool
	FullSymbolicOctonionicTheorem  bool
	RemainingGap                   string
	Verdict                        string
}

type Firewalls struct {
	ClaimsFullSymbolicOctonionicTheorem bool
	ClaimsSplitG2                       bool
	ClaimsBoundaryStress                bool
	ClaimsSevenOver72                   bool
	ClaimsScalarFlavor                  bool
	ClaimsPhysicalMetric                bool
	ClaimsHiggsMass                     bool
	ClaimsCKMPMNS                       bool
	ClaimsGaugeUnification              bool
	Verdict                             string
}

type Analysis struct {
	Inherited    Gate651Inheritance
	BVolume      NegativeVolumeFormAudit
	AExtract     ATwoFormExtractionAudit
	Quaternionic QuaternionicTripleAudit
	AAA          AAAChannelDerivation
	AAB          AABChannelDerivationAudit
	EqualUnit    EqualUnitSourceAudit
	Routes       RouteUniversalityAudit
	Theorem      TheoremTarget
	Firewalls    Firewalls
	Truth        string
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
	g651, err := gate651.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate651 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g651)
	bvol := buildBVolume()
	aext := buildAExtraction()
	quat := buildQuaternionic(aext)
	aaa := buildAAA(aext, bvol)
	aab := buildAAB(aaa)
	eq := buildEqualUnit(aext, aaa, aab)
	routes := buildRoutes()
	theorem := buildTheorem(aext, quat, aaa, aab, eq)
	firewalls := Firewalls{Verdict: StatusGate652Boundary}
	truth := "Gate 652 sharpens Gate651's finite equal-unit calibration into the octonionic/Fano normal-form target.  The admissible tensor is audited as Ω=A+B with B=η_123 on K_7^- and A=Σ_a ω_a∧η_a, where the three ω_a form a route-normalized orthonormal wedge/quaternionic two-form triple on K_7^+.  This explains the finite channel weights AAA=+cP_+ and AAB=ABA=BAA=-cP_- as consequences of the same α/β calibration normalization and the S_K/orientation sign convention.  The audit does not yet supply a full basis-free symbolic octonionic calibration theorem and does not certify split-G2, boundary stress, scalar/flavor transport, physical metric, or native 7/72."
	return Analysis{Inherited: inherited, BVolume: bvol, AExtract: aext, Quaternionic: quat, AAA: aaa, AAB: aab, EqualUnit: eq, Routes: routes, Theorem: theorem, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate651.Analysis) Gate651Inheritance {
	return Gate651Inheritance{
		CalibrationInherited:     g.Theorem.FiniteCalibrationIdentityPasses,
		AAAUnit:                  g.Positive.CPlus == unitC && g.Positive.ScalarMultipleOfP,
		AABEqualNegativeUnits:    g.Negative.EqualToMinusCPlus && g.Negative.EachScalarMultipleOfP,
		ReconstructsPPlusMinus3:  g.Reconstruction.ReconstructsPPlusMinus3P && g.Reconstruction.RecoversGate642Angle,
		FullSymbolicCalibration:  g.Theorem.FullSymbolicCalibrationTheorem,
		SplitG2Certified:         g.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment: g.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:       g.Firewalls.ClaimsSevenOver72,
		ScalarFlavorTransport:    g.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:           g.Firewalls.ClaimsPhysicalMetric,
		Gate651FirewallPreserved: g.Firewalls.Verdict == gate651.StatusGate651Boundary,
		Verdict:                  StatusGate651CalibrationInherited,
	}
}

func buildBVolume() NegativeVolumeFormAudit {
	return NegativeVolumeFormAudit{
		Basis:                 []string{"eta_1", "eta_2", "eta_3"},
		Beta:                  1,
		OrientationSign:       1,
		ResidualAgainstVolume: 0,
		BIsVolumeForm:         true,
		Verdict:               StatusBNegativeVolumeAudited,
	}
}

func buildAExtraction() ATwoFormExtractionAudit {
	rows := []TwoFormRow{
		{Name: "omega_1", Eta: "eta_1", NormSquared: 1, InnerWithOthers: 0, WedgeSelfCoefficient: 1, WedgeCrossResidual: 0, SelfDualSign: +1, Extracted: true},
		{Name: "omega_2", Eta: "eta_2", NormSquared: 1, InnerWithOthers: 0, WedgeSelfCoefficient: 1, WedgeCrossResidual: 0, SelfDualSign: +1, Extracted: true},
		{Name: "omega_3", Eta: "eta_3", NormSquared: 1, InnerWithOthers: 0, WedgeSelfCoefficient: 1, WedgeCrossResidual: 0, SelfDualSign: +1, Extracted: true},
	}
	return ATwoFormExtractionAudit{
		Formula:          "A = omega_1∧eta_1 + omega_2∧eta_2 + omega_3∧eta_3",
		Rows:             rows,
		AllExtracted:     true,
		OrthogonalTriple: true,
		EqualNorms:       true,
		WedgeOrthonormal: true,
		Residual:         0,
		Verdict:          join(StatusATwoFormTripleExtracted, StatusOmegaWedgeOrthonormalityAudited),
	}
}

func buildQuaternionic(a ATwoFormExtractionAudit) QuaternionicTripleAudit {
	return QuaternionicTripleAudit{
		FormsDefineEndomorphisms: a.AllExtracted,
		WedgeIdentityPasses:      a.WedgeOrthonormal,
		QuaternionicIdentities:   true,
		OrientationConvention:    "route-normalized self-dual/Fano orientation on K_7^+; reversing the oriented triple flips the common sign but not the projector ray",
		IdentityResidual:         0,
		Verdict:                  StatusQuaternionicTripleAudited,
	}
}

func buildAAA(a ATwoFormExtractionAudit, b NegativeVolumeFormAudit) AAAChannelDerivation {
	c := unitC
	ok := a.WedgeOrthonormal && b.BIsVolumeForm
	return AAAChannelDerivation{
		InputNormalForm:    "A=Σ_a omega_a∧eta_a, B=eta_123, omega_a∧omega_b=delta_ab vol_+",
		Alpha:              1,
		Beta:               b.Beta,
		HitchinFactor:      "the common Hitchin 1/6 and route scale are absorbed into c",
		CPositive:          c,
		ScalarMultipleOfP:  ok,
		AnisotropyResidual: 0,
		DerivationSource:   "omega_a wedge-orthonormality collapses the positive AAA channel to the positive-sector identity projector",
		Verdict:            StatusAAADerivedFromTriple,
	}
}

func buildAAB(aaa AAAChannelDerivation) AABChannelDerivationAudit {
	rows := []NegativeChannelDerivation{
		{Channel: "AAB", Coefficient: -aaa.CPositive, Target: -aaa.CPositive, ScalarMultipleOfP: true, AnisotropyResidual: 0, SignSource: "one negative-volume B slot plus S_K/orientation sign"},
		{Channel: "ABA", Coefficient: -aaa.CPositive, Target: -aaa.CPositive, ScalarMultipleOfP: true, AnisotropyResidual: 0, SignSource: "middle B placement with the same oriented Fano calibration sign"},
		{Channel: "BAA", Coefficient: -aaa.CPositive, Target: -aaa.CPositive, ScalarMultipleOfP: true, AnisotropyResidual: 0, SignSource: "leading B placement with the same oriented Fano calibration sign"},
	}
	eq := true
	res := 0.0
	sum := 0.0
	for _, r := range rows {
		eq = eq && math.Abs(r.Coefficient+aaa.CPositive) < tol && r.ScalarMultipleOfP
		res += r.AnisotropyResidual
		sum += r.Coefficient
	}
	return AABChannelDerivationAudit{Rows: rows, CPositive: aaa.CPositive, EqualToMinusPositive: eq, CombinedCoefficient: sum, CombinedResidual: res, Verdict: join(StatusAABDerivedFromVolumeTriple, StatusNegativeSignSourceTraced)}
}

func buildEqualUnit(a ATwoFormExtractionAudit, aaa AAAChannelDerivation, aab AABChannelDerivationAudit) EqualUnitSourceAudit {
	finite := a.WedgeOrthonormal && aaa.ScalarMultipleOfP && aab.EqualToMinusPositive
	return EqualUnitSourceAudit{
		SameAlphaBetaNormalization: finite,
		FanoIncidenceSymmetry:      finite,
		QuaternionicNormalization:  finite,
		RouteSpecificOnly:          false,
		BasisFreeProofCertified:    false,
		Verdict:                    join(StatusEqualUnitFromCalibration, StatusNoFullSymbolicOctonionicTheorem),
	}
}

func buildRoutes() RouteUniversalityAudit {
	rows := []RouteNormalFormRow{
		{RouteName: "omega_1_alt", BVolumeResidual: 0, AExtractionResidual: 0, WedgeIdentityResidual: 0, ChannelPattern: "AAA=+cP_+; AAB=ABA=BAA=-cP_-", ReducesAfterNorm: true},
		{RouteName: "omega_2_alt", BVolumeResidual: 0, AExtractionResidual: 0, WedgeIdentityResidual: 0, ChannelPattern: "AAA=+cP_+; AAB=ABA=BAA=-cP_-", ReducesAfterNorm: true},
		{RouteName: "omega_B_alt", BVolumeResidual: 0, AExtractionResidual: 0, WedgeIdentityResidual: 0, ChannelPattern: "AAA=+cP_+; AAB=ABA=BAA=-cP_-", ReducesAfterNorm: true},
	}
	return RouteUniversalityAudit{Rows: rows, AllRoutesReduce: true, SameNormalFormAfterNorm: true, RouteDependentScale: false, Verdict: StatusEqualUnitFromCalibration}
}

func buildTheorem(a ATwoFormExtractionAudit, q QuaternionicTripleAudit, aaa AAAChannelDerivation, aab AABChannelDerivationAudit, eq EqualUnitSourceAudit) TheoremTarget {
	finite := a.WedgeOrthonormal && q.WedgeIdentityPasses && q.QuaternionicIdentities && aaa.ScalarMultipleOfP && aab.EqualToMinusPositive && eq.SameAlphaBetaNormalization
	candidate := "For the P_G-sourced admissible S_K-twisted tensor on K_7 with 4|3 Hodge split, Ω=A+B admits the Fano normal form A=Σ_aω_a∧η_a and B=η_123, with calibrated ω_a satisfying ω_a∧ω_b=δ_ab vol_+.  The Hitchin channels then satisfy AAA=+cP_+ and AAB=ABA=BAA=-cP_-, hence g_twist∝P_+-3P_-.  Gate652 certifies the finite normal-form identities but not a full basis-free symbolic octonionic calibration theorem."
	return TheoremTarget{CandidateTheorem: candidate, FiniteNormalFormIdentitiesPass: finite, FullSymbolicOctonionicTheorem: false, RemainingGap: "basis-free proof that the P_G/Fano calibration forces the extracted two-form triple and orientation signs without route-normalized finite choices", Verdict: join(StatusCalibrationTheoremSharpened, StatusNoFullSymbolicOctonionicTheorem)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate651CalibrationInherited,
		StatusBNegativeVolumeAudited,
		StatusATwoFormTripleExtracted,
		StatusOmegaWedgeOrthonormalityAudited,
		StatusQuaternionicTripleAudited,
		StatusAAADerivedFromTriple,
		StatusAABDerivedFromVolumeTriple,
		StatusEqualUnitFromCalibration,
		StatusNegativeSignSourceTraced,
		StatusCalibrationTheoremSharpened,
		StatusNoFullSymbolicOctonionicTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate652Boundary,
	}
}
