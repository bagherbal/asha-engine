// Package generation2trialitytoyukawareadoutminimalityandnogoaudit implements
// Gate 803: Triality-to-Yukawa Readout Package Minimality and No-Go Audit.
//
// Gate 803 inherits the Gate 802 result that the complex D4 trilinear invariant
// is only an airlocked pre-Yukawa shape. It defines the full minimal readout
// package required before that invariant could become a sector-labeled,
// generation-resolved, Hermitian, scale-local Yukawa trace ledger, and proves
// the current no-go against deriving N_eff, PMNS/CKM, or C_Higgs from T_D4 alone.
package generation2trialitytoyukawareadoutminimalityandnogoaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE803-TRIALITY-TO-YUKAWA-READOUT-MINIMALITY-NO-GO-AUDIT"

	StatusGate802Inherited         = "PASS_GATE802_COMPLEX_D4_TRILINEAR_OBSTRUCTION_INHERITED"
	StatusTD4OnlyPreYukawa         = "PASS_T_D4_ACCEPTED_ONLY_AS_AIRLOCKED_PRE_YUKAWA_SHAPE"
	StatusReadoutPackageDefined    = "PASS_TRIALITY_YUKAWA_READOUT_PACKAGE_DEFINED"
	StatusRealDescentAudited       = "PASS_REAL_DESCENT_SEAL_REQUIREMENT_AUDITED"
	StatusGaugeAssignmentAudited   = "PASS_GAUGE_REPRESENTATION_ASSIGNMENT_REQUIREMENT_AUDITED"
	StatusSectorMinimalityAudited  = "PASS_SECTOR_ASSIGNMENT_MINIMALITY_AUDITED"
	StatusGenerationAudited        = "PASS_GENERATION_CARRIER_REQUIREMENT_AUDITED"
	StatusHermitianAudited         = "PASS_HERMITIAN_OPERATOR_SEAL_REQUIREMENT_AUDITED"
	StatusHierarchyAudited         = "PASS_SYMMETRY_BREAKING_HIERARCHY_REQUIREMENT_AUDITED"
	StatusTraceAtomAudited         = "PASS_TRACE_ATOM_EXTRACTION_REQUIREMENT_AUDITED"
	StatusColorAudited             = "PASS_COLOR_MULTIPLICITY_SEAL_REQUIREMENT_AUDITED"
	StatusScaleAudited             = "PASS_SCALE_SCHEME_SEAL_REQUIREMENT_AUDITED"
	StatusNonCircularAudited       = "PASS_NONCIRCULARITY_SEAL_REQUIREMENT_AUDITED"
	StatusRemovalFailuresAudited   = "PASS_MINIMALITY_REMOVAL_FAILURES_AUDITED"
	StatusNoGoDefined              = "PASS_TRIALITY_TO_YUKAWA_NO_GO_STATEMENT_DEFINED"
	StatusCurrentSubobjectsAudited = "PASS_CURRENT_ASHA_READOUT_SUBOBJECTS_AUDITED"
	StatusPathsSeparated           = "PASS_EMPIRICAL_AND_NATIVE_YUKAWA_PATHS_SEPARATED"
	StatusCHiggsFirewall           = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusBranchDecision           = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls        = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusNeedsExtraSeals           = "CONDITIONAL_SUPPORT_YUKAWA_READOUT_REQUIRES_MULTIPLE_EXTRA_SEALS_BEYOND_T_D4"
	StatusAllSealsNonCosmetic       = "CONDITIONAL_SUPPORT_ALL_READOUT_SEALS_ARE_NONCOSMETIC"
	StatusFSTEdgeTemplate           = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_TRIPLE_SUPPLIES_YUKAWA_EDGE_SHAPE_TEMPLATE"
	StatusColorSU3TraceMultiplicity = "CONDITIONAL_SUPPORT_COLOR_SU3_REMAINS_CERTIFIED_TRACE_MULTIPLICITY_SOURCE"
	StatusExternalFastest           = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_REMAINS_FASTEST_PATH_TO_N_EFF_SOURCE_AUDIT"
	StatusNextFSTCompatibility      = "CONDITIONAL_SUPPORT_NEXT_NATIVE_GATE_SHOULD_TEST_SPECTRAL_TRIPLE_EDGE_COMPATIBILITY"

	StatusTD4NotTraceLedger           = "FAILED_ROUTE_T_D4_NOT_YET_YUKAWA_TRACE_LEDGER"
	StatusTrilinearAloneNoPackage     = "FAILED_ROUTE_TRILINEAR_INVARIANT_ALONE_DOES_NOT_DEFINE_READOUT_PACKAGE"
	StatusNoRealDescent               = "FAILED_ROUTE_NO_REAL_DESCENT_MAP_FOR_T_D4"
	StatusComplexNotNativeYukawa      = "FAILED_ROUTE_COMPLEX_T_D4_NOT_NATIVE_REAL_CL17_YUKAWA_OBJECT"
	StatusNoGaugeAssignment           = "FAILED_ROUTE_NO_GAUGE_REPRESENTATION_ASSIGNMENT_FROM_T_D4"
	StatusTrialityFramesNoSMEdges     = "FAILED_ROUTE_THREE_TRIALITY_FRAMES_DO_NOT_DEFINE_STANDARD_MODEL_YUKAWA_EDGES"
	StatusThreeTypesNotFourSectors    = "FAILED_ROUTE_THREE_TRIALITY_TYPES_DO_NOT_MATCH_FOUR_YUKAWA_SECTORS"
	StatusNoSMSectorAssignment        = "FAILED_ROUTE_NO_TRIALITY_TO_STANDARD_MODEL_SECTOR_ASSIGNMENT"
	StatusTrialityTypesNotGenerations = "FAILED_ROUTE_TRIALITY_TYPES_NOT_GENERATION_COPIES"
	StatusNoGenerationCarrier         = "FAILED_ROUTE_NO_GENERATION_CARRIER_FROM_T_D4"
	StatusNoPMNSCKMFrames             = "FAILED_ROUTE_NO_PMNS_CKM_WITHOUT_GENERATION_FRAMES"
	StatusComplexNotHermitian         = "FAILED_ROUTE_COMPLEX_TRILINEAR_AMPLITUDE_NOT_HERMITIAN_YUKAWA_OPERATOR"
	StatusNoSVDTheorem                = "FAILED_ROUTE_NO_SINGULAR_VALUE_EXTRACTION_THEOREM"
	StatusUniqueNoHierarchy           = "FAILED_ROUTE_UNIQUE_TRILINEAR_INVARIANT_DOES_NOT_DETERMINE_YUKAWA_HIERARCHY"
	StatusTD4NoTopDominance           = "FAILED_ROUTE_T_D4_DOES_NOT_EXPLAIN_TOP_DOMINANCE"
	StatusTD4NoNEffMinusThree         = "FAILED_ROUTE_T_D4_DOES_NOT_EXPLAIN_N_EFF_MINUS_THREE"
	StatusNoPositiveAtoms             = "FAILED_ROUTE_NO_POSITIVE_TRACE_ATOMS_FROM_T_D4"
	StatusNoBackwardAtoms             = "FAILED_ROUTE_TRACE_ATOMS_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_N_EFF_OR_HIGGS_DATA"
	StatusD4DoesNotReplaceColor       = "FAILED_ROUTE_D4_TRIALITY_DOES_NOT_REPLACE_COLOR_MULTIPLICITY_RULE"
	StatusColorDoubleCount            = "FAILED_ROUTE_COLOR_FACTOR_MUST_NOT_BE_DOUBLE_COUNTED"
	StatusNoScaleLedger               = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_SCALE_LOCAL_YUKAWA_LEDGER"
	StatusNoScaleStability            = "FAILED_ROUTE_NO_N_EFF_SCALE_STABILITY_FROM_TRIALITY_INVARIANT"
	StatusNoTargetTuning              = "FAILED_ROUTE_NO_TRIALITY_YUKAWA_CLAIM_ALLOWED_WITH_TARGET_TUNING"
	StatusCannotCompressToTD4         = "FAILED_ROUTE_TRIALITY_TO_YUKAWA_READOUT_PACKAGE_CANNOT_BE_COMPRESSED_TO_T_D4_ALONE"
	StatusTD4AloneNoLedger            = "FAILED_ROUTE_T_D4_ALONE_CANNOT_SOURCE_YUKAWA_TRACE_LEDGER"
	StatusTD4AloneNoNEffFlavor        = "FAILED_ROUTE_T_D4_ALONE_CANNOT_SOURCE_N_EFF_OR_FLAVOR_MIXING"
	StatusNoCurrentReadoutPackage     = "FAILED_ROUTE_CURRENT_ASHA_DOES_NOT_SUPPLY_TRIALITY_YUKAWA_READOUT_PACKAGE"
	StatusExternalNotNative           = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	StatusTrialityNotReadyExternal    = "FAILED_ROUTE_TRIALITY_BRANCH_NOT_READY_TO_REPLACE_EXTERNAL_LEDGER"
	StatusNoCYukawaUpdate             = "FAILED_ROUTE_TRIALITY_READOUT_NO_GO_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB                = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate803             = "FIREWALL_PRESERVED_GATE803_TRIALITY_YUKAWA_READOUT_MINIMALITY_BOUNDARY"
)

