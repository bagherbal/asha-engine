// Package modulartimeflowvacuumselector implements Gate 362:
// Modular Time Flow / Dynamical Vacuum Selector Extension Audit.
package modulartimeflowvacuumselector

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE362-MODULAR-TIME-FLOW-DYNAMICAL-VACUUM-SELECTOR-EXTENSION-AUDIT"

	StatusPathBActivated                  = "CONDITIONAL_SUPPORT_PATH_B_DYNAMIC_EXTENSION_ACTIVATED"
	StatusReadmeShiftFormalized           = "CONDITIONAL_SUPPORT_PROJECT_README_FLOW_SHIFT_FORMALIZED"
	StatusFlowOperatorClassIntroduced     = "CONDITIONAL_SUPPORT_NEW_FLOW_OPERATOR_CLASS_INTRODUCED"
	StatusModularFlowAxiomsFormalized     = "CONDITIONAL_SUPPORT_MODULAR_FLOW_AXIOMS_FORMALIZED"
	StatusAdmissibilitySieveExecuted      = "CONDITIONAL_SUPPORT_FLOW_ADMISSIBILITY_SIEVE_EXECUTED"
	StatusLandscapePreservationFormalized = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_CONSTRAINT_FORMALIZED"
	StatusVacuumSelectorTargetFormalized  = "CONDITIONAL_SUPPORT_VACUUM_SELECTOR_TARGET_FORMALIZED"
	StatusPhaseIIIFocusDeclared           = "CONDITIONAL_SUPPORT_PHASE_III_FLOW_BASED_FOCUS_DECLARED"
	StatusMinimalExtensionForkInstalled   = "CONDITIONAL_SUPPORT_MINIMAL_DYNAMICAL_EXTENSION_FORK_INSTALLED"

	StatusTensionFlowOperatorNotConstructed = "CONDITIONAL_TENSION_EXPLICIT_MODULAR_FLOW_OPERATOR_NOT_CONSTRUCTED"
	StatusTensionNeedNonInvariantGenerator  = "CONDITIONAL_TENSION_FLOW_MUST_BREAK_UNITARY_FLAVOR_INVARIANCE"
	StatusTensionMustPreserveLandscape      = "CONDITIONAL_TENSION_FLOW_EXTENSION_MUST_NOT_DAMAGE_DERIVED_LANDSCAPE"
	StatusTensionVacuumSelectionStillTarget = "CONDITIONAL_TENSION_VACUUM_SELECTION_NOW_DYNAMIC_TARGET_NOT_STATIC_CORE_TASK"

	StatusFailedVacuumPointNotSelected    = "FAILED_ROUTE_VACUUM_POINT_NOT_SELECTED_YET"
	StatusFailedExplicitFlowKernelMissing = "FAILED_ROUTE_EXPLICIT_MODULAR_FLOW_KERNEL_NOT_DERIVED"
	StatusFailedCKMNotDerived             = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED_BY_FLOW_YET"
	StatusFailedYukawaNotDerived          = "FAILED_ROUTE_YUKAWA_COORDINATES_NOT_DERIVED_BY_FLOW_YET"
	StatusFailedCosmologicalNotDerived    = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED_BY_FLOW_YET"
	StatusFailedNoFinalToE                = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
)

const (
	inheritedGate        = 361
	startingVacuumInputs = 15
	sevenSealTarget      = 7
)

