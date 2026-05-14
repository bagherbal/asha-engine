// Package scalarbundlemap implements Gate 189: scalar-bundle map / H_Phi
// projector identification audit.
//
// Gate 188 constructed an exact branchwise unordered pair of complementary
// trace-2 projectors {P_A,P_B} on the abstract quartic contact module. Gate 189
// asks whether this pair can be canonically identified with the older physical
// active scalar carrier H_Phi, whose scalar response has an asymmetric high/high
// and low/low spectrum.
//
// The central distinction is compatibility versus trivialization. The abstract
// branch pair and the physical high/low pair have the same 2+2 dimensions, so
// branchwise scalar-bundle maps exist. But selecting P_A -> P_high rather than
// P_low breaks the eta -> -eta involution from Gate 188. Unless an already
// derived eta-odd finite source pulls back from matter/topology/contact data,
// the physical H_Phi bundle remains untrivialized.
package scalarbundlemap

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/branchprojector"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarvacuum"
	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalnormalization"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/charge"
)

type PhysicalProjectorRecord struct {
	Name                  string
	Meaning               string
	Matrix                string
	Trace                 float64
	IdempotenceResidual   float64
	OrthogonalResidual    float64
	CompletenessResidual  float64
	Dimension             int
	Eigenvalue            float64
	Multiplicity          int
	ProjectorLawSatisfied bool
}

type TargetCarrierAudit struct {
	CarrierName                 string
	Dimension                   int
	Spectrum                    []float64
	HighEigenvalue              float64
	LowEigenvalue               float64
	PairSplit                   float64
	HighMultiplicity            int
	LowMultiplicity             int
	PairDegenerate              bool
	AsymmetricHighLowSpectrum   bool
	HighProjector               PhysicalProjectorRecord
	LowProjector                PhysicalProjectorRecord
	ProjectorsOrthogonal        bool
	ProjectorsComplete          bool
	DimensionallyCompatible     bool
	PhysicalProjectorPairExists bool
	Verdict                     string
}

type AbstractBranchAudit struct {
	SourcePackage               string
	ProjectorPairConstructed    bool
	TraceTwoEach                bool
	DimensionTwoEach            bool
	EtaInvolutionSwapsPair      bool
	UnorderedPairInvariant      bool
	IndividualRootProjectors    int
	CanonicalBranchSelected     bool
	PhysicalScalarBundleDerived bool
	Verdict                     string
}

type IntertwinerAudit struct {
	AbstractPairDimensionallyMatchesTarget bool
	Assignments                            []string
	AssignmentCount                        int
	BranchwiseIntertwinersExist            bool
	CanonicalAssignmentDerived             bool
	RequiresEtaOrientationChoice           bool
	EtaToHighLowBreaksInvolution           bool
	PhysicalScalarBundleMapDerived         bool
	Verdict                                string
}

type SourceCandidateAudit struct {
	Name                     string
	Source                   string
	Domain                   string
	Available                bool
	ActsOnQuarticModule      bool
	ActsOnPhysicalHphi       bool
	EtaOdd                   bool
	SelectsHighLowAssignment bool
	Obstruction              string
	Verdict                  string
}

type OrientationSourceAudit struct {
	CandidatesAudited            []SourceCandidateAudit
	MatterSideOperatorsAudited   bool
	BLPullbackAudited            bool
	TopologicalSealAudited       bool
	ScalarResponseOrdersTarget   bool
	EtaOddSourceFound            bool
	CanonicalOrientationDerived  bool
	PhysicalHighLowPullbackFound bool
	Verdict                      string
}

type TrivializationAudit struct {
	AssumeAssignment              string
	MapsExist                     bool
	UniqueIntertwiner             bool
	ContinuousGaugeFreedom        string
	RealBundleFreedomDimension    int
	MetricFixedResidualFreedom    string
	ComplexFrameFreedom           string
	RequiresGaugeFrameChoice      bool
	SU2FrameDerived               bool
	CanonicalChangeOfBasisDerived bool
	Verdict                       string
}

