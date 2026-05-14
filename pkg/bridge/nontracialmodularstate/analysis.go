// Package nontracialmodularstate implements Gate 364:
// Nontracial Modular State Origin / Vacuum Density Matrix Derivation Audit.
package nontracialmodularstate

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE364-NONTRACIAL-MODULAR-STATE-ORIGIN-VACUUM-DENSITY-MATRIX-DERIVATION-AUDIT"

	StatusTopologicalSourcingFormalized = "CONDITIONAL_SUPPORT_TOPOLOGICAL_SOURCING_AUDIT_FORMALIZED"
	StatusTauMagnitudeStateAudited      = "CONDITIONAL_SUPPORT_TAU_MAGNITUDE_DENSITY_STATE_AUDITED"
	StatusKMSStateFormalized            = "CONDITIONAL_SUPPORT_EXPONENTIAL_KMS_STATE_FORMALIZED"
	StatusFlowActivationSieveExecuted   = "CONDITIONAL_SUPPORT_FLOW_ACTIVATION_SIEVE_EXECUTED"
	StatusModularTimeCapacityIdentified = "CONDITIONAL_SUPPORT_MODULAR_TIME_ACTIVATION_CAPACITY_IDENTIFIED"
	StatusParameterCensusUpdated        = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionTauMagnitudeResidualDegeneracy = "CONDITIONAL_TENSION_TAU_MAGNITUDE_STATE_HAS_RESIDUAL_12_DEGENERACY"
	StatusTensionKMSMapNotMandated              = "CONDITIONAL_TENSION_KMS_DENSITY_MAP_NOT_MANDATED_BY_CURRENT_CORE"
	StatusTensionNontracialCandidateIsAddress   = "CONDITIONAL_TENSION_NONTRACIAL_DENSITY_CANDIDATE_IS_STILL_VACUUM_ADDRESS"
	StatusTensionFlowActivatesButNotSelects     = "CONDITIONAL_TENSION_FLOW_ACTIVATES_TIME_BUT_DOES_NOT_SELECT_UNIQUE_VACUUM"

	StatusFailedNonTracialStateNotDerived = "FAILED_ROUTE_NONTRACIAL_STATE_NOT_NATIVELY_DERIVED"
	StatusFailedModularTimeNotActivated   = "FAILED_ROUTE_MODULAR_TIME_NOT_ACTIVATED_UNCONDITIONALLY"
	StatusFailedDensityMandateMissing     = "FAILED_ROUTE_TOPOLOGY_TO_DENSITY_MATRIX_MAP_NOT_DERIVED"
	StatusFailedFlavorVacuumNotSelected   = "FAILED_ROUTE_FLAVOR_VACUUM_NOT_SELECTED_BY_DENSITY_FLOW"
	StatusFailedCensusNotReduced          = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
	StatusFailedCKMNotDerived             = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_NONTRACIAL_STATE"
)

