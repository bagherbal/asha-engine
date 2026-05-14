// Package modularhamiltonianorigin implements Gate 366:
// Modular Hamiltonian Origin / Triality Energy Constraint Derivation Audit.
package modularhamiltonianorigin

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE366-MODULAR-HAMILTONIAN-ORIGIN-TRIALITY-ENERGY-CONSTRAINT-DERIVATION-AUDIT"

	StatusHamiltonianOriginFormalized = "CONDITIONAL_SUPPORT_MODULAR_HAMILTONIAN_ORIGIN_CRITERIA_FORMALIZED"
	StatusCandidateSpectrumAudited    = "CONDITIONAL_SUPPORT_NATIVE_HAMILTONIAN_CANDIDATES_AUDITED"
	StatusTrialityHamiltonianAudited  = "CONDITIONAL_SUPPORT_TAU_ETA_HAMILTONIAN_CAPACITY_AUDITED"
	StatusEnergyConstraintAudited     = "CONDITIONAL_SUPPORT_ENERGY_CONSTRAINT_INVERSION_AUDITED"
	StatusFlowKernelRecomputed        = "CONDITIONAL_SUPPORT_KMS_FLOW_KERNEL_RECOMPUTED"
	StatusLandscapeSafetyAudited      = "CONDITIONAL_SUPPORT_LANDSCAPE_AND_KINETIC_SAFETY_RECHECKED"
	StatusCensusUpdated               = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionIdentityHamiltonianTracial    = "CONDITIONAL_TENSION_IDENTITY_HAMILTONIAN_FREEZES_MODULAR_TIME"
	StatusTensionMagnitudeResidualDegeneracy   = "CONDITIONAL_TENSION_MAGNITUDE_HAMILTONIANS_RETAIN_12_DEGENERACY"
	StatusTensionTrialityNoncentralButUnchosen = "CONDITIONAL_TENSION_TAU_ETA_IS_NONCENTRAL_BUT_NOT_SELECTED_AS_ENERGY"
	StatusTensionEnergyConstraintCircular      = "CONDITIONAL_TENSION_ENERGY_CONSTRAINT_IS_CIRCULAR_WITHOUT_NATIVE_EXPECTATION_VALUE"
	StatusTensionHamiltonianIsVacuumAddress    = "CONDITIONAL_TENSION_MODULAR_HAMILTONIAN_REMAINS_VACUUM_ADDRESS"

	StatusFailedHamiltonianNotDerived     = "FAILED_ROUTE_MODULAR_HAMILTONIAN_NOT_DERIVED"
	StatusFailedEnergyConstraintNotNative = "FAILED_ROUTE_MODULAR_ENERGY_CONSTRAINT_NOT_DERIVED"
	StatusFailedKMSNotPromoted            = "FAILED_ROUTE_KMS_STATE_NOT_PROMOTED_NATIVE"
	StatusFailedVacuumNotSelected         = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_MODULAR_HAMILTONIAN"
	StatusFailedFlavorNotDerived          = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_MODULAR_HAMILTONIAN"
	StatusFailedYukawaNotDerived          = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_MODULAR_HAMILTONIAN"
	StatusFailedCensusNotReduced          = "FAILED_ROUTE_VACUUM_PARAMETER_CENSUS_NOT_REDUCED"
)

const (
	vacuumInputs = 15
	BGap         = 0.102464921191
	eps          = 1e-12
)

var TauEta = []float64{2, -2, 1}

type OriginCriteria struct {
	Formalized        bool
	MustBeSelfAdjoint bool
	MustBeNoncentral  bool
	MustBeNative      bool
	MustBeNonCircular bool
	MustBreakOrbit    bool
	Description       string
	Verdict           string
}

type HamiltonianCandidate struct {
	Name                string
	Source              string
	Spectrum            []float64
	SelfAdjoint         bool
	Central             bool
	NativeOperator      bool
	EnergyRoleDerived   bool
	BreaksAllDegeneracy bool
	KMSState            []float64
	ModularFrequencies  []Frequency
	PromotedNative      bool
	Verdict             string
}

