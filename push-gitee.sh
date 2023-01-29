#!/usr/bin/env bash

set -e

git remote set-url origin https://gitee.com/quant1x/gotdx.git
git push --all
git push --tags
git remote -vv