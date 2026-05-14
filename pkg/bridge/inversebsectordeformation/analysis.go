// Package inversebsectordeformation implements Gate 201: inverse B-sector
// deformation search / threshold prediction audit.
//
// Gate 200 showed that the Standard Model one-loop beta vector does not close
// the bottom-up convergence triangle under the quarantined Z-pole comparison.
// Gate 201 inverts that failure.  It asks what single-threshold beta deformation
// would be required to force all three gauge couplings to meet the quarantined
// topological benchmark u*=1, while keeping the finite-algebra and physical
// prediction firewalls sealed.
//
// The central result is deliberately strict: a threshold scale M_B alone does
// not determine a unique deformation vector.  The exact inverse family also
// depends on the still-sealed UV boundary scale M_*.  The package therefore
// exposes the exact family, searches known rational representation rows, audits
// universal-completion degeneracy, and compares the output against the finite
// B-sector/contact spectral inventory without promoting any match to a theorem.
package inversebsectordeformation

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugecouplingboundaryseal"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalboundaryviability"
)

type Rational = gaugecouplingboundaryseal.Rational

func R(num, den int64) Rational { return gaugecouplingboundaryseal.R(num, den) }

const gaugeCount = 3

var gaugeNames = [gaugeCount]string{"U1_GUT", "SU2L", "SU3C"}

type FloatTriple struct {
	U1GUT float64
	SU2L  float64
	SU3C  float64
}

func (v FloatTriple) At(i int) float64 {
	switch i {
	case 0:
		return v.U1GUT
	case 1:
		return v.SU2L
	case 2:
		return v.SU3C
	default:
		panic("bad gauge index")
	}
}

func (v FloatTriple) String() string {
	return fmt.Sprintf("(%.9g,%.9g,%.9g)", v.U1GUT, v.SU2L, v.SU3C)
}

type RationalTriple struct {
	U1GUT Rational
	SU2L  Rational
	SU3C  Rational
}

func (v RationalTriple) At(i int) Rational {
	switch i {
	case 0:
		return v.U1GUT
	case 1:
		return v.SU2L
	case 2:
		return v.SU3C
	default:
		panic("bad gauge index")
	}
}

func (v RationalTriple) Float() FloatTriple {
	return FloatTriple{ratFloat(v.U1GUT), ratFloat(v.SU2L), ratFloat(v.SU3C)}
}

func (v RationalTriple) String() string {
	return fmt.Sprintf("(%s,%s,%s)", v.U1GUT, v.SU2L, v.SU3C)
}

func ratFloat(r Rational) float64 { return float64(r.Num) / float64(r.Den) }

type BoundaryInputAudit struct {
	TopologicalU                          float64
	TopologicalAlphaInverse               float64
	DerivedSin2ThetaWSeed                 float64
	SeedCompatibleWithGUTNormalization    bool
	SMBetaVector                          RationalTriple
	SMBetaVectorAcceptedFromGate200       bool
	EmpiricalLedgerInherited              bool
	EmpiricalLedgerQuarantined            bool
	ObservedInputsUsedForFiniteDerivation bool
	BoundaryScaleDerived                  bool
	Verdict                               string
}

type InverseFamilyAudit struct {
	Equation                                string
	ThresholdLogSymbol                      string
	BoundaryLogSymbol                       string
	RequiredDeltaFormula                    string
	SingleThresholdScaleOnly                bool
	BoundaryScaleStillFree                  bool
	UnderdeterminedByOneContinuousParameter bool
	CanEvaluateIfBoundaryScaleSealed        bool
	MismatchTriangleClosedByConstruction    bool
	UOneBoundaryEnforcedByConstruction      bool
	PhysicalPredictionClaim                 bool
	Verdict                                 string
}

type RequiredDeformationPoint struct {
	ThresholdScaleGeV      float64
	BoundaryScaleGeV       float64
	ThresholdLogFromMZ     float64
	BoundaryLogFromMZ      float64
	LeverArmAboveThreshold float64
	RequiredDeltaB         FloatTriple
	ResidualU1GUT          float64
	ResidualSU2L           float64
	ResidualSU3C           float64
	MaxAbsResidual         float64
	TriangleArea           float64
	ValidOrderedScales     bool
	Verdict                string
}

type KnownRepresentation struct {
	Name                string
	Kind                string
	SMRepresentation    string
	DeltaB              RationalTriple
	StandardBetaRow     bool
	FiniteEngineDerived bool
	Notes               string
}

type RepresentationRawMatch struct {
	Candidate             KnownRepresentation
	LinearSystemSolvable  bool
	ThresholdLogFromMZ    float64
	BoundaryLogFromMZ     float64
	ThresholdScaleGeV     float64
	BoundaryScaleGeV      float64
	MaxAbsResidual        float64
	PositiveOrderedScales bool
	ExactUOneClosure      bool
	Verdict               string
}

