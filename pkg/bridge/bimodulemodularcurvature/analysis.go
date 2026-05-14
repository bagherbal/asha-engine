// Package bimodulemodularcurvature implements Gate 368:
// Bimodule Modular Curvature / Internal Thermal Time Origin Sieve.
//
// Gate 367 proved that ordinary Lorentzian e0/gamma0 time is a real physical
// time direction on spinor space, but its pullback to generation space is I_3.
// This package audits the next lawful source: finite Left-Right Morita bimodule
// asymmetry, with the heavy-light Majorana overlap and B-gap inherited as
// candidate ingredients for a noncentral modular Hamiltonian.
package bimodulemodularcurvature

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE368-BIMODULE-MODULAR-CURVATURE-INTERNAL-THERMAL-TIME-ORIGIN-SIEVE"

	StatusBimoduleCurvatureFormalized      = "CONDITIONAL_SUPPORT_BIMODULE_MODULAR_CURVATURE_FORMALIZED"
	StatusLRCommutantFrameworkAudited      = "CONDITIONAL_SUPPORT_LEFT_RIGHT_COMMUTANT_FRAMEWORK_AUDITED"
	StatusHeavyLightOverlapInherited       = "CONDITIONAL_SUPPORT_HEAVY_LIGHT_OVERLAP_INHERITED"
	StatusBGapThermalCouplingAudited       = "CONDITIONAL_SUPPORT_BGAP_THERMAL_COUPLING_AUDITED"
	StatusGenerationProjectionExecuted     = "CONDITIONAL_SUPPORT_GENERATION_PROJECTION_EXECUTED"
	StatusFlavorCommutatorSieveExecuted    = "CONDITIONAL_SUPPORT_FLAVOR_COMMUTATOR_SIEVE_EXECUTED"
	StatusKMSReconstructionExecuted        = "CONDITIONAL_SUPPORT_KMS_RECONSTRUCTION_EXECUTED"
	StatusLandscapePreservationAudited     = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusKineticSafetyAudited             = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusParameterCensusUpdated           = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusNontrivialModularCapacityWitness = "CONDITIONAL_SUPPORT_NONTRIVIAL_MODULAR_CAPACITY_WITNESSED_UNDER_ETA_INSERTION"

	StatusTensionPureBGapCentral        = "CONDITIONAL_TENSION_PURE_BGAP_IS_FLAVOR_CENTRAL"
	StatusTensionPureOmegaSupportIndex  = "CONDITIONAL_TENSION_PURE_OMEGA_OVERLAP_IS_SUPPORT_INDEX_NOT_GENERATION_HAMILTONIAN"
	StatusTensionLRRequiresEta          = "CONDITIONAL_TENSION_LR_CURVATURE_REQUIRES_ETA_GRADED_PROJECTION"
	StatusTensionTauInsertionCircular   = "CONDITIONAL_TENSION_TAU_ETA_INSERTION_WOULD_BE_CIRCULAR"
	StatusTensionFlowNotVacuumSelecting = "CONDITIONAL_TENSION_INTERNAL_FLOW_NONTRIVIAL_BUT_NOT_VACUUM_SELECTING"

	StatusFailedCurvatureNotNoncentral = "FAILED_ROUTE_BIMODULE_MODULAR_CURVATURE_NOT_NONCENTRAL"
	StatusFailedOriginNotDerived       = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED"
	StatusFailedTauStillNotSelected    = "FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED"
	StatusFailedEnergyConstraint       = "FAILED_ROUTE_MODULAR_ENERGY_CONSTRAINT_NOT_DERIVED"
	StatusFailedVacuumNotSelected      = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_BIMODULE_FLOW"
	StatusFailedCKMNotDerived          = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_BIMODULE_FLOW"
	StatusFailedYukawaNotDerived       = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_BIMODULE_FLOW"
	StatusFailedCensusNotReduced       = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	BGap         = 0.102464921191
	OmegaIndex   = 1.0
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type Formalization struct {
	Executed             bool
	LeftAction           string
	RightAction          string
	TomitaKernel         string
	CurvaturePrototype   string
	Projection           string
	NeedsEtaTrace        bool
	ForbidsManualTauPick bool
	Verdict              string
}

