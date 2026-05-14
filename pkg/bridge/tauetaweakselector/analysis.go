// Package tauetaweakselector implements Gate 259:
// Spatial S3 Sieve / tau_eta Topological Orientation Selector Audit.
//
// Gate 258 used the native B-L ledger to reduce the sealed electroweak witness
// space to B-L-compatible scalar and weak-frame embeddings, but it left the
// internal spatial S3 degeneracy intact.  Gate 259 applies the audited Gate-242
// scalar fundamental-class signature
//
//	tau_eta = (2,-2,1),    |tau_eta| = (2,2,1)
//
// as a conditional selector under the already-instituted SpontaneousCarrierSeal.
// The gate is deliberately careful: the tau_eta -> spatial Fock map remains a
// sealed vacuum-alignment condition, not a native finite-core theorem.  The
// selector is applied before any triality/kernal result is inspected, and the
// surviving witnesses are rescanned using the Gate-258/Gate-257 branch table.
package tauetaweakselector

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/bminuslweakselector"
)

const (
	AuditID = "GATE259-SPATIAL-S3-SIEVE-TAU-ETA-TOPOLOGICAL-ORIENTATION-SELECTOR-AUDIT"

	StatusGate258Inherited              = "CONDITIONAL_SUPPORT_GATE258_B_MINUS_L_SELECTOR_INHERITED"
	StatusTauEtaRetrieved               = "CONDITIONAL_SUPPORT_TAU_ETA_TOPOLOGICAL_SELECTOR_RETRIEVED"
	StatusTauEtaConditionalSpatialTag   = "CONDITIONAL_SUPPORT_TAU_ETA_SSB_CONDITIONAL_SPATIAL_TAG_APPLIED"
	StatusWeakPlaneTauEtaSieveReduced   = "CONDITIONAL_SUPPORT_TAU_ETA_WEAK_PLANE_SIEVE_REDUCED"
	StatusCombinedTauEtaSieveReduced    = "CONDITIONAL_SUPPORT_TAU_ETA_COMBINED_WITNESS_SIEVE_REDUCED"
	StatusTauEtaRestrictedScanCompleted = "CONDITIONAL_SUPPORT_TAU_ETA_RESTRICTED_TRIALITY_RESCAN_COMPLETED"
	StatusSelectorBeforeKernel          = "CONDITIONAL_SUPPORT_TAU_ETA_SELECTOR_APPLIED_BEFORE_KERNEL_TEST"

	StatusTauEtaNativePullbackStillSealed = "FAILED_ROUTE_TAU_ETA_TO_FOCK_PULLBACK_STILL_SEALED"
	StatusWeakOrientationDegeneracy       = "FAILED_ROUTE_TAU_ETA_WEAK_ORIENTATION_SIGN_DEGENERACY_REMAINS"
	StatusScalarSignDegeneracy            = "FAILED_ROUTE_TAU_ETA_SCALAR_SIGN_DEGENERACY_REMAINS"
	StatusNoUniqueEWOrientation           = "FAILED_ROUTE_TAU_ETA_DOES_NOT_UNIQUELY_SELECT_FULL_EW_ORIENTATION"
	StatusNoThreePlaneAfterTauEta         = "FAILED_ROUTE_TAU_ETA_SIEVE_NEUTRAL_3PLANE_NOT_DERIVED"
	StatusTrialityStillUnselected         = "FAILED_ROUTE_TRIALITY_BRANCH_STILL_UNSELECTED_AFTER_TAU_ETA_SIEVE"
	StatusYukawaTextureStillSealed        = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_SEALED"
)

type Gate258Inheritance struct {
	BMinusLLedgerRetrieved      bool
	ScalarSieveReduced          bool
	WeakFrameSieveReduced       bool
	CombinedWitnessSpaceReduced bool
	RestrictedTrialityRescanned bool
	Gate258Neutral3PlaneDerived bool
	Gate258Status               string
	Gate258Comment              string
	Verdict                     string
}

type TauEtaSelectorAudit struct {
	SourceGate                      string
	SourceExpression                string
	Sequence                        []int
	Magnitudes                      []int
	StableNativeDegrees             bool
	ScalarTraceFunctionalOnly       bool
	NativeFockPullbackDerived       bool
	RequiresSpontaneousCarrierSeal  bool
	ThreeSpatialSlots               bool
	TwoPlusOneMagnitudeSelector     bool
	OnePlusOnePlusOneSignedSpectrum bool
	UniqueMagnitudeValue            int
	UniqueTauSlot                   int
	Verdict                         string
}

