# dgate

A command-line interface for DeployGate

## Installation

```
$ go get github.com/DeployGate/dgate
```

## Usage

### Login

```
$ dgate login --email your@email.com --password your_password
```

### Push
Push/Update an application:

```
$ dgate push [package file path] --message release_message
```

See https://deploygate.com/docs/cli for more information.

## Development

### Setup

```
$ go get github.com/mattn/gom
$ gom install
```

### run & build

```
$ gom run *.go
$ gom build
```

## License

Copyright (C) 2015 DeployGate All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