type Summary struct {
	TestsAudited                      int
	BranchProjectorsInherited         bool
	PhysicalHighLowProjectorsVerified bool
	DimensionallyCompatible           bool
	IntertwinersExist                 bool
	CanonicalEtaHighLowAssignment     bool
	MatterPullbackBreaksEta           bool
	UniqueBundleTrivializationDerived bool
	PhysicalScalarBundleDerived       bool
	Comment                           string
}

type Firewall struct {
	UsesObservedInputForDerivation    bool
	UsesNumericRootApproximation      bool
	UsesIndividualRootDiagonalization bool
	UsesArbitraryEtaHighLowAssignment bool
	BranchProjectorPairInherited      bool
	PhysicalHighLowProjectorsVerified bool
	DimensionCompatibilityDerived     bool
	EtaOddSourceDerived               bool
	CanonicalEtaOrientationDerived    bool
	UniqueBundleTrivializationDerived bool
	ConditionalBundleMapsExist        bool
	PhysicalScalarBundleDerived       bool
	ChernWeilCarrierDerived           bool
	HeatKernelMatchingDerived         bool
	ThresholdCorrectedBetaDerived     bool
	AbsoluteCouplingPromoted          bool
	PhysicalConstantsDerived          bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalNullityBefore          int
	ConditionalNullityAfter           int
	ClosedStatements                  []string
	OpenRequirements                  []string
	RecommendedNextGate               string
	Verdict                           string
}

type Analysis struct {
	PreviousGate188 branchprojector.Analysis
	ScalarVacuum    scalarvacuum.Analysis
	ScalarComplex   scalarcomplex.Analysis
	Charge          charge.Analysis
	Topological     topologicalnormalization.Analysis

	AbstractBranch AbstractBranchAudit
	TargetCarrier  TargetCarrierAudit
	Intertwiner    IntertwinerAudit
	Sources        OrientationSourceAudit
	Trivialization TrivializationAudit
	Summary        Summary
	Firewall       Firewall
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := branchprojector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 188 input: %w", err)
			return
		}
		vac, err := scalarvacuum.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar-vacuum input: %w", err)
			return
		}
		complex, err := scalarcomplex.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar-complex input: %w", err)
			return
		}
		bl, err := charge.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build B-L charge input: %w", err)
			return
		}
		top, err := topologicalnormalization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build topological-normalization input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, vac, complex, bl, top, 1e-8)
	})
	return defaultA, defaultErr
}

