// Package generation2topologicalanomalyledger implements Gate 490:
// Topological Charge & Anomaly Cancellation Ledger.
//
// Gate 489 closed the native Yukawa/CKM selector branch. Gate 490 redirects the
// engine to a non-flavor invariant: anomaly cancellation. Unlike masses and
// mixing angles, chiral gauge-anomaly ledgers depend only on the discrete
// representation table. This package therefore audits the one-generation
// left-handed Weyl ledger with exact rational arithmetic and verifies that the
// Standard Model gauge anomalies cancel without importing any Yukawa, CKM, PMNS,
// mass, or Wolfenstein data.
//
// The result is deliberately bounded: the finite charge ledger cancels exactly
// and is stable under family replication, but this is not a Yukawa selector, not
// a CKM theorem, and not a derivation of flavor moduli. It is a topological
// consistency theorem for the already-admitted discrete representation ledger.
package generation2topologicalanomalyledger

import (
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/anomaly"
)

const (
	AuditID = "GATE490-TOPOLOGICAL-CHARGE-ANOMALY-CANCELLATION-LEDGER"

	StatusGate489Inherited              = "CONDITIONAL_SUPPORT_GATE489_FLAVOR_AIRLOCK_INHERITED"
	StatusChargeLedgerConstructed       = "CONDITIONAL_SUPPORT_NATIVE_DISCRETE_CHARGE_LEDGER_CONSTRUCTED"
	StatusExactGaugeAnomalyCancellation = "CONDITIONAL_SUPPORT_NATIVE_ANOMALY_CANCELLATION_PROVEN"
	StatusExactABJTriangleTracesZero    = "CONDITIONAL_SUPPORT_ABJ_TRIANGLE_TRACES_CANCEL_EXACTLY"
	StatusWittenSU2GlobalAnomalyCleared = "CONDITIONAL_SUPPORT_WITTEN_SU2_GLOBAL_ANOMALY_EVEN_DOUBLETS"
	StatusFamilyReplicationStable       = "CONDITIONAL_SUPPORT_FAMILY_REPLICATION_ANOMALY_STABLE"
	StatusExistingGate79Consistent      = "CONDITIONAL_SUPPORT_GATE79_ANOMALY_LEDGER_CONFIRMED"
	StatusFlavorMassIndependent         = "CONDITIONAL_SUPPORT_TOPOLOGICAL_LEDGER_FLAVOR_MASS_INDEPENDENT"
	StatusNoYukawaSelector              = "FAILED_ROUTE_ANOMALY_CANCELLATION_DOES_NOT_SELECT_YUKAWA_TEXTURE"
	StatusNoCKMJarlskog                 = "FAILED_ROUTE_ANOMALY_CANCELLATION_DOES_NOT_DERIVE_CKM_OR_JARLSKOG"
	StatusNoContinuumDynamics           = "FAILED_ROUTE_ANOMALY_LEDGER_DOES_NOT_DERIVE_CONTINUUM_DYNAMICS"
	StatusFirewallPreserved             = "FIREWALL_PRESERVED_NO_FLAVOR_DATA_IMPORTED"
	StatusGate491RedirectDefined        = "CONDITIONAL_SUPPORT_GATE491_SCALAR_EDGE_STABILITY_REDIRECT_DEFINED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                          bool
	Gate485NullKoideBaselineInherited bool
	Gate489FlavorAirlockClosed        bool
	YukawaEntriesEnvironmental        bool
	CKMOrientationEnvironmental       bool
	JarlskogEnvironmental             bool
	NoFlavorDataImported              bool
	Verdict                           string
	Reason                            string
}

type WeylMultiplet struct {
	Name                    string
	Description             string
	WeylMultiplicity        int
	Hypercharge             string
	BMinusL                 string
	WeakDoubletMultiplicity int
	ColorDynkinMultiplicity int
	ColorCubicMultiplicity  int
	ColorCubicSign          int
}

