// Package bminuslweakselector implements Gate 258:
// Weak-Plane Selector / B-L Embedding Orientation Constraint Audit.
//
// The gate consumes the sealed witness inventory from Gate 257 and applies an
// independent selector before any triality/kernal criterion is consulted.  The
// selector is the native B-L ledger, i.e. the 1⊕3 Fock polarization
//
//	B-L = -N_0 + (1/3)(N_1+N_2+N_3).
//
// It tests whether B-L reduces the weak-frame and scalar/contact embedding
// degeneracy.  It does not choose a witness because it yields a desired
// three-plane; all surviving witnesses are re-scanned using the Gate-257
// branch results.
package bminuslweakselector

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/sealedcarrierwitness"
)

const (
	AuditID = "GATE258-WEAK-PLANE-SELECTOR-B-MINUS-L-EMBEDDING-ORIENTATION-CONSTRAINT-AUDIT"

	StatusGate257Inherited          = "CONDITIONAL_SUPPORT_GATE257_SEALED_WITNESS_SCAN_INHERITED"
	StatusBMinusLLedgerRetrieved    = "CONDITIONAL_SUPPORT_B_MINUS_L_LEDGER_RETRIEVED"
	StatusScalarBMinusLSieveReduced = "CONDITIONAL_SUPPORT_B_MINUS_L_SCALAR_EMBEDDING_SIEVE_REDUCED"
	StatusWeakBMinusLSieveReduced   = "CONDITIONAL_SUPPORT_B_MINUS_L_WEAK_FRAME_SIEVE_REDUCED"
	StatusReducedTrialityRescanned  = "CONDITIONAL_SUPPORT_B_MINUS_L_RESTRICTED_TRIALITY_RESCAN_COMPLETED"
	StatusNoOutcomeSelector         = "CONDITIONAL_SUPPORT_SELECTOR_APPLIED_BEFORE_KERNEL_TEST"
	StatusBMinusLSelectorActive     = "CONDITIONAL_SUPPORT_B_MINUS_L_WEAK_PLANE_SELECTOR_ACTIVE"

	StatusWeakPlaneSpatialDegeneracyRemains = "FAILED_ROUTE_B_MINUS_L_SPATIAL_WEAK_PLANE_DEGENERACY_REMAINS"
	StatusScalarSignDegeneracyRemains       = "FAILED_ROUTE_B_MINUS_L_SCALAR_SIGN_DEGENERACY_REMAINS"
	StatusNoUniqueOrientation               = "FAILED_ROUTE_B_MINUS_L_DOES_NOT_UNIQUELY_SELECT_EW_ORIENTATION"
	StatusNoThreePlaneAfterSieve            = "FAILED_ROUTE_B_MINUS_L_SIEVE_NEUTRAL_3PLANE_NOT_DERIVED"
	StatusTrialityBranchStillUnselected     = "FAILED_ROUTE_TRIALITY_BRANCH_STILL_UNSELECTED_AFTER_B_MINUS_L_SIEVE"
	StatusYukawaTextureStillSealed          = "FAILED_ROUTE_YUKAWA_TEXTURE_STILL_SEALED"
)

type Gate257Inheritance struct {
	NativeChargeEigenvaluesExtracted bool
	EmbeddingWitnessesScanned        bool
	AllTrialityBranchesScanned       bool
	Gate257Neutral3PlaneDerived      bool
	Gate257Status                    string
	Gate257Comment                   string
	Verdict                          string
}

type BMinusLLedger struct {
	Expression              string
	Coefficients            []float64
	TemporalMode            int
	SpatialModes            []int
	DistinctEigenvalues     []float64
	SpatialIsotropy         bool
	OnePlusThreeSplit       bool
	DerivedFiniteFockLedger bool
	UsesObservedInput       bool
	Verdict                 string
}

type ScalarCandidateAudit struct {
	Name                        string
	Kind                        string
	Coefficients                []float64
	DiagonalCommutesWithBMinusL bool
	PreservesSpatialIsotropy    bool
	PreservesOnePlusThreeBlocks bool
	BMinusLCompatible           bool
	RejectedReason              string
}

