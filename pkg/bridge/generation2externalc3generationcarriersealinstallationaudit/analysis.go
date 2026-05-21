// Package generation2externalc3generationcarriersealinstallationaudit implements
// Gate 960: ExternalC3 GenerationCarrier Seal Installation Audit.
//
// Gate 960 follows the Gate 959 fork. The active ASHA board has no certified
// native generation carrier in the current certificate, so this gate installs
// C^3 only as an explicit external generation-carrier seal. This keeps R4 work
// moving while preserving that generation multiplicity is not native, R3 stays
// dual-sealed, and flavor/Yukawa/particle assignments remain blocked.
package generation2externalc3generationcarriersealinstallationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE960-GENERATION2EXTERNALC3GENERATIONCARRIERSEALINSTALLATIONAUDIT"
	InheritedStatus = "R4_REQUIRES_EXTERNAL_SEAL_OR_NEW_PARENT_AIRLOCK"
	Verdict         = "EXTERNAL_C3_GENERATION_CARRIER_SEAL_INSTALLED_R4_CAN_PROCEED_SEALED_NOT_NATIVE"
	Classification  = "R4_EXTERNAL_GENERATION_CARRIER_SEALED_NO_NATIVE_MULTIPLICITY_THEOREM"
	ShortStatus     = "R4_SEALED_GENERATION_CARRIER_AVAILABLE_NO_FLAVOR_MAP"
	NextGate        = "NEXT_GATE961_FLAVOR_ORIENTATION_MAP_PRECONDITION_AUDIT_UNDER_EXTERNAL_C3_AND_R3_DUALSEAL"
)

type SealStatus string

const (
	SealInstalled           SealStatus = "INSTALLED_AS_EXPLICIT_SEAL"
	SealInherited           SealStatus = "INHERITED_AND_PRESERVED"
	SealForbiddenPromotion  SealStatus = "FORBIDS_NATIVE_PROMOTION"
	SealDownstreamBlocked   SealStatus = "DOWNSTREAM_BLOCKED"
	SealNativeTheoremDenied SealStatus = "NOT_NATIVE_THEOREM"
)

type SealAudit struct {
	Name                       string
	Object                     string
	Status                     SealStatus
	AllowedRole                string
	ForbiddenRole              string
	Installed                  bool
	NativeGenerationTheorem    bool
	CompatibleWithR3DualSeal   bool
	FlavorOrientationProvided  bool
	IndividualYukawaProvided   bool
	PhysicalAssignmentProvided bool
	OfficialLedgerUpdate       bool
	Supports                   []string
	Firewalls                  []string
}

type Decision struct {
	ExternalC3SealInstalled       bool
	NativeGenerationCarrier       bool
	R4MayProceedUnderSeals        bool
	R3DualSealPreserved           bool
	FlavorOrientationMapAvailable bool
	IndividualYukawaAllowed       bool
	PhysicalAssignmentAllowed     bool
	OfficialLedgerUpdateAllowed   bool
	ParentAirlockStillOpen        bool
}

type Analysis struct {
	AuditID                    string
	Inherited                  string
	Verdict                    string
	Classification             string
	ShortStatus                string
	NextGate                   string
	ExternalCarrier            string
	GenerationCarrierAvailable bool
	NativeMultiplicityTheorem  bool
	Decision                   Decision
	Audits                     []SealAudit
	Supports                   []string
	Failures                   []string
	Final                      string
}

func BuildDefault() (Analysis, error) {
	audits := DefaultAudits()
	if len(audits) != 4 {
		return Analysis{}, fmt.Errorf("expected 4 seal audits, got %d", len(audits))
	}
	decision := Decision{
		ExternalC3SealInstalled:       true,
		NativeGenerationCarrier:       false,
		R4MayProceedUnderSeals:        true,
		R3DualSealPreserved:           true,
		FlavorOrientationMapAvailable: false,
		IndividualYukawaAllowed:       false,
		PhysicalAssignmentAllowed:     false,
		OfficialLedgerUpdateAllowed:   false,
		ParentAirlockStillOpen:        true,
	}
	a := Analysis{
		AuditID:                    AuditID,
		Inherited:                  InheritedStatus,
		Verdict:                    Verdict,
		Classification:             Classification,
		ShortStatus:                ShortStatus,
		NextGate:                   NextGate,
		ExternalCarrier:            "G_gen^seal = C^3",
		GenerationCarrierAvailable: true,
		NativeMultiplicityTheorem:  false,
		Decision:                   decision,
		Audits:                     audits,
		Supports:                   Supports(),
		Failures:                   Failures(),
		Final:                      "Gate 960 installs ExternalGenerationCarrierSeal(C3) as an explicit quarantined family multiplicity carrier. R4 may proceed only under the inherited R3 dual seal plus the external C3 generation-carrier seal. This does not derive native generation multiplicity, does not supply flavor orientation, does not assign particles, does not produce individual Yukawa values, and does not update official ledgers. The native parent-airlock route remains open as a separate infrastructure problem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 960 identity")
	}
	if !a.GenerationCarrierAvailable || a.NativeMultiplicityTheorem || a.Decision.NativeGenerationCarrier {
		return fmt.Errorf("Gate 960 must install seal but deny native multiplicity theorem")
	}
	if !a.Decision.ExternalC3SealInstalled || !a.Decision.R4MayProceedUnderSeals || !a.Decision.R3DualSealPreserved {
		return fmt.Errorf("Gate 960 must install external C3 seal and preserve R3 dual seal: %#v", a.Decision)
	}
	if a.Decision.FlavorOrientationMapAvailable || a.Decision.IndividualYukawaAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 960 overclaimed flavor/Yukawa/particle/official status: %#v", a.Decision)
	}
	if len(a.Audits) != 4 {
		return fmt.Errorf("expected 4 audits, got %d", len(a.Audits))
	}
	for _, au := range a.Audits {
		if au.NativeGenerationTheorem || au.FlavorOrientationProvided || au.IndividualYukawaProvided || au.PhysicalAssignmentProvided || au.OfficialLedgerUpdate {
			return fmt.Errorf("audit overclaimed native/downstream status: %#v", au)
		}
	}
	return nil
}