type ChargeLedger struct {
	Executed             bool
	Multiplets           []WeylMultiplet
	LeftHandedWeylStates int
	WeakDoubletCount     int
	WeakDoubletCountEven bool
	ContainsNuRConjugate bool
	DiscreteOnly         bool
	ObservedMassInput    bool
	ObservedMixingInput  bool
	Verdict              string
	Reason               string
}

type Moment struct {
	Symbol      string
	Value       string
	Cancels     bool
	Category    string
	Description string
}

type AnomalyAudit struct {
	Executed                   bool
	Moments                    []Moment
	AllPerturbativeGaugeCancel bool
	AllMixedGaugeGravityCancel bool
	SU2GlobalWittenCancels     bool
	ABJTriangleTraceCount      int
	ZeroTraceCount             int
	ExactRationalArithmetic    bool
	ExistingGate79StateCount   int
	ExistingGate79Cancels      bool
	ExistingGate79Consistent   bool
	Verdict                    string
	Reason                     string
}

type StabilityTheorem struct {
	Executed                       bool
	GenerationUniversal            bool
	FamilyReplicationPreservesZero bool
	FlavorMassIndependent          bool
	YukawaIndependent              bool
	CKMIndependent                 bool
	PMNSIndependent                bool
	GaugeStabilityLedgerSatisfied  bool
	YukawaTextureSelected          bool
	CKMJarlskogDerived             bool
	ContinuumDynamicsDerived       bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                    bool
	ObservedMassesImported      bool
	ObservedYukawaImported      bool
	ObservedCKMImported         bool
	ObservedPMNSImported        bool
	ObservedWolfensteinImported bool
	NativeYukawaMatrixWritten   bool
	NativeCKMMatrixWritten      bool
	NativeJarlskogWritten       bool
	NativeFlavorModuliChanged   bool
	NativeFlavorDimAfter        int
	KXYCoeffDimAfter            int
	Verdict                     string
	Reason                      string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Ledger      ChargeLedger
	Anomaly     AnomalyAudit
	Stability   StabilityTheorem
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Ledger = buildChargeLedger()
	a.Anomaly = buildAnomalyAudit(a.Ledger)
	a.Stability = buildStabilityTheorem(a.Ledger, a.Anomaly)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate485NullKoideBaselineInherited: true,
		Gate489FlavorAirlockClosed:        true,
		YukawaEntriesEnvironmental:        true,
		CKMOrientationEnvironmental:       true,
		JarlskogEnvironmental:             true,
		NoFlavorDataImported:              true,
		Verdict:                           StatusGate489Inherited,
		Reason:                            "Gate489 formally closed native Yukawa/CKM prediction. Gate490 therefore redirects to mass-independent topological consistency traces built only from discrete representation charges.",
	}
}

