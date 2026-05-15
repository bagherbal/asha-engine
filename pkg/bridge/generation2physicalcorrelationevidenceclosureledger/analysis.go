// Package generation2physicalcorrelationevidenceclosureledger implements Gate 551:
// Physical Correlation Evidence Board Sector Closure Ledger.
//
// Gate 550 proved that the evidence-board parser and governance plumbing can
// process a checksum-protected synthetic fixture while refusing to board it as
// real bridge evidence. Gate 551 closes the full physical-correlation evidence
// board sector: Schwinger source airlocks, authenticity, import switches,
// authorization, comparator harnesses, release review, evidence-board citation,
// and synthetic board rejection are frozen as bridge-only boundaries. No bridge
// evidence is accepted, no physical source is loaded, and no native theorem is
// written.
package generation2physicalcorrelationevidenceclosureledger

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticevidenceboardadapter"
)

const (
	AuditID = "GATE551-PHYSICAL-CORRELATION-EVIDENCE-BOARD-SECTOR-CLOSURE-LEDGER"

	StatusGate550SyntheticEvidenceInherited = "CONDITIONAL_SUPPORT_GATE550_SYNTHETIC_EVIDENCE_BOARD_INHERITED"
	StatusClosureLedgerEmitted              = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_EVIDENCE_BOARD_SECTOR_CLOSURE_LEDGER_EMITTED"
	StatusNativeFrontierFrozen              = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_NATIVE_FRONTIER_FROZEN"
	StatusBridgeFrontierMapped              = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_BRIDGE_FRONTIER_MAPPED"
	StatusEnvironmentalFrontierMapped       = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_ENVIRONMENTAL_FRONTIER_MAPPED"
	StatusSchwingerSourceBlockClosed        = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_SCHWINGER_SOURCE_BLOCK_CLOSED"
	StatusSourceAuthenticityBlockClosed     = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_SOURCE_AUTHENTICITY_BLOCK_CLOSED"
	StatusRealImportSwitchBlockClosed       = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_REAL_IMPORT_SWITCH_BLOCK_CLOSED"
	StatusAuthorizationBlockClosed          = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_AUTHORIZATION_BLOCK_CLOSED"
	StatusComparatorHarnessBlockClosed      = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_COMPARATOR_HARNESS_BLOCK_CLOSED"
	StatusReleaseReviewBlockClosed          = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_RELEASE_REVIEW_BLOCK_CLOSED"
	StatusEvidenceBoardBlockClosed          = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_CITATION_BOARD_BLOCK_CLOSED"
	StatusNoBridgeEvidenceBoarded           = "CONDITIONAL_SUPPORT_NO_PHYSICAL_CORRELATION_BRIDGE_EVIDENCE_BOARDED_IN_GATE551"
	StatusClosureFirewallMatrixComplete     = "CONDITIONAL_SUPPORT_GATE551_EVIDENCE_BOARD_FIREWALL_MATRIX_COMPLETE"

	StatusFailedClosureDoesNotDeriveSchwinger = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedClosureDoesNotProveOS         = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedClosureDoesNotGrantWick       = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedClosureDoesNotSelectHilbert   = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedClosureDoesNotDeriveHamilton  = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedClosureDoesNotGrantUnitary    = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedClosureDoesNotGrantGlobal     = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedClosureDoesNotSelectArrow     = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedClosureNoRealSource           = "FAILED_ROUTE_NO_AUTHENTICATED_NON_SYNTHETIC_CORRELATION_SOURCE_IN_GATE551"
	StatusFailedClosureNoBoardEntry           = "FAILED_ROUTE_NO_RELEASED_PHYSICAL_CORRELATION_BOARD_ENTRY_IN_GATE551"
	StatusFailedClosureEvidenceBoardNotNative = "FAILED_ROUTE_EVIDENCE_BOARD_SECTOR_CLOSURE_DOES_NOT_GRANT_NATIVE_LAW"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_GATE551_PHYSICAL_CORRELATION_EVIDENCE_BOARD_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked          = "FIREWALL_BLOCKED_GATE551_EVIDENCE_BOARD_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate550ManifestParsed        bool
	Gate550RowsAccepted          int
	Gate550ChecksumVerified      bool
	Gate550CitationParsed        bool
	Gate550UncertaintyParsed     bool
	Gate550ReproducibilityParsed bool
	Gate550RevocationParsed      bool
	Gate550VersionedIndexParsed  bool
	Gate550NativeDeltaZero       bool
	Gate550SyntheticBlocked      bool
	Gate550NoBridgeEvidence      bool
	Gate550NoRealSource          bool
	Gate550NativeWriteLocked     bool
	Gate550NoPhysicalClaims      bool
	Gate550RedirectsToGate551    bool

	Verdict string
	Reason  string
}

