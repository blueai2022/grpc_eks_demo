# grpc_eks_demo
Demo Golang API with gRPC, gRPC Gateway and EKS

## Overview:
User Auth, Proxy and Usage Tracking Service for backend Health API to recognize medical entities from text and output ICD10 code

## Run

### Run app locally with postgres
``docker compose up -d``

### Run app only
``make server``

### Running on AWS EKS
Auto deployment to EKS (through github action) has been turned on for any code push to main
