// Package generation2levelbchiggserrorbudgetandindependentinputsensitivityaudit implements
// Gate 792: Level-B Error Budget and Independent-Input Sensitivity Audit.
//
// Gate 791 made C_Higgs a clean Level-B Fermi-normalized test interface.
// Gate 792 quantifies first-order numerical sensitivities, controlled component
// removals, error-budget classes, and theorem-pressure priorities while keeping
// all PMNS/CKM, Yukawa, HistoryLoop, VEV/G_F, scalar-runtime, and pole-mass
// firewalls intact.
package generation2levelbchiggserrorbudgetandindependentinputsensitivityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE792-LEVEL-B-C-HIGGS-ERROR-BUDGET-INDEPENDENT-INPUT-SENSITIVITY-AUDIT"

	StatusGate791Inherited                    = "PASS_GATE791_LEVEL_B_TEST_INTERFACE_INHERITED"
	StatusAnalyticSensitivityComputed         = "PASS_ANALYTIC_SENSITIVITY_FORMULAS_COMPUTED"
	StatusRelativeElasticityCompleted         = "PASS_RELATIVE_ELASTICITY_AUDIT_COMPLETED"
	StatusAbsolutePerturbationComputed        = "PASS_ABSOLUTE_PERTURBATION_LEDGER_COMPUTED"
	StatusComponentRemovalDiagnosticsComputed = "PASS_COMPONENT_REMOVAL_DIAGNOSTICS_COMPUTED"
	StatusNumericalVsTheoremSeparated         = "PASS_NUMERICAL_SENSITIVITY_VERSUS_THEOREM_PRESSURE_SEPARATED"
	StatusErrorBudgetCategoriesDefined        = "PASS_ERROR_BUDGET_CATEGORIES_DEFINED"
	StatusMajorInputsClassified               = "PASS_MAJOR_INPUTS_CLASSIFIED_BY_ERROR_TYPE"
	StatusScientificTestabilityCompleted      = "PASS_SCIENTIFIC_TESTABILITY_AUDIT_COMPLETED"
	StatusNextBranchRecommendationRecorded    = "PASS_NEXT_BRANCH_RECOMMENDATION_RECORDED"
	StatusPhysicalFirewallsEnforced           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusExplicitLinearizedErrorBudget           = "CONDITIONAL_SUPPORT_C_HIGGS_HAS_EXPLICIT_LINEARIZED_ERROR_BUDGET"
	StatusNEffUnitRelativeLeverage                = "CONDITIONAL_SUPPORT_N_EFF_HAS_UNIT_RELATIVE_LEVERAGE_ON_C_HIGGS"
	StatusLHopfSecondSensitivity                  = "CONDITIONAL_SUPPORT_L_HOPF_IS_SECOND_NUMERICAL_SENSITIVITY_CHANNEL"
	StatusAbsoluteResponseStrongestToLAndN        = "CONDITIONAL_SUPPORT_ABSOLUTE_C_HIGGS_RESPONSE_IS_STRONGEST_TO_L_HOPF_AND_N_EFF"
	StatusYukawaParticipationNumericallyImportant = "CONDITIONAL_SUPPORT_YUKAWA_PARTICIPATION_IS_NUMERICALLY_IMPORTANT_RELATIVE_TO_BOUNDARY_MICROCORRECTIONS"
	StatusBoundaryMicroLowNumericalLeverage       = "CONDITIONAL_SUPPORT_BOUNDARY_MICROSTRUCTURE_IS_STRUCTURALLY_IMPORTANT_BUT_LOW_NUMERICAL_LEVERAGE_FOR_C_HIGGS"
	StatusNEffTopNumericalLeverage                = "CONDITIONAL_SUPPORT_N_EFF_IS_TOP_NUMERICAL_LEVERAGE_TARGET"
	StatusKappaOrientTopFlavorObstruction         = "CONDITIONAL_SUPPORT_KAPPA_ORIENT_IS_TOP_FLAVOR_THEOREM_OBSTRUCTION"
	StatusLHopfTopHistoryObstruction              = "CONDITIONAL_SUPPORT_L_HOPF_IS_TOP_HISTORY_TRANSPORT_OBSTRUCTION"
	StatusNEffReductionImprovesTestability        = "CONDITIONAL_SUPPORT_N_EFF_REDUCTION_MOST_IMPROVES_NUMERICAL_TESTABILITY"
	StatusGenerationAndHistoryImproveNative       = "CONDITIONAL_SUPPORT_GENERATION_MIXING_AND_HISTORYLOOP_REDUCTION_MOST_IMPROVE_NATIVE_STATUS"
	StatusTreeToPoleRequired                      = "CONDITIONAL_SUPPORT_TREE_TO_POLE_PACKAGE_REQUIRED_FOR_PHYSICAL_COMPARISON"
	StatusNEffNextNumericalBranch                 = "CONDITIONAL_SUPPORT_N_EFF_SOURCE_REDUCTION_IS_BEST_NEXT_NUMERICAL_TESTABILITY_BRANCH"

	StatusErrorBudgetNotNativeTheorem      = "FAILED_ROUTE_ERROR_BUDGET_NOT_NATIVE_THEOREM"
	StatusNEffNotNativeYukawa              = "FAILED_ROUTE_N_EFF_NOT_NATIVE_YUKAWA_THEOREM"
	StatusLHopfNotNativeHistoryLoop        = "FAILED_ROUTE_L_HOPF_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusKappaOrientNotNativeFlavor       = "FAILED_ROUTE_KAPPA_ORIENT_NOT_NATIVE_FLAVOR_THEOREM"
	StatusFWallNotNativeBoundary           = "FAILED_ROUTE_F_WALL_3_RED_NOT_NATIVE_BOUNDARY_RESPONSE_THEOREM"
	StatusLevelBSensitivityNotLevelC       = "FAILED_ROUTE_LEVEL_B_SENSITIVITY_NOT_LEVEL_C_PREDICTION"
	StatusTreeProxyNotPoleMass             = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusPoleComparisonRequiresCorrection = "FAILED_ROUTE_POLE_COMPARISON_REQUIRES_CORRECTION_PACKAGE"
	StatusObservedHiggsForbiddenInBudget   = "FAILED_ROUTE_OBSERVED_HIGGS_MASS_MUST_NOT_SOURCE_ERROR_BUDGET"
	StatusFirewallPreservedGate792         = "FIREWALL_PRESERVED_GATE792_LEVEL_B_ERROR_BUDGET_BOUNDARY"
)

