import os
import re
import matplotlib.pyplot as plt
from collections import defaultdict
import numpy as np

# === ПАРАМЕТРЫ ПОЛЬЗОВАТЕЛЯ ===
DATA_FILES = [
    "results/AggregateEff.txt",
    "results/filterEff.txt",
    "results/GenerateEff.txt",
    "results/MergeSort.txt",
    "results/QuickSortEff.txt",
]

FUNCTION_NAMES = [
    "Агрегация",
    "Фильтрация",
    "Генерация",
    "Сортировка слиянием",
    "Быстрая сортировка"
]

SELECTED_G_VALUES = [12, 500, 1000]
SELECTED_SIZE = 100000000  # <--- фиксированный размер выборки

# =============================

def parse_line(line):
    pattern = r'(\w+): ([\w\.\-NaN]+)'
    return {key: val for key, val in re.findall(pattern, line)}

def load_data():
    data_per_g = defaultdict(lambda: defaultdict(list))  # {g: {file_idx: [entry,...]}}
    for file_idx, file in enumerate(DATA_FILES):
        with open(file, 'r') as f:
            for line in f:
                if not line.strip():
                    continue
                entry = parse_line(line)
                try:
                    g = int(entry["g"])
                    size = int(entry["size"])
                    if size != SELECTED_SIZE:
                        continue  # пропустить записи с другим размером
                    data_per_g[g][file_idx].append({
                        "threads": int(entry["threads"]),
                        "size": size,
                        "timeP": float(entry["timeP"]) if entry["timeP"] != "NaN" else 0.0,
                        "SpeedUp": float(entry["SpeedUp"]) if entry["SpeedUp"] != "NaN" else 0.0,
                        "Eff": float(entry["Eff"]) if entry["Eff"] != "NaN" else 0.0,
                    })
                except (KeyError, ValueError):
                    continue
    return data_per_g



def plot_metrics(data_per_g, selected_g_values, function_names):
    metrics = {
        "timeP": "Время параллельного выполнения (сек)",
        "SpeedUp": "Ускорение",
        "Eff": "Эффективность"
    }

    for g in selected_g_values:
        if g not in data_per_g:
            print(f"[!] Нет данных для g = {g}")
            continue

        runs = data_per_g[g]

        for metric, ylabel in metrics.items():
            fig, ax = plt.subplots(figsize=(10, 6))

            for file_idx, entries in runs.items():
                entries = sorted(entries, key=lambda x: x["threads"])
                threads = [e["threads"] for e in entries]
                values = [e[metric] for e in entries]

                if not threads or not values:
                    continue

                label = function_names[file_idx] if file_idx < len(function_names) else f'Функция {file_idx+1}'
                ax.plot(threads, values, marker='o', label=label)

            ax.set_title(f'{ylabel} при g = {g}, size = {SELECTED_SIZE}')
            ax.set_xlabel('Количество потоков')
            ax.set_ylabel(ylabel)
            ax.legend()
            ax.grid(True)
            ax.set_xticks(sorted(set(threads)))
            ax.figure.canvas.manager.set_window_title(f'g = {g} | {metric}')
            plt.tight_layout()
            plt.show(block=True)
            
if __name__ == "__main__":
    data = load_data()
    plot_metrics(data, SELECTED_G_VALUES, FUNCTION_NAMES)