type FrontierRow struct {
	Gate        int
	Block       string
	Native      string
	Bridge      string
	Environment string
	Firewall    string
}

type ClosureLedger struct {
	Executed bool
	Rows     []FrontierRow

	RowCount                 int
	NativeRows               int
	BridgeRows               int
	EnvironmentalRows        int
	FailedRouteRows          int
	SchwingerSourceClosed    bool
	SourceAuthenticityClosed bool
	RealImportSwitchClosed   bool
	AuthorizationClosed      bool
	ComparatorHarnessClosed  bool
	ReleaseReviewClosed      bool
	EvidenceBoardClosed      bool
	NativeFrontierFrozen     bool
	BridgeFrontierMapped     bool
	EnvironmentalMapped      bool

	Verdict string
	Reason  string
}

type SectorGuard struct {
	Executed bool

	PhysicalSchwingerFunctionsImported bool
	AuthenticatedNonSyntheticSource    bool
	SourceAuthenticityAccepted         bool
	RealImportSwitchEnabled            bool
	OperatorIntentForRealImport        bool
	ComparatorAuthorized               bool
	ComparatorExecutedOnRealSource     bool
	ComparatorOutputReleased           bool
	BridgeEvidenceReleased             bool
	EvidenceBoardManifestImported      bool
	EvidenceBoardEntryAccepted         bool
	EvidenceEntriesAccepted            int
	EvidenceBoardCitationScopeActive   bool
	NativeDeltaZeroRequired            bool
	NativeWriteLocked                  bool
	NativeWriteAuthorization           bool
	NativeRegistryWrite                bool
	ClosureOnly                        bool

	Verdict  string
	Reason   string
	Failures []string
}

