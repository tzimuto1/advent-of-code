
def read_file(file, is_file_name=True):
    if not is_file_name:
        return file.rstrip()
    with open(file, "r") as f:
        return f.read().rstrip()

def read_lines(file, is_file_name=True):
    return read_file(file, is_file_name).split("\n")

def read_grid(file, is_file_name=True):
    lines = read_lines(file, is_file_name)
    return [[c for c in line] for line in lines]


if __name__ == "__main__":
    pass