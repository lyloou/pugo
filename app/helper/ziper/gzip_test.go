package ziper

import (
	"os"
	"testing"

	"encoding/base64"

	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	gzipString = "The compiler back end introduced in Go 1.7 for 64-bit x86 is now used on all architectures, and those architectures should see significant performance improvements. For instance, the CPU time required by our benchmark programs was reduced by 20-30% on 32-bit ARM systems. There are also some modest performance improvements in this release for 64-bit x86 systems. The compiler and linker have been made faster. Compile times should be improved by about 15% over Go 1.7. There is still more work to be done in this area: expect faster compilation speeds in future releases."
	gzipLength = 344
	gzipBase64 = "H4sIAAAAAAAC_3SRzWobQRCEX6UuvtnCP4kTcguG5BQIwXmA3p1a76CZ6U13r2S9fRhZVn7At4Eeuur7-nEmRq1LLjQMMm7BlpBbmKZ1ZH_iq-Jm8wGTGu7fXQ058PzxHtnRdI_VmaANUgrExjkHx1iNfglpCTGr898BfNa1JDgJz08tT3mUFlhok1qVNhK5LqY7VrbwDb6oITePPrpEzMTD95-IXAnjrzUbE4YDdDUMbONcxbZYTJ9MqmMvDuMLzXDA7fXV3fVFr3x3e4T5_OMb_ODB6hs8zrTel5DiCtdKVE30t_t1RTHnHlIozv9F_b37j-sup-S2pWGWHTGQDVUSMYkHbYOHl69HzrO04Zx9hJFB18DN-wvojna61CtFdnjkUlDViL3aFqF9Q9LGc2sxyifweeEYp-xTS4msDb6Q6Qg5rf1-r5i--R0AAP__kJqCeDwCAAA="
)

func TestGzip(t *testing.T) {
	Convey("Gzip", t, func() {
		data, err := Gzip([]byte(gzipString))
		So(err, ShouldBeNil)
		So(data, ShouldHaveLength, gzipLength)
	})

	Convey("GzipFile", t, func() {
		_, err := GzipFile("xyz.txt")
		So(err, ShouldNotBeNil)

		data, err := GzipFile("gzip_test.txt")
		So(err, ShouldBeNil)
		So(data, ShouldHaveLength, gzipLength)
	})

	Convey("GzipFileBase64", t, func() {
		_, err := GzipFileBase64("xyz.txt")
		So(err, ShouldNotBeNil)

		data, err := GzipFileBase64("gzip_test.txt")
		So(err, ShouldBeNil)
		So(data, ShouldEqual, gzipBase64)
	})
}

func TestUnGzip(t *testing.T) {
	Convey("UnGzip", t, func() {
		decodedBytes, _ := base64.URLEncoding.DecodeString(gzipBase64)
		data, err := UnGzip(decodedBytes)
		So(err, ShouldBeNil)
		So(string(data), ShouldEqual, gzipString)
	})

	Convey("UnGzipFile", t, func() {
		err := UnGzipFileBase64(gzipBase64, "gzip_test.txt2")
		So(err, ShouldBeNil)

		data, err := ioutil.ReadFile("gzip_test.txt2")
		So(err, ShouldBeNil)
		So(string(data), ShouldEqual, gzipString)

		os.Remove("gzip_test.txt2")
	})
}