type Firewall struct {
	Executed bool

	PhysicalSchwingerFunctionsLoaded bool
	PhysicalOSCertificateLoaded      bool
	PhysicalWickMapLoaded            bool
	PhysicalHilbertSpaceLoaded       bool
	PhysicalHamiltonianLoaded        bool
	UnitaryDynamicsLoaded            bool
	GlobalCausalityLoaded            bool
	TimeArrowLoaded                  bool
	ReleasedBridgeEvidence           bool
	BoardedBridgeEvidence            bool
	NativeSchwingerFunctionWrite     bool
	NativeOSPositivityWrite          bool
	NativeWickWrite                  bool
	NativeHilbertWrite               bool
	NativeHamiltonianWrite           bool
	NativeUnitaryDynamicsWrite       bool
	NativeGlobalCausalWrite          bool
	NativeTimeArrowWrite             bool
	NativeEvidenceBoardWrite         bool
	NativeRegistryWritten            bool

	Verdict string
	Reason  string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Closure     ClosureLedger
	Guard       SectorGuard
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
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
	g550, err := generation2syntheticevidenceboardadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate550 synthetic evidence-board adapter: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g550)}
	a.Closure = buildClosureLedger()
	a.Guard = buildGuard(a.Inheritance, a.Closure)
	a.Firewall = buildFirewall(a.Guard)
	a.Registry = buildRegistry()
	a.Next = buildNext()
	a.Truth = truth()
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2syntheticevidenceboardadapter.Analysis) Inheritance {
	noPhysical := !g.Firewall.PhysicalSchwingerFunctionsLoaded && !g.Firewall.OSPositivityCertificateLoaded && !g.Firewall.WickMapLoaded && !g.Firewall.HilbertSpaceReconstructed && !g.Firewall.HamiltonianSpectrumLoaded && !g.Firewall.UnitaryDynamicsLoaded && !g.Firewall.GlobalCausalityLoaded && !g.Firewall.TimeArrowLoaded
	return Inheritance{
		Executed:                     true,
		Gate550ManifestParsed:        g.Import.Loaded && g.Import.ManifestImported,
		Gate550RowsAccepted:          g.Import.AcceptedRows,
		Gate550ChecksumVerified:      g.Import.ChecksumVerified,
		Gate550CitationParsed:        g.Board.CitationScopeParsed,
		Gate550UncertaintyParsed:     g.Board.UncertaintyBudgetParsed,
		Gate550ReproducibilityParsed: g.Board.ReproducibilityParsed,
		Gate550RevocationParsed:      g.Board.RevocationHooksParsed,
		Gate550VersionedIndexParsed:  g.Board.VersionedIndexParsed,
		Gate550NativeDeltaZero:       g.Board.NativeDeltaZero,
		Gate550SyntheticBlocked:      g.Board.BlockedBecauseSynthetic && !g.Board.AcceptanceAllowed && !g.Board.BoardedAsBridgeEvidence,
		Gate550NoBridgeEvidence:      g.Import.EntriesAccepted == 0 && !g.Import.BoardedAsBridgeEvidence && !g.Firewall.BridgeEvidenceBoarded,
		Gate550NoRealSource:          !g.Firewall.RealSchwingerSourceImported && !g.Firewall.AuthenticatedRealSource,
		Gate550NativeWriteLocked:     g.Board.NativeWriteLocked && !g.Board.NativeWriteAuthorization && !g.Firewall.NativeRegistryWritten,
		Gate550NoPhysicalClaims:      noPhysical,
		Gate550RedirectsToGate551:    g.Next.Gate == 551,
		Verdict:                      StatusGate550SyntheticEvidenceInherited,
		Reason:                       "Gate551 inherits Gate550's checksum-verified synthetic evidence-board parser, governance metadata, zero-native-delta check, synthetic board rejection, no bridge evidence, no real source, and native-write lock.",
	}
}

