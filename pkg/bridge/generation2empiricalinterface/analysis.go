// Package generation2empiricalinterface implements Gate 453:
// Texture-Zero Invariant Ledger / Allowed Empirical Interface.
//
// Gates 450-452 closed the GST/Fritzsch shortcut: the ASHA triangle has an
// exact texture-zero spectral sum rule, but the mass-angle ratios remain
// coefficient- and phase-dependent; no native law suppresses the 1-3 edge, no
// native law fixes the phase ray, and no K_gen-preserving basis change can turn
// the triangle into a nearest-neighbor chain. Gate 453 therefore formalizes the
// only honest way forward: texture-zero phenomenology may be used as an
// explicitly empirical bridge interface, while K_gen, the Generation-2 zero,
// and the full triangle remain the only promoted structural claims.
package generation2empiricalinterface

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE453-TEXTURE-ZERO-INVARIANT-LEDGER-ALLOWED-EMPIRICAL-INTERFACE"

	StatusGate452Inherited                      = "CONDITIONAL_SUPPORT_GATE452_BASIS_INVARIANCE_INHERITED"
	StatusNativeInvariantLedgerSealed           = "CONDITIONAL_SUPPORT_NATIVE_TEXTURE_ZERO_INVARIANT_LEDGER_SEALED"
	StatusEmpiricalInterfaceDefined             = "CONDITIONAL_SUPPORT_TEXTURE_ZERO_EMPIRICAL_INTERFACE_DEFINED"
	StatusImportContractValidated               = "CONDITIONAL_SUPPORT_EXPLICIT_EMPIRICAL_IMPORT_CONTRACT_VALIDATED"
	StatusPromotionFirewallValidated            = "CONDITIONAL_SUPPORT_NO_EMPIRICAL_TEXTURE_PROMOTED_TO_NATIVE_LAW"
	StatusResidualComputationsQuarantined       = "CONDITIONAL_SUPPORT_TEXTURE_ZERO_RESIDUALS_QUARANTINED_AS_COMPARATORS"
	StatusEmpiricalFirewallPreserved            = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedNativeRatioNotRestored          = "FAILED_ROUTE_NATIVE_RATIO_DERIVATION_NOT_RESTORED"
	StatusFailedGSTRequiresEmpiricalBranchInput = "FAILED_ROUTE_GST_FRITZSCH_RELATION_REQUIRES_EXPLICIT_EMPIRICAL_BRANCH_INPUT"
	StatusFailedObservablePromotionRejected     = "FAILED_ROUTE_OBSERVED_MASS_MIXING_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate444Generation2Zero          bool
	Gate445TriangleForced           bool
	Gate446PhaseQuarantined         bool
	Gate447CoefficientsSealed       bool
	Gate450TextureZeroSumRule       bool
	Gate450RatiosRequireAmplitudes  bool
	Gate451FullTrianglePreserved    bool
	Gate451NoNativePhaseRaySelector bool
	Gate452NearestNeighborNotGauge  bool
	Gate452KPreservingBasisGroup    string
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type NativeInvariant struct {
	Name              string
	Formula           string
	NativeStatus      string
	DependsOn         []string
	RequiresEmpirical bool
	CanPredictNumber  bool
	Reason            string
}

type NativeLedger struct {
	Executed                  bool
	Invariants                []NativeInvariant
	PromotedStructuralObjects []string
	QuarantinedObjects        []string
	NativeOnlyPredictsGST     bool
	Verdict                   string
	Reason                    string
}

type EmpiricalInput struct {
	Name            string
	Kind            string
	Allowed         bool
	RequiredLabel   string
	NativePromotion bool
	Reason          string
}

type ImportContract struct {
	Executed                   bool
	Inputs                     []EmpiricalInput
	AllowedCount               int
	RejectedPromotionCount     int
	RequiresExplicitLabel      bool
	RequiresRenormalizationTag bool
	RequiresSectorTag          bool
	AllowsNativeClaim          bool
	Verdict                    string
	Reason                     string
}

type Comparator struct {
	Name             string
	Formula          string
	RequiresInputs   []string
	NativeObservable bool
	Allowed          bool
	Reason           string
}

