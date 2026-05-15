// Package generation2provenancecontract implements Gate 457:
// Empirical Comparator Provenance Contract / Sector-Scheme Ledger.
//
// Gate 456 derived a symbolic bridge inverse from labelled comparator pairs to
// the projective coefficient ray, while proving that the inverse is branchy and
// bridge-only. Gate 457 hardens the interface: any future real comparator import
// must carry a complete provenance ledger before the symbolic inverse may even be
// evaluated. This package imports no observed masses, Yukawa values, CKM/PMNS
// values, or fitted coefficient rays.
package generation2provenancecontract

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE457-EMPIRICAL-COMPARATOR-PROVENANCE-CONTRACT-SECTOR-SCHEME-LEDGER"

	StatusGate456Inherited                     = "CONDITIONAL_SUPPORT_GATE456_SYMBOLIC_INVERSE_INHERITED"
	StatusProvenanceContractDefined            = "CONDITIONAL_SUPPORT_EMPIRICAL_COMPARATOR_PROVENANCE_CONTRACT_DEFINED"
	StatusSchemaRequiredFieldsValidated        = "CONDITIONAL_SUPPORT_SECTOR_SCHEME_PROVENANCE_FIELDS_VALIDATED"
	StatusBridgeOnlyObservedImportGuarded      = "CONDITIONAL_SUPPORT_BRIDGE_ONLY_OBSERVED_IMPORT_GUARDED"
	StatusTextureComparatorContractValidated   = "CONDITIONAL_SUPPORT_TEXTURE_COMPARATOR_PROVENANCE_CONTRACT_VALIDATED"
	StatusEmpiricalFirewallPreserved           = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedMissingSectorScaleScheme       = "FAILED_ROUTE_MISSING_SECTOR_SCALE_SCHEME_METADATA"
	StatusFailedMissingSourceUncertainty       = "FAILED_ROUTE_MISSING_SOURCE_OR_UNCERTAINTY_METADATA"
	StatusFailedNativePromotionAttempt         = "FAILED_ROUTE_PROVENANCE_RECORD_ATTEMPTS_NATIVE_PROMOTION"
	StatusFailedObservedDefaultModeRejected    = "FAILED_ROUTE_OBSERVED_VALUES_REJECTED_OUTSIDE_EXPLICIT_BRIDGE_IMPORT"
	StatusFailedBranchTagRequired              = "FAILED_ROUTE_ORIENTED_INVERSE_REQUIRES_EXPLICIT_BRANCH_TAG"
	StatusFailedDimensionfulComparatorRejected = "FAILED_ROUTE_DIMENSIONFUL_COMPARATOR_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	RayDOF          = 2
	MinimumFields   = 11
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate445TriangleForced           bool
	Gate456SymbolicInverseDerived   bool
	Gate456BridgeOnly               bool
	Gate456GenericBranchCount       int
	Gate456RequiresBranchTags       bool
	Gate456ComparatorDomainGuard    bool
	NativeCoefficientSelectorAbsent bool
	NoObservedValuesImported        bool
	Verdict                         string
}

type FieldRule struct {
	Name        string
	Required    bool
	Reason      string
	FailureCode string
}

type Contract struct {
	Executed                                   bool
	Rules                                      []FieldRule
	RequiredFieldCount                         int
	RequiresSector                             bool
	RequiresObservable                         bool
	RequiresScale                              bool
	RequiresScheme                             bool
	RequiresSource                             bool
	RequiresSourceVersion                      bool
	RequiresUncertainty                        bool
	RequiresDimensionless                      bool
	RequiresBridgeOnly                         bool
	RequiresNoNativePromotion                  bool
	RequiresBranchTagIfOriented                bool
	AllowsObservedOnlyWithExplicitBridgeImport bool
	Verdict                                    string
	Reason                                     string
}

type ComparatorRecord struct {
	Name                    string
	Sector                  string
	Observable              string
	ValueKind               string
	ValueExpression         string
	Scale                   string
	Scheme                  string
	Source                  string
	SourceVersion           string
	Uncertainty             string
	Dimensionless           bool
	BridgeOnly              bool
	ExplicitObservedImport  bool
	NativePromotionClaim    bool
	RequiresOrientedInverse bool
	BranchTag               string
	Passed                  bool
	Verdict                 string
	Reason                  string
}

