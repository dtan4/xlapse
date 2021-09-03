package main

import (
	"image"
	"image/gif"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAppend(t *testing.T) {
	testcases := map[string]struct {
		name  string
		delay int
	}{
		"success": {
			name:  "001.jpg",
			delay: 10,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			filename := filepath.Join("testdata", tc.name)
			body, err := ioutil.ReadFile(filename)
			if err != nil {
				t.Fatal(err)
			}

			g := &Gif{
				gif: &gif.GIF{},
			}

			err = g.Append(body, tc.delay)
			if err != nil {
				t.Errorf("want no error, got %q", err.Error())
			}
		})
	}
}

func TestSave(t *testing.T) {
	testcases := map[string]struct {
		source string
	}{
		"success": {
			source: "001.gif",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			filename := filepath.Join("testdata", tc.source)

			f, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			i, err := gif.Decode(f)
			if err != nil {
				t.Fatal(err)
			}

			g := &Gif{
				gif: &gif.GIF{},
			}
			g.gif.Image = append(g.gif.Image, i.(*image.Paletted))
			g.gif.Delay = append(g.gif.Delay, 10)

			tmpdir, err := ioutil.TempDir("", "TestSaveToFile")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tmpdir)

			out := filepath.Join(tmpdir, "TestSaveToFile.gif")
			fout, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				t.Fatal(err)
			}
			defer fout.Close()

			if err := g.Save(fout); err != nil {
				t.Errorf("want no error, got %q", err.Error())
			}

			if _, err := os.Stat(out); err != nil {
				t.Errorf("GIF file %q is not created", out)
			}
		})
	}
}
