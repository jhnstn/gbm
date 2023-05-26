package version

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/andreyvit/diff"
)

func TestVersionUpdates(t *testing.T) {
	t.Run("It returns an error if the file type is unrecognized", func(t *testing.T) {
		if err := Update("what-the.yml", "1.1.2.3"); err == nil {
			t.Fatal("Expected an error")
		}
	})
}

func TestJSONUpdates(t *testing.T) {

	assertVersionChange := buildVersionAsserter(t, updateJsonVersion)

	t.Run("It updates versions in package.json files", func(t *testing.T) {
		expected := "package.version-9.9.9.json"
		target := "package.json"
		assertVersionChange(expected, target, "9.9.9")
	})

	t.Run("It honors whitespace in json files", func(t *testing.T) {
		testJson := []byte(`{"version"      :      "1.0.0"}`)

		got, err := updateJsonVersion(testJson, "9.9.9")
		assertNoError(t, err)

		want := []byte(`{"version"      :      "9.9.9"}`)
		assertNoDiff(t, got, want)
	})

	t.Run("It fails if a version can't be found in the json", func(t *testing.T) {
		testJson, err := os.ReadFile("testdata/package-no-version.json")
		if err != nil {
			t.Fatal(err)
		}

		if _, err := updateJsonVersion(testJson, "9.9.9"); err == nil {
			t.Fatal("Expected an error")
		}
	})
}

func TestGradleUpdates(t *testing.T) {

	assertVersionChange := buildVersionAsserter(t, updateGradleVersion)

	t.Run("It only updates the Gutenberg version", func(t *testing.T) {
		expected := "build.version-9.9.9.gradle"
		target := "build.gradle"
		assertVersionChange(expected, target, "v9.9.9")
	})

	t.Run("It matches alpha versions", func(t *testing.T) {
		expected := "build.version-9.9.9.gradle"
		target := "build.alpha.gradle"
		assertVersionChange(expected, target, "v9.9.9")
	})

	t.Run("It updates from a pr to a tag", func(t *testing.T) {
		expected := "build.version-9.9.9.gradle"
		target := "build-pr-123-asdfgh.gradle"
		assertVersionChange(expected, target, "v9.9.9")
	})

	t.Run("It updates from a tag to a pr", func(t *testing.T) {
		expected := "build-pr-123-asdfgh.gradle"
		target := "build.version-9.9.9.gradle"
		assertVersionChange(expected, target, "123-asdfgh")
	})

	t.Run("It errors if it can't find a gutenberg mobile version", func(t *testing.T) {
		testGradle, err := os.ReadFile("testdata/build-no-gb.gradle")
		if err != nil {
			t.Fatal(err)
		}

		if _, err := updateGradleVersion(testGradle, "v9.9.9"); err == nil {
			t.Fatal("Expected an error")
		}

	})

}

func TestRubyConfigUpdates(t *testing.T) {

	assertVersionChange := buildVersionAsserter(t, updateRubyFileVersion)
	t.Run("It updates to a tag from a commit", func(t *testing.T) {

		expected := "version-tag-v9.9.9.rb"
		target := "version-commit-e765ad9.rb"
		assertVersionChange(expected, target, "v9.9.9")
	})

	t.Run("It updates to a new tag", func(t *testing.T) {
		expected := "version-tag-v9.9.9.rb"
		target := "version.rb"

		assertVersionChange(expected, target, "v9.9.9")

	})

	t.Run("It updates to a commit from a tag", func(t *testing.T) {
		expected := "version-commit-e765ad9.rb"
		target := "version-tag-v9.9.9.rb"
		assertVersionChange(expected, target, "e765ad9")
	})

	t.Run("It updates to a new commit", func(t *testing.T) {
		expected := "version-commit-v999.rb"
		target := "version-commit-e765ad9.rb"
		assertVersionChange(expected, target, "v999")
	})

}

func buildVersionAsserter(t testing.TB, updater func([]byte, string) ([]byte, error)) func(string, string, string) {
	return func(expected, target, version string) {
		t.Helper()
		want, err := os.ReadFile(filepath.Join("testdata", expected))
		if err != nil {
			t.Fatal(err)
		}

		targetConfig, err := os.ReadFile(filepath.Join("testdata", target))

		if err != nil {
			t.Fatal(err)
		}

		got, err := updater(targetConfig, version)

		assertNoError(t, err)
		assertNoDiff(t, got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't expect one")
	}
}

func assertNoDiff(t testing.TB, got, want []byte) {
	t.Helper()
	if g, w := string(got), string(want); g != w {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(g, w))
	}
}

// func errorWithDiff(t testing.TB, got, want []byte) {
// t.Helper()
// t.Errorf("Result not as expected:\n%v", diff.LineDiff(string(got), string(want)))
// }