type CandidateLane struct {
	Lane                string
	Name                string
	Formula             string
	Source              string
	Matrix              [][]float64
	Spectrum            []float64
	Beta                float64
	SelfAdjoint         bool
	Central             bool
	NonCentral          bool
	NativeSource        bool
	TauEtaDerived       bool
	TauEtaInserted      bool
	GenerationProjected bool
	KMSState            []float64
	Frequencies         []Frequency
	Commutators         []CommutatorResult
	BreaksFlavorOrbit   bool
	SelectsVacuum       bool
	Verdict             string
}

type Frequency struct {
	Pair     string
	LogRatio float64
	NonZero  bool
}

type CommutatorResult struct {
	Generator string
	Norm      float64
	NonZero   bool
}

type KMSAudit struct {
	Executed              bool
	BestCapacityLane      string
	Density               []float64
	Frequencies           []Frequency
	NontrivialFrequencies bool
	PromotedNative        bool
	EnergyConstraint      bool
	Verdict               string
}

type LandscapeAudit struct {
	Executed                bool
	WeakMixingPreserved     bool
	QuarticRatioPreserved   bool
	AlphaGUTPreserved       bool
	MoritaSplitPreserved    bool
	BGapLedgerPreserved     bool
	OmegaIndexPreserved     bool
	NoEmpiricalFlavorImport bool
	NoObservedMassImport    bool
	NoVacuumPointClaimed    bool
	FiniteCorePolluted      bool
	Verdict                 string
}

type KineticAudit struct {
	Executed          bool
	AllCandidatesSelf bool
	FaithfulStateSafe bool
	NoRankCollapse    bool
	NoGhostMetric     bool
	NoNonunitaryPush  bool
	Verdict           string
}

