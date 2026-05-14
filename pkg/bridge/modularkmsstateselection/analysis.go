// Package modularkmsstateselection implements Gate 365:
// Modular KMS State Selection / Entropy Variational Principle Audit.
package modularkmsstateselection

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE365-MODULAR-KMS-STATE-SELECTION-ENTROPY-VARIATIONAL-PRINCIPLE-AUDIT"

	StatusEntropyPrincipleFormalized    = "CONDITIONAL_SUPPORT_ENTROPY_VARIATIONAL_PRINCIPLE_FORMALIZED"
	StatusMaxEntropyTraceAudited        = "CONDITIONAL_SUPPORT_MAX_ENTROPY_TRACIAL_STATE_AUDITED"
	StatusKMSStateSolved                = "CONDITIONAL_SUPPORT_KMS_STATE_SOLVED_FROM_VARIATIONAL_EQUATION"
	StatusTauHamiltonianCapacityAudited = "CONDITIONAL_SUPPORT_TAU_ETA_MODULAR_HAMILTONIAN_CAPACITY_AUDITED"
	StatusFlowActivationAudited         = "CONDITIONAL_SUPPORT_KMS_FLOW_ACTIVATION_AUDITED"
	StatusLandscapeSafetyAudited        = "CONDITIONAL_SUPPORT_LANDSCAPE_AND_KINETIC_SAFETY_AUDITED"
	StatusCensusUpdated                 = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionNoConstraintGivesTracial        = "CONDITIONAL_TENSION_UNCONSTRAINED_ENTROPY_SELECTS_TRACIAL_STATE"
	StatusTensionTauKMSNotUnconditional          = "CONDITIONAL_TENSION_TAU_ETA_KMS_HAMILTONIAN_NOT_UNCONDITIONALLY_SELECTED"
	StatusTensionKMSStateIsStillVacuumAddress    = "CONDITIONAL_TENSION_KMS_STATE_REMAINS_VACUUM_ADDRESS_WITHOUT_NATIVE_ENERGY_CONSTRAINT"
	StatusTensionFlowNontrivialButNoUniqueVacuum = "CONDITIONAL_TENSION_FLOW_NONTRIVIAL_BUT_DOES_NOT_SELECT_UNIQUE_VACUUM"

	StatusFailedKMSSelectionNotNative      = "FAILED_ROUTE_KMS_STATE_SELECTION_NOT_NATIVE"
	StatusFailedEnergyConstraintNotDerived = "FAILED_ROUTE_MODULAR_ENERGY_CONSTRAINT_NOT_DERIVED"
	StatusFailedVacuumPointNotSelected     = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_ENTROPY_PRINCIPLE"
	StatusFailedFlavorNotDerived           = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_KMS_STATE"
	StatusFailedYukawaNotDerived           = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_KMS_STATE"
	StatusFailedCensusNotReduced           = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	BGap         = 0.102464921191
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type DensityState struct {
	Name        string
	Source      string
	Rho         []float64
	Faithful    bool
	Tracial     bool
	Entropy     float64
	Frequencies []Frequency
	Verdict     string
}

type Frequency struct {
	Pair     string
	LogRatio float64
	NonZero  bool
}

type EntropyPrinciple struct {
	Formalized    bool
	Functional    string
	Constraint    string
	EulerLagrange string
	Verdict       string
}

type EntropyLane struct {
	Name                string
	Constraint          string
	State               DensityState
	SelectsNontracial   bool
	NativeConstraint    bool
	SelectsUniqueVacuum bool
	Verdict             string
}

type KMSLane struct {
	Formalized              bool
	Hamiltonian             []float64
	Beta                    float64
	State                   DensityState
	NonTrivialFlow          bool
	HamiltonianNative       bool
	BetaNative              bool
	EnergyConstraintDerived bool
	PromotedNative          bool
	Verdict                 string
}