type Span struct {
	AuditID       string
	InheritedGate int
	Path          string
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type DocumentationShift struct {
	Required           bool
	AppliedInReadme    bool
	AppliedInDocs      bool
	OldMisleadingFrame string
	NewFrame           string
	Directive          string
	Verdict            string
}

type FlowAxiom struct {
	Name        string
	Requirement string
	Reason      string
	Mandatory   bool
}

type FlowOperatorClass struct {
	Name                        string
	NewOperatorClass            bool
	NativeCandidate             bool
	BreaksUnitaryFlavorSymmetry bool
	PreservesLandscapeRatios    bool
	KineticSafeTarget           bool
	SelectsVacuumPoint          bool
	MissingTheorem              string
	Verdict                     string
}

type AdmissibilitySieve struct {
	Executed                  bool
	Axioms                    []FlowAxiom
	Candidate                 FlowOperatorClass
	PreservesLandscape        bool
	BreaksFlavorInvariant     bool
	ExplicitKernelConstructed bool
	VacuumSelected            bool
	RemainingInputs           int
	Verdict                   string
}

type Program struct {
	Formalized        bool
	Name              string
	ImmediateGate     string
	RequiredArtifacts []string
	ForbiddenMoves    []string
	SuccessCriterion  string
	FailureCriterion  string
	Verdict           string
}

type Census struct {
	StartingVacuumInputs int
	ReductionFromFlow    int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed          bool
	PathBActive       bool
	FlowKernelMissing bool
	RemainingInputs   int
	Status            string
	DirectAnswer      string
	NextGate          string
}

type Analysis struct {
	Span    Span
	Docs    DocumentationShift
	Sieve   AdmissibilitySieve
	Program Program
	Census  Census
	Summary Summary
	Truth   string
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
	span := compileSpan()
	docs := formalizeDocumentationShift()
	sieve := executeAdmissibilitySieve()
	program := formalizeProgram()
	census := updateCensus(sieve)
	summary := buildSummary(span, docs, sieve, program, census)
	truth := "Gate 362 activates Path B: the ASHA core is no longer asked to squeeze vacuum coordinates from the closed static algebra.  From this gate forward the project explicitly focuses on flow: a new modular/Lorentzian time operator class must be constructed that breaks unitary flavor invariance, preserves all rigid landscape theorems, remains kinetic-safe, and selects a vacuum point dynamically.  This gate installs the extension charter; it does not yet derive the flow kernel or reduce the 15 vacuum coordinates."
	return Analysis{Span: span, Docs: docs, Sieve: sieve, Program: program, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, Path: "B: minimal dynamical extension", AddsFit: false, Purpose: "introduce a new native time/modular-flow operator class after Gate 361 proved the static ASHA core complete as a landscape theory", Verdict: StatusPathBActivated}
}

func formalizeDocumentationShift() DocumentationShift {
	return DocumentationShift{
		Required:           true,
		AppliedInReadme:    true,
		AppliedInDocs:      true,
		OldMisleadingFrame: "continued static texture searches inside the already closed finite core",
		NewFrame:           "from Gate 362 onward, Phase III focuses on flow-based vacuum selection: modular/Lorentzian time dynamics, not more static resonances",
		Directive:          "Any future vacuum-selection gate must introduce a new flow operator, a theorem about that operator, or a no-go result about that operator; it must not merely try another static texture.",
		Verdict:            join(StatusReadmeShiftFormalized, StatusPhaseIIIFocusDeclared),
	}
}

func flowAxioms() []FlowAxiom {
	return []FlowAxiom{
		{Name: "new operator class", Requirement: "the selector must not be expressible as a finite combination of the Gate-361 closed static operators", Reason: "otherwise Gate 361 already applies", Mandatory: true},
		{Name: "non-unitary-invariant flavor action", Requirement: "the flow generator must act nontrivially on the flavor orbit", Reason: "unitary trace invariants leave CKM/PMNS directions flat", Mandatory: true},
		{Name: "landscape preservation", Requirement: "the flow must preserve sin²θ_W=3/8, λ/g²=1197/4624, α_GUT⁻¹=8π, v/M_P=2^(3/2)e^(-4π²), and Morita 1:3 structure", Reason: "a vacuum selector may choose coordinates, but it must not rewrite the laws", Mandatory: true},
		{Name: "kinetic safety", Requirement: "the flow must not create ghosts, rank collapse, or non-positive wave-function normalization", Reason: "Gate 301-302 positivity firewalls remain mandatory", Mandatory: true},
		{Name: "vacuum address output", Requirement: "the flow must select a discrete or continuous point in the 15-dimensional vacuum-coordinate space", Reason: "capacity or near-match is insufficient after Gate 361", Mandatory: true},
	}
}

func executeAdmissibilitySieve() AdmissibilitySieve {
	axioms := flowAxioms()
	candidate := FlowOperatorClass{
		Name:                        "modular/Lorentzian time-flow vacuum-address operator Θ_flow",
		NewOperatorClass:            true,
		NativeCandidate:             true,
		BreaksUnitaryFlavorSymmetry: false,
		PreservesLandscapeRatios:    true,
		KineticSafeTarget:           true,
		SelectsVacuumPoint:          false,
		MissingTheorem:              "explicit modular-flow kernel Θ_flow and its non-unitary flavor-gradient action are not yet constructed",
		Verdict:                     join(StatusFlowOperatorClassIntroduced, StatusTensionFlowOperatorNotConstructed, StatusFailedExplicitFlowKernelMissing),
	}
	verdict := join(StatusModularFlowAxiomsFormalized, StatusAdmissibilitySieveExecuted, StatusLandscapePreservationFormalized, StatusVacuumSelectorTargetFormalized, StatusTensionNeedNonInvariantGenerator, StatusTensionMustPreserveLandscape, StatusFailedVacuumPointNotSelected)
	return AdmissibilitySieve{Executed: true, Axioms: axioms, Candidate: candidate, PreservesLandscape: candidate.PreservesLandscapeRatios, BreaksFlavorInvariant: candidate.BreaksUnitaryFlavorSymmetry, ExplicitKernelConstructed: false, VacuumSelected: false, RemainingInputs: startingVacuumInputs, Verdict: verdict}
}

func formalizeProgram() Program {
	return Program{
		Formalized:    true,
		Name:          "Phase III Flow Program",
		ImmediateGate: "Gate 363 — Modular Spectral Flow Kernel / Vacuum Address Operator Construction Audit",
		RequiredArtifacts: []string{
			"a Lorentzian or Tomita-Takesaki modular generator Θ_flow",
			"a proof that Θ_flow is not a unitary trace invariant on flavor space",
			"a positive/kinetic-safe induced metric on moduli space",
			"a gradient or semigroup equation dX/ds = -∇_Θ S_eff(X)",
			"a preservation proof for all Gate-348/Gate-361 landscape ratios",
		},
		ForbiddenMoves: []string{
			"do not fit CKM/PMNS angles by assigning a texture by hand",
			"do not modify the derived landscape ratios to improve mass fits",
			"do not count parameter reductions from attractor capacity unless the UV coordinate becomes unnecessary",
			"do not use a non-unitary projector unless its kinetic-safe metric is derived",
		},
		SuccessCriterion: "derive an explicit Θ_flow whose fixed points or attractors select at least one previously quarantined vacuum coordinate without empirical input",
		FailureCriterion: "show that even the minimal modular/Lorentzian extension is flavor-invariant or dynamically degenerate",
		Verdict:          StatusMinimalExtensionForkInstalled,
	}
}

func updateCensus(s AdmissibilitySieve) Census {
	reduction := 0
	if s.VacuumSelected {
		reduction = 1
	}
	remaining := startingVacuumInputs - reduction
	verdict := join(StatusTensionVacuumSelectionStillTarget, StatusFailedCKMNotDerived, StatusFailedYukawaNotDerived, StatusFailedCosmologicalNotDerived)
	return Census{StartingVacuumInputs: startingVacuumInputs, ReductionFromFlow: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: verdict}
}

func buildSummary(span Span, docs DocumentationShift, sieve AdmissibilitySieve, program Program, census Census) Summary {
	statuses := []string{span.Verdict, docs.Verdict, sieve.Candidate.Verdict, sieve.Verdict, program.Verdict, census.Verdict, StatusFailedNoFinalToE}
	direct := "Path B is now active.  If the README or earlier gate language implied that the existing static ASHA core should keep deriving the vacuum by more texture searches, that is now superseded: from Gate 362 forward the project focuses on flow-based vacuum selection.  The current gate installs the modular/Lorentzian time-flow operator class and its admissibility rules, but the explicit Θ_flow kernel has not yet been derived."
	next := program.ImmediateGate
	return Summary{Executed: true, PathBActive: true, FlowKernelMissing: !sieve.ExplicitKernelConstructed, RemainingInputs: census.RemainingInputs, Status: join(statuses...), DirectAnswer: direct, NextGate: next}
}

func Statuses(a Analysis) []string {
	seen := map[string]bool{}
	add := func(v string) {
		for _, p := range strings.Split(v, ";") {
			p = strings.TrimSpace(p)
			if p != "" {
				seen[p] = true
			}
		}
	}
	add(a.Span.Verdict)
	add(a.Docs.Verdict)
	add(a.Sieve.Candidate.Verdict)
	add(a.Sieve.Verdict)
	add(a.Program.Verdict)
	add(a.Census.Verdict)
	add(a.Summary.Status)
	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s inherited=%d path=%s addsFit=%v purpose=%s verdict=%s", s.AuditID, s.InheritedGate, s.Path, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatDocs(d DocumentationShift) string {
	return fmt.Sprintf("required=%v readme=%v docs=%v old=%q new=%q directive=%q verdict=%s", d.Required, d.AppliedInReadme, d.AppliedInDocs, d.OldMisleadingFrame, d.NewFrame, d.Directive, d.Verdict)
}
func FormatCandidate(c FlowOperatorClass) string {
	return fmt.Sprintf("%s new=%v nativeCandidate=%v breaksFlavor=%v preservesLandscape=%v kineticSafeTarget=%v selectsVacuum=%v missing=%s verdict=%s", c.Name, c.NewOperatorClass, c.NativeCandidate, c.BreaksUnitaryFlavorSymmetry, c.PreservesLandscapeRatios, c.KineticSafeTarget, c.SelectsVacuumPoint, c.MissingTheorem, c.Verdict)
}
func FormatSieve(s AdmissibilitySieve) string {
	return fmt.Sprintf("executed=%v axioms=%d preservesLandscape=%v breaksFlavor=%v explicitKernel=%v vacuumSelected=%v remaining=%d candidate={%s} verdict=%s", s.Executed, len(s.Axioms), s.PreservesLandscape, s.BreaksFlavorInvariant, s.ExplicitKernelConstructed, s.VacuumSelected, s.RemainingInputs, FormatCandidate(s.Candidate), s.Verdict)
}
func FormatProgram(p Program) string {
	return fmt.Sprintf("formalized=%v name=%s next=%s artifacts=%d forbidden=%d success=%q failure=%q verdict=%s", p.Formalized, p.Name, p.ImmediateGate, len(p.RequiredArtifacts), len(p.ForbiddenMoves), p.SuccessCriterion, p.FailureCriterion, p.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d reduction=%d remaining=%d target=%d reached=%v verdict=%s", c.StartingVacuumInputs, c.ReductionFromFlow, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%v pathB=%v flowKernelMissing=%v remaining=%d next=%s status=%s", s.Executed, s.PathBActive, s.FlowKernelMissing, s.RemainingInputs, s.NextGate, s.Status)
}

func join(parts ...string) string {
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if strings.TrimSpace(p) != "" {
			out = append(out, p)
		}
	}
	return strings.Join(out, ";")
}
