#!/bin/bash

# Automatic technical test output generator
# Generates technical_test.txt in project root

echo "Running Golang Technical Test..."

OUTPUT_FILE="technical_test.txt"

# run go test runner & redirect to file
go run . > $OUTPUT_FILE

if [ $? -eq 0 ]; then
    echo "✔ technical_test.txt generated successfully."
    echo "Location: $(pwd)/$OUTPUT_FILE"
else
    echo "✖ Failed to generate technical_test.txt"
    exit 1
fi