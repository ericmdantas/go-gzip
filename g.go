package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func main() {
	var b bytes.Buffer

	w := gzip.NewWriter(&b)
	w.Write([]byte(`
        alkdjakj kjasd
         sdj 
         asdj a
         dj askd akdj asj 
         sdj asd jsdkj aksdj ksdj a
         d asjk ajsdk asjd ajd asd çasd a
         sd 
         sdj sk as
         dj skj s
         d asjasj
          sdjks jasdk sdçask 
          asdk asdk lk
           sdklsk askd laskd çakd 
           asd asdk aldk as
           ç çj!
           alkdjakj kjasd
         sdj 
         asdj a
         dj askd akdj asj 
         sdj asd jsdkj aksdj ksdj a
         d asjk ajsdk asjd ajd asd çasd a
         sd 
         sdj sk as
         dj skj s
         d asjasj
          sdjks jasdk sdçask 
          asdk asdk lk
           sdklsk askd laskd çakd 
           asd asdk aldk as
           ç çj!
           alkdjakj kjasd
         sdj 
         asdj a
         dj askd akdj asj 
         sdj asd jsdkj aksdj ksdj a
         d asjk ajsdk asjd ajd asd çasd a
         sd 
         sdj sk as
         dj skj s
         d asjasj
          sdjks jasdk sdçask 
          asdk asdk lk
           sdklsk askd laskd çakd 
           asd asdk aldk as
           ç çj!
    `))

	w.Close()
	w.Flush()

	ioutil.WriteFile("g.gz", b.Bytes(), 0644)
}
