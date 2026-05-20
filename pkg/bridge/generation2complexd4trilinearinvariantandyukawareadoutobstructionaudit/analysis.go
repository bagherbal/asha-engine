// Package generation2complexd4trilinearinvariantandyukawareadoutobstructionaudit implements
// Gate 802: Complex D4 Trilinear Invariant and Yukawa Readout Obstruction Audit.
//
// Gate 802 keeps the D4 branch inside the Gate 801 ComplexD4TrialityAirlock.
// It records the lawful complex Spin(8)/D4 trilinear invariant, audits its
// representation-theoretic role, and explicitly blocks any premature promotion
// to a Yukawa trace ledger, N_eff theorem, generation theorem, PMNS/CKM source,
// Georgi-Jarlskog theorem, scalar-Higgs update, or native real Cl(1,7) theorem.
package generation2complexd4trilinearinvariantandyukawareadoutobstructionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE802-COMPLEX-D4-TRILINEAR-INVARIANT-YUKAWA-READOUT-OBSTRUCTION-AUDIT"

	StatusGate801Inherited           = "PASS_GATE801_REAL_FORM_TRIALITY_AIRLOCK_INHERITED"
	StatusComplexD4AirlockInherited  = "PASS_COMPLEX_D4_AIRLOCK_STATUS_INHERITED"
	StatusComplexCarriersDefined     = "PASS_COMPLEX_TRIALITY_CARRIERS_DEFINED"
	StatusTrilinearDefined           = "PASS_COMPLEX_D4_TRILINEAR_INVARIANT_DEFINED"
	StatusMultiplicityAuditDefined   = "PASS_INVARIANT_MULTIPLICITY_AUDIT_DEFINED"
	StatusTrialityCovarianceDefined  = "PASS_TRIALITY_COVARIANCE_AUDIT_DEFINED"
	StatusReadoutRequirementsDefined = "PASS_TRIALITY_YUKAWA_READOUT_PACKAGE_REQUIREMENTS_DEFINED"
	StatusSectorObstructionAudited   = "PASS_SECTOR_ASSIGNMENT_OBSTRUCTION_AUDITED"
	StatusGenerationObstruction      = "PASS_GENERATION_OBSTRUCTION_AUDITED"
	StatusPositivityObstruction      = "PASS_POSITIVITY_AND_SINGULAR_VALUE_OBSTRUCTION_AUDITED"
	StatusTopDominanceObstruction    = "PASS_TOP_DOMINANCE_BREAKING_OBSTRUCTION_AUDITED"
	StatusGJReadoutObstruction       = "PASS_GEORGI_JARLSKOG_READOUT_OBSTRUCTION_AUDITED"
	StatusRealFormObstruction        = "PASS_REAL_FORM_OBSTRUCTION_PRESERVED"
	StatusLawfulUseRecorded          = "PASS_LAWFUL_USE_OF_T_D4_RECORDED"
	StatusCHiggsFirewallPreserved    = "PASS_C_HIGGS_FIREWALL_PRESERVED"
	StatusOutcomeClassification      = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchDecisionRecorded     = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls          = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusThreeEightDimTypes       = "CONDITIONAL_SUPPORT_COMPLEX_D4_HAS_THREE_EIGHT_DIMENSIONAL_REPRESENTATION_TYPES"
	StatusTD4PreYukawa             = "CONDITIONAL_SUPPORT_T_D4_IS_LAWFUL_COMPLEX_PRE_YUKAWA_INVARIANT_CANDIDATE"
	StatusUniqueInvariantShape     = "CONDITIONAL_SUPPORT_UNIQUE_TRILINEAR_INVARIANT_WOULD_GIVE_CANONICAL_COMPLEX_COUPLING_SHAPE"
	StatusTrialityCovariance       = "CONDITIONAL_SUPPORT_T_D4_MAY_BE_COVARIANT_UNDER_COMPLEX_TRIALITY_ACTION"
	StatusTD4UsefulAirlocked       = "CONDITIONAL_SUPPORT_T_D4_IS_USEFUL_AS_AIRLOCKED_PRE_YUKAWA_SHAPE"
	StatusComplexBranchInteresting = "CONDITIONAL_SUPPORT_COMPLEX_TRILINEAR_BRANCH_IS_STRUCTURALLY_INTERESTING_BUT_READOUT_BLOCKED"

	StatusT1NotNative                 = "FAILED_ROUTE_T1_COMPLEX_TRIALITY_NOT_NATIVE_CL17_THEOREM"
	StatusT1NotYukawa                 = "FAILED_ROUTE_T1_COMPLEX_TRIALITY_NOT_YUKAWA_READOUT_THEOREM"
	StatusComplexTypesNotGenerations  = "FAILED_ROUTE_COMPLEX_REPRESENTATION_TYPES_NOT_GENERATION_COPIES"
	StatusComplexNotNativeCarriers    = "FAILED_ROUTE_COMPLEX_CARRIERS_NOT_NATIVE_REAL_CL17_CARRIERS"
	StatusTD4NotSMYukawa              = "FAILED_ROUTE_T_D4_NOT_YET_STANDARD_MODEL_YUKAWA_OPERATOR"
	StatusTD4NotTraceLedger           = "FAILED_ROUTE_T_D4_NOT_YET_YUKAWA_TRACE_LEDGER"
	StatusUniqueNoEigenvalues         = "FAILED_ROUTE_UNIQUE_INVARIANT_DOES_NOT_DETERMINE_YUKAWA_EIGENVALUES"
	StatusNormalizationNoHierarchy    = "FAILED_ROUTE_INVARIANT_NORMALIZATION_NOT_YUKAWA_HIERARCHY"
	StatusCovarianceNotGenerations    = "FAILED_ROUTE_TRIALITY_COVARIANCE_NOT_GENERATION_TRIPLICATION"
	StatusCovarianceNotMixing         = "FAILED_ROUTE_TRIALITY_COVARIANCE_NOT_PMNS_CKM_MISALIGNMENT"
	StatusTrilinearNoSectorOps        = "FAILED_ROUTE_TRILINEAR_INVARIANT_ALONE_DOES_NOT_DEFINE_SECTOR_OPERATORS"
	StatusTrilinearNoAtoms            = "FAILED_ROUTE_TRILINEAR_INVARIANT_ALONE_DOES_NOT_DEFINE_TRACE_ATOMS"
	StatusTrilinearNoNEff             = "FAILED_ROUTE_TRILINEAR_INVARIANT_ALONE_DOES_NOT_DEFINE_N_EFF"
	StatusNoTrialityToSMSectors       = "FAILED_ROUTE_NO_TRIALITY_TO_STANDARD_MODEL_SECTOR_ASSIGNMENT"
	StatusThreeFramesNotFourSectors   = "FAILED_ROUTE_THREE_TRIALITY_FRAMES_DO_NOT_MATCH_FOUR_YUKAWA_SECTORS"
	StatusNoGaugeAssignment           = "FAILED_ROUTE_NO_GAUGE_REPRESENTATION_ASSIGNMENT_FROM_T_D4"
	StatusTrialityTypesNotGenerations = "FAILED_ROUTE_TRIALITY_TYPES_NOT_GENERATION_COPIES"
	StatusNoGenerationCarrier         = "FAILED_ROUTE_NO_GENERATION_CARRIER_FROM_T_D4"
	StatusNoPMNSCKMReadout            = "FAILED_ROUTE_NO_PMNS_CKM_READOUT_FROM_TRIALITY_INVARIANT"
	StatusComplexAmplitudeNotAtom     = "FAILED_ROUTE_COMPLEX_TRILINEAR_AMPLITUDE_NOT_POSITIVE_YUKAWA_ATOM"
	StatusNoHermitianSectorOperator   = "FAILED_ROUTE_NO_HERMITIAN_SECTOR_OPERATOR_FROM_T_D4"
	StatusNoSingularValueExtraction   = "FAILED_ROUTE_NO_SINGULAR_VALUE_EXTRACTION_THEOREM"
	StatusTD4NoTopDominance           = "FAILED_ROUTE_T_D4_DOES_NOT_EXPLAIN_TOP_DOMINANCE"
	StatusTD4NoNEffMinusThree         = "FAILED_ROUTE_T_D4_DOES_NOT_EXPLAIN_N_EFF_MINUS_THREE"
	StatusTD4NoScaleStability         = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_SCALE_STABILITY"
	StatusTD4NoGJ                     = "FAILED_ROUTE_T_D4_DOES_NOT_DERIVE_GEORGI_JARLSKOG_RATIOS"
	StatusTD4NoHighScaleClebsch       = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_HIGH_SCALE_CLEBSCH_FACTORS"
	StatusGJStillNeedsLedger          = "FAILED_ROUTE_GJ_DIAGNOSTIC_STILL_REQUIRES_MULTISCALE_YUKAWA_LEDGER"
	StatusComplexTD4NotRealCL17       = "FAILED_ROUTE_COMPLEX_T_D4_NOT_NATIVE_REAL_CL17_INVARIANT_WITHOUT_DESCENT"
	StatusNoRealDescentReadout        = "FAILED_ROUTE_NO_REAL_DESCENT_MAP_FOR_YUKAWA_READOUT"
	StatusTD4NotPhysicalScalarInput   = "FAILED_ROUTE_T_D4_NOT_DIRECT_PHYSICAL_OR_SCALAR_INPUT"
	StatusTD4NoCYukawaUpdate          = "FAILED_ROUTE_TRIALITY_TRILINEAR_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsStillLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallPreservedGate802    = "FIREWALL_PRESERVED_GATE802_COMPLEX_D4_TRILINEAR_READOUT_OBSTRUCTION_BOUNDARY"
)