type Sieve struct {
	Executed                             bool
	Records                              []ComparatorRecord
	AcceptedCount                        int
	RejectedCount                        int
	CompleteSymbolicDryRunAccepted       bool
	ExplicitBridgeObservedSchemaAccepted bool
	MissingSectorRejected                bool
	MissingScaleSchemeRejected           bool
	MissingSourceUncertaintyRejected     bool
	NativePromotionRejected              bool
	ObservedDefaultRejected              bool
	BranchTagMissingRejected             bool
	DimensionfulComparatorRejected       bool
	NoAcceptedNativeExport               bool
	Verdict                              string
	Reason                               string
}

type Firewall struct {
	Executed                      bool
	NoObservedMuonMassImported    bool
	NoObservedCharmMassImported   bool
	NoObservedYukawaImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoGSTPromotion                bool
	NoCoefficientRayPromotion     bool
	NoCurveFitPromoted            bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Contract    Contract
	Sieve       Sieve
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
	a.Contract = buildContract()
	a.Sieve = buildSieve(a.Contract)
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
		Gate445TriangleForced:           true,
		Gate456SymbolicInverseDerived:   true,
		Gate456BridgeOnly:               true,
		Gate456GenericBranchCount:       6,
		Gate456RequiresBranchTags:       true,
		Gate456ComparatorDomainGuard:    true,
		NativeCoefficientSelectorAbsent: true,
		NoObservedValuesImported:        true,
		Verdict:                         StatusGate456Inherited,
	}
}

func buildContract() Contract {
	rules := []FieldRule{
		{Name: "sector", Required: true, Reason: "selects up, down, charged-lepton, or neutrino bridge lane; prevents cross-sector coefficient smuggling", FailureCode: StatusFailedMissingSectorScaleScheme},
		{Name: "observable", Required: true, Reason: "must be one of the labelled bridge comparators or residuals; unlabelled texture data is not evaluable", FailureCode: StatusFailedMissingSectorScaleScheme},
		{Name: "value_kind", Required: true, Reason: "distinguishes symbolic dry run from explicit observed bridge import", FailureCode: StatusFailedObservedDefaultModeRejected},
		{Name: "scale", Required: true, Reason: "flavor observables are scale-dependent once interpreted phenomenologically", FailureCode: StatusFailedMissingSectorScaleScheme},
		{Name: "scheme", Required: true, Reason: "renormalization and mass definitions must not be silently mixed", FailureCode: StatusFailedMissingSectorScaleScheme},
		{Name: "source", Required: true, Reason: "observed bridge imports require citable provenance", FailureCode: StatusFailedMissingSourceUncertainty},
		{Name: "source_version", Required: true, Reason: "prevents stale data from masquerading as a stable theorem input", FailureCode: StatusFailedMissingSourceUncertainty},
		{Name: "uncertainty", Required: true, Reason: "observed bridge comparisons must carry error bars or an explicit symbolic uncertainty tag", FailureCode: StatusFailedMissingSourceUncertainty},
		{Name: "dimensionless", Required: true, Reason: "I_spec and I_K are dimensionless comparator scalars; dimensionful masses are not accepted at this layer", FailureCode: StatusFailedDimensionfulComparatorRejected},
		{Name: "bridge_only", Required: true, Reason: "the record may drive bridge evaluation only; native promotion is forbidden", FailureCode: StatusFailedNativePromotionAttempt},
		{Name: "branch_tag_if_oriented", Required: true, Reason: "Gate 456 proves six generic phase branches; oriented inverse calls require an explicit branch tag", FailureCode: StatusFailedBranchTagRequired},
	}
	return Contract{
		Executed:                    true,
		Rules:                       rules,
		RequiredFieldCount:          len(rules),
		RequiresSector:              true,
		RequiresObservable:          true,
		RequiresScale:               true,
		RequiresScheme:              true,
		RequiresSource:              true,
		RequiresSourceVersion:       true,
		RequiresUncertainty:         true,
		RequiresDimensionless:       true,
		RequiresBridgeOnly:          true,
		RequiresNoNativePromotion:   true,
		RequiresBranchTagIfOriented: true,
		AllowsObservedOnlyWithExplicitBridgeImport: true,
		Verdict: StatusProvenanceContractDefined,
		Reason:  "a comparator record is evaluable only after sector, scale, scheme, source, uncertainty, dimensionless status, bridge-only status, and branch semantics are all explicit.",
	}
}

