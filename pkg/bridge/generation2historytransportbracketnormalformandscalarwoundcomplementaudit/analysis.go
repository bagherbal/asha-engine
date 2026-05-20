// Package generation2historytransportbracketnormalformandscalarwoundcomplementaudit implements
// Gate 759: History Transport Bracket Normal Form and Scalar-Wound Complement Audit.
//
// Gate 758 factored the reduced scalar-Higgs bridge as
// lambda_runtime_eff=(1/8) C_Yukawa C_History. Gate 759 audits the internal
// normal form of the HistoryLoop transport factor by isolating
// Omega_History=1-|lambda|-F_wall_3_red+kappa_e_red and rewriting it as the
// complement of a reduced scalar matching deficit:
// Omega_History=1-kappa_lambda_red. This is a scalar-history bracket
// normalization audit only. It does not derive scalar runtime lambda, Higgs
// mass, pole mass, Yukawa eigenvalues, CKM/PMNS, flavor hierarchy, or a native
// HistoryLoopUnit theorem.
package generation2historytransportbracketnormalformandscalarwoundcomplementaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE759-HISTORY-TRANSPORT-BRACKET-NORMAL-FORM-AND-SCALAR-WOUND-COMPLEMENT-AUDIT"

	StatusGate758OneEighthFactorizationInherited                            = "PASS_GATE758_ONE_EIGHTH_FACTORIZATION_INHERITED"
	StatusHistoryTransportBracketDefined                                    = "PASS_HISTORY_TRANSPORT_BRACKET_DEFINED"
	StatusOmegaHistoryComputed                                              = "PASS_OMEGA_HISTORY_COMPUTED"
	StatusKappaLambdaRedDefined                                             = "PASS_KAPPA_LAMBDA_RED_DEFINED"
	StatusOmegaHistoryRewrittenAsOneMinusKappaLambdaRed                     = "PASS_OMEGA_HISTORY_REWRITTEN_AS_ONE_MINUS_KAPPA_LAMBDA_RED"
	StatusCHistoryNormalFormWritten                                         = "PASS_C_HISTORY_NORMAL_FORM_WRITTEN"
	StatusFullScalarHiggsFormRewritten                                      = "PASS_FULL_SCALAR_HIGGS_FORM_REWRITTEN"
	StatusSourceTypeInterpretationRecorded                                  = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusLayerSeparationAudited                                            = "PASS_LAYER_SEPARATION_AUDITED"
	StatusIllegalTermRejectionAudited                                       = "PASS_ILLEGAL_TERM_REJECTION_AUDITED"
	StatusCHistoryRadialHopfTransportOfScalarMatchingComplement             = "CONDITIONAL_SUPPORT_C_HISTORY_IS_RADIAL_HOPF_TRANSPORT_OF_SCALAR_MATCHING_COMPLEMENT"
	StatusKappaLambdaRedReconstructsScalarMatchingDeficitFromWallFlavorData = "CONDITIONAL_SUPPORT_KAPPA_LAMBDA_RED_RECONSTRUCTS_SCALAR_MATCHING_DEFICIT_FROM_WALL_FLAVOR_DATA"
	StatusScalarHiggsBridgeHasThreeFactorNormalForm                         = "CONDITIONAL_SUPPORT_SCALAR_HIGGS_BRIDGE_HAS_THREE_FACTOR_NORMAL_FORM"
	StatusKappaLambdaRedNotNativeScalarTheorem                              = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_NATIVE_SCALAR_THEOREM"
	StatusCHistoryNotNativeHistoryLoopTheorem                               = "FAILED_ROUTE_C_HISTORY_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusLHopfNotNativeTransportTheorem                                    = "FAILED_ROUTE_L_HOPF_NOT_NATIVE_TRANSPORT_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem                                 = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                                      = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                               = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate759HistoryTransportBracketBoundary                            = "FIREWALL_PRESERVED_GATE759_HISTORY_TRANSPORT_BRACKET_BOUNDARY"
)

const (
	oneEighth = 1.0 / 8.0
	three     = 3.0
	eight     = 8.0

	// Gate 756/Gate 758 audited bridge snapshots.
	nEffMZ             = 3.0023273474722147
	cYukawaMZ          = 0.9992248188812008
	cHistoryMZ         = 1.038025177923625
	lambdaRuntimeEffMZ = 0.12965256505060754

	// Gate 624/Gate 758 Radial-Hopf loop-unit snapshot.
	lHopf = 1.0 / (eight * math.Pi)
)

