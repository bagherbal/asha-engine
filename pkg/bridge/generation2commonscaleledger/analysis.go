// Package generation2commonscaleledger implements Gate 467:
// Common-Scale Running Ledger / Coefficient-Ray Comparator Design.
//
// Gate 466 proved that PDG-style mass rows alone cannot define the ASHA
// cylinder coordinates required by the CKM-null residual socket. Gate 467 is
// therefore not a numerical CKM run. It is a fail-closed schema theorem: it
// defines the minimal bridge-only data product that a future observed run must
// supply before alpha_u, phi_u, alpha_d, phi_d, and d_ud may be evaluated.
//
// The required data product is sector-indexed, common-scale/common-scheme, and
// rank-complete: each quark sector must carry a trace-zero spectrum invariant
// I_spec, an independent I_K comparator, complete branch tags {sigma_CP,n_C3},
// source/scale/scheme/uncertainty metadata, and bridge-only quarantine. CKM or
// Cabibbo values are explicitly rejected as ray-definition inputs.
package generation2commonscaleledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE467-COMMON-SCALE-RUNNING-LEDGER-COEFFICIENT-RAY-COMPARATOR-DESIGN"

	StatusGate466Inherited          = "CONDITIONAL_SUPPORT_GATE466_OBSERVED_ADAPTER_OBSTRUCTION_INHERITED"
	StatusProtocolDefined           = "CONDITIONAL_SUPPORT_COMMON_SCALE_COEFFICIENT_RAY_PROTOCOL_DEFINED"
	StatusCompleteSchemaAccepted    = "CONDITIONAL_SUPPORT_COMMON_SCALE_RAY_LEDGER_SCHEMA_ACCEPTED"
	StatusRankCompletenessValidated = "CONDITIONAL_SUPPORT_RANK_COMPLETE_COMPARATOR_REQUIREMENTS_VALIDATED"
	StatusBridgeOnlyDesignValidated = "CONDITIONAL_SUPPORT_COMMON_SCALE_COMPARATOR_DESIGN_BRIDGE_ONLY_VALIDATED"
	StatusFirewallPreserved         = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"

	StatusFailedMixedScaleRejected      = "FAILED_ROUTE_MIXED_SCALE_RUNNING_LEDGER_REJECTED"
	StatusFailedMissingIKRejected       = "FAILED_ROUTE_MISSING_IK_COMPARATOR_REJECTED"
	StatusFailedMissingBranchRejected   = "FAILED_ROUTE_MISSING_BRANCH_TAGS_REJECTED"
	StatusFailedMissingUncertainty      = "FAILED_ROUTE_UNCERTAINTY_PROPAGATION_MISSING"
	StatusFailedMassOnlyStillRankOne    = "FAILED_ROUTE_MASS_SPECTRA_ONLY_STILL_RANK_ONE"
	StatusFailedCabibboAsRayInput       = "FAILED_ROUTE_CABIBBO_USED_AS_RAY_INPUT_REJECTED"
	StatusFailedNativePromotionRejected = "FAILED_ROUTE_COMMON_SCALE_LEDGER_NATIVE_PROMOTION_REJECTED"
	StatusFailedObservedCKMNative       = "FAILED_ROUTE_OBSERVED_CKM_NATIVE_PREDICTION_REJECTED"
)

const (
	NativeFlavorDim       = 13
	KXYCoeffDim           = 9
	RequiredSectors       = 2
	RequiredMassesPerSect = 3
	ProjectiveRayDOF      = 2
	RequiredRayScalars    = 2
	RequiredBranchFields  = 2
	RequiredSchemaFields  = 12
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate454SpectrumOnlyRankOne       bool
	Gate456InverseRequiresTwoScalars bool
	Gate459BranchTagsRequired        bool
	Gate464CKMNullSocket             bool
	Gate465Airlock                   bool
	Gate466ObservedRowsImported      bool
	Gate466DUDUndefined              bool
	Gate466AlignmentNotComputable    bool
	NativeRegistryClean              bool
	Verdict                          string
}

