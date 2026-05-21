// Package generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit implements
// Gate 963: CanonicalFlavorSelector vs ExternalFlavorOrientationSeal Decision Audit.
//
// Gate 963 follows Gate 962's construction failure. It audits whether any
// current sealed-R4 object breaks the U(3) family-gauge ambiguity of the
// external C^3 generation carrier. If none does, it classifies downstream
// flavor-ledger work as requiring an explicit ExternalFlavorOrientationSeal.
package generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE963-GENERATION2CANONICALFLAVORSELECTORVSEXTERNALORIENTATIONSEALDECISIONAUDIT"
	InheritedStatus = "R4_EXTERNAL_C3_HAS_U3_FAMILY_ORBIT_NO_FLAVOR_BASIS"
	Verdict         = "NO_CANONICAL_FLAVOR_SELECTOR_FOUND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED_FOR_DOWNSTREAM_FLAVOR_LEDGER_TESTS"
	Classification  = "R4_FLAVOR_ORIENTATION_SOURCE_DECISION_EXTERNAL_SEAL_REQUIRED"
	ShortStatus     = "R4_REQUIRES_EXTERNAL_FLAVOR_ORIENTATION_SEAL"
	NextGate        = "NEXT_GATE964_EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLATION_AUDIT"
)

type SelectorStatus string

const (
	StatusCurrentCertificateBlocked SelectorStatus = "FAILED_NO_CANONICAL_SELECTOR_IN_CURRENT_CERTIFICATE"
	StatusU3Obstruction             SelectorStatus = "FAILED_U3_FAMILY_GAUGE_UNBROKEN"
	StatusSealRequired              SelectorStatus = "EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED"
	StatusDownstreamBoundary        SelectorStatus = "DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_ONLY_UNDER_SEAL"
)

type SelectorCandidate struct {
	Name                           string
	Source                         string
	Status                         SelectorStatus
	Audited                        bool
	CanonicalSelectorFound         bool
	BreaksU3FamilyGauge            bool
	SelectsCanonicalRepresentative bool
	UsesFlavorFormulaAsSource      bool
	UsesObservedFlavorData         bool
	UsesR3RowsAsFlavorBasis        bool
	UsesSocketAsFamilySelector     bool
	ProvidesNativeFlavorTheorem    bool
	ProvidesYukawaSpectrum         bool
	ProvidesCKMPMNS                bool
	AssignsParticles               bool
	OfficialLedgerUpdate           bool
	Supports                       []string
	Firewalls                      []string
}

