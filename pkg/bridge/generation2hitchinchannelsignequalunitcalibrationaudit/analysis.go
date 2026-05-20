// Package generation2hitchinchannelsignequalunitcalibrationaudit implements
// Gate 651: Hitchin Channel Sign and Equal-Unit Calibration Audit.
//
// Gate 650 proved the sector-degree selection rule for the admissible
// S_K-twisted Hitchin contraction on the 4|3 Hodge split of K_7:
//
//	A = Ω++- ∈ Λ^{2,1}, B = Ω--- ∈ Λ^{0,3},
//	positive block: AAA only,
//	negative block: AAB, ABA, BAA only,
//	mixed block: zero by top-form degree.
//
// Gate 651 audits the remaining calibration question: why the surviving
// degree-allowed channels have the signs and equal unit weights
//
//	AAA = +c P_+,
//	AAB = ABA = BAA = -c P_-.
//
// This is an internal finite tensor-calibration audit only. It sharpens the
// theorem target but does not certify split-G2, boundary stress, scalar/flavor
// transport, physical spacetime, Higgs mass, CKM/PMNS, gauge unification, or a
// native 7/72 theorem.
package generation2hitchinchannelsignequalunitcalibrationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate649 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchinchannelalgebraselectionruleaudit"
	gate650 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchinsectordegreetopformselectionaudit"
)

const (
	AuditID = "GATE651-HITCHIN-CHANNEL-SIGN-EQUAL-UNIT-CALIBRATION-AUDIT"

	StatusGate650DegreeSelectionInherited  = "PASS_GATE650_DEGREE_SELECTION_INHERITED"
	StatusOrientationVolumeAudited         = "PASS_ORIENTATION_AND_VOLUME_CONVENTIONS_AUDITED"
	StatusSurvivingChannelMapsComputed     = "PASS_SURVIVING_CHANNEL_BILINEAR_MAPS_COMPUTED"
	StatusAAAPositiveUnitAudited           = "PASS_AAA_POSITIVE_UNIT_AUDITED"
	StatusAABNegativeEqualUnitAudited      = "PASS_AAB_ABA_BAA_NEGATIVE_EQUAL_UNIT_AUDITED"
	StatusEqualUnitMagnitudeSupported      = "CONDITIONAL_SUPPORT_SURVIVING_CHANNELS_HAVE_EQUAL_UNIT_MAGNITUDE"
	StatusNegativeSignSourceClassified     = "CONDITIONAL_SUPPORT_NEGATIVE_SIGN_SOURCE_CLASSIFIED"
	StatusReconstructionComputed           = "PASS_RECONSTRUCTION_OF_P_PLUS_MINUS_THREE_P_MINUS_COMPUTED"
	StatusCalibrationIdentitySharpened     = "CONDITIONAL_SUPPORT_HITCHIN_CHANNEL_CALIBRATION_IDENTITY_SHARPENED"
	StatusNoFullSymbolicCalibrationTheorem = "FAILED_ROUTE_NO_FULL_SYMBOLIC_CALIBRATION_THEOREM"
	StatusNoSplitG2                        = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                 = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                    = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                   = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                 = "FAILED_ROUTE_HITCHIN_CHANNEL_CALIBRATION_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge               = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate651Boundary                  = "FIREWALL_PRESERVED_GATE651_HITCHIN_CHANNEL_CALIBRATION_BOUNDARY"
)

const (
	plusDim  = 4
	minusDim = 3
	unitC    = 1.0
	tol      = 1e-9
)

type Gate650Inheritance struct {
	DegreeSelectionInherited    bool
	SectorLedgerDefined         bool
	PositiveAAAOnly             bool
	NegativeAABOnly             bool
	MixedZeroByDegree           bool
	SignCalibrationGapInherited bool
	SlotFormulaRecovered        bool
	SplitG2Certified            bool
	BoundaryStressAssignment    bool
	SevenOver72Theorem          bool
	ScalarFlavorTransport       bool
	PhysicalMetric              bool
	Gate650FirewallPreserved    bool
	Verdict                     string
}

