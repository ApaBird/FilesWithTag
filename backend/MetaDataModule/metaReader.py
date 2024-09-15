import hachoir.metadata
from hachoir.parser import createParser
from hachoir.metadata import extractMetadata
import hachoir
import os

def ReadMeta(path):
    parser = createParser(path)
    metadata = extractMetadata(parser)

    ans = {}
    for line in metadata.exportPlaintext():
        print("--")
        split = line.replace("-", " ").split(":")
        name = split[0].title().replace(" ", "")
        val = split[1].strip()
        ans[name] = val

    return ans



