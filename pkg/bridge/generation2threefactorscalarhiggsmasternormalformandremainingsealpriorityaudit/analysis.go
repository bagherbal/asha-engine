// Package generation2threefactorscalarhiggsmasternormalformandremainingsealpriorityaudit implements
// Gate 760: Three-Factor Scalar-Higgs Master Normal Form and Remaining-Seal Priority Audit.
//
// Gate 759 rewrote the HistoryLoop transport factor as
// C_History=1+L_Hopf(1-kappa_lambda_red). Gate 760 records the resulting
// scalar-Higgs bridge as the current three-factor master normal form
// lambda_runtime_eff=(1/8)(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)] and orders
// the remaining unreduced seal targets. This is a master-normal-form and
// seal-priority audit only. It does not derive Yukawa eigenvalues, scalar
// runtime lambda, Higgs mass, pole mass, CKM/PMNS, flavor hierarchy, or a
// native HistoryLoopUnit theorem.
package generation2threefactorscalarhiggsmasternormalformandremainingsealpriorityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE760-THREE-FACTOR-SCALAR-HIGGS-MASTER-NORMAL-FORM-AND-REMAINING-SEAL-PRIORITY-AUDIT"

	StatusGate759HistoryTransportBracketInherited            = "PASS_GATE759_HISTORY_TRANSPORT_BRACKET_INHERITED"
	StatusThreeFactorMasterFormDefined                       = "PASS_THREE_FACTOR_MASTER_FORM_DEFINED"
	StatusMasterNumericalLedgerRecorded                      = "PASS_MASTER_NUMERICAL_LEDGER_RECORDED"
	StatusFactorSourceTypesAudited                           = "PASS_FACTOR_SOURCE_TYPES_AUDITED"
	StatusKappaLambdaRedExpansionRecorded                    = "PASS_KAPPA_LAMBDA_RED_EXPANSION_RECORDED"
	StatusRemainingSealPriorityAudited                       = "PASS_REMAINING_SEAL_PRIORITY_AUDITED"
	StatusNextReductionTargetsOrdered                        = "PASS_NEXT_REDUCTION_TARGETS_ORDERED"
	StatusPhysicalFirewallsEnforced                          = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusScalarHiggsBridgeHasThreeFactorMasterNormalForm    = "CONDITIONAL_SUPPORT_SCALAR_HIGGS_BRIDGE_HAS_THREE_FACTOR_MASTER_NORMAL_FORM"
	StatusKappaLambdaRedIsReconstructedScalarMatchingDeficit = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_RED_IS_RECONSTRUCTED_SCALAR_MATCHING_DEFICIT"
	StatusNextScalarSourceReductionTargetIsPradOrLHopf       = "CONDITIONAL_SUPPORT_NEXT_SCALAR_SOURCE_REDUCTION_TARGET_IS_P_RAD_OR_L_HOPF"
	StatusThreeFactorFormNotIndependentRuntimeTheorem        = "FAILED_ROUTE_THREE_FACTOR_FORM_NOT_INDEPENDENT_RUNTIME_THEOREM"
	StatusNEffNotNativeYukawaTheorem                         = "FAILED_ROUTE_N_EFF_NOT_NATIVE_YUKAWA_THEOREM"
	StatusLHopfNotNativeHistoryLoopTheorem                   = "FAILED_ROUTE_L_HOPF_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusKappaLambdaRedNotNativeScalarTheorem               = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_NATIVE_SCALAR_THEOREM"
	StatusNoNativePradSelector                               = "FAILED_ROUTE_NO_NATIVE_P_RAD_SELECTOR"
	StatusNoNativeBoundaryResponseGeneratingFunctionTheorem  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                       = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate760ThreeFactorMasterFormBoundary               = "FIREWALL_PRESERVED_GATE760_THREE_FACTOR_MASTER_FORM_BOUNDARY"
)

const (
	oneEighth = 1.0 / 8.0
	three     = 3.0
	eight     = 8.0

	// Gate 756/Gate 758/Gate 759 audited scalar-coordinate snapshots.
	nEffMZ             = 3.0023273474722147
	cYukawaMZ          = 0.9992248188812008
	kappaLambdaRedMZ   = 0.04432304306956136
	lambdaRuntimeEffMZ = 0.12965256505060754

	// Gate 624/Gate 758/Gate 759 Radial-Hopf loop-unit snapshot.
	lHopf = 1.0 / (eight * math.Pi)
)

