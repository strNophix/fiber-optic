from subprocess import run

GO_PACKAGES = [
    "github.com/gofiber/fiber/v2",
    "github.com/gofiber/swagger",
]

GO_INSTALL = [
    "github.com/swaggo/swag/cmd/swag@latest"
]

def main() -> int:
    run(["git", "init"])

    for package in GO_PACKAGES:
        run(["go", "get", "-u", package])

    for package in GO_INSTALL:
        run(["go", "install", package])

    run(["swag", "init", "-d", "internal/app"])

    return 0

if __name__ == "__main__":
    raise SystemExit(main())
