// Package generation2empiricaladapter implements Gate 455:
// Empirical Texture Adapter Stub / Dry-Run Firewall Test.
//
// Gate 454 proved that coefficient-ray observability is a bridge-rank matter:
// spectrum-only comparators have rank one, two labelled scalar comparators can
// identify a coefficient ray locally, and CP orientation still needs an
// explicit branch tag. Gate 455 turns that result into a machine-checkable
// adapter firewall. The default adapter imports no observed flavor data; it
// accepts only native ledgers and labelled symbolic dry-run bridge requests,
// while every attempt to promote spectrum fits, GST/Fritzsch relations, CKM,
// PMNS, Yukawa, or phase selectors to native ASHA law fails closed.
package generation2empiricaladapter

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE455-EMPIRICAL-TEXTURE-ADAPTER-DRY-RUN-FIREWALL-TEST"

	StatusGate454Inherited                          = "CONDITIONAL_SUPPORT_GATE454_RANK_PROTOCOL_INHERITED"
	StatusAdapterSchemaDefined                      = "CONDITIONAL_SUPPORT_EMPIRICAL_TEXTURE_ADAPTER_SCHEMA_DEFINED"
	StatusDryRunFirewallValidated                   = "CONDITIONAL_SUPPORT_EMPIRICAL_TEXTURE_ADAPTER_FIREWALL_VALIDATED"
	StatusLabelledLocalRayDryRunAllowed             = "CONDITIONAL_SUPPORT_LABELLED_LOCAL_RAY_DRY_RUN_ALLOWED"
	StatusLabelledOrientedRayDryRunAllowed          = "CONDITIONAL_SUPPORT_LABELLED_ORIENTED_RAY_DRY_RUN_ALLOWED"
	StatusNoObservedValuesImportedDefault           = "CONDITIONAL_SUPPORT_NO_OBSERVED_VALUES_IMPORTED_BY_DEFAULT"
	StatusBridgeOnlyExportsValidated                = "CONDITIONAL_SUPPORT_BRIDGE_ONLY_EXPORTS_VALIDATED"
	StatusEmpiricalFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedSpectrumOnlyNativePromotionRejected = "FAILED_ROUTE_ADAPTER_REJECTS_SPECTRUM_ONLY_NATIVE_PROMOTION"
	StatusFailedMissingMetadataRejected             = "FAILED_ROUTE_ADAPTER_REJECTS_MISSING_METADATA"
	StatusFailedGSTNativePromotionRejected          = "FAILED_ROUTE_ADAPTER_REJECTS_GST_NATIVE_PROMOTION"
	StatusFailedCKMPMNSNativeSelectorRejected       = "FAILED_ROUTE_ADAPTER_REJECTS_CKM_PMNS_NATIVE_SELECTOR"
	StatusFailedObservedValuesRejectedByDefault     = "FAILED_ROUTE_ADAPTER_REJECTS_OBSERVED_VALUES_IN_DRY_RUN_MODE"
	StatusFailedNativeCoefficientExportAbsent       = "FAILED_ROUTE_NATIVE_COEFFICIENT_EXPORT_ABSENT"
)

const (
	NativeFlavorDim  = 13
	KXYCoeffDim      = 9
	ProjectiveRayDOF = 2
	MinLocalScalars  = 2
	MinOrientedTags  = 3
)

const (
	ValueModeNone     = "none"
	ValueModeDummy    = "symbolic-dummy"
	ValueModeObserved = "observed"
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate450TextureZeroSumRule        bool
	Gate451FullTrianglePreserved     bool
	Gate452NearestNeighborNotGauge   bool
	Gate453EmpiricalInterfaceDefined bool
	Gate454ProjectiveRayDOF          int
	Gate454SpectrumOnlyRank          int
	Gate454MinimumLocalScalars       int
	Gate454CPBranchTagRequired       bool
	Gate454NativeSelectorAbsent      bool
	NoEmpiricalInputsImported        bool
	Verdict                          string
}

