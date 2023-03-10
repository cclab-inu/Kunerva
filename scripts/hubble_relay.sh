#!/bin/bash

kubectl -n kube-system port-forward service/hubble-relay 4245:80