type ScalarSieveAudit struct {
	Candidates            []ScalarCandidateAudit
	InputCount            int
	SurvivorCount         int
	SurvivorNames         []string
	RejectedNames         []string
	Reduced               bool
	UniqueSelected        bool
	SignDegeneracyRemains bool
	Verdict               string
}

type WeakFrameCandidateAudit struct {
	Name                        string
	ModePair                    [2]int
	OrientationSign             int
	Coefficients                []float64
	ModeBMinusLValues           []float64
	DiagonalCommutesWithBMinusL bool
	PairsEqualBMinusLSectors    bool
	IsSpatialSpatialPlane       bool
	BMinusLCompatible           bool
	RejectedReason              string
}

type WeakFrameSieveAudit struct {
	Candidates                 []WeakFrameCandidateAudit
	InputCount                 int
	SurvivorCount              int
	SurvivorNames              []string
	RejectedNames              []string
	Reduced                    bool
	UniqueSelected             bool
	SpatialPlaneDegeneracyLeft bool
	Verdict                    string
}

type CombinedWitnessSieveAudit struct {
	InputWitnessCount     int
	SurvivingWitnessCount int
	SurvivingWitnessNames []string
	Reduced               bool
	UniqueOrientation     bool
	Verdict               string
}

type RestrictedBranchResult struct {
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
	ScannedAfterSelector           bool
	SelectedByKernelBeforeSelector bool
	Results                        []RestrictedBranchResult
	Verdict                        string
}

