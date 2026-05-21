package generation2operatorlevelfinitetriplematrixrealizationaudit

import (
	"strings"
	"testing"
)

func TestGate854BasisAndCarrierDimensions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Basis.Complete || !a.Basis.JCopyIncluded || a.Basis.HLRank != 8 || a.Basis.HRMinRank != 7 || a.Basis.HPartMinRank != 15 || a.Basis.HFMinRank != 30 {
		t.Fatalf("bad minimal basis: %s", FormatBasis(a.Basis))
	}
	if a.Basis.AmbientPartRank != 16 || a.Basis.AmbientF != 32 || !a.Basis.RightPunctureOutside || !a.Basis.LeftKernelInHL {
		t.Fatalf("ambient/active fork not tracked: %s", FormatBasis(a.Basis))
	}
}

func TestGate854RhoGammaJSeals(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Rho.PreservesMinimalCarrier || a.Rho.PunctureForcedBackIntoCarrier || !a.Rho.HMayMixHPlusHMinus || a.Rho.HPlusHMinusNativeHEigensplit || a.Rho.NativeRepresentationProof {
		t.Fatalf("rho seal overpromoted or inconsistent: %s", FormatRho(a.Rho))
	}
	if !a.Gamma.SquareIdentity || !a.Gamma.ChiralityOddWithDFByBlock || a.Gamma.KOExtensionCertified || !a.Gamma.SupportLevelOnly {
		t.Fatalf("gamma seal invalid: %s", FormatGamma(a.Gamma))
	}
	if !a.J.ParticleOppositeExchange || !a.J.AntiunitaryFormal || a.J.KOSignsCertified || a.J.OppositeActionCompatibilityProved {
		t.Fatalf("J seal overpromoted: %s", FormatJ(a.J))
	}
}

func TestGate854SymbolicDMatrixSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.D
	if !d.YPlus1Zero || len(d.YTerms) != 3 || d.Rank != 14 || d.KernelRank != 1 || !d.SelfAdjointByBlock || !d.ChiralityOddByBlock || !d.ExtendedToJCopy {
		t.Fatalf("symbolic D matrix invalid: %s", FormatD(d))
	}
	if d.OperatorValuedMatrixCertified || d.NumericalYukawaMagnitudesCertified || d.UnbrokenHEquivariantTheorem {
		t.Fatalf("D matrix overpromoted: %s", FormatD(d))
	}
}

func TestGate854FirstOrderAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Checks.FirstOrderExecutableNextGate || a.Checks.FirstOrderProvedThisGate || a.Impact.FirstOrderProved || a.Impact.NativeFiniteTripleProof {
		t.Fatalf("first-order overpromoted: %s | %s", FormatChecks(a.Checks), FormatImpact(a.Impact))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate854Theorem(t *testing.T) {
	res := Generation2OperatorLevelFiniteTripleMatrixRealizationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