func Build(prev branchprojector.Analysis, vac scalarvacuum.Analysis, complex scalarcomplex.Analysis, bl charge.Analysis, top topologicalnormalization.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	if !prev.Firewall.BranchwiseProjectorPairDerived || !prev.Firewall.ConditionalScalarProjectorDerived {
		return Analysis{}, fmt.Errorf("Gate 189 requires Gate 188 branchwise scalar projectors")
	}
	if prev.Firewall.PhysicalScalarBundleDerived || prev.Firewall.CanonicalScalarProjectorDerived || prev.Firewall.CanonicalUniqueBranchDerived {
		return Analysis{}, fmt.Errorf("Gate 189 expects Gate 188 to stop before physical scalar-bundle identification")
	}
	if vac.ActiveRealDimension != 4 || !vac.LowPairSelected || vac.HighPairDimension != 2 || vac.LowPairDimension != 2 {
		return Analysis{}, fmt.Errorf("Gate 189 requires a physical H_Phi high/low 2+2 target spectrum")
	}

	abstract := auditAbstractBranch(prev)
	target, err := auditTargetCarrier(vac, eps)
	if err != nil {
		return Analysis{}, err
	}
	intertwiner := auditIntertwiner(abstract, target)
	sources := auditOrientationSources(vac, complex, bl, top)
	triv := auditTrivialization(intertwiner, complex)
	summary := Summary{
		TestsAudited:                      5,
		BranchProjectorsInherited:         abstract.ProjectorPairConstructed,
		PhysicalHighLowProjectorsVerified: target.PhysicalProjectorPairExists,
		DimensionallyCompatible:           target.DimensionallyCompatible && intertwiner.AbstractPairDimensionallyMatchesTarget,
		IntertwinersExist:                 intertwiner.BranchwiseIntertwinersExist,
		CanonicalEtaHighLowAssignment:     intertwiner.CanonicalAssignmentDerived || sources.CanonicalOrientationDerived,
		MatterPullbackBreaksEta:           sources.EtaOddSourceFound,
		UniqueBundleTrivializationDerived: triv.UniqueIntertwiner && triv.CanonicalChangeOfBasisDerived,
		PhysicalScalarBundleDerived:       false,
		Comment:                           "Gate 189 verifies dimensional compatibility between the branchwise abstract projector pair and the physical H_Phi high/low projector pair, then proves the eta-to-high/low assignment and bundle trivialization are not canonical with current finite data.",
	}
	fw := Firewall{
		UsesObservedInputForDerivation:    false,
		UsesNumericRootApproximation:      false,
		UsesIndividualRootDiagonalization: false,
		UsesArbitraryEtaHighLowAssignment: false,
		BranchProjectorPairInherited:      abstract.ProjectorPairConstructed,
		PhysicalHighLowProjectorsVerified: target.PhysicalProjectorPairExists,
		DimensionCompatibilityDerived:     target.DimensionallyCompatible,
		EtaOddSourceDerived:               sources.EtaOddSourceFound,
		CanonicalEtaOrientationDerived:    sources.CanonicalOrientationDerived || intertwiner.CanonicalAssignmentDerived,
		UniqueBundleTrivializationDerived: triv.UniqueIntertwiner && triv.CanonicalChangeOfBasisDerived,
		ConditionalBundleMapsExist:        intertwiner.BranchwiseIntertwinersExist,
		PhysicalScalarBundleDerived:       false,
		ChernWeilCarrierDerived:           false,
		HeatKernelMatchingDerived:         false,
		ThresholdCorrectedBetaDerived:     false,
		AbsoluteCouplingPromoted:          false,
		PhysicalConstantsDerived:          false,
		StrictNullityBefore:               prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:          prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:           prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"the physical H_Phi scalar response has exact 2+2 high/low projectors compatible in dimension with the branchwise abstract projector pair",
			"branchwise intertwiners exist after choosing an eta orientation and a gauge frame",
			"B-L/Fock, scalar response, scalar complex data, and the topological action seal do not currently provide an eta-odd pullback to select P_A -> P_high rather than P_low",
			"even after an assignment is assumed, the change-of-basis map has GL(2,R) x GL(2,R) frame freedom and is not unique",
		},
		OpenRequirements: []string{
			"derive an eta-odd finite source or prove that eta orientation is pure spontaneous/gauge data",
			"derive a canonical metric/complex/SU(2) frame on the branchwise 2D scalar planes before claiming a unique H_Phi trivialization",
			"construct a lawful scalar-bundle map only as conditional orientation data unless a finite source breaks the eta involution",
			"keep Chern-Weil, heat-kernel, threshold beta rows, absolute coupling promotion, and physical constants sealed until the scalar bundle is actually trivialized",
		},
		RecommendedNextGate: "Gate 190 — eta-odd scalar-orientation source / matter-pullback search audit",
		Verdict:             "Gate 189 is an orientation obstruction theorem: the 2+2 fibers match, but the physical H_Phi bundle remains untrivialized because eta-to-high/low assignment and the gauge frame are not canonically derived.",
	}
	truth := "Gate 189 proves compatibility without identification. The branchwise projectors {P_A,P_B} and the physical H_Phi projectors {P_high,P_low} are both complementary trace-2 pairs, so scalar-bundle maps exist after a spontaneous orientation choice. But mapping eta to high or low breaks the eta -> -eta involution, and no audited matter-side, scalar-side, or topological source currently supplies an eta-odd pullback. The physical scalar bundle is therefore not yet strictly derived."
	return Analysis{PreviousGate188: prev, ScalarVacuum: vac, ScalarComplex: complex, Charge: bl, Topological: top, AbstractBranch: abstract, TargetCarrier: target, Intertwiner: intertwiner, Sources: sources, Trivialization: triv, Summary: summary, Firewall: fw, TruthStatement: truth}, nil
}

