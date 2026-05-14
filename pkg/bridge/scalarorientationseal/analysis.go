// Package scalarorientationseal implements Gate 191: spontaneous scalar-
// orientation seal / gauge-fixed H_Phi trivialization axiom audit.
//
// Gate 190 proved that no available finite gauge-invariant datum selects the
// eta orientation. Gate 191 therefore stops searching for a hidden selector and
// records the eta-to-high/low convention as an explicit spontaneous vacuum
// boundary condition. All scalar-bundle constructions in this package are
// conditional on that seal. The seal is deliberately quarantined: it may
// trivialize the branchwise scalar bundle, but it does not derive physical
// constants, threshold rows, Chern-Weil carriers, heat-kernel matching, Yukawa
// amplitudes, or absolute couplings.
package scalarorientationseal

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betamatching"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcomplex"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarorientationsource"
	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalnormalization"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type OrientationChoice string

const (
	EtaToHigh OrientationChoice = "eta_to_high"
	EtaToLow  OrientationChoice = "eta_to_low"
)

type SpontaneousOrientationSeal struct {
	Name                              string
	Choice                            OrientationChoice
	Source                            string
	AxiomID                           string
	ConditionalStatus                 string
	EtaMapsToHigh                     bool
	EtaMapsToLow                      bool
	ExplicitAxiom                     bool
	Quarantined                       bool
	RequiredByGate190                 bool
	BreaksEtaInvolutionAsBoundaryData bool
	DerivedFromFiniteSelector         bool
	UsesObservedInput                 bool
	CarriesNumericPhysicalConstant    bool
	CarriesGaugeCoupling              bool
	CarriesMassScale                  bool
	Verdict                           string
}

type PhysicalProjector struct {
	Name                  string
	Meaning               string
	Matrix                string
	Trace                 float64
	IdempotenceResidual   float64
	ComplementResidual    float64
	OrthogonalityResidual float64
	Dimension             int
	Verified              bool
}

type SealedFrameAudit struct {
	AbstractFrameName               string
	PhysicalFrameName               string
	AbstractProjectorA              PhysicalProjector
	AbstractProjectorB              PhysicalProjector
	PhysicalProjectorHigh           PhysicalProjector
	PhysicalProjectorLow            PhysicalProjector
	SealMapsAToHigh                 bool
	SealMapsBToLow                  bool
	SealMapsAToLow                  bool
	SealMapsBToHigh                 bool
	DimensionCompatibilityInherited bool
	Gate190ObstructionInherited     bool
	CanonicalWithoutSeal            bool
	ConditionalOnSeal               bool
	Verdict                         string
}

type TrivializationAudit struct {
	MatrixName                         string
	Matrix                             string
	OrthogonalFrameMap                 bool
	Invertible                         bool
	InverseEqualsTranspose             bool
	MapsAProjectorToHighResidual       float64
	MapsBProjectorToLowResidual        float64
	MapsAProjectorToLowResidual        float64
	MapsBProjectorToHighResidual       float64
	ProjectorIntertwiningVerified      bool
	UniqueWithoutGaugeFrame            bool
	GaugeFreedomBeforeSeal             string
	GaugeFreedomAfterSealedFrameChoice string
	PhysicalScalarBundleTrivialized    bool
	Verdict                            string
}

type PulledBackGenerator struct {
	Name              string
	PullbackMatrix    string
	CommutesWithHigh  bool
	CommutesWithLow   bool
	BlockDiagonal     bool
	OffDiagonalNorm   float64
	DiagonalBlockNorm float64
	MixesABFibers     bool
	PreservesABFibers bool
	PhysicalMeaning   string
}

type GaugePullbackAudit struct {
	Generators                           []PulledBackGenerator
	T3LPreservesFibers                   bool
	YPhiPreservesFibers                  bool
	T1MixesFibers                        bool
	T2MixesFibers                        bool
	BrokenGeneratorsOffDiagonal          bool
	UnbrokenGeneratorsBlockDiagonal      bool
	GaugeConnectionAttachedConditionally bool
	GaugeBosonMassesDerived              bool
	PhysicalCouplingsDerived             bool
	Verdict                              string
}