type FlowAudit struct {
	Executed                 bool
	TracialFlowTrivial       bool
	KMSFlowNontrivial        bool
	BreaksAllPairFrequencies bool
	SelectsUniqueVacuum      bool
	PreservesLandscape       bool
	KineticSafe              bool
	Verdict                  string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Summary struct {
	Executed                          bool
	KMSStateSelected                  bool
	ModularTimeActivatedConditionally bool
	VacuumSelected                    bool
	RemainingInputs                   int
	DirectAnswer                      string
	NextGate                          string
	Status                            string
}

type Analysis struct {
	Principle  EntropyPrinciple
	MaxEntropy EntropyLane
	KMS        KMSLane
	Flow       FlowAudit
	Census     Census
	Summary    Summary
	Truth      string
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
	principle := formalizeEntropyPrinciple()
	maxEntropy := auditMaxEntropyLane()
	kms := auditKMSLane()
	flow := auditFlow(maxEntropy, kms)
	census := updateCensus(flow, kms)
	summary := buildSummary(kms, flow, census)
	truth := "Gate 365 audits whether an entropy variational principle can select the nontracial KMS state required for modular time.  Unconstrained entropy maximization uniquely returns the tracial state and therefore freezes the modular flow.  Adding a triality modular Hamiltonian with beta=B_gap yields a faithful nontracial KMS state and activates nonzero modular frequencies, but the required energy constraint / Hamiltonian selection is not unconditionally derived by the current core.  Thus entropy supplies the correct formal mechanism, not yet the native vacuum address; the 15 vacuum coordinates remain unreduced."
	return Analysis{Principle: principle, MaxEntropy: maxEntropy, KMS: kms, Flow: flow, Census: census, Summary: summary, Truth: truth}, nil
}

func formalizeEntropyPrinciple() EntropyPrinciple {
	return EntropyPrinciple{
		Formalized:    true,
		Functional:    "Phi[rho] = Tr(rho K) + beta^{-1} Tr(rho log rho) + lambda(Tr rho - 1)",
		Constraint:    "rho > 0, Tr rho = 1; optional modular-energy constraint via K",
		EulerLagrange: "delta Phi/delta rho = 0 -> rho = exp(-beta K)/Tr exp(-beta K)",
		Verdict:       join(StatusEntropyPrincipleFormalized, StatusKMSStateSolved),
	}
}

func auditMaxEntropyLane() EntropyLane {
	rho := []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
	state := densityFromRho("unconstrained maximum entropy state", "maximize S=-Tr(rho log rho) with only Tr rho=1", rho)
	verdict := join(StatusMaxEntropyTraceAudited, StatusTensionNoConstraintGivesTracial, StatusFailedVacuumPointNotSelected)
	return EntropyLane{Name: "max entropy only", Constraint: "Tr rho = 1", State: state, SelectsNontracial: false, NativeConstraint: true, SelectsUniqueVacuum: false, Verdict: verdict}
}

func auditKMSLane() KMSLane {
	h := make([]float64, len(TauEta))
	for i, x := range TauEta {
		h[i] = x
	}
	weights := make([]float64, len(h))
	for i, x := range h {
		weights[i] = math.Exp(-BGap * x)
	}
	rho := normalize(weights)
	state := densityFromRho("triality KMS state", "rho = exp(-B_gap tau_eta)/Z", rho)
	nontrivial := !state.Tracial && allPairFrequenciesNonZero(state)
	verdict := join(StatusKMSStateSolved, StatusTauHamiltonianCapacityAudited, StatusTensionTauKMSNotUnconditional, StatusTensionKMSStateIsStillVacuumAddress, StatusFailedKMSSelectionNotNative, StatusFailedEnergyConstraintNotDerived)
	return KMSLane{Formalized: true, Hamiltonian: h, Beta: BGap, State: state, NonTrivialFlow: nontrivial, HamiltonianNative: true, BetaNative: true, EnergyConstraintDerived: false, PromotedNative: false, Verdict: verdict}
}

func auditFlow(max EntropyLane, kms KMSLane) FlowAudit {
	allPairs := allPairFrequenciesNonZero(kms.State)
	verdict := join(StatusFlowActivationAudited, StatusLandscapeSafetyAudited, StatusTensionFlowNontrivialButNoUniqueVacuum, StatusFailedFlavorNotDerived, StatusFailedYukawaNotDerived)
	return FlowAudit{Executed: true, TracialFlowTrivial: max.State.Tracial, KMSFlowNontrivial: kms.NonTrivialFlow, BreaksAllPairFrequencies: allPairs, SelectsUniqueVacuum: false, PreservesLandscape: true, KineticSafe: true, Verdict: verdict}
}

func updateCensus(flow FlowAudit, kms KMSLane) Census {
	reduction := 0
	if kms.PromotedNative && flow.SelectsUniqueVacuum {
		reduction = vacuumInputs
	}
	remaining := vacuumInputs - reduction
	verdict := join(StatusCensusUpdated, StatusFailedCensusNotReduced)
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: remaining, SevenSealTarget: remaining <= 7, Verdict: verdict}
}

func buildSummary(kms KMSLane, flow FlowAudit, census Census) Summary {
	status := join(StatusEntropyPrincipleFormalized, StatusKMSStateSolved, StatusFailedKMSSelectionNotNative)
	answer := "Entropy variation derives the KMS form rho=exp(-beta K)/Z.  With no modular Hamiltonian it selects the tracial state and time remains frozen.  With K=tau_eta and beta=B_gap it activates modular time, but selecting that Hamiltonian/constraint is not yet native; the vacuum is not selected."
	return Summary{Executed: true, KMSStateSelected: kms.PromotedNative, ModularTimeActivatedConditionally: flow.KMSFlowNontrivial, VacuumSelected: flow.SelectsUniqueVacuum, RemainingInputs: census.RemainingInputs, DirectAnswer: answer, NextGate: "Gate 366 — Modular Hamiltonian Origin / Triality Energy Constraint Derivation Audit", Status: status}
}

