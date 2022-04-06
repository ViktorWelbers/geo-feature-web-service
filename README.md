# Geo API Backend

This Project implements APIs that help to generate geo-spatial datasets for Machine Learning purposes

To receive a feature vector for a specific lat-long and radius (in Meters) the following GET-Requests are available:

To Start the service, type `docker-compose up` at root level

<h2>DB Setup</h2>

The PostGIS database you need, is already accounted for in the docker-compose.yaml:
Credentials are:

> password = mysecretpassword
> user = postgres
> port = 5433

To load the OpenStreetMap data into PostGIS from scratch the following procedure needs to be done:

1. Download osm.pbf file from https://download.geofabrik.de/europe/germany/nordrhein-westfalen-latest.osm.pbf 

2. Install osmctools (on Ubuntu):

    `sudo apt install osmctools && wget -O - http://m.m.i24.cc/osmconvert.c | sudo cc -x c - -lz -O3 -o osmconvert`

3. Install osm2pgsql (on Ubuntu):

    `sudo apt-get install -y osm2pgsql`


4. Convert nordrhein-westfalen-latest.osm.pbf to "only nodes" for easy queries over single geometry:

    `osmconvert nordrhein-westfalen-latest.osm.pbf  --all-to-nodes --max-objects=180000000  --object-type-offset=18000000000+1 -o=nrw_nodes.osm`

5. Load nrw_nodes.osm from previous step into the PostGIS db:

    `osm2pgsql -U postgres -W -d postgres -H localhost -s -P 5433  -l nrw_nodes.osm`

<h2>Usage</h2>

To get a JSONL with aggregated entries for every feature:

`/v1/{LONGTITUDE}/{LATITUDE}/{RADIUS}`

<br/><br/>

To get a JSON that is easier to use for a frontend framework (such as React, Angular etc) you can use :

`/v2/{LONGTITUDE}/{LATITUDE}/{RADIUS}`

Here a JSON with a single key `data` is returned that stores the key value pairs