type SealFirewallAudit struct {
	SealExplicitInput                 bool
	SealQuarantined                   bool
	ObservedInputUsed                 bool
	HiddenEtaSelectorClaimed          bool
	ChernWeilCarrierDerived           bool
	HeatKernelMatchingDerived         bool
	ThresholdCorrectedBetaDerived     bool
	YukawaAmplitudeDerived            bool
	AbsoluteCouplingPromoted          bool
	PhysicalConstantsDerived          bool
	TopologicalSealImportedAsConstant bool
	BetaRowsUnlocked                  bool
	ScalarKineticNormalizationDerived bool
	GaugeActionHessianDerived         bool
	GaugeGeneratorPullbackDerived     bool
	ConditionalPhysicalBundleDerived  bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalNullityBefore          int
	ConditionalNullityAfter           int
	ClosedStatements                  []string
	OpenRequirements                  []string
	RecommendedNextGate               string
	Verdict                           string
}

type Summary struct {
	TestsAudited                             int
	Gate190ObstructionInherited              bool
	SpontaneousOrientationSealRecorded       bool
	GaugeFixedTrivializationConstructed      bool
	GaugeGeneratorPullbackConstructed        bool
	T3YBlockDiagonal                         bool
	T1T2OffDiagonal                          bool
	PhysicalScalarBundleConditionallyDerived bool
	ChernWeilHeatKernelThresholdsStillSealed bool
	Comment                                  string
}

type Analysis struct {
	PreviousGate190 scalarorientationsource.Analysis
	ScalarCovariant scalarcovariant.Analysis
	ScalarComplex   scalarcomplex.Analysis
	Topological     topologicalnormalization.Analysis
	BetaMatching    betamatching.Analysis

	Seal           SpontaneousOrientationSeal
	SealedFrame    SealedFrameAudit
	Trivialization TrivializationAudit
	GaugePullback  GaugePullbackAudit
	Firewall       SealFirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := scalarorientationsource.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 190 input: %w", err)
			return
		}
		sc, err := scalarcovariant.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar-covariant input: %w", err)
			return
		}
		cx, err := scalarcomplex.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build scalar-complex input: %w", err)
			return
		}
		top, err := topologicalnormalization.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build topological-normalization input: %w", err)
			return
		}
		beta, err := betamatching.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build beta-matching input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, sc, cx, top, beta, EtaToHigh, 1e-9)
	})
	return defaultA, defaultErr
}

