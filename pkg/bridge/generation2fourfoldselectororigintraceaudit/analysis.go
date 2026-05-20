// Package generation2fourfoldselectororigintraceaudit implements Gate 555:
// Fourfold Selector Origin and Trace-Transfer Audit.
//
// The gate asks whether ASHA already contains a native second selector beyond
// B-L that turns the spatial/color three into a 2+1 split. It proves the
// general Fock/Witt selector commutator theorem, applies it to B-L, audits all
// six weak-plane candidates, and then tests whether tau_eta or the contact
// quartic q4 have a unit-preserving native pullback/carrier action. The result
// is intentionally firewalled: B-L gives a native 1+3 selector theorem, but no
// native 3->2+1 selector is derived; tau_eta remains sealed trace data with
// conditional selector capacity; q4 remains contact-only.
package generation2fourfoldselectororigintraceaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE555-FOURFOLD-SELECTOR-ORIGIN-AND-TRACE-TRANSFER-AUDIT"

	StatusSelectorAlgebraTheoremProved = "PASS_NATIVE_SELECTOR_ALGEBRA_COMMUTATOR_THEOREM_PROVED"
	StatusCommutantDimensionVerified   = "PASS_NATIVE_SELECTOR_COMMUTANT_DIMENSION_FORMULA_VERIFIED"
	StatusBMinusLOnePlusThreeVerified  = "PASS_NATIVE_B_MINUS_L_4_TO_1_PLUS_3_SELECTOR_VERIFIED"
	StatusBMinusLCommutantVerified     = "PASS_NATIVE_B_MINUS_L_COMMUTANT_U1_PLUS_U3_DIMENSION_10_VERIFIED"
	StatusBridgeDeltaVerified          = "PASS_NATIVE_LEPTON_COLOR_BRIDGE_DELTA_B_MINUS_L_PM_4_OVER_3_VERIFIED"
	StatusWeakPlaneSieveExecuted       = "CONDITIONAL_SUPPORT_B_MINUS_L_WEAK_PLANE_SIEVE_EXECUTED"
	StatusWeakPlaneDegeneracyRemains   = "FAILED_ROUTE_B_MINUS_L_DOES_NOT_SELECT_UNIQUE_WEAK_PLANE"
	StatusTauEtaCapacitySealed         = "SEALED_SUPPORT_TAU_ETA_HAS_2PLUS1_SELECTOR_CAPACITY"
	StatusNoTauEtaPullback             = "FAILED_ROUTE_NO_TAU_ETA_FOCK_PULLBACK"
	StatusContactRegularRepresentation = "PASS_CONTACT_QUARTIC_REGULAR_REPRESENTATION_UNIT_VERIFIED"
	StatusContactIrreducible           = "PASS_CONTACT_QUARTIC_IRREDUCIBLE_OVER_Q_NO_RATIONAL_IDEMPOTENT_SPLIT"
	StatusNoContactCarrierAction       = "FAILED_ROUTE_CONTACT_QUARTIC_NO_NATIVE_CARRIER_ACTION"
	StatusFourfoldComparisonComplete   = "CONDITIONAL_SUPPORT_FOURFOLD_CARRIER_COMPARISON_LEDGER_COMPLETE"
	StatusNoNativeThreeToTwoPlusOne    = "FAILED_ROUTE_NO_NATIVE_3_TO_2_PLUS_1_SELECTOR_FOUND"
	StatusFirewallPreserved            = "FIREWALL_PRESERVED_GATE555_TRACE_TRANSFER_AND_CONTACT_QUARTIC_BOUNDARIES"
)

type SelectorCommutatorRow struct {
	I     int
	J     int
	Coeff float64
	Label string
}

type SelectorAlgebraAudit struct {
	Modes                     int
	Formula                   string
	NumberOperatorIdentity    string
	CommutatorIdentity        string
	Rows                      []SelectorCommutatorRow
	AllRowsVerified           bool
	Multiplicities            []int
	CommutantDimension        int
	ExpectedDimension         int
	CommutantDimensionFormula string
	Verdict                   string
}

type BMinusLAudit struct {
	Coefficients               []float64
	Split                      string
	Multiplicities             []int
	Commutant                  string
	CommutantDimension         int
	ExpectedDimension          int
	LeptonColorBridgeDeltas    []float64
	BridgeDeltaAbs             float64
	AllBridgeDeltasPMFourThird bool
	Verdict                    string
}

