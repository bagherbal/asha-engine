// Package generation2nullbaselineperturbation implements Gate 481:
// Null-Baseline Perturbation Ledger / Sector Transport Audit.
//
// Gate 480 derived a conditional C\ell(1,7) null-vacuum baseline
// alpha_vac=1 and I_K,vac=1/2. Gate 481 tests whether that baseline can be
// legally transported into physical quark/lepton sector coordinates. The key
// invariant result is that a shared baseline cancels in relative cylinder
// distances: only sector perturbations delta_alpha and delta_phi remain. Thus
// the null baseline can organize bridge ledgers, but it cannot by itself compute
// CKM/PMNS residuals or replace sector-specific I_K comparators.
package generation2nullbaselineperturbation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE481-NULL-BASELINE-PERTURBATION-LEDGER-SECTOR-TRANSPORT"

	StatusGate480Inherited     = "CONDITIONAL_SUPPORT_GATE480_NULL_BASELINE_INHERITED"
	StatusLedgerDefined        = "CONDITIONAL_SUPPORT_NULL_BASELINE_PERTURBATION_LEDGER_DEFINED"
	StatusBaselineCancellation = "CONDITIONAL_SUPPORT_SHARED_NULL_BASELINE_CANCELS_IN_RELATIVE_DISTANCE"
	StatusSyntheticTransportOK = "CONDITIONAL_SUPPORT_SYNTHETIC_PERTURBATION_TRANSPORT_BRIDGE_ONLY_VALIDATED"
	StatusFirewallPreserved    = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE481_TRANSPORT_AUDIT"

	StatusFailedTransportNotNative    = "FAILED_ROUTE_NULL_TO_SECTOR_TRANSPORT_NOT_NATIVE"
	StatusFailedPerturbationsUnforced = "FAILED_ROUTE_NULL_BASELINE_DOES_NOT_FIX_SECTOR_PERTURBATIONS"
	StatusFailedIKVacAsSectorIK       = "FAILED_ROUTE_I_K_VAC_HALF_CANNOT_REPLACE_SECTOR_I_K"
	StatusFailedCKMPMNSPrediction     = "FAILED_ROUTE_NULL_BASELINE_TRANSPORT_AS_CKM_PMNS_PREDICTION_REJECTED"
	StatusFailedNativePromotion       = "FAILED_ROUTE_NULL_PERTURBATION_LEDGER_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                 bool
	Gate480NullBaseline      bool
	AlphaVac                 float64
	IKVac                    float64
	Gate480SectorCoordsOpen  bool
	Gate480FirewallPreserved bool
	NativeRegistryClean      bool
	Verdict                  string
}

type TransportLaw struct {
	Executed                  bool
	BaselineAlpha             float64
	BaselinePhiGauge          float64
	FormulaAlpha              string
	FormulaPhi                string
	SharedBaselineAssumed     bool
	TransportPreviouslyNative bool
	PerturbationsNative       bool
	IKVacCanReplaceSectorIK   bool
	Verdict                   string
	Reason                    string
	Failures                  []string
}

type PerturbationRow struct {
	Sector     string
	DeltaAlpha float64
	DeltaPhi   float64
	Alpha      float64
	Phi        float64
	IK         float64
	BridgeOnly bool
	Synthetic  bool
	Provenance string
}

type PairDistance struct {
	Name              string
	LeftSector        string
	RightSector       string
	DeltaAlpha        float64
	DeltaPhi          float64
	Distance          float64
	BaselineCancelled bool
	Computed          bool
	BridgeOnly        bool
	NativePrediction  bool
	Verdict           string
	Reason            string
}

type PerturbationLedger struct {
	Executed               bool
	Rows                   []PerturbationRow
	Pairs                  []PairDistance
	AllRowsBridgeOnly      bool
	AllRowsSynthetic       bool
	QuarkDistanceComputed  bool
	LeptonDistanceComputed bool
	Verdict                string
	Reason                 string
}