type UniversalCompletionMatch struct {
	CandidateShape          KnownRepresentation
	LinearSystemSolvable    bool
	UniversalDelta          float64
	TotalDeltaB             FloatTriple
	ThresholdLogFromMZ      float64
	BoundaryLogFromMZ       float64
	ThresholdScaleGeV       float64
	BoundaryScaleGeV        float64
	MaxAbsResidual          float64
	PositiveOrderedScales   bool
	NonnegativeUniversalRow bool
	ExactUOneClosure        bool
	ConditionalAlive        bool
	FiniteDerived           bool
	Verdict                 string
}

type RepresentationSearchAudit struct {
	KnownRationalRowsAudited            int
	RawExactMatches                     []RepresentationRawMatch
	RawNoGoMatches                      []RepresentationRawMatch
	RawExactKnownRepresentationFound    bool
	UniversalCompletionMatches          []UniversalCompletionMatch
	ConditionalUniversalShapeMatchFound bool
	UniversalCompletionFiniteDerived    bool
	IntegerOrRationalTotalDeltaDerived  bool
	PhysicalRepresentationClaimed       bool
	Verdict                             string
}

type InternalSpectralAudit struct {
	BGapValue                           float64
	BPositiveModeCount                  int
	ContactPartialModeCount             int
	ScalarActiveModeCount               int
	DimensionlessSpectralAnchorsKnown   bool
	BGapHasRepresentationRow            bool
	ContactModesHaveRepresentationRows  bool
	ThresholdActivationRuleDerived      bool
	FiniteToContinuumMatchingDerived    bool
	StructuralBGapMatchFound            bool
	StructuralContactMatchFound         bool
	CountResonancePromoted              bool
	ConditionalAlgebraicPredictionAlive bool
	Verdict                             string
}

type FirewallAudit struct {
	Gate200Inherited                         bool
	ObservedInputsUsedForFiniteDerivation    bool
	TopologicalUOneDerived                   bool
	TopologicalUOneAssumedAsConditionalAudit bool
	BoundaryScaleDerived                     bool
	AbsoluteMassPredicted                    bool
	PhysicalUnificationClaimed               bool
	ThresholdCorrectedPhysicalFitClaimed     bool
	FiniteMatchingCorrectionsDerived         bool
	FiniteToContinuumNormalizationDerived    bool
	BsectorRepresentationDerived             bool
	StrictNullityBefore                      int
	StrictNullityAfter                       int
	PhenomenologyNullityBefore               int
	PhenomenologyNullityAfter                int
	PhysicalPredictionNullityBefore          int
	PhysicalPredictionNullityAfter           int
	RecommendedNextGate                      string
	OpenRequirements                         []string
	Verdict                                  string
}

type Summary struct {
	TestsAudited                           int
	Gate200Inherited                       bool
	InverseFamilyConstructed               bool
	SingleScaleUniquePredictionRejected    bool
	KnownRationalRepresentationNoGoLogged  bool
	ConditionalUniversalShapeMatchesLogged bool
	InternalBsectorMatchFound              bool
	NoPhysicalPredictionClaim              bool
	Comment                                string
}

type Analysis struct {
	PreviousGate200 topologicalboundaryviability.Analysis
	Threshold       threshold.Analysis
	Boundary        BoundaryInputAudit
	InverseFamily   InverseFamilyAudit
	BenchmarkPoint  RequiredDeformationPoint
	Representation  RepresentationSearchAudit
	Internal        InternalSpectralAudit
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
		prev, err := topologicalboundaryviability.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 200 input: %w", err)
			return
		}
		th, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build finite threshold inventory: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, th)
	})
	return defaultA, defaultErr
}

