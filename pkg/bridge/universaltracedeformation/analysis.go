// Package universaltracedeformation implements Gate 202: universal trace
// deformation / topological boundary offset audit.
//
// Gate 201 found that two non-universal rational threshold shapes close the
// quarantined u*=1 inverse system only after adding a real universal beta row.
// Gate 202 tests the logically cleaner interpretation: a universal beta row over
// the interval [M_B,M_*] is algebraically equivalent to shifting the topological
// boundary condition itself,
//
//	u_* -> u_* + delta_u,    delta_u = c_univ log(M_*/M_B)/(8*pi^2).
//
// The package then audits whether the already-derived finite B-sector spectral
// gap or the exact contact zeta/action traces provide a canonical, coefficient-
// free boundary-volume defect equal to the required delta_u.  It deliberately
// records failure as a useful obstruction: finite spectral scalars exist, but no
// theorem currently maps them to a universal gauge trace deformation.
package universaltracedeformation

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
)

const tolerance = 1e-9

type Gate201UniversalShape struct {
	Name                string
	RationalShapeDeltaB string
	UniversalBetaDelta  float64
	ThresholdScaleGeV   float64
	BoundaryScaleGeV    float64
	ConditionalAlive    bool
	FiniteDerived       bool
}

type Gate201Snapshot struct {
	ConditionalUniversalShapeMatchesLogged bool
	UniversalCompletionFiniteDerived       bool
	RepresentationPhysicalClaimed          bool
	PhysicalUnificationClaimed             bool
	ObservedInputsUsedForFiniteDerivation  bool
	NoPhysicalPredictionClaim              bool
	ThresholdAnchorsAvailable              bool
	ThresholdCorrectedBetaDerived          bool
	PhysicalMassUnitDerived                bool
	BGapValue                              float64
	ContactPartialModeCount                int
	ScalarActiveModeCount                  int
	StrictNullityAfter                     int
	PhysicalPredictionNullityAfter         int
	UniversalShapes                        []Gate201UniversalShape
	TruthStatement                         string
}

func DefaultGate201Snapshot() Gate201Snapshot {
	return Gate201Snapshot{
		ConditionalUniversalShapeMatchesLogged: true,
		UniversalCompletionFiniteDerived:       false,
		RepresentationPhysicalClaimed:          false,
		PhysicalUnificationClaimed:             false,
		ObservedInputsUsedForFiniteDerivation:  false,
		NoPhysicalPredictionClaim:              true,
		ThresholdAnchorsAvailable:              true,
		ThresholdCorrectedBetaDerived:          false,
		PhysicalMassUnitDerived:                false,
		BGapValue:                              0.1024649212,
		ContactPartialModeCount:                7,
		ScalarActiveModeCount:                  4,
		StrictNullityAfter:                     0,
		PhysicalPredictionNullityAfter:         4,
		UniversalShapes: []Gate201UniversalShape{
			{Name: "Dirac vectorlike quark doublet", RationalShapeDeltaB: "(2/15,2,4/3)", UniversalBetaDelta: 7.65295391, ThresholdScaleGeV: 1.46775e6, BoundaryScaleGeV: 2.40100e15, ConditionalAlive: true, FiniteDerived: false},
			{Name: "Weyl SU(2)L adjoint fermion", RationalShapeDeltaB: "(0,4/3,0)", UniversalBetaDelta: 10.1497543, ThresholdScaleGeV: 8.19808e6, BoundaryScaleGeV: 2.42277e14, ConditionalAlive: true, FiniteDerived: false},
		},
		TruthStatement: "Gate 201 found no raw known rational representation row. Two non-universal shapes close only with a real universal beta completion; the universal row is conditional phenomenology, not finite-derived B-sector physics.",
	}
}

func (s Gate201Snapshot) UniversalShapeCount() int { return len(s.UniversalShapes) }

type BoundaryOffsetEquivalenceAudit struct {
	AlphaInverseEquation                         string
	UEquation                                    string
	RequiredOffsetFormula                        string
	UniversalBetaShiftEquivalentToBoundaryOffset bool
	RelativeRunningUnaffectedByUniversalRow      bool
	BoundaryOffsetActsOnlyAsCommonIntercept      bool
	SignConventionChecked                        bool
	RequiresKnownLeverArm                        bool
	PhysicalPredictionClaim                      bool
	Verdict                                      string
}