type CancellationProof struct {
	Executed                bool
	DistanceFormula         string
	Substitution            string
	ReducedFormula          string
	BaselineAlphaCancels    bool
	BaselinePhiCancels      bool
	OnlyPerturbationsRemain bool
	Verdict                 string
	Reason                  string
}

type Firewall struct {
	Executed                         bool
	ObservedMassImported             bool
	CKMImported                      bool
	PMNSImported                     bool
	VacuumIKNativeBaseline           bool
	VacuumIKPhysicalSectorCoordinate bool
	SectorIKSolvedByBaseline         bool
	PerturbationsNative              bool
	DUDNativePrediction              bool
	DENuNativePrediction             bool
	CKMMatrixConstructed             bool
	PMNSMatrixConstructed            bool
	NativeRegistryWritten            bool
	NativeFlavorDimAfter             int
	KXYCoeffDimAfter                 int
	Verdict                          string
	Reason                           string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Transport   TransportLaw
	Proof       CancellationProof
	Ledger      PerturbationLedger
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
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	a := Analysis{Inheritance: buildInheritance()}
	a.Transport = buildTransport()
	a.Proof = buildProof()
	a.Ledger = buildLedger(a)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate480NullBaseline: true, AlphaVac: 1, IKVac: 0.5, Gate480SectorCoordsOpen: true, Gate480FirewallPreserved: true, NativeRegistryClean: true, Verdict: StatusGate480Inherited}
}

func buildTransport() TransportLaw {
	return TransportLaw{
		Executed:                  true,
		BaselineAlpha:             1,
		BaselinePhiGauge:          0,
		FormulaAlpha:              "alpha_s = alpha_vac + delta_alpha_s",
		FormulaPhi:                "phi_s = phi_vac + delta_phi_s",
		SharedBaselineAssumed:     true,
		TransportPreviouslyNative: false,
		PerturbationsNative:       false,
		IKVacCanReplaceSectorIK:   false,
		Verdict:                   StatusLedgerDefined,
		Reason:                    "Gate481 may define a bridge-only perturbation chart around the Gate480 null vacuum, but no prior native theorem transports the same baseline into physical sector coordinates or fixes the sector perturbations.",
		Failures:                  []string{StatusFailedTransportNotNative, StatusFailedPerturbationsUnforced, StatusFailedIKVacAsSectorIK},
	}
}

func buildProof() CancellationProof {
	return CancellationProof{
		Executed:                true,
		DistanceFormula:         "d_st=sqrt((alpha_t-alpha_s)^2+4 sin^2((phi_t-phi_s)/2))",
		Substitution:            "alpha_s=alpha_vac+delta_alpha_s, phi_s=phi_vac+delta_phi_s",
		ReducedFormula:          "d_st=sqrt((delta_alpha_t-delta_alpha_s)^2+4 sin^2((delta_phi_t-delta_phi_s)/2))",
		BaselineAlphaCancels:    true,
		BaselinePhiCancels:      true,
		OnlyPerturbationsRemain: true,
		Verdict:                 StatusBaselineCancellation,
		Reason:                  "a universal null baseline is common-mode data; relative quark/lepton mixing distances are sensitive only to differences of sector perturbations.",
	}
}

func buildLedger(a Analysis) PerturbationLedger {
	rows := []PerturbationRow{
		row("u", 0.020, 0.100),
		row("d", -0.030, -0.120),
		row("e", -0.050, 0.200),
		row("nu", 0.040, 0.620),
	}
	pairs := []PairDistance{
		pair("synthetic quark null-baseline residual", rows[0], rows[1]),
		pair("synthetic lepton null-baseline residual", rows[2], rows[3]),
	}
	allBridge, allSynthetic := true, true
	for _, r := range rows {
		allBridge = allBridge && r.BridgeOnly
		allSynthetic = allSynthetic && r.Synthetic
	}
	return PerturbationLedger{Executed: true, Rows: rows, Pairs: pairs, AllRowsBridgeOnly: allBridge, AllRowsSynthetic: allSynthetic, QuarkDistanceComputed: pairs[0].Computed, LeptonDistanceComputed: pairs[1].Computed, Verdict: StatusSyntheticTransportOK, Reason: "synthetic perturbation rows demonstrate the chart mechanics while remaining bridge-only; their distances are diagnostics, not CKM/PMNS predictions"}
}