func buildChargeLedger() ChargeLedger {
	multiplets := []WeylMultiplet{
		{Name: "Q_L", Description: "left quark weak doublets: (u_L,d_L) in three colors", WeylMultiplicity: 6, Hypercharge: "1/6", BMinusL: "1/3", WeakDoubletMultiplicity: 3, ColorDynkinMultiplicity: 2, ColorCubicMultiplicity: 2, ColorCubicSign: +1},
		{Name: "u_R^c", Description: "left-handed conjugate of right up quark", WeylMultiplicity: 3, Hypercharge: "-2/3", BMinusL: "-1/3", WeakDoubletMultiplicity: 0, ColorDynkinMultiplicity: 1, ColorCubicMultiplicity: 1, ColorCubicSign: -1},
		{Name: "d_R^c", Description: "left-handed conjugate of right down quark", WeylMultiplicity: 3, Hypercharge: "1/3", BMinusL: "-1/3", WeakDoubletMultiplicity: 0, ColorDynkinMultiplicity: 1, ColorCubicMultiplicity: 1, ColorCubicSign: -1},
		{Name: "L_L", Description: "left lepton weak doublet: (nu_L,e_L)", WeylMultiplicity: 2, Hypercharge: "-1/2", BMinusL: "-1", WeakDoubletMultiplicity: 1, ColorDynkinMultiplicity: 0, ColorCubicMultiplicity: 0, ColorCubicSign: 0},
		{Name: "e_R^c", Description: "left-handed conjugate of right electron", WeylMultiplicity: 1, Hypercharge: "1", BMinusL: "1", WeakDoubletMultiplicity: 0, ColorDynkinMultiplicity: 0, ColorCubicMultiplicity: 0, ColorCubicSign: 0},
		{Name: "nu_R^c", Description: "left-handed conjugate of right neutrino / sterile hypercharge row", WeylMultiplicity: 1, Hypercharge: "0", BMinusL: "1", WeakDoubletMultiplicity: 0, ColorDynkinMultiplicity: 0, ColorCubicMultiplicity: 0, ColorCubicSign: 0},
	}
	weyl, weakDoublets := 0, 0
	for _, m := range multiplets {
		weyl += m.WeylMultiplicity
		weakDoublets += m.WeakDoubletMultiplicity
	}
	return ChargeLedger{
		Executed:             true,
		Multiplets:           multiplets,
		LeftHandedWeylStates: weyl,
		WeakDoubletCount:     weakDoublets,
		WeakDoubletCountEven: weakDoublets%2 == 0,
		ContainsNuRConjugate: true,
		DiscreteOnly:         true,
		ObservedMassInput:    false,
		ObservedMixingInput:  false,
		Verdict:              StatusChargeLedgerConstructed,
		Reason:               "the one-generation left-handed Weyl ledger contains 16 discrete states and four weak doublets; entries are representation charges, not masses or mixing angles",
	}
}

func buildAnomalyAudit(l ChargeLedger) AnomalyAudit {
	moments := []Moment{
		ratMoment("Tr(Y)", gravY(l.Multiplets), "mixed gauge-gravity", "mixed gravitational-U(1)_Y anomaly"),
		ratMoment("Tr(Y^3)", u1Cubed(l.Multiplets, func(m WeylMultiplet) string { return m.Hypercharge }), "abelian triangle", "cubic U(1)_Y ABJ triangle anomaly"),
		ratMoment("SU(2)_L^2·Y", su2sqU1(l.Multiplets, func(m WeylMultiplet) string { return m.Hypercharge }), "nonabelian-abelian triangle", "weak-isospin squared with hypercharge insertion"),
		ratMoment("SU(3)_c^2·Y", su3sqU1(l.Multiplets, func(m WeylMultiplet) string { return m.Hypercharge }), "nonabelian-abelian triangle", "color squared with hypercharge insertion"),
		ratMoment("SU(3)_c^3", su3Cubed(l.Multiplets), "nonabelian triangle", "quark doublet fundamentals cancel right-conjugate antifundamentals"),
		{Symbol: "SU(2)_L^3", Value: "0", Cancels: true, Category: "nonabelian triangle", Description: "perturbative local SU(2) anomaly vanishes because the fundamental representation is pseudoreal"},
		{Symbol: "Witten SU(2)_L", Value: fmt.Sprintf("%d doublets", l.WeakDoubletCount), Cancels: l.WeakDoubletCountEven, Category: "global anomaly", Description: "global SU(2) anomaly is absent because the number of left weak doublets is even"},
		ratMoment("Tr(B-L)", gravBL(l.Multiplets), "B-L cross-check", "mixed gravitational-(B-L) ledger; nu_R row is required"),
		ratMoment("Tr((B-L)^3)", u1Cubed(l.Multiplets, func(m WeylMultiplet) string { return m.BMinusL }), "B-L cross-check", "cubic B-L ledger; nu_R row is required"),
	}
	zero := 0
	allPert := true
	allMixed := true
	for _, m := range moments {
		if m.Cancels {
			zero++
		}
		if m.Symbol != "Witten SU(2)_L" && !m.Cancels {
			allPert = false
		}
		if (m.Symbol == "Tr(Y)" || m.Symbol == "Tr(B-L)") && !m.Cancels {
			allMixed = false
		}
	}

	existing, err := anomaly.BuildDefault()
	existingStates := 0
	existingCancels := false
	if err == nil {
		existingStates = len(existing.States)
		existingCancels = existing.YAnomalyCancels && existing.BMinusLAnomalyCancels && existing.MixedAbelianCancels
	}
	return AnomalyAudit{
		Executed:                   true,
		Moments:                    moments,
		AllPerturbativeGaugeCancel: allPert,
		AllMixedGaugeGravityCancel: allMixed,
		SU2GlobalWittenCancels:     l.WeakDoubletCountEven,
		ABJTriangleTraceCount:      len(moments) - 1,
		ZeroTraceCount:             zero,
		ExactRationalArithmetic:    true,
		ExistingGate79StateCount:   existingStates,
		ExistingGate79Cancels:      existingCancels,
		ExistingGate79Consistent:   err == nil && existingStates == l.LeftHandedWeylStates && existingCancels,
		Verdict:                    StatusExactGaugeAnomalyCancellation,
		Reason:                     "all local gauge and mixed gauge-gravity anomaly moments vanish exactly as rational representation traces; the global SU(2) doublet count is even",
	}
}

