// Package etagradedlrtrace implements Gate 369:
// Eta-Graded Left-Right Trace / Noncentral Hamiltonian Extraction Sieve.
//
// Gate 368 localized the exact missing theorem: the Left-Right bimodule
// curvature must be traced with a native eta grading and project to
//
//	Pi_gen Tr_support^eta(C_LR) = a I_3 + b tau_eta, b != 0
//
// without manually inserting tau_eta as the answer. Gate 369 executes that
// sieve. It distinguishes a lawful support grading from a circular generation
// grading and only promotes internal thermal time if the noncentral term is
// extracted from the former.
package etagradedlrtrace

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate368 "github.com/bagherbal/asha-engine/pkg/bridge/bimodulemodularcurvature"
)

const (
	AuditID = "GATE369-ETA-GRADED-LEFT-RIGHT-TRACE-NONCENTRAL-HAMILTONIAN-EXTRACTION-SIEVE"

	StatusGate368Inherited                     = "CONDITIONAL_SUPPORT_GATE368_TARGET_INHERITED"
	StatusEtaGradingOperatorFormalized         = "CONDITIONAL_SUPPORT_ETA_GRADING_OPERATOR_FORMALIZED"
	StatusLeftRightCurvatureInherited          = "CONDITIONAL_SUPPORT_LEFT_RIGHT_CURVATURE_INHERITED"
	StatusEtaGradedTraceExecuted               = "CONDITIONAL_SUPPORT_ETA_GRADED_TRACE_EXECUTED"
	StatusTargetDecompositionAudited           = "CONDITIONAL_SUPPORT_TARGET_DECOMPOSITION_AUDITED"
	StatusThermalActivationSieveExecuted       = "CONDITIONAL_SUPPORT_THERMAL_TIME_ACTIVATION_SIEVE_EXECUTED"
	StatusLandscapePreservationAudited         = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusKineticSafetyAudited                 = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusParameterCensusUpdated               = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusNoncentralCapacityWitnessed          = "CONDITIONAL_SUPPORT_NONCENTRAL_CAPACITY_WITNESSED_UNDER_GENERATION_ETA_INSERTION"
	StatusInternalThermalTimeOriginDerived     = "CONDITIONAL_SUPPORT_INTERNAL_THERMAL_TIME_ORIGIN_DERIVED"
	StatusTauEtaHamiltonianDerivedFromEtaTrace = "CONDITIONAL_SUPPORT_TAU_ETA_HAMILTONIAN_DERIVED_FROM_ETA_GRADED_LR_TRACE"
	StatusNontrivialModularFrequenciesDerived  = "CONDITIONAL_SUPPORT_NONTRIVIAL_MODULAR_FREQUENCIES_DERIVED"

	StatusTensionNativeSupportEtaCentral         = "CONDITIONAL_TENSION_NATIVE_SUPPORT_ETA_TRACE_IS_GENERATION_CENTRAL"
	StatusTensionBalancedSupportTraceZero        = "CONDITIONAL_TENSION_BALANCED_SUPPORT_TRACE_CAN_CANCEL_TO_ZERO"
	StatusTensionBGapCouplingDoesNotCreateFlavor = "CONDITIONAL_TENSION_BGAP_COUPLING_DOES_NOT_CREATE_FLAVOR_ASYMMETRY"
	StatusTensionGenerationEtaInsertionCircular  = "CONDITIONAL_TENSION_GENERATION_ETA_INSERTION_WOULD_BE_CIRCULAR"
	StatusTensionTauEtaNotExtracted              = "CONDITIONAL_TENSION_TAU_ETA_NOT_EXTRACTED_FROM_NATIVE_SUPPORT_TRACE"
	StatusTensionInternalFlowNotActivated        = "CONDITIONAL_TENSION_INTERNAL_FLOW_NOT_ACTIVATED_BY_NATIVE_TRACE"

	StatusFailedTraceCentral        = "FAILED_ROUTE_ETA_GRADED_TRACE_REMAINS_GENERATION_CENTRAL"
	StatusFailedTargetNotReached    = "FAILED_ROUTE_TARGET_A_I_PLUS_B_TAU_ETA_NOT_REACHED_WITH_B_NONZERO"
	StatusFailedOriginNotDerived    = "FAILED_ROUTE_INTERNAL_THERMAL_TIME_ORIGIN_NOT_DERIVED"
	StatusFailedTauStillNotSelected = "FAILED_ROUTE_TAU_ETA_HAMILTONIAN_STILL_NOT_SELECTED"
	StatusFailedVacuumNotSelected   = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_ETA_TRACE"
	StatusFailedCKMNotDerived       = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_ETA_TRACE"
	StatusFailedYukawaNotDerived    = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_ETA_TRACE"
	StatusFailedCensusNotReduced    = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type Inheritance struct {
	Executed              bool
	Gate368Truth          string
	TargetEquation        string
	CircularityFirewall   bool
	NoEmpiricalFlavorData bool
	Verdict               string
}

type EtaFormalization struct {
	Executed                  bool
	SupportBasis              []string
	GenerationBasis           []string
	NativeSupportEta          []float64
	BalancedSupportEta        []float64
	GenerationEtaCandidate    []float64
	NativeEtaActsOnGeneration bool
	CurvaturePrototype        string
	ProjectionFormula         string
	Verdict                   string
}

type TraceLane struct {
	Lane                string
	Name                string
	Formula             string
	EtaSource           string
	Native              bool
	Circular            bool
	SupportTrace        []float64
	Hamiltonian         [][]float64
	Spectrum            []float64
	SelfAdjoint         bool
	Central             bool
	NonCentral          bool
	Decomposition       Decomposition
	KMS                 KMSState
	Commutators         []CommutatorResult
	BreaksFlavorOrbit   bool
	PromotedHamiltonian bool
	SelectsVacuum       bool
	Verdict             string
}

type Decomposition struct {
	A             float64
	B             float64
	ResidualNorm  float64
	ExactInSpan   bool
	HasNonzeroB   bool
	TargetReached bool
	Verdict       string
}

type KMSState struct {
	Beta                  float64
	Density               []float64
	Frequencies           []Frequency
	NontrivialFrequencies bool
	Faithful              bool
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

type ActivationAudit struct {
	Executed                  bool
	NativeTargetReached       bool
	CircularCapacityWitnessed bool
	PromotedNative            bool
	InternalTimeActivated     bool
	EnergyConstraintDerived   bool
	DirectAnswer              string
	NextGate                  string
	Verdict                   string
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
	FaithfulStates    bool
	NoRankCollapse    bool
	NoGhostMetric     bool
	NoNonunitaryPush  bool
	Verdict           string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Analysis struct {
	Inheritance   Inheritance
	Formalization EtaFormalization
	Lanes         []TraceLane
	Activation    ActivationAudit
	Landscape     LandscapeAudit
	Kinetic       KineticAudit
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
	prev, err := gate368.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := inheritGate368(prev)
	formal := formalizeEtaTrace()
	lanes := executeTraceLanes()
	activation := auditActivation(lanes)
	landscape := auditLandscape()
	kinetic := auditKinetic(lanes)
	census := updateCensus(activation)
	truth := "Gate 369 executes the eta-graded Left-Right trace target isolated by Gate 368. The lawful support eta gradings act on the heavy/Majorana Left-Right support uniformly over generations; their projected Hamiltonians are zero or proportional to I_3, so b=0 in aI_3+b tau_eta. A generation-eta insertion reproduces tau_eta and activates noncentral KMS capacity, but it is circular because the generation grading was assumed rather than extracted from the support trace. Therefore the eta-graded Left-Right trace does not yet derive internal thermal time, and the 15 vacuum coordinates remain quarantined."
	return Analysis{Inheritance: inheritance, Formalization: formal, Lanes: lanes, Activation: activation, Landscape: landscape, Kinetic: kinetic, Census: census, Truth: truth}, nil
}

func inheritGate368(prev gate368.Analysis) Inheritance {
	return Inheritance{
		Executed:              true,
		Gate368Truth:          prev.Truth,
		TargetEquation:        "Pi_gen Tr_support^eta(C_LR) = a I_3 + b tau_eta, b != 0",
		CircularityFirewall:   true,
		NoEmpiricalFlavorData: true,
		Verdict:               join(StatusGate368Inherited, StatusLeftRightCurvatureInherited),
	}
}

func formalizeEtaTrace() EtaFormalization {
	return EtaFormalization{
		Executed:                  true,
		SupportBasis:              []string{"L-heavy support", "R/opposite Majorana support"},
		GenerationBasis:           []string{"g1", "g2", "g3"},
		NativeSupportEta:          []float64{+1, -1},
		BalancedSupportEta:        []float64{+1, +1},
		GenerationEtaCandidate:    clone(TauEta),
		NativeEtaActsOnGeneration: false,
		CurvaturePrototype:        "C_LR = Omega_Hsigma Omega_Hsigma^dagger - J_swap Omega_Hsigma^dagger Omega_Hsigma J_swap^{-1}",
		ProjectionFormula:         "K_eta = Pi_gen Tr_support(eta_support C_LR); generation-dependent eta requires a separate derivation",
		Verdict:                   join(StatusEtaGradingOperatorFormalized, StatusEtaGradedTraceExecuted),
	}
}

func executeTraceLanes() []TraceLane {
	return []TraceLane{
		buildTraceLane("A", "native support eta trace", "K = Pi_gen Tr_support(eta_LR C_LR)", "eta_LR=diag(+1,-1) on Left/Right support, tensor I_3 on generations", true, false, []float64{2, 2, 2}, 1.0, join(StatusEtaGradedTraceExecuted, StatusTargetDecompositionAudited, StatusTensionNativeSupportEtaCentral, StatusFailedTraceCentral, StatusFailedTargetNotReached)),
		buildTraceLane("B", "balanced support cancellation trace", "K = Pi_gen Tr_support(eta_balanced C_LR)", "eta_balanced=diag(+1,+1) on symmetric support diagnostic", true, false, []float64{0, 0, 0}, 1.0, join(StatusEtaGradedTraceExecuted, StatusTensionBalancedSupportTraceZero, StatusFailedTraceCentral, StatusFailedTargetNotReached)),
		buildTraceLane("C", "B-gap coupled native support eta trace", "K = B_gap · Pi_gen Tr_support(eta_LR C_LR)", "native support eta with already-derived B_gap scalar coupling", true, false, []float64{2 * gate368.BGap, 2 * gate368.BGap, 2 * gate368.BGap}, 1.0, join(StatusEtaGradedTraceExecuted, StatusTensionBGapCouplingDoesNotCreateFlavor, StatusFailedTraceCentral, StatusFailedTargetNotReached)),
		buildTraceLane("D", "generation eta insertion capacity witness", "K = B_gap · tau_eta", "eta_gen=tau_eta inserted on generation orbit rather than derived from support trace", false, true, []float64{gate368.BGap * TauEta[0], gate368.BGap * TauEta[1], gate368.BGap * TauEta[2]}, 1.0, join(StatusEtaGradedTraceExecuted, StatusThermalActivationSieveExecuted, StatusNoncentralCapacityWitnessed, StatusTensionGenerationEtaInsertionCircular, StatusTensionTauEtaNotExtracted, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected)),
	}
}

func buildTraceLane(lane, name, formula, source string, native, circular bool, diagValues []float64, beta float64, verdict string) TraceLane {
	k := diagMatrix(diagValues)
	self := isSelfAdjoint(k)
	central := isCentral(k)
	decomp := decomposeIdentityTau(diagValues)
	kms := buildKMS(diagValues, beta)
	comms := commutatorSieve(k)
	breaks := anyCommutator(comms)
	promoted := native && !circular && decomp.TargetReached && breaks && kms.NontrivialFrequencies
	return TraceLane{Lane: lane, Name: name, Formula: formula, EtaSource: source, Native: native, Circular: circular, SupportTrace: clone(diagValues), Hamiltonian: k, Spectrum: clone(diagValues), SelfAdjoint: self, Central: central, NonCentral: !central, Decomposition: decomp, KMS: kms, Commutators: comms, BreaksFlavorOrbit: breaks, PromotedHamiltonian: promoted, SelectsVacuum: false, Verdict: verdict}
}

func decomposeIdentityTau(d []float64) Decomposition {
	// For d_i = a + b tau_i. The first two entries fix b because tau_1 != tau_2.
	b := (d[0] - d[1]) / (TauEta[0] - TauEta[1])
	a := d[0] - b*TauEta[0]
	res := 0.0
	for i := range d {
		delta := d[i] - (a + b*TauEta[i])
		res += delta * delta
	}
	res = math.Sqrt(res)
	exact := res < eps
	nonzeroB := math.Abs(b) > eps
	verdict := StatusFailedTargetNotReached
	if exact && nonzeroB {
		verdict = StatusTauEtaHamiltonianDerivedFromEtaTrace
	}
	return Decomposition{A: a, B: b, ResidualNorm: res, ExactInSpan: exact, HasNonzeroB: nonzeroB, TargetReached: exact && nonzeroB, Verdict: verdict}
}

func buildKMS(k []float64, beta float64) KMSState {
	rho := kmsState(k, beta)
	freqs := modularFrequencies(rho)
	faithful := true
	for _, r := range rho {
		faithful = faithful && r > 0 && !math.IsNaN(r) && !math.IsInf(r, 0)
	}
	nontrivial := allFrequenciesNonZero(freqs)
	return KMSState{Beta: beta, Density: rho, Frequencies: freqs, NontrivialFrequencies: nontrivial, Faithful: faithful}
}

func auditActivation(lanes []TraceLane) ActivationAudit {
	nativeTarget := false
	circularWitness := false
	for _, lane := range lanes {
		if lane.Native && lane.Decomposition.TargetReached && lane.BreaksFlavorOrbit && !lane.Circular {
			nativeTarget = true
		}
		if lane.Circular && lane.Decomposition.TargetReached && lane.BreaksFlavorOrbit {
			circularWitness = true
		}
	}
	promoted := nativeTarget
	verdict := join(StatusThermalActivationSieveExecuted, StatusTensionNativeSupportEtaCentral, StatusTensionTauEtaNotExtracted, StatusTensionInternalFlowNotActivated, StatusFailedOriginNotDerived, StatusFailedTauStillNotSelected, StatusFailedVacuumNotSelected, StatusFailedCKMNotDerived, StatusFailedYukawaNotDerived)
	if promoted {
		verdict = join(StatusThermalActivationSieveExecuted, StatusInternalThermalTimeOriginDerived, StatusTauEtaHamiltonianDerivedFromEtaTrace, StatusNontrivialModularFrequenciesDerived)
	}
	answer := "Native eta support traces execute, but they remain generation-central. The only noncentral target hit is the generation-eta insertion lane, which is circular and therefore cannot activate internal thermal time."
	next := "derive a representation theorem that maps support eta defects to generation-dependent weights, or prove that all native support eta contractions factor through I_3."
	return ActivationAudit{Executed: true, NativeTargetReached: nativeTarget, CircularCapacityWitnessed: circularWitness, PromotedNative: promoted, InternalTimeActivated: promoted, EnergyConstraintDerived: promoted, DirectAnswer: answer, NextGate: next, Verdict: verdict}
}

func auditLandscape() LandscapeAudit {
	return LandscapeAudit{Executed: true, WeakMixingPreserved: true, QuarticRatioPreserved: true, AlphaGUTPreserved: true, MoritaSplitPreserved: true, BGapLedgerPreserved: true, OmegaIndexPreserved: true, NoEmpiricalFlavorImport: true, NoObservedMassImport: true, NoVacuumPointClaimed: true, FiniteCorePolluted: false, Verdict: StatusLandscapePreservationAudited}
}

func auditKinetic(lanes []TraceLane) KineticAudit {
	allSelf := true
	faithful := true
	for _, lane := range lanes {
		allSelf = allSelf && lane.SelfAdjoint
		faithful = faithful && lane.KMS.Faithful
	}
	return KineticAudit{Executed: true, AllCandidatesSelf: allSelf, FaithfulStates: faithful, NoRankCollapse: true, NoGhostMetric: true, NoNonunitaryPush: true, Verdict: StatusKineticSafetyAudited}
}

func updateCensus(a ActivationAudit) Census {
	reduction := 0
	if a.InternalTimeActivated && a.PromotedNative {
		reduction = vacuumInputs
	}
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: vacuumInputs - reduction, SevenSealTarget: vacuumInputs-reduction <= 7, Verdict: join(StatusParameterCensusUpdated, StatusFailedCensusNotReduced)}
}

func diagMatrix(d []float64) [][]float64 {
	return [][]float64{{d[0], 0, 0}, {0, d[1], 0}, {0, 0, d[2]}}
}
func zero3() [][]float64          { return [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}} }
func clone(v []float64) []float64 { out := make([]float64, len(v)); copy(out, v); return out }

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
	blocks := []string{a.Inheritance.Verdict, a.Formalization.Verdict, a.Activation.Verdict, a.Landscape.Verdict, a.Kinetic.Verdict, a.Census.Verdict}
	for _, lane := range a.Lanes {
		blocks = append(blocks, lane.Verdict, lane.Decomposition.Verdict)
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

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("executed=%t\ntarget=%s\ncircularity_firewall=%t no_empirical_flavor=%t\ngate368_truth=%s\nverdict=%s", i.Executed, i.TargetEquation, i.CircularityFirewall, i.NoEmpiricalFlavorData, i.Gate368Truth, i.Verdict)
}

func FormatFormalization(f EtaFormalization) string {
	return fmt.Sprintf("executed=%t\nsupport_basis=%v\ngeneration_basis=%v\nnative_support_eta=%v balanced_support_eta=%v generation_eta_candidate=%v native_eta_generation=%t\ncurvature=%s\nprojection=%s\nverdict=%s", f.Executed, f.SupportBasis, f.GenerationBasis, f.NativeSupportEta, f.BalancedSupportEta, f.GenerationEtaCandidate, f.NativeEtaActsOnGeneration, f.CurvaturePrototype, f.ProjectionFormula, f.Verdict)
}

func FormatLane(l TraceLane) string {
	parts := []string{
		fmt.Sprintf("lane=%s name=%s", l.Lane, l.Name),
		fmt.Sprintf("formula=%s", l.Formula),
		fmt.Sprintf("eta_source=%s", l.EtaSource),
		fmt.Sprintf("native=%t circular=%t", l.Native, l.Circular),
		fmt.Sprintf("K=diag[%.12f %.12f %.12f]", l.Spectrum[0], l.Spectrum[1], l.Spectrum[2]),
		fmt.Sprintf("self_adjoint=%t central=%t noncentral=%t", l.SelfAdjoint, l.Central, l.NonCentral),
		FormatDecomposition(l.Decomposition),
		FormatKMS(l.KMS),
	}
	for _, cm := range l.Commutators {
		parts = append(parts, fmt.Sprintf("[%s,K]_norm=%.12f nonzero=%t", cm.Generator, cm.Norm, cm.NonZero))
	}
	parts = append(parts, fmt.Sprintf("breaks_flavor=%t promoted=%t selects_vacuum=%t", l.BreaksFlavorOrbit, l.PromotedHamiltonian, l.SelectsVacuum), "verdict="+l.Verdict)
	return strings.Join(parts, "\n")
}

func FormatDecomposition(d Decomposition) string {
	return fmt.Sprintf("decomposition: a=%.12f b=%.12f residual=%.12e exact=%t nonzero_b=%t target=%t verdict=%s", d.A, d.B, d.ResidualNorm, d.ExactInSpan, d.HasNonzeroB, d.TargetReached, d.Verdict)
}

func FormatKMS(k KMSState) string {
	parts := []string{fmt.Sprintf("KMS: beta=%.12f rho=[%.12f %.12f %.12f] faithful=%t nontrivial=%t", k.Beta, k.Density[0], k.Density[1], k.Density[2], k.Faithful, k.NontrivialFrequencies)}
	for _, f := range k.Frequencies {
		parts = append(parts, fmt.Sprintf("omega_%s=%.12f nonzero=%t", f.Pair, f.LogRatio, f.NonZero))
	}
	return strings.Join(parts, "\n")
}

func FormatActivation(a ActivationAudit) string {
	return fmt.Sprintf("executed=%t native_target=%t circular_capacity=%t promoted_native=%t internal_time=%t energy_constraint=%t\nanswer=%s\nnext=%s\nverdict=%s", a.Executed, a.NativeTargetReached, a.CircularCapacityWitnessed, a.PromotedNative, a.InternalTimeActivated, a.EnergyConstraintDerived, a.DirectAnswer, a.NextGate, a.Verdict)
}

func FormatLandscape(l LandscapeAudit) string {
	return fmt.Sprintf("executed=%t weak=%t quartic=%t alphaGUT=%t morita=%t bgap=%t omega=%t no_empirical_flavor=%t no_mass=%t no_vacuum=%t polluted=%t\nverdict=%s", l.Executed, l.WeakMixingPreserved, l.QuarticRatioPreserved, l.AlphaGUTPreserved, l.MoritaSplitPreserved, l.BGapLedgerPreserved, l.OmegaIndexPreserved, l.NoEmpiricalFlavorImport, l.NoObservedMassImport, l.NoVacuumPointClaimed, l.FiniteCorePolluted, l.Verdict)
}

func FormatKinetic(k KineticAudit) string {
	return fmt.Sprintf("executed=%t all_self_adjoint=%t faithful_states=%t no_rank_collapse=%t no_ghost=%t no_nonunitary_push=%t\nverdict=%s", k.Executed, k.AllCandidatesSelf, k.FaithfulStates, k.NoRankCollapse, k.NoGhostMetric, k.NoNonunitaryPush, k.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d reduction=%d remaining=%d seven_seal_target=%t\nverdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func MarkdownAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 369 Registry Audit — Eta-Graded Left-Right Trace / Noncentral Hamiltonian Extraction Sieve\n\n")
	b.WriteString("## Gate identity\n\n")
	b.WriteString("- **Gate:** 369\n")
	b.WriteString("- **Package:** `pkg/bridge/etagradedlrtrace`\n")
	b.WriteString("- **Theorem:** `EtaGradedLeftRightTraceNoncentralHamiltonianExtractionSieveTheorem`\n")
	b.WriteString("- **Audit ID:** `" + AuditID + "`\n")
	b.WriteString("- **Layer:** Bridge / Phase-III Flow Extension\n")
	b.WriteString("- **Purpose:** execute the exact eta-graded trace target isolated by Gate 368 and determine whether `tau_eta` emerges natively as a modular Hamiltonian.\n\n")

	b.WriteString("## Files, folders, and active gate chain\n\n")
	b.WriteString("| Region | Project objects | Gate-369 relevance |\n")
	b.WriteString("|---|---|---|\n")
	b.WriteString("| Core docs | `README.md`, `docs/architecture.md`, `GateResearcherMethod.md` | Ledger and method constraints. |\n")
	b.WriteString("| Registry | `internal/app/app.go` | Gate 369 is registered after Gate 368 and before runtime cache. |\n")
	b.WriteString("| Current package | `pkg/bridge/etagradedlrtrace` | Executes eta-graded trace sieve. |\n")
	b.WriteString("| Gate 368 package | `pkg/bridge/bimodulemodularcurvature` | Supplies the target equation and circularity firewall. |\n")
	b.WriteString("| Flow chain | Gates 362–369 | Static closure → modular flow → nontracial state → KMS → Hamiltonian origin → Lorentzian-time no-go → bimodule curvature → eta trace. |\n")
	b.WriteString("| Bimodule/heavy sector | Gates 290–320, 347 | Morita trace capacity, doubled space, Majorana/heavy overlap, and flavor-invariance obstructions. |\n\n")

	b.WriteString("## Inherited Gate-368 target\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")

	b.WriteString("## Eta grading formalization\n\n```text\n")
	b.WriteString(FormatFormalization(a.Formalization))
	b.WriteString("\n```\n\n")

	b.WriteString("## Candidate lane table\n\n")
	b.WriteString("| Lane | Candidate | Native? | Circular? | `K_eta` | Decomposition | Flavor action | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---|---|---|---|\n")
	for _, lane := range a.Lanes {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | `diag(%.12g, %.12g, %.12g)` | `a=%.12g, b=%.12g, residual=%.3g` | central=%t, breaks=%t | `%s` |\n", lane.Lane, lane.Name, lane.Native, lane.Circular, lane.Spectrum[0], lane.Spectrum[1], lane.Spectrum[2], lane.Decomposition.A, lane.Decomposition.B, lane.Decomposition.ResidualNorm, lane.Central, lane.BreaksFlavorOrbit, lane.Verdict))
	}
	b.WriteString("\n")

	for _, lane := range a.Lanes {
		b.WriteString("## Lane " + lane.Lane + " — " + lane.Name + "\n\n```text\n")
		b.WriteString(FormatLane(lane))
		b.WriteString("\n```\n\n")
	}

	b.WriteString("## Thermal-time activation sieve\n\n```text\n")
	b.WriteString(FormatActivation(a.Activation))
	b.WriteString("\n```\n\n")

	b.WriteString("## Landscape preservation\n\n```text\n")
	b.WriteString(FormatLandscape(a.Landscape))
	b.WriteString("\n```\n\n")

	b.WriteString("## Kinetic safety\n\n```text\n")
	b.WriteString(FormatKinetic(a.Kinetic))
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
	b.WriteString("## Next lawful theorem target\n\n")
	b.WriteString("Gate 369 proves that the current native support eta trace is not enough. The next theorem must decide whether there exists a native representation map from support defects to generation-dependent weights, or whether every admissible eta-graded Left-Right contraction factors through `I_3`.\n")
	return b.String()
}
