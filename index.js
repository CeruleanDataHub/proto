
const CURRENT_LOOP_PROTO_PATH = __dirname + '/currentLoop.proto';
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(
    CURRENT_LOOP_PROTO_PATH,
    {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true
    });
const CurrentLoopProto = grpc.loadPackageDefinition(packageDefinition);

module.exports = {
    CurrentLoopProto: CurrentLoopProto
};