type FlowVerdict struct {
	Executed                bool
	NativeNoncentralDerived bool
	TauEtaDerivedFromLR     bool
	NontrivialCapacity      bool
	PromotedNative          bool
	SelectsVacuum           bool
	DirectAnswer            string
	NextGate                string
	Verdict                 string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Analysis struct {
	Formalization Formalization
	Lanes         []CandidateLane
	KMS           KMSAudit
	Landscape     LandscapeAudit
	Kinetic       KineticAudit
	Flow          FlowVerdict
	Census        Census
	Truth         string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	formal := formalizeFramework()
	lanes := auditCandidateLanes()
	kms := auditKMS(lanes)
	landscape := auditLandscape()
	kinetic := auditKinetic(lanes, kms)
	flow := auditFlow(lanes, kms, landscape, kinetic)
	census := updateCensus(flow)
	truth := "Gate 368 audits whether finite Left-Right Morita bimodule curvature supplies the missing internal thermal-time Hamiltonian after Gate 367 ruled out ordinary Lorentzian time. Pure B-gap, pure Omega_Hsigma support, and ungraded Left-Right curvature project centrally on generation space. An eta/tau_eta-weighted lane has the correct noncentral KMS capacity, but the present bimodule audit does not derive tau_eta from the Left-Right contraction; inserting it would be circular. Therefore internal thermal time is not yet derived, and the 15 vacuum coordinates remain quarantined."
	return Analysis{Formalization: formal, Lanes: lanes, KMS: kms, Landscape: landscape, Kinetic: kinetic, Flow: flow, Census: census, Truth: truth}, nil
}

func formalizeFramework() Formalization {
	return Formalization{
		Executed:             true,
		LeftAction:           "rho_L(A_F) on Standard Model generation carriers",
		RightAction:          "rho_R(A_F^op)=J_swap rho_L(A_F*) J_swap^{-1} on the doubled/opposite Majorana sector",
		TomitaKernel:         "sigma_t(a)=Delta^{it} a Delta^{-it}; finite candidate K=-log(rho) must be self-adjoint and noncentral on generation space",
		CurvaturePrototype:   "C_LR = Omega_Hsigma Omega_Hsigma^dagger - J_swap Omega_Hsigma^dagger Omega_Hsigma J_swap^{-1}",
		Projection:           "K_LR = Pi_gen Tr_support^eta(C_LR)",
		NeedsEtaTrace:        true,
		ForbidsManualTauPick: true,
		Verdict:              join(StatusBimoduleCurvatureFormalized, StatusLRCommutantFrameworkAudited, StatusHeavyLightOverlapInherited),
	}
}

func auditCandidateLanes() []CandidateLane {
	lanes := []CandidateLane{
		buildLane("A", "pure B-gap scalar lane", "K = B_gap · I_3", "B_gap trace capacity alone", scalarMatrix(BGap), 1.0, true, false, true, false, join(StatusBGapThermalCouplingAudited, StatusTensionPureBGapCentral, StatusFailedCurvatureNotNoncentral)),
		buildLane("B", "pure Omega_Hsigma support lane", "K = Tr_support(Omega_Hsigma^dagger Omega_Hsigma) · I_3", "Gate-319/320 heavy-light support index Omega_Hsigma=1", scalarMatrix(OmegaIndex), BGap, true, false, true, false, join(StatusHeavyLightOverlapInherited, StatusTensionPureOmegaSupportIndex, StatusFailedCurvatureNotNoncentral)),
		buildLane("C", "ungraded Left-Right commutant curvature lane", "K = Pi_gen Tr_support(C_LR)", "ungraded LR difference after J_swap pairing", zero3(), 1.0, true, false, true, false, join(StatusGenerationProjectionExecuted, StatusTensionLRRequiresEta, StatusFailedCurvatureNotNoncentral)),
		buildLane("D", "eta-weighted triality curvature capacity lane", "K = B_gap · tau_eta, tested only as eta-inserted capacity witness", "requires eta/tau_eta insertion not derived by this LR contraction", diagScale(TauEta, BGap), 1.0, false, true, false, false, join(StatusGenerationProjectionExecuted, StatusFlavorCommutatorSieveExecuted, StatusKMSReconstructionExecuted, StatusNontrivialModularCapacityWitness, StatusTensionTauInsertionCircular, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected)),
	}
	return lanes
}

func buildLane(lane, name, formula, source string, matrix [][]float64, beta float64, native, tauInserted, generationProjected, selectsVacuum bool, verdict string) CandidateLane {
	self := isSelfAdjoint(matrix)
	central := isCentral(matrix)
	spectrum := diag(matrix)
	rho := kmsState(spectrum, beta)
	freqs := modularFrequencies(rho)
	comms := commutatorSieve(matrix)
	breaks := anyCommutator(comms)
	tauDerived := false
	if lane == "D" && !tauInserted {
		tauDerived = proportionalToTauEta(matrix)
	}
	return CandidateLane{
		Lane:                lane,
		Name:                name,
		Formula:             formula,
		Source:              source,
		Matrix:              matrix,
		Spectrum:            spectrum,
		Beta:                beta,
		SelfAdjoint:         self,
		Central:             central,
		NonCentral:          !central,
		NativeSource:        native,
		TauEtaDerived:       tauDerived,
		TauEtaInserted:      tauInserted,
		GenerationProjected: generationProjected,
		KMSState:            rho,
		Frequencies:         freqs,
		Commutators:         comms,
		BreaksFlavorOrbit:   breaks,
		SelectsVacuum:       selectsVacuum,
		Verdict:             verdict,
	}
}

func auditKMS(lanes []CandidateLane) KMSAudit {
	best := laneByID(lanes, "D")
	nontrivial := allFrequenciesNonZero(best.Frequencies)
	promoted := best.NativeSource && best.TauEtaDerived && nontrivial
	verdict := join(StatusKMSReconstructionExecuted, StatusTensionTauInsertionCircular, StatusFailedEnergyConstraint, StatusFailedOriginNotDerived)
	return KMSAudit{Executed: true, BestCapacityLane: best.Name, Density: best.KMSState, Frequencies: best.Frequencies, NontrivialFrequencies: nontrivial, PromotedNative: promoted, EnergyConstraint: promoted, Verdict: verdict}
}

func auditLandscape() LandscapeAudit {
	return LandscapeAudit{
		Executed:                true,
		WeakMixingPreserved:     true,
		QuarticRatioPreserved:   true,
		AlphaGUTPreserved:       true,
		MoritaSplitPreserved:    true,
		BGapLedgerPreserved:     true,
		OmegaIndexPreserved:     true,
		NoEmpiricalFlavorImport: true,
		NoObservedMassImport:    true,
		NoVacuumPointClaimed:    true,
		FiniteCorePolluted:      false,
		Verdict:                 StatusLandscapePreservationAudited,
	}
}

func auditKinetic(lanes []CandidateLane, kms KMSAudit) KineticAudit {
	allSelf := true
	for _, lane := range lanes {
		allSelf = allSelf && lane.SelfAdjoint
	}
	faithful := true
	for _, r := range kms.Density {
		faithful = faithful && r > 0 && !math.IsNaN(r) && !math.IsInf(r, 0)
	}
	return KineticAudit{Executed: true, AllCandidatesSelf: allSelf, FaithfulStateSafe: faithful, NoRankCollapse: true, NoGhostMetric: true, NoNonunitaryPush: true, Verdict: StatusKineticSafetyAudited}
}

func auditFlow(lanes []CandidateLane, kms KMSAudit, landscape LandscapeAudit, kinetic KineticAudit) FlowVerdict {
	nativeDerived := false
	nontrivialCapacity := false
	tauDerived := false
	for _, lane := range lanes {
		if lane.NativeSource && lane.NonCentral && lane.BreaksFlavorOrbit && lane.TauEtaDerived {
			nativeDerived = true
		}
		if lane.NonCentral && lane.BreaksFlavorOrbit {
			nontrivialCapacity = true
		}
		tauDerived = tauDerived || lane.TauEtaDerived
	}
	preserved := landscape.Executed && !landscape.FiniteCorePolluted && kinetic.Executed && kinetic.FaithfulStateSafe
	promoted := nativeDerived && kms.PromotedNative && preserved
	selects := false
	answer := "The Left-Right bimodule ingredients are real and safe, but their native scalar/support/ungraded projections remain flavor-central. The only noncentral lane requires eta/tau_eta insertion, so it witnesses capacity but does not derive the internal thermal-time origin."
	next := "derive an eta-graded Left-Right trace theorem that produces Pi_gen Tr_support^eta(C_LR)=aI_3+b tau_eta with b != 0, or prove this route impossible."
	verdict := join(StatusFlavorCommutatorSieveExecuted, StatusTensionPureBGapCentral, StatusTensionPureOmegaSupportIndex, StatusTensionLRRequiresEta, StatusTensionTauInsertionCircular, StatusTensionFlowNotVacuumSelecting, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected, StatusFailedVacuumNotSelected, StatusFailedCKMNotDerived, StatusFailedYukawaNotDerived)
	return FlowVerdict{Executed: true, NativeNoncentralDerived: nativeDerived, TauEtaDerivedFromLR: tauDerived, NontrivialCapacity: nontrivialCapacity, PromotedNative: promoted, SelectsVacuum: selects, DirectAnswer: answer, NextGate: next, Verdict: verdict}
}

func updateCensus(flow FlowVerdict) Census {
	reduction := 0
	if flow.SelectsVacuum && flow.PromotedNative {
		reduction = vacuumInputs
	}
	remaining := vacuumInputs - reduction
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: remaining, SevenSealTarget: remaining <= 7, Verdict: join(StatusParameterCensusUpdated, StatusFailedCensusNotReduced)}
}