func Build(prev topologicalboundaryviability.Analysis, th threshold.Analysis) (Analysis, error) {
	if !prev.Summary.PairwiseIntersectionsSolved || !prev.Summary.MismatchTriangleNonzero || !prev.Summary.EmpiricalLedgerQuarantined {
		return Analysis{}, fmt.Errorf("Gate 201 requires Gate 200 quarantined mismatch triangle")
	}
	if prev.Firewall.ObservedInputsUsedForFiniteDerivation || prev.Firewall.PhysicalUnificationClaimed || prev.Firewall.ThresholdCorrectedPhysicalFitClaimed {
		return Analysis{}, fmt.Errorf("Gate 201 refuses leaked finite derivation, unification claim, or threshold fit from Gate 200")
	}
	if !th.DimensionlessSpectralAnchorsAvailable || th.ThresholdCorrectedBetaDerived || th.PhysicalMassUnitDerived {
		return Analysis{}, fmt.Errorf("Gate 201 requires finite spectral anchors with threshold beta/mass firewalls still sealed")
	}

	boundary := auditBoundaryInputs(prev)
	family := auditInverseFamily()
	// A single numerical point is included only as a reproducible check of the
	// formula: M_B is set to the inherited Z-pole scale and M_* to the Gate-200
	// centroid.  This is not a preferred physical solution and is not used to
	// derive masses or representation rows.
	benchmark := RequiredDeformationAt(prev.Ledger.ScaleGeV, prev.Benchmark.CentroidScaleGeV, prev)
	rep := auditRepresentations(prev)
	internal := auditInternalSpectral(th, rep)
	fw := auditFirewall(prev, boundary, family, rep, internal)
	summary := Summary{
		TestsAudited:                           7,
		Gate200Inherited:                       fw.Gate200Inherited,
		InverseFamilyConstructed:               family.CanEvaluateIfBoundaryScaleSealed && family.MismatchTriangleClosedByConstruction && family.UOneBoundaryEnforcedByConstruction,
		SingleScaleUniquePredictionRejected:    family.SingleThresholdScaleOnly && family.BoundaryScaleStillFree && family.UnderdeterminedByOneContinuousParameter,
		KnownRationalRepresentationNoGoLogged:  !rep.RawExactKnownRepresentationFound && len(rep.RawNoGoMatches) == rep.KnownRationalRowsAudited,
		ConditionalUniversalShapeMatchesLogged: rep.ConditionalUniversalShapeMatchFound && !rep.UniversalCompletionFiniteDerived,
		InternalBsectorMatchFound:              internal.StructuralBGapMatchFound || internal.StructuralContactMatchFound,
		NoPhysicalPredictionClaim:              !fw.PhysicalUnificationClaimed && !fw.AbsoluteMassPredicted && !fw.ThresholdCorrectedPhysicalFitClaimed && fw.PhysicalPredictionNullityBefore == fw.PhysicalPredictionNullityAfter,
		Comment:                                "Gate 201 inverts the Gate 200 mismatch into an exact one-threshold beta-deformation family. A single threshold onset M_B is insufficient to pick a unique Δb vector because the UV boundary scale M_* remains sealed. Known rational one-row representation candidates do not close the u*=1 system as raw rows. Some non-universal rational shapes admit an exact closure only after adding a large universal beta row; this is recorded as conditional phenomenological predata, not as a finite B-sector theorem.",
	}
	truth := "Gate 201 converts the mismatch triangle into a precise inverse-threshold equation. The equation is useful because it tells the engine what a new continuum sector would have to contribute, but it also exposes a new no-go: M_B alone does not determine Δb. The current known rational representation library contains no raw single-row solution that heals the triangle at u*=1 with ordered positive scales. Universal-completion shape resonances exist, but the universal component is not finite-derived and the B-sector/contact spectral anchors still lack representation rows, activation rules, and matching maps. Therefore the B-sector deformation remains a conditional algebraic prediction target, not a derived physical particle threshold."

	return Analysis{
		PreviousGate200: prev,
		Threshold:       th,
		Boundary:        boundary,
		InverseFamily:   family,
		BenchmarkPoint:  benchmark,
		Representation:  rep,
		Internal:        internal,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  truth,
	}, nil
}

func auditBoundaryInputs(prev topologicalboundaryviability.Analysis) BoundaryInputAudit {
	seed := 3.0 / 8.0
	return BoundaryInputAudit{
		TopologicalU:                          1,
		TopologicalAlphaInverse:               4 * math.Pi,
		DerivedSin2ThetaWSeed:                 seed,
		SeedCompatibleWithGUTNormalization:    math.Abs(seed-0.375) < 1e-15,
		SMBetaVector:                          RationalTriple{R(41, 10), R(-19, 6), R(-7, 1)},
		SMBetaVectorAcceptedFromGate200:       prev.Triangle.PairwiseCount == 3,
		EmpiricalLedgerInherited:              prev.Ledger.Quarantined && prev.Ledger.ExplicitPhenomenologicalInput,
		EmpiricalLedgerQuarantined:            prev.Ledger.Quarantined,
		ObservedInputsUsedForFiniteDerivation: prev.Ledger.UsedForFiniteDerivation,
		BoundaryScaleDerived:                  false,
		Verdict:                               "u*=1, sin²θW=3/8, b_SM, and the Z-pole ledger are admitted only as a quarantined inverse-audit input set",
	}
}

func auditInverseFamily() InverseFamilyAudit {
	return InverseFamilyAudit{
		Equation:                                "A_i(M_Z) - b_i L_*/(2π) - Δb_i (L_*-L_B)/(2π) = 4π u_* with u_*=1",
		ThresholdLogSymbol:                      "L_B = log(M_B/M_Z)",
		BoundaryLogSymbol:                       "L_* = log(M_*/M_Z)",
		RequiredDeltaFormula:                    "Δb_i(L_*,L_B) = [2π(A_i(M_Z)-4π) - b_i L_*] / (L_*-L_B)",
		SingleThresholdScaleOnly:                true,
		BoundaryScaleStillFree:                  true,
		UnderdeterminedByOneContinuousParameter: true,
		CanEvaluateIfBoundaryScaleSealed:        true,
		MismatchTriangleClosedByConstruction:    true,
		UOneBoundaryEnforcedByConstruction:      true,
		PhysicalPredictionClaim:                 false,
		Verdict:                                 "the inverse system is exact, but M_B alone is insufficient; a boundary-scale seal or a representation-row theorem is still required",
	}
}