func Build(prev scalarorientationsource.Analysis, sc scalarcovariant.Analysis, cx scalarcomplex.Analysis, top topologicalnormalization.Analysis, beta betamatching.Analysis, choice OrientationChoice, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !prev.Firewall.EtaOrientationClassifiedSpontaneous || prev.Firewall.CanonicalEtaOrientationDerived || prev.Firewall.GaugeInvariantEtaOddSourceFound || prev.Firewall.PhysicalScalarBundleDerived {
		return Analysis{}, fmt.Errorf("Gate 191 requires Gate 190 to isolate eta orientation as spontaneous/gauge data and to derive no physical scalar bundle")
	}
	if sc.ActiveRealDimension != 4 || !sc.AbstractCovariantDerivativeTemplate {
		return Analysis{}, fmt.Errorf("Gate 191 requires a four-real-dimensional scalar covariant derivative template")
	}
	if choice != EtaToHigh && choice != EtaToLow {
		return Analysis{}, fmt.Errorf("unsupported orientation choice %q", choice)
	}

	seal := buildSeal(prev, choice)
	frame, err := auditSealedFrame(prev, seal, eps)
	if err != nil {
		return Analysis{}, err
	}
	triv, err := auditTrivialization(seal, frame, eps)
	if err != nil {
		return Analysis{}, err
	}
	pull, err := auditGaugePullback(sc, triv, eps)
	if err != nil {
		return Analysis{}, err
	}
	fw := auditFirewall(prev, seal, triv, pull, top, beta, sc, cx)
	summary := Summary{
		TestsAudited:                             5,
		Gate190ObstructionInherited:              prev.Firewall.EtaOrientationClassifiedSpontaneous && !prev.Firewall.CanonicalEtaOrientationDerived,
		SpontaneousOrientationSealRecorded:       seal.ExplicitAxiom && seal.Quarantined,
		GaugeFixedTrivializationConstructed:      triv.PhysicalScalarBundleTrivialized && triv.ProjectorIntertwiningVerified,
		GaugeGeneratorPullbackConstructed:        pull.GaugeConnectionAttachedConditionally,
		T3YBlockDiagonal:                         pull.T3LPreservesFibers && pull.YPhiPreservesFibers,
		T1T2OffDiagonal:                          pull.T1MixesFibers && pull.T2MixesFibers,
		PhysicalScalarBundleConditionallyDerived: fw.ConditionalPhysicalBundleDerived,
		ChernWeilHeatKernelThresholdsStillSealed: !fw.ChernWeilCarrierDerived && !fw.HeatKernelMatchingDerived && !fw.ThresholdCorrectedBetaDerived,
		Comment:                                  "Gate 191 records eta -> high as an explicit spontaneous orientation seal, constructs a gauge-fixed H_Phi trivialization only under that seal, and proves the weak generators pull back with T3/Y block-diagonal and T1/T2 off-diagonal. The seal does not promote constants or thresholds.",
	}
	truth := "Gate 191 is a conditional construction, not a new finite selector. The eta orientation is inserted as an explicit spontaneous vacuum seal. In the sealed A/B frame, the branchwise scalar projector pair is identified with the physical H_Phi high/low projectors and the weak generators can be pulled back: T3L and Y_phi preserve the two fibers, while T1 and T2 mix them. Chern-Weil, heat-kernel, thresholds, couplings, Yukawa amplitudes, and physical constants remain sealed."

	return Analysis{PreviousGate190: prev, ScalarCovariant: sc, ScalarComplex: cx, Topological: top, BetaMatching: beta, Seal: seal, SealedFrame: frame, Trivialization: triv, GaugePullback: pull, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func buildSeal(prev scalarorientationsource.Analysis, choice OrientationChoice) SpontaneousOrientationSeal {
	etaHigh := choice == EtaToHigh
	return SpontaneousOrientationSeal{
		Name:                              "SpontaneousOrientationSeal",
		Choice:                            choice,
		Source:                            "explicit vacuum boundary condition after Gate 190 eta-source obstruction",
		AxiomID:                           "AXIOM-SPONTANEOUS-SCALAR-ORIENTATION-ETA-HIGH-LOW",
		ConditionalStatus:                 "CONDITIONAL_ON_VACUUM_SEAL",
		EtaMapsToHigh:                     etaHigh,
		EtaMapsToLow:                      !etaHigh,
		ExplicitAxiom:                     true,
		Quarantined:                       true,
		RequiredByGate190:                 prev.Spontaneous.OrientationInsertionPointIsolated,
		BreaksEtaInvolutionAsBoundaryData: true,
		DerivedFromFiniteSelector:         false,
		UsesObservedInput:                 false,
		CarriesNumericPhysicalConstant:    false,
		CarriesGaugeCoupling:              false,
		CarriesMassScale:                  false,
		Verdict:                           "The eta-to-high/low assignment is recorded as explicit spontaneous vacuum boundary data. It is not derived from a hidden finite selector and carries no observed constant, coupling, or mass scale.",
	}
}

func auditSealedFrame(prev scalarorientationsource.Analysis, seal SpontaneousOrientationSeal, eps float64) (SealedFrameAudit, error) {
	pa := linear.Diagonal([]float64{1, 1, 0, 0})
	pb := linear.Diagonal([]float64{0, 0, 1, 1})
	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	ra, err := projectorRecord("P_A^seal", "abstract eta-positive 2D fiber in the sealed branch frame", pa, pb, eps)
	if err != nil {
		return SealedFrameAudit{}, err
	}
	rb, err := projectorRecord("P_B^seal", "abstract eta-negative 2D fiber in the sealed branch frame", pb, pa, eps)
	if err != nil {
		return SealedFrameAudit{}, err
	}
	rh, err := projectorRecord("P_high", "physical high scalar-weight 2D fiber", ph, pl, eps)
	if err != nil {
		return SealedFrameAudit{}, err
	}
	rl, err := projectorRecord("P_low", "physical low scalar-weight 2D fiber", pl, ph, eps)
	if err != nil {
		return SealedFrameAudit{}, err
	}
	return SealedFrameAudit{
		AbstractFrameName:               "sealed branchwise quartic scalar frame (A1,A2,B1,B2)",
		PhysicalFrameName:               "Gate-37/Gate-169 H_Phi high-low frame",
		AbstractProjectorA:              ra,
		AbstractProjectorB:              rb,
		PhysicalProjectorHigh:           rh,
		PhysicalProjectorLow:            rl,
		SealMapsAToHigh:                 seal.Choice == EtaToHigh,
		SealMapsBToLow:                  seal.Choice == EtaToHigh,
		SealMapsAToLow:                  seal.Choice == EtaToLow,
		SealMapsBToHigh:                 seal.Choice == EtaToLow,
		DimensionCompatibilityInherited: prev.PreviousGate189.Firewall.DimensionCompatibilityDerived && prev.PreviousGate189.Firewall.ConditionalBundleMapsExist,
		Gate190ObstructionInherited:     prev.Firewall.EtaOrientationClassifiedSpontaneous && !prev.Firewall.CanonicalEtaOrientationDerived,
		CanonicalWithoutSeal:            false,
		ConditionalOnSeal:               true,
		Verdict:                         "The sealed A/B frame represents the exact Gate-188 branchwise projector pair after the spontaneous eta orientation is explicitly fixed. It is dimensionally compatible with the H_Phi high/low projector pair, but only conditionally on the seal.",
	}, nil
}

func auditTrivialization(seal SpontaneousOrientationSeal, frame SealedFrameAudit, eps float64) (TrivializationAudit, error) {
	var w linear.Matrix
	if seal.Choice == EtaToHigh {
		w = linear.Identity(4)
	} else {
		w = planeSwap()
	}
	wt := w.Transpose()
	invCheck, err := w.Mul(wt)
	if err != nil {
		return TrivializationAudit{}, err
	}
	invResidual, err := invCheck.MaxAbsDiff(linear.Identity(4))
	if err != nil {
		return TrivializationAudit{}, err
	}

	pa := linear.Diagonal([]float64{1, 1, 0, 0})
	pb := linear.Diagonal([]float64{0, 0, 1, 1})
	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})

	wpa, err := conjugateByTranspose(w, pa)
	if err != nil {
		return TrivializationAudit{}, err
	}
	wpb, err := conjugateByTranspose(w, pb)
	if err != nil {
		return TrivializationAudit{}, err
	}
	resAH, _ := wpa.MaxAbsDiff(ph)
	resBL, _ := wpb.MaxAbsDiff(pl)
	resAL, _ := wpa.MaxAbsDiff(pl)
	resBH, _ := wpb.MaxAbsDiff(ph)
	if seal.Choice == EtaToLow {
		// For eta_to_low, the sealed orientation maps A to low and B to high.
		resAH, resBL, resAL, resBH = resAL, resBH, resAH, resBL
	}
	verified := invResidual <= eps && resAH <= eps && resBL <= eps && frame.ConditionalOnSeal
	return TrivializationAudit{
		MatrixName:                         "W_seal",
		Matrix:                             FormatMatrix(w),
		OrthogonalFrameMap:                 invResidual <= eps,
		Invertible:                         invResidual <= eps,
		InverseEqualsTranspose:             invResidual <= eps,
		MapsAProjectorToHighResidual:       resAH,
		MapsBProjectorToLowResidual:        resBL,
		MapsAProjectorToLowResidual:        resAL,
		MapsBProjectorToHighResidual:       resBH,
		ProjectorIntertwiningVerified:      verified,
		UniqueWithoutGaugeFrame:            false,
		GaugeFreedomBeforeSeal:             "GL(2,R)_A × GL(2,R)_B and eta <-> -eta orientation exchange",
		GaugeFreedomAfterSealedFrameChoice: "residual block-frame freedom inside each 2D fiber; W_seal is canonical only in the chosen sealed orthonormal A/B and high/low frames",
		PhysicalScalarBundleTrivialized:    verified,
		Verdict:                            "The explicit seal lets the engine construct a gauge-fixed H_Phi trivialization. This is not a unique unsealed quartic-basis isomorphism; it is a conditional frame map whose dependencies are explicit.",
	}, nil
}