type RequiredBoundaryOffset struct {
	ShapeName             string
	RationalShapeDeltaB   string
	UniversalBetaDelta    float64
	ThresholdScaleGeV     float64
	BoundaryScaleGeV      float64
	LeverArmLog           float64
	RequiredDeltaU        float64
	DefectAdjustedU       float64
	RequiredAlphaInvShift float64
	FiniteDerived         bool
	FromGate201           bool
	ConditionalOnly       bool
	Verdict               string
}

type TraceOffsetCandidate struct {
	Name                    string
	Source                  string
	Formula                 string
	ExactValue              string
	Value                   float64
	BranchFree              bool
	UsesObservedInput       bool
	RequiresSpectralTriple  bool
	RequiresCutoffFunction  bool
	RequiresRowSemantics    bool
	RequiresMatchingMap     bool
	CoefficientCanonical    bool
	CanonicalBoundaryOffset bool
	Verdict                 string
}

type AbsorptionTest struct {
	RequiredShape            string
	CandidateName            string
	CandidateSource          string
	RequiredDeltaU           float64
	CandidateDeltaU          float64
	SignedResidual           float64
	AbsResidual              float64
	PerfectlyAbsorbs         bool
	CandidateCanonical       bool
	ConditionalBridgeAllowed bool
	Verdict                  string
}

type FiniteTraceAudit struct {
	BGapValue                          float64
	ContactZetaValues                  int
	ContactActionCandidates            int
	CandidatesAudited                  int
	ExactBranchFreeContactTraces       int
	CanonicalBoundaryOffsetCandidates  int
	PerfectAbsorptions                 int
	CanonicalPerfectAbsorptions        int
	BGapCanonicalOffsetDerived         bool
	ContactZetaCanonicalOffsetDerived  bool
	UniversalVolumeDefectCanonicalized bool
	Verdict                            string
}

type FirewallAudit struct {
	Gate201Inherited                      bool
	Gate201UniversalShapesConditionalOnly bool
	Gate201PhysicalPredictionClaimed      bool
	ObservedInputsUsedForFiniteDerivation bool
	PerfectUOneBoundaryDerived            bool
	DefectAdjustedBoundaryDerived         bool
	BGapUsedAsPhysicalMass                bool
	BGapUsedAsBetaRow                     bool
	ContactZetaUsedAsBetaRow              bool
	ArbitraryCoefficientInserted          bool
	PhysicalUnificationClaimed            bool
	ThresholdCorrectedPhysicalFitClaimed  bool
	AbsoluteMassPredicted                 bool
	FiniteMatchingCorrectionsDerived      bool
	StrictNullityBefore                   int
	StrictNullityAfter                    int
	PhysicalPredictionNullityBefore       int
	PhysicalPredictionNullityAfter        int
	RecommendedNextGate                   string
	OpenRequirements                      []string
	Verdict                               string
}

type Summary struct {
	TestsAudited                  int
	EquivalenceEstablished        bool
	RequiredOffsetsComputed       bool
	FiniteTraceCandidatesAudited  bool
	NoCanonicalOffsetFound        bool
	UniversalCompletionNotDerived bool
	FailedRouteLogged             bool
	NoPhysicalPredictionClaim     bool
	Comment                       string
}

type Analysis struct {
	PreviousGate201 Gate201Snapshot

	Equivalence     BoundaryOffsetEquivalenceAudit
	RequiredOffsets []RequiredBoundaryOffset
	TraceCandidates []TraceOffsetCandidate
	AbsorptionTests []AbsorptionTest
	FiniteTrace     FiniteTraceAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultA, defaultErr = Build(DefaultGate201Snapshot())
	})
	return defaultA, defaultErr
}

