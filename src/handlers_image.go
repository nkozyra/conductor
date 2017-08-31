package conductor

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"

	"github.com/disintegration/gift"
	"github.com/gorilla/mux"
	_ "github.com/nfnt/resize"
)

func UserImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["id"]

	id, _ := strconv.ParseInt(sid, 10, 32)

	data := GetImage(id)

	decData, _ := base64.StdEncoding.DecodeString(data)

	im, _, err := image.Decode(bytes.NewReader(decData))
	if err != nil {

	}
	g := gift.New(
		gift.ResizeToFill(int(80), int(80), gift.NearestNeighborResampling, gift.TopAnchor),
	)
	dst := image.NewRGBA(g.Bounds(im.Bounds()))
	g.Draw(dst, im)

	buf := new(bytes.Buffer)
	op := jpeg.Options{Quality: 60}
	err = jpeg.Encode(buf, dst, &op)
	output := buf.Bytes()

	w.Header().Set("Content-type", "image/png")
	fmt.Fprintln(w, string(output))

}
