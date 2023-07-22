import subprocess

GO_PACKAGES = [
    "github.com/gofiber/fiber/v2"
]

def main() -> int:
    subprocess.run(["git", "init"])

    for package in GO_PACKAGES:
        subprocess.run(["go", "get", "-u", package])
    return 0

if __name__ == "__main__":
    raise SystemExit(main())
