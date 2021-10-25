package models

type Format string
type Currency string
type Gender string

const (
	Float      Format = "float"
	TCommerce5 Format = "tcommerce5"
	Opening    Format = "opening"
)

const (
	BRL Currency = "BRL"
)

type Source struct {
	Ds string `json:"ds,omitempty"`
}

type Deal struct {
	ID string `json:"id,omitempty"`
}

type Spec struct {
	Placement Placement `json:"placement,omitempty"`
	Sdk       string    `json:"sdk,omitempty"`
	TagID     string    `json:"tagid,omitempty"`
}

type Placement struct {
	Display Display `json:"display,omitempty"`
}

type Display struct {
	W   int64 `json:"w,omitempty"`
	H   int64 `json:"h,omitempty"`
	Pos int64 `json:"pos,omitempty"`
}

type Item struct {
	ID      int64    `json:"id,omitempty"`
	Qty     int64    `json:"qty,omitempty"`
	Flr     float64  `json:"flr,omitempty"`
	FlrCur  Currency `json:"flrcur,omitempty"`
	Deal    Deal     `json:"deal,omitempty"`
	Private int64    `json:"private,omitempty"`
	Spec    Spec     `json:"spec,omitempty"`
}

type Context struct {
	App    App    `json:"app,omitempty"`
	User   User   `json:"user,omitempty"`
	Device Device `json:"device,omitempty"`
}

type App struct {
	ID        string    `json:"id,omitempty"`
	Publisher Publisher `json:"publisher,omitempty"`
	Cat       []string  `json:"cat,omitempty"`
	Name      string    `json:"name,omitempty"`
}

type Publisher struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type User struct {
	ID       string   `json:"id,omitempty"`
	Data     []string `json:"data,omitempty"`
	Keywords string   `json:"keywords,omitempty"`
	Yob      int64    `json:"yob,omitempty"`
	Gender   Gender   `json:"gender,omitempty"`
	Geo      Geo      `json:"geo,omitempty"`
}

type Geo struct {
	Country string  `json:"country,omitempty"`
	Lon     float64 `json:"lon,omitempty"`
	City    string  `json:"city,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
	Zip     string  `json:"zip,omitempty"`
	Region  string  `json:"region,omitempty"`
	Type    int64   `json:"type,omitempty"`
}

type Device struct {
	Type    int64   `json:"type,omitempty"`
	Ua      string  `json:"ua,omitempty"`
	Dnt     int64   `json:"dnt,omitempty"`
	Lmt     int64   `json:"lmt,omitempty"`
	Make    string  `json:"make,omitempty"`
	Model   string  `json:"model,omitempty"`
	Os      string  `json:"os,omitempty"`
	Osv     string  `json:"osv,omitempty"`
	Hwv     string  `json:"hwv,omitempty"`
	H       int64   `json:"h,omitempty"`
	W       int64   `json:"w,omitempty"`
	Ppi     int64   `json:"ppi,omitempty"`
	PxRatio float64 `json:"pxratio,omitempty"`
	Js      int64   `json:"js,omitempty"`
	Lang    string  `json:"lang,omitempty"`
	Ip      string  `json:"ip,omitempty"`
	Ipv6    string  `json:"ipv6,omitempty"`
}

type Params struct {
	ID      string     `json:"id,omitempty"`
	Tmax    int64      `json:"tmax,omitempty"`
	At      int64      `json:"at,omitempty"`
	Cur     []Currency `json:"cur,omitempty"`
	Source  Source     `json:"source,omitempty"`
	Package int64      `json:"package,omitempty"`
	Item    []Item     `json:"item,omitempty"`
	Context Context    `json:"context,omitempty"`
}