type Gate759Inheritance struct {
	Inherited                       bool
	HistoryFormula                  string
	FullFormula                     string
	KappaLambdaRed                  float64
	Complement                      float64
	CHistory                        float64
	LambdaRuntimeEff                float64
	ThreeFactorFormAvailable        bool
	IndependentScalarRuntimeTheorem bool
	Verdict                         string
}

type MasterFormula struct {
	BaselineSymbol                  string
	CYukawaSymbol                   string
	CHistorySymbol                  string
	Formula                         string
	ExpandedFormula                 string
	CBaseline                       float64
	CYukawa                         float64
	CHistory                        float64
	TotalCorrection                 float64
	LambdaRuntimeEff                float64
	LambdaRuntimeFromMaster         float64
	MasterResidual                  float64
	Defined                         bool
	IndependentScalarRuntimeTheorem bool
	Verdict                         string
}

type MasterNumericalLedger struct {
	NEff                  float64
	CYukawa               float64
	LHopf                 float64
	KappaLambdaRed        float64
	KappaLambdaComplement float64
	CHistory              float64
	LambdaRuntimeEff      float64
	Recorded              bool
	Finite                bool
	Verdict               string
}

type FactorSourceTypes struct {
	BaselineSourceType       string
	CYukawaSourceType        string
	LHopfSourceType          string
	KappaLambdaRedSourceType string
	Audited                  bool
	BaselineScalarTheorem    bool
	NEffNativeYukawaTheorem  bool
	LHopfNativeTheorem       bool
	KappaNativeTheorem       bool
	Verdict                  string
}

type KappaLambdaRedExpansion struct {
	Definition                  string
	FWall3RedFormula            string
	KappaERedFormula            string
	PrimitiveInCurrentBridge    bool
	ReconstructedFromWallFlavor bool
	ExpansionRecorded           bool
	NativeScalarTheorem         bool
	BoundaryGeneratingFunction  bool
	Verdict                     string
}

type SealPriority struct {
	Rank                        int
	Symbol                      string
	Layer                       string
	Reason                      string
	Unreduced                   bool
	ActiveInRuntimeScalarNumber bool
}

type RemainingSealPriorityAudit struct {
	Priorities                  []SealPriority
	Audited                     bool
	Ordered                     bool
	ScalarReductionTarget       string
	FlavorYukawaReductionTarget string
	BoundaryReductionTarget     string
	NativePradSelector          bool
	NativeBoundaryGenerator     bool
	Verdict                     string
}

type PhysicalFirewalls struct {
	ThreeFactorIndependentRuntimeTheorem bool
	NEffNativeYukawaTheorem              bool
	LHopfNativeHistoryLoopTheorem        bool
	KappaLambdaRedNativeScalarTheorem    bool
	TreeProxyPoleMass                    bool
	HiggsSocketSealsHiggsMassTheorem     bool
	YukawaOperatorOrEigenvalueTheorem    bool
	HiggsMassOrPoleMassTheorem           bool
	Audited                              bool
	Verdict                              string
}

