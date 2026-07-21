# UserPreferences

The current user's personalization preferences.

Combines per-table column/density preferences and favorite repositories into a
single object, mirroring what the web app persists in local storage. Saving is a
full replace: the supplied object becomes the complete stored state.

Table preferences are persisted as a single JSON document
(:class:`~cloudsmith.users.models.UserTablePreference`), while favorites are
persisted relationally (:class:`~cloudsmith.users.models.UserFavoriteRepository`)
so they are cleaned up when a repository is deleted.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `TablePreferences`                                                                  | map[string][components.TablePreference](../../models/components/tablepreference.md) | :heavy_minus_sign:                                                                  | Table preferences, keyed by table identifier.                                       |
| `FavoriteRepositories`                                                              | []`string`                                                                          | :heavy_minus_sign:                                                                  | Identifiers (slug_perm) of the user's favorite repositories.                        |