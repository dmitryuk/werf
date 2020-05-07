	"path/filepath"

	"github.com/flant/werf/pkg/path_matcher"
func makeDiffParser(out io.Writer, pathMatcher path_matcher.PathMatcher) *diffParser {
		PathMatcher: pathMatcher,
	PathMatcher path_matcher.PathMatcher
			p.state = modifyFileDiff
			return p.handleIndexDiffLine(line)
		if strings.HasPrefix(line, "index ") {
			return p.handleIndexDiffLine(line)
		}
		if strings.HasPrefix(line, "index ") {
			return p.handleIndexDiffLine(line)
		}
		if strings.HasPrefix(line, "index ") {
			return p.handleIndexDiffLine(line)
		}
			if !p.PathMatcher.MatchPath(path) {
			newPath := p.trimFileBaseFilepath(path)
			if !p.PathMatcher.MatchPath(path) {
			newPath := p.trimFileBaseFilepath(path)
// TODO: remove index line from resulting patch completely in v1.2
func (p *diffParser) handleIndexDiffLine(line string) error {
	var prefix, hashes, suffix string

	parts := strings.SplitN(line, " ", 3)
	if len(parts) == 3 {
		prefix, hashes, suffix = parts[0], parts[1], parts[2]
	} else if len(parts) == 2 {
		prefix, hashes = parts[0], parts[1]
	} else {
		// unexpected format
		return p.writeOutLine(line)
	}

	hashesParts := strings.SplitN(hashes, "..", 2)
	if len(hashesParts) != 2 {
		// unexpected format
		return p.writeOutLine(line)
	}

	stripHashFunc := func(h string) string {
		// TODO: remove index line from resulting patch completely in v1.2
		if len(h) < 8 {
			return h
		}
		return h[:8]
	}

	var leftHashes []string
	for _, h := range strings.Split(hashesParts[0], ",") {
		leftHashes = append(leftHashes, stripHashFunc(h))
	}

	var rightHashes []string
	for _, h := range strings.Split(hashesParts[1], ",") {
		rightHashes = append(rightHashes, stripHashFunc(h))
	}

	var newLine string

	if suffix == "" {
		newLine = fmt.Sprintf("%s %s..%s", prefix, strings.Join(leftHashes, ","), strings.Join(rightHashes, ","))
	} else {
		newLine = fmt.Sprintf("%s %s..%s %s", prefix, strings.Join(leftHashes, ","), strings.Join(rightHashes, ","), suffix)
	}

	return p.writeOutLine(newLine)
}

	newPath := p.trimFileBaseFilepath(path)
	newPath := p.trimFileBaseFilepath(path)
	if strings.HasSuffix(line, " (commits not present)") {
		return fmt.Errorf("cannot handle \"commits not present\" in git diff line %q, check specified submodule commits are correct", line)
	}
	newPath := p.trimFileBaseFilepath(path)
func (p *diffParser) trimFileBaseFilepath(path string) string {
	return filepath.ToSlash(p.PathMatcher.TrimFileBaseFilepath(filepath.FromSlash(path)))
}

	newPath := p.trimFileBaseFilepath(path)