import subprocess

GO_PACKAGES = [
    "github.com/gofiber/fiber/v2"
]

def main() -> int:
    for package in GO_PACKAGES:
        subprocess.run(["go", "get", "-u", package])
    return 0

if __name__ == "__main__":
    raise SystemExit(main())