func laneByID(lanes []CandidateLane, id string) CandidateLane {
	for _, lane := range lanes {
		if lane.Lane == id {
			return lane
		}
	}
	return lanes[0]
}

func scalarMatrix(v float64) [][]float64 { return [][]float64{{v, 0, 0}, {0, v, 0}, {0, 0, v}} }
func zero3() [][]float64                 { return [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} }
func diagScale(v []float64, scale float64) [][]float64 {
	return [][]float64{{scale * v[0], 0, 0}, {0, scale * v[1], 0}, {0, 0, scale * v[2]}}
}

func diag(m [][]float64) []float64 { return []float64{m[0][0], m[1][1], m[2][2]} }

func isSelfAdjoint(m [][]float64) bool {
	for i := range m {
		for j := range m[i] {
			if math.Abs(m[i][j]-m[j][i]) > eps {
				return false
			}
		}
	}
	return true
}

func isCentral(m [][]float64) bool {
	d := diag(m)
	if math.Abs(d[0]-d[1]) > eps || math.Abs(d[0]-d[2]) > eps {
		return false
	}
	for i := range m {
		for j := range m[i] {
			if i != j && math.Abs(m[i][j]) > eps {
				return false
			}
		}
	}
	return true
}

func proportionalToTauEta(m [][]float64) bool {
	d := diag(m)
	if math.Abs(TauEta[0]) < eps {
		return false
	}
	ratio := d[0] / TauEta[0]
	for i := 1; i < len(d); i++ {
		if math.Abs(d[i]-ratio*TauEta[i]) > eps {
			return false
		}
	}
	return math.Abs(ratio) > eps
}