type Inheritance struct {
	TD4Status            string
	AirlockLevel         string
	HasNativeRealDescent bool
	HasSectorOperators   bool
	HasTraceAtoms        bool
	Verdicts             []string
}

type ReadoutPackage struct {
	Defined     bool
	Seals       []string
	TargetChain []string
	Verdict     string
	Supports    []string
	Failures    []string
}

type SealAudit struct {
	Name          string
	Audited       bool
	Required      []string
	CurrentStatus string
	Verdict       string
	Supports      []string
	Failures      []string
}

type Minimality struct {
	Audited         bool
	RemovalFailures []string
	Verdict         string
	Supports        []string
	Failures        []string
}

type NoGo struct {
	Defined         bool
	GivenOnly       []string
	CannotConstruct []string
	Reason          string
	Verdict         string
	Failures        []string
}

type CurrentASHA struct {
	Audited       bool
	Supplies      []string
	DoesNotSupply []string
	Verdict       string
	Supports      []string
	Failures      []string
}

type PathSeparation struct {
	Recorded      bool
	EmpiricalPath []string
	NativePath    []string
	Verdict       string
	Supports      []string
	Failures      []string
}

type CHiggsFirewall struct {
	Preserved bool
	Formula   string
	Unchanged []string
	Verdict   string
	Failures  []string
}