type SectorRayLedger struct {
	Sector                   string
	MassLabels               []string
	CommonScale              string
	CommonScheme             string
	Source                   string
	SourceVersion            string
	UncertaintyModel         string
	RunningMassesCommonScale bool
	TraceZeroProjection      bool
	HasISpec                 bool
	ISpecKind                string
	HasIK                    bool
	IKKind                   string
	HasCPOddSign             bool
	HasC3Sheet               bool
	BranchTagKind            string
	DimensionlessComparators bool
	BridgeOnly               bool
	NativePromotionClaim     bool
	CabibboAsRayInput        bool
}

type LedgerEvaluation struct {
	Ledger   SectorRayLedger
	Accepted bool
	Verdict  string
	Reason   string
	Failures []string
}

type Protocol struct {
	Executed                         bool
	RequiredSectors                  []string
	RequiresCommonScale              bool
	RequiresCommonScheme             bool
	RequiresThreeMassesPerSector     bool
	RequiresTraceZeroProjection      bool
	RequiresISpec                    bool
	RequiresIK                       bool
	RequiresBranchTags               bool
	RequiresSourceScaleScheme        bool
	RequiresUncertaintyPropagation   bool
	RequiresDimensionlessComparators bool
	RequiresBridgeOnly               bool
	RejectsCabibboAsRayInput         bool
	RejectsNativePromotion           bool
	ProjectiveRayDOF                 int
	SpectrumOnlyRank                 int
	RequiredRayScalars               int
	RequiredBranchFields             int
	RequiredSchemaFields             int
	Verdict                          string
	Reason                           string
}

type SchemaAudit struct {
	Executed                   bool
	Ledgers                    []LedgerEvaluation
	AcceptedLedgers            int
	RejectedLedgers            int
	CompleteUSectorAccepted    bool
	CompleteDSectorAccepted    bool
	MixedScaleRejected         bool
	MissingIKRejected          bool
	MissingBranchRejected      bool
	MissingUncertaintyRejected bool
	MassOnlyRejected           bool
	CabibboRayInputRejected    bool
	NativePromotionRejected    bool
	BothSectorsReady           bool
	DUDComputableIfNumeric     bool
	DUDComputedNow             bool
	Verdict                    string
	Reason                     string
}

type Firewall struct {
	Executed                      bool
	CommonScaleProtocolNative     bool
	EmpiricalCoordinatesNative    bool
	DUDNativePrediction           bool
	CKMNativePrediction           bool
	CabibboUsedAsRayInput         bool
	QuarkMassesAsTheoremInput     bool
	NativeRegistryWritten         bool
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
	Protocol    Protocol
	Schema      SchemaAudit
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
	a.Protocol = buildProtocol()
	a.Schema = buildSchemaAudit()
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
		Gate454SpectrumOnlyRankOne:       true,
		Gate456InverseRequiresTwoScalars: true,
		Gate459BranchTagsRequired:        true,
		Gate464CKMNullSocket:             true,
		Gate465Airlock:                   true,
		Gate466ObservedRowsImported:      true,
		Gate466DUDUndefined:              true,
		Gate466AlignmentNotComputable:    true,
		NativeRegistryClean:              true,
		Verdict:                          StatusGate466Inherited,
	}
}

func buildProtocol() Protocol {
	return Protocol{
		Executed:                         true,
		RequiredSectors:                  []string{"u", "d"},
		RequiresCommonScale:              true,
		RequiresCommonScheme:             true,
		RequiresThreeMassesPerSector:     true,
		RequiresTraceZeroProjection:      true,
		RequiresISpec:                    true,
		RequiresIK:                       true,
		RequiresBranchTags:               true,
		RequiresSourceScaleScheme:        true,
		RequiresUncertaintyPropagation:   true,
		RequiresDimensionlessComparators: true,
		RequiresBridgeOnly:               true,
		RejectsCabibboAsRayInput:         true,
		RejectsNativePromotion:           true,
		ProjectiveRayDOF:                 ProjectiveRayDOF,
		SpectrumOnlyRank:                 1,
		RequiredRayScalars:               RequiredRayScalars,
		RequiredBranchFields:             RequiredBranchFields,
		RequiredSchemaFields:             RequiredSchemaFields,
		Verdict:                          StatusProtocolDefined,
		Reason:                           "a future observed CKM bridge run needs common-scale sector spectra plus I_spec, I_K, and {sigma_CP,n_C3}; Cabibbo data may be a target residual only, never a ray-definition input",
	}
}

