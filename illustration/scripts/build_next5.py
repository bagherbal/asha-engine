#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

TOOLKIT_ROOT = Path(__file__).resolve().parents[1]
if str(TOOLKIT_ROOT) not in sys.path:
    sys.path.insert(0, str(TOOLKIT_ROOT))

from asha_figures.readme_next5 import next5_contracts, export_contract, validate_contract
from asha_figures.section4_8_svg import render_contract_svg
from asha_figures.style import STYLE


def main() -> None:
    parser = argparse.ArgumentParser(description="Build the next five ASHA README visual atlas figures.")
    parser.add_argument("--out", default="outputs/next5", help="Output directory, relative to toolkit root unless absolute.")
    parser.add_argument("--png", action="store_true", help="Also render PNG via cairosvg when available.")
    args = parser.parse_args()
    out_dir = Path(args.out)
    if not out_dir.is_absolute():
        out_dir = TOOLKIT_ROOT / out_dir
    out_dir.mkdir(parents=True, exist_ok=True)

    all_manifests = []
    for c in next5_contracts():
        checks = validate_contract(c)
        if checks["status"] == "FAIL":
            raise SystemExit(f"validation failed for {c.figure_id}: {checks}")
        svg_path = out_dir / f"{c.figure_id}.svg"
        json_path = out_dir / f"{c.figure_id}.geometry.json"
        manifest_path = out_dir / f"{c.figure_id}.manifest.json"
        render_contract_svg(c, svg_path)
        json_path.write_text(json.dumps(export_contract(c), indent=2, ensure_ascii=False), encoding="utf-8")
        manifest = {
            "artifact": svg_path.name,
            "companion_geometry": json_path.name,
            "status": checks["status"],
            "checks": checks,
            "aesthetic": "Transcendental Geometric Minimalism",
            "readme_anchor": c.readme_anchor,
            "notes": c.notes,
        }
        manifest_path.write_text(json.dumps(manifest, indent=2, ensure_ascii=False), encoding="utf-8")
        if args.png:
            try:
                import cairosvg
                png_path = out_dir / f"{c.figure_id}.png"
                cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=STYLE.width, output_height=STYLE.height)
                print(png_path)
            except Exception as exc:
                print(f"PNG export skipped for {c.figure_id}: {exc}")
        print(svg_path)
        print(json_path)
        print(manifest_path)
        all_manifests.append(manifest)
    (out_dir / "asha_readme_next5.manifest.json").write_text(json.dumps(all_manifests, indent=2, ensure_ascii=False), encoding="utf-8")

if __name__ == "__main__":
    main()