func Build(prev Gate201Snapshot) (Analysis, error) {
	if !prev.ConditionalUniversalShapeMatchesLogged || prev.RepresentationPhysicalClaimed || prev.PhysicalUnificationClaimed || len(prev.UniversalShapes) == 0 {
		return Analysis{}, fmt.Errorf("Gate 202 requires Gate 201 conditional universal shape resonances without physical claims")
	}
	if !prev.ThresholdAnchorsAvailable || prev.ThresholdCorrectedBetaDerived || prev.PhysicalMassUnitDerived {
		return Analysis{}, fmt.Errorf("Gate 202 requires dimensionless threshold anchors with mass/beta firewalls sealed")
	}
	if !contactZetaSnapshotFinite() {
		return Analysis{}, fmt.Errorf("Gate 202 requires the exact Gate-162 contact zeta snapshot")
	}

	eq := auditEquivalence()
	required := requiredOffsets(prev)
	candidates := traceCandidates(prev)
	tests := absorptionTests(required, candidates)
	ft := auditFiniteTrace(prev, candidates, tests)
	fw := auditFirewall(prev, ft)
	summary := Summary{
		TestsAudited:                  6,
		EquivalenceEstablished:        eq.UniversalBetaShiftEquivalentToBoundaryOffset && eq.SignConventionChecked,
		RequiredOffsetsComputed:       len(required) == prev.UniversalShapeCount() && len(required) > 0,
		FiniteTraceCandidatesAudited:  ft.CandidatesAudited == len(candidates) && len(candidates) > 0,
		NoCanonicalOffsetFound:        !ft.UniversalVolumeDefectCanonicalized && ft.CanonicalPerfectAbsorptions == 0,
		UniversalCompletionNotDerived: !fw.DefectAdjustedBoundaryDerived && !fw.ArbitraryCoefficientInserted,
		FailedRouteLogged:             !ft.UniversalVolumeDefectCanonicalized && !fw.PhysicalUnificationClaimed,
		NoPhysicalPredictionClaim:     !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed && !fw.AbsoluteMassPredicted && fw.PhysicalPredictionNullityBefore == fw.PhysicalPredictionNullityAfter,
		Comment:                       "Gate 202 proves the universal-beta/common-intercept equivalence, computes the boundary-offset values required by the Gate-201 conditional shapes, and audits B-gap/contact-zeta finite trace candidates. The finite algebra has real trace data, but no canonical coefficient-free map from those traces to the required universal boundary offset; therefore the route is logged as a failed/blocked bridge rather than as a prediction.",
	}
	truth := "Gate 202 establishes a useful algebraic identity: a universal beta row over [M_B,M_*] is equivalent to a common topological boundary offset delta_u. Applying this to the Gate-201 conditional shapes gives two required offsets, but the currently available finite B-sector gap and contact zeta/action traces do not canonically equal or derive those offsets. Because no trace-to-boundary functional, spectral triple coefficient, gauge kinetic map, or matching theorem exists, the universal-completion source remains a FAILED_ROUTE obstruction, not a B-sector prediction."

	return Analysis{
		PreviousGate201: prev,
		Equivalence:     eq,
		RequiredOffsets: required,
		TraceCandidates: candidates,
		AbsorptionTests: tests,
		FiniteTrace:     ft,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  truth,
	}, nil
}

func auditEquivalence() BoundaryOffsetEquivalenceAudit {
	return BoundaryOffsetEquivalenceAudit{
		AlphaInverseEquation:  "A_i(M_Z)=4πu_* + b_i L_*/(2π) + (r_i+c_univ)(L_*-L_B)/(2π)",
		UEquation:             "A_i(M_Z)=4π(u_*+δ_u) + b_i L_*/(2π) + r_i(L_*-L_B)/(2π)",
		RequiredOffsetFormula: "δ_u = c_univ (L_*-L_B)/(8π²)",
		UniversalBetaShiftEquivalentToBoundaryOffset: true,
		RelativeRunningUnaffectedByUniversalRow:      true,
		BoundaryOffsetActsOnlyAsCommonIntercept:      true,
		SignConventionChecked:                        true,
		RequiresKnownLeverArm:                        true,
		PhysicalPredictionClaim:                      false,
		Verdict:                                      "universal beta deformation is exactly a common alpha-inverse intercept shift, hence a boundary-u offset; it cannot repair relative-running data by itself",
	}
}