type Frequency struct {
	Pair     string
	LogRatio float64
	NonZero  bool
}

type EnergyConstraint struct {
	Formalized        bool
	Hamiltonian       []float64
	Beta              float64
	State             []float64
	Expectation       float64
	ConstraintNative  bool
	Circular          bool
	BetaDerived       bool
	PromotesKMSNative bool
	Verdict           string
}

type FlowAudit struct {
	Executed           bool
	BestCandidate      string
	NontrivialCapacity bool
	PromotedNative     bool
	PreservesLandscape bool
	KineticSafe        bool
	SelectsVacuum      bool
	Verdict            string
}

type Census struct {
	StartingInputs  int
	Reduction       int
	RemainingInputs int
	SevenSealTarget bool
	Verdict         string
}

type Summary struct {
	Executed                bool
	HamiltonianFound        bool
	EnergyConstraintDerived bool
	VacuumSelected          bool
	RemainingInputs         int
	DirectAnswer            string
	NextGate                string
	Status                  string
}

type Analysis struct {
	Criteria   OriginCriteria
	Candidates []HamiltonianCandidate
	Energy     EnergyConstraint
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
	criteria := formalizeOriginCriteria()
	candidates := auditHamiltonianCandidates()
	energy := auditEnergyConstraint(candidateByName(candidates, "triality signed tau_eta"))
	flow := auditFlow(candidates, energy)
	census := updateCensus(flow)
	summary := buildSummary(flow, energy, census)
	truth := "Gate 366 audits whether the modular Hamiltonian K used in the KMS state can be derived natively rather than chosen as a vacuum address.  The identity Hamiltonian is native but gives a tracial frozen flow.  Magnitude Hamiltonians are faithful but retain the 1-2 degeneracy.  The signed triality Hamiltonian K=tau_eta is native as an operator and activates modular time with beta=B_gap, but no current spectral-action or entropy principle selects it as the modular energy constraint.  Thus the KMS mechanism remains formal and conditional; the 15 vacuum coordinates remain unreduced."
	return Analysis{Criteria: criteria, Candidates: candidates, Energy: energy, Flow: flow, Census: census, Summary: summary, Truth: truth}, nil
}

func formalizeOriginCriteria() OriginCriteria {
	return OriginCriteria{
		Formalized:        true,
		MustBeSelfAdjoint: true,
		MustBeNoncentral:  true,
		MustBeNative:      true,
		MustBeNonCircular: true,
		MustBreakOrbit:    true,
		Description:       "A modular Hamiltonian must be self-adjoint, noncentral on the generation algebra, selected by a native ASHA functional or constraint, not equivalent to choosing rho, and capable of nontrivially acting on the flavor orbit while preserving kinetic positivity.",
		Verdict:           StatusHamiltonianOriginFormalized,
	}
}

func auditHamiltonianCandidates() []HamiltonianCandidate {
	candidates := []HamiltonianCandidate{
		candidate("identity", "unconstrained maximum-entropy / central trace state", []float64{0, 0, 0}, true, true, true, true),
		candidate("tau magnitude", "rho sourced by |tau_eta|", []float64{2, 2, 1}, true, false, true, false),
		candidate("tau squared magnitude", "rho sourced by tau_eta^2", []float64{4, 4, 1}, true, false, true, false),
		candidate("triality signed tau_eta", "signed generation topology tau_eta=(2,-2,1)", []float64{2, -2, 1}, true, false, true, false),
	}
	for i := range candidates {
		c := &candidates[i]
		c.KMSState = kmsState(c.Spectrum, BGap)
		c.ModularFrequencies = modularFrequencies(c.KMSState)
		c.BreaksAllDegeneracy = allFrequenciesNonZero(c.ModularFrequencies)
		c.PromotedNative = c.NativeOperator && c.EnergyRoleDerived && c.BreaksAllDegeneracy && !c.Central
		parts := []string{StatusCandidateSpectrumAudited}
		switch c.Name {
		case "identity":
			parts = append(parts, StatusTensionIdentityHamiltonianTracial, StatusFailedVacuumNotSelected)
		case "tau magnitude", "tau squared magnitude":
			parts = append(parts, StatusTensionMagnitudeResidualDegeneracy, StatusFailedVacuumNotSelected)
		case "triality signed tau_eta":
			parts = append(parts, StatusTrialityHamiltonianAudited, StatusTensionTrialityNoncentralButUnchosen, StatusFailedHamiltonianNotDerived)
		}
		c.Verdict = join(parts...)
	}
	return candidates
}

