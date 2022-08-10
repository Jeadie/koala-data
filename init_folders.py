from typing import Dict, List
import os


def main():
    koala_folder = {
        "GPS": ["Input", "Output", "Graph"],
        "Scatter": ["Input", "Output", "Graph"],
        "Habitat":  ["Input", "Output", "Graph"],
    }
    create_folder_structure("./", koala_folder)

def create_folder_structure(base_folder: str, folder_structure: Dict[str, List[str]]):
    for k, v in folder_structure.items():
        for i in v:
            path = f"{base_folder}/{k}/{i}"
            os.makedirs(path)

if __name__ == "__main__":
    main()