type AdapterSchema struct {
	Executed                       bool
	Name                           string
	DefaultValueMode               string
	RequiredLabels                 []string
	AllowedOperations              []string
	RejectedOperations             []string
	AllowsObservedValuesByDefault  bool
	AllowsNativeCoefficientExport  bool
	AllowsGSTAsNativeLaw           bool
	AllowsCKMPMNSAsNativeSelectors bool
	Verdict                        string
	Reason                         string
}

type AdapterRequest struct {
	Name                     string
	Operation                string
	ValueMode                string
	ComparatorCount          int
	HasSectorTag             bool
	HasRenormalizationScale  bool
	HasRenormalizationScheme bool
	HasBridgeLabel           bool
	HasCPBranchTag           bool
	UsesSpectrum             bool
	UsesMasses               bool
	UsesYukawa               bool
	UsesCKM                  bool
	UsesPMNS                 bool
	ClaimsLocalRay           bool
	ClaimsOrientedRay        bool
	ClaimsNativeCoefficient  bool
	ClaimsGSTNative          bool
	ClaimsPhaseNative        bool
	Allowed                  bool
	Classification           string
	Verdict                  string
	Reason                   string
}

type AdapterSieve struct {
	Executed                            bool
	Requests                            []AdapterRequest
	AllowedCount                        int
	RejectedCount                       int
	NativeLedgerAllowed                 bool
	LocalRayDryRunAllowed               bool
	OrientedRayDryRunAllowed            bool
	SpectrumOnlyNativePromotionRejected bool
	MissingMetadataRejected             bool
	GSTNativePromotionRejected          bool
	CKMPMNSNativeSelectorRejected       bool
	ObservedValuesRejectedByDefault     bool
	NativeCoefficientExportRejected     bool
	AnyForbiddenAccepted                bool
	Verdict                             string
	Reason                              string
}

type DryRunExport struct {
	Executed                 bool
	ActualObservedValueCount int
	DummyComparatorCount     int
	NativeExportCount        int
	BridgeExportCount        int
	NativePromotionBlocked   bool
	SchemaFailuresFailClosed bool
	Verdict                  string
	Reason                   string
}

type Firewall struct {
	Executed                       bool
	NoObservedMuonMassImported     bool
	NoObservedCharmMassImported    bool
	NoObservedYukawaImported       bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoCurveFitPromoted             bool
	NoGSTPromotion                 bool
	NoNativeCoefficientRayValue    bool
	KGenStillForced                bool
	XTriangleStillForced           bool
	TextureZeroSumRuleStillBridge  bool
	YPhaseStillQuarantined         bool
	SectorCoefficientsStillSealed  bool
	CPOrientationStillBranchTagged bool
	NativeFlavorDimAfter           int
	KXYCoeffDimAfter               int
	Verdict                        string
	Reason                         string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      AdapterSchema
	Sieve       AdapterSieve
	Export      DryRunExport
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
	a.Schema = buildSchema()
	a.Sieve = buildSieve(a.Schema)
	a.Export = buildExport(a.Sieve)
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
		Executed:                         true,
		Gate444KGenForced:                true,
		Gate445TriangleForced:            true,
		Gate450TextureZeroSumRule:        true,
		Gate451FullTrianglePreserved:     true,
		Gate452NearestNeighborNotGauge:   true,
		Gate453EmpiricalInterfaceDefined: true,
		Gate454ProjectiveRayDOF:          ProjectiveRayDOF,
		Gate454SpectrumOnlyRank:          1,
		Gate454MinimumLocalScalars:       MinLocalScalars,
		Gate454CPBranchTagRequired:       true,
		Gate454NativeSelectorAbsent:      true,
		NoEmpiricalInputsImported:        true,
		Verdict:                          StatusGate454Inherited,
	}
}