type Inheritance struct {
	ClAlgebra              string
	VolumeSquare           int
	RealChiralityCertified bool
	ComplexifiedSpin       string
	OuterAutomorphism      string
	TrialityLevel          string
	Verdicts               []string
}

type ComplexCarriers struct {
	Defined            bool
	VectorDimC         int
	SpinorPlusDimC     int
	SpinorMinusDimC    int
	Permutation        string
	GenerationCopies   bool
	NativeRealCarriers bool
	Verdict            string
	Supports           []string
	Failures           []string
}

type TrilinearInvariant struct {
	Defined     bool
	Formula     string
	GammaAction string
	Equivariant bool
	NonZero     bool
	Blocked     []string
	Verdict     string
	Supports    []string
	Failures    []string
}

type InvariantMultiplicity struct {
	Audited                 bool
	HomDimension            int
	CanonicalShapeUpToScale bool
	DeterminesEigenvalues   bool
	DeterminesHierarchy     bool
	Verdict                 string
	Supports                []string
	Failures                []string
}

type TrialityCovariance struct {
	Audited                bool
	CyclicStable           bool
	GenerationTriplication bool
	MixingReadout          bool
	Verdict                string
	Supports               []string
	Failures               []string
}

type ReadoutRequirements struct {
	Defined  bool
	Items    []string
	Verdict  string
	Failures []string
}