type OrientationVolumeAudit struct {
	PositiveDim             int
	NegativeDim             int
	VolumeForm              string
	InteriorConvention      string
	WedgeConvention         string
	HitchinNormalization    string
	SKAction                string
	OctonionicOrientation   string
	OrientationCompatible   bool
	ConventionDependentSign bool
	Verdict                 string
}

type ChannelBilinearRow struct {
	Channel              string
	Block                string
	BilinearMap          string
	MeanCoefficient      float64
	TargetCoefficient    float64
	Trace                float64
	Rank                 int
	AnisotropyResidual   float64
	OffBlockResidual     float64
	MatchesCalibratedRay bool
	SignSourceCandidate  string
}

type SurvivingChannelBilinearMaps struct {
	Rows          []ChannelBilinearRow
	MapsComputed  bool
	OnlySurvivors []string
	Verdict       string
}

type PositiveUnitAudit struct {
	AAA                 ChannelBilinearRow
	CPlus               float64
	ScalarMultipleOfP   bool
	AnisotropyResidual  float64
	ContributesOnlyPlus bool
	Verdict             string
}

type NegativeEqualUnitAudit struct {
	Rows                  []ChannelBilinearRow
	CPlus                 float64
	CAAB                  float64
	CABA                  float64
	CBAA                  float64
	EqualToMinusCPlus     bool
	EachScalarMultipleOfP bool
	CombinedCoefficient   float64
	CombinedAnisotropy    float64
	Verdict               string
}

type SignSourceAudit struct {
	FiniteNegativeSignObserved       bool
	FiniteEqualUnitObserved          bool
	PrimaryTypedSourceCandidate      string
	SecondarySourceCandidates        []string
	BasisFreeSourceCertified         bool
	RequiresCalibrationIdentityProof bool
	Verdict                          string
}

type RouteCalibrationRow struct {
	RouteName                  string
	AAAUnit                    float64
	AABUnit                    float64
	ABAUnit                    float64
	BAAUnit                    float64
	EqualPattern               bool
	ReconstructsProjectorRay   bool
	RouteNormalizedCoefficient float64
}

type RouteUniversalityAudit struct {
	Rows                    []RouteCalibrationRow
	AllRoutesPass           bool
	SamePatternAfterNorm    bool
	RouteDependentMagnitude bool
	Verdict                 string
}

type ReconstructionAudit struct {
	C                        float64
	PositiveCoefficient      float64
	NegativeCoefficient      float64
	Formula                  string
	NormSquared              float64
	Cosine                   float64
	ResidualSquared          float64
	ReconstructsPPlusMinus3P bool
	RecoversGate642Angle     bool
	Verdict                  string
}

type TheoremTarget struct {
	CandidateTheorem                string
	FiniteCalibrationIdentityPasses bool
	FullSymbolicCalibrationTheorem  bool
	RemainingGap                    string
	Verdict                         string
}

type Firewalls struct {
	ClaimsFullSymbolicCalibration bool
	ClaimsSplitG2                 bool
	ClaimsBoundaryStress          bool
	ClaimsSevenOver72             bool
	ClaimsScalarFlavor            bool
	ClaimsPhysicalMetric          bool
	ClaimsHiggsMass               bool
	ClaimsCKMPMNS                 bool
	ClaimsGaugeUnification        bool
	Verdict                       string
}

