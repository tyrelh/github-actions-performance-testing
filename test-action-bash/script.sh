#! /usr/bin/env bash

COUNT=100000

for i in $(seq 1 $COUNT); do
  # Generate a UUID
    uuid=$(uuidgen)
    echo "Generated UUID: $uuid"
done

for i in $(seq 1 $COUNT); do
  # Parse a JSON string
    json='{"name": "GitHub Copilot", "age": 1}'
    name=$(echo $json | jq -r '.name')
    echo "Parsed name from JSON: $name"
done

echo "Done"