func buildSchema() AdapterSchema {
	return AdapterSchema{
		Executed:         true,
		Name:             "generation-2 texture-zero empirical adapter dry-run schema",
		DefaultValueMode: ValueModeDummy,
		RequiredLabels: []string{
			"bridge-only provenance label",
			"charged sector tag",
			"renormalization scale tag",
			"renormalization scheme tag",
			"CP branch tag for oriented phase claims",
		},
		AllowedOperations: []string{
			"native structural ledger with no observed values",
			"labelled symbolic spectrum residual",
			"labelled symbolic local coefficient-ray dry run using I_spec and I_K",
			"labelled symbolic oriented coefficient-ray dry run using I_spec, I_K, and explicit CP branch tag",
		},
		RejectedOperations: []string{
			"spectrum-only native coefficient claim",
			"GST/Fritzsch relation promoted to native law",
			"missing sector/scale/scheme/provenance metadata",
			"CKM or PMNS value used as native phase selector",
			"observed flavor values imported in default dry-run mode",
		},
		AllowsObservedValuesByDefault:  false,
		AllowsNativeCoefficientExport:  false,
		AllowsGSTAsNativeLaw:           false,
		AllowsCKMPMNSAsNativeSelectors: false,
		Verdict:                        StatusAdapterSchemaDefined,
		Reason:                         "the adapter is a dry-run bridge schema: it can validate labelled symbolic comparator paths, but cannot export native coefficient values or silently consume observed flavor data.",
	}
}

func buildSieve(schema AdapterSchema) AdapterSieve {
	seeds := []AdapterRequest{
		{
			Name:            "native structural family ledger",
			Operation:       "native-ledger",
			ValueMode:       ValueModeNone,
			ComparatorCount: 0,
			Classification:  "candidate",
		},
		{
			Name:                     "symbolic spectrum residual only",
			Operation:                "spectrum-residual",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          1,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			UsesSpectrum:             true,
			Classification:           "candidate",
		},
		{
			Name:                     "symbolic local coefficient-ray dry run",
			Operation:                "local-ray-comparator",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          2,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			UsesSpectrum:             true,
			ClaimsLocalRay:           true,
			Classification:           "candidate",
		},
		{
			Name:                     "symbolic oriented coefficient-ray dry run",
			Operation:                "oriented-ray-comparator",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          3,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			HasCPBranchTag:           true,
			UsesSpectrum:             true,
			ClaimsLocalRay:           true,
			ClaimsOrientedRay:        true,
			Classification:           "candidate",
		},
		{
			Name:                     "spectrum-only native coefficient claim",
			Operation:                "native-coefficient-from-spectrum",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          1,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			UsesSpectrum:             true,
			UsesMasses:               true,
			ClaimsNativeCoefficient:  true,
			Classification:           "candidate",
		},
		{
			Name:            "local-ray comparator missing metadata",
			Operation:       "local-ray-comparator",
			ValueMode:       ValueModeDummy,
			ComparatorCount: 2,
			HasSectorTag:    true,
			HasBridgeLabel:  true,
			UsesSpectrum:    true,
			ClaimsLocalRay:  true,
			Classification:  "candidate",
		},
		{
			Name:                     "GST relation as ASHA law",
			Operation:                "gst-native-law",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          2,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			ClaimsGSTNative:          true,
			Classification:           "candidate",
		},
		{
			Name:                     "CKM/PMNS phase as native selector",
			Operation:                "phase-selector-native",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          3,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			HasCPBranchTag:           true,
			UsesCKM:                  true,
			UsesPMNS:                 true,
			ClaimsPhaseNative:        true,
			Classification:           "candidate",
		},
		{
			Name:                     "observed-value dry-run import",
			Operation:                "observed-local-ray-import",
			ValueMode:                ValueModeObserved,
			ComparatorCount:          2,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			UsesSpectrum:             true,
			UsesMasses:               true,
			ClaimsLocalRay:           true,
			Classification:           "candidate",
		},
	}
	requests := make([]AdapterRequest, len(seeds))
	s := AdapterSieve{Executed: true, Requests: requests}
	for i, r := range seeds {
		requests[i] = classify(schema, r)
		cr := requests[i]
		if cr.Allowed {
			s.AllowedCount++
		} else {
			s.RejectedCount++
		}
		switch cr.Classification {
		case "allowed-native-ledger":
			s.NativeLedgerAllowed = true
		case "allowed-local-ray-dry-run":
			s.LocalRayDryRunAllowed = true
		case "allowed-oriented-ray-dry-run":
			s.OrientedRayDryRunAllowed = true
		case "rejected-spectrum-only-native-promotion":
			s.SpectrumOnlyNativePromotionRejected = true
			s.NativeCoefficientExportRejected = true
		case "rejected-missing-metadata":
			s.MissingMetadataRejected = true
		case "rejected-gst-native-promotion":
			s.GSTNativePromotionRejected = true
		case "rejected-ckm-pmns-native-selector":
			s.CKMPMNSNativeSelectorRejected = true
		case "rejected-observed-values-dry-run":
			s.ObservedValuesRejectedByDefault = true
		}
		if forbidden(cr) && cr.Allowed {
			s.AnyForbiddenAccepted = true
		}
	}
	s.Verdict = StatusDryRunFirewallValidated
	s.Reason = "the dry-run adapter accepts native ledgers and fully-labelled symbolic bridge comparators, while spectrum-only native promotion, missing metadata, GST promotion, CKM/PMNS native selectors, and observed-value imports fail closed."
	return s
}