func auditAbstractBranch(prev branchprojector.Analysis) AbstractBranchAudit {
	return AbstractBranchAudit{
		SourcePackage:               "pkg/bridge/branchprojector",
		ProjectorPairConstructed:    prev.Projectors.ProjectorPairConstructed,
		TraceTwoEach:                prev.Projectors.TraceTwoEach,
		DimensionTwoEach:            prev.Projectors.DimensionTwoEach,
		EtaInvolutionSwapsPair:      prev.Factors.FactorSwapInvolutionPreserved,
		UnorderedPairInvariant:      prev.Factors.FactorSwapInvolutionPreserved && !prev.Projectors.CanonicalBranchSelected,
		IndividualRootProjectors:    prev.Projectors.IndividualRootProjectors,
		CanonicalBranchSelected:     prev.Projectors.CanonicalBranchSelected,
		PhysicalScalarBundleDerived: prev.Projectors.PhysicalScalarBundleDerived,
		Verdict:                     "Gate 188 supplies an unordered pair of exact trace-2 projectors. The eta involution exchanges P_A and P_B, so neither member is yet physically high or low.",
	}
}

func auditTargetCarrier(vac scalarvacuum.Analysis, eps float64) (TargetCarrierAudit, error) {
	highP := linear.Diagonal([]float64{1, 1, 0, 0})
	lowP := linear.Diagonal([]float64{0, 0, 1, 1})
	h, err := projectorRecord("P_high", "high scalar-response pair", highP, lowP, vac.HighPairEigenvalue, vac.HighPairDimension, eps)
	if err != nil {
		return TargetCarrierAudit{}, err
	}
	l, err := projectorRecord("P_low", "low scalar-response pair", lowP, highP, vac.LowPairEigenvalue, vac.LowPairDimension, eps)
	if err != nil {
		return TargetCarrierAudit{}, err
	}
	sum, _ := highP.Add(lowP)
	complete := residualToIdentity(sum) <= eps
	orth, _ := highP.Mul(lowP)
	orthogonal := orth.FrobeniusNorm() <= eps
	pairDeg := vac.HighPairDimension == 2 && vac.LowPairDimension == 2 && vac.PairSplit > eps
	return TargetCarrierAudit{
		CarrierName:                 "H_Phi active scalar carrier",
		Dimension:                   vac.ActiveRealDimension,
		Spectrum:                    append([]float64(nil), vac.ActiveSpectrum...),
		HighEigenvalue:              vac.HighPairEigenvalue,
		LowEigenvalue:               vac.LowPairEigenvalue,
		PairSplit:                   vac.PairSplit,
		HighMultiplicity:            vac.HighPairDimension,
		LowMultiplicity:             vac.LowPairDimension,
		PairDegenerate:              pairDeg,
		AsymmetricHighLowSpectrum:   pairDeg && math.Abs(vac.HighPairEigenvalue-vac.LowPairEigenvalue) > eps,
		HighProjector:               h,
		LowProjector:                l,
		ProjectorsOrthogonal:        orthogonal,
		ProjectorsComplete:          complete,
		DimensionallyCompatible:     pairDeg && h.Dimension == 2 && l.Dimension == 2,
		PhysicalProjectorPairExists: h.ProjectorLawSatisfied && l.ProjectorLawSatisfied && orthogonal && complete,
		Verdict:                     "The physical H_Phi target has a real asymmetric high/high + low/low scalar-response spectrum with two complementary trace-2 projectors.",
	}, nil
}

func projectorRecord(name, meaning string, p, other linear.Matrix, eig float64, mult int, eps float64) (PhysicalProjectorRecord, error) {
	tr, err := p.Trace()
	if err != nil {
		return PhysicalProjectorRecord{}, err
	}
	sq, err := p.Mul(p)
	if err != nil {
		return PhysicalProjectorRecord{}, err
	}
	idem, err := sq.Sub(p)
	if err != nil {
		return PhysicalProjectorRecord{}, err
	}
	orth, err := p.Mul(other)
	if err != nil {
		return PhysicalProjectorRecord{}, err
	}
	sum, err := p.Add(other)
	if err != nil {
		return PhysicalProjectorRecord{}, err
	}
	complete := residualToIdentity(sum)
	law := math.Abs(tr-2) <= eps && idem.FrobeniusNorm() <= eps && orth.FrobeniusNorm() <= eps && complete <= eps && mult == 2
	return PhysicalProjectorRecord{
		Name:                  name,
		Meaning:               meaning,
		Matrix:                diagonalMatrixString(p),
		Trace:                 tr,
		IdempotenceResidual:   idem.FrobeniusNorm(),
		OrthogonalResidual:    orth.FrobeniusNorm(),
		CompletenessResidual:  complete,
		Dimension:             int(math.Round(tr)),
		Eigenvalue:            eig,
		Multiplicity:          mult,
		ProjectorLawSatisfied: law,
	}, nil
}

