#!/usr/bin/env python
# /// script
# requires-python = ">=3.10"
# dependencies = [
#     "matplotlib",
#     "pandas",
# ]
# ///

import argparse
import json
import re
from pathlib import Path

import pandas as pd
import matplotlib.pyplot as plt

parser = argparse.ArgumentParser(description="Plot fetch benchmark results.")
parser.add_argument(
    "json_file",
    nargs="?",
    help="Path to fetch-benchmarks.json (optional, latest date dir will be used if omitted)"
)
args = parser.parse_args()

if args.json_file:
    json_file = Path(args.json_file)
    latest_dir = json_file.parent
else:
    benchmarks_dir = Path(__file__).parent
    date_pattern = re.compile(r"\d{4}-\d{2}-\d{2}")
    latest_dir = max(
        (d for d in benchmarks_dir.iterdir() if d.is_dir() and date_pattern.fullmatch(d.name)),
        key=lambda d: d.name
    )
    json_file = latest_dir / "fetch-benchmarks.json"

with open(json_file) as f:
    data = json.load(f)

results = data["results"]
names = [r["command"] for r in results]
times_ms = [r["mean"] * 1000 for r in results]

df = pd.DataFrame({
    "Fetch Utility": names,
    "Mean Time (ms)": times_ms
})
df.sort_values("Mean Time (ms)", inplace=True)

plt.figure(figsize=(10, 6))
plt.barh(df["Fetch Utility"], df["Mean Time (ms)"])
plt.xlabel("Mean Time (milliseconds)")
plt.title(f"Fetch Benchmark – {latest_dir.name}")
plt.tight_layout()

output_path = latest_dir / "fetch-benchmarks.png"
plt.savefig(output_path, dpi=150)
print(f"✅ Plot saved to {output_path}")

