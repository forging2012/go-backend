#!/usr/bin/env bash
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/system_info_reader.go -package=mock github.com/MongoDBNavigator/go-backend/domain/system/repository SystemInfoReader
mockgen -destination=${GOPATH}/src/github.com/MongoDBNavigator/go-backend/tests/mock/jwt_middleware.go -package=mock github.com/MongoDBNavigator/go-backend/user_interface/resource/middleware Middleware