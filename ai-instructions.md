## Handling Structs and Pointers

When working with structs in Go, always treat structs as pointer receivers, pointer arguments, and pointer return values by default.
Additionally, use pointer types for struct fields whenever possible.

The main reason is to maintain consistency so that client code always handles structs through pointers.
Efficiency gained from pointer usage is the second reason. If working with pointers makes the code harder to use or read, feel free to use value types instead.

Basically, structs should be handled using pointers, but there are exceptions.
If there are established conventions or good practices in Go that specify whether to use values or pointers, prioritize following those conventions/good practices.
For example, `sync.RWMutex` is almost never stored as a pointer in a struct field.

In all other cases, use value types only when the drawbacks of using pointers outweigh the benefits.

Please treat primitive values as value types by default.

## Function and Method Signatures

Aim for at most four parameters and at most four return values per function or method. If either count would exceed four, look for a design that is easier to understand and use at the call site. Group related parameters or return values into a struct when they form a coherent concept. Consider splitting the function when it has distinct responsibilities.

A larger signature is acceptable when explicit parameters or return values make the call site clearer, when the signature follows an established Go convention, or when an interface or callback requires it.

If known requirements make additional parameters or return values likely, consider them when designing the initial signature.

## Use constants whenever possible.

If a function repeatedly uses fixed strings, numbers, time.Duration, or similar values, define them as constants whenever possible and use those constants within the function.
The same applies to defining default values.

## About Backward Compatibility

This project has not been released yet.
Therefore, you do not need to write code that considers backward compatibility, such as providing function aliases.

## About Tests

Please write tests using Ginkgo.
If the project does not already include `github.com/onsi/gomega` and `github.com/onsi/ginkgo`, install them when creating tests.

Please create test files on a per-file basis, rather than having a single test file per package.

For each package, create exactly one Ginkgo *_suite_test.go file.
Always keep the *_suite_test.go file separate from the test files that contain the actual test cases.

When writing tests, be aware that the implementation may be incorrect. In such cases, do not write tests just to make them pass. Instead, implement the test as the correct and intended test, even if it would fail, and mark the test as skipped.
Additionally, add a comment to that test saying: FIXME: consider fixing the implementation. After that, ask me whether the implementation should be fixed so that the test can pass.

## What to Include After Code Fixes or Implementation

After making changes, please always include the key points to review. Also mention anything else that caught your attention while modifying the code, as well as any areas where the correctness of the implementation should be verified.

If the changes are extremely minor, you do not need to include review points. Instead, simply provide a brief summary of what changed.

Please also provide a concise, one-line Git commit message.

## About Language

Although the instructions and conversation will be in Japanese, please write all code comments in English. The implementation output may use Japanese or other languages, but all code comments should be in English.

Existing README files may be written in either Japanese or English.

Although the languages are currently mixed, they will eventually be separated into `README.md` for the English version and `README.ja.md` for the Japanese version.

For now, write all new content added to `README.md` in English. Leave the existing Japanese content unchanged.

If there is a special need to use another language, I will explicitly instruct you to do so.

## When I Need to Make a Decision

When I need to choose among several options, please explain the factors I should consider, along with the advantages and disadvantages of each option.

When recommending an option, do not treat a smaller scope or a lower volume of changes as an advantage or use the amount of required modification as a primary decision factor.

In general, prioritize the option that results in better specifications, design, and code quality, even if it requires broader or more extensive changes.

Consider compatibility loss a disadvantage only when the proposed changes would cause significant compatibility issues. Compatibility should be taken into account only when the current version is `v1.0.0` or later.

The current version is still below `v1.0.0`, so backward compatibility does not currently need to be prioritized.

## When Refining a Plan

When working on a task to refine a plan, the plan may contain ambiguous points or unresolved decisions.

When you need me to make a decision about an unresolved point, do not use a selection dialog. Selection dialogs provide too little space for detailed explanations, which can make it easy to overlook important considerations.

At the beginning of the refinement process, if there are decisions that require my input, first summarize all of the relevant discussion points and unresolved issues in a list.

I will then respond to each point in the chat and provide my decisions there.

## Git Operations

Git may be used only for read-only inspection of the repository.

You may run commands that inspect the current state, differences, or history, such as:

* `git status`
* `git diff`
* `git log`
* `git show`
* `git blame`
* `git branch --show-current`

Do not perform any Git operation that modifies the working tree, staging area, branches, tags, repository history, remotes, or Git configuration.

In particular, do not run commands such as:

* `git add`, `git commit`, or `git stash`
* `git checkout`, `git switch`, or `git restore`
* `git reset`, `git revert`, `git merge`, or `git rebase`
* `git branch`, `git tag`, or commands that create, rename, or delete references
* `git push`, `git pull`, or `git fetch`
* `git config`
* Any command that modifies files under the `.git` directory

Do not create commits or otherwise change Git state unless the user explicitly instructs you to do so.

## Docker

You may freely start and stop Docker containers as needed.

If running containers would be helpful for verifying that the application works correctly, please do so and perform the necessary tests.

If you create any unnecessary data or other temporary resources during testing, please clean them up as much as possible afterward.

## References to TODO Files, Implementation Plans, and Design Documents

TODO files, implementation plans, and design documents will not be made public under any circumstances.

Therefore, do not include any direct or indirect references to these files or documents in either source code or comments.

### References in Source Code

Do not include identifiers, string literals, constants, or other code elements that refer to these files or documents.

### References in Comments

Do not include references to these files or documents in code comments, including:

- File names, file paths, URLs, or document titles that identify these files or documents.
- Explicit references, such as "See TODO.md" or "As described in the implementation plan."
- Indirect references that assume the reader has access to these documents, such as "Implemented according to the design document" or "See Phase 2 of the plan."

Comments must be self-contained and understandable without access to these documents. Ordinary TODO comments describing pending work are allowed as long as they do not reference private files or documents.

### Temporary References

You may temporarily include such references during implementation if necessary, but you must ensure that all of them are completely removed before completing the task.