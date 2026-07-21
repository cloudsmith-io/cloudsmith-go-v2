# Geolocation

Geolocation data of the request IP address.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `City`                                      | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | City.                                       |
| `Country`                                   | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | Country.                                    |
| `CountryCode`                               | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | Country code (alpha2 ISO-3166-1).           |
| `Region`                                    | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | Region.                                     |
| `IPAddress`                                 | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | IP address. IPv4 or IPv6.                   |