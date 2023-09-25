#!/bin/bash
docker container run --name hatchet --rm -it -e PORT=3000 -e MONGO_URI=mongodb://useradmin:thepianohasbeendrinking@host.docker.internal:27017 -p 3000:3000 hatchet
