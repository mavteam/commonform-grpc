// servers
import grpcServer from './grpc-server'

// logging
import pino from 'pino'

// set log level
const logLevel = process.env.LOG_LEVEL || 'info'
const serverLog = pino({
  level: logLevel
})

// gracefully shutdown
const trap = function () {
  serverLog.trace({ event: 'shutdown signal' })
  cleanup()
}
process.on('SIGTERM', trap)
process.on('SIGQUIT', trap)
process.on('SIGINT', trap)
process.on('uncaughtException', function (exception) {
  serverLog.error({ exception: exception })
  cleanup()
})

// clean everything up
var grpcS
function cleanup () {
  grpcS.tryShutdown(function () {
    serverLog.trace({ event: 'grpc server gracefully shutdown' })
  })

  serverLog.warn({ event: 'engine shutdown' })
  process.exit(0)
};

// turn on grpc server
const grpcServerHost = process.env.GRPC_SERVER_HOST || '127.0.0.1'
const grpcServerPort = process.env.GRPC_SERVER_PORT || 8081
serverLog.debug({ event: 'booting grpc server' })
grpcS = grpcServer(grpcServerHost, grpcServerPort, serverLog)
grpcS.start()
serverLog.warn({ event: 'grpc server listening', host: grpcServerHost, port: grpcServerPort })