const (
	vacuumInputs = 15
	BGap         = 0.102464921191
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type DensityState struct {
	Name          string
	Source        string
	Rho           []float64
	Faithful      bool
	Tracial       bool
	NativeData    bool
	MandatedMap   bool
	ResidualDeg12 bool
	Frequencies   []Frequency
	Verdict       string
}

type Frequency struct {
	Element  string
	Ratio    float64
	LogRatio float64
	NonZero  bool
}

type TopologicalSourcing struct {
	Formalized            bool
	SignedTauDensityValid bool
	MagnitudeState        DensityState
	SquaredMagnitudeState DensityState
	NativeNonTracialFound bool
	Reason                string
	Verdict               string
}

type KMSAudit struct {
	Formalized  bool
	Hamiltonian string
	Beta        float64
	State       DensityState
	NonTrivial  bool
	Mandated    bool
	Verdict     string
}

type FlowActivation struct {
	Executed                    bool
	CandidateStates             []DensityState
	AnyNonTrivial               bool
	AnyMandatedNativeNontracial bool
	BreaksAllPairFrequencies    bool
	SelectsUniqueVacuum         bool
	RemainingInputs             int
	Verdict                     string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Summary struct {
	Executed        bool
	NonTracialFound bool
	TimeActivated   bool
	VacuumSelected  bool
	RemainingInputs int
	DirectAnswer    string
	NextGate        string
	Status          string
}

type Analysis struct {
	Topological TopologicalSourcing
	KMS         KMSAudit
	Flow        FlowActivation
	Census      Census
	Summary     Summary
	Truth       string
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
	topo := auditTopologicalSourcing()
	kms := auditKMSState()
	flow := executeFlowActivation(topo, kms)
	census := updateCensus(flow)
	summary := buildSummary(flow, census)
	truth := "Gate 364 audits whether ASHA's native generation topology can produce the nontracial faithful modular state required by Gate 363.  Tau-derived and KMS-like density candidates activate nontrivial modular frequencies, but the current core does not mandate a unique topology-to-density map.  Tau magnitude/square states retain a 1-2 degeneracy, while the sign-sensitive KMS lane requires choosing K=B_gap tau_eta as the modular Hamiltonian.  Therefore modular time has capacity but is not unconditionally activated, and the 15 vacuum coordinates remain quarantined."
	return Analysis{Topological: topo, KMS: kms, Flow: flow, Census: census, Summary: summary, Truth: truth}, nil
}

func auditTopologicalSourcing() TopologicalSourcing {
	mag := make([]float64, len(TauEta))
	sq := make([]float64, len(TauEta))
	for i, x := range TauEta {
		mag[i] = math.Abs(x)
		sq[i] = x * x
	}
	magState := stateFromWeights("tau magnitude density", "rho_i proportional to |tau_i|", mag, true, false)
	sqState := stateFromWeights("tau squared-magnitude density", "rho_i proportional to tau_i^2", sq, true, false)
	signedValid := true
	for _, x := range TauEta {
		if x < 0 {
			signedValid = false
		}
	}
	verdict := join(StatusTopologicalSourcingFormalized, StatusTauMagnitudeStateAudited, StatusTensionTauMagnitudeResidualDegeneracy, StatusFailedDensityMandateMissing)
	return TopologicalSourcing{
		Formalized:            true,
		SignedTauDensityValid: signedValid,
		MagnitudeState:        magState,
		SquaredMagnitudeState: sqState,
		NativeNonTracialFound: false,
		Reason:                "tau_eta is native and non-democratic, but converting it into rho requires an additional positive-state prescription.  Signed tau is not a density because it has a negative component; |tau| and tau^2 are positive but keep rho_1=rho_2 and therefore do not fully break generation degeneracy.",
		Verdict:               verdict,
	}
}

func auditKMSState() KMSAudit {
	weights := make([]float64, len(TauEta))
	// KMS-like state rho ∝ exp(-B_gap tau_eta).  This is a capacity witness: it uses derived data,
	// but the rule K=B_gap tau_eta is not mandated by the current ASHA core.
	for i, x := range TauEta {
		weights[i] = math.Exp(-BGap * x)
	}
	state := stateFromWeights("KMS tau_eta density", "rho proportional to exp(-B_gap tau_eta)", weights, true, false)
	nontrivial := !state.Tracial
	verdict := join(StatusKMSStateFormalized, StatusModularTimeCapacityIdentified, StatusTensionKMSMapNotMandated, StatusTensionNontracialCandidateIsAddress, StatusFailedNonTracialStateNotDerived)
	return KMSAudit{Formalized: true, Hamiltonian: "K_flow = B_gap * diag(tau_eta)", Beta: 1, State: state, NonTrivial: nontrivial, Mandated: false, Verdict: verdict}
}

func executeFlowActivation(topo TopologicalSourcing, kms KMSAudit) FlowActivation {
	states := []DensityState{topo.MagnitudeState, topo.SquaredMagnitudeState, kms.State}
	anyNonTrivial := false
	anyMandated := false
	breaksAll := false
	for _, s := range states {
		if !s.Tracial {
			anyNonTrivial = true
		}
		if s.MandatedMap && !s.Tracial {
			anyMandated = true
		}
		if allPairFrequenciesNonZero(s) {
			breaksAll = true
		}
	}
	verdict := join(StatusFlowActivationSieveExecuted, StatusModularTimeCapacityIdentified, StatusTensionFlowActivatesButNotSelects, StatusFailedModularTimeNotActivated, StatusFailedFlavorVacuumNotSelected, StatusFailedCKMNotDerived)
	if !anyMandated {
		verdict = join(verdict, StatusFailedNonTracialStateNotDerived)
	}
	return FlowActivation{Executed: true, CandidateStates: states, AnyNonTrivial: anyNonTrivial, AnyMandatedNativeNontracial: anyMandated, BreaksAllPairFrequencies: breaksAll, SelectsUniqueVacuum: false, RemainingInputs: vacuumInputs, Verdict: verdict}
}

func updateCensus(flow FlowActivation) Census {
	reduction := 0
	if flow.AnyMandatedNativeNontracial && flow.SelectsUniqueVacuum {
		reduction = 1 // Not reached: reserved for future native density theorem.
	}
	remaining := vacuumInputs - reduction
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: remaining, SevenSealTarget: remaining <= 7, Verdict: join(StatusParameterCensusUpdated, StatusFailedCensusNotReduced)}
}

func buildSummary(flow FlowActivation, census Census) Summary {
	status := join(StatusTopologicalSourcingFormalized, StatusKMSStateFormalized, StatusFlowActivationSieveExecuted, StatusParameterCensusUpdated, StatusFailedNonTracialStateNotDerived, StatusFailedFlavorVacuumNotSelected)
	answer := "Topology supplies nontracial density candidates and KMS capacity, but the current finite core does not mandate a unique faithful nontracial modular state.  Modular time can be activated conditionally, yet the density matrix remains a vacuum-address input rather than a derived selector."
	return Summary{Executed: true, NonTracialFound: flow.AnyMandatedNativeNontracial, TimeActivated: flow.AnyMandatedNativeNontracial, VacuumSelected: flow.SelectsUniqueVacuum, RemainingInputs: census.RemainingInputs, DirectAnswer: answer, NextGate: "Gate 365 — Modular KMS State Selection / Entropy Variational Principle Audit", Status: status}
}

func stateFromWeights(name, source string, weights []float64, nativeData, mandated bool) DensityState {
	sum := 0.0
	faithful := true
	for _, w := range weights {
		if w <= 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			faithful = false
		}
		sum += w
	}
	rho := make([]float64, len(weights))
	if faithful && sum > 0 {
		for i, w := range weights {
			rho[i] = w / sum
		}
	}
	tracial := true
	for i := 1; i < len(rho); i++ {
		if math.Abs(rho[i]-rho[0]) > eps {
			tracial = false
			break
		}
	}
	residual12 := len(rho) >= 2 && math.Abs(rho[0]-rho[1]) < eps
	verdict := StatusTopologicalSourcingFormalized
	if !tracial {
		verdict = join(verdict, StatusModularTimeCapacityIdentified)
	}
	if residual12 {
		verdict = join(verdict, StatusTensionTauMagnitudeResidualDegeneracy)
	}
	if !mandated {
		verdict = join(verdict, StatusTensionNontracialCandidateIsAddress)
	}
	return DensityState{Name: name, Source: source, Rho: rho, Faithful: faithful, Tracial: tracial, NativeData: nativeData, MandatedMap: mandated, ResidualDeg12: residual12, Frequencies: frequencies(rho), Verdict: verdict}
}

