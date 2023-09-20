# tesla-api

This application will connect to the Tesla API, makes a request, and shows the response.

Please see https://tesla-api.timdorr.com/ for details on using the Tesla API.

In order to make a request, `tesla-api` needs an access token.  This can be obtained in
several ways; please see https://tesla-api.timdorr.com/api-basics/authentication for details.

If you use `teslamate`, the access token can be obtained automatically from its
database; this is handled automatically if you provide the `docker-compose` file, or if
it is found in the current directory.

Usage examples:

- `tesla-api GET /api/1/users/me`
- `tesla-api GET /api/1/vehicles`
- `tesla-api GET /api/1/products`
- `tesla-api GET /api/1/vehicles/{id}/vehicle_data`
- `tesla-api GET /api/1/vehicles/929653650881721/nearby_charging_sites`
- `tesla-api GET /api/1/vehicles/929653650881721/release_notes`
- `tesla-api POST /api/1/vehicles/{id}/wake_up`
- `tesla-api POST /api/1/vehicles/{id}/honk_horn`
- `tesla-api POST /api/1/vehicles/{id}/flash_lights`
- `tesla-api POST /api/1/vehicles/{id}/remote_start_drive`
- `tesla-api POST /api/1/vehicles/{id}/door_unlock`
- `tesla-api POST /api/1/vehicles/{id}/door_lock`
- `tesla-api POST /api/1/vehicles/{id}/actuate_trunk '{"which": "rear"}`
- `tesla-api POST /api/1/vehicles/{id}/window_control '{"command": "close", "lat": 0, "lon": 0}`