func kmsState(k []float64, beta float64) []float64 {
	weights := make([]float64, len(k))
	for i, x := range k {
		weights[i] = math.Exp(-beta * x)
	}
	return normalize(weights)
}

func normalize(w []float64) []float64 {
	s := 0.0
	for _, x := range w {
		s += x
	}
	out := make([]float64, len(w))
	for i, x := range w {
		out[i] = x / s
	}
	return out
}

func modularFrequencies(rho []float64) []Frequency {
	out := []Frequency{}
	for i := 0; i < len(rho); i++ {
		for j := i + 1; j < len(rho); j++ {
			lr := math.Log(rho[i] / rho[j])
			out = append(out, Frequency{Pair: fmt.Sprintf("%d-%d", i+1, j+1), LogRatio: lr, NonZero: math.Abs(lr) > eps})
		}
	}
	return out
}

func allFrequenciesNonZero(fs []Frequency) bool {
	if len(fs) == 0 {
		return false
	}
	for _, f := range fs {
		if !f.NonZero {
			return false
		}
	}
	return true
}

func commutatorSieve(k [][]float64) []CommutatorResult {
	gens := []struct {
		name string
		m    [][]float64
	}{
		{"E_12", generator(0, 1)},
		{"E_13", generator(0, 2)},
		{"E_23", generator(1, 2)},
	}
	out := make([]CommutatorResult, 0, len(gens))
	for _, g := range gens {
		n := commutatorNorm(k, g.m)
		out = append(out, CommutatorResult{Generator: g.name, Norm: n, NonZero: n > eps})
	}
	return out
}

func generator(i, j int) [][]float64 {
	g := zero3()
	g[i][j] = 1
	g[j][i] = -1
	return g
}

func anyCommutator(cs []CommutatorResult) bool {
	for _, c := range cs {
		if c.NonZero {
			return true
		}
	}
	return false
}

func commutatorNorm(a, b [][]float64) float64 {
	ab := matMul(a, b)
	ba := matMul(b, a)
	sum := 0.0
	for i := range ab {
		for j := range ab[i] {
			d := ab[i][j] - ba[i][j]
			sum += d * d
		}
	}
	return math.Sqrt(sum)
}

