// Package generation2boundaryscalarprojectorselectorfactorizationfirewallaudit implements
// Gate 687: Boundary Scalar / Projector Selector Factorization Firewall Audit.
//
// Gate 686 separated the active bridge response into a boundary scalar,
// a Boolean-octonionic support-selected projector, and ordinary augmented trace
// scalarization. Gate 687 audits the firewall behind that separation: the
// boundary scalar S_split acts as S_split I_H72, hence is central in the
// finite/augmented endomorphism algebra and cannot by itself impose the support
// equations P_B P=P and P_G P=P. The projector identity must therefore remain
// selected by the separate native Boolean-octonionic support sieve at the
// current theorem level.
//
// This is a bridge-layer factorization firewall audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native 7/72 theorem, or a native projector-activation theorem.
package generation2boundaryscalarprojectorselectorfactorizationfirewallaudit

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE687-BOUNDARY-SCALAR-PROJECTOR-SELECTOR-FACTORIZATION-FIREWALL-AUDIT"

	StatusGate686SupportMinimalityInherited      = "PASS_GATE686_SUPPORT_MINIMALITY_INHERITED"
	StatusScalarActionCommutes                   = "PASS_SCALAR_ACTION_COMMUTES_WITH_PROJECTOR_ALGEBRA"
	StatusScalarAloneCannotSelectIdentity        = "PASS_SCALAR_ALONE_CANNOT_SELECT_PROJECTOR_IDENTITY"
	StatusNativeSupportSelectorRecorded          = "PASS_NATIVE_SUPPORT_SELECTOR_RECORDED"
	StatusResponseFactorizationWritten           = "PASS_RESPONSE_FACTORIZATION_WRITTEN"
	StatusThreeSealDecompositionDefined          = "PASS_THREE_SEAL_DECOMPOSITION_DEFINED"
	StatusActiveResponseFactorsScalarAndSelector = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_FACTORS_INTO_BOUNDARY_SCALAR_AND_NATIVE_PROJECTOR_SELECTOR"
	StatusProjectorIdentityNativeSupportSealed   = "CONDITIONAL_SUPPORT_PROJECTOR_IDENTITY_SELECTION_IS_NATIVE_SUPPORT_SEALED"
	StatusSSplitAloneDoesNotImposeSupport        = "FAILED_ROUTE_S_SPLIT_ALONE_DOES_NOT_IMPOSE_BOOLEAN_OCTONIONIC_SUPPORT"
	StatusNoBoundaryScalarToSupportCoupling      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_SCALAR_TO_SUPPORT_COUPLING_THEOREM"
	StatusNoNativeProjectorActivationTheorem     = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_ACTIVATION_THEOREM"
	StatusNoNativeSevenOver72Theorem             = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate687FactorizationBoundary           = "FIREWALL_PRESERVED_GATE687_SCALAR_PROJECTOR_FACTORIZATION_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	booleanRank       = 56
	octonionicRank    = 14
	k7Dimension       = 7
	uPlusVDimension   = booleanRank + octonionicRank - k7Dimension
	w7Dimension       = lambda4Dimension - uPlusVDimension
	auditedDBase      = 0.0001256552099683575
	auditedSSplit     = 0.0012924448188162962
)

type Gate686Inheritance struct {
	SupportMinimalityInherited bool
	RankSevenTraceInherited    bool
	BooleanSupportRequired     bool
	OctonionicSupportRequired  bool
	SelectedProjector          string
	BoundaryScalar             string
	TraceScalarization         string
	H72Dimension               int
	K7Dimension                int
	DBase                      float64
	SSplit                     float64
	PriorFirewallPreserved     bool
	Verdict                    string
}

type ScalarActionAudit struct {
	ScalarOperator              string
	CommutesWithPB              bool
	CommutesWithPG              bool
	CommutesWithAnyProjector    bool
	CentralAction               bool
	CarriesProjectorDirection   bool
	CanDistinguishPK7FromPW7    bool
	CanImposeBooleanSupport     bool
	CanImposeOctonionicSupport  bool
	OnlyScalesSelectedProjector bool
	Verdict                     string
}

type ProjectorCandidate struct {
	Name                    string
	Rank                    int
	Carrier                 string
	PassesBooleanSupport    bool
	PassesOctonionicSupport bool
	SelectedByNativeSupport bool
	DistinguishedByScalar   bool
	ScalarActionDescription string
	SupportSelectionVerdict string
}