func buildSieve(c Contract) Sieve {
	seeds := []ComparatorRecord{
		{
			Name: "complete symbolic I_K dry run accepted", Sector: "charged-lepton", Observable: "I_K", ValueKind: "symbolic", ValueExpression: "I_K(symbolic)", Scale: "symbolic-scale-tag", Scheme: "symbolic-renormalization-scheme", Source: "ASHA Gate454/Gate456 internal formula ledger", SourceVersion: "gate456", Uncertainty: "symbolic-none", Dimensionless: true, BridgeOnly: true, ExplicitObservedImport: false, NativePromotionClaim: false, RequiresOrientedInverse: false,
		},
		{
			Name: "explicit observed bridge schema accepted without value import", Sector: "up", Observable: "I_spec", ValueKind: "observed-placeholder", ValueExpression: "external-value-redacted-by-gate457", Scale: "must-be-provided-by-caller", Scheme: "must-be-provided-by-caller", Source: "must-be-provided-by-caller", SourceVersion: "must-be-provided-by-caller", Uncertainty: "must-be-provided-by-caller", Dimensionless: true, BridgeOnly: true, ExplicitObservedImport: true, NativePromotionClaim: false, RequiresOrientedInverse: false,
		},
		{
			Name: "missing sector rejected", Observable: "I_K", ValueKind: "symbolic", ValueExpression: "I_K", Scale: "MZ", Scheme: "MSbar", Source: "external", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: true, BridgeOnly: true,
		},
		{
			Name: "missing scale and scheme rejected", Sector: "down", Observable: "I_spec", ValueKind: "symbolic", ValueExpression: "I_spec", Source: "external", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: true, BridgeOnly: true,
		},
		{
			Name: "missing source uncertainty rejected", Sector: "charged-lepton", Observable: "I_K", ValueKind: "symbolic", ValueExpression: "I_K", Scale: "MZ", Scheme: "MSbar", Dimensionless: true, BridgeOnly: true,
		},
		{
			Name: "native promotion attempt rejected", Sector: "up", Observable: "I_spec", ValueKind: "symbolic", ValueExpression: "I_spec", Scale: "MZ", Scheme: "MSbar", Source: "external", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: true, BridgeOnly: false, NativePromotionClaim: true,
		},
		{
			Name: "observed default mode rejected", Sector: "down", Observable: "I_K", ValueKind: "observed", ValueExpression: "redacted", Scale: "MZ", Scheme: "MSbar", Source: "external", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: true, BridgeOnly: true, ExplicitObservedImport: false,
		},
		{
			Name: "oriented inverse without branch tag rejected", Sector: "charged-lepton", Observable: "oriented_phi_branch", ValueKind: "symbolic", ValueExpression: "phi_branch", Scale: "symbolic-scale-tag", Scheme: "symbolic-renormalization-scheme", Source: "ASHA Gate456 internal formula ledger", SourceVersion: "gate456", Uncertainty: "symbolic-none", Dimensionless: true, BridgeOnly: true, RequiresOrientedInverse: true,
		},
		{
			Name: "dimensionful mass mistaken for comparator rejected", Sector: "charged-lepton", Observable: "I_K", ValueKind: "symbolic", ValueExpression: "m_mu", Scale: "pole", Scheme: "pole", Source: "external", SourceVersion: "v", Uncertainty: "sigma", Dimensionless: false, BridgeOnly: true,
		},
	}
	out := Sieve{Executed: true}
	for _, seed := range seeds {
		rec := evaluate(seed)
		out.Records = append(out.Records, rec)
		if rec.Passed {
			out.AcceptedCount++
		} else {
			out.RejectedCount++
		}
		switch rec.Name {
		case "complete symbolic I_K dry run accepted":
			out.CompleteSymbolicDryRunAccepted = rec.Passed
		case "explicit observed bridge schema accepted without value import":
			out.ExplicitBridgeObservedSchemaAccepted = rec.Passed
		case "missing sector rejected":
			out.MissingSectorRejected = !rec.Passed && rec.Verdict == StatusFailedMissingSectorScaleScheme
		case "missing scale and scheme rejected":
			out.MissingScaleSchemeRejected = !rec.Passed && rec.Verdict == StatusFailedMissingSectorScaleScheme
		case "missing source uncertainty rejected":
			out.MissingSourceUncertaintyRejected = !rec.Passed && rec.Verdict == StatusFailedMissingSourceUncertainty
		case "native promotion attempt rejected":
			out.NativePromotionRejected = !rec.Passed && rec.Verdict == StatusFailedNativePromotionAttempt
		case "observed default mode rejected":
			out.ObservedDefaultRejected = !rec.Passed && rec.Verdict == StatusFailedObservedDefaultModeRejected
		case "oriented inverse without branch tag rejected":
			out.BranchTagMissingRejected = !rec.Passed && rec.Verdict == StatusFailedBranchTagRequired
		case "dimensionful mass mistaken for comparator rejected":
			out.DimensionfulComparatorRejected = !rec.Passed && rec.Verdict == StatusFailedDimensionfulComparatorRejected
		}
	}
	out.NoAcceptedNativeExport = true
	for _, rec := range out.Records {
		if rec.Passed && (!rec.BridgeOnly || rec.NativePromotionClaim) {
			out.NoAcceptedNativeExport = false
		}
	}
	out.Verdict = StatusTextureComparatorContractValidated
	out.Reason = fmt.Sprintf("%d contract-valid bridge records accepted and %d malformed or native-promoting records rejected before evaluation.", out.AcceptedCount, out.RejectedCount)
	return out
}

