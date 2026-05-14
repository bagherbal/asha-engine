// Package modularspectralflowkernel implements Gate 363:
// Modular Spectral Flow Kernel / Vacuum Address Operator Construction Audit.
package modularspectralflowkernel

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE363-MODULAR-SPECTRAL-FLOW-KERNEL-VACUUM-ADDRESS-OPERATOR-CONSTRUCTION-AUDIT"

	StatusModularOperatorFormalized    = "CONDITIONAL_SUPPORT_MODULAR_OPERATOR_FORMALIZED"
	StatusTomitaLinkFormalized         = "CONDITIONAL_SUPPORT_TOMITA_TAKESAKI_J_LINK_FORMALIZED"
	StatusGradientFlowFormalized       = "CONDITIONAL_SUPPORT_GRADIENT_FLOW_EQUATION_FORMALIZED"
	StatusFlavorOrbitSieveExecuted     = "CONDITIONAL_SUPPORT_FLAVOR_ORBIT_SIEVE_EXECUTED"
	StatusLandscapePreservationAudited = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_SIEVE_AUDITED"
	StatusKineticSafetyAudited         = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusNonTracialCapacityIdentified = "CONDITIONAL_SUPPORT_NONTRACIAL_MODULAR_STATE_CAPACITY_IDENTIFIED"

	StatusTensionTracialFlowTrivial        = "CONDITIONAL_TENSION_NATIVE_TRACIAL_STATE_GENERATES_TRIVIAL_MODULAR_FLOW"
	StatusTensionNonTracialStateNotDerived = "CONDITIONAL_TENSION_NONTRACIAL_STATE_REQUIRED_BUT_NOT_DERIVED"
	StatusTensionRhoWouldBeVacuumAddress   = "CONDITIONAL_TENSION_MODULAR_DENSITY_MATRIX_IS_ITSELF_A_VACUUM_ADDRESS"

	StatusFailedKernelNotConstructed      = "FAILED_ROUTE_MODULAR_FLOW_KERNEL_NOT_CONSTRUCTED_NATIVELY"
	StatusFailedDegeneracyNotBroken       = "FAILED_ROUTE_DYNAMICAL_FLAVOR_BREAKING_NOT_VERIFIED"
	StatusFailedVacuumNotSelected         = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_BY_MODULAR_FLOW"
	StatusFailedNonTracialStateNotDerived = "FAILED_ROUTE_NONTRACIAL_FAITHFUL_STATE_NOT_DERIVED"
	StatusFailedCKMNotDerived             = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_MODULAR_FLOW"
	StatusFailedYukawaNotDerived          = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_MODULAR_FLOW"
)

const (
	vacuumInputs = 15
	eps          = 1e-12
)

type ModularState struct {
	Name               string
	Probabilities      []float64
	Faithful           bool
	Tracial            bool
	NativeDerived      bool
	ModularHamiltonian []float64
	Verdict            string
}

type ModularOperator struct {
	Formalized      bool
	State           ModularState
	DeltaDefined    bool
	SigmaDefinition string
	JLink           string
	KGenerator      []float64
	NonTrivial      bool
	Verdict         string
}

type FlowEquation struct {
	Formalized      bool
	Equation        string
	EffectiveAction string
	Metric          string
	RequiresRho     bool
	Verdict         string
}

type MatrixElementAction struct {
	Element          string
	LambdaRatio      float64
	Frequency        float64
	MagnitudeChanged bool
	Detail           string
}

type FlavorOrbitSieve struct {
	Executed                    bool
	NativeTracialActions        []MatrixElementAction
	NonTracialCandidateActions  []MatrixElementAction
	NativeBreaksFlavorOrbit     bool
	CandidateBreaksFlavorOrbit  bool
	CandidateSelectsUniquePoint bool
	DegeneracyBroken            bool
	RemainingInputs             int
	Verdict                     string
}

