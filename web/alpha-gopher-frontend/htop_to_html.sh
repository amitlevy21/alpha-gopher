#!/bin/sh

dst=$1

echo q | htop -C | aha --line-fix > $dst