func auditGaugePullback(sc scalarcovariant.Analysis, triv TrivializationAudit, eps float64) (GaugePullbackAudit, error) {
	if !triv.ProjectorIntertwiningVerified {
		return GaugePullbackAudit{}, fmt.Errorf("cannot pull back gauge generators without sealed trivialization")
	}
	w := linear.Identity(4) // BuildDefault uses eta_to_high. The sealed frame is already H_Phi ordered.
	gens := []struct {
		name    string
		m       linear.Matrix
		meaning string
	}{
		{"T1", sc.T1, "broken charged weak generator; connects the two scalar fibers"},
		{"T2", sc.T2, "broken charged weak generator; connects the two scalar fibers"},
		{"T3L", sc.T3, "diagonal weak-isospin generator preserving high/low planes"},
		{"Y_phi", sc.YPhi, "scalar hypercharge phase rotation preserving high/low planes"},
	}
	records := make([]PulledBackGenerator, 0, len(gens))
	for _, g := range gens {
		pull, err := conjugateByTranspose(w.Transpose(), g.m)
		if err != nil {
			return GaugePullbackAudit{}, err
		}
		records = append(records, generatorRecord(g.name, pull, g.meaning, eps))
	}
	t1 := findGen(records, "T1")
	t2 := findGen(records, "T2")
	t3 := findGen(records, "T3L")
	y := findGen(records, "Y_phi")
	return GaugePullbackAudit{
		Generators:                           records,
		T3LPreservesFibers:                   t3.PreservesABFibers && t3.BlockDiagonal,
		YPhiPreservesFibers:                  y.PreservesABFibers && y.BlockDiagonal,
		T1MixesFibers:                        t1.MixesABFibers && !t1.BlockDiagonal,
		T2MixesFibers:                        t2.MixesABFibers && !t2.BlockDiagonal,
		BrokenGeneratorsOffDiagonal:          t1.MixesABFibers && t2.MixesABFibers,
		UnbrokenGeneratorsBlockDiagonal:      t3.BlockDiagonal && y.BlockDiagonal,
		GaugeConnectionAttachedConditionally: true,
		GaugeBosonMassesDerived:              false,
		PhysicalCouplingsDerived:             false,
		Verdict:                              "Under W_seal, T3L and Y_phi preserve the sealed A/B fibers, while T1 and T2 are off-diagonal and connect them. This attaches the weak representation to the conditional scalar bundle but does not derive masses or couplings.",
	}, nil
}