func RequiredDeformationAt(thresholdScaleGeV, boundaryScaleGeV float64, prev topologicalboundaryviability.Analysis) RequiredDeformationPoint {
	lb := math.Log(thresholdScaleGeV / prev.Ledger.ScaleGeV)
	ls := math.Log(boundaryScaleGeV / prev.Ledger.ScaleGeV)
	lever := ls - lb
	beta := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	obs := observedTriple(prev)
	topAlphaInv := 4 * math.Pi
	var d [gaugeCount]float64
	var res [gaugeCount]float64
	if lever > 0 && finite(lever) {
		for i := 0; i < gaugeCount; i++ {
			d[i] = (2*math.Pi*(obs.At(i)-topAlphaInv) - beta.At(i)*ls) / lever
			res[i] = obs.At(i) - beta.At(i)*ls/(2*math.Pi) - d[i]*lever/(2*math.Pi) - topAlphaInv
		}
	} else {
		for i := 0; i < gaugeCount; i++ {
			d[i] = math.NaN()
			res[i] = math.NaN()
		}
	}
	maxRes := maxAbs(res[:])
	return RequiredDeformationPoint{
		ThresholdScaleGeV:      thresholdScaleGeV,
		BoundaryScaleGeV:       boundaryScaleGeV,
		ThresholdLogFromMZ:     lb,
		BoundaryLogFromMZ:      ls,
		LeverArmAboveThreshold: lever,
		RequiredDeltaB:         FloatTriple{d[0], d[1], d[2]},
		ResidualU1GUT:          res[0],
		ResidualSU2L:           res[1],
		ResidualSU3C:           res[2],
		MaxAbsResidual:         maxRes,
		TriangleArea:           0,
		ValidOrderedScales:     thresholdScaleGeV >= prev.Ledger.ScaleGeV && boundaryScaleGeV > thresholdScaleGeV && lever > 0 && maxRes < 1e-8,
		Verdict:                "formula check only; this point uses Gate-200 comparison scales and is not a physical mass prediction",
	}
}

func auditRepresentations(prev topologicalboundaryviability.Analysis) RepresentationSearchAudit {
	library := knownRepresentationLibrary()
	raw := make([]RepresentationRawMatch, 0, len(library))
	univ := make([]UniversalCompletionMatch, 0, len(library))
	rawExact := make([]RepresentationRawMatch, 0)
	rawNoGo := make([]RepresentationRawMatch, 0)
	cond := make([]UniversalCompletionMatch, 0)
	for _, c := range library {
		r := solveRawRepresentation(prev, c)
		raw = append(raw, r)
		if r.ExactUOneClosure {
			rawExact = append(rawExact, r)
		} else {
			rawNoGo = append(rawNoGo, r)
		}
		u := solveUniversalCompletion(prev, c)
		univ = append(univ, u)
		if u.ConditionalAlive {
			cond = append(cond, u)
		}
	}
	sort.Slice(cond, func(i, j int) bool {
		if math.Abs(cond[i].UniversalDelta-cond[j].UniversalDelta) == 0 {
			return cond[i].CandidateShape.Name < cond[j].CandidateShape.Name
		}
		return cond[i].UniversalDelta < cond[j].UniversalDelta
	})
	verdict := "no raw known rational representation row closes the u*=1 inverse system; universal-completion shape resonances are conditional and not finite-derived"
	return RepresentationSearchAudit{
		KnownRationalRowsAudited:            len(library),
		RawExactMatches:                     rawExact,
		RawNoGoMatches:                      rawNoGo,
		RawExactKnownRepresentationFound:    len(rawExact) > 0,
		UniversalCompletionMatches:          cond,
		ConditionalUniversalShapeMatchFound: len(cond) > 0,
		UniversalCompletionFiniteDerived:    false,
		IntegerOrRationalTotalDeltaDerived:  false,
		PhysicalRepresentationClaimed:       false,
		Verdict:                             verdict,
	}
}