type LandscapePreservation struct {
	Audited                bool
	PreservesWeakAngle     bool
	PreservesQuarticRatio  bool
	PreservesAlphaGUT      bool
	PreservesMoritaSplit   bool
	PreservesKineticSafety bool
	Reason                 string
	Verdict                string
}

type Summary struct {
	Executed          bool
	KernelFormalized  bool
	NativeKernelFound bool
	VacuumSelected    bool
	RemainingInputs   int
	DirectAnswer      string
	NextGate          string
	Status            string
}

type Analysis struct {
	Operator  ModularOperator
	Flow      FlowEquation
	Flavor    FlavorOrbitSieve
	Landscape LandscapePreservation
	Summary   Summary
	Truth     string
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
	state := nativeTracialState(3)
	op := formalizeOperator(state)
	flow := formalizeGradient()
	flavor := executeFlavorOrbitSieve(op)
	landscape := auditLandscape(op, flavor)
	summary := buildSummary(op, flow, flavor, landscape)
	truth := "Gate 363 constructs the Tomita-Takesaki modular-flow formalism for the finite flavor algebra.  For the native tracial state inherited from closed ASHA spectral traces, the modular operator is trivial and the flow cannot break CKM/PMNS degeneracy.  A nontracial faithful density matrix would create a nontrivial modular Hamiltonian, but that density matrix is itself a vacuum-address input and is not derived by the current core.  Therefore the gate formalizes the correct Path-B kernel target while preserving the firewall: modular flow needs a native nontracial state theorem before it can select the vacuum."
	return Analysis{Operator: op, Flow: flow, Flavor: flavor, Landscape: landscape, Summary: summary, Truth: truth}, nil
}

func nativeTracialState(n int) ModularState {
	p := make([]float64, n)
	for i := range p {
		p[i] = 1 / float64(n)
	}
	return buildState("native finite spectral trace state", p, true)
}

func candidateNonTracialState() ModularState {
	// Deliberately only a capacity witness: positive and faithful, but not derived.
	p := []float64{0.50, 0.30, 0.20}
	return buildState("nontracial faithful candidate density", p, false)
}

func buildState(name string, p []float64, native bool) ModularState {
	faithful := true
	for _, x := range p {
		if x <= 0 {
			faithful = false
		}
	}
	tracial := true
	for i := 1; i < len(p); i++ {
		if math.Abs(p[i]-p[0]) > eps {
			tracial = false
			break
		}
	}
	k := make([]float64, len(p))
	for i, x := range p {
		if x > 0 {
			k[i] = -math.Log(x)
		} else {
			k[i] = math.Inf(1)
		}
	}
	verdict := StatusModularOperatorFormalized
	if tracial {
		verdict = join(StatusModularOperatorFormalized, StatusTensionTracialFlowTrivial)
	}
	if !native && !tracial {
		verdict = join(StatusNonTracialCapacityIdentified, StatusTensionNonTracialStateNotDerived, StatusFailedNonTracialStateNotDerived)
	}
	return ModularState{Name: name, Probabilities: p, Faithful: faithful, Tracial: tracial, NativeDerived: native, ModularHamiltonian: k, Verdict: verdict}
}

func formalizeOperator(state ModularState) ModularOperator {
	nontrivial := false
	if len(state.ModularHamiltonian) > 1 {
		for i := 1; i < len(state.ModularHamiltonian); i++ {
			if math.Abs(state.ModularHamiltonian[i]-state.ModularHamiltonian[0]) > eps {
				nontrivial = true
				break
			}
		}
	}
	verdict := join(StatusModularOperatorFormalized, StatusTomitaLinkFormalized)
	if !nontrivial {
		verdict = join(verdict, StatusTensionTracialFlowTrivial, StatusFailedDegeneracyNotBroken)
	}
	return ModularOperator{
		Formalized:      true,
		State:           state,
		DeltaDefined:    state.Faithful,
		SigmaDefinition: "sigma_t(A) = Delta^{it} A Delta^{-it}; in finite density-matrix form sigma_t(E_ij)=(rho_i/rho_j)^{it} E_ij",
		JLink:           "Tomita modular conjugation is linked to the ASHA real structure J by requiring J_flow to preserve the J_swap particle-antiparticle pairing while acting on the GNS/state representation; the current native trace state makes this link formal but dynamically trivial.",
		KGenerator:      state.ModularHamiltonian,
		NonTrivial:      nontrivial,
		Verdict:         verdict,
	}
}