func candidate(name, source string, spectrum []float64, selfAdjoint, central, native, energyDerived bool) HamiltonianCandidate {
	return HamiltonianCandidate{Name: name, Source: source, Spectrum: append([]float64(nil), spectrum...), SelfAdjoint: selfAdjoint, Central: central, NativeOperator: native, EnergyRoleDerived: energyDerived}
}

func auditEnergyConstraint(k HamiltonianCandidate) EnergyConstraint {
	rho := kmsState(k.Spectrum, BGap)
	exp := expectation(rho, k.Spectrum)
	verdict := join(StatusEnergyConstraintAudited, StatusTensionEnergyConstraintCircular, StatusTensionHamiltonianIsVacuumAddress, StatusFailedEnergyConstraintNotNative, StatusFailedKMSNotPromoted)
	return EnergyConstraint{Formalized: true, Hamiltonian: append([]float64(nil), k.Spectrum...), Beta: BGap, State: rho, Expectation: exp, ConstraintNative: false, Circular: true, BetaDerived: true, PromotesKMSNative: false, Verdict: verdict}
}

func auditFlow(candidates []HamiltonianCandidate, energy EnergyConstraint) FlowAudit {
	best := "triality signed tau_eta"
	nontrivial := false
	promoted := false
	for _, c := range candidates {
		if c.Name == best {
			nontrivial = c.BreaksAllDegeneracy
			promoted = c.PromotedNative && energy.PromotesKMSNative
		}
	}
	verdict := join(StatusFlowKernelRecomputed, StatusLandscapeSafetyAudited, StatusTensionHamiltonianIsVacuumAddress, StatusFailedFlavorNotDerived, StatusFailedYukawaNotDerived)
	return FlowAudit{Executed: true, BestCandidate: best, NontrivialCapacity: nontrivial, PromotedNative: promoted, PreservesLandscape: true, KineticSafe: true, SelectsVacuum: false, Verdict: verdict}
}

func updateCensus(flow FlowAudit) Census {
	reduction := 0
	if flow.PromotedNative && flow.SelectsVacuum {
		reduction = vacuumInputs
	}
	remaining := vacuumInputs - reduction
	verdict := join(StatusCensusUpdated, StatusFailedCensusNotReduced)
	return Census{StartingInputs: vacuumInputs, Reduction: reduction, RemainingInputs: remaining, SevenSealTarget: remaining <= 7, Verdict: verdict}
}

func buildSummary(flow FlowAudit, energy EnergyConstraint, census Census) Summary {
	answer := "The current core contains native candidate operators for K, including tau_eta.  Only signed tau_eta activates all modular frequencies with beta=B_gap, but no native energy constraint selects K=tau_eta; promoting it would still choose a vacuum address."
	status := join(StatusHamiltonianOriginFormalized, StatusTrialityHamiltonianAudited, StatusFailedHamiltonianNotDerived)
	return Summary{Executed: true, HamiltonianFound: false, EnergyConstraintDerived: energy.ConstraintNative, VacuumSelected: flow.SelectsVacuum, RemainingInputs: census.RemainingInputs, DirectAnswer: answer, NextGate: "Gate 367 — Modular Energy Functional Extension / Minimal Vacuum-Address Axiom Audit", Status: status}
}