func knownRepresentationLibrary() []KnownRepresentation {
	return []KnownRepresentation{
		{Name: "real scalar gauge singlet", Kind: "real scalar", SMRepresentation: "(1,1,0)", DeltaB: RationalTriple{R(0, 1), R(0, 1), R(0, 1)}, StandardBetaRow: true, Notes: "zero beta row; cannot alter the triangle"},
		{Name: "complex scalar singlet Y=1", Kind: "complex scalar", SMRepresentation: "(1,1,1)", DeltaB: RationalTriple{R(1, 5), R(0, 1), R(0, 1)}, StandardBetaRow: true, Notes: "hypercharge-only scalar row"},
		{Name: "Higgs-like scalar doublet", Kind: "complex scalar", SMRepresentation: "(1,2,1/2)", DeltaB: RationalTriple{R(1, 10), R(1, 6), R(0, 1)}, StandardBetaRow: true, Notes: "one extra complex scalar doublet"},
		{Name: "real SU(2)L triplet scalar", Kind: "real scalar", SMRepresentation: "(1,3,0)", DeltaB: RationalTriple{R(0, 1), R(1, 3), R(0, 1)}, StandardBetaRow: true, Notes: "adjoint real scalar row"},
		{Name: "complex SU(2)L triplet scalar", Kind: "complex scalar", SMRepresentation: "(1,3,0)", DeltaB: RationalTriple{R(0, 1), R(2, 3), R(0, 1)}, StandardBetaRow: true, Notes: "complex adjoint scalar row"},
		{Name: "Weyl SU(2)L adjoint fermion", Kind: "Weyl fermion", SMRepresentation: "(1,3,0)", DeltaB: RationalTriple{R(0, 1), R(4, 3), R(0, 1)}, StandardBetaRow: true, Notes: "wino-like adjoint shape"},
		{Name: "Weyl SU(3)c adjoint fermion", Kind: "Weyl fermion", SMRepresentation: "(8,1,0)", DeltaB: RationalTriple{R(0, 1), R(0, 1), R(2, 1)}, StandardBetaRow: true, Notes: "gluino-like adjoint shape"},
		{Name: "Dirac vectorlike charged lepton", Kind: "Dirac fermion", SMRepresentation: "(1,1,-1)", DeltaB: RationalTriple{R(4, 5), R(0, 1), R(0, 1)}, StandardBetaRow: true, Notes: "vectorlike singlet lepton"},
		{Name: "Dirac vectorlike lepton doublet", Kind: "Dirac fermion", SMRepresentation: "(1,2,-1/2)", DeltaB: RationalTriple{R(2, 5), R(2, 3), R(0, 1)}, StandardBetaRow: true, Notes: "vectorlike weak doublet"},
		{Name: "Dirac vectorlike up quark", Kind: "Dirac fermion", SMRepresentation: "(3,1,2/3)", DeltaB: RationalTriple{R(16, 15), R(0, 1), R(2, 3)}, StandardBetaRow: true, Notes: "vectorlike color triplet singlet"},
		{Name: "Dirac vectorlike down quark", Kind: "Dirac fermion", SMRepresentation: "(3,1,-1/3)", DeltaB: RationalTriple{R(4, 15), R(0, 1), R(2, 3)}, StandardBetaRow: true, Notes: "vectorlike color triplet singlet"},
		{Name: "Dirac vectorlike quark doublet", Kind: "Dirac fermion", SMRepresentation: "(3,2,1/6)", DeltaB: RationalTriple{R(2, 15), R(2, 1), R(4, 3)}, StandardBetaRow: true, Notes: "vectorlike color/weak doublet shape"},
		{Name: "one chiral SM generation aggregate", Kind: "Weyl fermion aggregate", SMRepresentation: "Q+u+d+L+e", DeltaB: RationalTriple{R(4, 3), R(4, 3), R(4, 3)}, StandardBetaRow: true, Notes: "universal one-generation aggregate; invisible to triangle closure"},
	}
}

func solveRawRepresentation(prev topologicalboundaryviability.Analysis, c KnownRepresentation) RepresentationRawMatch {
	obs := observedTriple(prev)
	beta := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	d := c.DeltaB.Float()
	s := targetSTriple(obs)
	// Solve the gauge-1/gauge-2 absolute equations:
	// S_i = (b_i + d_i)L_* - d_i L_B.
	a11, a12 := beta.At(0)+d.At(0), -d.At(0)
	a21, a22 := beta.At(1)+d.At(1), -d.At(1)
	det := a11*a22 - a12*a21
	out := RepresentationRawMatch{Candidate: c}
	if math.Abs(det) < 1e-12 || !finite(det) {
		out.Verdict = "singular raw inverse system for this row; no unique (M_B,M_*) solution"
		return out
	}
	x := (s.At(0)*a22 - a12*s.At(1)) / det
	y := (a11*s.At(1) - s.At(0)*a21) / det
	res := rawResiduals(s, beta, d, x, y)
	maxRes := maxAbs(res[:])
	ordered := x > y && y > 0 && finite(x) && finite(y)
	out.LinearSystemSolvable = true
	out.BoundaryLogFromMZ = x
	out.ThresholdLogFromMZ = y
	out.BoundaryScaleGeV = scaleFromLog(prev, x)
	out.ThresholdScaleGeV = scaleFromLog(prev, y)
	out.MaxAbsResidual = maxRes
	out.PositiveOrderedScales = ordered
	out.ExactUOneClosure = ordered && maxRes < 1e-7
	if out.ExactUOneClosure {
		out.Verdict = "raw known rational representation row closes the inverse system; this would still be phenomenology, not finite derivation"
	} else {
		out.Verdict = "raw row rejected: it either gives non-ordered scales or fails the third u*=1 equation"
	}
	return out
}