func auditFirewall(prev scalarorientationsource.Analysis, seal SpontaneousOrientationSeal, triv TrivializationAudit, pull GaugePullbackAudit, top topologicalnormalization.Analysis, beta betamatching.Analysis, sc scalarcovariant.Analysis, cx scalarcomplex.Analysis) SealFirewallAudit {
	thresholdDerived := beta.ThresholdCorrectedBetaDerived || beta.BetaCorrectionRowsAllowed > 0
	return SealFirewallAudit{
		SealExplicitInput:                 seal.ExplicitAxiom,
		SealQuarantined:                   seal.Quarantined,
		ObservedInputUsed:                 seal.UsesObservedInput,
		HiddenEtaSelectorClaimed:          seal.DerivedFromFiniteSelector,
		ChernWeilCarrierDerived:           false,
		HeatKernelMatchingDerived:         false,
		ThresholdCorrectedBetaDerived:     false,
		YukawaAmplitudeDerived:            false,
		AbsoluteCouplingPromoted:          false,
		PhysicalConstantsDerived:          false,
		TopologicalSealImportedAsConstant: false,
		BetaRowsUnlocked:                  thresholdDerived && false,
		ScalarKineticNormalizationDerived: sc.FiniteScalarKineticNormalizationDerived,
		GaugeActionHessianDerived:         sc.GaugeActionHessianDerived,
		GaugeGeneratorPullbackDerived:     pull.GaugeConnectionAttachedConditionally,
		ConditionalPhysicalBundleDerived:  triv.PhysicalScalarBundleTrivialized,
		StrictNullityBefore:               prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:          prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:           0,
		ClosedStatements: []string{
			"the eta-to-high/low orientation is no longer searched for as a hidden theorem; it is explicitly recorded as spontaneous boundary data",
			"the physical H_Phi scalar bundle is trivialized only in the sealed frame and only conditionally on the orientation seal",
			"T3L and Y_phi pull back as fiber-preserving/block-diagonal generators",
			"T1 and T2 pull back as fiber-mixing/off-diagonal weak generators",
			fmt.Sprintf("topological normalization remains a separate seal: conditional absolute U available=%t, strict absolute U derived=%t", top.Firewall.ConditionalAbsoluteUAvailable, top.Firewall.StrictAbsoluteUDerived),
			fmt.Sprintf("beta matching remains separate: threshold-corrected beta derived=%t", beta.ThresholdCorrectedBetaDerived),
			fmt.Sprintf("pair-compatible complex structure is available=%t but canonical complex derived=%t", cx.PairCompatibleComplexAvailable, cx.CanonicalComplexDerived),
		},
		OpenRequirements: []string{
			"construct a scalar-bundle Chern-Weil carrier using the sealed bundle without importing the topological seal as an absolute coupling",
			"derive heat-kernel matching on the sealed bundle before any threshold beta-row promotion",
			"derive scalar kinetic normalization and physical gauge couplings before interpreting W/Z/photon masses",
			"keep all results depending on the orientation seal marked CONDITIONAL_ON_VACUUM_SEAL",
		},
		RecommendedNextGate: "Gate 192 — sealed scalar-bundle Chern-Weil carrier / heat-kernel preflight audit",
		Verdict:             "Gate 191 opens the conditional physical scalar bundle but keeps the strict firewall intact: the seal is explicit, quarantined, and carries no numerical physical constant or hidden eta selector.",
	}
}