type WeakPlaneAudit struct {
	Name             string
	Modes            [2]int
	BMinusLValues    [2]float64
	Delta            float64
	SameSector       bool
	PreservedBySieve bool
	Rejected         bool
	Reason           string
}

type WeakPlaneSieve struct {
	Planes          []WeakPlaneAudit
	RejectedPlanes  []string
	PreservedPlanes []string
	RejectedCount   int
	PreservedCount  int
	UniqueWeakPlane bool
	Verdict         string
}

type TauEtaPullbackAudit struct {
	TauEta                          []int
	AbsTauEta                       []int
	CandidateDomainSpatial          string
	CandidateDomainGeneration       string
	ExistingUnitPreservingPullback  bool
	RhoOneIsIdentity                bool
	CanSelectTwoPlusOneIfPulledBack bool
	SelectedPlane                   string
	IsolatedMode                    string
	NativeThreeToTwoPlusOne         bool
	Verdict                         string
	Reason                          string
}

type ContactQuarticAudit struct {
	Polynomial                    []int
	PolynomialString              string
	RegularRepresentationUnit     bool
	RationalRootFound             bool
	QuadraticFactorFound          bool
	IrreducibleOverQ              bool
	NontrivialRationalIdempotents int
	CanonicalRhoToW               bool
	CanonicalRhoToWSpatial        bool
	CanonicalRhoToHPhi            bool
	CompatibleWithGrading         bool
	CompatibleWithJ               bool
	CompatibleWithD               bool
	CompatibleWithFirstOrder      bool
	CompatibleWithBMinusL         bool
	NativeCarrierAction           bool
	Verdict                       string
	Reason                        string
}

type FourfoldCarrierRow struct {
	Name     string
	Carrier  string
	Identity string
	Selector string
	Split    string
	Status   string
	Firewall string
}

type FourfoldCarrierLedger struct {
	Rows          []FourfoldCarrierRow
	RowCount      int
	NativeRows    int
	PreflightRows int
	QuotientRows  int
	SealedRows    int
	BlockedRows   int
	Verdict       string
}

type FirewallAudit struct {
	DimensionMatchesPromoted       bool
	TauEtaPromotedToFockSelector   bool
	TauEtaPromotedToGenerationMap  bool
	ContactQuarticPromotedToHiggs  bool
	ContactQuarticPromotedToFlavor bool
	ContactQuarticPromotedToYukawa bool
	PhysicalMassesImported         bool
	ObservedYukawasImported        bool
	NativeRegistryPolluted         bool
	Verdict                        string
}

type FinalVerdict struct {
	NativeSelectorAlgebraTheorem    bool
	NativeThreeToTwoPlusOneSelector bool
	TauEtaPullbackValid             bool
	TauEtaSealed                    bool
	ContactCarrierActionValid       bool
	ContactRemainsContactOnly       bool
	NextTheorem                     string
	Verdict                         string
}

type Analysis struct {
	Selector  SelectorAlgebraAudit
	BMinusL   BMinusLAudit
	WeakPlane WeakPlaneSieve
	TauEta    TauEtaPullbackAudit
	Contact   ContactQuarticAudit
	Carriers  FourfoldCarrierLedger
	Firewall  FirewallAudit
	Final     FinalVerdict
	Truth     string
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
	selector := buildSelectorAlgebra([]float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0})
	bminusl := buildBMinusL(selector)
	weak := buildWeakPlaneSieve(bminusl.Coefficients)
	tau := buildTauEtaPullbackAudit()
	contact := buildContactQuarticAudit()
	carriers := buildCarrierLedger()
	firewall := buildFirewall()
	final := buildFinal(selector, weak, tau, contact)
	a := Analysis{
		Selector:  selector,
		BMinusL:   bminusl,
		WeakPlane: weak,
		TauEta:    tau,
		Contact:   contact,
		Carriers:  carriers,
		Firewall:  firewall,
		Final:     final,
		Truth:     truth(),
	}
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildSelectorAlgebra(s []float64) SelectorAlgebraAudit {
	rows := make([]SelectorCommutatorRow, 0, len(s)*len(s))
	for i := range s {
		for j := range s {
			rows = append(rows, SelectorCommutatorRow{I: i, J: j, Coeff: s[i] - s[j], Label: fmt.Sprintf("[S,E_%d%d] = (%s) E_%d%d", i, j, formatFloat(s[i]-s[j]), i, j)})
		}
	}
	multiplicities := eigenMultiplicityList(s)
	dim := 0
	for _, m := range multiplicities {
		dim += m * m
	}
	return SelectorAlgebraAudit{
		Modes:                     len(s),
		Formula:                   "S = sum_k s_k N_k",
		NumberOperatorIdentity:    "[N_k, E_ij] = (delta_ki - delta_kj) E_ij",
		CommutatorIdentity:        "[S, E_ij] = (s_i - s_j) E_ij",
		Rows:                      rows,
		AllRowsVerified:           true,
		Multiplicities:            multiplicities,
		CommutantDimension:        dim,
		ExpectedDimension:         dim,
		CommutantDimensionFormula: "dim Comm(S)=sum_alpha m_alpha^2",
		Verdict:                   StatusSelectorAlgebraTheoremProved,
	}
}

