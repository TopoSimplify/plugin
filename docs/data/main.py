import json

from shapely import wkt
from shapely.geometry import mapping


def wkt2geojson(wkt_file, geojson_file):
    with open(wkt_file) as fid:
        lines = fid.readlines()
        with open(geojson_file, "w") as fout:
            for line in lines:
                line = wkt.loads(line)
                line = mapping(line)
                line = json.dumps(line)
                fout.write(line + "\n")


def sidedness():
    wkt2geojson("data/sidedness.wkt", "data/sidedness.json")
    wkt2geojson("data/sidedness_const.wkt", "data/sidedness_const.json")

def planar():
    wkt2geojson("data/planar.wkt", "data/planar.json")
    wkt2geojson("data/planar_const.wkt", "data/planar_const.json")

def non_planar():
    wkt2geojson("data/non_planar.wkt", "data/non_planar.json")
    wkt2geojson("data/non_planar_const.wkt", "data/non_planar_const.json")


if __name__ == '__main__':
    sidedness()
    planar()
    non_planar()