func solveUniversalCompletion(prev topologicalboundaryviability.Analysis, c KnownRepresentation) UniversalCompletionMatch {
	obs := observedTriple(prev)
	beta := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	r := c.DeltaB.Float()
	s := targetSTriple(obs)
	// Let d_i = r_i + c_univ and z=L_*-L_B.  Differences remove c_univ:
	// S_i-S_j = (b_i-b_j)L_* + (r_i-r_j)z.
	b12, r12, s12 := beta.At(0)-beta.At(1), r.At(0)-r.At(1), s.At(0)-s.At(1)
	b13, r13, s13 := beta.At(0)-beta.At(2), r.At(0)-r.At(2), s.At(0)-s.At(2)
	det := b12*r13 - r12*b13
	out := UniversalCompletionMatch{CandidateShape: c}
	if math.Abs(det) < 1e-12 || !finite(det) {
		out.Verdict = "universal-completion shape system is singular; the row is triangle-invisible or degenerate"
		return out
	}
	x := (s12*r13 - r12*s13) / det
	z := (b12*s13 - s12*b13) / det
	if math.Abs(z) < 1e-12 || !finite(z) {
		out.Verdict = "universal-completion lever arm is singular"
		return out
	}
	cu := (s.At(0) - beta.At(0)*x - r.At(0)*z) / z
	y := x - z
	d := FloatTriple{r.At(0) + cu, r.At(1) + cu, r.At(2) + cu}
	res := rawResiduals(s, beta, d, x, y)
	maxRes := maxAbs(res[:])
	ordered := x > y && y > 0 && finite(x) && finite(y)
	nonneg := cu >= 0 && finite(cu)
	exact := ordered && nonneg && maxRes < 1e-7
	out.LinearSystemSolvable = true
	out.UniversalDelta = cu
	out.TotalDeltaB = d
	out.BoundaryLogFromMZ = x
	out.ThresholdLogFromMZ = y
	out.BoundaryScaleGeV = scaleFromLog(prev, x)
	out.ThresholdScaleGeV = scaleFromLog(prev, y)
	out.MaxAbsResidual = maxRes
	out.PositiveOrderedScales = ordered
	out.NonnegativeUniversalRow = nonneg
	out.ExactUOneClosure = exact
	out.ConditionalAlive = exact
	out.FiniteDerived = false
	if exact {
		out.Verdict = "conditional shape resonance: rational non-universal row closes only after adding a real universal beta row that is not finite-derived"
	} else {
		out.Verdict = "universal completion rejected or unphysical under ordered-scale/nonnegative-universal criteria"
	}
	return out
}

func auditInternalSpectral(th threshold.Analysis, rep RepresentationSearchAudit) InternalSpectralAudit {
	return InternalSpectralAudit{
		BGapValue:                           th.BGap,
		BPositiveModeCount:                  len(th.BPositiveEigenvalues),
		ContactPartialModeCount:             len(th.ContactPartialOverlap),
		ScalarActiveModeCount:               len(th.ScalarActiveSpectrum),
		DimensionlessSpectralAnchorsKnown:   th.DimensionlessSpectralAnchorsAvailable,
		BGapHasRepresentationRow:            false,
		ContactModesHaveRepresentationRows:  false,
		ThresholdActivationRuleDerived:      th.ThresholdActivationRuleDerived,
		FiniteToContinuumMatchingDerived:    th.FiniteToContinuumMatchingDerived,
		StructuralBGapMatchFound:            false,
		StructuralContactMatchFound:         false,
		CountResonancePromoted:              false,
		ConditionalAlgebraicPredictionAlive: rep.ConditionalUniversalShapeMatchFound,
		Verdict:                             "B-sector gap and contact partial-overlap modes remain finite spectral anchors only; no representation row, activation law, or beta tensor is derived, so no structural B-sector match is claimed",
	}
}