type Obstruction struct {
	Audited  bool
	Name     string
	Details  []string
	Verdict  string
	Failures []string
}

type RealFormObstruction struct {
	Preserved bool
	Requires  []string
	Verdict   string
	Failures  []string
}

type LawfulUse struct {
	Recorded bool
	Allowed  []string
	Blocked  []string
	Verdict  string
	Supports []string
	Failures []string
}

type CHiggsFirewall struct {
	Preserved bool
	Formula   string
	Unchanged []string
	Verdict   string
	Failures  []string
}

type Outcomes struct {
	Recorded bool
	Items    []string
	Verdict  string
	Support  string
}

type BranchDecision struct {
	Recorded    bool
	NextNative  string
	Alternative string
	Verdict     string
}

type Firewalls struct {
	Enforced      bool
	NoYukawa      bool
	NoEigenvalues bool
	NoPMNSCKM     bool
	NoFlavor      bool
	NoNEff        bool
	NoGJ          bool
	NoScalar      bool
	NoPoleMass    bool
	NoVEVGF       bool
	NoNativeCL17  bool
	NoHistoryLoop bool
	Verdict       string
}

type Analysis struct {
	Inheritance    Inheritance
	Carriers       ComplexCarriers
	Trilinear      TrilinearInvariant
	Multiplicity   InvariantMultiplicity
	Covariance     TrialityCovariance
	Readout        ReadoutRequirements
	Sector         Obstruction
	Generation     Obstruction
	Positivity     Obstruction
	TopDominance   Obstruction
	GeorgiJarlskog Obstruction
	RealForm       RealFormObstruction
	Lawful         LawfulUse
	CHiggs         CHiggsFirewall
	Outcome        Outcomes
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Analysis, error) {
	inherit := Inheritance{
		ClAlgebra: "Cl(1,7) ≅ Mat(16,R)", VolumeSquare: -1,
		RealChiralityCertified: false,
		ComplexifiedSpin:       "spin(1,7)_C ≅ so(8,C)",
		OuterAutomorphism:      "Out(D4)≅S3 after complexification",
		TrialityLevel:          "T1 — complex D4 triality only",
		Verdicts:               []string{StatusGate801Inherited, StatusComplexD4AirlockInherited, StatusT1NotNative, StatusT1NotYukawa},
	}
	if inherit.VolumeSquare != -1 || inherit.RealChiralityCertified || !strings.Contains(inherit.TrialityLevel, "T1") {
		return Analysis{}, fmt.Errorf("Gate 802 expects Gate 801 T1 complex-only airlock inheritance")
	}
	carriers := ComplexCarriers{
		Defined: true, VectorDimC: 8, SpinorPlusDimC: 8, SpinorMinusDimC: 8,
		Permutation:      "V_C ↔ S+_C ↔ S-_C under complex D4 outer action",
		GenerationCopies: false, NativeRealCarriers: false,
		Verdict:  StatusComplexCarriersDefined,
		Supports: []string{StatusThreeEightDimTypes},
		Failures: []string{StatusComplexTypesNotGenerations, StatusComplexNotNativeCarriers},
	}
	if carriers.VectorDimC != 8 || carriers.SpinorPlusDimC != 8 || carriers.SpinorMinusDimC != 8 {
		return Analysis{}, fmt.Errorf("complex D4 carrier dimensions must be 8,8,8")
	}
	trilinear := TrilinearInvariant{
		Defined:     true,
		Formula:     "T_D4(v,ψ+,ψ-)=<γ(v)ψ+,ψ->",
		GammaAction: "γ(v): S+_C -> S-_C",
		Equivariant: true,
		NonZero:     true,
		Blocked:     []string{"Standard Model Yukawa operator", "Yukawa singular-value ledger", "generation theorem", "N_eff readout"},
		Verdict:     StatusTrilinearDefined,
		Supports:    []string{StatusTD4PreYukawa},
		Failures:    []string{StatusTD4NotSMYukawa, StatusTD4NotTraceLedger},
	}
	multiplicity := InvariantMultiplicity{
		Audited: true, HomDimension: 1, CanonicalShapeUpToScale: true,
		DeterminesEigenvalues: false, DeterminesHierarchy: false,
		Verdict:  StatusMultiplicityAuditDefined,
		Supports: []string{StatusUniqueInvariantShape},
		Failures: []string{StatusUniqueNoEigenvalues, StatusNormalizationNoHierarchy},
	}
	covariance := TrialityCovariance{
		Audited: true, CyclicStable: true, GenerationTriplication: false, MixingReadout: false,
		Verdict:  StatusTrialityCovarianceDefined,
		Supports: []string{StatusTrialityCovariance},
		Failures: []string{StatusCovarianceNotGenerations, StatusCovarianceNotMixing},
	}
	readout := ReadoutRequirements{
		Defined: true,
		Items: []string{
			"sector assignment V_C,S+_C,S-_C -> Standard Model sector carriers or pre-sector carriers",
			"operator extraction T_D4 + symmetry-breaking data -> Y_u,Y_d,Y_e,Y_nu",
			"singular-value extraction Y_sector -> y_i",
			"trace atom map y_i -> x_i=y_i^2",
			"color/generation bookkeeping with color factor counted exactly once",
			"scale convention M_Z or high-scale ledger with RG transport",
			"positivity/reality condition from complex invariant to real positive atoms",
			"breaking/deformation explaining hierarchy, top dominance, and N_eff-3",
		},
		Verdict:  StatusReadoutRequirementsDefined,
		Failures: []string{StatusTrilinearNoSectorOps, StatusTrilinearNoAtoms, StatusTrilinearNoNEff},
	}
	sector := Obstruction{
		Audited: true, Name: "sector assignment obstruction",
		Details:  []string{"Y_u,Y_d,Y_e,Y_nu require gauge representation data, chirality, generation labels, color factors, and Higgs coupling edges", "complex triality carriers are only V_C,S+_C,S-_C"},
		Verdict:  StatusSectorObstructionAudited,
		Failures: []string{StatusNoTrialityToSMSectors, StatusThreeFramesNotFourSectors, StatusNoGaugeAssignment},
	}
	generation := Obstruction{
		Audited: true, Name: "generation obstruction",
		Details:  []string{"triality couples three representation types", "it does not supply three family copies G_gen", "PMNS/CKM require sector operators and diagonalization frames"},
		Verdict:  StatusGenerationObstruction,
		Failures: []string{StatusTrialityTypesNotGenerations, StatusNoGenerationCarrier, StatusNoPMNSCKMReadout},
	}
	positivity := Obstruction{
		Audited: true, Name: "positivity and singular-value obstruction",
		Details:  []string{"complex multilinear amplitude is not a positive atom", "Yukawa trace ledger needs x_i=y_i^2 ≥ 0", "requires Hermitian sector operators, adjoint convention, SVD, normalization, scale, and positivity proof"},
		Verdict:  StatusPositivityObstruction,
		Failures: []string{StatusComplexAmplitudeNotAtom, StatusNoHermitianSectorOperator, StatusNoSingularValueExtraction},
	}
	top := Obstruction{
		Audited: true, Name: "top-dominance and rest-pressure obstruction",
		Details:  []string{"certified baseline remains color-tripled top dominance", "T_D4 does not explain why one top-like colored channel dominates", "T_D4 does not explain N_eff-3 or scale stability"},
		Verdict:  StatusTopDominanceObstruction,
		Failures: []string{StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree, StatusTD4NoScaleStability},
	}
	gj := Obstruction{
		Audited: true, Name: "Georgi-Jarlskog readout obstruction",
		Details:  []string{"GJ requires high-scale down/lepton Yukawa ratios", "T_D4 supplies neither RG transport nor Clebsch coefficients nor GUT embedding"},
		Verdict:  StatusGJReadoutObstruction,
		Failures: []string{StatusTD4NoGJ, StatusTD4NoHighScaleClebsch, StatusGJStillNeedsLedger},
	}
	realForm := RealFormObstruction{
		Preserved: true,
		Requires:  []string{"RealDescentMap from complex triality invariant to real Cl(1,7) typed object", "preserve real structure", "preserve bilinear signatures", "preserve Clifford compatibility", "preserve positivity/reality of trace atoms", "preserve sector readout"},
		Verdict:   StatusRealFormObstruction,
		Failures:  []string{StatusComplexTD4NotRealCL17, StatusNoRealDescentReadout},
	}
	lawful := LawfulUse{
		Recorded: true,
		Allowed:  []string{"provide an airlocked pre-Yukawa coupling shape", "test whether later sector assignment can be equivariant", "guide representation-theoretic constraint search", "define future obstruction target"},
		Blocked:  []string{"derive N_eff", "derive top dominance", "derive Yukawa eigenvalues", "derive PMNS/CKM", "derive Georgi-Jarlskog", "modify C_Higgs"},
		Verdict:  StatusLawfulUseRecorded,
		Supports: []string{StatusTD4UsefulAirlocked},
		Failures: []string{StatusTD4NotPhysicalScalarInput},
	}
	chiggs := CHiggsFirewall{
		Preserved: true,
		Formula:   "C_Higgs=(3/N_eff)C_History",
		Unchanged: []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs", "lambda_H_bridge", "m_H_tree_proxy"},
		Verdict:   StatusCHiggsFirewallPreserved,
		Failures:  []string{StatusTD4NoCYukawaUpdate, StatusCHiggsStillLevelB},
	}
	outcome := Outcomes{
		Recorded: true,
		Items:    []string{"O1 complex D4 trilinear invariant is lawful as airlocked pre-Yukawa shape", "O2 no Yukawa trace-readout package certified", "O3 no native real Cl(1,7) descent certified", "O4 no N_eff, PMNS/CKM, GJ, or scalar-Higgs update follows"},
		Verdict:  StatusOutcomeClassification,
		Support:  StatusComplexBranchInteresting,
	}
	branch := BranchDecision{
		Recorded:    true,
		NextNative:  "Gate 803 — Triality-to-Yukawa Readout Package Minimality and No-Go Audit",
		Alternative: "Gate 803 — External Yukawa Ledger Acquisition and Sector Contribution Audit",
		Verdict:     StatusBranchDecisionRecorded,
	}
	firewalls := Firewalls{Enforced: true, NoYukawa: true, NoEigenvalues: true, NoPMNSCKM: true, NoFlavor: true, NoNEff: true, NoGJ: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoNativeCL17: true, NoHistoryLoop: true, Verdict: StatusFirewallPreservedGate802}
	return Analysis{Inheritance: inherit, Carriers: carriers, Trilinear: trilinear, Multiplicity: multiplicity, Covariance: covariance, Readout: readout, Sector: sector, Generation: generation, Positivity: positivity, TopDominance: top, GeorgiJarlskog: gj, RealForm: realForm, Lawful: lawful, CHiggs: chiggs, Outcome: outcome, Branch: branch, Firewalls: firewalls, Truth: "Gate 802 certifies T_D4 only as a complex-airlocked pre-Yukawa invariant shape; it does not supply sector operators, trace atoms, N_eff, PMNS/CKM, GJ ratios, or native real Cl(1,7) descent.", Final: "Gate 802 keeps the triality branch precise: the complex D4 trilinear invariant is lawful as an airlocked pre-Yukawa coupling shape, but it does not provide the missing physics. The next native question is the minimal TrialityYukawaReadoutPackage needed to turn the trilinear into sector operators and trace atoms, and whether ASHA can supply any part of it."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate801Inherited, StatusComplexD4AirlockInherited, StatusComplexCarriersDefined,
		StatusTrilinearDefined, StatusMultiplicityAuditDefined, StatusTrialityCovarianceDefined,
		StatusReadoutRequirementsDefined, StatusSectorObstructionAudited, StatusGenerationObstruction,
		StatusPositivityObstruction, StatusTopDominanceObstruction, StatusGJReadoutObstruction,
		StatusRealFormObstruction, StatusLawfulUseRecorded, StatusCHiggsFirewallPreserved,
		StatusOutcomeClassification, StatusBranchDecisionRecorded, StatusPhysicalFirewalls,
		StatusThreeEightDimTypes, StatusTD4PreYukawa, StatusUniqueInvariantShape,
		StatusTrialityCovariance, StatusTD4UsefulAirlocked, StatusComplexBranchInteresting,
		StatusT1NotNative, StatusT1NotYukawa, StatusComplexTypesNotGenerations,
		StatusComplexNotNativeCarriers, StatusTD4NotSMYukawa, StatusTD4NotTraceLedger,
		StatusUniqueNoEigenvalues, StatusNormalizationNoHierarchy, StatusCovarianceNotGenerations,
		StatusCovarianceNotMixing, StatusTrilinearNoSectorOps, StatusTrilinearNoAtoms,
		StatusTrilinearNoNEff, StatusNoTrialityToSMSectors, StatusThreeFramesNotFourSectors,
		StatusNoGaugeAssignment, StatusTrialityTypesNotGenerations, StatusNoGenerationCarrier,
		StatusNoPMNSCKMReadout, StatusComplexAmplitudeNotAtom, StatusNoHermitianSectorOperator,
		StatusNoSingularValueExtraction, StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree,
		StatusTD4NoScaleStability, StatusTD4NoGJ, StatusTD4NoHighScaleClebsch,
		StatusGJStillNeedsLedger, StatusComplexTD4NotRealCL17, StatusNoRealDescentReadout,
		StatusTD4NotPhysicalScalarInput, StatusTD4NoCYukawaUpdate, StatusCHiggsStillLevelB,
		StatusFirewallPreservedGate802,
	}
}

