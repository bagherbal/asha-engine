#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import sys
from pathlib import Path

TOOLKIT_ROOT = Path(__file__).resolve().parents[1]
if str(TOOLKIT_ROOT) not in sys.path:
    sys.path.insert(0, str(TOOLKIT_ROOT))

from asha_figures.contact_vacuum import build_section2_geometry, export_section2_json_dict, validate_section2_geometry
from asha_figures.style import STYLE
from asha_figures.section2_svg import render_section2_svg


def main() -> None:
    parser = argparse.ArgumentParser(description="Build ASHA Section 2 Boolean/G2 Contact Vacuum figure artifacts.")
    parser.add_argument("--out", default="outputs/section2", help="Output directory, relative to toolkit root unless absolute.")
    parser.add_argument("--png", action="store_true", help="Also render PNG via cairosvg when available.")
    args = parser.parse_args()

    toolkit_root = TOOLKIT_ROOT
    out_dir = Path(args.out)
    if not out_dir.is_absolute():
        out_dir = toolkit_root / out_dir
    out_dir.mkdir(parents=True, exist_ok=True)

    geometry = build_section2_geometry(width=STYLE.width, height=STYLE.height)
    checks = validate_section2_geometry(geometry)
    if checks["status"] != "PASS_LOG_RANK_TOPOLOGICAL_CONTACT_VACUUM":
        raise SystemExit(f"validation failed: {checks}")

    svg_path = out_dir / "asha_section2_boolean_g2_contact_vacuum_k7.svg"
    json_path = out_dir / "asha_section2_boolean_g2_contact_vacuum_k7.geometry.json"
    manifest_path = out_dir / "asha_section2_boolean_g2_contact_vacuum_k7.manifest.json"

    render_section2_svg(geometry, svg_path)
    json_path.write_text(json.dumps(export_section2_json_dict(geometry), indent=2, ensure_ascii=False), encoding="utf-8")
    manifest_path.write_text(json.dumps({
        "artifact": svg_path.name,
        "companion_geometry": json_path.name,
        "status": checks["status"],
        "checks": checks,
        "aesthetic": "Transcendental Geometric Minimalism",
        "notes": [
            "P_B is represented as a broad frosted support lens with rank 56.",
            "P_G is represented as a sharper G2/octonionic support lens with rank 14.",
            "K7 is the luminous intersection payload with rank 7.",
            "Areas scale as ln(rank), satisfying the requested topological/logarithmic accuracy level.",
            "The figure does not claim literal high-dimensional embedding geometry."
        ],
    }, indent=2, ensure_ascii=False), encoding="utf-8")

    if args.png:
        try:
            import cairosvg
            png_path = out_dir / "asha_section2_boolean_g2_contact_vacuum_k7.png"
            cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=STYLE.width, output_height=STYLE.height)
            print(png_path)
        except Exception as exc:
            print(f"PNG export skipped: {exc}")

    print(svg_path)
    print(json_path)
    print(manifest_path)


if __name__ == "__main__":
    main()