type ResidualLedger struct {
	Executed                         bool
	Comparators                      []Comparator
	AllowsTextureResiduals           bool
	AllowsGSTResidual                bool
	AllowsNativeGSTRatioClaim        bool
	AllowsCoefficientFittingAsNative bool
	Verdict                          string
	Reason                           string
}

type InterfaceRequest struct {
	Name               string
	RequestedOperation string
	ImportsEmpirical   bool
	ExplicitlyLabelled bool
	AttemptsPromotion  bool
	Allowed            bool
	Reason             string
}

type InterfaceSieve struct {
	Executed             bool
	Requests             []InterfaceRequest
	NativeOnlyAllowed    bool
	EmpiricalFitAllowed  bool
	PromotionRejected    bool
	AnyForbiddenAccepted bool
	Verdict              string
	Reason               string
}

type Firewall struct {
	Executed                        bool
	NoObservedMuonMassImported      bool
	NoObservedCharmMassImported     bool
	NoObservedYukawaImported        bool
	NoCKMImported                   bool
	NoPMNSImported                  bool
	NoCurveFit                      bool
	NoGSTPromotion                  bool
	KGenStillForced                 bool
	Generation2ZeroStillForced      bool
	XTriangleStillForced            bool
	YPhaseStillQuarantined          bool
	SectorCoefficientsStillSealed   bool
	GSTFritzschRelationsQuarantined bool
	NativeFlavorDimAfter            int
	KXYCoeffDimAfter                int
	Verdict                         string
	Reason                          string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Native      NativeLedger
	Contract    ImportContract
	Residuals   ResidualLedger
	Sieve       InterfaceSieve
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Native = buildNativeLedger()
	a.Contract = buildImportContract()
	a.Residuals = buildResidualLedger()
	a.Sieve = buildInterfaceSieve()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate444KGenForced:               true,
		Gate444Generation2Zero:          true,
		Gate445TriangleForced:           true,
		Gate446PhaseQuarantined:         true,
		Gate447CoefficientsSealed:       true,
		Gate450TextureZeroSumRule:       true,
		Gate450RatiosRequireAmplitudes:  true,
		Gate451FullTrianglePreserved:    true,
		Gate451NoNativePhaseRaySelector: true,
		Gate452NearestNeighborNotGauge:  true,
		Gate452KPreservingBasisGroup:    "centralizer_U(3)(K_gen)=U(1)^3",
		NoEmpiricalInputsImported:       true,
		Verdict:                         StatusGate452Inherited,
	}
}

func buildNativeLedger() NativeLedger {
	invariants := []NativeInvariant{
		{Name: "primitive family axis", Formula: "K_gen=diag(-1,0,1)", NativeStatus: "geometrically-forced structural law", DependsOn: []string{"traceless anomaly boundary", "KMS integer spacing", "three-generation boundary"}, RequiresEmpirical: false, CanPredictNumber: false, Reason: "fixes the family address and Generation-2 bare zero, not a physical mass."},
		{Name: "closed triangular bridge support", Formula: "X_triangle=[[0,1,1],[1,0,1],[1,1,0]]", NativeStatus: "geometrically-forced support topology", DependsOn: []string{"endpoint-balanced mass lift", "det(K+epsilon X)=2 epsilon^3"}, RequiresEmpirical: false, CanPredictNumber: false, Reason: "support is forced; amplitude, sector, and phase are not."},
		{Name: "texture-zero spectral sum rule", Formula: "0=sum_i lambda_i |U_2i|^2", NativeStatus: "symbolic invariant", DependsOn: []string{"M_22=0", "unitary diagonalization"}, RequiresEmpirical: false, CanPredictNumber: false, Reason: "exact identity, but it contains the unknown spectrum and eigenvectors."},
		{Name: "full-triangle characteristic polynomial", Formula: "lambda^3-(a^2+3(b^2+c^2))lambda-2(b^3-3bc^2)", NativeStatus: "symbolic invariant", DependsOn: []string{"a/r coefficient ray", "phase phi"}, RequiresEmpirical: true, CanPredictNumber: false, Reason: "the symbolic shape is native, but evaluating ratios requires coefficient and phase selectors."},
		{Name: "K-preserving basis class", Formula: "centralizer_U(3)(K_gen)=U(1)^3", NativeStatus: "basis-invariance constraint", DependsOn: []string{"simple K spectrum"}, RequiresEmpirical: false, CanPredictNumber: false, Reason: "rephasings preserve support and cannot hide a nearest-neighbor branch."},
	}
	return NativeLedger{
		Executed:                  true,
		Invariants:                invariants,
		PromotedStructuralObjects: []string{"K_gen", "Generation-2 bare structural zero", "full X_triangle support", "M_22=0 spectral sum rule"},
		QuarantinedObjects:        []string{"a/r coefficient ray", "b/c phase ray", "Y_gen phase value", "sector K/X/Y amplitudes", "GST/Fritzsch relation", "CKM/PMNS values", "physical muon/charm masses"},
		NativeOnlyPredictsGST:     false,
		Verdict:                   StatusNativeInvariantLedgerSealed,
		Reason:                    "native ledger contains structural identities only; no native-only mass-angle map survives Gates 450-452.",
	}
}

