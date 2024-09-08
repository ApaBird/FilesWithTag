from hachoir.parser import createParser
from hachoir.metadata import extractMetadata

def handlerMeta(order):
    if order.find("GetMeta") != -1:
        print(order)
        return ReadMeta(order.replace("GetMeta", ""))
    else:
        return None

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