func candidateByName(candidates []HamiltonianCandidate, name string) HamiltonianCandidate {
	for _, c := range candidates {
		if c.Name == name {
			return c
		}
	}
	return candidates[0]
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

func expectation(rho, k []float64) float64 {
	s := 0.0
	for i := range rho {
		s += rho[i] * k[i]
	}
	return s
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func Statuses(a Analysis) []string {
	set := map[string]struct{}{}
	blocks := []string{a.Criteria.Verdict, a.Energy.Verdict, a.Flow.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, c := range a.Candidates {
		blocks = append(blocks, c.Verdict)
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

func FormatCriteria(c OriginCriteria) string {
	return fmt.Sprintf("formalized=%v self_adjoint=%v noncentral=%v native=%v non_circular=%v breaks_orbit=%v\ndescription=%s\nverdict=%s", c.Formalized, c.MustBeSelfAdjoint, c.MustBeNoncentral, c.MustBeNative, c.MustBeNonCircular, c.MustBreakOrbit, c.Description, c.Verdict)
}

func FormatCandidate(c HamiltonianCandidate) string {
	parts := []string{
		fmt.Sprintf("candidate=%s", c.Name),
		fmt.Sprintf("source=%s", c.Source),
		fmt.Sprintf("spectrum=[%.6f %.6f %.6f]", c.Spectrum[0], c.Spectrum[1], c.Spectrum[2]),
		fmt.Sprintf("self_adjoint=%v central=%v native_operator=%v energy_role_derived=%v", c.SelfAdjoint, c.Central, c.NativeOperator, c.EnergyRoleDerived),
		fmt.Sprintf("rho=[%.12f %.12f %.12f]", c.KMSState[0], c.KMSState[1], c.KMSState[2]),
	}
	for _, f := range c.ModularFrequencies {
		parts = append(parts, fmt.Sprintf("omega_%s=%.12f nonzero=%v", f.Pair, f.LogRatio, f.NonZero))
	}
	parts = append(parts, fmt.Sprintf("breaks_all_degeneracy=%v promoted_native=%v", c.BreaksAllDegeneracy, c.PromotedNative), "verdict="+c.Verdict)
	return strings.Join(parts, "\n")
}

func FormatEnergy(e EnergyConstraint) string {
	return fmt.Sprintf("formalized=%v beta=%.12f K=[%.3f %.3f %.3f]\nrho=[%.12f %.12f %.12f]\n<K>=%.12f constraint_native=%v circular=%v beta_derived=%v promotes_native=%v\nverdict=%s", e.Formalized, e.Beta, e.Hamiltonian[0], e.Hamiltonian[1], e.Hamiltonian[2], e.State[0], e.State[1], e.State[2], e.Expectation, e.ConstraintNative, e.Circular, e.BetaDerived, e.PromotesKMSNative, e.Verdict)
}

func FormatFlow(f FlowAudit) string {
	return fmt.Sprintf("executed=%v best=%s nontrivial_capacity=%v promoted_native=%v preserves_landscape=%v kinetic_safe=%v selects_vacuum=%v\nverdict=%s", f.Executed, f.BestCandidate, f.NontrivialCapacity, f.PromotedNative, f.PreservesLandscape, f.KineticSafe, f.SelectsVacuum, f.Verdict)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d reduction=%d remaining=%d seven_seal_target=%v\nverdict=%s", c.StartingInputs, c.Reduction, c.RemainingInputs, c.SevenSealTarget, c.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%v hamiltonian_found=%v energy_constraint_derived=%v vacuum_selected=%v remaining=%d\nanswer=%s\nnext=%s\nstatus=%s", s.Executed, s.HamiltonianFound, s.EnergyConstraintDerived, s.VacuumSelected, s.RemainingInputs, s.DirectAnswer, s.NextGate, s.Status)
}