func canonicalLedgers() []SectorRayLedger {
	base := SectorRayLedger{
		CommonScale:              "mu_star_common_bridge_scale",
		CommonScheme:             "MS-bar-or-explicit-running-scheme",
		Source:                   "bridge-running-ledger-source",
		SourceVersion:            "versioned external running calculation",
		UncertaintyModel:         "covariance-or-interval-propagation-declared",
		RunningMassesCommonScale: true,
		TraceZeroProjection:      true,
		HasISpec:                 true,
		ISpecKind:                "dimensionless trace-zero spectrum invariant",
		HasIK:                    true,
		IKKind:                   "dimensionless K-axis comparator independent of spectrum invariant",
		HasCPOddSign:             true,
		HasC3Sheet:               true,
		BranchTagKind:            "bridge-only {sigma_CP,n_C3}",
		DimensionlessComparators: true,
		BridgeOnly:               true,
	}
	u := base
	u.Sector = "u"
	u.MassLabels = []string{"m_u(mu_star)", "m_c(mu_star)", "m_t(mu_star)"}
	d := base
	d.Sector = "d"
	d.MassLabels = []string{"m_d(mu_star)", "m_s(mu_star)", "m_b(mu_star)"}
	return []SectorRayLedger{u, d}
}

func rejectedLedgers() []SectorRayLedger {
	mk := func(name string) SectorRayLedger {
		l := canonicalLedgers()[0]
		l.Sector = name
		return l
	}
	mixed := mk("u")
	mixed.CommonScale = "mixed: 2 GeV, mu=m_c, mu=m_t"
	mixed.RunningMassesCommonScale = false

	missingIK := mk("u")
	missingIK.HasIK = false
	missingIK.IKKind = ""

	missingBranch := mk("d")
	missingBranch.HasCPOddSign = false
	missingBranch.HasC3Sheet = false
	missingBranch.BranchTagKind = ""

	missingUncertainty := mk("d")
	missingUncertainty.UncertaintyModel = ""

	massOnly := mk("u")
	massOnly.HasIK = false
	massOnly.HasCPOddSign = false
	massOnly.HasC3Sheet = false
	massOnly.BranchTagKind = ""
	massOnly.IKKind = ""

	cabibboAsRay := mk("u-d")
	cabibboAsRay.MassLabels = []string{"|V_us|"}
	cabibboAsRay.CabibboAsRayInput = true

	nativeClaim := mk("u")
	nativeClaim.NativePromotionClaim = true

	return []SectorRayLedger{mixed, missingIK, missingBranch, missingUncertainty, massOnly, cabibboAsRay, nativeClaim}
}

func EvaluateLedger(l SectorRayLedger) LedgerEvaluation {
	failures := []string{}
	if l.NativePromotionClaim {
		failures = append(failures, StatusFailedNativePromotionRejected)
	}
	if l.CabibboAsRayInput {
		failures = append(failures, StatusFailedCabibboAsRayInput)
	}
	if !l.RunningMassesCommonScale || l.CommonScale == "" || strings.Contains(l.CommonScale, "mixed") || l.CommonScheme == "" {
		failures = append(failures, StatusFailedMixedScaleRejected)
	}
	if len(l.MassLabels) != RequiredMassesPerSect && !l.CabibboAsRayInput {
		failures = append(failures, StatusFailedMixedScaleRejected)
	}
	if !l.HasIK || l.IKKind == "" {
		failures = append(failures, StatusFailedMissingIKRejected)
	}
	if !l.HasCPOddSign || !l.HasC3Sheet || l.BranchTagKind == "" {
		failures = append(failures, StatusFailedMissingBranchRejected)
	}
	if l.UncertaintyModel == "" || l.Source == "" || l.SourceVersion == "" {
		failures = append(failures, StatusFailedMissingUncertainty)
	}
	if !l.HasIK && (!l.HasCPOddSign || !l.HasC3Sheet) {
		failures = append(failures, StatusFailedMassOnlyStillRankOne)
	}
	if !l.TraceZeroProjection || !l.HasISpec || !l.DimensionlessComparators || !l.BridgeOnly {
		failures = append(failures, StatusFailedMassOnlyStillRankOne)
	}
	if len(failures) > 0 {
		return LedgerEvaluation{Ledger: l, Accepted: false, Verdict: failures[0], Reason: "sector ledger fails the common-scale rank-complete comparator contract", Failures: failures}
	}
	return LedgerEvaluation{Ledger: l, Accepted: true, Verdict: StatusCompleteSchemaAccepted, Reason: "sector ledger has common scale/scheme, trace-zero spectrum invariant, I_K, branch tags, uncertainty model, and bridge-only quarantine"}
}

