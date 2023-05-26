package version

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func Update(file, version string) error {
	// get the file type
	ext := filepath.Ext(file)

	var replacer func([]byte, string) ([]byte, error)

	switch ext {
	case ".json":
		replacer = updateJsonVersion

	case "./gradle":
		replacer = updateGradleVersion

	case ".rb":
		replacer = updateRubyFileVersion

	default:
		return fmt.Errorf("unable to update %s: unknown file type %s", file, ext)
	}

	buff, err := readFile(file)
	if err != nil {
		return nil
	}
	update, err := replacer(buff, version)
	if err != nil {
		return err
	}
	if err := writeUpdate(update, file); err != nil {
		return err
	}
	return nil
}

func updateJsonVersion(json []byte, version string) ([]byte, error) {
	re := regexp.MustCompile(`("version"\s*:\s*)"(?:.*)"`)

	if match := re.Match(json); !match {
		return nil, errors.New("cannot find a version in the json file")
	}
	repl := fmt.Sprintf(`$1"%s"`, version)
	return re.ReplaceAll(json, []byte(repl)), nil
}

func updateGradleVersion(gradle []byte, version string) ([]byte, error) {
	re := regexp.MustCompile(`(gutenbergMobileVersion\s*=\s*)'(?:.*)'`)

	if match := re.Match(gradle); !match {
		return nil, errors.New("cannot find a version in the gradle file")
	}

	repl := fmt.Sprintf(`$1'%s'`, version)
	return re.ReplaceAll(gradle, []byte(repl)), nil
}

func updateRubyFileVersion(rubyConfig []byte, version string) ([]byte, error) {
	versionRegexp := regexp.MustCompile(`v\d+\.\d+\.\d+`)
	tagLineRegexp := regexp.MustCompile(`([\r\n]\s*)tag:.*`)
	commitLineRegexp := regexp.MustCompile(`([\r\n]\s*)#*\s*commit:.*`)

	if versionRegexp.MatchString(version) {

		tagRepl := []byte(fmt.Sprintf(`${1}tag: '%s'`, version))

		if commitLineRegexp.Match(rubyConfig) {

			updated := tagLineRegexp.ReplaceAll(rubyConfig, []byte(""))
			return commitLineRegexp.ReplaceAll(updated, tagRepl), nil
		} else {
			return tagLineRegexp.ReplaceAll(rubyConfig, tagRepl), nil
		}

	} else {
		commitRepl := []byte(fmt.Sprintf(`${1}commit: '%s'`, version))

		if tagLineRegexp.Match(rubyConfig) {
			updated := commitLineRegexp.ReplaceAll(rubyConfig, []byte(""))
			return tagLineRegexp.ReplaceAll(updated, commitRepl), nil
		} else {
			return commitLineRegexp.ReplaceAll(rubyConfig, commitRepl), nil
		}
	}
}

func readFile(file string) ([]byte, error) {
	buff, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func writeUpdate(update []byte, file string) error {
	f, err := os.Create(file)
	defer f.Close()

	if err != nil {
		return err
	}

	if _, err := f.Write(update); err != nil {
		return err
	}

	return nil
}
