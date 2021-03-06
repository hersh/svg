package svg

import mt "github.com/rustyoz/Mtransform"

type Circle struct {
	Id        string `xml:"id,attr"`
	Transform string `xml:"transform,attr"`
	Style     string `xml:"style,attr"`
	Cx        string `xml:"cx,attr"`
	Cy        string `xml:"cy,attr"`
	Radius    string `xml:"r,attr"`

	transform mt.Transform
	group     *Group
}