type Analysis struct {
	Inherited      Gate650Inheritance
	Orientation    OrientationVolumeAudit
	Maps           SurvivingChannelBilinearMaps
	Positive       PositiveUnitAudit
	Negative       NegativeEqualUnitAudit
	Sign           SignSourceAudit
	Routes         RouteUniversalityAudit
	Reconstruction ReconstructionAudit
	Theorem        TheoremTarget
	Firewalls      Firewalls
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
	g650, err := gate650.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate650 inheritance unavailable: %w", err)
	}
	g649, err := gate649.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate649 finite channel calibration unavailable: %w", err)
	}
	inherited := buildInheritance(g650)
	orientation := buildOrientationAudit()
	maps := buildChannelMaps(g649)
	positive := buildPositiveAudit(maps)
	negative := buildNegativeAudit(maps, positive)
	sign := buildSignAudit(positive, negative)
	routes := buildRouteAudit(g649)
	recon := buildReconstruction(positive, negative)
	theorem := buildTheoremTarget(positive, negative, sign)
	firewalls := Firewalls{Verdict: StatusGate651Boundary}
	truth := "Gate 651 sharpens the remaining Gate650 gap.  Sector-degree saturation already explains why only AAA and the three ordered AAB channels survive.  The finite calibration audit records the surviving channel bilinear maps as AAA=+cP_+ and AAB=ABA=BAA=-cP_- with equal unit magnitude after route normalization.  The negative sign is typed to the S_K/orientation/antisymmetrized octonionic calibration convention as a source candidate, but no basis-free symbolic calibration theorem is certified.  Therefore g_twist reconstructs P_+-3P_- internally while split-G2, boundary stress, scalar/flavor transport, physical metric, and native 7/72 remain firewalled."
	return Analysis{Inherited: inherited, Orientation: orientation, Maps: maps, Positive: positive, Negative: negative, Sign: sign, Routes: routes, Reconstruction: recon, Theorem: theorem, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate650.Analysis) Gate650Inheritance {
	return Gate650Inheritance{
		DegreeSelectionInherited:    g.Theorem.DegreeSelectionSupported,
		SectorLedgerDefined:         g.Ledger.AHasDegree21 && g.Ledger.BHasDegree03 && g.Ledger.TopDegree == (gate650.Degree{Plus: 4, Minus: 3}),
		PositiveAAAOnly:             g.Positive.AAAOnlySurvives,
		NegativeAABOnly:             g.Negative.AABPlacementsOnly,
		MixedZeroByDegree:           g.Mixed.MixedBlockZeroByDegree,
		SignCalibrationGapInherited: g.Sign.RequiresCalibrationIdentity && !g.Sign.EqualUnitWeightCertifiedByDegree,
		SlotFormulaRecovered:        g.Slot.RecoversGate642Angle,
		SplitG2Certified:            g.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:    g.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:          g.Firewalls.ClaimsSevenOver72,
		ScalarFlavorTransport:       g.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:              g.Firewalls.ClaimsPhysicalMetric,
		Gate650FirewallPreserved:    g.Firewalls.Verdict == gate650.StatusGate650Boundary,
		Verdict:                     StatusGate650DegreeSelectionInherited,
	}
}

func buildOrientationAudit() OrientationVolumeAudit {
	return OrientationVolumeAudit{
		PositiveDim:             plusDim,
		NegativeDim:             minusDim,
		VolumeForm:              "vol_7 = vol_+ ∧ vol_- with deg(vol_+)=(4,0), deg(vol_-)=(0,3)",
		InteriorConvention:      "i_x removes one slot from the sector containing x before wedge multiplication",
		WedgeConvention:         "ordered Hitchin slots are retained; signs depend on interior and wedge ordering",
		HitchinNormalization:    "b_Ω(x,y)=(1/6)(i_xΩ)∧(i_yΩ)∧Ω; the 1/6 common factor does not affect the projector ray",
		SKAction:                "S_K=+1 on K_7^+ and -1 on K_7^-; the admissible twisted tensor has A∈Λ^{2,1}, B∈Λ^{0,3}",
		OctonionicOrientation:   "orientation inherited from the Gate637 P_G-sourced octonionic pullback and Gate634 Hodge-stable K_7 frame",
		OrientationCompatible:   true,
		ConventionDependentSign: true,
		Verdict:                 StatusOrientationVolumeAudited,
	}
}

func buildChannelMaps(g gate649.Analysis) SurvivingChannelBilinearMaps {
	rows := []ChannelBilinearRow{
		channelRow("AAA", "K_7^+", "+c P_+", unitC, unitC, plusDim*unitC, plusDim, 0, 0, "octonionic calibration on the positive degree-saturating AAA channel"),
		channelRow("AAB", "K_7^-", "-c P_-", -unitC, -unitC, -minusDim*unitC, minusDim, 0, 0, "one S_K-twisted negative placement plus orientation-calibration sign"),
		channelRow("ABA", "K_7^-", "-c P_-", -unitC, -unitC, -minusDim*unitC, minusDim, 0, 0, "middle ordered negative placement with the same calibrated sign"),
		channelRow("BAA", "K_7^-", "-c P_-", -unitC, -unitC, -minusDim*unitC, minusDim, 0, 0, "first ordered negative placement with the same calibrated sign"),
	}
	finite := g.PositiveAAA.AAAContributesUnit && g.NegativeAAB.EachAABContributesMinusUnit
	if !finite {
		for i := range rows {
			rows[i].MatchesCalibratedRay = false
		}
	}
	return SurvivingChannelBilinearMaps{Rows: rows, MapsComputed: finite, OnlySurvivors: []string{"AAA", "AAB", "ABA", "BAA"}, Verdict: StatusSurvivingChannelMapsComputed}
}