func buildClosureLedger() ClosureLedger {
	rows := []FrontierRow{
		{Gate: 536, Block: "physical Schwinger source schema", Native: "no", Bridge: "source-ledger airlock", Environment: "actual S_n family and constructive measure", Firewall: "schema is not correlation physics"},
		{Gate: 537, Block: "synthetic Schwinger parser", Native: "no", Bridge: "synthetic 19-row parser", Environment: "none imported", Firewall: "fake S_n cannot become physical correlators"},
		{Gate: 538, Block: "source authenticity schema", Native: "no", Bridge: "13-row provenance sieve", Environment: "non-synthetic source identity", Firewall: "authenticity schema is not source authenticity"},
		{Gate: 539, Block: "synthetic authenticity adapter", Native: "no", Bridge: "checksum/provenance parser", Environment: "none imported", Firewall: "synthetic source rejected as physical source"},
		{Gate: 540, Block: "real import switch", Native: "no", Bridge: "default-off real-source switch", Environment: "operator intent, access grant, checksum", Firewall: "no source import by default"},
		{Gate: 541, Block: "real-looking negative control", Native: "no", Bridge: "default-deny rejection proof", Environment: "unauthenticated source rejected", Firewall: "real-looking metadata is not trusted source data"},
		{Gate: 542, Block: "authorization manifest schema", Native: "no", Bridge: "14-row authorization airlock", Environment: "operator signature and review", Firewall: "authorization schema does not run comparator"},
		{Gate: 543, Block: "synthetic authorization adapter", Native: "no", Bridge: "quarantined dry-run authorization", Environment: "none imported", Firewall: "synthetic authorization cannot authorize real import"},
		{Gate: 544, Block: "comparator execution harness", Native: "no", Bridge: "16-row comparator contract", Environment: "authenticated non-synthetic source", Firewall: "harness does not execute physics"},
		{Gate: 545, Block: "synthetic comparator output", Native: "no", Bridge: "quarantine result plumbing", Environment: "none imported", Firewall: "synthetic output remains quarantined"},
		{Gate: 546, Block: "release-review airlock", Native: "no", Bridge: "15-row release-review schema", Environment: "human review and reproducibility", Firewall: "release schema is not bridge evidence"},
		{Gate: 547, Block: "synthetic release-review adapter", Native: "no", Bridge: "release-review parser", Environment: "none imported", Firewall: "synthetic output cannot be released"},
		{Gate: 548, Block: "physical-correlation import/release closure", Native: "no", Bridge: "sector closure ledger", Environment: "released bridge evidence absent", Firewall: "closure ledger does not derive dynamics"},
		{Gate: 549, Block: "physical-correlation evidence board", Native: "no", Bridge: "17-row citation/governance airlock", Environment: "future released bridge evidence", Firewall: "evidence board is not native law"},
		{Gate: 550, Block: "synthetic evidence-board adapter", Native: "no", Bridge: "checksum-governed synthetic board parser", Environment: "none imported", Firewall: "synthetic board entry rejected"},
	}
	c := ClosureLedger{Executed: true, Rows: rows, RowCount: len(rows), NativeFrontierFrozen: true, BridgeFrontierMapped: true, EnvironmentalMapped: true, SchwingerSourceClosed: true, SourceAuthenticityClosed: true, RealImportSwitchClosed: true, AuthorizationClosed: true, ComparatorHarnessClosed: true, ReleaseReviewClosed: true, EvidenceBoardClosed: true, Verdict: StatusClosureLedgerEmitted, Reason: "Gates 536-550 are closed as a physical-correlation evidence-board frontier ledger: every source, comparator, release, and evidence-board lane remains bridge-only until authenticated non-synthetic source chains and review exist; even then native writes stay forbidden."}
	for _, r := range rows {
		if r.Native != "" {
			c.NativeRows++
		}
		if r.Bridge != "" {
			c.BridgeRows++
		}
		if r.Environment != "" {
			c.EnvironmentalRows++
		}
		if r.Firewall != "" {
			c.FailedRouteRows++
		}
	}
	return c
}

func buildGuard(i Inheritance, c ClosureLedger) SectorGuard {
	g := SectorGuard{Executed: true,
		PhysicalSchwingerFunctionsImported: false,
		AuthenticatedNonSyntheticSource:    false,
		SourceAuthenticityAccepted:         false,
		RealImportSwitchEnabled:            false,
		OperatorIntentForRealImport:        false,
		ComparatorAuthorized:               false,
		ComparatorExecutedOnRealSource:     false,
		ComparatorOutputReleased:           false,
		BridgeEvidenceReleased:             false,
		EvidenceBoardManifestImported:      false,
		EvidenceBoardEntryAccepted:         false,
		EvidenceEntriesAccepted:            0,
		EvidenceBoardCitationScopeActive:   false,
		NativeDeltaZeroRequired:            true,
		NativeWriteLocked:                  true,
		NativeWriteAuthorization:           false,
		NativeRegistryWrite:                false,
		ClosureOnly:                        true,
		Verdict:                            StatusNoBridgeEvidenceBoarded,
		Reason:                             "Gate551 is a closure ledger only: no non-synthetic source, released comparator output, release review, evidence-board entry, or native registry write exists.",
		Failures:                           []string{StatusFailedClosureNoRealSource, StatusFailedClosureNoBoardEntry, StatusFailedClosureEvidenceBoardNotNative},
	}
	if !i.Gate550SyntheticBlocked || !i.Gate550NoBridgeEvidence || !i.Gate550NoRealSource || !i.Gate550NativeWriteLocked || !c.EvidenceBoardClosed {
		g.Verdict = "FAILED_ROUTE_GATE551_CLOSURE_INPUTS_INCOMPLETE"
		g.Reason = "Gate551 cannot close the evidence-board sector unless Gate550 synthetic rejection, no-bridge-evidence, no-real-source, and native lock all hold."
	}
	return g
}

