#!/usr/bin/env python3
"""Convert JSON input to YAML output.

Usage:
    cat data.json | python3 json2yaml.py > data.yaml
"""
import sys
import json
import yaml

def main():
    try:
        data = json.load(sys.stdin)
    except json.JSONDecodeError as e:
        sys.stderr.write(f"Invalid JSON: {e}\n")
        sys.exit(1)
    yaml.safe_dump(data, sys.stdout, default_flow_style=False, sort_keys=False)

if __name__ == "__main__":
    main()