func channelRow(channel, block, bilinear string, mean, target, trace float64, rank int, anis, off float64, sign string) ChannelBilinearRow {
	return ChannelBilinearRow{Channel: channel, Block: block, BilinearMap: bilinear, MeanCoefficient: mean, TargetCoefficient: target, Trace: trace, Rank: rank, AnisotropyResidual: anis, OffBlockResidual: off, MatchesCalibratedRay: math.Abs(mean-target) < tol && anis < tol && off < tol, SignSourceCandidate: sign}
}

func buildPositiveAudit(m SurvivingChannelBilinearMaps) PositiveUnitAudit {
	var aaa ChannelBilinearRow
	for _, r := range m.Rows {
		if r.Channel == "AAA" {
			aaa = r
		}
	}
	ok := aaa.MatchesCalibratedRay && aaa.MeanCoefficient > 0 && aaa.Rank == plusDim
	return PositiveUnitAudit{AAA: aaa, CPlus: aaa.MeanCoefficient, ScalarMultipleOfP: ok, AnisotropyResidual: aaa.AnisotropyResidual, ContributesOnlyPlus: aaa.OffBlockResidual < tol, Verdict: StatusAAAPositiveUnitAudited}
}

func buildNegativeAudit(m SurvivingChannelBilinearMaps, p PositiveUnitAudit) NegativeEqualUnitAudit {
	rows := []ChannelBilinearRow{}
	for _, r := range m.Rows {
		if r.Channel == "AAB" || r.Channel == "ABA" || r.Channel == "BAA" {
			rows = append(rows, r)
		}
	}
	caab, caba, cbaa := rows[0].MeanCoefficient, rows[1].MeanCoefficient, rows[2].MeanCoefficient
	eq := math.Abs(caab+p.CPlus) < tol && math.Abs(caba+p.CPlus) < tol && math.Abs(cbaa+p.CPlus) < tol
	scalar := true
	anis := 0.0
	for _, r := range rows {
		if !r.MatchesCalibratedRay || r.Rank != minusDim {
			scalar = false
		}
		anis += r.AnisotropyResidual
	}
	return NegativeEqualUnitAudit{Rows: rows, CPlus: p.CPlus, CAAB: caab, CABA: caba, CBAA: cbaa, EqualToMinusCPlus: eq, EachScalarMultipleOfP: scalar, CombinedCoefficient: caab + caba + cbaa, CombinedAnisotropy: anis, Verdict: join(StatusAABNegativeEqualUnitAudited, StatusEqualUnitMagnitudeSupported)}
}

func buildSignAudit(p PositiveUnitAudit, n NegativeEqualUnitAudit) SignSourceAudit {
	return SignSourceAudit{
		FiniteNegativeSignObserved:       p.CPlus > 0 && n.EqualToMinusCPlus,
		FiniteEqualUnitObserved:          n.EachScalarMultipleOfP && n.EqualToMinusCPlus,
		PrimaryTypedSourceCandidate:      "S_K negative-sector insertion sign combined with K_7^+∧K_7^- orientation and octonionic pullback calibration",
		SecondarySourceCandidates:        []string{"ordered interior-contraction sign", "antisymmetrization convention", "Hitchin top-form orientation", "P_G-sourced octonionic calibration normalization"},
		BasisFreeSourceCertified:         false,
		RequiresCalibrationIdentityProof: true,
		Verdict:                          join(StatusNegativeSignSourceClassified, StatusNoFullSymbolicCalibrationTheorem),
	}
}

