package geoctx

import "context"

const (
	locateKey = "geo-locate-key"
	geoKey    = "geo-key"
)

func SetGeo(parent context.Context, g string) context.Context {
	return context.WithValue(parent, geoKey, g)
}

func GetGeoInfo(ctx context.Context) string {
	v := ctx.Value(geoKey)

	geo, ok := v.(string)
	if !ok {
		return ""
	}
	return geo
}

func SetGeoLocate(parent context.Context, lng, lat float64) context.Context {
	return context.WithValue(
		parent,
		locateKey,
		map[string]float64{
			"Lat": lat,
			"Lng": lng,
		})
}

func GetGeoLocate(ctx context.Context) map[string]float64 {
	v := ctx.Value(locateKey)

	gloc, ok := v.(map[string]float64)
	if !ok {
		return map[string]float64{}
	}
	return gloc
}
