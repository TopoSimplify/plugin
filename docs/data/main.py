import json
import os.path

from shapely import wkt
from shapely.geometry import mapping, shape


def wkt2geojson(wkt_file, geojson_file):
    with open(wkt_file) as fid:
        lines = fid.readlines()
        with open(geojson_file, "w") as fout:
            for line in lines:
                line = wkt.loads(line)
                line = mapping(line)
                line = json.dumps(line)
                fout.write(line + "\n")


def geojson2wkt(geojson_file, wkt_file):
    with open(geojson_file) as fid:
        lines = fid.readlines()
        with open(wkt_file, "w") as fout:
            for line in lines:
                line = json.loads(line)
                line = shape(line['geometry'])
                fout.write(line.wkt + "\n")


def sidedness():
    wkt2geojson("data/sidedness.wkt", "data/sidedness.json")
    wkt2geojson("data/sidedness_const.wkt", "data/sidedness_const.json")


def planar():
    wkt2geojson("data/planar.wkt", "data/planar.json")
    wkt2geojson("data/planar_const.wkt", "data/planar_const.json")


def non_planar():
    wkt2geojson("data/non_planar.wkt", "data/non_planar.json")
    wkt2geojson("data/non_planar_const.wkt", "data/non_planar_const.json")


def feature_class():
    wkt2geojson("data/feature_class.wkt", "data/feature_class.json")
    wkt2geojson("data/feature_class_const.wkt", "data/feature_class_const.json")


def geojson_2_wkt(json_file):
    base, _ = os.path.splitext(json_file)
    wkt_file = base + ".wkt"
    geojson2wkt(json_file, wkt_file)


if __name__ == '__main__':
    # sidedness()
    # planar()
    # non_planar()
    # feature_class()

    geojson_2_wkt("output/out_feat_class.json")
