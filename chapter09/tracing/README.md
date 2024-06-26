# Tracing with OpenTelemetry and Jaeger

TODO: fill in the walk through

## Running this example
- `docker compose up -d`
- Once started the client application will periodically send requests to the server. Distributed traces will be collected for the requests and responses, then exported for analysis in Jaeger. To view the traces in Jaeger, open http://localhost:16686.

If you see something like:
```bash
docker-compose up -d
Traceback (most recent call last):
  File "urllib3/connectionpool.py", line 670, in urlopen
  File "urllib3/connectionpool.py", line 392, in _make_request
```
This indicates that you aren't running docker. Make sure you have docker installed and it is running.

## Tearing down this example
- `docker compose down`

## Influences / Credit
The code in this demo was heavily influenced from the example application in https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/examples/demo
which carries the following license.
```
// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
```

See also: [OTEL License](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/10cfdaac1387b4df7a525c3050ce18ec8f8068be/LICENSE)