type Decision struct {
	InheritedU3OrbitOnly                                 bool
	CanonicalFlavorSelectorFound                         bool
	CanonicalRepresentativeSelected                      bool
	U3FamilyGaugeFreedomRetained                         bool
	CurrentASHADataBreaksU3Gauge                         bool
	ExternalFlavorOrientationSealRequired                bool
	ExternalFlavorOrientationSealNative                  bool
	ExternalFlavorOrientationSealCanSelectRepresentative bool
	DownstreamFlavorLedgerTestsAllowedUnderSeal          bool
	NativeFlavorTheoremCertified                         bool
	YukawaEigenvaluesDerived                             bool
	PhysicalParticlesAssigned                            bool
	CKMPMNSDerived                                       bool
	OfficialLedgerUpdateAllowed                          bool
	R3DualSealPreserved                                  bool
	ExternalGenerationCarrierSealPreserved               bool
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	NextGate       string
	Problem        string
	RequiredSeal   string
	Decision       Decision
	Candidates     []SelectorCandidate
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	candidates := DefaultCandidates()
	if len(candidates) != 8 {
		return Analysis{}, fmt.Errorf("expected 8 selector/decision candidates, got %d", len(candidates))
	}
	decision := Decision{
		InheritedU3OrbitOnly:                                 true,
		CanonicalFlavorSelectorFound:                         false,
		CanonicalRepresentativeSelected:                      false,
		U3FamilyGaugeFreedomRetained:                         true,
		CurrentASHADataBreaksU3Gauge:                         false,
		ExternalFlavorOrientationSealRequired:                true,
		ExternalFlavorOrientationSealNative:                  false,
		ExternalFlavorOrientationSealCanSelectRepresentative: true,
		DownstreamFlavorLedgerTestsAllowedUnderSeal:          true,
		NativeFlavorTheoremCertified:                         false,
		YukawaEigenvaluesDerived:                             false,
		PhysicalParticlesAssigned:                            false,
		CKMPMNSDerived:                                       false,
		OfficialLedgerUpdateAllowed:                          false,
		R3DualSealPreserved:                                  true,
		ExternalGenerationCarrierSealPreserved:               true,
	}
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		Problem:        "Gate 962 produced only [Phi_flav]_{U(3)}: an orientation orbit of the sealed C^3 family carrier, not a canonical flavor basis.",
		RequiredSeal:   "ExternalFlavorOrientationSeal selecting a representative Phi_flav^seal in [Phi_flav]_{U(3)} for downstream ledger tests only",
		Decision:       decision,
		Candidates:     candidates,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 963 searches the current sealed R4 certificate for a CanonicalFlavorSelector and finds none. The external C^3 carrier retains U(3) family-gauge freedom; A_F^orient, R3 tracebody, socket/Fock structures, boundary activation, K7 remnants, and Boolean-octonionic projectors do not select a canonical flavor representative. Downstream flavor-ledger tests therefore require an explicit ExternalFlavorOrientationSeal. That seal may select a representative for tests, but it is not a native flavor theorem and does not derive Yukawa values, CKM/PMNS, particle assignments, or official ledger updates.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 963 identity")
	}
	if !a.Decision.InheritedU3OrbitOnly || !a.Decision.U3FamilyGaugeFreedomRetained || a.Decision.CurrentASHADataBreaksU3Gauge {
		return fmt.Errorf("Gate 963 must inherit U3 orbit obstruction: %#v", a.Decision)
	}
	if a.Decision.CanonicalFlavorSelectorFound || a.Decision.CanonicalRepresentativeSelected {
		return fmt.Errorf("Gate 963 must not certify canonical selector or representative: %#v", a.Decision)
	}
	if !a.Decision.ExternalFlavorOrientationSealRequired || a.Decision.ExternalFlavorOrientationSealNative || !a.Decision.ExternalFlavorOrientationSealCanSelectRepresentative {
		return fmt.Errorf("Gate 963 must require external orientation seal as non-native representative selector: %#v", a.Decision)
	}
	if !a.Decision.DownstreamFlavorLedgerTestsAllowedUnderSeal {
		return fmt.Errorf("Gate 963 must allow downstream tests only under seal: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremCertified || a.Decision.YukawaEigenvaluesDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.CKMPMNSDerived || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 963 overclaimed flavor theorem or downstream values: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved {
		return fmt.Errorf("Gate 963 must preserve inherited seals: %#v", a.Decision)
	}
	for _, c := range a.Candidates {
		if !c.Audited {
			return fmt.Errorf("candidate not audited: %#v", c)
		}
		if c.CanonicalSelectorFound || c.BreaksU3FamilyGauge || c.SelectsCanonicalRepresentative || c.ProvidesNativeFlavorTheorem || c.ProvidesYukawaSpectrum || c.ProvidesCKMPMNS || c.AssignsParticles || c.OfficialLedgerUpdate {
			return fmt.Errorf("candidate overclaimed flavor selector/theorem/value: %#v", c)
		}
		if c.UsesFlavorFormulaAsSource || c.UsesObservedFlavorData || c.UsesR3RowsAsFlavorBasis || c.UsesSocketAsFamilySelector {
			return fmt.Errorf("candidate used forbidden source as selector: %#v", c)
		}
	}
	return nil
}

