from pathlib import Path
import subprocess

BIN = 'bin.wasm'


def get_size(path: Path) -> int:
    # build binary
    cmd = [
        'tinygo', 'build',
        '-o', BIN,
        '-target', 'wasm-unknown',
        '-tags', f'none {path.stem}',
        str(path),
    ]
    subprocess.run(cmd, check=True)

    # strip debug symobols
    subprocess.run(['wasm-strip', BIN], check=True)

    # optimize for size
    cmd = ['wasm-opt', '-Oz', '--all-features', '-o', BIN, BIN]
    subprocess.run(cmd, check=True)

    size = len(Path(BIN).read_bytes())
    Path(BIN).unlink()
    return size


print('| function     | tinymath | stdlib | ratio |')
print('| ------------ | -------- | ------ | ----- |')
root = Path(__file__).parent / 'size_bench'
for tiny_path in (root / 'tiny').iterdir():
    std_path = root / 'std' / tiny_path.name
    tiny_size = get_size(tiny_path)
    std_size = get_size(std_path)
    ratio = int(tiny_size / std_size * 100)
    print(f'| {tiny_path.stem:12} | {tiny_size:>8} | {std_size:>6} | {ratio:>4}% |')  # noqa: E501