func buildFirewall(g SectorGuard) Firewall {
	return Firewall{Executed: true,
		PhysicalSchwingerFunctionsLoaded: false,
		PhysicalOSCertificateLoaded:      false,
		PhysicalWickMapLoaded:            false,
		PhysicalHilbertSpaceLoaded:       false,
		PhysicalHamiltonianLoaded:        false,
		UnitaryDynamicsLoaded:            false,
		GlobalCausalityLoaded:            false,
		TimeArrowLoaded:                  false,
		ReleasedBridgeEvidence:           g.BridgeEvidenceReleased,
		BoardedBridgeEvidence:            g.EvidenceBoardEntryAccepted,
		NativeSchwingerFunctionWrite:     false,
		NativeOSPositivityWrite:          false,
		NativeWickWrite:                  false,
		NativeHilbertWrite:               false,
		NativeHamiltonianWrite:           false,
		NativeUnitaryDynamicsWrite:       false,
		NativeGlobalCausalWrite:          false,
		NativeTimeArrowWrite:             false,
		NativeEvidenceBoardWrite:         g.NativeRegistryWrite,
		NativeRegistryWritten:            g.NativeRegistryWrite,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "Gate551 freezes the evidence-board sector as bridge-only; no physical correlation, OS/Wick/Hilbert/Hamiltonian/dynamics/causality/time-arrow, evidence-board, or registry write is admitted natively.",
	}
}

func buildRegistry() RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No native law is written at Gate551.", "The evidence-board closure ledger records zero native delta for the full physical-correlation import/release/evidence-board sector."},
		BridgeEntries:        []string{"Gate551 closes Gates 536-550 as a bridge-only evidence pipeline: source schema, parsers, authenticity, switch, authorization, comparator harness, release review, evidence board, and synthetic rejection.", "Evidence may be cited only after authenticated non-synthetic source chain, comparator execution, release review, evidence-board governance, revocation hooks, and zero-native-write proof."},
		EnvironmentalEntries: []string{"Physical Schwinger functions, constructive measures, OS certificates, Wick maps, Hamiltonian spectra, reproducibility reports, uncertainty budgets, and real source chains remain environmental/bridge inputs.", "No released bridge evidence or board entry exists in Gate551."},
		FailedRoutes:         []string{StatusFailedClosureDoesNotDeriveSchwinger, StatusFailedClosureDoesNotProveOS, StatusFailedClosureDoesNotGrantWick, StatusFailedClosureDoesNotSelectHilbert, StatusFailedClosureDoesNotDeriveHamilton, StatusFailedClosureDoesNotGrantUnitary, StatusFailedClosureDoesNotGrantGlobal, StatusFailedClosureDoesNotSelectArrow, StatusFailedClosureNoRealSource, StatusFailedClosureNoBoardEntry, StatusFailedClosureEvidenceBoardNotNative, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"A future universal environmental ledger may summarize all non-native choices across flavor, coupling scales, topology, dimensionality, Wick/Hilbert/OS, Schwinger sources, and evidence boards without promoting them to native ASHA law."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 552, Title: "Universal Environmental Boundary Ledger Airlock", Reason: "Gate551 closes the physical-correlation evidence-board sector. The next safe gate is a master environmental boundary ledger that cross-indexes all quarantined universe-specific inputs without relaxing any native firewall.", PrimaryTask: "Define a unified, bridge-only/environmental ledger for flavor data, coupling scales, topology, dimensional split, Wick/Hilbert/OS structures, Schwinger sources, comparator evidence, and evidence-board entries, with mandatory zero-native-delta governance."}
}