func classify(schema AdapterSchema, r AdapterRequest) AdapterRequest {
	hasMetadata := r.HasSectorTag && r.HasRenormalizationScale && r.HasRenormalizationScheme && r.HasBridgeLabel
	if r.ValueMode == ValueModeObserved && !schema.AllowsObservedValuesByDefault {
		r.Allowed = false
		r.Classification = "rejected-observed-values-dry-run"
		r.Verdict = StatusFailedObservedValuesRejectedByDefault
		r.Reason = "default Gate455 mode is a dry-run adapter; observed flavor values require a later explicit empirical run mode and cannot enter this registry theorem."
		return r
	}
	if r.ClaimsGSTNative {
		r.Allowed = false
		r.Classification = "rejected-gst-native-promotion"
		r.Verdict = StatusFailedGSTNativePromotionRejected
		r.Reason = "GST/Fritzsch relations were quarantined by Gates 450-452 and may not be relabelled as ASHA law by the adapter."
		return r
	}
	if r.ClaimsPhaseNative || ((r.UsesCKM || r.UsesPMNS) && (r.ClaimsNativeCoefficient || r.ClaimsOrientedRay || r.ClaimsPhaseNative)) {
		r.Allowed = false
		r.Classification = "rejected-ckm-pmns-native-selector"
		r.Verdict = StatusFailedCKMPMNSNativeSelectorRejected
		r.Reason = "CKM/PMNS or CP-phase data can be compared only as labelled empirical bridge information; it cannot select the native phase ray."
		return r
	}
	if r.ClaimsNativeCoefficient {
		r.Allowed = false
		r.Classification = "rejected-spectrum-only-native-promotion"
		r.Verdict = StatusFailedSpectrumOnlyNativePromotionRejected
		r.Reason = "Gate454 proves spectrum-only rank one and Gate455 forbids native coefficient export, so this path is rejected."
		return r
	}
	if r.Operation == "native-ledger" && r.ValueMode == ValueModeNone && r.ComparatorCount == 0 {
		r.Allowed = true
		r.Classification = "allowed-native-ledger"
		r.Verdict = StatusBridgeOnlyExportsValidated
		r.Reason = "the adapter may render the native structural ledger because it imports no observed values and exports no coefficient ray."
		return r
	}
	if !hasMetadata {
		r.Allowed = false
		r.Classification = "rejected-missing-metadata"
		r.Verdict = StatusFailedMissingMetadataRejected
		r.Reason = "bridge comparators must carry sector, scale, scheme, and bridge-only provenance labels."
		return r
	}
	if r.ClaimsOrientedRay {
		if r.ComparatorCount >= MinOrientedTags && r.HasCPBranchTag {
			r.Allowed = true
			r.Classification = "allowed-oriented-ray-dry-run"
			r.Verdict = StatusLabelledOrientedRayDryRunAllowed
			r.Reason = "three symbolic comparator tags including an explicit CP branch tag satisfy the Gate454 oriented-ray protocol, but only as bridge metadata."
			return r
		}
		r.Allowed = false
		r.Classification = "rejected-missing-metadata"
		r.Verdict = StatusFailedMissingMetadataRejected
		r.Reason = "oriented phase claims require the CP branch tag in addition to local-ray metadata."
		return r
	}
	if r.ClaimsLocalRay {
		if r.ComparatorCount >= MinLocalScalars {
			r.Allowed = true
			r.Classification = "allowed-local-ray-dry-run"
			r.Verdict = StatusLabelledLocalRayDryRunAllowed
			r.Reason = "two labelled symbolic scalar comparators satisfy local Gate454 rank, with no native promotion."
			return r
		}
		r.Allowed = false
		r.Classification = "rejected-spectrum-only-native-promotion"
		r.Verdict = StatusFailedSpectrumOnlyNativePromotionRejected
		r.Reason = "a single spectrum comparator cannot identify the coefficient ray."
		return r
	}
	if r.Operation == "spectrum-residual" && r.ComparatorCount == 1 {
		r.Allowed = true
		r.Classification = "allowed-spectrum-residual-only"
		r.Verdict = StatusBridgeOnlyExportsValidated
		r.Reason = "a spectrum-only residual is allowed as a labelled bridge comparator, but it explicitly carries no coefficient-ray identification claim."
		return r
	}
	r.Allowed = false
	r.Classification = "rejected-unclassified"
	r.Verdict = StatusFailedNativeCoefficientExportAbsent
	r.Reason = "the request does not match any allowed dry-run bridge operation."
	return r
}