func requiredOffsets(prev Gate201Snapshot) []RequiredBoundaryOffset {
	matches := append([]Gate201UniversalShape(nil), prev.UniversalShapes...)
	sort.Slice(matches, func(i, j int) bool { return matches[i].UniversalBetaDelta < matches[j].UniversalBetaDelta })
	out := make([]RequiredBoundaryOffset, 0, len(matches))
	for _, m := range matches {
		lever := math.Log(m.BoundaryScaleGeV / m.ThresholdScaleGeV)
		du := m.UniversalBetaDelta * lever / (8 * math.Pi * math.Pi)
		out = append(out, RequiredBoundaryOffset{
			ShapeName:             m.Name,
			RationalShapeDeltaB:   m.RationalShapeDeltaB,
			UniversalBetaDelta:    m.UniversalBetaDelta,
			ThresholdScaleGeV:     m.ThresholdScaleGeV,
			BoundaryScaleGeV:      m.BoundaryScaleGeV,
			LeverArmLog:           lever,
			RequiredDeltaU:        du,
			DefectAdjustedU:       1 + du,
			RequiredAlphaInvShift: 4 * math.Pi * du,
			FiniteDerived:         m.FiniteDerived,
			FromGate201:           true,
			ConditionalOnly:       m.ConditionalAlive,
			Verdict:               "required offset is inherited from Gate-201 phenomenological universal completion, not derived from finite algebra",
		})
	}
	return out
}

func traceCandidates(prev Gate201Snapshot) []TraceOffsetCandidate {
	out := []TraceOffsetCandidate{
		{
			Name:                    "B-sector first spectral gap",
			Source:                  "Gate-201 inherited threshold/B-sector vacuum action spectrum",
			Formula:                 "δ_gap := λ_B,first+",
			ExactValue:              fmt.Sprintf("%.10f", prev.BGapValue),
			Value:                   prev.BGapValue,
			BranchFree:              true,
			UsesObservedInput:       false,
			RequiresMatchingMap:     true,
			CoefficientCanonical:    false,
			CanonicalBoundaryOffset: false,
			Verdict:                 "finite dimensionless spectral anchor; no theorem maps it to a universal boundary-u volume defect",
		},
	}
	for _, z := range contactZetaSnapshot() {
		out = append(out, TraceOffsetCandidate{
			Name:                    "contact zeta trace " + z.Name,
			Source:                  "Gate-162 contactzeta seven-root zeta ledger snapshot",
			Formula:                 z.Formula,
			ExactValue:              z.ExactValue,
			Value:                   rationalFloat(z.ExactValue),
			BranchFree:              true,
			UsesObservedInput:       false,
			RequiresSpectralTriple:  true,
			RequiresCutoffFunction:  true,
			RequiresMatchingMap:     true,
			CoefficientCanonical:    false,
			CanonicalBoundaryOffset: false,
			Verdict:                 "exact branch-free finite zeta value; without spectral-triple/cutoff/gauge-map data it is not a boundary offset",
		})
	}
	for _, c := range contactActionSnapshot() {
		out = append(out, TraceOffsetCandidate{
			Name:                    "contact action scalar " + c.Name,
			Source:                  "Gate-162 contactzeta action functional candidate snapshot",
			Formula:                 c.Formula,
			ExactValue:              c.ExactValue,
			Value:                   rationalFloat(c.ExactValue),
			BranchFree:              true,
			UsesObservedInput:       false,
			RequiresSpectralTriple:  true,
			RequiresCutoffFunction:  true,
			RequiresMatchingMap:     true,
			CoefficientCanonical:    false,
			CanonicalBoundaryOffset: false,
			Verdict:                 "exact action-level scalar; coefficient/cutoff/gauge-map firewalls forbid interpreting it as δ_u",
		})
	}
	return out
}

