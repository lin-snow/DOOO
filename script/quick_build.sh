# Description: Quick build script is used to build the project without docker.

# clean the app directory
cd ../app && rm -rf ./*

# add dist & config directory
mkdir dist && mkdir config

# build Frontend Directory
cd ../Web/v1

# build Frontend
npm run build

# copy Frontend dist to app directory
cp -r dist/* ../../app/dist

# build Backend Directory
cd ../../

# copy config file to app directory
cp config/config.yaml app/config

# build Backend
go build main.go

# copy Backend to app directory
cp main.exe ./app

# change directory to app
cd app

# Start the app
./main.exe



