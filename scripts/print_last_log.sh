#!/usr/bin/env bash

RESULT=$({ ls -t logs/prod/*; ls -t logs/debug/*; } | grep -w log | head -1)
cat "$RESULT"