func frequencies(rho []float64) []Frequency {
	labels := []string{"E12", "E13", "E23"}
	pairs := [][2]int{{0, 1}, {0, 2}, {1, 2}}
	out := make([]Frequency, 0, len(pairs))
	for k, pair := range pairs {
		r := rho[pair[0]] / rho[pair[1]]
		l := math.Log(r)
		out = append(out, Frequency{Element: labels[k], Ratio: r, LogRatio: l, NonZero: math.Abs(l) > eps})
	}
	return out
}

func allPairFrequenciesNonZero(s DensityState) bool {
	if len(s.Frequencies) == 0 {
		return false
	}
	for _, f := range s.Frequencies {
		if !f.NonZero {
			return false
		}
	}
	return true
}

func Statuses(a Analysis) []string {
	var out []string
	add := func(s string) {
		for _, part := range strings.Split(s, "|") {
			part = strings.TrimSpace(part)
			if part != "" {
				out = append(out, part)
			}
		}
	}
	add(a.Topological.MagnitudeState.Verdict)
	add(a.Topological.SquaredMagnitudeState.Verdict)
	add(a.Topological.Verdict)
	add(a.KMS.State.Verdict)
	add(a.KMS.Verdict)
	add(a.Flow.Verdict)
	add(a.Census.Verdict)
	add(a.Summary.Status)
	sort.Strings(out)
	dedup := out[:0]
	for _, s := range out {
		if len(dedup) == 0 || dedup[len(dedup)-1] != s {
			dedup = append(dedup, s)
		}
	}
	return dedup
}

