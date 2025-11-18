package token

// create access token (jwt)
// sub -> user id; aud -> api/resource server (when registering app, give the aud, like finances.napps.com); exp -> expiration time in epoch (15 minutes);
// scope -> permissions (read:users, write:posts); iss -> issue (auth server identifier)
// stored client side in secure cookie, ideally
// sent back in response body

// create refresh token (opaque)
// expiration is 30 days
// HTTP-only secure cookie
// Store in DB

// create id token (jwt)
// sub -> user id; name, email, preferred_username -> identity claims; iat, exp -> issued at, expiration (15 minutes), iss -> issuer; aud -> audience (client app)
// only used by client-side app
// sent back in response body

// validate access token
