# State

## Server configuration

Getting important server information: version, limits, features, etc.

Can be used for checking server availability (ping). Works without authentication.

### Request
{% include "../codegen/samples/state.request.yml" %}

### Response
{% include "../codegen/samples/state.response.yml" %}

## Emulate server errors

### Request
{% include "../codegen/samples/fake-error.request.yml" %}

### Response
{% include "../codegen/samples/fake-error.response.yml" %}

## Global badge
Number of unread notifications. Should be displayed as a badge on 
the application icon, on browser tab, etc.

### Request
{% include "../codegen/samples/badge.request.yml" %}

### Response
{% include "../codegen/samples/badge.response.yml" %}