func FormatFrequency(fs []Frequency) string {
	parts := make([]string, 0, len(fs))
	for _, f := range fs {
		parts = append(parts, fmt.Sprintf("%s ratio=%.12g log=%.12g nonzero=%t", f.Element, f.Ratio, f.LogRatio, f.NonZero))
	}
	return strings.Join(parts, "; ")
}

func FormatState(s DensityState) string {
	return fmt.Sprintf("%s source=%q rho=%v faithful=%t tracial=%t nativeData=%t mandated=%t residual12=%t freqs=[%s] verdict=%s", s.Name, s.Source, s.Rho, s.Faithful, s.Tracial, s.NativeData, s.MandatedMap, s.ResidualDeg12, FormatFrequency(s.Frequencies), s.Verdict)
}

func FormatTopological(t TopologicalSourcing) string {
	return fmt.Sprintf("formalized=%t signedDensityValid=%t nativeNontracial=%t magnitude={%s} squared={%s} reason=%q verdict=%s", t.Formalized, t.SignedTauDensityValid, t.NativeNonTracialFound, FormatState(t.MagnitudeState), FormatState(t.SquaredMagnitudeState), t.Reason, t.Verdict)
}

func FormatKMS(k KMSAudit) string {
	return fmt.Sprintf("formalized=%t H=%q beta=%.12g nontrivial=%t mandated=%t state={%s} verdict=%s", k.Formalized, k.Hamiltonian, k.Beta, k.NonTrivial, k.Mandated, FormatState(k.State), k.Verdict)
}

func FormatFlow(f FlowActivation) string {
	parts := make([]string, 0, len(f.CandidateStates))
	for _, s := range f.CandidateStates {
		parts = append(parts, FormatState(s))
	}
	return fmt.Sprintf("executed=%t anyNontrivial=%t mandated=%t breaksAll=%t selected=%t remaining=%d states=[%s] verdict=%s", f.Executed, f.AnyNonTrivial, f.AnyMandatedNativeNontracial, f.BreaksAllPairFrequencies, f.SelectsUniqueVacuum, f.RemainingInputs, strings.Join(parts, " || "), f.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d reduction=%d remaining=%d seven=%t verdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t nontracialFound=%t timeActivated=%t vacuumSelected=%t remaining=%d next=%q answer=%q status=%s", s.Executed, s.NonTracialFound, s.TimeActivated, s.VacuumSelected, s.RemainingInputs, s.NextGate, s.DirectAnswer, s.Status)
}

func join(parts ...string) string {
	var out []string
	for _, p := range parts {
		for _, item := range strings.Split(p, "|") {
			item = strings.TrimSpace(item)
			if item != "" {
				out = append(out, item)
			}
		}
	}
	return strings.Join(out, "|")
}