func formalizeGradient() FlowEquation {
	return FlowEquation{
		Formalized:      true,
		Equation:        "dX/ds = - grad_Theta S_eff(X)",
		EffectiveAction: "S_eff plus modular penalty/relative-entropy direction generated by K_Theta=-log(rho_flow)",
		Metric:          "positive kinetic metric required by Gate 301/302; no ghost or rank-collapse directions admitted",
		RequiresRho:     true,
		Verdict:         StatusGradientFlowFormalized,
	}
}

func executeFlavorOrbitSieve(op ModularOperator) FlavorOrbitSieve {
	nativeActions := actionsForState(op.State)
	cand := candidateNonTracialState()
	candidateActions := actionsForState(cand)
	candidateBreaks := !cand.Tracial
	nativeBreaks := op.NonTrivial
	verdict := join(StatusFlavorOrbitSieveExecuted, StatusTensionTracialFlowTrivial, StatusFailedDegeneracyNotBroken, StatusFailedVacuumNotSelected)
	if candidateBreaks {
		verdict = join(verdict, StatusNonTracialCapacityIdentified, StatusTensionNonTracialStateNotDerived, StatusTensionRhoWouldBeVacuumAddress)
	}
	return FlavorOrbitSieve{
		Executed:                    true,
		NativeTracialActions:        nativeActions,
		NonTracialCandidateActions:  candidateActions,
		NativeBreaksFlavorOrbit:     nativeBreaks,
		CandidateBreaksFlavorOrbit:  candidateBreaks,
		CandidateSelectsUniquePoint: false,
		DegeneracyBroken:            false,
		RemainingInputs:             vacuumInputs,
		Verdict:                     verdict,
	}
}

func actionsForState(state ModularState) []MatrixElementAction {
	labels := []string{"E12", "E13", "E23"}
	pairs := [][2]int{{0, 1}, {0, 2}, {1, 2}}
	out := make([]MatrixElementAction, 0, len(pairs))
	for idx, pair := range pairs {
		i, j := pair[0], pair[1]
		ratio := state.Probabilities[i] / state.Probabilities[j]
		freq := math.Log(ratio)
		out = append(out, MatrixElementAction{
			Element:          labels[idx],
			LambdaRatio:      ratio,
			Frequency:        freq,
			MagnitudeChanged: false,
			Detail:           fmt.Sprintf("sigma_t(%s) = exp(i t %.12g) %s", labels[idx], freq, labels[idx]),
		})
	}
	return out
}

func auditLandscape(op ModularOperator, flavor FlavorOrbitSieve) LandscapePreservation {
	kineticSafe := op.State.Faithful
	verdict := join(StatusLandscapePreservationAudited, StatusKineticSafetyAudited)
	if !op.NonTrivial {
		verdict = join(verdict, StatusTensionTracialFlowTrivial)
	}
	return LandscapePreservation{
		Audited:                true,
		PreservesWeakAngle:     true,
		PreservesQuarticRatio:  true,
		PreservesAlphaGUT:      true,
		PreservesMoritaSplit:   true,
		PreservesKineticSafety: kineticSafe,
		Reason:                 "A modular flow generated inside the generation/vacuum orbit can preserve central gauge/Morita invariants if rho_flow commutes with the landscape projectors; the native tracial state certainly preserves them but is dynamically trivial.",
		Verdict:                verdict,
	}
}