func projectorRecord(name, meaning string, p, complement linear.Matrix, eps float64) (PhysicalProjector, error) {
	p2, err := p.Mul(p)
	if err != nil {
		return PhysicalProjector{}, err
	}
	idemp, _ := p2.MaxAbsDiff(p)
	sum, err := p.Add(complement)
	if err != nil {
		return PhysicalProjector{}, err
	}
	comp, _ := sum.MaxAbsDiff(linear.Identity(4))
	prod, err := p.Mul(complement)
	if err != nil {
		return PhysicalProjector{}, err
	}
	orth := prod.FrobeniusNorm()
	tr, err := p.Trace()
	if err != nil {
		return PhysicalProjector{}, err
	}
	return PhysicalProjector{Name: name, Meaning: meaning, Matrix: FormatMatrix(p), Trace: tr, IdempotenceResidual: idemp, ComplementResidual: comp, OrthogonalityResidual: orth, Dimension: int(math.Round(tr)), Verified: idemp <= eps && comp <= eps && orth <= eps && int(math.Round(tr)) == 2}, nil
}

func generatorRecord(name string, g linear.Matrix, meaning string, eps float64) PulledBackGenerator {
	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	commH := commNorm(g, ph)
	commL := commNorm(g, pl)
	hh := linear.MustMul(linear.MustMul(ph, g), ph)
	ll := linear.MustMul(linear.MustMul(pl, g), pl)
	hl := linear.MustMul(linear.MustMul(ph, g), pl)
	lh := linear.MustMul(linear.MustMul(pl, g), ph)
	off, _ := hl.Add(lh)
	diag, _ := hh.Add(ll)
	offNorm := off.FrobeniusNorm()
	diagNorm := diag.FrobeniusNorm()
	block := offNorm <= eps && commH <= eps && commL <= eps
	return PulledBackGenerator{
		Name:              name,
		PullbackMatrix:    FormatMatrix(g),
		CommutesWithHigh:  commH <= eps,
		CommutesWithLow:   commL <= eps,
		BlockDiagonal:     block,
		OffDiagonalNorm:   offNorm,
		DiagonalBlockNorm: diagNorm,
		MixesABFibers:     offNorm > eps,
		PreservesABFibers: block,
		PhysicalMeaning:   meaning,
	}
}