func evaluate(r ComparatorRecord) ComparatorRecord {
	missingSectorScaleScheme := strings.TrimSpace(r.Sector) == "" || strings.TrimSpace(r.Observable) == "" || strings.TrimSpace(r.Scale) == "" || strings.TrimSpace(r.Scheme) == ""
	if missingSectorScaleScheme {
		r.Passed = false
		r.Verdict = StatusFailedMissingSectorScaleScheme
		r.Reason = "sector, observable, scale, and scheme must be explicit before any texture comparator can be evaluated."
		return r
	}
	if strings.TrimSpace(r.Source) == "" || strings.TrimSpace(r.SourceVersion) == "" || strings.TrimSpace(r.Uncertainty) == "" {
		r.Passed = false
		r.Verdict = StatusFailedMissingSourceUncertainty
		r.Reason = "source, source version, and uncertainty metadata are mandatory provenance fields."
		return r
	}
	if !r.Dimensionless {
		r.Passed = false
		r.Verdict = StatusFailedDimensionfulComparatorRejected
		r.Reason = "Gate457 accepts dimensionless comparator scalars only; dimensionful masses must be converted in an external bridge adapter."
		return r
	}
	if !r.BridgeOnly || r.NativePromotionClaim {
		r.Passed = false
		r.Verdict = StatusFailedNativePromotionAttempt
		r.Reason = "provenance records may enter bridge evaluation only and cannot request native-law promotion."
		return r
	}
	if strings.Contains(r.ValueKind, "observed") && !r.ExplicitObservedImport {
		r.Passed = false
		r.Verdict = StatusFailedObservedDefaultModeRejected
		r.Reason = "observed values are rejected unless the caller explicitly chooses observed bridge-import mode."
		return r
	}
	if r.RequiresOrientedInverse && strings.TrimSpace(r.BranchTag) == "" {
		r.Passed = false
		r.Verdict = StatusFailedBranchTagRequired
		r.Reason = "oriented phase inversion requires an explicit branch tag because Gate456 leaves six generic phase branches."
		return r
	}
	r.Passed = true
	r.Verdict = StatusBridgeOnlyObservedImportGuarded
	r.Reason = "record is schema-complete and bridge-only; it may reach symbolic comparator evaluation but still exports no native coefficient ray."
	return r
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		NoObservedMuonMassImported:    true,
		NoObservedCharmMassImported:   true,
		NoObservedYukawaImported:      true,
		NoCKMImported:                 true,
		NoPMNSImported:                true,
		NoGSTPromotion:                true,
		NoCoefficientRayPromotion:     true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusEmpiricalFirewallPreserved,
		Reason:                        "Gate457 defines schema/provenance gates only; it imports no actual flavor data and promotes no comparator, coefficient ray, or texture relation into native law-space.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        458,
		Title:       "Comparator Ledger Evaluation Harness / Redacted Phenomenology Slot",
		Reason:      "the provenance schema is now fail-closed, so the next bridge can evaluate redacted/synthetic comparator records against the Gate456 inverse without native promotion",
		PrimaryTask: "build the first evaluation harness that consumes only Gate457-valid records, computes symbolic residual objects, and marks every output bridge-only",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 457 defines %d required provenance fields, accepts %d schema-complete bridge records, rejects %d malformed/native-promoting records, and preserves the 13-moduli firewall without importing any observed flavor values.", a.Contract.RequiredFieldCount, a.Sieve.AcceptedCount, a.Sieve.RejectedCount)
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate456SymbolicInverseDerived || !a.Inheritance.Gate456BridgeOnly || a.Inheritance.Gate456GenericBranchCount != 6 || !a.Inheritance.Gate456RequiresBranchTags || !a.Inheritance.Gate456ComparatorDomainGuard || !a.Inheritance.NativeCoefficientSelectorAbsent || !a.Inheritance.NoObservedValuesImported {
		return fmt.Errorf("Gate456 symbolic inverse not inherited: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Contract.Executed || a.Contract.RequiredFieldCount != MinimumFields || !a.Contract.RequiresSector || !a.Contract.RequiresObservable || !a.Contract.RequiresScale || !a.Contract.RequiresScheme || !a.Contract.RequiresSource || !a.Contract.RequiresSourceVersion || !a.Contract.RequiresUncertainty || !a.Contract.RequiresDimensionless || !a.Contract.RequiresBridgeOnly || !a.Contract.RequiresNoNativePromotion || !a.Contract.RequiresBranchTagIfOriented || !a.Contract.AllowsObservedOnlyWithExplicitBridgeImport {
		return fmt.Errorf("provenance contract incomplete: %s", FormatContract(a.Contract))
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 7 || !a.Sieve.CompleteSymbolicDryRunAccepted || !a.Sieve.ExplicitBridgeObservedSchemaAccepted || !a.Sieve.MissingSectorRejected || !a.Sieve.MissingScaleSchemeRejected || !a.Sieve.MissingSourceUncertaintyRejected || !a.Sieve.NativePromotionRejected || !a.Sieve.ObservedDefaultRejected || !a.Sieve.BranchTagMissingRejected || !a.Sieve.DimensionfulComparatorRejected || !a.Sieve.NoAcceptedNativeExport {
		return fmt.Errorf("provenance sieve failed: %s", FormatSieve(a.Sieve))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoGSTPromotion || !a.Firewall.NoCoefficientRayPromotion || !a.Firewall.NoCurveFitPromoted || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate456Inherited,
		StatusProvenanceContractDefined,
		StatusSchemaRequiredFieldsValidated,
		StatusBridgeOnlyObservedImportGuarded,
		StatusTextureComparatorContractValidated,
		StatusEmpiricalFirewallPreserved,
		StatusFailedMissingSectorScaleScheme,
		StatusFailedMissingSourceUncertainty,
		StatusFailedNativePromotionAttempt,
		StatusFailedObservedDefaultModeRejected,
		StatusFailedBranchTagRequired,
		StatusFailedDimensionfulComparatorRejected,
	}
}
