package geo

import (
	"context"

	"github.com/shinshin8/golang-grpc-middleware/server/ctxgeo"

	geo "github.com/martinlindhe/google-geolocate"
	"google.golang.org/grpc"
)

// Reverse geocode
func UnaryServerInterceptor(k string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		client := geo.NewGoogleGeo(k)

		geolocation := ctxgeo.GetGeoLocate(ctx)

		res, _ := client.ReverseGeocode(&geo.Point{
			Lat: geolocation["Lat"],
			Lng: geolocation["Lng"],
		})

		ctxgeo.SetGeo(ctx, res)

		return handler(ctx, err)
	}
}