const (
	pSnapshot              = 7.0 / 72.0
	sSnapshot              = 0.0012924448188162962
	xiBoundarySnapshot     = 0.0503471644870914
	kappaOrientSnapshot    = 0.00550633006471245
	kappaBoundarySnapshot  = -2.775846236678231e-6
	kappaERedSnapshot      = 0.005503554218475772
	fWall3Snapshot         = 0.00012565521035653708
	absLambdaSnapshot      = 0.049700942077680596
	kappaLambdaRedSnapshot = 0.04432304306956136
	lHopfSnapshot          = 0.039788735772973836
	nEffSnapshot           = 3.0023273474722147
	cYukawaSnapshot        = 0.9992248188812008
	cHistorySnapshot       = 1.038025177923625
	cHiggsSnapshot         = 1.0372205204048603
)

type Gate791Inheritance struct {
	Inherited       bool
	CleanTestObject bool
	LevelBOnly      bool
	Verdict         string
}

type AnalyticSensitivity struct {
	Computed              bool
	DCByNEff              float64
	DCByLHopf             float64
	DCByKappaLambdaRed    float64
	DCByAbsLambda         float64
	DCByKappaOrient       float64
	DCByS                 float64
	DCByXiBoundary        float64
	FormulaNEff           string
	FormulaLHopf          string
	FormulaKappaLambdaRed string
	Verdict               string
}

type RelativeElasticity struct {
	Completed          bool
	ENEff              float64
	ELHopf             float64
	EAbsLambda         float64
	EKappaLambdaRed    float64
	EKappaOrient       float64
	ES                 float64
	EXiBoundary        float64
	NEffUnitLeverage   bool
	LHopfSecondChannel bool
	Verdict            string
}