func truth() string {
	return "Gate551 closes the physical-correlation evidence-board sector: Gates 536-550 now form a complete bridge-only pipeline from source schema to evidence-board rejection, with no authenticated non-synthetic source, no released bridge evidence, no board entry, and no native physical dynamics theorem."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate550ManifestParsed || a.Inheritance.Gate550RowsAccepted != 17 || !a.Inheritance.Gate550ChecksumVerified || !a.Inheritance.Gate550CitationParsed || !a.Inheritance.Gate550UncertaintyParsed || !a.Inheritance.Gate550ReproducibilityParsed || !a.Inheritance.Gate550RevocationParsed || !a.Inheritance.Gate550VersionedIndexParsed || !a.Inheritance.Gate550NativeDeltaZero || !a.Inheritance.Gate550SyntheticBlocked || !a.Inheritance.Gate550NoBridgeEvidence || !a.Inheritance.Gate550NoRealSource || !a.Inheritance.Gate550NativeWriteLocked || !a.Inheritance.Gate550NoPhysicalClaims || !a.Inheritance.Gate550RedirectsToGate551 {
		bad = append(bad, "Gate550 inheritance incomplete")
	}
	if !a.Closure.Executed || a.Closure.RowCount != 15 || !a.Closure.NativeFrontierFrozen || !a.Closure.BridgeFrontierMapped || !a.Closure.EnvironmentalMapped || !a.Closure.SchwingerSourceClosed || !a.Closure.SourceAuthenticityClosed || !a.Closure.RealImportSwitchClosed || !a.Closure.AuthorizationClosed || !a.Closure.ComparatorHarnessClosed || !a.Closure.ReleaseReviewClosed || !a.Closure.EvidenceBoardClosed {
		bad = append(bad, "closure ledger incomplete")
	}
	if a.Guard.PhysicalSchwingerFunctionsImported || a.Guard.AuthenticatedNonSyntheticSource || a.Guard.SourceAuthenticityAccepted || a.Guard.RealImportSwitchEnabled || a.Guard.OperatorIntentForRealImport || a.Guard.ComparatorAuthorized || a.Guard.ComparatorExecutedOnRealSource || a.Guard.ComparatorOutputReleased || a.Guard.BridgeEvidenceReleased || a.Guard.EvidenceBoardManifestImported || a.Guard.EvidenceBoardEntryAccepted || a.Guard.EvidenceEntriesAccepted != 0 || a.Guard.EvidenceBoardCitationScopeActive || !a.Guard.NativeDeltaZeroRequired || !a.Guard.NativeWriteLocked || a.Guard.NativeWriteAuthorization || a.Guard.NativeRegistryWrite || !a.Guard.ClosureOnly {
		bad = append(bad, "sector guard leaked")
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidence || a.Firewall.BoardedBridgeEvidence || a.Firewall.NativeEvidenceBoardWrite || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate551 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: manifest=%t rows=%d checksum=%t citation=%t uncertainty=%t reproducibility=%t revocation=%t version=%t delta_zero=%t synthetic_blocked=%t no_bridge=%t no_real=%t native_locked=%t no_physical=%t redirects=%t; %s", x.Verdict, x.Gate550ManifestParsed, x.Gate550RowsAccepted, x.Gate550ChecksumVerified, x.Gate550CitationParsed, x.Gate550UncertaintyParsed, x.Gate550ReproducibilityParsed, x.Gate550RevocationParsed, x.Gate550VersionedIndexParsed, x.Gate550NativeDeltaZero, x.Gate550SyntheticBlocked, x.Gate550NoBridgeEvidence, x.Gate550NoRealSource, x.Gate550NativeWriteLocked, x.Gate550NoPhysicalClaims, x.Gate550RedirectsToGate551, x.Reason)
}

func FormatClosure(x ClosureLedger) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: rows=%d native_rows=%d bridge_rows=%d env_rows=%d failed_rows=%d source=%t authenticity=%t switch=%t authorization=%t comparator=%t release=%t board=%t native_frozen=%t bridge_mapped=%t env_mapped=%t; %s", StatusClosureLedgerEmitted, StatusNativeFrontierFrozen, StatusBridgeFrontierMapped, StatusEnvironmentalFrontierMapped, StatusSchwingerSourceBlockClosed, StatusSourceAuthenticityBlockClosed, StatusRealImportSwitchBlockClosed, StatusAuthorizationBlockClosed, StatusComparatorHarnessBlockClosed, StatusReleaseReviewBlockClosed, StatusEvidenceBoardBlockClosed, StatusClosureFirewallMatrixComplete, x.RowCount, x.NativeRows, x.BridgeRows, x.EnvironmentalRows, x.FailedRouteRows, x.SchwingerSourceClosed, x.SourceAuthenticityClosed, x.RealImportSwitchClosed, x.AuthorizationClosed, x.ComparatorHarnessClosed, x.ReleaseReviewClosed, x.EvidenceBoardClosed, x.NativeFrontierFrozen, x.BridgeFrontierMapped, x.EnvironmentalMapped, x.Reason)
}