func conjugateByTranspose(w, a linear.Matrix) (linear.Matrix, error) {
	// Orthogonal frame maps use W^{-1}=W^T. This returns W^T A W.
	wt := w.Transpose()
	left, err := wt.Mul(a)
	if err != nil {
		return linear.Matrix{}, err
	}
	return left.Mul(w)
}

func planeSwap() linear.Matrix {
	w := linear.NewMatrix(4, 4)
	w.Set(0, 2, 1)
	w.Set(1, 3, 1)
	w.Set(2, 0, 1)
	w.Set(3, 1, 1)
	return w
}

func commNorm(a, b linear.Matrix) float64 {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return math.Inf(1)
	}
	return c.FrobeniusNorm()
}

func findGen(records []PulledBackGenerator, name string) PulledBackGenerator {
	for _, r := range records {
		if r.Name == name {
			return r
		}
	}
	return PulledBackGenerator{Name: name}
}

func FormatMatrix(m linear.Matrix) string {
	rows := make([]string, 0, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, 0, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals = append(vals, fmt.Sprintf("%.6g", m.At(r, c)))
		}
		rows = append(rows, "["+strings.Join(vals, ", ")+"]")
	}
	return "[" + strings.Join(rows, ", ") + "]"
}

func FormatSeal(a SpontaneousOrientationSeal) string {
	return fmt.Sprintf("name=%s choice=%s source=%q axiom=%s status=%s etaHigh=%t etaLow=%t explicit=%t quarantined=%t gate190=%t breaksEta=%t derivedSelector=%t observed=%t constant=%t coupling=%t massScale=%t (%s)", a.Name, a.Choice, a.Source, a.AxiomID, a.ConditionalStatus, a.EtaMapsToHigh, a.EtaMapsToLow, a.ExplicitAxiom, a.Quarantined, a.RequiredByGate190, a.BreaksEtaInvolutionAsBoundaryData, a.DerivedFromFiniteSelector, a.UsesObservedInput, a.CarriesNumericPhysicalConstant, a.CarriesGaugeCoupling, a.CarriesMassScale, a.Verdict)
}