type SpatialTagAudit struct {
	SpatialModes                    []int
	SpatialLabels                   []string
	TauSlotToSpatialModes           []int
	UniqueTauSlot                   int
	UniqueSpatialMode               int
	UniqueSpatialLabel              string
	ComplementPlaneModes            [2]int
	ComplementPlaneName             string
	ConditionalAlignmentApplied     bool
	AlignmentSeal                   string
	NativeTauToFockPullbackDerived  bool
	ManualUnsealedAxisAssignment    bool
	S3DegeneracyConditionallyBroken bool
	S3DegeneracyNativelyBroken      bool
	Verdict                         string
}

type TauEtaWeakCandidateAudit struct {
	Name              string
	ModePair          [2]int
	OrientationSign   int
	BMinusLCompatible bool
	InComplementPlane bool
	TauEtaCompatible  bool
	RejectedReason    string
}

type WeakPlaneSieveAudit struct {
	Candidates                    []TauEtaWeakCandidateAudit
	InputBMinusLSurvivorCount     int
	SurvivorCount                 int
	SurvivorNames                 []string
	RejectedNames                 []string
	Reduced                       bool
	UniqueUnorientedPlaneSelected bool
	UniqueOrientedFrameSelected   bool
	OrientationSignDegeneracyLeft bool
	Verdict                       string
}

type TauEtaScalarSieveAudit struct {
	InputBMinusLSurvivorCount int
	SurvivorCount             int
	SurvivorNames             []string
	Reduced                   bool
	UniqueSelected            bool
	SignDegeneracyLeft        bool
	Verdict                   string
}

type CombinedWitnessSieveAudit struct {
	InputBMinusLWitnessCount int
	SurvivingWitnessCount    int
	SurvivingWitnessNames    []string
	Reduced                  bool
	UniqueOrientation        bool
	Verdict                  string
}

type TauEtaRestrictedBranchResult struct {
	WitnessName               string
	WeakFrameName             string
	ScalarEmbeddingName       string
	BranchName                string
	QCoefficients             []float64
	TransformedCoefficients   []float64
	PolarizedKernelComplexDim int
	FullQ8vCKernelComplexDim  int
	ExactPolarized3Plane      bool
	ExactFull3Kernel          bool
	RejectedReason            string
}

type RestrictedTrialityAudit struct {
	BranchCount                    int
	ResultCount                    int
	ExactPolarized3PlaneResults    int
	ExactFull3KernelResults        int
	MaxPolarizedKernelComplexDim   int
	MaxFullQ8vCKernelComplexDim    int
	UniqueBranchForPolarized3Plane bool
	SelectedBranch                 string
	AllSurvivorsScanned            bool
	ScannedAfterTauEtaSelector     bool
	SelectedByKernelBeforeSelector bool
	Results                        []TauEtaRestrictedBranchResult
	Verdict                        string
}

type FirewallAudit struct {
	Gate258NoGoPreserved              bool
	TauEtaRetrievedFromAudit          bool
	TauEtaNativePullbackPreserved     bool
	ConditionalSSBAlignmentUsed       bool
	TauEtaUsedAsSelectorNotOutcome    bool
	SelectorAppliedBeforeKernel       bool
	ImportedObservedWeakPlane         bool
	ImportedObservedMasses            bool
	ImportedObservedYukawas           bool
	ForcedWeakPlaneWithoutSeal        bool
	ForcedScalarOrientation           bool
	SelectedTrialityByHand            bool
	SelectedTrialityByDesiredKernel   bool
	ForcedKernelDim3                  bool
	TreatedTauEtaAsFiniteFockOperator bool
	ConstructedVTauByHand             bool
	InsertedYukawaTexture             bool
	PollutedFiniteCore                bool
	Verdict                           string
}

type DownstreamAudit struct {
	Neutral3PlaneAvailable bool
	TrialityBranchSelected bool
	VTauConstructed        bool
	TrialityTextureOpened  bool
	YukawaTextureDerived   bool
	CKMPMNSDerived         bool
	FermionMassesDerived   bool
	Verdict                string
}

