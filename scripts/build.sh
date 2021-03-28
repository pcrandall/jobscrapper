#! /bin/bash
go mod tidy

curDir=$(pwd)
today=$(date +'%F')
releaseDir=$curDir/releases/$today
# mkdir -p $releaseDir/{config,site}
mkdir -p $releaseDir/


#embed files into binary
packr2

# GOOS=linux GOARCH=386 go build .
go build .
cp -r config ${releaseDir}/
cp -r site ${releaseDir}/
mv jobScrapper ${releaseDir}/jobScrapper