func auditIntertwiner(a AbstractBranchAudit, t TargetCarrierAudit) IntertwinerAudit {
	match := a.ProjectorPairConstructed && a.TraceTwoEach && a.DimensionTwoEach && t.DimensionallyCompatible && t.PhysicalProjectorPairExists
	assignments := []string{"P_A -> P_high, P_B -> P_low", "P_A -> P_low, P_B -> P_high"}
	return IntertwinerAudit{
		AbstractPairDimensionallyMatchesTarget: match,
		Assignments:                            assignments,
		AssignmentCount:                        len(assignments),
		BranchwiseIntertwinersExist:            match,
		CanonicalAssignmentDerived:             false,
		RequiresEtaOrientationChoice:           match,
		EtaToHighLowBreaksInvolution:           match && a.EtaInvolutionSwapsPair && t.AsymmetricHighLowSpectrum,
		PhysicalScalarBundleMapDerived:         false,
		Verdict:                                "There are exactly two eta-orientation assignments between the unordered abstract pair and the ordered high/low target pair. Existing data proves existence after a choice, not a canonical choice.",
	}
}

func auditOrientationSources(vac scalarvacuum.Analysis, complex scalarcomplex.Analysis, bl charge.Analysis, top topologicalnormalization.Analysis) OrientationSourceAudit {
	candidates := []SourceCandidateAudit{
		{
			Name:                     "Gate-37/86 scalar response high-low ordering",
			Source:                   "scalarpotential/scalarvacuum",
			Domain:                   "physical H_Phi active frame",
			Available:                vac.LowPairSelected && vac.HighPairDimension == 2 && vac.LowPairDimension == 2,
			ActsOnQuarticModule:      false,
			ActsOnPhysicalHphi:       true,
			EtaOdd:                   false,
			SelectsHighLowAssignment: false,
			Obstruction:              "orders the target high/low planes but supplies no pullback from the abstract quartic projector pair and no eta-odd action",
			Verdict:                  "target asymmetry exists, but it is not an abstract-branch orientation selector",
		},
		{
			Name:                     "B-L / Fock charge polarization",
			Source:                   "pkg/matter/charge",
			Domain:                   "one-particle/Fock matter carrier",
			Available:                bl.ChargePolarizesOnePlusThree,
			ActsOnQuarticModule:      false,
			ActsOnPhysicalHphi:       false,
			EtaOdd:                   false,
			SelectsHighLowAssignment: false,
			Obstruction:              "B-L supplies a 1+3 matter polarization, not a 2+2 eta-odd operator on the quartic scalar branch; no canonical scalar pullback is derived",
			Verdict:                  "audited and obstructed as final high/low scalar-orientation source",
		},
		{
			Name:                     "pair-compatible scalar complex structure",
			Source:                   "pkg/bridge/scalarcomplex",
			Domain:                   "physical active scalar planes",
			Available:                complex.PairCompatibleComplexAvailable,
			ActsOnQuarticModule:      false,
			ActsOnPhysicalHphi:       true,
			EtaOdd:                   false,
			SelectsHighLowAssignment: false,
			Obstruction:              "a commuting complex candidate exists on the physical pair planes, but its signs/orientations are already recorded as noncanonical and do not orient eta on the quartic branch",
			Verdict:                  "supports future frame data, but not a scalar-bundle identification theorem",
		},
		{
			Name:                     "topological action seal S_top=8*pi^2",
			Source:                   "pkg/bridge/topologicalnormalization",
			Domain:                   "dimensionless topological/gauge normalization branch",
			Available:                top.Input.TopologicalSealAvailable,
			ActsOnQuarticModule:      false,
			ActsOnPhysicalHphi:       false,
			EtaOdd:                   false,
			SelectsHighLowAssignment: false,
			Obstruction:              "the seal is scalar-orientation neutral and already conditional on continuum index/trace bridges; it does not couple to eta",
			Verdict:                  "no eta-breaking scalar orientation source",
		},
	}
	etaOdd := false
	selects := false
	for _, c := range candidates {
		etaOdd = etaOdd || (c.Available && c.EtaOdd)
		selects = selects || (c.Available && c.SelectsHighLowAssignment)
	}
	return OrientationSourceAudit{
		CandidatesAudited:            candidates,
		MatterSideOperatorsAudited:   true,
		BLPullbackAudited:            true,
		TopologicalSealAudited:       true,
		ScalarResponseOrdersTarget:   vac.LowPairSelected,
		EtaOddSourceFound:            etaOdd,
		CanonicalOrientationDerived:  selects,
		PhysicalHighLowPullbackFound: false,
		Verdict:                      "No audited finite operator supplies an eta-odd pullback that distinguishes P_A -> P_high from P_A -> P_low. The high/low target is ordered, but the abstract branch pair is still eta-symmetric.",
	}
}