func forbidden(r AdapterRequest) bool {
	return r.ClaimsNativeCoefficient || r.ClaimsGSTNative || r.ClaimsPhaseNative || r.ValueMode == ValueModeObserved || (!r.HasRenormalizationScale && r.Operation != "native-ledger") || (!r.HasRenormalizationScheme && r.Operation != "native-ledger")
}

func buildExport(s AdapterSieve) DryRunExport {
	dummy := 0
	bridge := 0
	for _, r := range s.Requests {
		if r.Allowed && r.ValueMode == ValueModeDummy {
			dummy += r.ComparatorCount
			bridge++
		}
	}
	return DryRunExport{
		Executed:                 true,
		ActualObservedValueCount: 0,
		DummyComparatorCount:     dummy,
		NativeExportCount:        0,
		BridgeExportCount:        bridge,
		NativePromotionBlocked:   s.NativeCoefficientExportRejected && s.GSTNativePromotionRejected && s.CKMPMNSNativeSelectorRejected,
		SchemaFailuresFailClosed: s.MissingMetadataRejected && s.ObservedValuesRejectedByDefault && !s.AnyForbiddenAccepted,
		Verdict:                  StatusNoObservedValuesImportedDefault,
		Reason:                   "the default adapter run exports only bridge-labelled dry-run validations and imports zero observed numerical flavor values.",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                       true,
		NoObservedMuonMassImported:     a.Export.ActualObservedValueCount == 0,
		NoObservedCharmMassImported:    a.Export.ActualObservedValueCount == 0,
		NoObservedYukawaImported:       true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoCurveFitPromoted:             a.Export.NativeExportCount == 0,
		NoGSTPromotion:                 a.Sieve.GSTNativePromotionRejected,
		NoNativeCoefficientRayValue:    a.Sieve.NativeCoefficientExportRejected && a.Export.NativeExportCount == 0,
		KGenStillForced:                a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:           a.Inheritance.Gate445TriangleForced,
		TextureZeroSumRuleStillBridge:  a.Inheritance.Gate450TextureZeroSumRule,
		YPhaseStillQuarantined:         a.Inheritance.Gate454CPBranchTagRequired,
		SectorCoefficientsStillSealed:  a.Inheritance.Gate454NativeSelectorAbsent,
		CPOrientationStillBranchTagged: a.Inheritance.Gate454CPBranchTagRequired,
		NativeFlavorDimAfter:           NativeFlavorDim,
		KXYCoeffDimAfter:               KXYCoeffDim,
		Verdict:                        StatusEmpiricalFirewallPreserved,
		Reason:                         "Gate455 validates the adapter firewall only; it does not consume or promote observed flavor data, CKM/PMNS values, Yukawas, or coefficient-ray values.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        456,
		Title:       "Symbolic Coefficient-Ray Inversion / Branch-Caustic Map",
		Reason:      "after the dry-run adapter is fail-closed, the next native-safe task is to derive the exact symbolic inverse map from labelled comparators to the bridge ray and mark caustics/branch degeneracies",
		PrimaryTask: "derive alpha from I_K, derive cos(3 phi) from I_spec and alpha, identify sin(3 phi)=0 caustics, and keep all values bridge-labelled",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 455 validates the empirical texture adapter as a dry-run firewall: %d requests are accepted only as native ledgers or bridge-labelled symbolic comparator paths, while %d requests are rejected, including spectrum-only native promotion, missing metadata, GST promotion, CKM/PMNS native selection, and observed-value imports.", a.Sieve.AllowedCount, a.Sieve.RejectedCount)
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || a.Inheritance.Gate454ProjectiveRayDOF != ProjectiveRayDOF || a.Inheritance.Gate454SpectrumOnlyRank != 1 || a.Inheritance.Gate454MinimumLocalScalars != MinLocalScalars || !a.Inheritance.Gate454CPBranchTagRequired || !a.Inheritance.Gate454NativeSelectorAbsent || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("Gate454 rank protocol not inherited: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Schema.Executed || a.Schema.AllowsObservedValuesByDefault || a.Schema.AllowsNativeCoefficientExport || a.Schema.AllowsGSTAsNativeLaw || a.Schema.AllowsCKMPMNSAsNativeSelectors || len(a.Schema.RequiredLabels) < 4 {
		return fmt.Errorf("adapter schema leaks forbidden path: %s", FormatSchema(a.Schema))
	}
	if !a.Sieve.Executed || a.Sieve.AllowedCount != 4 || a.Sieve.RejectedCount != 5 || !a.Sieve.NativeLedgerAllowed || !a.Sieve.LocalRayDryRunAllowed || !a.Sieve.OrientedRayDryRunAllowed || !a.Sieve.SpectrumOnlyNativePromotionRejected || !a.Sieve.MissingMetadataRejected || !a.Sieve.GSTNativePromotionRejected || !a.Sieve.CKMPMNSNativeSelectorRejected || !a.Sieve.ObservedValuesRejectedByDefault || !a.Sieve.NativeCoefficientExportRejected || a.Sieve.AnyForbiddenAccepted {
		return fmt.Errorf("adapter sieve failed: %s", FormatSieve(a.Sieve))
	}
	if !a.Export.Executed || a.Export.ActualObservedValueCount != 0 || a.Export.NativeExportCount != 0 || a.Export.BridgeExportCount < 3 || !a.Export.NativePromotionBlocked || !a.Export.SchemaFailuresFailClosed {
		return fmt.Errorf("dry-run export leaked observed/native value: %s", FormatExport(a.Export))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoCurveFitPromoted || !a.Firewall.NoGSTPromotion || !a.Firewall.NoNativeCoefficientRayValue || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate454Inherited,
		StatusAdapterSchemaDefined,
		StatusDryRunFirewallValidated,
		StatusLabelledLocalRayDryRunAllowed,
		StatusLabelledOrientedRayDryRunAllowed,
		StatusNoObservedValuesImportedDefault,
		StatusBridgeOnlyExportsValidated,
		StatusEmpiricalFirewallPreserved,
		StatusFailedSpectrumOnlyNativePromotionRejected,
		StatusFailedMissingMetadataRejected,
		StatusFailedGSTNativePromotionRejected,
		StatusFailedCKMPMNSNativeSelectorRejected,
		StatusFailedObservedValuesRejectedByDefault,
		StatusFailedNativeCoefficientExportAbsent,
	}
}

func join(xs []string) string {
	if len(xs) == 0 {
		return "∅"
	}
	return strings.Join(xs, ", ")
}