func DefaultAudits() []SealAudit {
	return []SealAudit{
		{
			Name:                    "seal declaration",
			Object:                  "ExternalGenerationCarrierSeal(C3), G_gen^seal = C^3",
			Status:                  SealInstalled,
			AllowedRole:             "family multiplicity carrier under explicit seal",
			ForbiddenRole:           "native ASHA generation theorem",
			Installed:               true,
			NativeGenerationTheorem: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_SEAL_INSTALLED",
				"CONDITIONAL_SUPPORT_G_GEN_SEAL_EQUALS_C3_AVAILABLE_AS_QUARANTINED_FAMILY_CARRIER",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_EXTERNAL_C3_SEAL_DOES_NOT_DERIVE_GENERATION_MULTIPLICITY",
			},
		},
		{
			Name:                     "R3 dual-seal compatibility",
			Object:                   "ScalarSourceSeal(S_split) plus PostOrientationSeal(A_F^orient)",
			Status:                   SealInherited,
			AllowedRole:              "aggregate R3 tracebody inherited only under explicit dual seal",
			ForbiddenRole:            "native R3 promotion or official diagnostic update",
			CompatibleWithR3DualSeal: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_SEAL_COMPATIBLE_WITH_R3_DUALSEAL",
				"CONDITIONAL_SUPPORT_R4_MAY_PROCEED_ONLY_WITH_SCALAR_SOURCE_AND_POST_ORIENTATION_SEALS_VISIBLE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:                      "no flavor orientation yet",
			Object:                    "three family slots without orientation map",
			Status:                    SealDownstreamBlocked,
			AllowedRole:               "slot carrier only; no labeling or orientation",
			ForbiddenRole:             "electron/muon/tau, quark, CKM, PMNS, hierarchy, or spectrum assignment",
			FlavorOrientationProvided: false,
			IndividualYukawaProvided:  false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_PROVIDES_ONLY_THREE_FAMILY_SLOTS",
				"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_BECOMES_NEXT_EXPLICIT_WOUND",
			},
			Firewalls: []string{
				"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
			},
		},
		{
			Name:                    "parent-airlock route remains native dream",
			Object:                  "D4/Spin(8) parent-board airlock to Lambda4/K7",
			Status:                  SealForbiddenPromotion,
			AllowedRole:             "future native infrastructure project",
			ForbiddenRole:           "implicit support for current sealed C3 carrier",
			NativeGenerationTheorem: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_PARENT_AIRLOCK_REMAINS_SEPARATE_NATIVE_ROUTE",
				"CONDITIONAL_SUPPORT_EXTERNAL_SEAL_DOES_NOT_CLOSE_PARENT_AIRLOCK_WOUND",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET",
				"FAILED_ROUTE_NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_SEAL_COMPATIBLE_WITH_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_PROVIDES_ONLY_THREE_FAMILY_SLOTS",
		"CONDITIONAL_SUPPORT_R4_CAN_PROCEED_UNDER_EXTERNAL_GENERATION_CARRIER_SEAL",
		"CONDITIONAL_SUPPORT_PARENT_AIRLOCK_REMAINS_SEPARATE_NATIVE_ROUTE",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_BECOMES_NEXT_EXPLICIT_WOUND",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_EXTERNAL_C3_SEAL_DOES_NOT_DERIVE_GENERATION_MULTIPLICITY",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
		"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_SEAL_COMPATIBLE_WITH_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_PROVIDES_ONLY_THREE_FAMILY_SLOTS",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_BECOMES_NEXT_EXPLICIT_WOUND",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}

func AuditSupports(audits []SealAudit) []string {
	var out []string
	for _, au := range audits {
		out = append(out, au.Supports...)
	}
	return out
}

func AuditFailures(audits []SealAudit) []string {
	var out []string
	for _, au := range audits {
		out = append(out, au.Firewalls...)
	}
	return out
}

func AuditNotes(audits []SealAudit) []string {
	out := make([]string, 0, len(audits))
	for _, au := range audits {
		out = append(out, au.Name+" => "+string(au.Status))
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