type ScalarIndistinguishabilityAudit struct {
	Candidates                 []ProjectorCandidate
	AllRankSevenScaled         bool
	ScalarSeparatesCandidates  bool
	SupportSeparatesCandidates bool
	PK7SelectedBySupport       bool
	PW7RejectedBySupport       bool
	Verdict                    string
}

type SupportSelectionAudit struct {
	Constraints             []string
	ImageInBooleanSector    bool
	ImageInOctonionicSector bool
	ImageInIntersection     bool
	IntersectionDimension   int
	RankEqualsIntersection  bool
	SelectedProjector       string
	IndependentOfSSplit     bool
	Verdict                 string
}

type ThreeSealDecompositionAudit struct {
	BoundaryAmplitudeSeal             string
	NativeProjectorSelectorSeal       string
	TraceScalarizationSeal            string
	BoundaryScalarControlsAmplitude   bool
	ProjectorSelectorControlsIdentity bool
	TraceControlsScalarResponse       bool
	DBaseApproximation                string
	Verdict                           string
}

type FactorizationAudit struct {
	ActiveResponse                 string
	FactorizationRequired          bool
	BoundaryScalarFactor           string
	ProjectorSelectorFactor        string
	TraceFactor                    string
	SSplitAloneSelectsIdentity     bool
	ProjectorIdentitySupportSealed bool
	NativeCouplingProved           bool
	Verdict                        string
}

type NoGoAudit struct {
	BlockedRoute                 string
	Reason                       string
	ScalarCommutatorData         []string
	ScalarDirectionInformation   bool
	BoundaryScalarImposesSupport bool
	NoGoCertified                bool
	Verdict                      string
}

type MissingTheoremAudit struct {
	FutureTargets []string
	PreciseGap    string
	Verdict       string
}

type VerdictDiscipline struct {
	ClaimsScalarSelectsProjector        bool
	ClaimsScalarImposesSupport          bool
	ClaimsBoundaryScalarSupportCoupling bool
	ClaimsProjectorActivation           bool
	ClaimsNativeSevenOver72             bool
	ClaimsBoundaryStressDerivation      bool
	ClaimsScalarRGMatching              bool
	ClaimsHiggsMass                     bool
	ClaimsGaugeUnification              bool
	ClaimsFlavorDerivation              bool
	Verdict                             string
}

