import re

DATA_FILES = [
    "copy_res/AggregateEff.txt",
    "copy_res/filterEff.txt",
    "copy_res/GenerateEff.txt",
    "copy_res/MergeSort.txt",
    "copy_res/QuickSortEff.txt",
]

SELECTED_G_VALUES = [12, 500, 1000]

def parse_line(line):
    """Парсит строку и возвращает словарь параметров."""
    pattern = (
        r"size: (\d+), g: (\d+), threads: (\d+), timeS: ([\d.]+), "
        r"timeP: ([\d.]+), SpeedUp: ([\d.]+), Eff: ([\d.]+), trueg: (\d+)"
    )
    match = re.match(pattern, line.strip())
    if match:
        keys = ["size", "g", "threads", "timeS", "timeP", "SpeedUp", "Eff", "trueg"]
        values = [int(match.group(1)), int(match.group(2)), int(match.group(3)),
                  float(match.group(4)), float(match.group(5)), float(match.group(6)),
                  float(match.group(7)), int(match.group(8))]
        return dict(zip(keys, values))
    return None

def find_min_time_params():
    results = {}

    for file in DATA_FILES:
        with open(file, "r") as f:
            lines = f.readlines()

        parsed = [parse_line(line) for line in lines]
        parsed = [entry for entry in parsed if entry and entry["g"] in SELECTED_G_VALUES]

        file_results = {}
        for g_value in SELECTED_G_VALUES:
            g_entries = [entry for entry in parsed if entry["g"] == g_value]
            if g_entries:
                best_entry = min(g_entries, key=lambda x: x["timeP"])
                file_results[g_value] = best_entry

        results[file] = file_results

    return results

# Пример использования:
if __name__ == "__main__":
    import pprint
    pprint.pprint(find_min_time_params())
