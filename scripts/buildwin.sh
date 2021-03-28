#! /bin/bash

#embed config.yml
packr2

# embed icon in the executable
rsrc -ico assets/favicon.ico

curDir=$(pwd)
today=$(date +'%F')
releaseDir=$curDir/releases/$today/windows
mkdir -p $releaseDir/{config,site}
mkdir -p $releaseDir/

GOOS=windows GOARCH=386 go build .
cp -r config ${releaseDir}/
cp -r site ${releaseDir}/
cp $curDir/{runWithConfig.bat,runWithoutSearch.bat} ${releaseDir}/
mv $curDir/jobScrapper.exe ${releaseDir}/jobScrapper.exe
