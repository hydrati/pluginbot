package paspider

type PALocalizationDownloadEntry struct {
	Lang string `json:"lang"`
	Link string `json:"link"`
	Hash string `json:"hash"`
}

type PALocalizationDownloadEntryMap map[string]*PALocalizationDownloadEntry

type SpiderResult struct {
	Version string `json:"version"`
	Hash    string `json:"hash"` // MD5
	Url     string `json:"url"`
}