func auditFirewall(prev topologicalboundaryviability.Analysis, boundary BoundaryInputAudit, family InverseFamilyAudit, rep RepresentationSearchAudit, internal InternalSpectralAudit) FirewallAudit {
	return FirewallAudit{
		Gate200Inherited:                         prev.Summary.PairwiseIntersectionsSolved && prev.Summary.MismatchTriangleNonzero,
		ObservedInputsUsedForFiniteDerivation:    boundary.ObservedInputsUsedForFiniteDerivation,
		TopologicalUOneDerived:                   false,
		TopologicalUOneAssumedAsConditionalAudit: true,
		BoundaryScaleDerived:                     boundary.BoundaryScaleDerived,
		AbsoluteMassPredicted:                    false,
		PhysicalUnificationClaimed:               false,
		ThresholdCorrectedPhysicalFitClaimed:     false,
		FiniteMatchingCorrectionsDerived:         false,
		FiniteToContinuumNormalizationDerived:    false,
		BsectorRepresentationDerived:             internal.StructuralBGapMatchFound || internal.StructuralContactMatchFound,
		StrictNullityBefore:                      prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                       prev.Firewall.StrictNullityAfter,
		PhenomenologyNullityBefore:               0,
		PhenomenologyNullityAfter:                0,
		PhysicalPredictionNullityBefore:          prev.Firewall.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:           prev.Firewall.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                      "Gate 202 — canonical B-sector/contact representation-row construction or universal-completion source audit",
		OpenRequirements: []string{
			"derive or seal the UV boundary scale M_* before emitting a unique Δb(M_B)",
			"derive a B-sector/contact representation row under SU(3)c × SU(2)L × U(1)Y",
			"derive an activation/decoupling rule for B-sector/contact modes",
			"derive whether any universal beta row is a complete multiplet, regulator tower, or forbidden fit parameter",
			"derive finite matching corrections before claiming a physical threshold-corrected RG fit",
		},
		Verdict: "inverse threshold equations are useful conditional predata; all finite, mass, matching, and physical-prediction firewalls remain sealed",
	}
}

func observedTriple(prev topologicalboundaryviability.Analysis) FloatTriple {
	return FloatTriple{prev.Ledger.Alpha1GUTInverse, prev.Ledger.Alpha2Inverse, prev.Ledger.Alpha3Inverse}
}

func targetSTriple(obs FloatTriple) FloatTriple {
	topAlphaInv := 4 * math.Pi
	return FloatTriple{
		2 * math.Pi * (obs.U1GUT - topAlphaInv),
		2 * math.Pi * (obs.SU2L - topAlphaInv),
		2 * math.Pi * (obs.SU3C - topAlphaInv),
	}
}

func rawResiduals(s FloatTriple, beta FloatTriple, d FloatTriple, x, y float64) [gaugeCount]float64 {
	return [gaugeCount]float64{
		(beta.At(0)+d.At(0))*x - d.At(0)*y - s.At(0),
		(beta.At(1)+d.At(1))*x - d.At(1)*y - s.At(1),
		(beta.At(2)+d.At(2))*x - d.At(2)*y - s.At(2),
	}
}

func scaleFromLog(prev topologicalboundaryviability.Analysis, l float64) float64 {
	if !finite(l) || l > 700 || l < -700 {
		return math.NaN()
	}
	return prev.Ledger.ScaleGeV * math.Exp(l)
}

func maxAbs(xs []float64) float64 {
	m := 0.0
	for _, x := range xs {
		if !finite(x) {
			return math.Inf(1)
		}
		if math.Abs(x) > m {
			m = math.Abs(x)
		}
	}
	return m
}

func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }

func FormatBoundary(b BoundaryInputAudit) string {
	return fmt.Sprintf("u*=%.6g alpha*^-1=%.9g sin2=%.6g seedOK=%t bSM=%s gate200=%t ledger=%t quarantined=%t observedFinite=%t MstarDerived=%t", b.TopologicalU, b.TopologicalAlphaInverse, b.DerivedSin2ThetaWSeed, b.SeedCompatibleWithGUTNormalization, b.SMBetaVector, b.SMBetaVectorAcceptedFromGate200, b.EmpiricalLedgerInherited, b.EmpiricalLedgerQuarantined, b.ObservedInputsUsedForFiniteDerivation, b.BoundaryScaleDerived)
}

func FormatInverseFamily(f InverseFamilyAudit) string {
	return fmt.Sprintf("eq=%q formula=%q threshold=%s boundary=%s singleMB=%t MstarFree=%t underdetermined=%t evaluableIfSealed=%t closesTriangle=%t enforcesU1=%t physicalClaim=%t", f.Equation, f.RequiredDeltaFormula, f.ThresholdLogSymbol, f.BoundaryLogSymbol, f.SingleThresholdScaleOnly, f.BoundaryScaleStillFree, f.UnderdeterminedByOneContinuousParameter, f.CanEvaluateIfBoundaryScaleSealed, f.MismatchTriangleClosedByConstruction, f.UOneBoundaryEnforcedByConstruction, f.PhysicalPredictionClaim)
}

func FormatPoint(p RequiredDeformationPoint) string {
	return fmt.Sprintf("LB=%.9g L*=%.9g lever=%.9g MB=%.9g GeV M*=%.9g GeV Δb=%s maxResidual=%.3g area=%.3g ordered=%t", p.ThresholdLogFromMZ, p.BoundaryLogFromMZ, p.LeverArmAboveThreshold, p.ThresholdScaleGeV, p.BoundaryScaleGeV, p.RequiredDeltaB, p.MaxAbsResidual, p.TriangleArea, p.ValidOrderedScales)
}

