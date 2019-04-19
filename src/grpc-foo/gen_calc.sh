#!/bin/bash
protoc calculator/calcpb/calculator.proto --go_out=plugins=grpc:.