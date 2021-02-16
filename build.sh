#!/bin/bash
ruta=$(pwd)"/starwars";
export GOOS=linux;
export GOARCH=amd64;
cd $ruta;
go build;
zip starwars.zip starwars
aws s3api put-object --bucket starwars-mercadolibre --key stawars.zip --body starwars.zip;