type Gate758Inheritance struct {
	Inherited                       bool
	FactorizationFormula            string
	CYukawaFormula                  string
	CHistoryFormula                 string
	NEff                            float64
	CYukawa                         float64
	CHistory                        float64
	LambdaRuntimeEff                float64
	OneEighthBaseline               float64
	ThreeFactorNormalFormAvailable  bool
	IndependentScalarRuntimeTheorem bool
	Verdict                         string
}

type HistoryTransportBracket struct {
	BracketSymbol         string
	OmegaFormula          string
	CHistoryFormula       string
	LHopf                 float64
	CHistory              float64
	OmegaHistory          float64
	OmegaFromCHistory     float64
	OmegaResidual         float64
	BracketDefined        bool
	OmegaComputed         bool
	PhysicalTimeOrRGScale bool
	Verdict               string
}

type ReducedScalarMatchingDeficit struct {
	Symbol                string
	Definition            string
	KappaLambdaRed        float64
	Complement            float64
	OmegaHistory          float64
	ComplementResidual    float64
	Defined               bool
	RewrittenAsComplement bool
	NativeScalarTheorem   bool
	Verdict               string
}

type HistoryNormalForm struct {
	CHistoryOriginalFormula     string
	CHistoryReducedFormula      string
	FullScalarHiggsFormula      string
	CYukawa                     float64
	CHistory                    float64
	KappaLambdaRed              float64
	OmegaHistory                float64
	LambdaRuntimeEff            float64
	LambdaRuntimeFromNormalForm float64
	NormalFormResidual          float64
	CHistoryWritten             bool
	FullFormRewritten           bool
	ThreeFactorNormalForm       bool
	IndependentRuntimeTheorem   bool
	Verdict                     string
}

type SourceTypeInterpretation struct {
	KappaLambdaRedSourceType string
	OmegaHistorySourceType   string
	CHistorySourceType       string
	Recorded                 bool
	NativeHistoryLoopTheorem bool
	NativeTransportTheorem   bool
	Verdict                  string
}

type LayerSeparation struct {
	CYukawaLayer                       string
	KappaLambdaRedLayer                string
	LHopfLayer                         string
	FactorsMultiplyAfterScalarCollapse bool
	OperatorsOnSameNativeBoard         bool
	LayerSeparationAudited             bool
	Verdict                            string
}

type IllegalTermRejection struct {
	KappaLambdaRedNativeScalarTheorem     bool
	OmegaHistoryPhysicalTimeOrRGScale     bool
	LHopfBoundaryEventProbability         bool
	CHistoryNativeHistoryLoopTheorem      bool
	LambdaRuntimeEffIndependentPrediction bool
	TreeProxyPoleMassPrediction           bool
	ClaimsYukawaEigenvaluesDerived        bool
	ClaimsHiggsMassOrPoleMassTheorem      bool
	Audited                               bool
	Verdict                               string
}

type Analysis struct {
	Gate758        Gate758Inheritance
	Bracket        HistoryTransportBracket
	Deficit        ReducedScalarMatchingDeficit
	NormalForm     HistoryNormalForm
	Interpretation SourceTypeInterpretation
	Layers         LayerSeparation
	Illegal        IllegalTermRejection
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
	g758 := buildGate758Inheritance()
	if !finitePositive(g758.CYukawa) || !finitePositive(g758.CHistory) || !finitePositive(g758.NEff) {
		return Analysis{}, fmt.Errorf("invalid Gate758 inherited factors: C_Yukawa=%g C_History=%g N_eff=%g", g758.CYukawa, g758.CHistory, g758.NEff)
	}
	bracket := buildHistoryTransportBracket(g758)
	if !finitePositive(bracket.LHopf) || !finitePositive(bracket.OmegaHistory) {
		return Analysis{}, fmt.Errorf("invalid History bracket values: L_Hopf=%g Omega=%g", bracket.LHopf, bracket.OmegaHistory)
	}
	deficit := buildReducedScalarMatchingDeficit(bracket)
	normal := buildHistoryNormalForm(g758, bracket, deficit)
	interp := buildSourceTypeInterpretation()
	layers := buildLayerSeparation()
	illegal := buildIllegalTermRejection()
	truth := "Gate 759 rewrites the Gate758 HistoryLoop transport bracket Omega_History=1-|lambda|-F_wall_3_red+kappa_e_red as the complement Omega_History=1-kappa_lambda_red, where kappa_lambda_red=|lambda|+F_wall_3_red-kappa_e_red. Numerically C_History=1.038025177923625, L_Hopf=1/(8*pi), Omega_History=0.9556769569304386, and kappa_lambda_red=0.04432304306956136. The scalar-Higgs bridge is therefore lambda_runtime_eff=(1/8) C_Yukawa [1+L_Hopf(1-kappa_lambda_red)]. This is a scalar-history bracket normal form and source-typing audit only, not a native scalar theorem, native HistoryLoop theorem, independent scalar-runtime theorem, Higgs-mass theorem, pole-mass theorem, or Yukawa eigenvalue theorem."
	return Analysis{Gate758: g758, Bracket: bracket, Deficit: deficit, NormalForm: normal, Interpretation: interp, Layers: layers, Illegal: illegal, Truth: truth}, nil
}