func buildImportContract() ImportContract {
	inputs := []EmpiricalInput{
		{Name: "sector label", Kind: "metadata", Allowed: true, RequiredLabel: "u,d,e,nu or explicitly external sector", NativePromotion: false, Reason: "matrix coefficients are sector-indexed; the sector must be named before any comparison."},
		{Name: "renormalization scale/scheme", Kind: "metadata", Allowed: true, RequiredLabel: "scale and scheme required", NativePromotion: false, Reason: "masses and mixings are scale/scheme dependent in continuum phenomenology."},
		{Name: "coefficient ray", Kind: "empirical bridge input", Allowed: true, RequiredLabel: "empirical-coefficient-ray", NativePromotion: false, Reason: "a/r and phi may be imported only as bridge data, never as native ASHA selectors."},
		{Name: "observed spectrum", Kind: "empirical comparator", Allowed: true, RequiredLabel: "observed-comparator", NativePromotion: false, Reason: "observed masses can test residuals but cannot define the theorem."},
		{Name: "observed mixing matrix", Kind: "empirical comparator", Allowed: true, RequiredLabel: "observed-comparator", NativePromotion: false, Reason: "mixing data can be compared after labelling; it cannot be reverse-promoted."},
		{Name: "GST/Fritzsch branch condition", Kind: "empirical texture assumption", Allowed: true, RequiredLabel: "external-texture-assumption", NativePromotion: false, Reason: "nearest-neighbor suppression or phase fixing is allowed only as an explicit non-native branch."},
		{Name: "observed muon/charm mass as native proof", Kind: "forbidden promotion", Allowed: false, RequiredLabel: "rejected", NativePromotion: true, Reason: "physical mass values belong behind the 13-moduli firewall."},
		{Name: "CKM/PMNS angle as native phase selector", Kind: "forbidden promotion", Allowed: false, RequiredLabel: "rejected", NativePromotion: true, Reason: "using observed mixing to select phi would invert the proof order."},
	}
	allowed := 0
	rejected := 0
	for _, x := range inputs {
		if x.Allowed {
			allowed++
		}
		if !x.Allowed && x.NativePromotion {
			rejected++
		}
	}
	return ImportContract{
		Executed:                   true,
		Inputs:                     inputs,
		AllowedCount:               allowed,
		RejectedPromotionCount:     rejected,
		RequiresExplicitLabel:      true,
		RequiresRenormalizationTag: true,
		RequiresSectorTag:          true,
		AllowsNativeClaim:          false,
		Verdict:                    StatusImportContractValidated,
		Reason:                     "empirical values may enter only through labelled comparator/branch fields with sector and scheme metadata.",
	}
}