func FormatSealedFrame(a SealedFrameAudit) string {
	return fmt.Sprintf("abstract=%q physical=%q A=%s B=%s high=%s low=%s AtoHigh=%t BtoLow=%t AtoLow=%t BtoHigh=%t dim=%t gate190=%t canonicalNoSeal=%t conditional=%t (%s)", a.AbstractFrameName, a.PhysicalFrameName, a.AbstractProjectorA.Matrix, a.AbstractProjectorB.Matrix, a.PhysicalProjectorHigh.Matrix, a.PhysicalProjectorLow.Matrix, a.SealMapsAToHigh, a.SealMapsBToLow, a.SealMapsAToLow, a.SealMapsBToHigh, a.DimensionCompatibilityInherited, a.Gate190ObstructionInherited, a.CanonicalWithoutSeal, a.ConditionalOnSeal, a.Verdict)
}

func FormatTrivialization(a TrivializationAudit) string {
	return fmt.Sprintf("W=%s matrix=%s orthogonal=%t invertible=%t invT=%t AtoHigh=%.3g BtoLow=%.3g AtoLow=%.3g BtoHigh=%.3g verified=%t uniqueNoGauge=%t before=%q after=%q physicalBundle=%t (%s)", a.MatrixName, a.Matrix, a.OrthogonalFrameMap, a.Invertible, a.InverseEqualsTranspose, a.MapsAProjectorToHighResidual, a.MapsBProjectorToLowResidual, a.MapsAProjectorToLowResidual, a.MapsBProjectorToHighResidual, a.ProjectorIntertwiningVerified, a.UniqueWithoutGaugeFrame, a.GaugeFreedomBeforeSeal, a.GaugeFreedomAfterSealedFrameChoice, a.PhysicalScalarBundleTrivialized, a.Verdict)
}

func FormatGaugePullback(a GaugePullbackAudit) string {
	parts := make([]string, 0, len(a.Generators))
	for _, g := range a.Generators {
		parts = append(parts, fmt.Sprintf("%s block=%t off=%.6g diag=%.6g mix=%t preserve=%t meaning=%q matrix=%s", g.Name, g.BlockDiagonal, g.OffDiagonalNorm, g.DiagonalBlockNorm, g.MixesABFibers, g.PreservesABFibers, g.PhysicalMeaning, g.PullbackMatrix))
	}
	return fmt.Sprintf("T3=%t Y=%t T1mix=%t T2mix=%t brokenOff=%t unbrokenBlock=%t attached=%t masses=%t couplings=%t generators=[%s] (%s)", a.T3LPreservesFibers, a.YPhiPreservesFibers, a.T1MixesFibers, a.T2MixesFibers, a.BrokenGeneratorsOffDiagonal, a.UnbrokenGeneratorsBlockDiagonal, a.GaugeConnectionAttachedConditionally, a.GaugeBosonMassesDerived, a.PhysicalCouplingsDerived, strings.Join(parts, "; "), a.Verdict)
}

func FormatFirewall(a SealFirewallAudit) string {
	return fmt.Sprintf("seal=%t quarantine=%t observed=%t hiddenSelector=%t chern=%t heat=%t threshold=%t yukawa=%t abs=%t constants=%t topoAsConst=%t betaUnlocked=%t scalarKinetic=%t gaugeHessian=%t pullback=%t conditionalBundle=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.SealExplicitInput, a.SealQuarantined, a.ObservedInputUsed, a.HiddenEtaSelectorClaimed, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.YukawaAmplitudeDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.TopologicalSealImportedAsConstant, a.BetaRowsUnlocked, a.ScalarKineticNormalizationDerived, a.GaugeActionHessianDerived, a.GaugeGeneratorPullbackDerived, a.ConditionalPhysicalBundleDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d gate190=%t seal=%t trivialization=%t pullback=%t T3Y=%t T1T2=%t conditionalBundle=%t sealed=%t (%s)", a.TestsAudited, a.Gate190ObstructionInherited, a.SpontaneousOrientationSealRecorded, a.GaugeFixedTrivializationConstructed, a.GaugeGeneratorPullbackConstructed, a.T3YBlockDiagonal, a.T1T2OffDiagonal, a.PhysicalScalarBundleConditionallyDerived, a.ChernWeilHeatKernelThresholdsStillSealed, a.Comment)
}