func buildSchemaAudit() SchemaAudit {
	ledgers := []LedgerEvaluation{}
	for _, l := range canonicalLedgers() {
		ledgers = append(ledgers, EvaluateLedger(l))
	}
	for _, l := range rejectedLedgers() {
		ledgers = append(ledgers, EvaluateLedger(l))
	}
	s := SchemaAudit{Executed: true, Ledgers: ledgers}
	for _, e := range ledgers {
		if e.Accepted {
			s.AcceptedLedgers++
			if e.Ledger.Sector == "u" {
				s.CompleteUSectorAccepted = true
			}
			if e.Ledger.Sector == "d" {
				s.CompleteDSectorAccepted = true
			}
		} else {
			s.RejectedLedgers++
			for _, f := range e.Failures {
				switch f {
				case StatusFailedMixedScaleRejected:
					s.MixedScaleRejected = true
				case StatusFailedMissingIKRejected:
					s.MissingIKRejected = true
				case StatusFailedMissingBranchRejected:
					s.MissingBranchRejected = true
				case StatusFailedMissingUncertainty:
					s.MissingUncertaintyRejected = true
				case StatusFailedMassOnlyStillRankOne:
					s.MassOnlyRejected = true
				case StatusFailedCabibboAsRayInput:
					s.CabibboRayInputRejected = true
				case StatusFailedNativePromotionRejected:
					s.NativePromotionRejected = true
				}
			}
		}
	}
	s.BothSectorsReady = s.CompleteUSectorAccepted && s.CompleteDSectorAccepted
	s.DUDComputableIfNumeric = s.BothSectorsReady
	s.DUDComputedNow = false
	s.Verdict = StatusBridgeOnlyDesignValidated
	s.Reason = "the schema accepts one complete u ledger and one complete d ledger, rejects every incomplete or unsafe variant, and still does not compute d_ud because Gate467 is a design/provenance contract, not an observed-value run"
	return s
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		CommonScaleProtocolNative:     false,
		EmpiricalCoordinatesNative:    false,
		DUDNativePrediction:           false,
		CKMNativePrediction:           false,
		CabibboUsedAsRayInput:         false,
		QuarkMassesAsTheoremInput:     false,
		NativeRegistryWritten:         false,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate467 defines the missing empirical bridge schema only; it does not turn common-scale ledgers, I_K, branch tags, d_ud, or CKM residuals into native ASHA law-space",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        468,
		Title:       "Common-Scale Synthetic Inversion Run / Uncertainty Propagation Harness",
		Reason:      "Gate467 defines the complete u/d comparator data product; the next safe step is to exercise the full inverse and d_ud formula on synthetic or redacted complete ledgers with interval propagation.",
		PrimaryTask: "apply the Gate456 inverse and Gate464 d_ud socket to rank-complete synthetic common-scale u/d records, while keeping all numerical coordinates bridge-only",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate466AlignmentNotComputable || !a.Inheritance.Gate456InverseRequiresTwoScalars || !a.Inheritance.Gate459BranchTagsRequired || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate467 inheritance incomplete")
	}
	if !a.Protocol.Executed || !a.Protocol.RequiresCommonScale || !a.Protocol.RequiresCommonScheme || !a.Protocol.RequiresISpec || !a.Protocol.RequiresIK || !a.Protocol.RequiresBranchTags || !a.Protocol.RequiresUncertaintyPropagation || !a.Protocol.RequiresBridgeOnly || !a.Protocol.RejectsCabibboAsRayInput || !a.Protocol.RejectsNativePromotion || a.Protocol.ProjectiveRayDOF != ProjectiveRayDOF || a.Protocol.SpectrumOnlyRank != 1 || a.Protocol.RequiredRayScalars != RequiredRayScalars || a.Protocol.RequiredBranchFields != RequiredBranchFields {
		return fmt.Errorf("common-scale protocol not complete")
	}
	if !a.Schema.Executed || a.Schema.AcceptedLedgers != RequiredSectors || a.Schema.RejectedLedgers < 7 || !a.Schema.CompleteUSectorAccepted || !a.Schema.CompleteDSectorAccepted || !a.Schema.MixedScaleRejected || !a.Schema.MissingIKRejected || !a.Schema.MissingBranchRejected || !a.Schema.MissingUncertaintyRejected || !a.Schema.MassOnlyRejected || !a.Schema.CabibboRayInputRejected || !a.Schema.NativePromotionRejected || !a.Schema.BothSectorsReady || !a.Schema.DUDComputableIfNumeric || a.Schema.DUDComputedNow {
		return fmt.Errorf("schema audit did not accept/reject expected ledgers")
	}
	if !a.Firewall.Executed || a.Firewall.CommonScaleProtocolNative || a.Firewall.EmpiricalCoordinatesNative || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CabibboUsedAsRayInput || a.Firewall.QuarkMassesAsTheoremInput || a.Firewall.NativeRegistryWritten || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall violated by common-scale design")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Schema.BothSectorsReady && a.Schema.DUDComputableIfNumeric && !a.Schema.DUDComputedNow && !a.Firewall.CKMNativePrediction {
		return "Gate 467 defines the exact bridge-only data product missing from Gate 466: common-scale u/d running spectra, I_spec, I_K, complete branch tags, provenance, and uncertainty propagation. It proves that a future d_ud calculation is allowed only after this rank-complete ledger exists; Cabibbo data remains a target residual, not an input coordinate or native prediction."
	}
	return "Gate 467 failed to define a safe common-scale comparator design."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t spectrum_rank_one=%t inverse_two_scalars=%t branch_tags=%t ckm_socket=%t airlock=%t gate466_rows=%t gate466_dud_undefined=%t gate466_not_computable=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate454SpectrumOnlyRankOne, x.Gate456InverseRequiresTwoScalars, x.Gate459BranchTagsRequired, x.Gate464CKMNullSocket, x.Gate465Airlock, x.Gate466ObservedRowsImported, x.Gate466DUDUndefined, x.Gate466AlignmentNotComputable, x.NativeRegistryClean, x.Verdict)
}