func FormatCarriers(c ComplexCarriers) string {
	return fmt.Sprintf("V_C=%d S+_C=%d S-_C=%d permutation=%s generation_copies=%v native_real=%v supports=[%s] failures=[%s]", c.VectorDimC, c.SpinorPlusDimC, c.SpinorMinusDimC, c.Permutation, c.GenerationCopies, c.NativeRealCarriers, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatTrilinear(t TrilinearInvariant) string {
	return fmt.Sprintf("%s; gamma=%s; equivariant=%v nonzero=%v blocked=[%s] supports=[%s] failures=[%s]", t.Formula, t.GammaAction, t.Equivariant, t.NonZero, strings.Join(t.Blocked, "; "), strings.Join(t.Supports, "; "), strings.Join(t.Failures, "; "))
}

func FormatMultiplicity(m InvariantMultiplicity) string {
	return fmt.Sprintf("Hom-dim=%d canonical_up_to_scale=%v determines_eigenvalues=%v determines_hierarchy=%v supports=[%s] failures=[%s]", m.HomDimension, m.CanonicalShapeUpToScale, m.DeterminesEigenvalues, m.DeterminesHierarchy, strings.Join(m.Supports, "; "), strings.Join(m.Failures, "; "))
}

func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("%s details=[%s] failures=[%s]", o.Name, strings.Join(o.Details, "; "), strings.Join(o.Failures, "; "))
}

func FormatReadout(r ReadoutRequirements) string {
	return fmt.Sprintf("items=[%s] failures=[%s]", strings.Join(r.Items, "; "), strings.Join(r.Failures, "; "))
}

func FormatRealForm(r RealFormObstruction) string {
	return fmt.Sprintf("preserved=%v requires=[%s] failures=[%s]", r.Preserved, strings.Join(r.Requires, "; "), strings.Join(r.Failures, "; "))
}

func FormatLawfulUse(l LawfulUse) string {
	return fmt.Sprintf("allowed=[%s] blocked=[%s] supports=[%s] failures=[%s]", strings.Join(l.Allowed, "; "), strings.Join(l.Blocked, "; "), strings.Join(l.Supports, "; "), strings.Join(l.Failures, "; "))
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