func auditTrivialization(i IntertwinerAudit, complex scalarcomplex.Analysis) TrivializationAudit {
	mapsExist := i.BranchwiseIntertwinersExist
	return TrivializationAudit{
		AssumeAssignment:              "for example P_A -> P_high and P_B -> P_low",
		MapsExist:                     mapsExist,
		UniqueIntertwiner:             false,
		ContinuousGaugeFreedom:        "GL(2,R)_A->high x GL(2,R)_B->low before metric/complex frame fixing",
		RealBundleFreedomDimension:    8,
		MetricFixedResidualFreedom:    "O(2) x O(2), or SO(2) x SO(2) after orientation is imposed",
		ComplexFrameFreedom:           "pair-compatible complex structure is available but noncanonical; an SU(2)/U(2)-style frame choice is still required",
		RequiresGaugeFrameChoice:      mapsExist,
		SU2FrameDerived:               complex.FullScalarSU2Recovered,
		CanonicalChangeOfBasisDerived: false,
		Verdict:                       "Even after an eta-to-high/low assignment is assumed, no unique change-of-basis matrix W is derived. The bundle map needs additional gauge-frame/trivialization data.",
	}
}

func residualToIdentity(m linear.Matrix) float64 {
	id := linear.Identity(m.Rows())
	d, err := m.Sub(id)
	if err != nil {
		return math.Inf(1)
	}
	return d.FrobeniusNorm()
}

func diagonalMatrixString(m linear.Matrix) string {
	vals := make([]string, 0, m.Rows())
	for i := 0; i < m.Rows(); i++ {
		vals = append(vals, fmt.Sprintf("%.0f", m.At(i, i)))
	}
	return "diag(" + strings.Join(vals, ",") + ")"
}

func FormatAbstractBranch(a AbstractBranchAudit) string {
	return fmt.Sprintf("source=%s constructed=%t trace2=%t dim2=%t etaSwap=%t unordered=%t rootProjectors=%d canonicalBranch=%t physicalBundle=%t (%s)", a.SourcePackage, a.ProjectorPairConstructed, a.TraceTwoEach, a.DimensionTwoEach, a.EtaInvolutionSwapsPair, a.UnorderedPairInvariant, a.IndividualRootProjectors, a.CanonicalBranchSelected, a.PhysicalScalarBundleDerived, a.Verdict)
}