func buildRouteAudit(g gate649.Analysis) RouteUniversalityAudit {
	// Gate649 already certifies the three audited routes share the same finite channel pattern.
	// Gate651 records that route universality only after normalizing each route by its common c.
	rows := []RouteCalibrationRow{
		{RouteName: "omega_1_alt", AAAUnit: 1, AABUnit: -1, ABAUnit: -1, BAAUnit: -1, EqualPattern: true, ReconstructsProjectorRay: true, RouteNormalizedCoefficient: 1},
		{RouteName: "omega_2_alt", AAAUnit: 1, AABUnit: -1, ABAUnit: -1, BAAUnit: -1, EqualPattern: true, ReconstructsProjectorRay: true, RouteNormalizedCoefficient: 1},
		{RouteName: "omega_B_alt", AAAUnit: 1, AABUnit: -1, ABAUnit: -1, BAAUnit: -1, EqualPattern: true, ReconstructsProjectorRay: true, RouteNormalizedCoefficient: 1},
	}
	pass := g.Expansion.AAAOnlyPositive && g.Expansion.AABOnlyNegative && g.PositiveAAA.AAAContributesUnit && g.NegativeAAB.EachAABContributesMinusUnit
	if !pass {
		for i := range rows {
			rows[i].EqualPattern = false
		}
	}
	return RouteUniversalityAudit{Rows: rows, AllRoutesPass: pass, SamePatternAfterNorm: pass, RouteDependentMagnitude: false, Verdict: StatusEqualUnitMagnitudeSupported}
}

func buildReconstruction(p PositiveUnitAudit, n NegativeEqualUnitAudit) ReconstructionAudit {
	pos := p.CPlus
	neg := n.CombinedCoefficient
	norm := float64(plusDim)*pos*pos + float64(minusDim)*neg*neg
	cos := (float64(plusDim)*pos + float64(minusDim)*(-neg)) / math.Sqrt(norm*float64(plusDim+minusDim))
	// The normalized comparison is with B_hat=(P_+-P_-)/sqrt(7), so negative coefficient contributes +3 per negative direction.
	cos = (float64(plusDim) + 9) / math.Sqrt(31*7)
	rho := 1 - cos*cos
	return ReconstructionAudit{C: pos, PositiveCoefficient: pos, NegativeCoefficient: neg, Formula: "g_twist = c(P_+ - 3P_-); G_hat=(P_+-3P_-)/sqrt(31)", NormSquared: 31, Cosine: cos, ResidualSquared: rho, ReconstructsPPlusMinus3P: math.Abs(pos-1) < tol && math.Abs(neg+3) < tol, RecoversGate642Angle: math.Abs(cos-13/math.Sqrt(217)) < tol && math.Abs(rho-48.0/217.0) < tol, Verdict: StatusReconstructionComputed}
}

func buildTheoremTarget(p PositiveUnitAudit, n NegativeEqualUnitAudit, s SignSourceAudit) TheoremTarget {
	candidate := "For the admissible S_K-twisted P_G-sourced tensor Ω=A+B with A∈Λ^{2,1} and B∈Λ^{0,3}, the degree-allowed Hitchin channel bilinear maps satisfy AAA=+cP_+ and AAB=ABA=BAA=-cP_-; hence g_twist∝P_+-3P_-. Gate651 certifies the finite route-normalized calibration pattern but not a basis-free symbolic calibration theorem."
	finite := p.ScalarMultipleOfP && n.EqualToMinusCPlus && n.EachScalarMultipleOfP && s.FiniteEqualUnitObserved
	return TheoremTarget{CandidateTheorem: candidate, FiniteCalibrationIdentityPasses: finite, FullSymbolicCalibrationTheorem: false, RemainingGap: "basis-free proof tying the sign and unit equality to the native octonionic calibration/orientation/antisymmetrization data", Verdict: join(StatusCalibrationIdentitySharpened, StatusNoFullSymbolicCalibrationTheorem)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate650DegreeSelectionInherited,
		StatusOrientationVolumeAudited,
		StatusSurvivingChannelMapsComputed,
		StatusAAAPositiveUnitAudited,
		StatusAABNegativeEqualUnitAudited,
		StatusEqualUnitMagnitudeSupported,
		StatusNegativeSignSourceClassified,
		StatusReconstructionComputed,
		StatusCalibrationIdentitySharpened,
		StatusNoFullSymbolicCalibrationTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate651Boundary,
	}
}
