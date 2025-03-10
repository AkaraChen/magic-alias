#!/usr/bin/env bats

load test_helper

@test "ma command exists and runs" {
  run ma
  [ "$status" -eq 0 ]
  [[ "$output" == *"Magic Alias is a modern, user-friendly tool"* ]]
}

@test "ma --help displays help information" {
  run ma --help
  [ "$status" -eq 0 ]
  [[ "$output" == *"Usage:"* ]]
  [[ "$output" == *"Available Commands:"* ]]
  [[ "$output" == *"Flags:"* ]]
}

@test "ma with invalid command shows error" {
  run ma nonexistentcommand
  [ "$status" -ne 0 ]
  [[ "$output" == *"unknown command"* ]]
}