func DefaultCandidates() []SelectorCandidate {
	return []SelectorCandidate{
		{
			Name:                   "current ASHA internal selector search",
			Source:                 "A_F^orient, R3 tracebody, socket ledger, Fock 1+3/B-L, boundary activation, K7 remnants, Boolean-octonionic projectors",
			Status:                 StatusCurrentCertificateBlocked,
			Audited:                true,
			CanonicalSelectorFound: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_CURRENT_CERTIFICATE_SEARCH_COMPLETED_FOR_CANONICAL_FLAVOR_SELECTOR",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE",
				"FAILED_ROUTE_A_F_ORIENT_R3_TRACEBODY_SOCKET_LEDGER_AND_PROJECTOR_DATA_DO_NOT_SELECT_FAMILY_BASIS",
			},
		},
		{
			Name:                   "U3 family-gauge obstruction",
			Source:                 "external C^3_gen,seal admits U(3) family-gauge rotations",
			Status:                 StatusU3Obstruction,
			Audited:                true,
			CanonicalSelectorFound: false,
			BreaksU3FamilyGauge:    false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_RETAINS_U3_FAMILY_GAUGE_FREEDOM",
			},
			Firewalls: []string{
				"FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA",
				"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_OF_FLAVOR_ORIENTATION_ORBIT",
			},
		},
		{
			Name:                        "ExternalFlavorOrientationSeal",
			Source:                      "choose Phi_flav^seal in [Phi_flav]_{U(3)} under explicit quarantine",
			Status:                      StatusSealRequired,
			Audited:                     true,
			CanonicalSelectorFound:      false,
			ProvidesNativeFlavorTheorem: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_CAN_SELECT_REPRESENTATIVE",
				"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED_FOR_DOWNSTREAM_FLAVOR_LEDGER_TESTS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
				"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CKM_PMNS_THEOREM",
			},
		},
		{
			Name:                       "A_F^orient selector reuse",
			Source:                     "post-orientation finite algebra interface",
			Status:                     StatusCurrentCertificateBlocked,
			Audited:                    true,
			UsesSocketAsFamilySelector: false,
			Firewalls: []string{
				"FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR",
				"FAILED_ROUTE_SOCKET_ORIENTATION_NOT_FAMILY_ORIENTATION",
			},
		},
		{
			Name:                    "R3 tracebody row matching",
			Source:                  "dual-sealed aggregate R3 row multiset",
			Status:                  StatusCurrentCertificateBlocked,
			Audited:                 true,
			UsesR3RowsAsFlavorBasis: false,
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
				"FAILED_ROUTE_R3_TRACEBODY_CANNOT_ORIENT_EXTERNAL_C3_BY_ROW_MATCHING",
			},
		},
		{
			Name:                       "Fock P3 / B-L selector reuse",
			Source:                     "Fock/projective 1+3 and B-L socket/internal-charge selectors",
			Status:                     StatusCurrentCertificateBlocked,
			Audited:                    true,
			UsesSocketAsFamilySelector: false,
			Firewalls: []string{
				"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_ORIENTATION",
			},
		},
		{
			Name:                      "flavor-ledger formulas as selector",
			Source:                    "epsilon_e, kappa_e, Koide, CKM/PMNS, observed masses, flavor wall residuals",
			Status:                    StatusCurrentCertificateBlocked,
			Audited:                   true,
			UsesFlavorFormulaAsSource: false,
			UsesObservedFlavorData:    false,
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
			},
		},
		{
			Name:                   "downstream permission boundary",
			Source:                 "sealed diagnostics after ExternalFlavorOrientationSeal",
			Status:                 StatusDownstreamBoundary,
			Audited:                true,
			ProvidesYukawaSpectrum: false,
			ProvidesCKMPMNS:        false,
			AssignsParticles:       false,
			OfficialLedgerUpdate:   false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_ONLY_UNDER_EXTERNAL_ORIENTATION_SEAL",
				"CONDITIONAL_SUPPORT_R3_DUALSEAL_EXTERNAL_C3_SEAL_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_MUST_REMAIN_VISIBLE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_CURRENT_CERTIFICATE_SEARCH_COMPLETED_FOR_CANONICAL_FLAVOR_SELECTOR",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_RETAINS_U3_FAMILY_GAUGE_FREEDOM",
		"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_CAN_SELECT_REPRESENTATIVE",
		"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_REQUIRED_FOR_DOWNSTREAM_FLAVOR_LEDGER_TESTS",
		"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_ONLY_UNDER_EXTERNAL_ORIENTATION_SEAL",
		"CONDITIONAL_SUPPORT_R3_DUALSEAL_EXTERNAL_C3_SEAL_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_MUST_REMAIN_VISIBLE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA",
		"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_OF_FLAVOR_ORIENTATION_ORBIT",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR",
		"FAILED_ROUTE_SOCKET_ORIENTATION_NOT_FAMILY_ORIENTATION",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBODY_CANNOT_ORIENT_EXTERNAL_C3_BY_ROW_MATCHING",
		"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_ORIENTATION",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_RETAINS_U3_FAMILY_GAUGE_FREEDOM",
		"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_CAN_SELECT_REPRESENTATIVE",
		"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_ONLY_UNDER_EXTERNAL_ORIENTATION_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA",
		"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_OF_FLAVOR_ORIENTATION_ORBIT",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
	}
}

func CandidateSupports(candidates []SelectorCandidate) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Supports...)
	}
	return out
}

func CandidateFailures(candidates []SelectorCandidate) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Firewalls...)
	}
	return out
}

func CandidateNotes(candidates []SelectorCandidate) []string {
	out := make([]string, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, c.Name+" => "+string(c.Status))
	}
	return out
}

func appendAll(parts ...[]string) []string {
	var out []string
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func containsAll(hay []string, needles []string) bool {
	set := map[string]bool{}
	for _, h := range hay {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}

func stringsJoin(v []string) string { return strings.Join(v, "; ") }
