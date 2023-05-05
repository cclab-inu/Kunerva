#!/bin/bash

kubectl annotate pod nginx-86c57db685-fxv2v io.cilium.proxy-visibility="<Ingress/80/TCP/HTTP>,<Egress/80/TCP/HTTP>"