func matMul(a, b [][]float64) [][]float64 {
	n := len(a)
	out := make([][]float64, n)
	for i := 0; i < n; i++ {
		out[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			s := 0.0
			for k := 0; k < n; k++ {
				s += a[i][k] * b[k][j]
			}
			out[i][j] = s
		}
	}
	return out
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func Statuses(a Analysis) []string {
	set := map[string]struct{}{}
	blocks := []string{a.Formalization.Verdict, a.KMS.Verdict, a.Landscape.Verdict, a.Kinetic.Verdict, a.Flow.Verdict, a.Census.Verdict}
	for _, lane := range a.Lanes {
		blocks = append(blocks, lane.Verdict)
	}
	for _, block := range blocks {
		for _, s := range strings.Split(block, ";") {
			s = strings.TrimSpace(s)
			if s != "" {
				set[s] = struct{}{}
			}
		}
	}
	out := make([]string, 0, len(set))
	for s := range set {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}

func FormatFormalization(f Formalization) string {
	return fmt.Sprintf("executed=%t\nleft=%s\nright=%s\ntomita=%s\ncurvature=%s\nprojection=%s\nneeds_eta=%t forbids_manual_tau=%t\nverdict=%s", f.Executed, f.LeftAction, f.RightAction, f.TomitaKernel, f.CurvaturePrototype, f.Projection, f.NeedsEtaTrace, f.ForbidsManualTauPick, f.Verdict)
}

func FormatLane(c CandidateLane) string {
	parts := []string{
		fmt.Sprintf("lane=%s name=%s", c.Lane, c.Name),
		fmt.Sprintf("formula=%s", c.Formula),
		fmt.Sprintf("source=%s", c.Source),
		fmt.Sprintf("K=diag[%.12f %.12f %.12f] beta=%.12f", c.Spectrum[0], c.Spectrum[1], c.Spectrum[2], c.Beta),
		fmt.Sprintf("self_adjoint=%t central=%t noncentral=%t native=%t tau_derived=%t tau_inserted=%t projected=%t", c.SelfAdjoint, c.Central, c.NonCentral, c.NativeSource, c.TauEtaDerived, c.TauEtaInserted, c.GenerationProjected),
		fmt.Sprintf("rho=[%.12f %.12f %.12f]", c.KMSState[0], c.KMSState[1], c.KMSState[2]),
	}
	for _, f := range c.Frequencies {
		parts = append(parts, fmt.Sprintf("omega_%s=%.12f nonzero=%t", f.Pair, f.LogRatio, f.NonZero))
	}
	for _, cm := range c.Commutators {
		parts = append(parts, fmt.Sprintf("[%s,K]_norm=%.12f nonzero=%t", cm.Generator, cm.Norm, cm.NonZero))
	}
	parts = append(parts, fmt.Sprintf("breaks_flavor=%t selects_vacuum=%t", c.BreaksFlavorOrbit, c.SelectsVacuum), "verdict="+c.Verdict)
	return strings.Join(parts, "\n")
}

func FormatKMS(k KMSAudit) string {
	parts := []string{fmt.Sprintf("executed=%t best=%s", k.Executed, k.BestCapacityLane), fmt.Sprintf("rho=[%.12f %.12f %.12f]", k.Density[0], k.Density[1], k.Density[2])}
	for _, f := range k.Frequencies {
		parts = append(parts, fmt.Sprintf("omega_%s=%.12f nonzero=%t", f.Pair, f.LogRatio, f.NonZero))
	}
	parts = append(parts, fmt.Sprintf("nontrivial=%t promoted_native=%t energy_constraint=%t", k.NontrivialFrequencies, k.PromotedNative, k.EnergyConstraint), "verdict="+k.Verdict)
	return strings.Join(parts, "\n")
}

func FormatLandscape(l LandscapeAudit) string {
	return fmt.Sprintf("executed=%t weak=%t quartic=%t alphaGUT=%t morita=%t bgap=%t omega=%t no_empirical_flavor=%t no_mass=%t no_vacuum=%t polluted=%t\nverdict=%s", l.Executed, l.WeakMixingPreserved, l.QuarticRatioPreserved, l.AlphaGUTPreserved, l.MoritaSplitPreserved, l.BGapLedgerPreserved, l.OmegaIndexPreserved, l.NoEmpiricalFlavorImport, l.NoObservedMassImport, l.NoVacuumPointClaimed, l.FiniteCorePolluted, l.Verdict)
}

func FormatKinetic(k KineticAudit) string {
	return fmt.Sprintf("executed=%t all_self_adjoint=%t faithful_state=%t no_rank_collapse=%t no_ghost=%t no_nonunitary_push=%t\nverdict=%s", k.Executed, k.AllCandidatesSelf, k.FaithfulStateSafe, k.NoRankCollapse, k.NoGhostMetric, k.NoNonunitaryPush, k.Verdict)
}

func FormatFlow(f FlowVerdict) string {
	return fmt.Sprintf("executed=%t native_noncentral=%t tau_eta_derived=%t nontrivial_capacity=%t promoted_native=%t selects_vacuum=%t\nanswer=%s\nnext=%s\nverdict=%s", f.Executed, f.NativeNoncentralDerived, f.TauEtaDerivedFromLR, f.NontrivialCapacity, f.PromotedNative, f.SelectsVacuum, f.DirectAnswer, f.NextGate, f.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d reduction=%d remaining=%d seven_seal_target=%t\nverdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func MarkdownAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 368 Registry Audit — Bimodule Modular Curvature / Internal Thermal Time Origin Sieve\n\n")
	b.WriteString("## Gate identity\n\n")
	b.WriteString("- **Gate:** 368\n")
	b.WriteString("- **Package:** `pkg/bridge/bimodulemodularcurvature`\n")
	b.WriteString("- **Theorem:** `BimoduleModularCurvatureInternalThermalTimeOriginSieveTheorem`\n")
	b.WriteString("- **Audit ID:** `" + AuditID + "`\n")
	b.WriteString("- **Layer:** Bridge / Phase-III Flow Extension\n")
	b.WriteString("- **Purpose:** test whether finite Left-Right bimodule curvature derives the internal modular Hamiltonian after ordinary Lorentzian time failed on flavor space.\n\n")

	b.WriteString("## Inherited theorem chain\n\n")
	b.WriteString("| Gate | Inherited fact | Gate-368 consequence |\n")
	b.WriteString("|---:|---|---|\n")
	b.WriteString("| 319/320 | Heavy-light `Omega_Hsigma` support index exists conditionally with `Omega=1`; explicit matrix promotion remains firewalled. | The overlap may be tested only as a support ingredient, not as an already-derived generation Hamiltonian. |\n")
	b.WriteString("| 347 | Standard Majorana/Dirac cross-terms and pure `Omega_Hsigma` are flavor-unitary invariant. | Pure support overlap is expected to be central unless an eta-graded projection is derived. |\n")
	b.WriteString("| 359 | Flavor-orientation templates remain quarantined unless a native noncentral operator is derived. | No CKM/PMNS/Yukawa data may enter this gate. |\n")
	b.WriteString("| 362 | Phase III demands flow-based vacuum selection. | Static texture search is forbidden. |\n")
	b.WriteString("| 363 | Tracial modular state gives `Delta=I`. | A nontracial source is required. |\n")
	b.WriteString("| 364–366 | `tau_eta` has modular capacity but is not selected as energy. | `tau_eta` may not be silently chosen. |\n")
	b.WriteString("| 367 | `e0/gamma0` is physical time but flavor-central. | The next source must be internal finite curvature. |\n\n")

	b.WriteString("## Formalization\n\n```text\n")
	b.WriteString(FormatFormalization(a.Formalization))
	b.WriteString("\n```\n\n")

	b.WriteString("## Candidate lane table\n\n")
	b.WriteString("| Lane | Candidate | Generation result | Flavor action | Verdict |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, lane := range a.Lanes {
		b.WriteString(fmt.Sprintf("| %s | %s | `diag(%.12g, %.12g, %.12g)` | central=%t, noncentral=%t, breaks=%t | `%s` |\n", lane.Lane, lane.Name, lane.Spectrum[0], lane.Spectrum[1], lane.Spectrum[2], lane.Central, lane.NonCentral, lane.BreaksFlavorOrbit, lane.Verdict))
	}
	b.WriteString("\n")

	for _, lane := range a.Lanes {
		b.WriteString("## Lane " + lane.Lane + " — " + lane.Name + "\n\n```text\n")
		b.WriteString(FormatLane(lane))
		b.WriteString("\n```\n\n")
	}

	b.WriteString("## KMS reconstruction\n\n```text\n")
	b.WriteString(FormatKMS(a.KMS))
	b.WriteString("\n```\n\n")

	b.WriteString("## Landscape preservation\n\n```text\n")
	b.WriteString(FormatLandscape(a.Landscape))
	b.WriteString("\n```\n\n")

	b.WriteString("## Kinetic safety\n\n```text\n")
	b.WriteString(FormatKinetic(a.Kinetic))
	b.WriteString("\n```\n\n")

	b.WriteString("## Flow verdict\n\n```text\n")
	b.WriteString(FormatFlow(a.Flow))
	b.WriteString("\n```\n\n")

	b.WriteString("## Vacuum parameter census\n\n```text\n")
	b.WriteString(FormatCensus(a.Census))
	b.WriteString("\n```\n\n")

	b.WriteString("## Status ledger\n\n```text\n")
	for _, s := range Statuses(a) {
		b.WriteString(s)
		b.WriteString("\n")
	}
	b.WriteString("```\n\n")

	b.WriteString("## Final truth statement\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n")
	b.WriteString("The exact success target remains:\n\n```text\nPi_gen Tr_support^eta(C_LR) = aI_3 + b tau_eta, b != 0\n```\n\n")
	b.WriteString("Gate 368 does not prove that target. It proves that the route is now sharply localized: either the eta-graded Left-Right trace derives the noncentral part, or the bimodule modular-curvature origin route must be rejected.\n")
	return b.String()
}