type AbsolutePerturbationLedger struct {
	Computed             bool
	Perturbation         float64
	DeltaFromNEff        float64
	DeltaFromLHopf       float64
	DeltaFromKappaOrient float64
	DeltaFromAbsLambda   float64
	DeltaFromS           float64
	DeltaFromXiBoundary  float64
	Strongest            []string
	Verdict              string
}

type ComponentRemovalDiagnostics struct {
	Computed                          bool
	CHiggsTopColor                    float64
	TopColorShift                     float64
	CHiggsBoundaryCorrectionRemoved   float64
	BoundaryCorrectionShift           float64
	CHiggsCubicRemoved                float64
	CubicRemovalShift                 float64
	YukawaImportant                   bool
	BoundaryMicroLowNumericalLeverage bool
	Verdict                           string
}

type SensitivityVsTheoremPressure struct {
	Separated                bool
	NumericalRanking         []string
	TheoremPressureRanking   []string
	NEffTopNumerical         bool
	KappaOrientTopFlavor     bool
	LHopfTopHistory          bool
	FWallLowLeverageButTyped bool
	Verdict                  string
}

type ErrorBudgetCategories struct {
	Defined               bool
	Categories            []string
	Classifications       []string
	NEffClassified        bool
	KappaOrientClassified bool
	LHopfClassified       bool
	FWallClassified       bool
	ComparisonClassified  bool
	Verdict               string
}

type ScientificTestability struct {
	Completed                 bool
	NumericalSharpnessTarget  string
	NativeClosureTargets      []string
	EmpiricalComparisonTarget string
	NEffBestNumerical         bool
	GenerationHistoryNative   bool
	TreeToPoleRequired        bool
	Verdict                   string
}

type NextBranchRecommendation struct {
	Recorded     bool
	Recommended  string
	Reason       string
	Alternatives []string
	Verdict      string
}

type Firewalls struct {
	Enforced                         bool
	ErrorBudgetNativeTheorem         bool
	LargestSensitivityDeepestTheorem bool
	KappaOrientSmallMeansUnimportant bool
	FWallLowLeverageDisposable       bool
	CHiggsPoleMassPrediction         bool
	TreeProxyPoleMass                bool
	LevelBLevelCPrediction           bool
	NEffNativeYukawa                 bool
	LHopfNativeHistoryLoop           bool
	KappaOrientNativeFlavor          bool
	ObservedHiggsSourcesBudget       bool
	Verdict                          string
}