type Analysis struct {
	Gate759        Gate759Inheritance
	Master         MasterFormula
	Ledger         MasterNumericalLedger
	Sources        FactorSourceTypes
	KappaExpansion KappaLambdaRedExpansion
	SealPriority   RemainingSealPriorityAudit
	Firewalls      PhysicalFirewalls
	Truth          string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g759 := buildGate759Inheritance()
	if !finitePositive(g759.CHistory) || !finitePositive(g759.Complement) || !finitePositive(g759.LambdaRuntimeEff) {
		return Analysis{}, fmt.Errorf("invalid Gate759 inherited values: C_History=%g complement=%g lambda=%g", g759.CHistory, g759.Complement, g759.LambdaRuntimeEff)
	}
	master := buildMasterFormula(g759)
	if !finitePositive(master.CYukawa) || !finitePositive(master.CHistory) || !finitePositive(master.LambdaRuntimeFromMaster) {
		return Analysis{}, fmt.Errorf("invalid Gate760 master factors: C_Yukawa=%g C_History=%g lambda=%g", master.CYukawa, master.CHistory, master.LambdaRuntimeFromMaster)
	}
	ledger := buildMasterNumericalLedger(master)
	sources := buildFactorSourceTypes()
	expansion := buildKappaLambdaRedExpansion()
	priority := buildRemainingSealPriorityAudit()
	firewalls := buildPhysicalFirewalls()
	truth := "Gate 760 records the current scalar-Higgs bridge as lambda_runtime_eff=(1/8)(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]. Numerically N_eff=3.0023273474722147, C_Yukawa=0.9992248188812008, L_Hopf=1/(8*pi), kappa_lambda_red=0.04432304306956136, C_History=1.038025177923625, and lambda_runtime_eff=0.12965256505060754. The master formula is a three-factor scalar-coordinate normal form: one-eighth top-color scalar proxy baseline, finite Yukawa trace participation correction, and Radial-Hopf transport of a reduced scalar matching complement. The remaining seal priority for scalar source reduction is P_rad/L_Hopf first, then N_eff for Yukawa/flavor reduction, then F_wall_3_red for boundary-response reduction. This is not an independent scalar-runtime theorem, native Yukawa theorem, native HistoryLoop theorem, native scalar matching theorem, Higgs-mass theorem, or pole-mass theorem."
	return Analysis{Gate759: g759, Master: master, Ledger: ledger, Sources: sources, KappaExpansion: expansion, SealPriority: priority, Firewalls: firewalls, Truth: truth}, nil
}

func buildGate759Inheritance() Gate759Inheritance {
	cHistory := 1.0 + lHopf*(1.0-kappaLambdaRedMZ)
	return Gate759Inheritance{
		Inherited:                       true,
		HistoryFormula:                  "C_History=1+L_Hopf(1-kappa_lambda_red)",
		FullFormula:                     "lambda_runtime_eff=(1/8) C_Yukawa [1+L_Hopf(1-kappa_lambda_red)]",
		KappaLambdaRed:                  kappaLambdaRedMZ,
		Complement:                      1.0 - kappaLambdaRedMZ,
		CHistory:                        cHistory,
		LambdaRuntimeEff:                lambdaRuntimeEffMZ,
		ThreeFactorFormAvailable:        true,
		IndependentScalarRuntimeTheorem: false,
		Verdict: strings.Join([]string{
			StatusGate759HistoryTransportBracketInherited,
			StatusScalarHiggsBridgeHasThreeFactorMasterNormalForm,
			StatusThreeFactorFormNotIndependentRuntimeTheorem,
		}, "; "),
	}
}

func buildMasterFormula(g Gate759Inheritance) MasterFormula {
	cYukawa := three / nEffMZ
	cHistory := 1.0 + lHopf*(1.0-g.KappaLambdaRed)
	total := cYukawa * cHistory
	lambda := oneEighth * total
	return MasterFormula{
		BaselineSymbol:                  "C_baseline=1/8",
		CYukawaSymbol:                   "C_Yukawa=3/N_eff",
		CHistorySymbol:                  "C_History=1+L_Hopf(1-kappa_lambda_red)",
		Formula:                         "lambda_runtime_eff=C_baseline C_Yukawa C_History",
		ExpandedFormula:                 "lambda_runtime_eff=(1/8)(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]",
		CBaseline:                       oneEighth,
		CYukawa:                         cYukawa,
		CHistory:                        cHistory,
		TotalCorrection:                 total,
		LambdaRuntimeEff:                g.LambdaRuntimeEff,
		LambdaRuntimeFromMaster:         lambda,
		MasterResidual:                  lambda - g.LambdaRuntimeEff,
		Defined:                         true,
		IndependentScalarRuntimeTheorem: false,
		Verdict: strings.Join([]string{
			StatusThreeFactorMasterFormDefined,
			StatusScalarHiggsBridgeHasThreeFactorMasterNormalForm,
			StatusThreeFactorFormNotIndependentRuntimeTheorem,
		}, "; "),
	}
}

