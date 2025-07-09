#!/bin/bash

# Version bump script for twig
# Usage: ./scripts/bump-version.sh [major|minor|patch]

set -e

BUMP_TYPE=${1:-patch}
VERSION_FILE="VERSION"

if [ ! -f "$VERSION_FILE" ]; then
    echo "Error: VERSION file not found"
    exit 1
fi

# Read current version
CURRENT_VERSION=$(cat "$VERSION_FILE")
echo "Current version: $CURRENT_VERSION"

# Parse version components
IFS='.' read -r major minor patch <<< "$CURRENT_VERSION"

# Bump version based on type
case $BUMP_TYPE in
    major)
        NEW_MAJOR=$((major + 1))
        NEW_VERSION="$NEW_MAJOR.0.0"
        ;;
    minor)
        NEW_MINOR=$((minor + 1))
        NEW_VERSION="$major.$NEW_MINOR.0"
        ;;
    patch)
        NEW_PATCH=$((patch + 1))
        NEW_VERSION="$major.$minor.$NEW_PATCH"
        ;;
    *)
        echo "Error: Invalid bump type. Use major, minor, or patch"
        exit 1
        ;;
esac

echo "New version: $NEW_VERSION"

# Update VERSION file
echo "$NEW_VERSION" > "$VERSION_FILE"

# Build the binary with new version
echo "Building binary with version $NEW_VERSION..."
go build -ldflags="-s -w -X main.version=$NEW_VERSION" -o twig main.go generator.go

echo "Version bumped to $NEW_VERSION"
echo "To create a release, run:"
echo "  git add VERSION"
echo "  git commit -m \"chore: bump version to $NEW_VERSION\""
echo "  git tag -a \"v$NEW_VERSION\" -m \"Release v$NEW_VERSION\""
echo "  git push origin main --tags" 