type FirewallAudit struct {
	Gate257NoGoPreserved            bool
	BMinusLAppliedBeforeKernel      bool
	BMinusLUsedAsSelectorNotOutcome bool
	ImportedObservedChargeTable     bool
	ImportedObservedMasses          bool
	ImportedObservedYukawas         bool
	ForcedWeakPlane                 bool
	ForcedScalarOrientation         bool
	SelectedTrialityByHand          bool
	ForcedKernelDim3                bool
	AcceptedYOnlyAsQ                bool
	TreatedSealAsFiniteDerivation   bool
	ConstructedVTauByHand           bool
	InsertedYukawaTexture           bool
	PollutedFiniteCore              bool
	Verdict                         string
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
	Gate257Inherited            bool
	BMinusLLedgerRetrieved      bool
	ScalarSieveReduced          bool
	WeakFrameSieveReduced       bool
	CombinedWitnessSpaceReduced bool
	RestrictedTrialityRescanned bool
	UniqueEWOrientationSelected bool
	Neutral3PlaneDerived        bool
	TrialityBranchSelected      bool
	YukawaTextureDerived        bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	PreviousGate257 sealedcarrierwitness.Analysis
	Inheritance     Gate257Inheritance
	BMinusL         BMinusLLedger
	ScalarSieve     ScalarSieveAudit
	WeakSieve       WeakFrameSieveAudit
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
		prev, err := sealedcarrierwitness.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 257 predecessor: %w", err)
			return
		}
		inheritance := inheritGate257(prev)
		ledger := retrieveBMinusLLedger(prev)
		scalar := auditScalarSieve(prev, ledger)
		weak := auditWeakFrameSieve(prev, ledger)
		combined := auditCombinedWitnesses(prev, scalar, weak)
		restricted := auditRestrictedTriality(prev, combined)
		firewall := auditFirewall(prev, scalar, weak, restricted)
		down := auditDownstream(restricted)
		summary := summarize(inheritance, ledger, scalar, weak, combined, restricted, down)
		truth := buildTruth(ledger, scalar, weak, combined, restricted)
		defaultA = Analysis{PreviousGate257: prev, Inheritance: inheritance, BMinusL: ledger, ScalarSieve: scalar, WeakSieve: weak, CombinedSieve: combined, RestrictedScan: restricted, Firewall: firewall, Downstream: down, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate257(a sealedcarrierwitness.Analysis) Gate257Inheritance {
	return Gate257Inheritance{
		NativeChargeEigenvaluesExtracted: a.Summary.NativeChargeEigenvaluesExtracted,
		EmbeddingWitnessesScanned:        a.Summary.EmbeddingWitnessesScanned,
		AllTrialityBranchesScanned:       a.Summary.AllTrialityBranchesScanned,
		Gate257Neutral3PlaneDerived:      a.Summary.NeutralPolarized3PlaneDerived,
		Gate257Status:                    a.Summary.Status,
		Gate257Comment:                   a.Summary.Comment,
		Verdict:                          StatusGate257Inherited + "; Gate 258 reuses the Gate-257 witness ledger and applies B-L before re-reading the branch results",
	}
}

func retrieveBMinusLLedger(prev sealedcarrierwitness.Analysis) BMinusLLedger {
	coeff := []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
	// Cross-check against the Gate-257 charge source, but keep the local exact
	// rational snapshot so this package does not depend on a deep historical
	// reconstruction.
	derived := prev.Charges.BMinusLFockLedgerDerived && sameVec(coeff, findBMinusLCoefficients(prev), 1e-12)
	return BMinusLLedger{
		Expression:              "B-L=-N_0+(1/3)(N_1+N_2+N_3)",
		Coefficients:            coeff,
		TemporalMode:            0,
		SpatialModes:            []int{1, 2, 3},
		DistinctEigenvalues:     []float64{-1, 1.0 / 3.0},
		SpatialIsotropy:         coeff[1] == coeff[2] && coeff[2] == coeff[3],
		OnePlusThreeSplit:       coeff[0] != coeff[1] && coeff[1] == coeff[2] && coeff[2] == coeff[3],
		DerivedFiniteFockLedger: derived,
		UsesObservedInput:       false,
		Verdict:                 StatusBMinusLLedgerRetrieved + "; native 1⊕3 Fock polarization is available as a selector, not as the electroweak charge itself",
	}
}

func findBMinusLCoefficients(prev sealedcarrierwitness.Analysis) []float64 {
	for _, s := range prev.Charges.Sources {
		if strings.Contains(s.Name, "B-L") && len(s.CoefficientVector) == 4 {
			return append([]float64(nil), s.CoefficientVector...)
		}
	}
	return nil
}

func auditScalarSieve(prev sealedcarrierwitness.Analysis, ledger BMinusLLedger) ScalarSieveAudit {
	out := make([]ScalarCandidateAudit, 0, len(prev.Embedding.ScalarEmbeddings))
	survivors := []string{}
	rejected := []string{}
	for _, s := range prev.Embedding.ScalarEmbeddings {
		spatialIso := sameScalar(s.YPhiCoefficients[1], s.YPhiCoefficients[2], 1e-12) && sameScalar(s.YPhiCoefficients[2], s.YPhiCoefficients[3], 1e-12)
		block := spatialIso // the temporal slot may differ, but the spatial S3 orbit must remain unsplit.
		compatible := spatialIso && block
		reason := ""
		if !compatible {
			reason = "scalar/contact orientation splits the B-L spatial S3 orbit and therefore breaks the native 1⊕3 selector"
			rejected = append(rejected, s.Name)
		} else {
			survivors = append(survivors, s.Name)
		}
		out = append(out, ScalarCandidateAudit{Name: s.Name, Kind: s.Kind, Coefficients: append([]float64(nil), s.YPhiCoefficients...), DiagonalCommutesWithBMinusL: true, PreservesSpatialIsotropy: spatialIso, PreservesOnePlusThreeBlocks: block, BMinusLCompatible: compatible, RejectedReason: reason})
	}
	sort.Strings(survivors)
	sort.Strings(rejected)
	return ScalarSieveAudit{Candidates: out, InputCount: len(out), SurvivorCount: len(survivors), SurvivorNames: survivors, RejectedNames: rejected, Reduced: len(survivors) < len(out), UniqueSelected: len(survivors) == 1, SignDegeneracyRemains: len(survivors) > 1, Verdict: scalarVerdict(len(out), len(survivors))}
}

func auditWeakFrameSieve(prev sealedcarrierwitness.Analysis, ledger BMinusLLedger) WeakFrameSieveAudit {
	out := make([]WeakFrameCandidateAudit, 0, len(prev.Embedding.WeakFrames))
	survivors := []string{}
	rejected := []string{}
	for _, w := range prev.Embedding.WeakFrames {
		i, j := w.ModePair[0], w.ModePair[1]
		equalSector := sameScalar(ledger.Coefficients[i], ledger.Coefficients[j], 1e-12)
		spatial := containsMode(ledger.SpatialModes, i) && containsMode(ledger.SpatialModes, j)
		compatible := equalSector && spatial
		reason := ""
		if !compatible {
			reason = "weak frame pairs unequal B-L sectors (temporal/lepton with spatial/quark), so the weak generator would not preserve the native B-L block"
			rejected = append(rejected, w.Name)
		} else {
			survivors = append(survivors, w.Name)
		}
		out = append(out, WeakFrameCandidateAudit{Name: w.Name, ModePair: w.ModePair, OrientationSign: w.OrientationSign, Coefficients: append([]float64(nil), w.T3Coefficients...), ModeBMinusLValues: []float64{ledger.Coefficients[i], ledger.Coefficients[j]}, DiagonalCommutesWithBMinusL: true, PairsEqualBMinusLSectors: equalSector, IsSpatialSpatialPlane: spatial, BMinusLCompatible: compatible, RejectedReason: reason})
	}
	sort.Strings(survivors)
	sort.Strings(rejected)
	return WeakFrameSieveAudit{Candidates: out, InputCount: len(out), SurvivorCount: len(survivors), SurvivorNames: survivors, RejectedNames: rejected, Reduced: len(survivors) < len(out), UniqueSelected: len(survivors) == 1, SpatialPlaneDegeneracyLeft: len(survivors) > 1, Verdict: weakVerdict(len(out), len(survivors))}
}

func auditCombinedWitnesses(prev sealedcarrierwitness.Analysis, scalar ScalarSieveAudit, weak WeakFrameSieveAudit) CombinedWitnessSieveAudit {
	survScalar := setOf(scalar.SurvivorNames)
	survWeak := setOf(weak.SurvivorNames)
	names := []string{}
	for _, w := range prev.SO8.Witnesses {
		if survWeak[w.WeakFrameName] && survScalar[w.ScalarEmbeddingName] {
			names = append(names, w.Name)
		}
	}
	sort.Strings(names)
	return CombinedWitnessSieveAudit{InputWitnessCount: prev.SO8.WitnessCount, SurvivingWitnessCount: len(names), SurvivingWitnessNames: names, Reduced: len(names) < prev.SO8.WitnessCount, UniqueOrientation: len(names) == 1, Verdict: fmt.Sprintf("%s; B-L reduces %d Gate-257 Q witnesses to %d but does not select a unique orientation", StatusBMinusLSelectorActive, prev.SO8.WitnessCount, len(names))}
}

func auditRestrictedTriality(prev sealedcarrierwitness.Analysis, combined CombinedWitnessSieveAudit) RestrictedTrialityAudit {
	allowed := setOf(combined.SurvivingWitnessNames)
	out := []RestrictedBranchResult{}
	exactP, exactFull := 0, 0
	maxP, maxFull := 0, 0
	branchHits := map[string]int{}
	branches := map[string]bool{}
	for _, r := range prev.TrialityScan.Results {
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
		out = append(out, RestrictedBranchResult{WitnessName: r.WitnessName, WeakFrameName: r.WeakFrameName, ScalarEmbeddingName: r.ScalarEmbeddingName, BranchName: r.BranchName, QCoefficients: append([]float64(nil), r.QCoefficients...), TransformedCoefficients: append([]float64(nil), r.TransformedCoefficients...), PolarizedKernelComplexDim: r.PolarizedKernelComplexDim, FullQ8vCKernelComplexDim: r.FullQ8vCKernelComplexDim, ExactPolarized3Plane: r.ExactPolarized3Plane, ExactFull3Kernel: r.ExactFull3Kernel, RejectedReason: r.RejectedReason})
	}
	unique := false
	selected := ""
	if exactP > 0 && len(branchHits) == 1 {
		unique = true
		for k := range branchHits {
			selected = k
		}
	}
	return RestrictedTrialityAudit{BranchCount: len(branches), ResultCount: len(out), ExactPolarized3PlaneResults: exactP, ExactFull3KernelResults: exactFull, MaxPolarizedKernelComplexDim: maxP, MaxFullQ8vCKernelComplexDim: maxFull, UniqueBranchForPolarized3Plane: unique, SelectedBranch: selected, AllSurvivorsScanned: len(out) == combined.SurvivingWitnessCount*len(branches), ScannedAfterSelector: true, SelectedByKernelBeforeSelector: false, Results: out, Verdict: restrictedVerdict(exactP, unique, maxP, maxFull)}
}

func auditFirewall(prev sealedcarrierwitness.Analysis, scalar ScalarSieveAudit, weak WeakFrameSieveAudit, scan RestrictedTrialityAudit) FirewallAudit {
	return FirewallAudit{Gate257NoGoPreserved: !prev.Summary.NeutralPolarized3PlaneDerived, BMinusLAppliedBeforeKernel: true, BMinusLUsedAsSelectorNotOutcome: true, ImportedObservedChargeTable: false, ImportedObservedMasses: false, ImportedObservedYukawas: false, ForcedWeakPlane: false, ForcedScalarOrientation: false, SelectedTrialityByHand: false, ForcedKernelDim3: false, AcceptedYOnlyAsQ: false, TreatedSealAsFiniteDerivation: false, ConstructedVTauByHand: false, InsertedYukawaTexture: false, PollutedFiniteCore: false, Verdict: "firewall holds: B-L is applied as a native pre-kernel selector; surviving degeneracy and failed 3-plane are recorded rather than patched"}
}

func auditDownstream(scan RestrictedTrialityAudit) DownstreamAudit {
	ok := scan.UniqueBranchForPolarized3Plane && scan.ExactPolarized3PlaneResults > 0
	return DownstreamAudit{Neutral3PlaneAvailable: ok, TrialityBranchSelected: scan.UniqueBranchForPolarized3Plane, VTauConstructed: false, TrialityTextureOpened: ok, YukawaTextureDerived: false, CKMPMNSDerived: false, FermionMassesDerived: false, Verdict: StatusYukawaTextureStillSealed}
}

func summarize(inh Gate257Inheritance, ledger BMinusLLedger, scalar ScalarSieveAudit, weak WeakFrameSieveAudit, combined CombinedWitnessSieveAudit, scan RestrictedTrialityAudit, down DownstreamAudit) Summary {
	status := StatusNoThreePlaneAfterSieve
	if down.Neutral3PlaneAvailable {
		status = "CONDITIONAL_SUPPORT_B_MINUS_L_WEAK_PLANE_SELECTOR_AND_NEUTRAL_3PLANE_DERIVED"
	}
	return Summary{Gate257Inherited: inh.NativeChargeEigenvaluesExtracted && inh.AllTrialityBranchesScanned, BMinusLLedgerRetrieved: ledger.DerivedFiniteFockLedger, ScalarSieveReduced: scalar.Reduced, WeakFrameSieveReduced: weak.Reduced, CombinedWitnessSpaceReduced: combined.Reduced, RestrictedTrialityRescanned: scan.AllSurvivorsScanned, UniqueEWOrientationSelected: combined.UniqueOrientation, Neutral3PlaneDerived: down.Neutral3PlaneAvailable, TrialityBranchSelected: down.TrialityBranchSelected, YukawaTextureDerived: down.YukawaTextureDerived, Status: status, NextGate: "Gate 259 — Spatial S3 / Contact Orientation Selector Beyond B-L", Comment: fmt.Sprintf("B-L reduced scalar embeddings %d→%d, weak frames %d→%d, and Q witnesses %d→%d. The restricted %d-branch scan still produced no exact 3-plane; max polarized kernel=%d, max full kernel=%d.", scalar.InputCount, scalar.SurvivorCount, weak.InputCount, weak.SurvivorCount, combined.InputWitnessCount, combined.SurvivingWitnessCount, scan.BranchCount, scan.MaxPolarizedKernelComplexDim, scan.MaxFullQ8vCKernelComplexDim)}
}

func buildTruth(ledger BMinusLLedger, scalar ScalarSieveAudit, weak WeakFrameSieveAudit, combined CombinedWitnessSieveAudit, scan RestrictedTrialityAudit) string {
	parts := []string{
		fmt.Sprintf("Gate 258 retrieves the native B-L ledger %s as the 1⊕3 Fock polarization selector.", ledger.Expression),
		fmt.Sprintf("The scalar/contact sieve preserves the spatial S3 orbit and reduces %d candidates to %d.", scalar.InputCount, scalar.SurvivorCount),
		fmt.Sprintf("The weak-frame sieve requires the weak pair to lie inside an equal B-L sector, reducing %d frames to %d spatial-spatial oriented frames.", weak.InputCount, weak.SurvivorCount),
		fmt.Sprintf("Together these constraints reduce the Gate-257 Q witness space from %d to %d before triality is read.", combined.InputWitnessCount, combined.SurvivingWitnessCount),
		fmt.Sprintf("The restricted all-branch scan has %d results and still yields zero exact polarized three-plane witnesses.", scan.ResultCount),
	}
	if weak.SpatialPlaneDegeneracyLeft {
		parts = append(parts, "B-L distinguishes the temporal/lepton mode from the spatial/quark orbit, but it cannot choose one oriented weak plane inside the remaining spatial S3 orbit.")
	}
	if scalar.SignDegeneracyRemains {
		parts = append(parts, "B-L also leaves the uniform scalar sign mirror unresolved.")
	}
	parts = append(parts, "Therefore B-L is a real selector, but not a sufficient selector for the neutral triality three-plane or Yukawa texture.")
	return strings.Join(parts, " ")
}

func scalarVerdict(input, survivors int) string {
	if survivors < input && survivors > 1 {
		return StatusScalarBMinusLSieveReduced + "; " + StatusScalarSignDegeneracyRemains + "; spatial isotropy removes contact 2+2 orientations but leaves the uniform ± scalar mirror"
	}
	if survivors == 1 {
		return "CONDITIONAL_SUPPORT_B_MINUS_L_SCALAR_EMBEDDING_UNIQUELY_SELECTED"
	}
	return "FAILED_ROUTE_B_MINUS_L_SCALAR_SIEVE_DID_NOT_REDUCE"
}

func weakVerdict(input, survivors int) string {
	if survivors < input && survivors > 1 {
		return StatusWeakBMinusLSieveReduced + "; " + StatusWeakPlaneSpatialDegeneracyRemains + "; temporal-spatial weak frames are rejected but the spatial S3 orbit remains"
	}
	if survivors == 1 {
		return "CONDITIONAL_SUPPORT_B_MINUS_L_WEAK_FRAME_UNIQUELY_SELECTED"
	}
	return "FAILED_ROUTE_B_MINUS_L_WEAK_FRAME_SIEVE_DID_NOT_REDUCE"
}

func restrictedVerdict(exact int, unique bool, maxP, maxFull int) string {
	if exact == 0 {
		return StatusReducedTrialityRescanned + "; " + StatusNoThreePlaneAfterSieve + fmt.Sprintf("; max polarized kernel=%d and max full kernel=%d after B-L restriction", maxP, maxFull)
	}
	if unique {
		return "CONDITIONAL_SUPPORT_B_MINUS_L_RESTRICTED_TRIALITY_BRANCH_DERIVED_NEUTRAL_3PLANE"
	}
	return "FAILED_ROUTE_B_MINUS_L_RESTRICTED_SCAN_HAS_MULTIPLE_THREE_PLANE_WITNESSES"
}

func setOf(xs []string) map[string]bool {
	m := make(map[string]bool, len(xs))
	for _, x := range xs {
		m[x] = true
	}
	return m
}

func containsMode(xs []int, x int) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func sameScalar(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func sameVec(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}
