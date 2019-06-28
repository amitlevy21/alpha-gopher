#!/bin/sh

dst=$1

echo q | htop| aha --line-fix > $dst