type Analysis struct {
	Gate791        Gate791Inheritance
	Sensitivity    AnalyticSensitivity
	Elasticity     RelativeElasticity
	Perturbations  AbsolutePerturbationLedger
	Removal        ComponentRemovalDiagnostics
	Pressure       SensitivityVsTheoremPressure
	ErrorBudget    ErrorBudgetCategories
	Testability    ScientificTestability
	Next           NextBranchRecommendation
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

func BuildDefault() (Analysis, error) {
	cYukawa := 3.0 / nEffSnapshot
	cHistory := 1.0 + lHopfSnapshot*(1.0-kappaLambdaRedSnapshot)
	cHiggs := cYukawa * cHistory
	if !closeAbs(cYukawa, cYukawaSnapshot, 1e-15) || !closeAbs(cHistory, cHistorySnapshot, 1e-15) || !closeAbs(cHiggs, cHiggsSnapshot, 1e-15) {
		return Analysis{}, fmt.Errorf("frozen C_Higgs ledger mismatch: C_Y=%.17g C_Hist=%.17g C=%.17g", cYukawa, cHistory, cHiggs)
	}

	bCoeff := -5.0/3.0 + xiBoundarySnapshot*pSnapshot
	dCByNEff := -cHiggs / nEffSnapshot
	dCByLHopf := cYukawa * (1.0 - kappaLambdaRedSnapshot)
	dCByKappaLambdaRed := -cYukawa * lHopfSnapshot
	dCByAbsLambda := dCByKappaLambdaRed
	dCByKappaOrient := cYukawa * lHopfSnapshot * (1.0 - pSnapshot*sSnapshot*sSnapshot)
	dKappaLambdaByS := pSnapshot + 2.0*pSnapshot*kappaOrientSnapshot*sSnapshot + 4.0*pSnapshot*bCoeff*sSnapshot*sSnapshot*sSnapshot - 6.0*pSnapshot*pSnapshot*sSnapshot*sSnapshot - 2.0*bCoeff*sSnapshot
	dCByS := -cYukawa * lHopfSnapshot * dKappaLambdaByS
	dCByXi := cYukawa * lHopfSnapshot * pSnapshot * sSnapshot * sSnapshot * (1.0 - pSnapshot*sSnapshot*sSnapshot)

	elasticNEff := dCByNEff * nEffSnapshot / cHiggs
	elasticLHopf := dCByLHopf * lHopfSnapshot / cHiggs
	elasticAbsLambda := dCByAbsLambda * absLambdaSnapshot / cHiggs
	elasticKappaLambda := dCByKappaLambdaRed * kappaLambdaRedSnapshot / cHiggs
	elasticKappaOrient := dCByKappaOrient * kappaOrientSnapshot / cHiggs
	elasticS := dCByS * sSnapshot / cHiggs
	elasticXi := dCByXi * xiBoundarySnapshot / cHiggs

	perturbation := 1e-6
	// Controlled replacements: these are diagnostics, not theorem promotion.
	cTopColor := cHistory
	cNoBoundaryCorrection := cHiggsWith(kappaOrientSnapshot, 0.0, true, true)
	cNoCubic := cHiggsWith(kappaOrientSnapshot, kappaBoundarySnapshot, false, true)

	checks := []struct {
		name string
		got  float64
		want float64
		tol  float64
	}{
		{"dC/dN_eff", dCByNEff, -0.34547216221380384, 1e-15},
		{"dC/dL_Hopf", dCByLHopf, 0.9549361341977547, 1e-15},
		{"dC/dkappa_lambda", dCByKappaLambdaRed, -0.03975789229626174, 1e-15},
		{"dC/dkappa_orient", dCByKappaOrient, 0.039757885839527426, 1e-15},
		{"dC/ds", dCByS, -0.004036181730287719, 1e-15},
		{"dC/dxi", dCByXi, 6.456733266535678e-09, 1e-21},
		{"E_N_eff", elasticNEff, -1.0, 1e-15},
		{"E_L_Hopf", elasticLHopf, 0.03663223082862708, 1e-15},
		{"E_abs_lambda", elasticAbsLambda, -0.0019050960362564661, 1e-18},
		{"E_kappa_lambda", elasticKappaLambda, -0.0016989547911319298, 1e-18},
		{"E_kappa_orient", elasticKappaOrient, 0.00021106412551705377, 1e-18},
		{"E_s", elasticS, -5.029347243414711e-06, 1e-20},
		{"E_xi", elasticXi, 3.1341282342992825e-10, 1e-24},
		{"top-color shift", cTopColor - cHiggs, 0.0008046575187645733, 1e-16},
		{"boundary correction removal shift", cNoBoundaryCorrection - cHiggs, 1.1036177793855018e-7, 1e-18},
		{"cubic removal shift", cNoCubic - cHiggs, -1.6224799281872038e-12, 1e-24},
	}
	for _, c := range checks {
		if !closeAbs(c.got, c.want, c.tol) {
			return Analysis{}, fmt.Errorf("%s mismatch: got %.17g want %.17g", c.name, c.got, c.want)
		}
	}

	return Analysis{
		Gate791: Gate791Inheritance{
			Inherited:       true,
			CleanTestObject: true,
			LevelBOnly:      true,
			Verdict:         StatusGate791Inherited,
		},
		Sensitivity: AnalyticSensitivity{
			Computed:              true,
			DCByNEff:              dCByNEff,
			DCByLHopf:             dCByLHopf,
			DCByKappaLambdaRed:    dCByKappaLambdaRed,
			DCByAbsLambda:         dCByAbsLambda,
			DCByKappaOrient:       dCByKappaOrient,
			DCByS:                 dCByS,
			DCByXiBoundary:        dCByXi,
			FormulaNEff:           "partial C_Higgs / partial N_eff = -C_Higgs/N_eff",
			FormulaLHopf:          "partial C_Higgs / partial L_Hopf = (3/N_eff)(1-kappa_lambda_red)",
			FormulaKappaLambdaRed: "partial C_Higgs / partial kappa_lambda_red = -(3/N_eff)L_Hopf",
			Verdict:               StatusAnalyticSensitivityComputed,
		},
		Elasticity: RelativeElasticity{
			Completed:          true,
			ENEff:              elasticNEff,
			ELHopf:             elasticLHopf,
			EAbsLambda:         elasticAbsLambda,
			EKappaLambdaRed:    elasticKappaLambda,
			EKappaOrient:       elasticKappaOrient,
			ES:                 elasticS,
			EXiBoundary:        elasticXi,
			NEffUnitLeverage:   closeAbs(elasticNEff, -1.0, 1e-15),
			LHopfSecondChannel: math.Abs(elasticLHopf) > math.Abs(elasticAbsLambda),
			Verdict:            StatusRelativeElasticityCompleted,
		},
		Perturbations: AbsolutePerturbationLedger{
			Computed:             true,
			Perturbation:         perturbation,
			DeltaFromNEff:        dCByNEff * perturbation,
			DeltaFromLHopf:       dCByLHopf * perturbation,
			DeltaFromKappaOrient: dCByKappaOrient * perturbation,
			DeltaFromAbsLambda:   dCByAbsLambda * perturbation,
			DeltaFromS:           dCByS * perturbation,
			DeltaFromXiBoundary:  dCByXi * perturbation,
			Strongest: []string{
				"L_Hopf absolute perturbation channel",
				"N_eff absolute perturbation channel",
			},
			Verdict: StatusAbsolutePerturbationComputed,
		},
		Removal: ComponentRemovalDiagnostics{
			Computed:                          true,
			CHiggsTopColor:                    cTopColor,
			TopColorShift:                     cTopColor - cHiggs,
			CHiggsBoundaryCorrectionRemoved:   cNoBoundaryCorrection,
			BoundaryCorrectionShift:           cNoBoundaryCorrection - cHiggs,
			CHiggsCubicRemoved:                cNoCubic,
			CubicRemovalShift:                 cNoCubic - cHiggs,
			YukawaImportant:                   (cTopColor - cHiggs) > 1e-4,
			BoundaryMicroLowNumericalLeverage: math.Abs(cNoBoundaryCorrection-cHiggs) < 1e-6 && math.Abs(cNoCubic-cHiggs) < 1e-9,
			Verdict:                           StatusComponentRemovalDiagnosticsComputed,
		},
		Pressure: SensitivityVsTheoremPressure{
			Separated: true,
			NumericalRanking: []string{
				"1. N_eff",
				"2. L_Hopf",
				"3. |lambda| / kappa_lambda_red / kappa_orient",
				"4. s",
				"5. xi_boundary and fine boundary correction terms",
			},
			TheoremPressureRanking: []string{
				"1. GenerationMixingOperatorSeal: needed for kappa_orient",
				"2. Yukawa operator/eigenvector theorem: needed for N_eff and eventually flavor mixing",
				"3. RadialHessianHopfTransportSeal: needed for native L_Hopf",
				"4. BoundaryExteriorResponsePackageSeal: needed for native F_wall_3_red",
				"5. Boundary scalar wall source theorem: needed for |lambda|, s, xi_boundary",
				"6. Electroweak scale / pole package: needed for physical mass comparison",
			},
			NEffTopNumerical:         true,
			KappaOrientTopFlavor:     true,
			LHopfTopHistory:          true,
			FWallLowLeverageButTyped: true,
			Verdict:                  StatusNumericalVsTheoremSeparated,
		},
		ErrorBudget: ErrorBudgetCategories{
			Defined: true,
			Categories: []string{
				"Type I — numerical uncertainty",
				"Type II — convention uncertainty",
				"Type III — theorem uncertainty",
				"Type IV — comparison uncertainty",
			},
			Classifications: []string{
				"N_eff: Type I + Type III",
				"kappa_orient: Type I + Type III",
				"L_Hopf: Type II + Type III",
				"F_wall_3_red: Type II + Type III",
				"|lambda|, s, xi_boundary: Type I + Type III bridge-coordinate uncertainty",
				"G_F / v: Type I + external scale seal",
				"tree-to-pole correction: Type IV",
			},
			NEffClassified:        true,
			KappaOrientClassified: true,
			LHopfClassified:       true,
			FWallClassified:       true,
			ComparisonClassified:  true,
			Verdict:               StatusErrorBudgetCategoriesDefined,
		},
		Testability: ScientificTestability{
			Completed:                true,
			NumericalSharpnessTarget: "N_eff",
			NativeClosureTargets: []string{
				"GenerationMixingOperatorSeal",
				"RadialHessianHopfTransportSeal",
				"Yukawa operator/eigenvector theorem",
			},
			EmpiricalComparisonTarget: "tree-to-pole correction package",
			NEffBestNumerical:         true,
			GenerationHistoryNative:   true,
			TreeToPoleRequired:        true,
			Verdict:                   StatusScientificTestabilityCompleted,
		},
		Next: NextBranchRecommendation{
			Recorded:    true,
			Recommended: "Gate 793 — N_eff Yukawa Trace Participation Source and Scale-Stability Audit",
			Reason:      "N_eff has unit relative elasticity in C_Higgs, controls the Yukawa dilution factor, and remains a sealed aggregate trace-participation ledger.",
			Alternatives: []string{
				"GenerationMixingOperator Construction Candidate Audit",
				"RadialHessian Hopf Transport Native-Law Audit",
				"Tree-to-Pole Correction Package Audit",
				"BoundaryExteriorResponsePackage Native-Law Audit",
			},
			Verdict: StatusNextBranchRecommendationRecorded,
		},
		Firewalls: Firewalls{
			Enforced:                         true,
			ErrorBudgetNativeTheorem:         false,
			LargestSensitivityDeepestTheorem: false,
			KappaOrientSmallMeansUnimportant: false,
			FWallLowLeverageDisposable:       false,
			CHiggsPoleMassPrediction:         false,
			TreeProxyPoleMass:                false,
			LevelBLevelCPrediction:           false,
			NEffNativeYukawa:                 false,
			LHopfNativeHistoryLoop:           false,
			KappaOrientNativeFlavor:          false,
			ObservedHiggsSourcesBudget:       false,
			Verdict:                          StatusFirewallPreservedGate792,
		},
		Truth:          "N_eff dominates relative numerical leverage on C_Higgs, but theorem risk remains distributed across flavor orientation, Yukawa participation, HistoryLoop transport, and boundary-response seals.",
		FinalStatement: "Gate 792 finds that N_eff is the highest numerical leverage input because C_Higgs has unit relative sensitivity to it. The largest theorem risks remain kappa_orient, L_Hopf, and N_eff: flavor orientation, HistoryLoop transport, and Yukawa trace participation. The recommended next gate is Gate 793 — N_eff Yukawa Trace Participation Source and Scale-Stability Audit, because reducing N_eff most directly improves the numerical scientific testability of the Level-B C_Higgs interface.",
	}, nil
}

func cHiggsWith(kappaOrient, kappaBoundary float64, includeCubic bool, includeBoundaryCorrection bool) float64 {
	ke := kappaOrient
	if includeBoundaryCorrection {
		ke += kappaBoundary
	}
	fWall := pSnapshot*sSnapshot + ke*pSnapshot*sSnapshot*sSnapshot
	if includeCubic {
		fWall -= 2.0 * pSnapshot * pSnapshot * sSnapshot * sSnapshot * sSnapshot
	}
	kappaLambda := absLambdaSnapshot + fWall - ke
	return (3.0 / nEffSnapshot) * (1.0 + lHopfSnapshot*(1.0-kappaLambda))
}

func Statuses() []string {
	return []string{
		StatusGate791Inherited,
		StatusAnalyticSensitivityComputed,
		StatusRelativeElasticityCompleted,
		StatusAbsolutePerturbationComputed,
		StatusComponentRemovalDiagnosticsComputed,
		StatusNumericalVsTheoremSeparated,
		StatusErrorBudgetCategoriesDefined,
		StatusMajorInputsClassified,
		StatusScientificTestabilityCompleted,
		StatusNextBranchRecommendationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusExplicitLinearizedErrorBudget,
		StatusNEffUnitRelativeLeverage,
		StatusLHopfSecondSensitivity,
		StatusAbsoluteResponseStrongestToLAndN,
		StatusYukawaParticipationNumericallyImportant,
		StatusBoundaryMicroLowNumericalLeverage,
		StatusNEffTopNumericalLeverage,
		StatusKappaOrientTopFlavorObstruction,
		StatusLHopfTopHistoryObstruction,
		StatusNEffReductionImprovesTestability,
		StatusGenerationAndHistoryImproveNative,
		StatusTreeToPoleRequired,
		StatusNEffNextNumericalBranch,
		StatusErrorBudgetNotNativeTheorem,
		StatusNEffNotNativeYukawa,
		StatusLHopfNotNativeHistoryLoop,
		StatusKappaOrientNotNativeFlavor,
		StatusFWallNotNativeBoundary,
		StatusLevelBSensitivityNotLevelC,
		StatusTreeProxyNotPoleMass,
		StatusPoleComparisonRequiresCorrection,
		StatusObservedHiggsForbiddenInBudget,
		StatusFirewallPreservedGate792,
	}
}

func FormatSensitivity(s AnalyticSensitivity) string {
	return fmt.Sprintf("dC/dN_eff=%.17g; dC/dL_Hopf=%.17g; dC/dkappa_lambda_red=%.17g; dC/d|lambda|=%.17g; dC/dkappa_orient=%.17g; dC/ds=%.17g; dC/dxi_boundary=%.17g", s.DCByNEff, s.DCByLHopf, s.DCByKappaLambdaRed, s.DCByAbsLambda, s.DCByKappaOrient, s.DCByS, s.DCByXiBoundary)
}

func FormatElasticity(e RelativeElasticity) string {
	return fmt.Sprintf("E_N_eff=%.17g; E_L_Hopf=%.17g; E_|lambda|=%.17g; E_kappa_lambda_red=%.17g; E_kappa_orient=%.17g; E_s=%.17g; E_xi_boundary=%.17g", e.ENEff, e.ELHopf, e.EAbsLambda, e.EKappaLambdaRed, e.EKappaOrient, e.ES, e.EXiBoundary)
}

func FormatPerturbations(p AbsolutePerturbationLedger) string {
	return fmt.Sprintf("for delta=%.1e: N_eff %.17g; L_Hopf %.17g; kappa_orient %.17g; |lambda| %.17g; s %.17g; xi_boundary %.17g", p.Perturbation, p.DeltaFromNEff, p.DeltaFromLHopf, p.DeltaFromKappaOrient, p.DeltaFromAbsLambda, p.DeltaFromS, p.DeltaFromXiBoundary)
}

func FormatRemoval(r ComponentRemovalDiagnostics) string {
	return fmt.Sprintf("top-color C=%.17g shift=%.17g; remove kappa_boundary C=%.17g shift=%.17g; remove cubic C=%.17g shift=%.17g", r.CHiggsTopColor, r.TopColorShift, r.CHiggsBoundaryCorrectionRemoved, r.BoundaryCorrectionShift, r.CHiggsCubicRemoved, r.CubicRemovalShift)
}

func FormatPressure(p SensitivityVsTheoremPressure) string {
	return "numerical: " + strings.Join(p.NumericalRanking, "; ") + " | theorem: " + strings.Join(p.TheoremPressureRanking, "; ")
}

func containsAll(haystack []string, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if strings.Contains(h, n) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

var buildOnce struct {
	sync.Once
	a   Analysis
	err error
}

func Cached() (Analysis, error) {
	buildOnce.Do(func() { buildOnce.a, buildOnce.err = BuildDefault() })
	return buildOnce.a, buildOnce.err
}
