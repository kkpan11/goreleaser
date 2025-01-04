# Deno

<!-- md:version v2.6-unreleased -->

<!-- md:alpha -->

You can now build TypeScript binaries using `deno compile` and GoReleaser!

Simply set the `builder` to `deno`, for instance:

```yaml title=".goreleaser.yaml"
builds:
  # You can have multiple builds defined as a yaml list
  - #
    # ID of the build.
    #
    # Default: Project directory name.
    id: "my-build"

    # Use deno.
    builder: deno

    # Binary name.
    # Can be a path (e.g. `bin/app`) to wrap the binary in a directory.
    #
    # Default: Project directory name.
    binary: program

    # List of targets to be built, in Deno's format.
    #
    # See: https://docs.deno.com/runtime/reference/cli/compile/#supported-targets
    # Default: [ "x86_64-pc-windows-msvc", "x86_64-apple-darwin", "aarch64-apple-darwin", "x86_64-unknown-linux-gnu", "aarch64-unknown-linux-gnu" ]
    targets:
      - linux-x64-modern
      - darwin-arm64

    # Path to project's (sub)directory containing the code.
    # This is the working directory for the `deno compile` command(s).
    #
    # Default: '.'.
    dir: my-app

    # Main entry point.
    #
    # Default: 'main.ts'.
    main: "file.ts"

    # Set a specific deno binary to use when building.
    # It is safe to ignore this option in most cases.
    #
    # Default: 'deno'.
    # Templates: allowed.
    tool: "deno-wrapper"

    # Sets the command to run to build.
    #
    # Default: 'compile'.
    command: not-build

    # Custom flags.
    #
    # Templates: allowed.
    # Default: [].
    flags:
      - --allow-all

    # Custom environment variables to be set during the builds.
    # Invalid environment variables will be ignored.
    #
    # Default: os.Environ() ++ env config section.
    # Templates: allowed.
    env:
      - FOO=bar
```

Some options are not supported yet[^fail], but it should be usable for
most projects already!

You can see more details about builds [here](./builds.md).

## Caveats

GoReleaser will translate Deno's Os/Arch pair into a GOOS/GOARCH pair, so
templates should work the same as before.
The original target name is available in templates as `.Target`, and so is the
the ABI and Vendor as `.Abi` and `.Vendor`, respectively.

[^fail]:
    GoReleaser will error if you try to use them. Give it a try with
    `goreleaser r --snapshot --clean`.

<!-- md:templates -->