func buildMasterNumericalLedger(m MasterFormula) MasterNumericalLedger {
	return MasterNumericalLedger{
		NEff:                  nEffMZ,
		CYukawa:               m.CYukawa,
		LHopf:                 lHopf,
		KappaLambdaRed:        kappaLambdaRedMZ,
		KappaLambdaComplement: 1.0 - kappaLambdaRedMZ,
		CHistory:              m.CHistory,
		LambdaRuntimeEff:      m.LambdaRuntimeFromMaster,
		Recorded:              true,
		Finite:                finitePositive(nEffMZ) && finitePositive(m.CYukawa) && finitePositive(lHopf) && finitePositive(m.CHistory) && finitePositive(m.LambdaRuntimeFromMaster),
		Verdict: strings.Join([]string{
			StatusMasterNumericalLedgerRecorded,
			StatusScalarHiggsBridgeHasThreeFactorMasterNormalForm,
		}, "; "),
	}
}

func buildFactorSourceTypes() FactorSourceTypes {
	return FactorSourceTypes{
		BaselineSourceType:       "top-color scalar proxy baseline: (3/8) gauge/spectral coefficient times 1/3 top-color participation shadow",
		CYukawaSourceType:        "finite Yukawa trace participation correction: b/a^2=1/N_eff, C_Yukawa=3/N_eff",
		LHopfSourceType:          "Radial-Hopf loop unit candidate: Tr_K7+(rho_plus (1/(2*pi)) P_rad)",
		KappaLambdaRedSourceType: "reduced scalar matching deficit: |lambda|+F_wall_3_red-kappa_e_red",
		Audited:                  true,
		BaselineScalarTheorem:    false,
		NEffNativeYukawaTheorem:  false,
		LHopfNativeTheorem:       false,
		KappaNativeTheorem:       false,
		Verdict: strings.Join([]string{
			StatusFactorSourceTypesAudited,
			StatusKappaLambdaRedIsReconstructedScalarMatchingDeficit,
			StatusNEffNotNativeYukawaTheorem,
			StatusLHopfNotNativeHistoryLoopTheorem,
			StatusKappaLambdaRedNotNativeScalarTheorem,
		}, "; "),
	}
}

func buildKappaLambdaRedExpansion() KappaLambdaRedExpansion {
	return KappaLambdaRedExpansion{
		Definition:                  "kappa_lambda_red=|lambda|+F_wall_3_red(s)-kappa_e_red",
		FWall3RedFormula:            "F_wall_3_red(s)=p_K7 s+kappa_e_red p_K7 s^2-2p_K7^2 s^3",
		KappaERedFormula:            "kappa_e_red=sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p_K7 s^2",
		PrimitiveInCurrentBridge:    false,
		ReconstructedFromWallFlavor: true,
		ExpansionRecorded:           true,
		NativeScalarTheorem:         false,
		BoundaryGeneratingFunction:  false,
		Verdict: strings.Join([]string{
			StatusKappaLambdaRedExpansionRecorded,
			StatusKappaLambdaRedIsReconstructedScalarMatchingDeficit,
			StatusKappaLambdaRedNotNativeScalarTheorem,
			StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		}, "; "),
	}
}

