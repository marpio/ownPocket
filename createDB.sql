CREATE VIRTUAL TABLE bookamrks USING fts4(url, title, firstParagraph, readableContent, searchableContent, tokenize=porter);