func row(sector string, da, dp float64) PerturbationRow {
	alpha := 1 + da
	phi := dp
	return PerturbationRow{Sector: sector, DeltaAlpha: da, DeltaPhi: dp, Alpha: alpha, Phi: phi, IK: IK(alpha), BridgeOnly: true, Synthetic: true, Provenance: "Gate481 synthetic perturbation ledger"}
}

func pair(name string, l, r PerturbationRow) PairDistance {
	da := r.DeltaAlpha - l.DeltaAlpha
	dp := wrapPi(r.DeltaPhi - l.DeltaPhi)
	d := math.Sqrt(da*da + 4*math.Pow(math.Sin(dp/2), 2))
	return PairDistance{Name: name, LeftSector: l.Sector, RightSector: r.Sector, DeltaAlpha: da, DeltaPhi: dp, Distance: d, BaselineCancelled: true, Computed: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusSyntheticTransportOK, Reason: "computed from synthetic perturbation deltas after common null baseline cancellation"}
}

func IK(alpha float64) float64 { return alpha / math.Sqrt(alpha*alpha+3) }
func wrapPi(x float64) float64 {
	for x <= -math.Pi {
		x += 2 * math.Pi
	}
	for x > math.Pi {
		x -= 2 * math.Pi
	}
	return x
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, ObservedMassImported: false, CKMImported: false, PMNSImported: false, VacuumIKNativeBaseline: true, VacuumIKPhysicalSectorCoordinate: false, SectorIKSolvedByBaseline: false, PerturbationsNative: false, DUDNativePrediction: false, DENuNativePrediction: false, CKMMatrixConstructed: false, PMNSMatrixConstructed: false, NativeRegistryWritten: false, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate481 keeps I_K,vac=1/2 as a common vacuum baseline; sector perturbations, CKM/PMNS residuals, and physical I_K values remain bridge/firewalled"}
}