func buildRemainingSealPriorityAudit() RemainingSealPriorityAudit {
	priorities := []SealPriority{
		{Rank: 1, Symbol: "P_rad", Layer: "ScalarVacuumDirectionSeal / Radial-Hopf source", Reason: "needed to source L_Hopf as Tr(rho_plus[(1/(2*pi))P_rad]); strongest scalar-runtime source-reduction pressure point", Unreduced: true, ActiveInRuntimeScalarNumber: true},
		{Rank: 2, Symbol: "n", Layer: "TwistorSelectorSeal / Higgs socket direction", Reason: "needed to define J_H(n), Hopf phase direction, and the sealed Higgs socket", Unreduced: true, ActiveInRuntimeScalarNumber: true},
		{Rank: 3, Symbol: "N_eff", Layer: "finite Yukawa trace participation", Reason: "reduced from b/a^2 but still depends on sealed Yukawa singular-value ledger", Unreduced: true, ActiveInRuntimeScalarNumber: true},
		{Rank: 4, Symbol: "kappa_e_red", Layer: "reduced flavor-wall deficit", Reason: "strongly source-typed but still depends on empirical/bridge theta13 and J_CKM", Unreduced: true, ActiveInRuntimeScalarNumber: true},
		{Rank: 5, Symbol: "F_wall_3_red", Layer: "cubic boundary-history response", Reason: "strong closure, but no native raw-moment generating-function theorem", Unreduced: true, ActiveInRuntimeScalarNumber: true},
		{Rank: 6, Symbol: "q", Layer: "hypercharge normalization / Higgs socket interface", Reason: "important for representation interface, but not directly active in the scalar runtime number after trace collapse", Unreduced: true, ActiveInRuntimeScalarNumber: false},
	}
	return RemainingSealPriorityAudit{
		Priorities:                  priorities,
		Audited:                     true,
		Ordered:                     prioritiesAreStrict(priorities),
		ScalarReductionTarget:       "P_rad / L_Hopf",
		FlavorYukawaReductionTarget: "N_eff",
		BoundaryReductionTarget:     "F_wall_3_red",
		NativePradSelector:          false,
		NativeBoundaryGenerator:     false,
		Verdict: strings.Join([]string{
			StatusRemainingSealPriorityAudited,
			StatusNextReductionTargetsOrdered,
			StatusNextScalarSourceReductionTargetIsPradOrLHopf,
			StatusNoNativePradSelector,
			StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		}, "; "),
	}
}

func buildPhysicalFirewalls() PhysicalFirewalls {
	return PhysicalFirewalls{
		ThreeFactorIndependentRuntimeTheorem: false,
		NEffNativeYukawaTheorem:              false,
		LHopfNativeHistoryLoopTheorem:        false,
		KappaLambdaRedNativeScalarTheorem:    false,
		TreeProxyPoleMass:                    false,
		HiggsSocketSealsHiggsMassTheorem:     false,
		YukawaOperatorOrEigenvalueTheorem:    false,
		HiggsMassOrPoleMassTheorem:           false,
		Audited:                              true,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusThreeFactorFormNotIndependentRuntimeTheorem,
			StatusNEffNotNativeYukawaTheorem,
			StatusLHopfNotNativeHistoryLoopTheorem,
			StatusKappaLambdaRedNotNativeScalarTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate760ThreeFactorMasterFormBoundary,
		}, "; "),
	}
}

func finitePositive(x float64) bool {
	return x > 0 && !math.IsNaN(x) && !math.IsInf(x, 0)
}

