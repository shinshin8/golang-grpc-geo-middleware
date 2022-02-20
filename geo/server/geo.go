package geo

import (
	"context"

	geo "github.com/martinlindhe/google-geolocate"
	"github.com/shinshin8/golang-grpc-middleware/geo/geoctx"
	"google.golang.org/grpc"
)

// ReverseGeocodeUnaryServerInterceptor returns a new unary interceptors that puts Geocode into context using Google Map API key.
func ReverseGeocodeUnaryServerInterceptor(apiKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		client := geo.NewGoogleGeo(apiKey)

		geolocation := geoctx.GetGeoLocate(ctx)

		res, err := client.ReverseGeocode(&geo.Point{
			Lat: geolocation["Lat"],
			Lng: geolocation["Lng"],
		})

		if err != nil {
			geoctx.SetGeo(ctx, "")
		} else {
			geoctx.SetGeo(ctx, res)
		}

		return handler(ctx, err)
	}
}

// GeocodeUnaryServerInterceptor returns a new unary interceptors that puts Geolocate into context using Google Map API key.
func GeocodeUnaryServerInterceptor(apiKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		client := geo.NewGoogleGeo(apiKey)

		geocode := geoctx.GetGeoInfo(ctx)

		res, _ := client.Geocode(geocode)
		if err != nil {
			geoctx.SetGeoLocate(ctx, 0, 0)
		} else {
			geoctx.SetGeoLocate(ctx, res.Lng, res.Lat)
		}

		return handler(ctx, err)
	}
}