func FormatProtocol(x Protocol) string {
	return fmt.Sprintf("executed=%t sectors=%s common_scale=%t common_scheme=%t three_masses=%t trace_zero=%t I_spec=%t I_K=%t branch_tags=%t metadata=%t uncertainty=%t dimensionless=%t bridge_only=%t reject_cabibbo_as_ray=%t reject_native=%t ray_dof=%d spectrum_rank=%d ray_scalars=%d branch_fields=%d schema_fields=%d verdict=%s reason=%s", x.Executed, strings.Join(x.RequiredSectors, ","), x.RequiresCommonScale, x.RequiresCommonScheme, x.RequiresThreeMassesPerSector, x.RequiresTraceZeroProjection, x.RequiresISpec, x.RequiresIK, x.RequiresBranchTags, x.RequiresSourceScaleScheme, x.RequiresUncertaintyPropagation, x.RequiresDimensionlessComparators, x.RequiresBridgeOnly, x.RejectsCabibboAsRayInput, x.RejectsNativePromotion, x.ProjectiveRayDOF, x.SpectrumOnlyRank, x.RequiredRayScalars, x.RequiredBranchFields, x.RequiredSchemaFields, x.Verdict, x.Reason)
}

func FormatLedger(x SectorRayLedger) string {
	return fmt.Sprintf("sector=%s masses=%s scale=%s scheme=%s source=%s version=%s uncertainty=%s common_running=%t trace_zero=%t I_spec=%t I_K=%t cp_sign=%t c3_sheet=%t dimensionless=%t bridge_only=%t cabibbo_as_ray=%t native_claim=%t", x.Sector, strings.Join(x.MassLabels, ","), x.CommonScale, x.CommonScheme, x.Source, x.SourceVersion, x.UncertaintyModel, x.RunningMassesCommonScale, x.TraceZeroProjection, x.HasISpec, x.HasIK, x.HasCPOddSign, x.HasC3Sheet, x.DimensionlessComparators, x.BridgeOnly, x.CabibboAsRayInput, x.NativePromotionClaim)
}

