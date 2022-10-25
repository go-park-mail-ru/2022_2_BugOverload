#!/usr/bin/env bash

RESULT=$(ls -t | grep -w log | head -1)
cat "$RESULT"