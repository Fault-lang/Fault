#!/bin/bash

echo "Testing Fault TUI fixes..."
echo ""

# Test 1: Traditional CLI mode (should work)
echo "✓ Test 1: Traditional CLI mode (AST)"
./fault -f generator/testdata/booleans.fspec -m ast > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "  ✓ CLI mode works"
else
    echo "  ✗ CLI mode failed"
    exit 1
fi

# Test 2: Build check
echo ""
echo "✓ Test 2: Binary built successfully"
ls -lh ./fault | grep -q "fault$"
if [ $? -eq 0 ]; then
    echo "  ✓ Binary exists"
else
    echo "  ✗ Binary not found"
    exit 1
fi

echo ""
echo "========================================="
echo "Automated tests passed!"
echo ""
echo "To test TUI mode interactively:"
echo "  1. Run: ./fault"
echo "  2. Enter a file path (e.g., generator/testdata/booleans.fspec)"
echo "  3. Select mode (try 'check' or 'ast')"
echo "  4. Select input format (fspec)"
echo "  5. Select output format (log)"
echo "  6. Should show results (no longer hang!)"
echo "  7. Press 'q' to quit"
echo ""
echo "If it still hangs, the debug info will show dimensions"
echo "========================================="
