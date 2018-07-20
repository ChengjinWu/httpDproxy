#!/bin/bash

#rewrite sohudadx as sohudsp
sed -i 's/package sohuadx/package sohudsp/1' sohudsp.proto

# 非当前目录的proto文件必须指定proto_path参数为文件所在路径
protoc --go_out=adx adx.proto && \
protoc --go_out=google google.proto && \
protoc --go_out=mega mega.proto && \
protoc --go_out=sohudsp sohudsp.proto && \
protoc --go_out=tanx tanx.proto && \
protoc --go_out=gdt gdt.proto && \
protoc --proto_path=../../util/web/proto --go_out=../../util/web ../../util/web/proto/info.proto
protoc --go_out=momo momo.proto
protoc --go_out=nativex nativex.proto
protoc --go_out=adinall adinall.proto