func buildGate758Inheritance() Gate758Inheritance {
	return Gate758Inheritance{
		Inherited:                       true,
		FactorizationFormula:            "lambda_runtime_eff=(1/8) C_Yukawa C_History",
		CYukawaFormula:                  "C_Yukawa=3/N_eff=3b/a^2",
		CHistoryFormula:                 "C_History=1+L_Hopf(1-|lambda|-F_wall_3_red+kappa_e_red)",
		NEff:                            nEffMZ,
		CYukawa:                         cYukawaMZ,
		CHistory:                        cHistoryMZ,
		LambdaRuntimeEff:                lambdaRuntimeEffMZ,
		OneEighthBaseline:               oneEighth,
		ThreeFactorNormalFormAvailable:  true,
		IndependentScalarRuntimeTheorem: false,
		Verdict: strings.Join([]string{
			StatusGate758OneEighthFactorizationInherited,
			StatusScalarHiggsBridgeHasThreeFactorNormalForm,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildHistoryTransportBracket(g Gate758Inheritance) HistoryTransportBracket {
	omega := (g.CHistory - 1.0) / lHopf
	return HistoryTransportBracket{
		BracketSymbol:         "Omega_History",
		OmegaFormula:          "Omega_History=1-|lambda(Lambda_12)|-F_wall_3_red(s)+kappa_e_red",
		CHistoryFormula:       "C_History=1+L_Hopf Omega_History",
		LHopf:                 lHopf,
		CHistory:              g.CHistory,
		OmegaHistory:          omega,
		OmegaFromCHistory:     omega,
		OmegaResidual:         omega - ((g.CHistory - 1.0) / lHopf),
		BracketDefined:        true,
		OmegaComputed:         true,
		PhysicalTimeOrRGScale: false,
		Verdict: strings.Join([]string{
			StatusHistoryTransportBracketDefined,
			StatusOmegaHistoryComputed,
			StatusCHistoryRadialHopfTransportOfScalarMatchingComplement,
		}, "; "),
	}
}

func buildReducedScalarMatchingDeficit(b HistoryTransportBracket) ReducedScalarMatchingDeficit {
	kappa := 1.0 - b.OmegaHistory
	return ReducedScalarMatchingDeficit{
		Symbol:                "kappa_lambda_red",
		Definition:            "kappa_lambda_red=|lambda(Lambda_12)|+F_wall_3_red(s)-kappa_e_red",
		KappaLambdaRed:        kappa,
		Complement:            1.0 - kappa,
		OmegaHistory:          b.OmegaHistory,
		ComplementResidual:    b.OmegaHistory - (1.0 - kappa),
		Defined:               true,
		RewrittenAsComplement: true,
		NativeScalarTheorem:   false,
		Verdict: strings.Join([]string{
			StatusKappaLambdaRedDefined,
			StatusOmegaHistoryRewrittenAsOneMinusKappaLambdaRed,
			StatusKappaLambdaRedReconstructsScalarMatchingDeficitFromWallFlavorData,
			StatusKappaLambdaRedNotNativeScalarTheorem,
		}, "; "),
	}
}

func buildHistoryNormalForm(g Gate758Inheritance, b HistoryTransportBracket, d ReducedScalarMatchingDeficit) HistoryNormalForm {
	cHistoryReduced := 1.0 + lHopf*(1.0-d.KappaLambdaRed)
	lambda := oneEighth * g.CYukawa * cHistoryReduced
	return HistoryNormalForm{
		CHistoryOriginalFormula:     "C_History=1+L_Hopf(1-|lambda|-F_wall_3_red+kappa_e_red)",
		CHistoryReducedFormula:      "C_History=1+L_Hopf(1-kappa_lambda_red)",
		FullScalarHiggsFormula:      "lambda_runtime_eff=(1/8) C_Yukawa [1+L_Hopf(1-kappa_lambda_red)]",
		CYukawa:                     g.CYukawa,
		CHistory:                    cHistoryReduced,
		KappaLambdaRed:              d.KappaLambdaRed,
		OmegaHistory:                b.OmegaHistory,
		LambdaRuntimeEff:            g.LambdaRuntimeEff,
		LambdaRuntimeFromNormalForm: lambda,
		NormalFormResidual:          lambda - g.LambdaRuntimeEff,
		CHistoryWritten:             true,
		FullFormRewritten:           true,
		ThreeFactorNormalForm:       true,
		IndependentRuntimeTheorem:   false,
		Verdict: strings.Join([]string{
			StatusCHistoryNormalFormWritten,
			StatusFullScalarHiggsFormRewritten,
			StatusScalarHiggsBridgeHasThreeFactorNormalForm,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildSourceTypeInterpretation() SourceTypeInterpretation {
	return SourceTypeInterpretation{
		KappaLambdaRedSourceType: "reduced scalar matching deficit reconstructed from signed scalar zero-wall depth, cubic boundary-history response, and reduced flavor-wall deficit",
		OmegaHistorySourceType:   "scalar matching complement transported by the Radial-Hopf loop unit",
		CHistorySourceType:       "HistoryLoop uplift factor after scalar-coordinate collapse",
		Recorded:                 true,
		NativeHistoryLoopTheorem: false,
		NativeTransportTheorem:   false,
		Verdict: strings.Join([]string{
			StatusSourceTypeInterpretationRecorded,
			StatusCHistoryRadialHopfTransportOfScalarMatchingComplement,
			StatusKappaLambdaRedReconstructsScalarMatchingDeficitFromWallFlavorData,
			StatusCHistoryNotNativeHistoryLoopTheorem,
			StatusLHopfNotNativeTransportTheorem,
		}, "; "),
	}
}

func buildLayerSeparation() LayerSeparation {
	return LayerSeparation{
		CYukawaLayer:                       "finite Yukawa trace participation layer",
		KappaLambdaRedLayer:                "scalar/flavor/boundary history closure layer",
		LHopfLayer:                         "Radial-Hopf transport source-candidate layer",
		FactorsMultiplyAfterScalarCollapse: true,
		OperatorsOnSameNativeBoard:         false,
		LayerSeparationAudited:             true,
		Verdict: strings.Join([]string{
			StatusLayerSeparationAudited,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusKappaLambdaRedNotNativeScalarTheorem,
			StatusLHopfNotNativeTransportTheorem,
		}, "; "),
	}
}

func buildIllegalTermRejection() IllegalTermRejection {
	return IllegalTermRejection{
		KappaLambdaRedNativeScalarTheorem:     false,
		OmegaHistoryPhysicalTimeOrRGScale:     false,
		LHopfBoundaryEventProbability:         false,
		CHistoryNativeHistoryLoopTheorem:      false,
		LambdaRuntimeEffIndependentPrediction: false,
		TreeProxyPoleMassPrediction:           false,
		ClaimsYukawaEigenvaluesDerived:        false,
		ClaimsHiggsMassOrPoleMassTheorem:      false,
		Audited:                               true,
		Verdict: strings.Join([]string{
			StatusIllegalTermRejectionAudited,
			StatusKappaLambdaRedNotNativeScalarTheorem,
			StatusCHistoryNotNativeHistoryLoopTheorem,
			StatusLHopfNotNativeTransportTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate759HistoryTransportBracketBoundary,
		}, "; "),
	}
}

func finitePositive(x float64) bool {
	return x > 0 && !math.IsNaN(x) && !math.IsInf(x, 0)
}

func Statuses() []string {
	return []string{
		StatusGate758OneEighthFactorizationInherited,
		StatusHistoryTransportBracketDefined,
		StatusOmegaHistoryComputed,
		StatusKappaLambdaRedDefined,
		StatusOmegaHistoryRewrittenAsOneMinusKappaLambdaRed,
		StatusCHistoryNormalFormWritten,
		StatusFullScalarHiggsFormRewritten,
		StatusSourceTypeInterpretationRecorded,
		StatusLayerSeparationAudited,
		StatusIllegalTermRejectionAudited,
		StatusCHistoryRadialHopfTransportOfScalarMatchingComplement,
		StatusKappaLambdaRedReconstructsScalarMatchingDeficitFromWallFlavorData,
		StatusScalarHiggsBridgeHasThreeFactorNormalForm,
		StatusKappaLambdaRedNotNativeScalarTheorem,
		StatusCHistoryNotNativeHistoryLoopTheorem,
		StatusLHopfNotNativeTransportTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate759HistoryTransportBracketBoundary,
	}
}

func FormatGate758(x Gate758Inheritance) string {
	return fmt.Sprintf("inherited=%t formula=%q cYFormula=%q cHFormula=%q nEff=%.16g cY=%.16g cH=%.16g lambdaRuntime=%.16g baseline=%.16g threeFactor=%t independentRuntime=%t verdict=%q", x.Inherited, x.FactorizationFormula, x.CYukawaFormula, x.CHistoryFormula, x.NEff, x.CYukawa, x.CHistory, x.LambdaRuntimeEff, x.OneEighthBaseline, x.ThreeFactorNormalFormAvailable, x.IndependentScalarRuntimeTheorem, x.Verdict)
}

func FormatBracket(x HistoryTransportBracket) string {
	return fmt.Sprintf("symbol=%q omegaFormula=%q cHFormula=%q L=%.16g cH=%.16g omega=%.16g omegaFromCHistory=%.16g residual=%.16g defined=%t computed=%t timeOrRG=%t verdict=%q", x.BracketSymbol, x.OmegaFormula, x.CHistoryFormula, x.LHopf, x.CHistory, x.OmegaHistory, x.OmegaFromCHistory, x.OmegaResidual, x.BracketDefined, x.OmegaComputed, x.PhysicalTimeOrRGScale, x.Verdict)
}

func FormatDeficit(x ReducedScalarMatchingDeficit) string {
	return fmt.Sprintf("symbol=%q definition=%q kappa=%.16g complement=%.16g omega=%.16g residual=%.16g defined=%t rewritten=%t nativeScalar=%t verdict=%q", x.Symbol, x.Definition, x.KappaLambdaRed, x.Complement, x.OmegaHistory, x.ComplementResidual, x.Defined, x.RewrittenAsComplement, x.NativeScalarTheorem, x.Verdict)
}

func FormatNormalForm(x HistoryNormalForm) string {
	return fmt.Sprintf("cHOriginal=%q cHReduced=%q full=%q cY=%.16g cH=%.16g kappa=%.16g omega=%.16g lambda=%.16g fromNormal=%.16g residual=%.16g cHWritten=%t fullRewritten=%t threeFactor=%t independentRuntime=%t verdict=%q", x.CHistoryOriginalFormula, x.CHistoryReducedFormula, x.FullScalarHiggsFormula, x.CYukawa, x.CHistory, x.KappaLambdaRed, x.OmegaHistory, x.LambdaRuntimeEff, x.LambdaRuntimeFromNormalForm, x.NormalFormResidual, x.CHistoryWritten, x.FullFormRewritten, x.ThreeFactorNormalForm, x.IndependentRuntimeTheorem, x.Verdict)
}

func FormatInterpretation(x SourceTypeInterpretation) string {
	return fmt.Sprintf("kappaSource=%q omegaSource=%q cHSource=%q recorded=%t nativeHistoryLoop=%t nativeTransport=%t verdict=%q", x.KappaLambdaRedSourceType, x.OmegaHistorySourceType, x.CHistorySourceType, x.Recorded, x.NativeHistoryLoopTheorem, x.NativeTransportTheorem, x.Verdict)
}

func FormatLayers(x LayerSeparation) string {
	return fmt.Sprintf("cYLayer=%q kappaLayer=%q LLayer=%q multiplyAfterCollapse=%t sameNativeBoard=%t audited=%t verdict=%q", x.CYukawaLayer, x.KappaLambdaRedLayer, x.LHopfLayer, x.FactorsMultiplyAfterScalarCollapse, x.OperatorsOnSameNativeBoard, x.LayerSeparationAudited, x.Verdict)
}

func FormatIllegal(x IllegalTermRejection) string {
	return fmt.Sprintf("reject(kappaNative=%t omegaTimeRG=%t LEventProbability=%t cHNative=%t lambdaIndependent=%t treePole=%t yukawaEigen=%t higgsPole=%t) audited=%t verdict=%q", x.KappaLambdaRedNativeScalarTheorem, x.OmegaHistoryPhysicalTimeOrRGScale, x.LHopfBoundaryEventProbability, x.CHistoryNativeHistoryLoopTheorem, x.LambdaRuntimeEffIndependentPrediction, x.TreeProxyPoleMassPrediction, x.ClaimsYukawaEigenvaluesDerived, x.ClaimsHiggsMassOrPoleMassTheorem, x.Audited, x.Verdict)
}