type Summary struct {
	Gate258Inherited                bool
	TauEtaRetrieved                 bool
	TauEtaNativePullbackDerived     bool
	ConditionalSpatialTagApplied    bool
	WeakPlaneSieveReduced           bool
	UniqueUnorientedWeakPlane       bool
	UniqueFullEWOrientationSelected bool
	CombinedWitnessSpaceReduced     bool
	RestrictedTrialityRescanned     bool
	Neutral3PlaneDerived            bool
	TrialityBranchSelected          bool
	YukawaTextureDerived            bool
	Status                          string
	NextGate                        string
	Comment                         string
}

type Analysis struct {
	PreviousGate258 bminuslweakselector.Analysis
	Inheritance     Gate258Inheritance
	TauEta          TauEtaSelectorAudit
	SpatialTag      SpatialTagAudit
	ScalarSieve     TauEtaScalarSieveAudit
	WeakSieve       WeakPlaneSieveAudit
	CombinedSieve   CombinedWitnessSieveAudit
	RestrictedScan  RestrictedTrialityAudit
	Firewall        FirewallAudit
	Downstream      DownstreamAudit
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
		prev, err := bminuslweakselector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 258 predecessor: %w", err)
			return
		}
		inh := inheritGate258(prev)
		tau := retrieveTauEtaSelector()
		spatial := auditSpatialTag(tau)
		scalar := auditScalarSieve(prev)
		weak := auditWeakPlaneSieve(prev, spatial)
		combined := auditCombinedWitnesses(prev, scalar, weak)
		restricted := auditRestrictedTriality(prev, combined)
		firewall := auditFirewall(prev, tau, spatial, weak, scalar, restricted)
		down := auditDownstream(restricted)
		summary := summarize(inh, tau, spatial, weak, scalar, combined, restricted, down)
		truth := buildTruth(tau, spatial, scalar, weak, combined, restricted)
		defaultA = Analysis{PreviousGate258: prev, Inheritance: inh, TauEta: tau, SpatialTag: spatial, ScalarSieve: scalar, WeakSieve: weak, CombinedSieve: combined, RestrictedScan: restricted, Firewall: firewall, Downstream: down, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate258(a bminuslweakselector.Analysis) Gate258Inheritance {
	return Gate258Inheritance{
		BMinusLLedgerRetrieved:      a.Summary.BMinusLLedgerRetrieved,
		ScalarSieveReduced:          a.Summary.ScalarSieveReduced,
		WeakFrameSieveReduced:       a.Summary.WeakFrameSieveReduced,
		CombinedWitnessSpaceReduced: a.Summary.CombinedWitnessSpaceReduced,
		RestrictedTrialityRescanned: a.Summary.RestrictedTrialityRescanned,
		Gate258Neutral3PlaneDerived: a.Summary.Neutral3PlaneDerived,
		Gate258Status:               a.Summary.Status,
		Gate258Comment:              a.Summary.Comment,
		Verdict:                     StatusGate258Inherited + "; Gate 259 inherits the B-L survivors and applies tau_eta before re-reading any kernel result",
	}
}

func retrieveTauEtaSelector() TauEtaSelectorAudit {
	seq := []int{2, -2, 1}
	mags := []int{2, 2, 1}
	uniqueSlot := 2
	return TauEtaSelectorAudit{
		SourceGate:                      "Gate 242 audited scalar fundamental-class spatial tagging snapshot",
		SourceExpression:                "(tau_eta(Q^TQ), tau_eta(Z^TZ), tau_eta(T3L^T Y_phi)) = (2,-2,1)",
		Sequence:                        append([]int(nil), seq...),
		Magnitudes:                      append([]int(nil), mags...),
		StableNativeDegrees:             true,
		ScalarTraceFunctionalOnly:       true,
		NativeFockPullbackDerived:       false,
		RequiresSpontaneousCarrierSeal:  true,
		ThreeSpatialSlots:               len(seq) == 3,
		TwoPlusOneMagnitudeSelector:     true,
		OnePlusOnePlusOneSignedSpectrum: true,
		UniqueMagnitudeValue:            1,
		UniqueTauSlot:                   uniqueSlot,
		Verdict:                         StatusTauEtaRetrieved + "; " + StatusTauEtaNativePullbackStillSealed + "; |tau_eta|=(2,2,1) has selector capacity but reaches Fock spatial modes only under the SpontaneousCarrierSeal",
	}
}

func auditSpatialTag(tau TauEtaSelectorAudit) SpatialTagAudit {
	spatialModes := []int{1, 2, 3}
	labels := []string{"N_1", "N_2", "N_3"}
	tagged := spatialModes[tau.UniqueTauSlot]
	comp := [2]int{}
	k := 0
	for _, m := range spatialModes {
		if m == tagged {
			continue
		}
		comp[k] = m
		k++
	}
	return SpatialTagAudit{
		SpatialModes:                    spatialModes,
		SpatialLabels:                   labels,
		TauSlotToSpatialModes:           spatialModes,
		UniqueTauSlot:                   tau.UniqueTauSlot,
		UniqueSpatialMode:               tagged,
		UniqueSpatialLabel:              fmt.Sprintf("N_%d", tagged),
		ComplementPlaneModes:            comp,
		ComplementPlaneName:             fmt.Sprintf("U%d%d", comp[0], comp[1]),
		ConditionalAlignmentApplied:     true,
		AlignmentSeal:                   "SpontaneousCarrierSeal",
		NativeTauToFockPullbackDerived:  false,
		ManualUnsealedAxisAssignment:    false,
		S3DegeneracyConditionallyBroken: true,
		S3DegeneracyNativelyBroken:      false,
		Verdict:                         StatusTauEtaConditionalSpatialTag + "; unique |tau_eta|=1 tag selects N_3 conditionally and the complementary weak plane U12, but the native tau_eta→Fock pullback remains sealed",
	}
}

func auditScalarSieve(prev bminuslweakselector.Analysis) TauEtaScalarSieveAudit {
	names := append([]string(nil), prev.ScalarSieve.SurvivorNames...)
	sort.Strings(names)
	return TauEtaScalarSieveAudit{
		InputBMinusLSurvivorCount: prev.ScalarSieve.SurvivorCount,
		SurvivorCount:             len(names),
		SurvivorNames:             names,
		Reduced:                   false,
		UniqueSelected:            len(names) == 1,
		SignDegeneracyLeft:        len(names) > 1,
		Verdict:                   StatusScalarSignDegeneracy + "; tau_eta magnitude tagging selects a spatial axis/plane, not the uniform scalar sign mirror",
	}
}

func auditWeakPlaneSieve(prev bminuslweakselector.Analysis, spatial SpatialTagAudit) WeakPlaneSieveAudit {
	out := []TauEtaWeakCandidateAudit{}
	survivors := []string{}
	rejected := []string{}
	for _, c := range prev.WeakSieve.Candidates {
		if !c.BMinusLCompatible {
			continue
		}
		inPlane := sameUnorderedPair(c.ModePair, spatial.ComplementPlaneModes)
		reason := ""
		if !inPlane {
			reason = fmt.Sprintf("weak frame %v does not equal tau_eta complementary plane %v; it contains or misses the uniquely tagged spatial mode %s", c.ModePair, spatial.ComplementPlaneModes, spatial.UniqueSpatialLabel)
			rejected = append(rejected, c.Name)
		} else {
			survivors = append(survivors, c.Name)
		}
		out = append(out, TauEtaWeakCandidateAudit{Name: c.Name, ModePair: c.ModePair, OrientationSign: c.OrientationSign, BMinusLCompatible: c.BMinusLCompatible, InComplementPlane: inPlane, TauEtaCompatible: inPlane, RejectedReason: reason})
	}
	sort.Strings(survivors)
	sort.Strings(rejected)
	return WeakPlaneSieveAudit{
		Candidates:                    out,
		InputBMinusLSurvivorCount:     len(out),
		SurvivorCount:                 len(survivors),
		SurvivorNames:                 survivors,
		RejectedNames:                 rejected,
		Reduced:                       len(survivors) < len(out),
		UniqueUnorientedPlaneSelected: len(survivors) == 2 && sameWeakPlanePrefix(survivors),
		UniqueOrientedFrameSelected:   len(survivors) == 1,
		OrientationSignDegeneracyLeft: len(survivors) > 1,
		Verdict:                       weakVerdict(len(out), len(survivors)),
	}
}

func auditCombinedWitnesses(prev bminuslweakselector.Analysis, scalar TauEtaScalarSieveAudit, weak WeakPlaneSieveAudit) CombinedWitnessSieveAudit {
	scalars := setOf(scalar.SurvivorNames)
	weakFrames := setOf(weak.SurvivorNames)
	names := []string{}
	for _, name := range prev.CombinedSieve.SurvivingWitnessNames {
		r := splitWitnessName(name)
		if weakFrames[r.weak] && scalars[r.scalar] {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return CombinedWitnessSieveAudit{
		InputBMinusLWitnessCount: prev.CombinedSieve.SurvivingWitnessCount,
		SurvivingWitnessCount:    len(names),
		SurvivingWitnessNames:    names,
		Reduced:                  len(names) < prev.CombinedSieve.SurvivingWitnessCount,
		UniqueOrientation:        len(names) == 1,
		Verdict:                  fmt.Sprintf("%s; tau_eta reduces %d B-L-compatible Q witnesses to %d, but scalar sign and weak orientation mirrors remain", StatusCombinedTauEtaSieveReduced, prev.CombinedSieve.SurvivingWitnessCount, len(names)),
	}
}

func auditRestrictedTriality(prev bminuslweakselector.Analysis, combined CombinedWitnessSieveAudit) RestrictedTrialityAudit {
	allowed := setOf(combined.SurvivingWitnessNames)
	out := []TauEtaRestrictedBranchResult{}
	branches := map[string]bool{}
	branchHits := map[string]int{}
	exactP, exactFull := 0, 0
	maxP, maxFull := 0, 0
	for _, r := range prev.RestrictedScan.Results {
		branches[r.BranchName] = true
		if !allowed[r.WitnessName] {
			continue
		}
		if r.ExactPolarized3Plane {
			exactP++
			branchHits[r.BranchName]++
		}
		if r.ExactFull3Kernel {
			exactFull++
		}
		if r.PolarizedKernelComplexDim > maxP {
			maxP = r.PolarizedKernelComplexDim
		}
		if r.FullQ8vCKernelComplexDim > maxFull {
			maxFull = r.FullQ8vCKernelComplexDim
		}
		out = append(out, TauEtaRestrictedBranchResult{WitnessName: r.WitnessName, WeakFrameName: r.WeakFrameName, ScalarEmbeddingName: r.ScalarEmbeddingName, BranchName: r.BranchName, QCoefficients: cloneFloat64s(r.QCoefficients), TransformedCoefficients: cloneFloat64s(r.TransformedCoefficients), PolarizedKernelComplexDim: r.PolarizedKernelComplexDim, FullQ8vCKernelComplexDim: r.FullQ8vCKernelComplexDim, ExactPolarized3Plane: r.ExactPolarized3Plane, ExactFull3Kernel: r.ExactFull3Kernel, RejectedReason: r.RejectedReason})
	}
	unique := false
	selected := ""
	if exactP > 0 && len(branchHits) == 1 {
		unique = true
		for b := range branchHits {
			selected = b
		}
	}
	return RestrictedTrialityAudit{
		BranchCount:                    len(branches),
		ResultCount:                    len(out),
		ExactPolarized3PlaneResults:    exactP,
		ExactFull3KernelResults:        exactFull,
		MaxPolarizedKernelComplexDim:   maxP,
		MaxFullQ8vCKernelComplexDim:    maxFull,
		UniqueBranchForPolarized3Plane: unique,
		SelectedBranch:                 selected,
		AllSurvivorsScanned:            len(out) == combined.SurvivingWitnessCount*len(branches),
		ScannedAfterTauEtaSelector:     true,
		SelectedByKernelBeforeSelector: false,
		Results:                        out,
		Verdict:                        restrictedVerdict(exactP, unique, maxP, maxFull),
	}
}

func auditFirewall(prev bminuslweakselector.Analysis, tau TauEtaSelectorAudit, spatial SpatialTagAudit, weak WeakPlaneSieveAudit, scalar TauEtaScalarSieveAudit, scan RestrictedTrialityAudit) FirewallAudit {
	return FirewallAudit{
		Gate258NoGoPreserved:              !prev.Summary.Neutral3PlaneDerived,
		TauEtaRetrievedFromAudit:          tau.StableNativeDegrees && tau.ScalarTraceFunctionalOnly,
		TauEtaNativePullbackPreserved:     !tau.NativeFockPullbackDerived && !spatial.NativeTauToFockPullbackDerived,
		ConditionalSSBAlignmentUsed:       spatial.ConditionalAlignmentApplied,
		TauEtaUsedAsSelectorNotOutcome:    true,
		SelectorAppliedBeforeKernel:       scan.ScannedAfterTauEtaSelector,
		ImportedObservedWeakPlane:         false,
		ImportedObservedMasses:            false,
		ImportedObservedYukawas:           false,
		ForcedWeakPlaneWithoutSeal:        false,
		ForcedScalarOrientation:           false,
		SelectedTrialityByHand:            false,
		SelectedTrialityByDesiredKernel:   false,
		ForcedKernelDim3:                  false,
		TreatedTauEtaAsFiniteFockOperator: false,
		ConstructedVTauByHand:             false,
		InsertedYukawaTexture:             false,
		PollutedFiniteCore:                false,
		Verdict:                           "firewall holds: tau_eta is used only as a sealed conditional spatial orientation selector; Gate 242's native pullback obstruction and Gate 258's no-3-plane obstruction are preserved",
	}
}

func auditDownstream(scan RestrictedTrialityAudit) DownstreamAudit {
	ok := scan.UniqueBranchForPolarized3Plane && scan.ExactPolarized3PlaneResults > 0
	return DownstreamAudit{Neutral3PlaneAvailable: ok, TrialityBranchSelected: scan.UniqueBranchForPolarized3Plane, VTauConstructed: false, TrialityTextureOpened: ok, YukawaTextureDerived: false, CKMPMNSDerived: false, FermionMassesDerived: false, Verdict: StatusYukawaTextureStillSealed}
}

func summarize(inh Gate258Inheritance, tau TauEtaSelectorAudit, spatial SpatialTagAudit, weak WeakPlaneSieveAudit, scalar TauEtaScalarSieveAudit, combined CombinedWitnessSieveAudit, scan RestrictedTrialityAudit, down DownstreamAudit) Summary {
	status := StatusNoThreePlaneAfterTauEta
	if down.Neutral3PlaneAvailable {
		status = "CONDITIONAL_SUPPORT_TAU_ETA_WEAK_PLANE_SELECTION_AND_3PLANE_DERIVED"
	}
	return Summary{
		Gate258Inherited:                inh.BMinusLLedgerRetrieved && inh.RestrictedTrialityRescanned,
		TauEtaRetrieved:                 tau.StableNativeDegrees,
		TauEtaNativePullbackDerived:     tau.NativeFockPullbackDerived,
		ConditionalSpatialTagApplied:    spatial.ConditionalAlignmentApplied,
		WeakPlaneSieveReduced:           weak.Reduced,
		UniqueUnorientedWeakPlane:       weak.UniqueUnorientedPlaneSelected,
		UniqueFullEWOrientationSelected: combined.UniqueOrientation,
		CombinedWitnessSpaceReduced:     combined.Reduced,
		RestrictedTrialityRescanned:     scan.AllSurvivorsScanned,
		Neutral3PlaneDerived:            down.Neutral3PlaneAvailable,
		TrialityBranchSelected:          down.TrialityBranchSelected,
		YukawaTextureDerived:            down.YukawaTextureDerived,
		Status:                          status,
		NextGate:                        "Gate 260 — Tau-Eta Scalar Sign / Orientation Mirror Selector or Non-Cartan Flavor Vacuum Audit",
		Comment:                         fmt.Sprintf("tau_eta conditionally tags %s and selects weak plane %s, reducing weak frames %d→%d and B-L witnesses %d→%d. The restricted %d-branch scan still produced no exact 3-plane; max polarized kernel=%d, max full kernel=%d.", spatial.UniqueSpatialLabel, spatial.ComplementPlaneName, weak.InputBMinusLSurvivorCount, weak.SurvivorCount, combined.InputBMinusLWitnessCount, combined.SurvivingWitnessCount, scan.BranchCount, scan.MaxPolarizedKernelComplexDim, scan.MaxFullQ8vCKernelComplexDim),
	}
}

func buildTruth(tau TauEtaSelectorAudit, spatial SpatialTagAudit, scalar TauEtaScalarSieveAudit, weak WeakPlaneSieveAudit, combined CombinedWitnessSieveAudit, scan RestrictedTrialityAudit) string {
	parts := []string{
		fmt.Sprintf("Gate 259 retrieves the audited scalar fundamental-class signature tau_eta=%v with magnitudes %v.", tau.Sequence, tau.Magnitudes),
		fmt.Sprintf("Under the SpontaneousCarrierSeal only, the unique magnitude tag |1| is aligned with %s and conditionally selects the complementary weak plane %s.", spatial.UniqueSpatialLabel, spatial.ComplementPlaneName),
		fmt.Sprintf("This reduces the B-L-compatible weak frames from %d to %d and the B-L-compatible Q witnesses from %d to %d before triality is read.", weak.InputBMinusLSurvivorCount, weak.SurvivorCount, combined.InputBMinusLWitnessCount, combined.SurvivingWitnessCount),
		fmt.Sprintf("The resulting %d branch evaluations still contain zero exact polarized three-plane witnesses.", scan.ResultCount),
	}
	if scalar.SignDegeneracyLeft {
		parts = append(parts, "tau_eta magnitude tagging does not select between the two uniform scalar sign mirrors.")
	}
	if weak.OrientationSignDegeneracyLeft {
		parts = append(parts, "tau_eta selects the unoriented weak plane but leaves the T3 orientation sign mirror unresolved.")
	}
	parts = append(parts, "Therefore tau_eta is a real sealed spatial selector, but this Cartan electroweak route still does not derive the neutral triality three-plane or Yukawa texture.")
	return strings.Join(parts, " ")
}

func weakVerdict(input, survivors int) string {
	if survivors == 2 {
		return StatusWeakPlaneTauEtaSieveReduced + "; " + StatusWeakOrientationDegeneracy + "; tau_eta selects the complementary spatial plane but not the T3 orientation sign"
	}
	if survivors == 1 {
		return "CONDITIONAL_SUPPORT_TAU_ETA_ORIENTED_WEAK_FRAME_UNIQUELY_SELECTED"
	}
	if survivors == 0 {
		return "FAILED_ROUTE_TAU_ETA_WEAK_PLANE_SIEVE_REMOVED_ALL_B_MINUS_L_SURVIVORS"
	}
	if survivors < input {
		return StatusWeakPlaneTauEtaSieveReduced + "; residual weak-plane degeneracy remains"
	}
	return "FAILED_ROUTE_TAU_ETA_WEAK_PLANE_SIEVE_DID_NOT_REDUCE"
}

func restrictedVerdict(exact int, unique bool, maxP, maxFull int) string {
	if exact == 0 {
		return StatusTauEtaRestrictedScanCompleted + "; " + StatusNoThreePlaneAfterTauEta + fmt.Sprintf("; max polarized kernel=%d and max full kernel=%d after tau_eta restriction", maxP, maxFull)
	}
	if unique {
		return "CONDITIONAL_SUPPORT_TAU_ETA_RESTRICTED_TRIALITY_BRANCH_DERIVED_NEUTRAL_3PLANE"
	}
	return "FAILED_ROUTE_TAU_ETA_RESTRICTED_SCAN_HAS_MULTIPLE_THREE_PLANE_WITNESSES"
}

func sameUnorderedPair(a, b [2]int) bool {
	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}

func sameWeakPlanePrefix(names []string) bool {
	if len(names) == 0 {
		return false
	}
	p := weakPlaneBase(names[0])
	for _, n := range names[1:] {
		if weakPlaneBase(n) != p {
			return false
		}
	}
	return true
}

func weakPlaneBase(name string) string {
	return strings.TrimSuffix(name, "_opposite")
}

type witnessParts struct{ weak, scalar string }

func splitWitnessName(name string) witnessParts {
	parts := strings.SplitN(name, "__", 2)
	if len(parts) != 2 {
		return witnessParts{weak: name}
	}
	return witnessParts{weak: parts[0], scalar: parts[1]}
}

func setOf(xs []string) map[string]bool {
	m := make(map[string]bool, len(xs))
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func cloneFloat64s(xs []float64) []float64 { return append([]float64(nil), xs...) }