func FormatGuard(x SectorGuard) string {
	return fmt.Sprintf("%s;%s;%s;%s: schwinger=%t auth_source=%t authenticity=%t switch=%t intent=%t comparator_auth=%t comparator_real=%t output_released=%t bridge_released=%t board_manifest=%t board_entry=%t entries=%d citation_active=%t delta_required=%t native_lock=%t native_auth=%t registry=%t closure_only=%t failures=%s; %s", x.Verdict, StatusFailedClosureNoRealSource, StatusFailedClosureNoBoardEntry, StatusFailedClosureEvidenceBoardNotNative, x.PhysicalSchwingerFunctionsImported, x.AuthenticatedNonSyntheticSource, x.SourceAuthenticityAccepted, x.RealImportSwitchEnabled, x.OperatorIntentForRealImport, x.ComparatorAuthorized, x.ComparatorExecutedOnRealSource, x.ComparatorOutputReleased, x.BridgeEvidenceReleased, x.EvidenceBoardManifestImported, x.EvidenceBoardEntryAccepted, x.EvidenceEntriesAccepted, x.EvidenceBoardCitationScopeActive, x.NativeDeltaZeroRequired, x.NativeWriteLocked, x.NativeWriteAuthorization, x.NativeRegistryWrite, x.ClosureOnly, strings.Join(x.Failures, ","), x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: schwinger=%t os=%t wick=%t hilbert=%t ham=%t unitary=%t global=%t arrow=%t released=%t boarded=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t native_board=%t registry=%t; %s", StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedClosureDoesNotDeriveSchwinger, StatusFailedClosureDoesNotProveOS, StatusFailedClosureDoesNotGrantWick, StatusFailedClosureDoesNotSelectHilbert, StatusFailedClosureDoesNotDeriveHamilton, StatusFailedClosureDoesNotGrantUnitary, StatusFailedClosureDoesNotGrantGlobal, StatusFailedClosureDoesNotSelectArrow, StatusFailedClosureNoRealSource, StatusFailedClosureNoBoardEntry, StatusFailedClosureEvidenceBoardNotNative, x.PhysicalSchwingerFunctionsLoaded, x.PhysicalOSCertificateLoaded, x.PhysicalWickMapLoaded, x.PhysicalHilbertSpaceLoaded, x.PhysicalHamiltonianLoaded, x.UnitaryDynamicsLoaded, x.GlobalCausalityLoaded, x.TimeArrowLoaded, x.ReleasedBridgeEvidence, x.BoardedBridgeEvidence, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeEvidenceBoardWrite, x.NativeRegistryWritten, x.Reason)
}

func Statuses() []string {
	return []string{StatusGate550SyntheticEvidenceInherited, StatusClosureLedgerEmitted, StatusNativeFrontierFrozen, StatusBridgeFrontierMapped, StatusEnvironmentalFrontierMapped, StatusSchwingerSourceBlockClosed, StatusSourceAuthenticityBlockClosed, StatusRealImportSwitchBlockClosed, StatusAuthorizationBlockClosed, StatusComparatorHarnessBlockClosed, StatusReleaseReviewBlockClosed, StatusEvidenceBoardBlockClosed, StatusNoBridgeEvidenceBoarded, StatusClosureFirewallMatrixComplete, StatusFailedClosureDoesNotDeriveSchwinger, StatusFailedClosureDoesNotProveOS, StatusFailedClosureDoesNotGrantWick, StatusFailedClosureDoesNotSelectHilbert, StatusFailedClosureDoesNotDeriveHamilton, StatusFailedClosureDoesNotGrantUnitary, StatusFailedClosureDoesNotGrantGlobal, StatusFailedClosureDoesNotSelectArrow, StatusFailedClosureNoRealSource, StatusFailedClosureNoBoardEntry, StatusFailedClosureEvidenceBoardNotNative, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
