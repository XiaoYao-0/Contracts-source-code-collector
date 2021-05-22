#!/bin/bash

if [ ! -d "CollectTool"  ];then
  mkdir "CollectTool"
fi
touch CollectTool/conf.json
echo "{\"contract_number\":0,\"proxy\":\"\",\"api_key\":\"\",\"storage_dir\":\"./contracts\"}" > CollectTool/conf.json
go build -o CollectTool/Collect
chmod -R 777 CollectTool