func FormatEvaluation(x LedgerEvaluation) string {
	return fmt.Sprintf("ledger={%s} accepted=%t verdict=%s failures=%s reason=%s", FormatLedger(x.Ledger), x.Accepted, x.Verdict, strings.Join(x.Failures, ","), x.Reason)
}

func FormatSchema(x SchemaAudit) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d u=%t d=%t mixed_scale_rejected=%t missing_IK=%t missing_branch=%t missing_uncertainty=%t mass_only=%t cabibbo_as_ray=%t native_promotion=%t both_ready=%t dud_computable_if_numeric=%t dud_computed_now=%t verdict=%s reason=%s", x.Executed, x.AcceptedLedgers, x.RejectedLedgers, x.CompleteUSectorAccepted, x.CompleteDSectorAccepted, x.MixedScaleRejected, x.MissingIKRejected, x.MissingBranchRejected, x.MissingUncertaintyRejected, x.MassOnlyRejected, x.CabibboRayInputRejected, x.NativePromotionRejected, x.BothSectorsReady, x.DUDComputableIfNumeric, x.DUDComputedNow, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t protocol_native=%t coords_native=%t dud_native=%t ckm_native=%t cabibbo_as_ray=%t masses_theorem_input=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.CommonScaleProtocolNative, x.EmpiricalCoordinatesNative, x.DUDNativePrediction, x.CKMNativePrediction, x.CabibboUsedAsRayInput, x.QuarkMassesAsTheoremInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{
		StatusGate466Inherited,
		StatusProtocolDefined,
		StatusCompleteSchemaAccepted,
		StatusRankCompletenessValidated,
		StatusBridgeOnlyDesignValidated,
		StatusFirewallPreserved,
		StatusFailedMixedScaleRejected,
		StatusFailedMissingIKRejected,
		StatusFailedMissingBranchRejected,
		StatusFailedMissingUncertainty,
		StatusFailedMassOnlyStillRankOne,
		StatusFailedCabibboAsRayInput,
		StatusFailedNativePromotionRejected,
		StatusFailedObservedCKMNative,
	}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 467 Registry Audit — Common-Scale Running Ledger / Coefficient-Ray Comparator Design\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusBridgeOnlyDesignValidated + "`\n\n")
	b.WriteString("Gate 467 does not compute `d_ud`. It defines the exact bridge-only input ledger that was missing in Gate 466. A future observed run must supply common-scale/common-scheme u and d sector spectra, the independent `I_K` comparator, complete `{sigma_CP,n_C3}` branch tags, provenance, and uncertainty propagation before the ASHA cylinder socket is allowed to evaluate coordinates.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Required protocol\n\n")
	b.WriteString(FormatProtocol(a.Protocol) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("I_spec = 2 cos(3 phi)/(alpha^2+3)^(3/2)\n")
	b.WriteString("I_K    = alpha/sqrt(alpha^2+3)\n")
	b.WriteString("branch = {sigma_CP, n_C3}\n")
	b.WriteString("d_ud   = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n")
	b.WriteString("```\n\n")

	b.WriteString("## Schema audit\n\n")
	b.WriteString(FormatSchema(a.Schema) + "\n\n")
	b.WriteString("| Sector / probe | Accepted | Verdict | Ledger |\n")
	b.WriteString("|---|---:|---|---|\n")
	for _, e := range a.Schema.Ledgers {
		b.WriteString(fmt.Sprintf("| `%s` | %t | `%s` | %s |\n", esc(e.Ledger.Sector), e.Accepted, esc(e.Verdict), esc(FormatLedger(e.Ledger))))
	}
	b.WriteString("\n")
	b.WriteString("The two accepted ledgers are schemas only: they prove that a complete bridge data product can be represented. They do not contain numerical coordinates and do not evaluate `d_ud`. Every mass-only, mixed-scale, missing-`I_K`, missing-branch, missing-uncertainty, Cabibbo-as-coordinate, or native-promotion attempt fails closed.\n\n")

	b.WriteString("## Native firewall proof\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("The common-scale ledger is an empirical bridge contract, not a native theorem. `K_gen` and `X_triangle` remain structural; `Y_phase`, sector coefficients, branch tags, and any future coordinates remain quarantined. Cabibbo/CKM data may only be a residual target, never an input used to define the ray.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
