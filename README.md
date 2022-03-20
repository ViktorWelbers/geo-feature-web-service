# Geo API Backend

This Project implements APIs that help to generate geo-spatial datasets for Machine Learning purposes

Frontend is tbd

To setup the PostGIS database you need to setup the following docker image runing as a service:

`docker pull postgis/postgis`

To make the service available in the local network a port mapping needs to be done from 5432 to a free port in the local network.
So in order to start a server you can run the following command:


`docker run --name postgis-server -e POSTGRES_PASSWORD=mysecretpassword -p 5433:5432 -d postgis/postgis `

Connect to postgis DB and create Hstore extension:

`sudo -u postgres psql -h localhost -p 5433`

 `CREATE extension hstore;`


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



To set up a local nominatim API in docker to query for postalcodes or cities for certain GPS-Coordinates, use the following command:

`
docker run -it --rm --shm-size=1g 
-e PBF_URL=https://download.geofabrik.de/europe/germany/nordrhein-westfalen-latest.osm.pbf 
-e REPLICATION_URL= http://download.geofabrik.de/europe/germany/nordrhein-westfalen-updates/ 
-e IMPORT_WIKIPEDIA=false 
-e NOMINATIM_PASSWORD=very_secure_password 
-v nominatim-data:/var/lib/postgresql/12/main 
-p 8080:8080 
--name nominatim 
mediagis/nominatim:3.7
`

