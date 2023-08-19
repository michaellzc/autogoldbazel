# autogold 🔫 bazel

This repo demonstrates how autogold doesn't work in bazel mode even with `ENABLE_BAZEL_PACKAGES_LOAD_HACK=true` is set.

First, verify it works fine with vanila Go:

```sh
go test -v ./version/.
```

Then, verify it fails with bazel:

```sh
bazel test //version/...
```

The error is:

```diff
        -- want
        ++ got
        @@ -1,9 +1 @@
        &version.Version{semver: &version.Version{
                segments: []int64{
                        1,
                        2,
                        3,
                },
                si:       3,
                original: "1.2.3",
        }}
        valast: format: 4:33: expected operand, found 'go'
```

`found 'go'`?

It turned out the bazel hack from `autogold`, [bazelPackagePathToName] infers the package name from the last component of the full package name. 

In this example, we use a package `github.com/hashicorp/go-version` but the package name is actually `package version`. 

Therefore, autogold will return `go-version` when `bazelPackagePathToName` is called. Obviously, `go-version` is not valid syntax for Go.

## FAQ

### How to run local copy of `autogold` and `valast`?

Clone them somewhere, and run:

```sh
cd valast
gazelle -repo_root=. -go_prefix=github.com/hexops/valast
touch WORKSPACE
```

```sh
cd autogold
gazelle -repo_root=. -go_prefix=github.com/hexops/autogold/v2
touch WORKSPACE
```

Then, replace `"@com_github_hexops_valast//:go_default_library"` with `"@com_github_hexops_valast//:valast",` in `autogold/BUILD.bazel`. 

Modify `local_repositiry` in `WORKSPACE` and `replace` directive in `go.mod`.

[bazelPackagePathToName]: https://github.com/hexops/autogold/blob/fbc07f54ce22cbc86b340025df0e15f699ddc267/bazel.go#L93-L108