type Analysis struct {
	Inherited            Gate686Inheritance
	ScalarAction         ScalarActionAudit
	Indistinguishability ScalarIndistinguishabilityAudit
	SupportSelection     SupportSelectionAudit
	ThreeSeal            ThreeSealDecompositionAudit
	Factorization        FactorizationAudit
	NoGo                 NoGoAudit
	Missing              MissingTheoremAudit
	Discipline           VerdictDiscipline
	Truth                string
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
	inherited := buildInheritance()
	scalar := buildScalarAction()
	indistinguishability := buildIndistinguishability()
	support := buildSupportSelection()
	threeSeal := buildThreeSealDecomposition()
	factorization := buildFactorization(scalar, support)
	noGo := buildNoGo(scalar)
	missing := MissingTheoremAudit{
		FutureTargets: []string{
			"BoundaryScalarToNativeSupportCouplingTheorem",
			"HistoryResponseFactorizationTheorem",
			StatusNoNativeProjectorActivationTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native coupling principle explaining why the physical history response factorizes as boundary scalar amplitude times the Boolean-octonionic intersection projector, rather than allowing S_split I_H72 to select projector support by scalar action alone",
		Verdict: strings.Join([]string{
			StatusNoBoundaryScalarToSupportCoupling,
			StatusNoNativeProjectorActivationTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate687FactorizationBoundary}
	truth := "Gate 687 sharpens the Gate686 activation gap. S_split is a scalar boundary coordinate and acts as S_split I_H72, so it commutes with P_B, P_G, and every candidate projector. A central scalar can scale P_K7, P_W7, or any other rank-seven projector, but it carries no projector-direction information and cannot impose P_B P=P or P_G P=P. Therefore the active bridge must remain factorized at the current theorem level: boundary amplitude S_split times the separately support-selected native projector P_K7, followed by ordinary augmented trace scalarization. The missing object is now a boundary-scalar-to-native-support coupling/factorization theorem, not a trace or scalar-centrality theorem."
	return Analysis{Inherited: inherited, ScalarAction: scalar, Indistinguishability: indistinguishability, SupportSelection: support, ThreeSeal: threeSeal, Factorization: factorization, NoGo: noGo, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance() Gate686Inheritance {
	return Gate686Inheritance{
		SupportMinimalityInherited: true,
		RankSevenTraceInherited:    true,
		BooleanSupportRequired:     true,
		OctonionicSupportRequired:  true,
		SelectedProjector:          "P_K7",
		BoundaryScalar:             "S_split=lambda(Lambda_12)+(R_3-1)",
		TraceScalarization:         "Tr_H72(S_split P_K7)/Tr_H72(I)=(7/72)S_split",
		H72Dimension:               h72Dimension,
		K7Dimension:                k7Dimension,
		DBase:                      auditedDBase,
		SSplit:                     auditedSSplit,
		PriorFirewallPreserved:     true,
		Verdict:                    StatusGate686SupportMinimalityInherited,
	}
}

func buildScalarAction() ScalarActionAudit {
	return ScalarActionAudit{
		ScalarOperator:              "S_split I_H72",
		CommutesWithPB:              true,
		CommutesWithPG:              true,
		CommutesWithAnyProjector:    true,
		CentralAction:               true,
		CarriesProjectorDirection:   false,
		CanDistinguishPK7FromPW7:    false,
		CanImposeBooleanSupport:     false,
		CanImposeOctonionicSupport:  false,
		OnlyScalesSelectedProjector: true,
		Verdict: strings.Join([]string{
			StatusScalarActionCommutes,
			StatusScalarAloneCannotSelectIdentity,
			StatusSSplitAloneDoesNotImposeSupport,
		}, "; "),
	}
}

func buildIndistinguishability() ScalarIndistinguishabilityAudit {
	candidates := []ProjectorCandidate{
		{
			Name:                    "P_K7",
			Rank:                    k7Dimension,
			Carrier:                 "K_7=Im(P_B)∩Im(P_G)",
			PassesBooleanSupport:    true,
			PassesOctonionicSupport: true,
			SelectedByNativeSupport: true,
			DistinguishedByScalar:   false,
			ScalarActionDescription: "S_split P_K7 is the scalar multiple of the already selected K_7 projector",
			SupportSelectionVerdict: StatusProjectorIdentityNativeSupportSealed,
		},
		{
			Name:                    "P_W7",
			Rank:                    w7Dimension,
			Carrier:                 "W_7=(Im(P_B)+Im(P_G))^perp inside Lambda^4 R^8",
			PassesBooleanSupport:    false,
			PassesOctonionicSupport: false,
			SelectedByNativeSupport: false,
			DistinguishedByScalar:   false,
			ScalarActionDescription: "S_split P_W7 is equally a scalar multiple of a rank-seven projector; scalar action alone does not reject it",
			SupportSelectionVerdict: StatusSSplitAloneDoesNotImposeSupport,
		},
		{
			Name:                    "P_arbitrary7",
			Rank:                    k7Dimension,
			Carrier:                 "generic rank-seven subspace of H_72",
			PassesBooleanSupport:    false,
			PassesOctonionicSupport: false,
			SelectedByNativeSupport: false,
			DistinguishedByScalar:   false,
			ScalarActionDescription: "S_split P_arbitrary7 has the same central-scalar form once that projector is supplied externally",
			SupportSelectionVerdict: StatusSSplitAloneDoesNotImposeSupport,
		},
	}
	return ScalarIndistinguishabilityAudit{
		Candidates:                 candidates,
		AllRankSevenScaled:         true,
		ScalarSeparatesCandidates:  false,
		SupportSeparatesCandidates: true,
		PK7SelectedBySupport:       true,
		PW7RejectedBySupport:       true,
		Verdict: strings.Join([]string{
			StatusScalarAloneCannotSelectIdentity,
			StatusNativeSupportSelectorRecorded,
			StatusProjectorIdentityNativeSupportSealed,
		}, "; "),
	}
}

func buildSupportSelection() SupportSelectionAudit {
	return SupportSelectionAudit{
		Constraints: []string{
			"P^2=P",
			"P^T=P",
			"rank(P)=7",
			"P_B P=P",
			"P_G P=P",
		},
		ImageInBooleanSector:    true,
		ImageInOctonionicSector: true,
		ImageInIntersection:     true,
		IntersectionDimension:   k7Dimension,
		RankEqualsIntersection:  true,
		SelectedProjector:       "P_K7",
		IndependentOfSSplit:     true,
		Verdict: strings.Join([]string{
			StatusNativeSupportSelectorRecorded,
			StatusProjectorIdentityNativeSupportSealed,
		}, "; "),
	}
}

func buildThreeSealDecomposition() ThreeSealDecompositionAudit {
	return ThreeSealDecompositionAudit{
		BoundaryAmplitudeSeal:             "S_split=lambda(Lambda_12)+(R_3-1)",
		NativeProjectorSelectorSeal:       "rank seven + Boolean support + octonionic support => P_K7",
		TraceScalarizationSeal:            "Tr_H72(S_split P_K7)/72=(7/72)S_split",
		BoundaryScalarControlsAmplitude:   true,
		ProjectorSelectorControlsIdentity: true,
		TraceControlsScalarResponse:       true,
		DBaseApproximation:                "D_base ≈ TraceScalarizationSeal(BoundaryAmplitudeSeal · NativeProjectorSelectorSeal)",
		Verdict: strings.Join([]string{
			StatusResponseFactorizationWritten,
			StatusThreeSealDecompositionDefined,
			StatusActiveResponseFactorsScalarAndSelector,
		}, "; "),
	}
}

func buildFactorization(scalar ScalarActionAudit, support SupportSelectionAudit) FactorizationAudit {
	return FactorizationAudit{
		ActiveResponse:                 "R_split = S_split · P_selected with P_selected=P_K7 under the native support sieve",
		FactorizationRequired:          scalar.CentralAction && !scalar.CarriesProjectorDirection && support.SelectedProjector == "P_K7",
		BoundaryScalarFactor:           "BoundaryAmplitudeSeal(S_split)",
		ProjectorSelectorFactor:        "NativeProjectorSelectorSeal(P_K7)",
		TraceFactor:                    "TraceScalarizationSeal(Tr_H72/72)",
		SSplitAloneSelectsIdentity:     false,
		ProjectorIdentitySupportSealed: true,
		NativeCouplingProved:           false,
		Verdict: strings.Join([]string{
			StatusResponseFactorizationWritten,
			StatusActiveResponseFactorsScalarAndSelector,
			StatusProjectorIdentityNativeSupportSealed,
			StatusNoBoundaryScalarToSupportCoupling,
		}, "; "),
	}
}

func buildNoGo(scalar ScalarActionAudit) NoGoAudit {
	return NoGoAudit{
		BlockedRoute: "S_split alone => P_B P=P and P_G P=P",
		Reason:       "S_split I_H72 is central; central scalar multiplication has no subspace orientation and cannot encode Boolean or octonionic support equations.",
		ScalarCommutatorData: []string{
			"[S_split I_H72, P_B]=0",
			"[S_split I_H72, P_G]=0",
			"[S_split I_H72, P]=0 for every candidate projector P",
		},
		ScalarDirectionInformation:   scalar.CarriesProjectorDirection,
		BoundaryScalarImposesSupport: scalar.CanImposeBooleanSupport || scalar.CanImposeOctonionicSupport,
		NoGoCertified:                scalar.CentralAction && !scalar.CarriesProjectorDirection && !scalar.CanImposeBooleanSupport && !scalar.CanImposeOctonionicSupport,
		Verdict: strings.Join([]string{
			StatusScalarActionCommutes,
			StatusSSplitAloneDoesNotImposeSupport,
			StatusNoBoundaryScalarToSupportCoupling,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate686SupportMinimalityInherited,
		StatusScalarActionCommutes,
		StatusScalarAloneCannotSelectIdentity,
		StatusNativeSupportSelectorRecorded,
		StatusResponseFactorizationWritten,
		StatusThreeSealDecompositionDefined,
		StatusActiveResponseFactorsScalarAndSelector,
		StatusProjectorIdentityNativeSupportSealed,
		StatusSSplitAloneDoesNotImposeSupport,
		StatusNoBoundaryScalarToSupportCoupling,
		StatusNoNativeProjectorActivationTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate687FactorizationBoundary,
	}
}

func sortedCandidateNames(candidates []ProjectorCandidate) []string {
	out := make([]string, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, c.Name)
	}
	sort.Strings(out)
	return out
}

func FormatInheritance(x Gate686Inheritance) string {
	return fmt.Sprintf("minimalityInherited=%t rankTrace=%t boolSupport=%t octSupport=%t selected=%q scalar=%q trace=%q h72=%d k7=%d dbase=%.18g ssplit=%.18g firewall=%t verdict=%q", x.SupportMinimalityInherited, x.RankSevenTraceInherited, x.BooleanSupportRequired, x.OctonionicSupportRequired, x.SelectedProjector, x.BoundaryScalar, x.TraceScalarization, x.H72Dimension, x.K7Dimension, x.DBase, x.SSplit, x.PriorFirewallPreserved, x.Verdict)
}

func FormatScalarAction(x ScalarActionAudit) string {
	return fmt.Sprintf("operator=%q commutePB=%t commutePG=%t commuteP=%t central=%t direction=%t distinguishPK7PW7=%t imposeBool=%t imposeOct=%t scalesOnly=%t verdict=%q", x.ScalarOperator, x.CommutesWithPB, x.CommutesWithPG, x.CommutesWithAnyProjector, x.CentralAction, x.CarriesProjectorDirection, x.CanDistinguishPK7FromPW7, x.CanImposeBooleanSupport, x.CanImposeOctonionicSupport, x.OnlyScalesSelectedProjector, x.Verdict)
}

func FormatCandidate(x ProjectorCandidate) string {
	return fmt.Sprintf("name=%q rank=%d carrier=%q bool=%t oct=%t supportSelected=%t scalarDistinguishes=%t scalar=%q supportVerdict=%q", x.Name, x.Rank, x.Carrier, x.PassesBooleanSupport, x.PassesOctonionicSupport, x.SelectedByNativeSupport, x.DistinguishedByScalar, x.ScalarActionDescription, x.SupportSelectionVerdict)
}

func FormatIndistinguishability(x ScalarIndistinguishabilityAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] names=[%s] allScaled=%t scalarSeparates=%t supportSeparates=%t pk7=%t pw7Rejected=%t verdict=%q", strings.Join(parts, " | "), strings.Join(sortedCandidateNames(x.Candidates), ", "), x.AllRankSevenScaled, x.ScalarSeparatesCandidates, x.SupportSeparatesCandidates, x.PK7SelectedBySupport, x.PW7RejectedBySupport, x.Verdict)
}

func FormatSupportSelection(x SupportSelectionAudit) string {
	return fmt.Sprintf("constraints=[%s] inBoolean=%t inOct=%t inIntersection=%t intersectionDim=%d rankEquals=%t selected=%q independentOfSSplit=%t verdict=%q", strings.Join(x.Constraints, "; "), x.ImageInBooleanSector, x.ImageInOctonionicSector, x.ImageInIntersection, x.IntersectionDimension, x.RankEqualsIntersection, x.SelectedProjector, x.IndependentOfSSplit, x.Verdict)
}

func FormatThreeSeal(x ThreeSealDecompositionAudit) string {
	return fmt.Sprintf("boundary=%q selector=%q trace=%q amplitude=%t identity=%t scalarResponse=%t dbase=%q verdict=%q", x.BoundaryAmplitudeSeal, x.NativeProjectorSelectorSeal, x.TraceScalarizationSeal, x.BoundaryScalarControlsAmplitude, x.ProjectorSelectorControlsIdentity, x.TraceControlsScalarResponse, x.DBaseApproximation, x.Verdict)
}

func FormatFactorization(x FactorizationAudit) string {
	return fmt.Sprintf("response=%q required=%t boundaryFactor=%q projectorFactor=%q traceFactor=%q ssplitSelectsIdentity=%t supportSealed=%t couplingProved=%t verdict=%q", x.ActiveResponse, x.FactorizationRequired, x.BoundaryScalarFactor, x.ProjectorSelectorFactor, x.TraceFactor, x.SSplitAloneSelectsIdentity, x.ProjectorIdentitySupportSealed, x.NativeCouplingProved, x.Verdict)
}

func FormatNoGo(x NoGoAudit) string {
	return fmt.Sprintf("blocked=%q reason=%q commutators=[%s] scalarDirection=%t imposesSupport=%t noGo=%t verdict=%q", x.BlockedRoute, x.Reason, strings.Join(x.ScalarCommutatorData, "; "), x.ScalarDirectionInformation, x.BoundaryScalarImposesSupport, x.NoGoCertified, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("future=[%s] precise=%q verdict=%q", strings.Join(x.FutureTargets, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsScalarSelectsProjector=%t claimsScalarImposesSupport=%t claimsCoupling=%t claimsProjectorActivation=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsScalarSelectsProjector, x.ClaimsScalarImposesSupport, x.ClaimsBoundaryScalarSupportCoupling, x.ClaimsProjectorActivation, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