func FormatTargetCarrier(a TargetCarrierAudit) string {
	return fmt.Sprintf("carrier=%s dim=%d spectrum=%v high=%.12g x%d low=%.12g x%d split=%.12g pair=%t asym=%t PH={%s tr=%.0f idem=%.2g orth=%.2g complete=%.2g} PL={%s tr=%.0f idem=%.2g orth=%.2g complete=%.2g} orthogonal=%t complete=%t compatible=%t exists=%t (%s)", a.CarrierName, a.Dimension, a.Spectrum, a.HighEigenvalue, a.HighMultiplicity, a.LowEigenvalue, a.LowMultiplicity, a.PairSplit, a.PairDegenerate, a.AsymmetricHighLowSpectrum, a.HighProjector.Matrix, a.HighProjector.Trace, a.HighProjector.IdempotenceResidual, a.HighProjector.OrthogonalResidual, a.HighProjector.CompletenessResidual, a.LowProjector.Matrix, a.LowProjector.Trace, a.LowProjector.IdempotenceResidual, a.LowProjector.OrthogonalResidual, a.LowProjector.CompletenessResidual, a.ProjectorsOrthogonal, a.ProjectorsComplete, a.DimensionallyCompatible, a.PhysicalProjectorPairExists, a.Verdict)
}

func FormatIntertwiner(a IntertwinerAudit) string {
	return fmt.Sprintf("matches=%t assignments=%v count=%d exists=%t canonical=%t etaChoice=%t breaksEta=%t physicalMap=%t (%s)", a.AbstractPairDimensionallyMatchesTarget, a.Assignments, a.AssignmentCount, a.BranchwiseIntertwinersExist, a.CanonicalAssignmentDerived, a.RequiresEtaOrientationChoice, a.EtaToHighLowBreaksInvolution, a.PhysicalScalarBundleMapDerived, a.Verdict)
}

func FormatSources(a OrientationSourceAudit) string {
	parts := make([]string, 0, len(a.CandidatesAudited))
	for _, c := range a.CandidatesAudited {
		parts = append(parts, fmt.Sprintf("%s available=%t quartic=%t hphi=%t etaOdd=%t selects=%t obstruction=%q", c.Name, c.Available, c.ActsOnQuarticModule, c.ActsOnPhysicalHphi, c.EtaOdd, c.SelectsHighLowAssignment, c.Obstruction))
	}
	return fmt.Sprintf("matter=%t BL=%t top=%t scalarOrders=%t etaOdd=%t canonical=%t pullback=%t candidates=[%s] (%s)", a.MatterSideOperatorsAudited, a.BLPullbackAudited, a.TopologicalSealAudited, a.ScalarResponseOrdersTarget, a.EtaOddSourceFound, a.CanonicalOrientationDerived, a.PhysicalHighLowPullbackFound, strings.Join(parts, "; "), a.Verdict)
}

func FormatTrivialization(a TrivializationAudit) string {
	return fmt.Sprintf("assume=%q maps=%t unique=%t freedom=%q realDim=%d metricFixed=%q complex=%q gaugeFrame=%t su2=%t canonicalW=%t (%s)", a.AssumeAssignment, a.MapsExist, a.UniqueIntertwiner, a.ContinuousGaugeFreedom, a.RealBundleFreedomDimension, a.MetricFixedResidualFreedom, a.ComplexFrameFreedom, a.RequiresGaugeFrameChoice, a.SU2FrameDerived, a.CanonicalChangeOfBasisDerived, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d branch=%t target=%t compatible=%t maps=%t assignment=%t matterEta=%t uniqueW=%t physical=%t (%s)", a.TestsAudited, a.BranchProjectorsInherited, a.PhysicalHighLowProjectorsVerified, a.DimensionallyCompatible, a.IntertwinersExist, a.CanonicalEtaHighLowAssignment, a.MatterPullbackBreaksEta, a.UniqueBundleTrivializationDerived, a.PhysicalScalarBundleDerived, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t numeric=%t rootDiag=%t arbitraryEta=%t branch=%t target=%t compatible=%t etaOdd=%t etaOrient=%t uniqueW=%t conditionalMaps=%t physicalBundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.UsesObservedInputForDerivation, a.UsesNumericRootApproximation, a.UsesIndividualRootDiagonalization, a.UsesArbitraryEtaHighLowAssignment, a.BranchProjectorPairInherited, a.PhysicalHighLowProjectorsVerified, a.DimensionCompatibilityDerived, a.EtaOddSourceDerived, a.CanonicalEtaOrientationDerived, a.UniqueBundleTrivializationDerived, a.ConditionalBundleMapsExist, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}