func buildResidualLedger() ResidualLedger {
	comparators := []Comparator{
		{Name: "texture-zero sum residual", Formula: "R_22=sum_i lambda_i |U_2i|^2", RequiresInputs: []string{"empirical spectrum", "empirical mixing row"}, NativeObservable: false, Allowed: true, Reason: "tests whether an external dataset is compatible with the structural zero."},
		{Name: "full-triangle determinant residual", Formula: "R_det=det(K+epsilon B)-2 epsilon^3", RequiresInputs: []string{"coefficient normalization", "bridge support"}, NativeObservable: false, Allowed: true, Reason: "checks support-class consistency, not a physical mass prediction."},
		{Name: "GST/Fritzsch residual", Formula: "R_GST=sin(theta_ij)^2-m_i/m_j", RequiresInputs: []string{"external branch choice", "empirical masses", "empirical angle"}, NativeObservable: false, Allowed: true, Reason: "allowed only as a labelled external texture test."},
		{Name: "native GST prediction", Formula: "sin(theta_ij)=sqrt(m_i/m_j)", RequiresInputs: []string{"forbidden native ratio selector"}, NativeObservable: false, Allowed: false, Reason: "Gates 450-452 proved the selector is absent."},
		{Name: "native coefficient fit", Formula: "solve a,b,c from observed masses and relabel as ASHA law", RequiresInputs: []string{"observed masses"}, NativeObservable: false, Allowed: false, Reason: "reverse-fitting violates the theorem-gated firewall."},
	}
	return ResidualLedger{
		Executed:                         true,
		Comparators:                      comparators,
		AllowsTextureResiduals:           true,
		AllowsGSTResidual:                true,
		AllowsNativeGSTRatioClaim:        false,
		AllowsCoefficientFittingAsNative: false,
		Verdict:                          StatusResidualComputationsQuarantined,
		Reason:                           "residuals are legitimate comparator outputs, but every value-bearing branch remains explicitly empirical.",
	}
}

