from __future__ import annotations

import json
from pathlib import Path
from typing import Dict, List

KEY_FILES = [
    "README.md",
    "docs/summaries/formula_ledger.md",
    "docs/summaries/essential_ontological_tower_map.md",
    "docs/audits/gates/gate1349_final_source_typed_toe_ledger.md",
    "docs/audits/gates/gate387_registry_audit.md",
    "docs/audits/gates/gate973_registry_audit.md",
]

KEY_TERMS = [
    "V_8", "X_4", "P_4", "eta", "Omega", "contact", "K_7", "A_F",
    "N_Q", "W_Q", "S_{ASHA}", "S_grav", "Theorem != Bridge",
    "3,", "4,", "7,", "27", "56", "70", "72", "PMNS", "Majorana",
]


def compact_lines(text: str, terms: List[str], limit: int = 42) -> List[str]:
    lines = text.splitlines()
    out = []
    for i, line in enumerate(lines, start=1):
        normalized = line.strip()
        if not normalized:
            continue
        if any(term in normalized for term in terms):
            out.append(f"L{i}: {normalized}")
        if len(out) >= limit:
            break
    return out


def build_source_digest(repo_root: str | Path) -> Dict:
    root = Path(repo_root)
    files = []
    gates_dir = root / "docs" / "audits" / "gates"
    gate_count = len(list(gates_dir.glob("*.md"))) if gates_dir.exists() else 0
    for rel in KEY_FILES:
        path = root / rel
        if path.exists():
            text = path.read_text(encoding="utf-8", errors="replace")
            files.append({
                "path": rel,
                "bytes": path.stat().st_size,
                "matched_lines": compact_lines(text, KEY_TERMS),
            })
    return {
        "repo_root": str(root),
        "gate_markdown_count": gate_count,
        "key_files": files,
        "note": "Digest is used to ground the visualization contract; generator remains deterministic and source-controlled.",
    }


def write_source_digest(repo_root: str | Path, out_path: str | Path) -> Dict:
    digest = build_source_digest(repo_root)
    Path(out_path).write_text(json.dumps(digest, indent=2), encoding="utf-8")
    return digest