func prioritiesAreStrict(xs []SealPriority) bool {
	if len(xs) == 0 {
		return false
	}
	for i, x := range xs {
		if x.Rank != i+1 || x.Symbol == "" || x.Layer == "" || x.Reason == "" {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate759HistoryTransportBracketInherited,
		StatusThreeFactorMasterFormDefined,
		StatusMasterNumericalLedgerRecorded,
		StatusFactorSourceTypesAudited,
		StatusKappaLambdaRedExpansionRecorded,
		StatusRemainingSealPriorityAudited,
		StatusNextReductionTargetsOrdered,
		StatusPhysicalFirewallsEnforced,
		StatusScalarHiggsBridgeHasThreeFactorMasterNormalForm,
		StatusKappaLambdaRedIsReconstructedScalarMatchingDeficit,
		StatusNextScalarSourceReductionTargetIsPradOrLHopf,
		StatusThreeFactorFormNotIndependentRuntimeTheorem,
		StatusNEffNotNativeYukawaTheorem,
		StatusLHopfNotNativeHistoryLoopTheorem,
		StatusKappaLambdaRedNotNativeScalarTheorem,
		StatusNoNativePradSelector,
		StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate760ThreeFactorMasterFormBoundary,
	}
}

func FormatGate759(x Gate759Inheritance) string {
	return fmt.Sprintf("inherited=%t historyFormula=%q fullFormula=%q kappa=%.16g complement=%.16g cHistory=%.16g lambdaRuntime=%.16g threeFactor=%t independentRuntime=%t verdict=%q", x.Inherited, x.HistoryFormula, x.FullFormula, x.KappaLambdaRed, x.Complement, x.CHistory, x.LambdaRuntimeEff, x.ThreeFactorFormAvailable, x.IndependentScalarRuntimeTheorem, x.Verdict)
}

func FormatMaster(x MasterFormula) string {
	return fmt.Sprintf("baseline=%q cY=%q cH=%q formula=%q expanded=%q baselineValue=%.16g cYValue=%.16g cHValue=%.16g total=%.16g lambda=%.16g fromMaster=%.16g residual=%.16g defined=%t independentRuntime=%t verdict=%q", x.BaselineSymbol, x.CYukawaSymbol, x.CHistorySymbol, x.Formula, x.ExpandedFormula, x.CBaseline, x.CYukawa, x.CHistory, x.TotalCorrection, x.LambdaRuntimeEff, x.LambdaRuntimeFromMaster, x.MasterResidual, x.Defined, x.IndependentScalarRuntimeTheorem, x.Verdict)
}

func FormatLedger(x MasterNumericalLedger) string {
	return fmt.Sprintf("nEff=%.16g cY=%.16g L=%.16g kappa=%.16g complement=%.16g cH=%.16g lambda=%.16g recorded=%t finite=%t verdict=%q", x.NEff, x.CYukawa, x.LHopf, x.KappaLambdaRed, x.KappaLambdaComplement, x.CHistory, x.LambdaRuntimeEff, x.Recorded, x.Finite, x.Verdict)
}

func FormatSources(x FactorSourceTypes) string {
	return fmt.Sprintf("baselineSource=%q cYSource=%q LSource=%q kappaSource=%q audited=%t baselineTheorem=%t nEffNative=%t LNative=%t kappaNative=%t verdict=%q", x.BaselineSourceType, x.CYukawaSourceType, x.LHopfSourceType, x.KappaLambdaRedSourceType, x.Audited, x.BaselineScalarTheorem, x.NEffNativeYukawaTheorem, x.LHopfNativeTheorem, x.KappaNativeTheorem, x.Verdict)
}

func FormatKappaExpansion(x KappaLambdaRedExpansion) string {
	return fmt.Sprintf("definition=%q fWall=%q kappaE=%q primitive=%t reconstructed=%t recorded=%t nativeScalar=%t boundaryGenerator=%t verdict=%q", x.Definition, x.FWall3RedFormula, x.KappaERedFormula, x.PrimitiveInCurrentBridge, x.ReconstructedFromWallFlavor, x.ExpansionRecorded, x.NativeScalarTheorem, x.BoundaryGeneratingFunction, x.Verdict)
}

func FormatSealPriority(x RemainingSealPriorityAudit) string {
	items := make([]string, 0, len(x.Priorities))
	for _, p := range x.Priorities {
		items = append(items, fmt.Sprintf("%d:%s[%s] active=%t", p.Rank, p.Symbol, p.Layer, p.ActiveInRuntimeScalarNumber))
	}
	return fmt.Sprintf("priorities=%q audited=%t ordered=%t scalarTarget=%q yukawaTarget=%q boundaryTarget=%q nativePrad=%t nativeBoundaryGenerator=%t verdict=%q", strings.Join(items, " | "), x.Audited, x.Ordered, x.ScalarReductionTarget, x.FlavorYukawaReductionTarget, x.BoundaryReductionTarget, x.NativePradSelector, x.NativeBoundaryGenerator, x.Verdict)
}

func FormatFirewalls(x PhysicalFirewalls) string {
	return fmt.Sprintf("reject(threeFactorRuntime=%t nEffNative=%t LNative=%t kappaNative=%t treePole=%t socketHiggsMass=%t yukawaEigen=%t higgsPole=%t) audited=%t verdict=%q", x.ThreeFactorIndependentRuntimeTheorem, x.NEffNativeYukawaTheorem, x.LHopfNativeHistoryLoopTheorem, x.KappaLambdaRedNativeScalarTheorem, x.TreeProxyPoleMass, x.HiggsSocketSealsHiggsMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.HiggsMassOrPoleMassTheorem, x.Audited, x.Verdict)
}