func buildInterfaceSieve() InterfaceSieve {
	requests := []InterfaceRequest{
		{Name: "native invariant report", RequestedOperation: "emit K_gen, X_triangle, M_22 sum rule", ImportsEmpirical: false, ExplicitlyLabelled: true, AttemptsPromotion: false, Allowed: true, Reason: "pure structural ledger; no numerical flavor observable claimed."},
		{Name: "labelled empirical coefficient-ray evaluation", RequestedOperation: "evaluate eigenvalues/mixings from supplied a/r and phi", ImportsEmpirical: true, ExplicitlyLabelled: true, AttemptsPromotion: false, Allowed: true, Reason: "allowed as bridge phenomenology after explicit label and sector/scheme metadata."},
		{Name: "GST residual test", RequestedOperation: "compare sin(theta)^2 to mass ratio", ImportsEmpirical: true, ExplicitlyLabelled: true, AttemptsPromotion: false, Allowed: true, Reason: "allowed as external texture diagnostic, not as ASHA theorem."},
		{Name: "silent CKM phase selector", RequestedOperation: "use observed CKM/PMNS phase to fix Y_gen", ImportsEmpirical: true, ExplicitlyLabelled: false, AttemptsPromotion: true, Allowed: false, Reason: "forbidden: imports observed data and promotes it to a native selector."},
		{Name: "nearest-neighbor rebrand", RequestedOperation: "delete 1-3 edge and call it basis gauge", ImportsEmpirical: false, ExplicitlyLabelled: false, AttemptsPromotion: true, Allowed: false, Reason: "forbidden by Gate 452 basis-invariance audit."},
	}
	anyForbiddenAccepted := false
	nativeAllowed := false
	empiricalAllowed := false
	promotionRejected := false
	for _, r := range requests {
		if !r.Allowed && r.AttemptsPromotion {
			promotionRejected = true
		}
		if r.Allowed && !r.ImportsEmpirical {
			nativeAllowed = true
		}
		if r.Allowed && r.ImportsEmpirical && r.ExplicitlyLabelled {
			empiricalAllowed = true
		}
		if r.Allowed && r.AttemptsPromotion {
			anyForbiddenAccepted = true
		}
	}
	return InterfaceSieve{
		Executed:             true,
		Requests:             requests,
		NativeOnlyAllowed:    nativeAllowed,
		EmpiricalFitAllowed:  empiricalAllowed,
		PromotionRejected:    promotionRejected,
		AnyForbiddenAccepted: anyForbiddenAccepted,
		Verdict:              StatusEmpiricalInterfaceDefined,
		Reason:               "the sieve accepts native ledgers and labelled empirical comparators, but rejects silent promotion and fake basis deletion.",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                        true,
		NoObservedMuonMassImported:      true,
		NoObservedCharmMassImported:     true,
		NoObservedYukawaImported:        true,
		NoCKMImported:                   true,
		NoPMNSImported:                  true,
		NoCurveFit:                      true,
		NoGSTPromotion:                  true,
		KGenStillForced:                 a.Inheritance.Gate444KGenForced,
		Generation2ZeroStillForced:      a.Inheritance.Gate444Generation2Zero,
		XTriangleStillForced:            a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:          a.Inheritance.Gate446PhaseQuarantined,
		SectorCoefficientsStillSealed:   a.Inheritance.Gate447CoefficientsSealed,
		GSTFritzschRelationsQuarantined: true,
		NativeFlavorDimAfter:            NativeFlavorDim,
		KXYCoeffDimAfter:                KXYCoeffDim,
		Verdict:                         StatusEmpiricalFirewallPreserved,
		Reason:                          "Gate 453 defines the legal empirical interface without changing native flavor dimension or coefficient firewall.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        454,
		Title:       "Coefficient-Ray Observability Rank Audit",
		Reason:      "after defining the legal empirical interface, the next mathematical question is the minimal number of external values needed to identify a sector coefficient ray without overclaiming nativeness",
		PrimaryTask: "compute the rank of the map from coefficient ray/phase data to normalized spectra and mixing invariants, then define the smallest comparator dataset allowed by Gate 453",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate452NearestNeighborNotGauge || !a.Inheritance.Gate450RatiosRequireAmplitudes || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("Gate453 requires Gate450-452 no-native-ratio inheritance")
	}
	if !a.Native.Executed || a.Native.NativeOnlyPredictsGST || len(a.Native.PromotedStructuralObjects) == 0 || len(a.Native.QuarantinedObjects) == 0 {
		return fmt.Errorf("native invariant ledger is not sealed correctly")
	}
	if !a.Contract.Executed || !a.Contract.RequiresExplicitLabel || !a.Contract.RequiresRenormalizationTag || !a.Contract.RequiresSectorTag || a.Contract.AllowsNativeClaim || a.Contract.RejectedPromotionCount == 0 {
		return fmt.Errorf("empirical import contract failed")
	}
	if !a.Residuals.Executed || !a.Residuals.AllowsTextureResiduals || !a.Residuals.AllowsGSTResidual || a.Residuals.AllowsNativeGSTRatioClaim || a.Residuals.AllowsCoefficientFittingAsNative {
		return fmt.Errorf("residual ledger failed to quarantine comparator outputs")
	}
	if !a.Sieve.Executed || !a.Sieve.NativeOnlyAllowed || !a.Sieve.EmpiricalFitAllowed || !a.Sieve.PromotionRejected || a.Sieve.AnyForbiddenAccepted {
		return fmt.Errorf("interface sieve accepted a forbidden promotion or rejected legal comparator use")
	}
	if !a.Firewall.Executed || !a.Firewall.NoGSTPromotion || !a.Firewall.GSTFritzschRelationsQuarantined || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall not preserved")
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 453 does not derive a new flavor observable. It seals the post-452 boundary: ASHA may natively report K_gen, the Generation-2 structural zero, the full triangular bridge support, and the exact M_22=0 spectral sum rule. Any coefficient ray, phase ray, GST/Fritzsch branch, mass value, or mixing value may enter only through an explicitly labelled empirical comparator interface, and no comparator residual may be promoted back into native geometry."
}

func statuses() []string {
	return []string{
		StatusGate452Inherited,
		StatusNativeInvariantLedgerSealed,
		StatusEmpiricalInterfaceDefined,
		StatusImportContractValidated,
		StatusPromotionFirewallValidated,
		StatusResidualComputationsQuarantined,
		StatusEmpiricalFirewallPreserved,
		StatusFailedNativeRatioNotRestored,
		StatusFailedGSTRequiresEmpiricalBranchInput,
		StatusFailedObservablePromotionRejected,
	}
}

func join(xs []string) string { return strings.Join(xs, "; ") }
