#!/bin/bash
echo "Running go bench to generate current benchmarks"
go test -v -bench=. -benchmem -count=6 -benchtime=500ms ./... > current_bench.out
benchstat old_bench.out current_bench.out 
mv current_bench.out old_bench.out