func buildStabilityTheorem(l ChargeLedger, aa AnomalyAudit) StabilityTheorem {
	stable := aa.AllPerturbativeGaugeCancel && aa.AllMixedGaugeGravityCancel && aa.SU2GlobalWittenCancels
	return StabilityTheorem{
		Executed:                       true,
		GenerationUniversal:            true,
		FamilyReplicationPreservesZero: stable,
		FlavorMassIndependent:          true,
		YukawaIndependent:              true,
		CKMIndependent:                 true,
		PMNSIndependent:                true,
		GaugeStabilityLedgerSatisfied:  stable,
		YukawaTextureSelected:          false,
		CKMJarlskogDerived:             false,
		ContinuumDynamicsDerived:       false,
		Verdict:                        StatusFamilyReplicationStable,
		Reason:                         fmt.Sprintf("one generation cancels exactly, so repeating the same generation-universal ledger multiplies zero by N; this proves topological gauge consistency but does not select flavor data; weak doublets=%d", l.WeakDoubletCount),
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                    true,
		ObservedMassesImported:      false,
		ObservedYukawaImported:      false,
		ObservedCKMImported:         false,
		ObservedPMNSImported:        false,
		ObservedWolfensteinImported: false,
		NativeYukawaMatrixWritten:   false,
		NativeCKMMatrixWritten:      false,
		NativeJarlskogWritten:       false,
		NativeFlavorModuliChanged:   false,
		NativeFlavorDimAfter:        NativeFlavorDim,
		KXYCoeffDimAfter:            KXYCoeffDim,
		Verdict:                     StatusFirewallPreserved,
		Reason:                      "Gate490 imports no masses, Yukawa entries, CKM/PMNS data, Wolfenstein parameters, or Jarlskog value and does not reopen the closed flavor branch",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"the one-generation discrete chiral representation ledger is anomaly-balanced: Tr(Y), Tr(Y^3), SU(2)^2·Y, SU(3)^2·Y, SU(3)^3, and local SU(2)^3 all vanish",
			"the global SU(2) Witten anomaly is absent because the finite ledger contains four left weak doublets per generation",
			"anomaly cancellation is generation-universal and remains zero under family replication",
		},
		BridgeEntries: []string{
			"the ledger uses the already-admitted standard-orientation charge branch; it is a representation-consistency theorem, not a flavor-selector theorem",
			"B-L cancellation with nu_R is retained as a consistency cross-check, not as a new physical U(1) coupling derivation",
		},
		EnvironmentalEntries: []string{
			"Yukawa entries, quark/lepton masses, CKM, PMNS, Wolfenstein parameters, and Jarlskog remain environmental/bridge data",
		},
		FailedRoutes: []string{
			StatusNoYukawaSelector,
			StatusNoCKMJarlskog,
			StatusNoContinuumDynamics,
		},
		OpenTheorems: []string{
			StatusGate491RedirectDefined,
			"search for native scalar-edge stability or continuum-permission conditions that remain independent of flavor moduli",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        491,
		Title:       "Scalar-Edge Stability and Higgs One-Form Positivity Audit",
		Reason:      "Gate490 proves the topological charge ledger is stable and flavor-independent. The next non-flavor native frontier should test whether the finite Higgs/edge action has a positivity or stability theorem independent of Yukawa textures.",
		PrimaryTask: "audit scalar-edge Hessian positivity, Goldstone directions, and allowed continuum-matching permissions without importing masses or flavor data",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate485NullKoideBaselineInherited || !a.Inheritance.Gate489FlavorAirlockClosed || !a.Inheritance.YukawaEntriesEnvironmental || !a.Inheritance.CKMOrientationEnvironmental || !a.Inheritance.JarlskogEnvironmental || !a.Inheritance.NoFlavorDataImported {
		return fmt.Errorf("Gate490 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Ledger.Executed || len(a.Ledger.Multiplets) != 6 || a.Ledger.LeftHandedWeylStates != 16 || a.Ledger.WeakDoubletCount != 4 || !a.Ledger.WeakDoubletCountEven || !a.Ledger.ContainsNuRConjugate || !a.Ledger.DiscreteOnly || a.Ledger.ObservedMassInput || a.Ledger.ObservedMixingInput {
		return fmt.Errorf("Gate490 charge ledger invalid: %+v", a.Ledger)
	}
	if !a.Anomaly.Executed || !a.Anomaly.AllPerturbativeGaugeCancel || !a.Anomaly.AllMixedGaugeGravityCancel || !a.Anomaly.SU2GlobalWittenCancels || !a.Anomaly.ExactRationalArithmetic || a.Anomaly.ZeroTraceCount != len(a.Anomaly.Moments) || !a.Anomaly.ExistingGate79Consistent {
		return fmt.Errorf("Gate490 anomaly audit invalid: %+v", a.Anomaly)
	}
	if !a.Stability.Executed || !a.Stability.GenerationUniversal || !a.Stability.FamilyReplicationPreservesZero || !a.Stability.FlavorMassIndependent || !a.Stability.YukawaIndependent || !a.Stability.CKMIndependent || !a.Stability.PMNSIndependent || !a.Stability.GaugeStabilityLedgerSatisfied || a.Stability.YukawaTextureSelected || a.Stability.CKMJarlskogDerived || a.Stability.ContinuumDynamicsDerived {
		return fmt.Errorf("Gate490 stability theorem invalid: %+v", a.Stability)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassesImported || a.Firewall.ObservedYukawaImported || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.NativeYukawaMatrixWritten || a.Firewall.NativeCKMMatrixWritten || a.Firewall.NativeJarlskogWritten || a.Firewall.NativeFlavorModuliChanged || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate490 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate490 proves the non-flavor topological charge ledger: the one-generation left-handed Weyl representation table has %d states and %d weak doublets, and every audited local/mixed anomaly trace vanishes exactly by rational arithmetic. This gives ASHA a native, mass-independent gauge-stability ledger. It does not reopen Yukawa/CKM prediction: no mass, texture, CKM matrix, Jarlskog value, or continuum coupling is derived. Zero anomaly is law-space consistency, not flavor history.", a.Ledger.LeftHandedWeylStates, a.Ledger.WeakDoubletCount)
}

func ratMoment(symbol string, v *big.Rat, category, desc string) Moment {
	return Moment{Symbol: symbol, Value: ratString(v), Cancels: v.Sign() == 0, Category: category, Description: desc}
}

func gravY(ms []WeylMultiplet) *big.Rat {
	return u1Linear(ms, func(m WeylMultiplet) string { return m.Hypercharge })
}
func gravBL(ms []WeylMultiplet) *big.Rat {
	return u1Linear(ms, func(m WeylMultiplet) string { return m.BMinusL })
}

func u1Linear(ms []WeylMultiplet, charge func(WeylMultiplet) string) *big.Rat {
	out := rat(0)
	for _, m := range ms {
		term := mul(intRat(m.WeylMultiplicity), parseRat(charge(m)))
		out.Add(out, term)
	}
	return out
}

func u1Cubed(ms []WeylMultiplet, charge func(WeylMultiplet) string) *big.Rat {
	out := rat(0)
	for _, m := range ms {
		q := parseRat(charge(m))
		q3 := mul(q, q, q)
		term := mul(intRat(m.WeylMultiplicity), q3)
		out.Add(out, term)
	}
	return out
}

func su2sqU1(ms []WeylMultiplet, charge func(WeylMultiplet) string) *big.Rat {
	out := rat(0)
	for _, m := range ms {
		if m.WeakDoubletMultiplicity == 0 {
			continue
		}
		term := mul(intRat(m.WeakDoubletMultiplicity), parseRat("1/2"), parseRat(charge(m)))
		out.Add(out, term)
	}
	return out
}

func su3sqU1(ms []WeylMultiplet, charge func(WeylMultiplet) string) *big.Rat {
	out := rat(0)
	for _, m := range ms {
		if m.ColorDynkinMultiplicity == 0 {
			continue
		}
		term := mul(intRat(m.ColorDynkinMultiplicity), parseRat("1/2"), parseRat(charge(m)))
		out.Add(out, term)
	}
	return out
}

func su3Cubed(ms []WeylMultiplet) *big.Rat {
	out := rat(0)
	for _, m := range ms {
		if m.ColorCubicMultiplicity == 0 {
			continue
		}
		term := intRat(m.ColorCubicMultiplicity * m.ColorCubicSign)
		out.Add(out, term)
	}
	return out
}

func parseRat(s string) *big.Rat {
	x, ok := new(big.Rat).SetString(s)
	if !ok {
		panic("invalid rational: " + s)
	}
	return x
}

func rat(n int64) *big.Rat  { return new(big.Rat).SetInt64(n) }
func intRat(n int) *big.Rat { return new(big.Rat).SetInt64(int64(n)) }

func mul(xs ...*big.Rat) *big.Rat {
	out := new(big.Rat).SetInt64(1)
	for _, x := range xs {
		out.Mul(out, x)
	}
	return out
}

func ratString(x *big.Rat) string {
	if x.Sign() == 0 {
		return "0"
	}
	if x.IsInt() {
		return x.Num().String()
	}
	return x.RatString()
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: Gate485_null=%t Gate489_airlock=%t Yukawa_env=%t CKM_env=%t J_env=%t flavor_data=%t; %s", x.Verdict, x.Gate485NullKoideBaselineInherited, x.Gate489FlavorAirlockClosed, x.YukawaEntriesEnvironmental, x.CKMOrientationEnvironmental, x.JarlskogEnvironmental, !x.NoFlavorDataImported, x.Reason)
}

func FormatLedger(x ChargeLedger) string {
	return fmt.Sprintf("%s: multiplets=%d Weyl_states=%d weak_doublets=%d even=%t nu_R^c=%t discrete_only=%t mass_input=%t mixing_input=%t; %s", x.Verdict, len(x.Multiplets), x.LeftHandedWeylStates, x.WeakDoubletCount, x.WeakDoubletCountEven, x.ContainsNuRConjugate, x.DiscreteOnly, x.ObservedMassInput, x.ObservedMixingInput, x.Reason)
}

func FormatAnomaly(x AnomalyAudit) string {
	return fmt.Sprintf("%s: moments=%d zero=%d perturbative=%t mixed_gravity=%t Witten_SU2=%t exact_rational=%t Gate79(states=%d,cancels=%t); %s", x.Verdict, len(x.Moments), x.ZeroTraceCount, x.AllPerturbativeGaugeCancel, x.AllMixedGaugeGravityCancel, x.SU2GlobalWittenCancels, x.ExactRationalArithmetic, x.ExistingGate79StateCount, x.ExistingGate79Cancels, x.Reason)
}

func FormatStability(x StabilityTheorem) string {
	return fmt.Sprintf("%s: generation_universal=%t family_replication_zero=%t mass_independent=%t Yukawa_independent=%t CKM_independent=%t PMNS_independent=%t gauge_stable=%t Yukawa_selector=%t CKM_J=%t continuum=%t; %s", x.Verdict, x.GenerationUniversal, x.FamilyReplicationPreservesZero, x.FlavorMassIndependent, x.YukawaIndependent, x.CKMIndependent, x.PMNSIndependent, x.GaugeStabilityLedgerSatisfied, x.YukawaTextureSelected, x.CKMJarlskogDerived, x.ContinuumDynamicsDerived, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: masses=%t Yukawa=%t CKM=%t PMNS=%t Wolfenstein=%t native_Yukawa=%t native_CKM=%t native_J=%t flavor_changed=%t dim=%d KXY=%d; %s", x.Verdict, x.ObservedMassesImported, x.ObservedYukawaImported, x.ObservedCKMImported, x.ObservedPMNSImported, x.ObservedWolfensteinImported, x.NativeYukawaMatrixWritten, x.NativeCKMMatrixWritten, x.NativeJarlskogWritten, x.NativeFlavorModuliChanged, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 490 Registry Audit — Topological Charge & Anomaly Cancellation Ledger\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusChargeLedgerConstructed,
		StatusExactGaugeAnomalyCancellation,
		StatusExactABJTriangleTracesZero,
		StatusWittenSU2GlobalAnomalyCleared,
		StatusFamilyReplicationStable,
		StatusFlavorMassIndependent,
		StatusNoYukawaSelector,
		StatusNoCKMJarlskog,
		StatusFirewallPreserved,
	} {
		b.WriteString("- `" + v + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate489 closed the native Yukawa/CKM selector branch. Gate490 therefore audits only discrete topological charge traces. It does not attempt to predict masses, mixing angles, CP phases, or flavor textures.\n\n")
	b.WriteString("## Native charge sieve\n\n")
	b.WriteString(FormatLedger(a.Ledger) + "\n\n")
	b.WriteString("| Multiplet | Description | Weyl multiplicity | Y | B-L | Weak doublets | Color Dynkin copies | Color cubic copies/sign |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|\n")
	for _, m := range a.Ledger.Multiplets {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %s | %s | %d | %d | %d×%+d |\n", m.Name, m.Description, m.WeylMultiplicity, m.Hypercharge, m.BMinusL, m.WeakDoubletMultiplicity, m.ColorDynkinMultiplicity, m.ColorCubicMultiplicity, m.ColorCubicSign))
	}
	b.WriteString("\n## Anomaly cancellation audit\n\n")
	b.WriteString(FormatAnomaly(a.Anomaly) + "\n\n")
	b.WriteString("| Trace | Exact value | Cancels? | Category | Meaning |\n")
	b.WriteString("|---|---:|---:|---|---|\n")
	for _, m := range a.Anomaly.Moments {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %s | %s |\n", m.Symbol, m.Value, m.Cancels, m.Category, m.Description))
	}
	b.WriteString("\nThe local ABJ triangle traces vanish exactly. The SU(2) local cubic anomaly is structurally zero because the doublet is pseudoreal. The global SU(2) anomaly is also absent because the one-generation ledger contains four weak doublets, an even number.\n\n")
	b.WriteString("## Stability theorem\n\n")
	b.WriteString(FormatStability(a.Stability) + "\n\n")
	b.WriteString("Since the cancellation occurs per generation, generation replication multiplies each zero trace by the number of families. This is a topological stability result and remains independent of Yukawa values, masses, CKM, PMNS, and Jarlskog data.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No flavor data entered the theorem. No native Yukawa matrix, CKM matrix, PMNS matrix, Jarlskog invariant, or flavor-moduli update was written.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
