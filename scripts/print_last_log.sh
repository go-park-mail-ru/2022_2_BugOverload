#!/usr/bin/env bash

RESULT=$({ ls -t logs/prod/* 2> /dev/null || ls -t logs/debug/* 2> /dev/null; } | grep -w log | head -1)
cat "$RESULT"