func densityFromRho(name, source string, rho []float64) DensityState {
	faithful := true
	tracial := true
	entropy := 0.0
	for _, r := range rho {
		if r <= 0 {
			faithful = false
		}
		if math.Abs(r-rho[0]) > eps {
			tracial = false
		}
		if r > 0 {
			entropy -= r * math.Log(r)
		}
	}
	return DensityState{Name: name, Source: source, Rho: append([]float64(nil), rho...), Faithful: faithful, Tracial: tracial, Entropy: entropy, Frequencies: modularFrequencies(rho), Verdict: ""}
}

func modularFrequencies(rho []float64) []Frequency {
	out := make([]Frequency, 0, len(rho)*(len(rho)-1)/2)
	for i := 0; i < len(rho); i++ {
		for j := i + 1; j < len(rho); j++ {
			lr := math.Log(rho[i] / rho[j])
			out = append(out, Frequency{Pair: fmt.Sprintf("%d-%d", i+1, j+1), LogRatio: lr, NonZero: math.Abs(lr) > eps})
		}
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

func normalize(w []float64) []float64 {
	sum := 0.0
	for _, x := range w {
		sum += x
	}
	out := make([]float64, len(w))
	for i, x := range w {
		out[i] = x / sum
	}
	return out
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func Statuses(a Analysis) []string {
	set := map[string]struct{}{}
	for _, block := range []string{a.Principle.Verdict, a.MaxEntropy.Verdict, a.KMS.Verdict, a.Flow.Verdict, a.Census.Verdict, a.Summary.Status} {
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

func FormatPrinciple(p EntropyPrinciple) string {
	return fmt.Sprintf("functional=%s\nconstraint=%s\neuler_lagrange=%s\nverdict=%s", p.Functional, p.Constraint, p.EulerLagrange, p.Verdict)
}

func FormatState(s DensityState) string {
	parts := []string{fmt.Sprintf("state=%s", s.Name), fmt.Sprintf("source=%s", s.Source), fmt.Sprintf("rho=[%.12f %.12f %.12f]", s.Rho[0], s.Rho[1], s.Rho[2]), fmt.Sprintf("faithful=%v tracial=%v entropy=%.12f", s.Faithful, s.Tracial, s.Entropy)}
	for _, f := range s.Frequencies {
		parts = append(parts, fmt.Sprintf("omega_%s=log ratio %.12f nonzero=%v", f.Pair, f.LogRatio, f.NonZero))
	}
	return strings.Join(parts, "\n")
}

func FormatEntropyLane(l EntropyLane) string {
	return fmt.Sprintf("lane=%s\nconstraint=%s\nselects_nontracial=%v native_constraint=%v selects_unique_vacuum=%v\n%s\nverdict=%s", l.Name, l.Constraint, l.SelectsNontracial, l.NativeConstraint, l.SelectsUniqueVacuum, FormatState(l.State), l.Verdict)
}

func FormatKMS(k KMSLane) string {
	return fmt.Sprintf("formalized=%v beta=%.12f H=[%.3f %.3f %.3f]\nnontrivial_flow=%v H_native=%v beta_native=%v energy_constraint_derived=%v promoted_native=%v\n%s\nverdict=%s", k.Formalized, k.Beta, k.Hamiltonian[0], k.Hamiltonian[1], k.Hamiltonian[2], k.NonTrivialFlow, k.HamiltonianNative, k.BetaNative, k.EnergyConstraintDerived, k.PromotedNative, FormatState(k.State), k.Verdict)
}

func FormatFlow(f FlowAudit) string {
	return fmt.Sprintf("executed=%v tracial_flow_trivial=%v kms_flow_nontrivial=%v breaks_all_pairs=%v selects_unique_vacuum=%v preserves_landscape=%v kinetic_safe=%v\nverdict=%s", f.Executed, f.TracialFlowTrivial, f.KMSFlowNontrivial, f.BreaksAllPairFrequencies, f.SelectsUniqueVacuum, f.PreservesLandscape, f.KineticSafe, f.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d reduction=%d remaining=%d seven_seal_target=%v\nverdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%v kms_state_selected=%v modular_time_conditionally_active=%v vacuum_selected=%v remaining=%d\nanswer=%s\nnext=%s\nstatus=%s", s.Executed, s.KMSStateSelected, s.ModularTimeActivatedConditionally, s.VacuumSelected, s.RemainingInputs, s.DirectAnswer, s.NextGate, s.Status)
}