func FormatRawMatch(m RepresentationRawMatch) string {
	return fmt.Sprintf("%s %s Δb=%s solvable=%t LB=%.6g L*=%.6g MB=%.6g M*=%.6g residual=%.3g ordered=%t exact=%t verdict=%s", m.Candidate.Name, m.Candidate.SMRepresentation, m.Candidate.DeltaB, m.LinearSystemSolvable, m.ThresholdLogFromMZ, m.BoundaryLogFromMZ, m.ThresholdScaleGeV, m.BoundaryScaleGeV, m.MaxAbsResidual, m.PositiveOrderedScales, m.ExactUOneClosure, m.Verdict)
}

func FormatRawMatches(xs []RepresentationRawMatch, max int) string {
	if max <= 0 || max > len(xs) {
		max = len(xs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatRawMatch(xs[i]))
	}
	if max < len(xs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(xs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUniversalMatch(m UniversalCompletionMatch) string {
	return fmt.Sprintf("%s shape=%s c_univ=%.9g totalΔb=%s LB=%.6g L*=%.6g MB=%.6g M*=%.6g residual=%.3g ordered=%t nonneg=%t alive=%t finite=%t", m.CandidateShape.Name, m.CandidateShape.DeltaB, m.UniversalDelta, m.TotalDeltaB, m.ThresholdLogFromMZ, m.BoundaryLogFromMZ, m.ThresholdScaleGeV, m.BoundaryScaleGeV, m.MaxAbsResidual, m.PositiveOrderedScales, m.NonnegativeUniversalRow, m.ConditionalAlive, m.FiniteDerived)
}

func FormatUniversalMatches(xs []UniversalCompletionMatch, max int) string {
	if max <= 0 || max > len(xs) {
		max = len(xs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatUniversalMatch(xs[i]))
	}
	if max < len(xs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(xs)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRepresentation(r RepresentationSearchAudit) string {
	return fmt.Sprintf("rows=%d rawExact=%t rawNoGos=%d condUniversal=%t condCount=%d univFinite=%t rationalTotal=%t physicalClaim=%t raw=%s universal=%s", r.KnownRationalRowsAudited, r.RawExactKnownRepresentationFound, len(r.RawNoGoMatches), r.ConditionalUniversalShapeMatchFound, len(r.UniversalCompletionMatches), r.UniversalCompletionFiniteDerived, r.IntegerOrRationalTotalDeltaDerived, r.PhysicalRepresentationClaimed, FormatRawMatches(r.RawNoGoMatches, 4), FormatUniversalMatches(r.UniversalCompletionMatches, 4))
}

func FormatInternal(i InternalSpectralAudit) string {
	return fmt.Sprintf("Bgap=%.10f Bpos=%d contactPartial=%d scalarActive=%d anchors=%t Brep=%t contactRep=%t activation=%t matching=%t Bmatch=%t contactMatch=%t countPromoted=%t conditionalAlive=%t verdict=%s", i.BGapValue, i.BPositiveModeCount, i.ContactPartialModeCount, i.ScalarActiveModeCount, i.DimensionlessSpectralAnchorsKnown, i.BGapHasRepresentationRow, i.ContactModesHaveRepresentationRows, i.ThresholdActivationRuleDerived, i.FiniteToContinuumMatchingDerived, i.StructuralBGapMatchFound, i.StructuralContactMatchFound, i.CountResonancePromoted, i.ConditionalAlgebraicPredictionAlive, i.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g200=%t observedFinite=%t u1Derived=%t u1Conditional=%t MstarDerived=%t mass=%t unification=%t thresholdFit=%t finiteMatch=%t contNorm=%t BsectorRep=%t strict=%d->%d pheno=%d->%d pred=%d->%d next=%s", f.Gate200Inherited, f.ObservedInputsUsedForFiniteDerivation, f.TopologicalUOneDerived, f.TopologicalUOneAssumedAsConditionalAudit, f.BoundaryScaleDerived, f.AbsoluteMassPredicted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.FiniteMatchingCorrectionsDerived, f.FiniteToContinuumNormalizationDerived, f.BsectorRepresentationDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhenomenologyNullityBefore, f.PhenomenologyNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d g200=%t family=%t singleScaleNoGo=%t repNoGo=%t condUniversal=%t internalMatch=%t noPrediction=%t comment=%s", s.TestsAudited, s.Gate200Inherited, s.InverseFamilyConstructed, s.SingleScaleUniquePredictionRejected, s.KnownRationalRepresentationNoGoLogged, s.ConditionalUniversalShapeMatchesLogged, s.InternalBsectorMatchFound, s.NoPhysicalPredictionClaim, s.Comment)
}
