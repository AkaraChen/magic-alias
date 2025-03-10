#!/usr/bin/env bash

# Set up test environment
setup() {
  # Create a temporary directory for test files
  export TEMP_DIR="$(mktemp -d)"
  export HOME="$TEMP_DIR"
  export MAGIC_ALIAS_DIR="$HOME/.magic-alias"
  export TEST_RC_FILE="$HOME/.bashrc"
  
  # Create necessary directories
  mkdir -p "$MAGIC_ALIAS_DIR"
  touch "$TEST_RC_FILE"
  
  # Mock SHELL environment variable to use bash
  export SHELL="/bin/bash"
}

# Clean up after tests
teardown() {
  # Remove temporary directory
  rm -rf "$TEMP_DIR"
}

# Helper to check if a string is in a file
assert_file_contains() {
  local file="$1"
  local expected="$2"
  
  if ! grep -q "$expected" "$file"; then
    echo "Expected file $file to contain: $expected"
    echo "File contents:"
    cat "$file"
    return 1
  fi
}

# Helper to check if a string is not in a file
assert_file_not_contains() {
  local file="$1"
  local unexpected="$2"
  
  if grep -q "$unexpected" "$file"; then
    echo "Expected file $file to NOT contain: $unexpected"
    echo "File contents:"
    cat "$file"
    return 1
  fi
}

# Helper to check if a file exists
assert_file_exists() {
  local file="$1"
  
  if [ ! -f "$file" ]; then
    echo "Expected file $file to exist"
    return 1
  fi
}

# Helper to check if a file does not exist
assert_file_not_exists() {
  local file="$1"
  
  if [ -f "$file" ]; then
    echo "Expected file $file to NOT exist"
    return 1
  fi
}
