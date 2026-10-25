# kube-health — Kubernetes Cluster Debugger CLI

A lightweight Go CLI to diagnose Kubernetes cluster health - nodes, pods, events, namespaces. Built for platform engineers.

## Why this exists
`kubectl get pods` is slow when clusters break. This tool aggregates health checks in one binary < 10MB.

## Features
- ✅ Node health (Ready, MemoryPressure, DiskPressure)
- ✅ Pod crash loop detection & restart counts
- ✅ Namespace listing with pod density
- ✅ kubeconfig auto-discovery (~/.kube/config)
- ✅ Color-coded terminal output

## Tech Stack
- Go 1.27, client-go v0.37.1, apimachinery v0.37.1

## Quick Start
```bash
go mod tidy
go run main.go

# Build
go build -o kube-health
./kube-health
