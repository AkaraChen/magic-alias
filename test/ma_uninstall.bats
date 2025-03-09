#!/usr/bin/env bats

load test_helper

@test "ma uninstall removes configuration from shell rc file" {
  # First initialize
  run ma init
  [ "$status" -eq 0 ]
  
  # Verify the configuration was added
  assert_file_contains "$TEST_RC_FILE" "# Magic Alias (ma)"
  
  # Run uninstall with --yes flag to skip confirmation
  run ma uninstall --yes
  [ "$status" -eq 0 ]
  [[ "$output" == *"Uninstall Magic Alias successfully"* ]]
  
  # Verify the configuration was removed
  assert_file_not_contains "$TEST_RC_FILE" "# Magic Alias (ma)"
  assert_file_not_contains "$TEST_RC_FILE" "export PATH=\"\$PATH:\$HOME/.magic-alias\""
}

@test "ma uninstall with --remove-aliases removes aliases directory" {
  # First initialize
  run ma init
  [ "$status" -eq 0 ]
  
  # Create some test aliases
  ma add test1 ls
  ma add test2 pwd
  
  # Verify the aliases directory exists
  [ -d "$MAGIC_ALIAS_DIR" ]
  
  # Run uninstall with --yes and --remove-aliases flags
  run ma uninstall --yes --remove-aliases
  [ "$status" -eq 0 ]
  [[ "$output" == *"All aliases have been removed"* ]]
  
  # Verify the aliases directory is gone or empty
  if [ -d "$MAGIC_ALIAS_DIR" ]; then
    # If the directory still exists, it should be empty
    run ls -A "$MAGIC_ALIAS_DIR"
    [ "$output" = "" ]
  fi
}

@test "ma uninstall creates backup of rc file" {
  # First initialize
  run ma init
  [ "$status" -eq 0 ]
  
  # Run uninstall with --yes flag
  run ma uninstall --yes
  [ "$status" -eq 0 ]
  [[ "$output" == *"Backup created"* ]]
  
  # Extract the backup file path from the output
  backup_file=$(echo "$output" | grep "Backup created" | sed -E 's/.*Backup created.*backup: ([^ ]+).*/\1/')
  
  # Verify the backup file exists
  [ -f "$backup_file" ]
}

@test "ma uninstall fails when not initialized" {
  # Don't initialize first
  
  # Run uninstall
  run ma uninstall --yes
  [ "$status" -ne 0 ]
  [[ "$output" == *"magic alias line not found in rc file"* ]]
}