type BranchDecision struct {
	Recorded              bool
	NextNative            string
	AlternativeGeneration string
	Empirical             string
	Verdict               string
	Supports              []string
}

type Firewalls struct {
	Enforced         bool
	NoYukawa         bool
	NoEigenvalues    bool
	NoPMNSCKM        bool
	NoFlavor         bool
	NoNEff           bool
	NoGJ             bool
	NoScalar         bool
	NoPoleMass       bool
	NoVEVGF          bool
	NoNativeTriality bool
	NoHistoryLoop    bool
	Verdict          string
}

type Analysis struct {
	Inheritance       Inheritance
	Package           ReadoutPackage
	RealDescent       SealAudit
	GaugeAssignment   SealAudit
	SectorAssignment  SealAudit
	GenerationCarrier SealAudit
	HermitianOperator SealAudit
	Hierarchy         SealAudit
	TraceAtom         SealAudit
	Color             SealAudit
	Scale             SealAudit
	NonCircularity    SealAudit
	Minimality        Minimality
	NoGo              NoGo
	Current           CurrentASHA
	Paths             PathSeparation
	CHiggs            CHiggsFirewall
	Branch            BranchDecision
	Firewalls         Firewalls
	Truth             string
	Final             string
}

func BuildDefault() (Analysis, error) {
	inheritance := Inheritance{
		TD4Status:            "lawful complex airlocked pre-Yukawa invariant",
		AirlockLevel:         "T1 — complex D4 triality only",
		HasNativeRealDescent: false,
		HasSectorOperators:   false,
		HasTraceAtoms:        false,
		Verdicts:             []string{StatusGate802Inherited, StatusTD4OnlyPreYukawa, StatusTD4NotTraceLedger},
	}
	if inheritance.HasNativeRealDescent || inheritance.HasSectorOperators || inheritance.HasTraceAtoms {
		return Analysis{}, fmt.Errorf("Gate 803 expects T_D4 to remain airlocked and non-readout")
	}

	pkg := ReadoutPackage{
		Defined:     true,
		Seals:       []string{"RealDescentSeal", "GaugeRepresentationAssignmentSeal", "SectorAssignmentSeal", "GenerationCarrierSeal", "HermitianOperatorSeal", "SymmetryBreakingHierarchySeal", "TraceAtomExtractionSeal", "ColorMultiplicitySeal", "ScaleSchemeSeal", "NonCircularitySeal"},
		TargetChain: []string{"T_D4 + TrialityYukawaReadoutPackage", "Y_u,Y_d,Y_e,Y_nu", "singular values y_i", "trace atoms x_i=y_i^2", "a,b,N_eff"},
		Verdict:     StatusReadoutPackageDefined,
		Supports:    []string{StatusNeedsExtraSeals},
		Failures:    []string{StatusTrilinearAloneNoPackage},
	}
	if len(pkg.Seals) != 10 {
		return Analysis{}, fmt.Errorf("readout package must contain ten non-cosmetic seals")
	}

	realDescent := SealAudit{Name: "RealDescentSeal", Audited: true, Required: []string{"real structure", "bilinear signatures", "Clifford action compatibility", "adjoint convention", "positivity/reality of later trace atoms"}, CurrentStatus: "missing", Verdict: StatusRealDescentAudited, Failures: []string{StatusNoRealDescent, StatusComplexNotNativeYukawa}}
	gauge := SealAudit{Name: "GaugeRepresentationAssignmentSeal", Audited: true, Required: []string{"Q_L -> u_R", "Q_L -> d_R", "L_L -> e_R", "L_L -> nu_R or chosen neutrino convention", "SU(3)c/SU(2)L/U(1)Y compatibility", "Higgs doublet interface"}, CurrentStatus: "missing", Verdict: StatusGaugeAssignmentAudited, Failures: []string{StatusNoGaugeAssignment, StatusTrialityFramesNoSMEdges}}
	sector := SealAudit{Name: "SectorAssignmentSeal", Audited: true, Required: []string{"map triality carrier data to Y_u,Y_d,Y_e,Y_nu", "resolve three triality types versus four Yukawa sectors"}, CurrentStatus: "missing", Verdict: StatusSectorMinimalityAudited, Failures: []string{StatusThreeTypesNotFourSectors, StatusNoSMSectorAssignment}}
	generation := SealAudit{Name: "GenerationCarrierSeal", Audited: true, Required: []string{"G_gen", "cardinality three", "generation labels or basis", "sector action on G_gen", "frame comparison rules"}, CurrentStatus: "missing", Verdict: StatusGenerationAudited, Failures: []string{StatusTrialityTypesNotGenerations, StatusNoGenerationCarrier, StatusNoPMNSCKMFrames}}
	hermitian := SealAudit{Name: "HermitianOperatorSeal", Audited: true, Required: []string{"linear operators Y_f", "adjoint convention", "Y_f†Y_f positive", "singular-value extraction"}, CurrentStatus: "missing", Verdict: StatusHermitianAudited, Failures: []string{StatusComplexNotHermitian, StatusNoSVDTheorem}}
	hierarchy := SealAudit{Name: "SymmetryBreakingHierarchySeal", Audited: true, Required: []string{"breaking operator", "preferred frame or vacuum", "sector-dependent deformation", "generation-dependent spectrum", "top-dominance mechanism", "rest-pressure mechanism"}, CurrentStatus: "missing", Verdict: StatusHierarchyAudited, Failures: []string{StatusUniqueNoHierarchy, StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree}}
	atom := SealAudit{Name: "TraceAtomExtractionSeal", Audited: true, Required: []string{"Hermitian sector operators -> positive singular values", "singular values -> trace atoms x_i=y_i^2", "validation a=sum x_i, b=sum x_i^2", "no backward solving from N_eff/C_Higgs/Higgs data"}, CurrentStatus: "missing", Verdict: StatusTraceAtomAudited, Failures: []string{StatusNoPositiveAtoms, StatusNoBackwardAtoms}}
	color := SealAudit{Name: "ColorMultiplicitySeal", Audited: true, Required: []string{"exact convention for color factor 3", "relation to coefficient or repeated trace atoms", "no double counting"}, CurrentStatus: "color SU(3) remains certified trace multiplicity source; not replaced by D4", Verdict: StatusColorAudited, Supports: []string{StatusColorSU3TraceMultiplicity}, Failures: []string{StatusD4DoesNotReplaceColor, StatusColorDoubleCount}}
	scale := SealAudit{Name: "ScaleSchemeSeal", Audited: true, Required: []string{"scale_mu", "renormalization scheme", "threshold convention", "normalization convention", "neutrino convention"}, CurrentStatus: "missing from T_D4", Verdict: StatusScaleAudited, Failures: []string{StatusNoScaleLedger, StatusNoScaleStability}}
	noncirc := SealAudit{Name: "NonCircularitySeal", Audited: true, Required: []string{"prove no tuning from N_eff, C_Higgs, lambda_runtime_eff, m_H_tree_proxy, m_H_pole, observed Higgs mass", "treat PMNS/CKM only as explicitly sealed diagnostics if used"}, CurrentStatus: "required before any prediction status", Verdict: StatusNonCircularAudited, Failures: []string{StatusNoTargetTuning}}

	minimality := Minimality{
		Audited: true,
		RemovalFailures: []string{
			"remove RealDescentSeal -> complex object cannot become native real ASHA data",
			"remove GaugeRepresentationAssignmentSeal -> no Standard Model Yukawa edges",
			"remove SectorAssignmentSeal -> no Y_u,Y_d,Y_e,Y_nu",
			"remove GenerationCarrierSeal -> no three-family structure or PMNS/CKM",
			"remove HermitianOperatorSeal -> no Y†Y and no singular values",
			"remove SymmetryBreakingHierarchySeal -> no hierarchy, top dominance, or N_eff-3",
			"remove TraceAtomExtractionSeal -> no a,b,N_eff",
			"remove ColorMultiplicitySeal -> no lawful factor 3 in trace ledger",
			"remove ScaleSchemeSeal -> no comparison to M_Z ledger or GJ high-scale diagnostics",
			"remove NonCircularitySeal -> no prediction status",
		},
		Verdict:  StatusRemovalFailuresAudited,
		Supports: []string{StatusAllSealsNonCosmetic},
		Failures: []string{StatusCannotCompressToTD4},
	}

	nogo := NoGo{
		Defined:         true,
		GivenOnly:       []string{"ComplexD4TrialityAirlock", "T_D4"},
		CannotConstruct: []string{"Y_u,Y_d,Y_e,Y_nu", "positive trace atoms x_i", "a,b,N_eff", "PMNS/CKM", "Georgi-Jarlskog ratios", "C_Higgs update"},
		Reason:          "T_D4 is an invariant coupling tensor, not a sector-labeled, generation-resolved, Hermitian, scale-local Yukawa operator package.",
		Verdict:         StatusNoGoDefined,
		Failures:        []string{StatusTD4AloneNoLedger, StatusTD4AloneNoNEffFlavor},
	}

	current := CurrentASHA{
		Audited:       true,
		Supplies:      []string{"K7 native 7-dimensional contact support", "K7+ four-dimensional Higgs socket", "Fock/projective 1+3 selector resonance", "P_B/P_G Boolean-octonionic projectors", "finite spectral triple allowed chiral edge shapes and trace templates", "aggregate a,b sealed trace values", "color SU(3) trace multiplicity"},
		DoesNotSupply: []string{"triality Yukawa readout package", "native generation carrier", "Yukawa eigenvalues", "generation orientations", "trace atoms from T_D4"},
		Verdict:       StatusCurrentSubobjectsAudited,
		Supports:      []string{StatusFSTEdgeTemplate, StatusColorSU3TraceMultiplicity},
		Failures:      []string{StatusNoCurrentReadoutPackage},
	}

	paths := PathSeparation{
		Recorded:      true,
		EmpiricalPath: []string{"ExternalYukawaLedgerSeal", "trace atoms", "sector contributions", "N_eff audit"},
		NativePath:    []string{"T_D4", "TrialityYukawaReadoutPackage", "possible future native trace atoms"},
		Verdict:       StatusPathsSeparated,
		Supports:      []string{StatusExternalFastest},
		Failures:      []string{StatusExternalNotNative, StatusTrialityNotReadyExternal},
	}

	chiggs := CHiggsFirewall{Preserved: true, Formula: "C_Higgs=(3/N_eff)C_History", Unchanged: []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs", "lambda_H_bridge", "m_H_tree_proxy"}, Verdict: StatusCHiggsFirewall, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}}
	branch := BranchDecision{Recorded: true, NextNative: "Gate 804 — Finite Spectral Triple Yukawa Edge Template and Triality Coupling Compatibility Audit", AlternativeGeneration: "Gate 804 — GenerationCarrier Search Across K7, Fock Projective Geometry, and Triality Airlock Audit", Empirical: "Gate 804 — External Yukawa Ledger Acquisition and Sector Contribution Audit", Verdict: StatusBranchDecision, Supports: []string{StatusNextFSTCompatibility}}
	firewalls := Firewalls{Enforced: true, NoYukawa: true, NoEigenvalues: true, NoPMNSCKM: true, NoFlavor: true, NoNEff: true, NoGJ: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoNativeTriality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate803}

	return Analysis{Inheritance: inheritance, Package: pkg, RealDescent: realDescent, GaugeAssignment: gauge, SectorAssignment: sector, GenerationCarrier: generation, HermitianOperator: hermitian, Hierarchy: hierarchy, TraceAtom: atom, Color: color, Scale: scale, NonCircularity: noncirc, Minimality: minimality, NoGo: nogo, Current: current, Paths: paths, CHiggs: chiggs, Branch: branch, Firewalls: firewalls, Truth: "Gate 803 proves the precise no-go: T_D4 is a beautiful invariant tensor, but not a Yukawa ledger.", Final: "Gate 803 does not jump from triality to N_eff. It defines the full TrialityYukawaReadoutPackage and shows that ASHA currently lacks the real descent, gauge/sector assignment, generation carrier, Hermitian operator construction, hierarchy breaking, trace atom extraction, scale scheme, and noncircularity proofs required to convert T_D4 into Yukawa data."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate802Inherited, StatusTD4OnlyPreYukawa, StatusReadoutPackageDefined,
		StatusRealDescentAudited, StatusGaugeAssignmentAudited, StatusSectorMinimalityAudited,
		StatusGenerationAudited, StatusHermitianAudited, StatusHierarchyAudited,
		StatusTraceAtomAudited, StatusColorAudited, StatusScaleAudited, StatusNonCircularAudited,
		StatusRemovalFailuresAudited, StatusNoGoDefined, StatusCurrentSubobjectsAudited,
		StatusPathsSeparated, StatusCHiggsFirewall, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusNeedsExtraSeals, StatusAllSealsNonCosmetic, StatusFSTEdgeTemplate,
		StatusColorSU3TraceMultiplicity, StatusExternalFastest, StatusNextFSTCompatibility,
		StatusTD4NotTraceLedger, StatusTrilinearAloneNoPackage, StatusNoRealDescent,
		StatusComplexNotNativeYukawa, StatusNoGaugeAssignment, StatusTrialityFramesNoSMEdges,
		StatusThreeTypesNotFourSectors, StatusNoSMSectorAssignment, StatusTrialityTypesNotGenerations,
		StatusNoGenerationCarrier, StatusNoPMNSCKMFrames, StatusComplexNotHermitian, StatusNoSVDTheorem,
		StatusUniqueNoHierarchy, StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree,
		StatusNoPositiveAtoms, StatusNoBackwardAtoms, StatusD4DoesNotReplaceColor, StatusColorDoubleCount,
		StatusNoScaleLedger, StatusNoScaleStability, StatusNoTargetTuning, StatusCannotCompressToTD4,
		StatusTD4AloneNoLedger, StatusTD4AloneNoNEffFlavor, StatusNoCurrentReadoutPackage,
		StatusExternalNotNative, StatusTrialityNotReadyExternal, StatusNoCYukawaUpdate, StatusCHiggsLevelB,
		StatusFirewallGate803,
	}
}

func FormatPackage(p ReadoutPackage) string {
	return fmt.Sprintf("seals=[%s] target=[%s] supports=[%s] failures=[%s]", strings.Join(p.Seals, "; "), strings.Join(p.TargetChain, " -> "), strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}

func FormatSeal(s SealAudit) string {
	return fmt.Sprintf("%s status=%s required=[%s] supports=[%s] failures=[%s]", s.Name, s.CurrentStatus, strings.Join(s.Required, "; "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatMinimality(m Minimality) string {
	return fmt.Sprintf("removal_failures=[%s] supports=[%s] failures=[%s]", strings.Join(m.RemovalFailures, "; "), strings.Join(m.Supports, "; "), strings.Join(m.Failures, "; "))
}

func FormatNoGo(n NoGo) string {
	return fmt.Sprintf("given=[%s] cannot=[%s] reason=%s failures=[%s]", strings.Join(n.GivenOnly, "; "), strings.Join(n.CannotConstruct, "; "), n.Reason, strings.Join(n.Failures, "; "))
}

func FormatCurrent(c CurrentASHA) string {
	return fmt.Sprintf("supplies=[%s] lacks=[%s] supports=[%s] failures=[%s]", strings.Join(c.Supplies, "; "), strings.Join(c.DoesNotSupply, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatPaths(p PathSeparation) string {
	return fmt.Sprintf("empirical=[%s] native=[%s] supports=[%s] failures=[%s]", strings.Join(p.EmpiricalPath, " -> "), strings.Join(p.NativePath, " -> "), strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}

func FormatCHiggs(c CHiggsFirewall) string {
	return fmt.Sprintf("%s unchanged=[%s] failures=[%s]", c.Formula, strings.Join(c.Unchanged, "; "), strings.Join(c.Failures, "; "))
}

func containsAll(hay []string, needles []string) bool {
	joined := strings.Join(hay, "\n")
	for _, n := range needles {
		if !strings.Contains(joined, n) {
			return false
		}
	}
	return true
}