func absorptionTests(required []RequiredBoundaryOffset, candidates []TraceOffsetCandidate) []AbsorptionTest {
	out := make([]AbsorptionTest, 0, len(required)*len(candidates))
	for _, req := range required {
		for _, cand := range candidates {
			residual := cand.Value - req.RequiredDeltaU
			perfect := math.Abs(residual) <= tolerance
			canonicalBridge := perfect && cand.CanonicalBoundaryOffset && cand.CoefficientCanonical && !cand.RequiresMatchingMap
			verdict := "rejected: candidate does not exactly and canonically absorb the required universal boundary offset"
			if canonicalBridge {
				verdict = "conditional bridge allowed: exact canonical trace offset absorbs the required universal row"
			} else if perfect {
				verdict = "numeric equality without canonical trace-to-boundary theorem; still rejected"
			}
			out = append(out, AbsorptionTest{
				RequiredShape:            req.ShapeName,
				CandidateName:            cand.Name,
				CandidateSource:          cand.Source,
				RequiredDeltaU:           req.RequiredDeltaU,
				CandidateDeltaU:          cand.Value,
				SignedResidual:           residual,
				AbsResidual:              math.Abs(residual),
				PerfectlyAbsorbs:         perfect,
				CandidateCanonical:       cand.CanonicalBoundaryOffset,
				ConditionalBridgeAllowed: canonicalBridge,
				Verdict:                  verdict,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if math.Abs(out[i].AbsResidual-out[j].AbsResidual) > 1e-15 {
			return out[i].AbsResidual < out[j].AbsResidual
		}
		if out[i].RequiredShape != out[j].RequiredShape {
			return out[i].RequiredShape < out[j].RequiredShape
		}
		return out[i].CandidateName < out[j].CandidateName
	})
	return out
}

func auditFiniteTrace(prev Gate201Snapshot, candidates []TraceOffsetCandidate, tests []AbsorptionTest) FiniteTraceAudit {
	exactBranchFree := 0
	canonical := 0
	for _, c := range candidates {
		if c.BranchFree && !c.UsesObservedInput {
			exactBranchFree++
		}
		if c.CanonicalBoundaryOffset {
			canonical++
		}
	}
	perfect := 0
	canonicalPerfect := 0
	for _, t := range tests {
		if t.PerfectlyAbsorbs {
			perfect++
		}
		if t.ConditionalBridgeAllowed {
			canonicalPerfect++
		}
	}
	return FiniteTraceAudit{
		BGapValue:                          prev.BGapValue,
		ContactZetaValues:                  len(contactZetaSnapshot()),
		ContactActionCandidates:            len(contactActionSnapshot()),
		CandidatesAudited:                  len(candidates),
		ExactBranchFreeContactTraces:       exactBranchFree,
		CanonicalBoundaryOffsetCandidates:  canonical,
		PerfectAbsorptions:                 perfect,
		CanonicalPerfectAbsorptions:        canonicalPerfect,
		BGapCanonicalOffsetDerived:         false,
		ContactZetaCanonicalOffsetDerived:  false,
		UniversalVolumeDefectCanonicalized: canonicalPerfect > 0,
		Verdict:                            "B-gap and contact zeta traces are real finite data, but no canonical trace-to-boundary offset functional exists; universal volume defect remains uncanonicalized",
	}
}

func auditFirewall(prev Gate201Snapshot, ft FiniteTraceAudit) FirewallAudit {
	return FirewallAudit{
		Gate201Inherited:                      prev.ConditionalUniversalShapeMatchesLogged && prev.NoPhysicalPredictionClaim,
		Gate201UniversalShapesConditionalOnly: prev.ConditionalUniversalShapeMatchesLogged && !prev.UniversalCompletionFiniteDerived,
		Gate201PhysicalPredictionClaimed:      prev.RepresentationPhysicalClaimed || prev.PhysicalUnificationClaimed,
		ObservedInputsUsedForFiniteDerivation: prev.ObservedInputsUsedForFiniteDerivation,
		PerfectUOneBoundaryDerived:            false,
		DefectAdjustedBoundaryDerived:         ft.UniversalVolumeDefectCanonicalized,
		BGapUsedAsPhysicalMass:                false,
		BGapUsedAsBetaRow:                     false,
		ContactZetaUsedAsBetaRow:              false,
		ArbitraryCoefficientInserted:          false,
		PhysicalUnificationClaimed:            false,
		ThresholdCorrectedPhysicalFitClaimed:  false,
		AbsoluteMassPredicted:                 false,
		FiniteMatchingCorrectionsDerived:      false,
		StrictNullityBefore:                   prev.StrictNullityAfter,
		StrictNullityAfter:                    prev.StrictNullityAfter,
		PhysicalPredictionNullityBefore:       prev.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:        prev.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                   "Gate 203 — universal beta source classification / complete-multiplet versus regulator-trace audit",
		OpenRequirements: []string{
			"derive a canonical finite trace-to-boundary-u functional with fixed coefficient",
			"derive a finite spectral triple/cutoff/gauge-kinetic map that turns zeta/action scalars into coupling intercepts",
			"derive whether the universal row is a complete multiplet tower, regulator/scheme trace, or forbidden fit parameter",
			"derive representation, activation, decoupling, and matching laws before using B-sector/contact data as thresholds",
			"derive or seal M_* independently of the phenomenological inverse-threshold solution",
		},
		Verdict: "Gate 202 blocks the B-gap/contact-zeta boundary-offset route under current axioms; no physical fit or finite prediction is claimed",
	}
}

type contactTraceSnapshot struct {
	Name       string
	Formula    string
	ExactValue string
}

func contactZetaSnapshot() []contactTraceSnapshot {
	return []contactTraceSnapshot{
		{Name: "ζ(0)", Formula: "ζ_contact(0)", ExactValue: "7"},
		{Name: "ζ(1)", Formula: "ζ_contact(1)", ExactValue: "7993/542"},
		{Name: "ζ(2)", Formula: "ζ_contact(2)", ExactValue: "10529233/293764"},
		{Name: "ζ(3)", Formula: "ζ_contact(3)", ExactValue: "15529024549/159220088"},
		{Name: "ζ(4)", Formula: "ζ_contact(4)", ExactValue: "24783201328945/86297287696"},
	}
}

func contactActionSnapshot() []contactTraceSnapshot {
	return []contactTraceSnapshot{
		{Name: "dimension term", Formula: "ζ(0)", ExactValue: "7"},
		{Name: "inverse trace", Formula: "ζ(1)", ExactValue: "7993/542"},
		{Name: "inverse quadratic trace", Formula: "ζ(2)", ExactValue: "10529233/293764"},
		{Name: "inverse cubic trace", Formula: "ζ(3)", ExactValue: "15529024549/159220088"},
		{Name: "inverse quartic trace", Formula: "ζ(4)", ExactValue: "24783201328945/86297287696"},
		{Name: "inverse mean", Formula: "ζ(1)/7", ExactValue: "7993/3794"},
		{Name: "inverse quadratic shape", Formula: "ζ(2)/ζ(1)^2", ExactValue: "10529233/63888049"},
		{Name: "positive-inverse balance", Formula: "Tr(Ω)ζ(1)/49", ExactValue: "231797/199185"},
		{Name: "full determinant", Formula: "prod(λ_i)", ExactValue: "271/29160"},
		{Name: "reciprocal determinant", Formula: "1/prod(λ_i)", ExactValue: "29160/271"},
	}
}

func contactZetaSnapshotFinite() bool {
	zs := contactZetaSnapshot()
	acts := contactActionSnapshot()
	return len(zs) == 5 && len(acts) == 10 && zs[0].ExactValue == "7" && zs[1].ExactValue == "7993/542"
}

func rationalFloat(s string) float64 {
	r := new(big.Rat)
	if _, ok := r.SetString(s); !ok {
		return math.NaN()
	}
	f, _ := new(big.Float).SetPrec(256).SetRat(r).Float64()
	return f
}

func FormatEquivalence(e BoundaryOffsetEquivalenceAudit) string {
	return fmt.Sprintf("alphaEq=%q uEq=%q formula=%q equivalent=%t relativeUnaffected=%t intercept=%t sign=%t leverRequired=%t physicalClaim=%t", e.AlphaInverseEquation, e.UEquation, e.RequiredOffsetFormula, e.UniversalBetaShiftEquivalentToBoundaryOffset, e.RelativeRunningUnaffectedByUniversalRow, e.BoundaryOffsetActsOnlyAsCommonIntercept, e.SignConventionChecked, e.RequiresKnownLeverArm, e.PhysicalPredictionClaim)
}

func FormatRequiredOffset(r RequiredBoundaryOffset) string {
	return fmt.Sprintf("%s shape=%s c_univ=%.9g lever=%.9g δu=%.9g u_defect=%.9g 4πδu=%.9g MB=%.9g M*=%.9g finite=%t conditional=%t", r.ShapeName, r.RationalShapeDeltaB, r.UniversalBetaDelta, r.LeverArmLog, r.RequiredDeltaU, r.DefectAdjustedU, r.RequiredAlphaInvShift, r.ThresholdScaleGeV, r.BoundaryScaleGeV, r.FiniteDerived, r.ConditionalOnly)
}

func FormatRequiredOffsets(rs []RequiredBoundaryOffset) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, FormatRequiredOffset(r))
	}
	return strings.Join(parts, "; ")
}

