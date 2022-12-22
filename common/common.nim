
proc getFile*(fileName: string) : seq[string] =
    var lines : seq[string]
    let f = open(fileName)
    defer:
        f.close()

    var line : string
    while f.readLine(line):
        lines.add(line)

    lines