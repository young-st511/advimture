#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "usage: scripts/release.sh major|minor|patch" >&2
}

bump="${1:-}"
case "$bump" in
  major|minor|patch) ;;
  *)
    usage
    exit 2
    ;;
esac

if ! git diff --quiet || ! git diff --cached --quiet; then
  echo "release requires a clean working tree. Commit the completed ExecPlan first." >&2
  exit 1
fi

branch="$(git branch --show-current)"
if [ "$branch" != "main" ]; then
  echo "release must run from main; current branch is $branch" >&2
  exit 1
fi

git fetch --tags origin

latest_tag="$(git tag --list 'v[0-9]*.[0-9]*.[0-9]*' --sort=-v:refname | head -n 1 || true)"
if [ -n "$latest_tag" ]; then
  base="${latest_tag#v}"
else
  base="$(tr -d '[:space:]' < VERSION)"
fi

if [[ ! "$base" =~ ^[0-9]+[.][0-9]+[.][0-9]+$ ]]; then
  echo "invalid base version: $base" >&2
  exit 1
fi

IFS=. read -r major minor patch <<<"$base"
case "$bump" in
  major)
    major=$((major + 1))
    minor=0
    patch=0
    ;;
  minor)
    minor=$((minor + 1))
    patch=0
    ;;
  patch)
    patch=$((patch + 1))
    ;;
esac

next="$major.$minor.$patch"
tag="v$next"

if git rev-parse "$tag" >/dev/null 2>&1; then
  echo "tag already exists: $tag" >&2
  exit 1
fi

printf "%s\n" "$next" > VERSION
make release-check

git add VERSION
git commit -m "Release $tag"
git tag -a "$tag" -m "Advimture $tag"

git push origin "$branch"
git push origin "$tag"

echo "Released $tag. GitHub Actions will publish the release and update Homebrew."