func buildNext() NextStep {
	return NextStep{Gate: 482, Title: "Null-baseline sector deformation source search", Reason: "Gate481 shows a shared null baseline cancels from relative distances, leaving sector perturbations unfixed.", PrimaryTask: "audit whether finite algebraic orientation, chirality, or Higgs-edge operators can natively source the sector perturbations delta_alpha and delta_phi without observed CKM/PMNS data"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullBaseline || a.Inheritance.AlphaVac != 1 || a.Inheritance.IKVac != 0.5 || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate481 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Transport.Executed || !a.Transport.SharedBaselineAssumed || a.Transport.TransportPreviouslyNative || a.Transport.PerturbationsNative || a.Transport.IKVacCanReplaceSectorIK {
		return fmt.Errorf("Gate481 transport invalid: %+v", a.Transport)
	}
	if !a.Proof.Executed || !a.Proof.BaselineAlphaCancels || !a.Proof.BaselinePhiCancels || !a.Proof.OnlyPerturbationsRemain {
		return fmt.Errorf("Gate481 cancellation proof invalid: %+v", a.Proof)
	}
	if !a.Ledger.Executed || !a.Ledger.AllRowsBridgeOnly || !a.Ledger.AllRowsSynthetic || !a.Ledger.QuarkDistanceComputed || !a.Ledger.LeptonDistanceComputed {
		return fmt.Errorf("Gate481 synthetic ledger invalid: %+v", a.Ledger)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassImported || a.Firewall.CKMImported || a.Firewall.PMNSImported || a.Firewall.VacuumIKPhysicalSectorCoordinate || a.Firewall.SectorIKSolvedByBaseline || a.Firewall.PerturbationsNative || a.Firewall.DUDNativePrediction || a.Firewall.DENuNativePrediction || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate481 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate481 result: I_K,vac=%.12g remains a common null-vacuum baseline; common baseline terms cancel from relative distances, leaving only bridge-only perturbations. Synthetic d_ud=%.12g and d_eν=%.12g are diagnostics, not CKM/PMNS predictions.", a.Inheritance.IKVac, a.Ledger.Pairs[0].Distance, a.Ledger.Pairs[1].Distance)
}

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 481 Registry Audit — Null-Baseline Perturbation Ledger / Sector Transport Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusLedgerDefined + "`\n\n")
	b.WriteString("Gate 481 accepts the Gate 480 result `alpha_vac=1`, `I_K,vac=0.5` only as a **common null-vacuum baseline**. It then asks whether this baseline can be transported into physical sector coordinates. The answer is negative at the native-law level: the transport chart can be defined, but the sector perturbations remain bridge-only.\n\n")
	b.WriteString("## Transport chart\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha_s = alpha_vac + delta_alpha_s\n")
	b.WriteString("phi_s   = phi_vac   + delta_phi_s\n")
	b.WriteString("alpha_vac = 1\n")
	b.WriteString("I_K,vac = 1/2\n")
	b.WriteString("```\n\n")
	b.WriteString("The chart is useful, but it is not a sector-coordinate theorem. No native rule in the current atlas fixes `delta_alpha_s` or `delta_phi_s`.\n\n")
	b.WriteString("## Baseline cancellation proof\n\n")
	b.WriteString("Starting from the cylinder distance:\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Proof.DistanceFormula + "\n")
	b.WriteString(a.Proof.Substitution + "\n")
	b.WriteString(a.Proof.ReducedFormula + "\n")
	b.WriteString("```\n\n")
	b.WriteString("Because `alpha_vac` and `phi_vac` are common-mode baseline data, they cancel. Relative mixing diagnostics are controlled by sector perturbations, not by the null baseline itself.\n\n")
	b.WriteString("## Synthetic bridge-only dry run\n\n")
	b.WriteString("| sector | delta_alpha | delta_phi | alpha | phi | I_K | bridge_only |\n|---|---:|---:|---:|---:|---:|---|\n")
	for _, r := range a.Ledger.Rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %t |\n", r.Sector, fmtFloat(r.DeltaAlpha), fmtFloat(r.DeltaPhi), fmtFloat(r.Alpha), fmtFloat(r.Phi), fmtFloat(r.IK), r.BridgeOnly))
	}
	b.WriteString("\n| pair | delta_alpha | delta_phi | distance | native_prediction |\n|---|---:|---:|---:|---|\n")
	for _, p := range a.Ledger.Pairs {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %t |\n", p.Name, fmtFloat(p.DeltaAlpha), fmtFloat(p.DeltaPhi), fmtFloat(p.Distance), p.NativePrediction))
	}
	b.WriteString("\nThese values are synthetic diagnostics proving the socket mechanics only. They are not CKM/PMNS predictions.\n\n")
	b.WriteString("## Rejected promotions\n\n```text\n")
	for _, s := range []string{StatusFailedTransportNotNative, StatusFailedPerturbationsUnforced, StatusFailedIKVacAsSectorIK, StatusFailedCKMPMNSPrediction, StatusFailedNativePromotion} {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Numerical output\n\n```text\n")
	b.WriteString(fmt.Sprintf("I_K,vac = %.12f\n", a.Inheritance.IKVac))
	b.WriteString(fmt.Sprintf("synthetic d_ud  = %.12f\n", a.Ledger.Pairs[0].Distance))
	b.WriteString(fmt.Sprintf("synthetic d_eν  = %.12f\n", a.Ledger.Pairs[1].Distance))
	b.WriteString("physical d_ud  = undefined\n")
	b.WriteString("physical d_eν  = undefined\n")
	b.WriteString("CKM/PMNS       = not constructed\n")
	b.WriteString("```\n\n")
	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s: %s\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	return b.String()
}