func FormatTraceCandidate(c TraceOffsetCandidate) string {
	return fmt.Sprintf("%s source=%s formula=%s value=%s≈%.9g branchFree=%t observed=%t rowSemantics=%t triple=%t cutoff=%t matching=%t coeff=%t canonicalOffset=%t", c.Name, c.Source, c.Formula, c.ExactValue, c.Value, c.BranchFree, c.UsesObservedInput, c.RequiresRowSemantics, c.RequiresSpectralTriple, c.RequiresCutoffFunction, c.RequiresMatchingMap, c.CoefficientCanonical, c.CanonicalBoundaryOffset)
}

func FormatTraceCandidates(cs []TraceOffsetCandidate, max int) string {
	if max <= 0 || max > len(cs) {
		max = len(cs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatTraceCandidate(cs[i]))
	}
	if max < len(cs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(cs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAbsorptionTest(t AbsorptionTest) string {
	return fmt.Sprintf("shape=%s candidate=%s source=%s requiredδu=%.9g candidateδu=%.9g residual=%.9g perfect=%t canonical=%t allowed=%t", t.RequiredShape, t.CandidateName, t.CandidateSource, t.RequiredDeltaU, t.CandidateDeltaU, t.SignedResidual, t.PerfectlyAbsorbs, t.CandidateCanonical, t.ConditionalBridgeAllowed)
}

func FormatAbsorptionTests(ts []AbsorptionTest, max int) string {
	if max <= 0 || max > len(ts) {
		max = len(ts)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatAbsorptionTest(ts[i]))
	}
	if max < len(ts) {
		parts = append(parts, fmt.Sprintf("... +%d", len(ts)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatFiniteTrace(f FiniteTraceAudit) string {
	return fmt.Sprintf("Bgap=%.10f zeta=%d action=%d candidates=%d branchFree=%d canonicalCandidates=%d perfect=%d canonicalPerfect=%d BgapOffset=%t zetaOffset=%t volumeCanonical=%t", f.BGapValue, f.ContactZetaValues, f.ContactActionCandidates, f.CandidatesAudited, f.ExactBranchFreeContactTraces, f.CanonicalBoundaryOffsetCandidates, f.PerfectAbsorptions, f.CanonicalPerfectAbsorptions, f.BGapCanonicalOffsetDerived, f.ContactZetaCanonicalOffsetDerived, f.UniversalVolumeDefectCanonicalized)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g201=%t conditionalShapes=%t g201Prediction=%t observedFinite=%t u1Derived=%t defectDerived=%t Bmass=%t Bbeta=%t zetaBeta=%t coeffInserted=%t unification=%t fit=%t mass=%t matching=%t strict=%d->%d prediction=%d->%d next=%s", f.Gate201Inherited, f.Gate201UniversalShapesConditionalOnly, f.Gate201PhysicalPredictionClaimed, f.ObservedInputsUsedForFiniteDerivation, f.PerfectUOneBoundaryDerived, f.DefectAdjustedBoundaryDerived, f.BGapUsedAsPhysicalMass, f.BGapUsedAsBetaRow, f.ContactZetaUsedAsBetaRow, f.ArbitraryCoefficientInserted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.AbsoluteMassPredicted, f.FiniteMatchingCorrectionsDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d equivalence=%t offsets=%t candidates=%t noCanonical=%t universalNotDerived=%t failedRoute=%t noPrediction=%t comment=%s", s.TestsAudited, s.EquivalenceEstablished, s.RequiredOffsetsComputed, s.FiniteTraceCandidatesAudited, s.NoCanonicalOffsetFound, s.UniversalCompletionNotDerived, s.FailedRouteLogged, s.NoPhysicalPredictionClaim, s.Comment)
}
