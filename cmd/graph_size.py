import matplotlib.pyplot as plt
import re

FUNCTION_NAMES = [
    "Агрегация",
    "Фильтрация",
    "Генерация",
    "Сортировка слиянием",
    "Быстрая сортировка"
]

DATA_FILES = [
    "results/Agg.txt",
    "results/Fil.txt",
    "results/Gen.txt",
    "results/MS.txt",
    "results/QS.txt",
]

# Конфигурация: имя функции -> (g, threads)
config = {
    "Агрегация": (500, 5),
    "Фильтрация": (500, 3),
    "Генерация": (500, 12),
    "Сортировка слиянием": (1000, 11),
    "Быстрая сортировка": (1000, 12),
}

# Размерности, по которым строим график
sizes = [5_000_000, 10_000_000, 15_000_000, 20_000_000, 25_000_000, 50_000_000, 75_000_000, 100_000_000]

# Функция парсинга строки
def parse_line(line):
    pattern = (
        r"size: (\d+), g: (\d+), threads: (\d+), timeS: ([\d.]+), "
        r"timeP: ([\d.]+), SpeedUp: ([\d.]+), Eff: ([\d.]+), trueg: (\d+)"
    )
    match = re.match(pattern, line.strip())
    if match:
        return {
            "size": int(match.group(1)),
            "g": int(match.group(2)),
            "threads": int(match.group(3)),
            "timeS": float(match.group(4)),
            "timeP": float(match.group(5)),
            "SpeedUp": float(match.group(6)),
            "Eff": float(match.group(7)),
            "trueg": int(match.group(8)),
        }
    return None

# Основная функция
def plot_all_speedups_on_one_figure():
    plt.figure(figsize=(10, 6))

    for i, (func_name, file_path) in enumerate(zip(FUNCTION_NAMES, DATA_FILES)):
        target_g, target_threads = config[func_name]

        with open(file_path, 'r', encoding='utf-8') as f:
            lines = f.readlines()

        data = [parse_line(line) for line in lines]
        filtered_data = {
            d["size"]: d["SpeedUp"]
            for d in data
            if d and d["g"] == target_g and d["threads"] == target_threads and d["size"] in sizes
        }

        # Сортировка и проверка
        sorted_sizes = sorted(filtered_data.keys())
        speedups = [filtered_data[size] for size in sorted_sizes]

        if not sorted_sizes:
            print(f"Нет данных для '{func_name}' при g={target_g}, threads={target_threads}")
            continue

        plt.plot(sorted_sizes, speedups, marker='o', label=func_name)

    # Оформление общего графика
    plt.title("Ускорение от размера данных")
    plt.xlabel("Размер данных")
    plt.ylabel("Ускорение")
    plt.xscale("log")
    plt.grid(True, which='both', linestyle='--', linewidth=0.5)
    plt.legend()
    plt.tight_layout()
    plt.show()

# Запуск
if __name__ == "__main__":
    plot_all_speedups_on_one_figure()