func buildSummary(op ModularOperator, flow FlowEquation, flavor FlavorOrbitSieve, landscape LandscapePreservation) Summary {
	statuses := []string{StatusModularOperatorFormalized, StatusGradientFlowFormalized, StatusFlavorOrbitSieveExecuted, StatusLandscapePreservationAudited}
	if !op.NonTrivial {
		statuses = append(statuses, StatusFailedKernelNotConstructed, StatusFailedDegeneracyNotBroken, StatusFailedVacuumNotSelected)
	}
	direct := "Gate 363 constructs the modular spectral-flow formalism but does not derive a native vacuum selector.  With the native tracial finite spectral state, Delta is the identity and sigma_t is trivial on flavor.  A nontracial faithful state would break the flavor orbit, but deriving that state is exactly the missing vacuum-address theorem."
	return Summary{Executed: true, KernelFormalized: op.Formalized && flow.Formalized, NativeKernelFound: op.NonTrivial && op.State.NativeDerived, VacuumSelected: flavor.DegeneracyBroken, RemainingInputs: flavor.RemainingInputs, DirectAnswer: direct, NextGate: "Gate 364 — Nontracial Modular State Origin / Vacuum Density Matrix Derivation Audit", Status: join(statuses...)}
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
	add(a.Operator.State.Verdict)
	add(a.Operator.Verdict)
	add(a.Flow.Verdict)
	add(a.Flavor.Verdict)
	add(a.Landscape.Verdict)
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

func FormatState(s ModularState) string {
	return fmt.Sprintf("%s faithful=%t tracial=%t native=%t rho=%v K=%v verdict=%s", s.Name, s.Faithful, s.Tracial, s.NativeDerived, s.Probabilities, s.ModularHamiltonian, s.Verdict)
}

func FormatOperator(o ModularOperator) string {
	return fmt.Sprintf("formalized=%t delta=%t nontrivial=%t state={%s} sigma=%q J=%q verdict=%s", o.Formalized, o.DeltaDefined, o.NonTrivial, FormatState(o.State), o.SigmaDefinition, o.JLink, o.Verdict)
}

func FormatFlow(f FlowEquation) string {
	return fmt.Sprintf("formalized=%t equation=%q action=%q metric=%q requiresRho=%t verdict=%s", f.Formalized, f.Equation, f.EffectiveAction, f.Metric, f.RequiresRho, f.Verdict)
}

func FormatActions(actions []MatrixElementAction) string {
	parts := make([]string, 0, len(actions))
	for _, a := range actions {
		parts = append(parts, fmt.Sprintf("%s ratio=%.12g freq=%.12g magnitudeChanged=%t", a.Element, a.LambdaRatio, a.Frequency, a.MagnitudeChanged))
	}
	return strings.Join(parts, "; ")
}

func FormatFlavor(f FlavorOrbitSieve) string {
	return fmt.Sprintf("executed=%t nativeBreaks=%t candidateBreaks=%t selected=%t remaining=%d native=[%s] candidate=[%s] verdict=%s", f.Executed, f.NativeBreaksFlavorOrbit, f.CandidateBreaksFlavorOrbit, f.DegeneracyBroken, f.RemainingInputs, FormatActions(f.NativeTracialActions), FormatActions(f.NonTracialCandidateActions), f.Verdict)
}

func FormatLandscape(l LandscapePreservation) string {
	return fmt.Sprintf("audited=%t weak=%t quartic=%t alpha=%t morita=%t kinetic=%t reason=%q verdict=%s", l.Audited, l.PreservesWeakAngle, l.PreservesQuarticRatio, l.PreservesAlphaGUT, l.PreservesMoritaSplit, l.PreservesKineticSafety, l.Reason, l.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t kernelFormalized=%t nativeKernelFound=%t vacuumSelected=%t remaining=%d next=%q answer=%q status=%s", s.Executed, s.KernelFormalized, s.NativeKernelFound, s.VacuumSelected, s.RemainingInputs, s.NextGate, s.DirectAnswer, s.Status)
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