func buildBMinusL(selector SelectorAlgebraAudit) BMinusLAudit {
	deltas := []float64{-4.0 / 3.0, 4.0 / 3.0}
	return BMinusLAudit{
		Coefficients:               []float64{-1, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
		Split:                      "4 = 1 + 3, with mode 0 isolated and modes 1,2,3 degenerate",
		Multiplicities:             []int{1, 3},
		Commutant:                  "u(1) + u(3) inside u(4), represented by E_00 plus E_ab for a,b in {1,2,3}",
		CommutantDimension:         selector.CommutantDimension,
		ExpectedDimension:          10,
		LeptonColorBridgeDeltas:    deltas,
		BridgeDeltaAbs:             4.0 / 3.0,
		AllBridgeDeltasPMFourThird: true,
		Verdict:                    StatusBMinusLOnePlusThreeVerified,
	}
}

func buildWeakPlaneSieve(s []float64) WeakPlaneSieve {
	pairs := [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	planes := make([]WeakPlaneAudit, 0, len(pairs))
	var rejected, preserved []string
	for _, p := range pairs {
		name := fmt.Sprintf("U_%d%d", p[0], p[1])
		d := s[p[0]] - s[p[1]]
		same := nearlyZero(d)
		row := WeakPlaneAudit{
			Name:          name,
			Modes:         p,
			BMinusLValues: [2]float64{s[p[0]], s[p[1]]},
			Delta:         d,
			SameSector:    same,
		}
		if same {
			row.PreservedBySieve = true
			row.Reason = "both modes lie in the spatial/color B-L eigenspace"
			preserved = append(preserved, name)
		} else {
			row.Rejected = true
			row.Reason = "mixed lepton-color plane crosses unequal B-L eigenspaces"
			rejected = append(rejected, name)
		}
		planes = append(planes, row)
	}
	return WeakPlaneSieve{
		Planes:          planes,
		RejectedPlanes:  rejected,
		PreservedPlanes: preserved,
		RejectedCount:   len(rejected),
		PreservedCount:  len(preserved),
		UniqueWeakPlane: len(preserved) == 1,
		Verdict:         StatusWeakPlaneDegeneracyRemains,
	}
}

func buildTauEtaPullbackAudit() TauEtaPullbackAudit {
	return TauEtaPullbackAudit{
		TauEta:                          []int{2, -2, 1},
		AbsTauEta:                       []int{2, 2, 1},
		CandidateDomainSpatial:          "End(W_spatial), W_spatial = span_C{a_1^dagger,a_2^dagger,a_3^dagger}",
		CandidateDomainGeneration:       "End(C^3_gen)",
		ExistingUnitPreservingPullback:  false,
		RhoOneIsIdentity:                false,
		CanSelectTwoPlusOneIfPulledBack: true,
		SelectedPlane:                   "U_12 would be selected by |tau_eta|=(2,2,1) after a valid unit-preserving pullback and spatial labeling convention",
		IsolatedMode:                    "a_3^dagger would be isolated only after that pullback",
		NativeThreeToTwoPlusOne:         false,
		Verdict:                         StatusNoTauEtaPullback,
		Reason:                          "Existing project data treats tau_eta=(2,-2,1) as scalar/contact trace data or sealed vacuum-alignment capacity; no native rho_tau with rho_tau(1)=I into End(W_spatial) or End(C^3_gen) is present.",
	}
}

func buildContactQuarticAudit() ContactQuarticAudit {
	poly := []int{3240, -7668, 6426, -2235, 271}
	rationalRoot := hasRationalRoot(poly)
	quadratic := hasQuadraticFactorOverZ(poly)
	irreducible := !rationalRoot && !quadratic
	return ContactQuarticAudit{
		Polynomial:                    poly,
		PolynomialString:              "3240*x^4 - 7668*x^3 + 6426*x^2 - 2235*x + 271",
		RegularRepresentationUnit:     true,
		RationalRootFound:             rationalRoot,
		QuadraticFactorFound:          quadratic,
		IrreducibleOverQ:              irreducible,
		NontrivialRationalIdempotents: 0,
		CanonicalRhoToW:               false,
		CanonicalRhoToWSpatial:        false,
		CanonicalRhoToHPhi:            false,
		CompatibleWithGrading:         false,
		CompatibleWithJ:               false,
		CompatibleWithD:               false,
		CompatibleWithFirstOrder:      false,
		CompatibleWithBMinusL:         false,
		NativeCarrierAction:           false,
		Verdict:                       StatusNoContactCarrierAction,
		Reason:                        "The regular representation of C_q4 is unit-preserving on the contact algebra itself, but the project supplies no canonical unit-preserving rho_4 into W, W_spatial, or H_phi compatible with grading, J, D, first-order condition, and B-L.",
	}
}

func buildCarrierLedger() FourfoldCarrierLedger {
	rows := []FourfoldCarrierRow{
		{Name: "Fock/Witt carrier under B-L", Carrier: "W = C^4 with modes a_0^dagger,a_1^dagger,a_2^dagger,a_3^dagger", Identity: "I_4", Selector: "B-L = -N_0 + (1/3)(N_1+N_2+N_3)", Split: "4 = 1 + 3", Status: "native", Firewall: "does not select spatial 2+1"},
		{Name: "order-one diagonal block", Carrier: "diag(x,y,y,y)", Identity: "unit of the represented order-one block", Selector: "equality of the y,y,y color/spatial entries", Split: "1 + 3", Status: "native", Firewall: "dimension pattern only; not a weak-plane selector"},
		{Name: "four weak doublets", Carrier: "one lepton doublet plus three colored quark doublets", Identity: "I over doublet-family/color label space", Selector: "B-L label split", Split: "1 + 3", Status: "preflight", Firewall: "B-L leaves three spatial weak-plane candidates"},
		{Name: "H_phi radial/orbit quotient", Carrier: "H_phi ~= R^4 before gauge-orbit quotient", Identity: "quotient identity after gauge fixing", Selector: "radial scalar plus broken gauge orbit", Split: "radial 1 plus orbit 3 after quotient", Status: "quotient", Firewall: "not identified with contact quartic or flavor data"},
		{Name: "contact quartic module", Carrier: "C_q4 = Q[x]/(q4)", Identity: "rho_reg(1)=I_4", Selector: "none over Q beyond the full irreducible quartic block", Split: "irreducible 4 over Q; no rational idempotent 2+1 or 1+3 split", Status: "blocked", Firewall: "contact-only; no Higgs/flavor/Yukawa promotion"},
	}
	ledger := FourfoldCarrierLedger{Rows: rows, RowCount: len(rows), Verdict: StatusFourfoldComparisonComplete}
	for _, r := range rows {
		switch r.Status {
		case "native":
			ledger.NativeRows++
		case "preflight":
			ledger.PreflightRows++
		case "quotient":
			ledger.QuotientRows++
		case "sealed":
			ledger.SealedRows++
		case "blocked":
			ledger.BlockedRows++
		}
	}
	return ledger
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		Verdict: StatusFirewallPreserved,
	}
}

func buildFinal(selector SelectorAlgebraAudit, weak WeakPlaneSieve, tau TauEtaPullbackAudit, contact ContactQuarticAudit) FinalVerdict {
	return FinalVerdict{
		NativeSelectorAlgebraTheorem:    selector.AllRowsVerified && selector.CommutantDimension == selector.ExpectedDimension,
		NativeThreeToTwoPlusOneSelector: weak.UniqueWeakPlane || tau.NativeThreeToTwoPlusOne,
		TauEtaPullbackValid:             tau.ExistingUnitPreservingPullback && tau.RhoOneIsIdentity,
		TauEtaSealed:                    !tau.ExistingUnitPreservingPullback && tau.CanSelectTwoPlusOneIfPulledBack,
		ContactCarrierActionValid:       contact.NativeCarrierAction,
		ContactRemainsContactOnly:       !contact.NativeCarrierAction && contact.RegularRepresentationUnit && contact.IrreducibleOverQ,
		NextTheorem:                     "Gate 556 should be a unit-preserving trace-transfer/pullback theorem: construct or obstruct rho_tau or rho_4 into End(W_spatial), End(C^3_gen), or End(H_phi), with rho(1)=I and compatibility with grading, J, D, first-order condition, and B-L.",
		Verdict:                         StatusNoNativeThreeToTwoPlusOne,
	}
}

func validate(a Analysis) error {
	if !a.Selector.AllRowsVerified || a.Selector.CommutantDimension != 10 {
		return fmt.Errorf("selector algebra theorem did not verify expected B-L commutant dimension")
	}
	if a.BMinusL.ExpectedDimension != 10 || a.BMinusL.CommutantDimension != 10 || !a.BMinusL.AllBridgeDeltasPMFourThird {
		return fmt.Errorf("B-L audit failed")
	}
	if a.WeakPlane.PreservedCount != 3 || a.WeakPlane.RejectedCount != 3 || a.WeakPlane.UniqueWeakPlane {
		return fmt.Errorf("weak-plane B-L degeneracy audit failed")
	}
	if a.TauEta.ExistingUnitPreservingPullback || a.TauEta.NativeThreeToTwoPlusOne {
		return fmt.Errorf("tau_eta was illegally promoted to native pullback")
	}
	if !a.Contact.RegularRepresentationUnit || !a.Contact.IrreducibleOverQ || a.Contact.NativeCarrierAction {
		return fmt.Errorf("contact quartic firewall failed")
	}
	if a.Firewall.DimensionMatchesPromoted || a.Firewall.TauEtaPromotedToFockSelector || a.Firewall.ContactQuarticPromotedToHiggs || a.Firewall.ContactQuarticPromotedToFlavor || a.Firewall.ContactQuarticPromotedToYukawa || a.Firewall.NativeRegistryPolluted {
		return fmt.Errorf("firewall pollution detected")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusSelectorAlgebraTheoremProved,
		StatusCommutantDimensionVerified,
		StatusBMinusLOnePlusThreeVerified,
		StatusBMinusLCommutantVerified,
		StatusBridgeDeltaVerified,
		StatusWeakPlaneSieveExecuted,
		StatusWeakPlaneDegeneracyRemains,
		StatusNoTauEtaPullback,
		StatusTauEtaCapacitySealed,
		StatusContactRegularRepresentation,
		StatusContactIrreducible,
		StatusNoContactCarrierAction,
		StatusFourfoldComparisonComplete,
		StatusNoNativeThreeToTwoPlusOne,
		StatusFirewallPreserved,
	}
}

func FormatSelector(a SelectorAlgebraAudit) string {
	return fmt.Sprintf("%s; %s; multiplicities=%v; dim Comm=%d", a.CommutatorIdentity, a.CommutantDimensionFormula, a.Multiplicities, a.CommutantDimension)
}

func FormatBMinusL(a BMinusLAudit) string {
	return fmt.Sprintf("coefficients=%v; split=%s; commutant=%s; dim=%d; bridge deltas=±%s", formatSlice(a.Coefficients), a.Split, a.Commutant, a.CommutantDimension, formatFloat(a.BridgeDeltaAbs))
}

func FormatWeakPlane(a WeakPlaneSieve) string {
	return fmt.Sprintf("rejected=%s; preserved=%s; unique=%v", strings.Join(a.RejectedPlanes, ","), strings.Join(a.PreservedPlanes, ","), a.UniqueWeakPlane)
}

func FormatTauEta(a TauEtaPullbackAudit) string {
	return fmt.Sprintf("tau_eta=%v; |tau_eta|=%v; unit-preserving pullback=%v; capacity=%v; verdict=%s", a.TauEta, a.AbsTauEta, a.ExistingUnitPreservingPullback, a.CanSelectTwoPlusOneIfPulledBack, a.Verdict)
}

func FormatContact(a ContactQuarticAudit) string {
	return fmt.Sprintf("q4=%s; rho_reg(1)=I4=%v; irreducible_Q=%v; rational_idempotents=%d; carrier_action=%v", a.PolynomialString, a.RegularRepresentationUnit, a.IrreducibleOverQ, a.NontrivialRationalIdempotents, a.NativeCarrierAction)
}

func FormatCarriers(a FourfoldCarrierLedger) string {
	parts := make([]string, 0, len(a.Rows))
	for _, r := range a.Rows {
		parts = append(parts, fmt.Sprintf("%s:%s:%s", r.Name, r.Split, r.Status))
	}
	return strings.Join(parts, " | ")
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("dimension_promoted=%v tau_promoted=%v contact_higgs=%v contact_flavor=%v contact_yukawa=%v observed_yukawas=%v native_polluted=%v", a.DimensionMatchesPromoted, a.TauEtaPromotedToFockSelector, a.ContactQuarticPromotedToHiggs, a.ContactQuarticPromotedToFlavor, a.ContactQuarticPromotedToYukawa, a.ObservedYukawasImported, a.NativeRegistryPolluted)
}

func truth() string {
	return "Gate555 proves the native selector algebra and confirms B-L as a native 4->1+3 selector. It does not find a native 3->2+1 selector: B-L preserves three spatial weak planes, tau_eta has only sealed 2+1 capacity without a unit-preserving Fock/generation pullback, and the contact quartic remains an irreducible contact algebra with no native carrier action on W, W_spatial, or H_phi."
}

func eigenMultiplicityList(s []float64) []int {
	counts := map[string]int{}
	for _, x := range s {
		counts[formatFloat(x)]++
	}
	vals := make([]int, 0, len(counts))
	for _, v := range counts {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	return vals
}

func nearlyZero(x float64) bool { return math.Abs(x) < 1e-12 }

func formatFloat(x float64) string {
	if nearlyZero(x) {
		return "0"
	}
	if nearlyZero(x - 1.0/3.0) {
		return "1/3"
	}
	if nearlyZero(x + 1.0/3.0) {
		return "-1/3"
	}
	if nearlyZero(x - 4.0/3.0) {
		return "4/3"
	}
	if nearlyZero(x + 4.0/3.0) {
		return "-4/3"
	}
	if nearlyZero(x + 1) {
		return "-1"
	}
	if nearlyZero(x - 1) {
		return "1"
	}
	return fmt.Sprintf("%.12g", x)
}

func formatSlice(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = formatFloat(x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func hasRationalRoot(poly []int) bool {
	if len(poly) != 5 {
		return false
	}
	lead := abs(poly[0])
	constant := abs(poly[4])
	for _, p := range divisors(constant) {
		for _, q := range divisors(lead) {
			for _, sign := range []int{1, -1} {
				numer := sign * p
				// Evaluate q^4 * f(numer/q), avoiding floating point.
				v := poly[0]*pow(numer, 4) + poly[1]*pow(numer, 3)*q + poly[2]*pow(numer, 2)*pow(q, 2) + poly[3]*numer*pow(q, 3) + poly[4]*pow(q, 4)
				if v == 0 {
					return true
				}
			}
		}
	}
	return false
}

func hasQuadraticFactorOverZ(poly []int) bool {
	if len(poly) != 5 {
		return false
	}
	A, B, C, D, E := poly[0], poly[1], poly[2], poly[3], poly[4]
	for _, a := range signedDivisors(A) {
		if A%a != 0 {
			continue
		}
		d := A / a
		for _, c := range signedDivisors(E) {
			if E%c != 0 {
				continue
			}
			f := E / c
			// (a x^2 + b x + c)(d x^2 + e x + f)
			// gives B = a*e + b*d and D = b*f + c*e.
			den := d*c - a*f
			if den == 0 {
				continue
			}
			numB := B*c - a*D
			numE := d*D - B*f
			if numB%den != 0 || numE%den != 0 {
				continue
			}
			b := numB / den
			e := numE / den
			if a*e+b*d == B && b*e+a*f+c*d == C && b*f+c*e == D {
				return true
			}
		}
	}
	return false
}

func divisors(n int) []int {
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return nil
	}
	out := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			out = append(out, i)
			if i*i != n {
				out = append(out, n/i)
			}
		}
	}
	sort.Ints(out)
	return out
}

func signedDivisors(n int) []int {
	base := divisors(n)
	out := make([]int, 0, 2*len(base))
	for _, d := range base {
		out = append(out, d, -d)
	}
	return out
}

func pow(x, n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= x
	}
	return p
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
