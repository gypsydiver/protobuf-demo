import grpc
import greeter_pb2

channel = grpc.insecure_channel(':5000')
stub = greeter_pb2.GreeterStub(channel)

info = greeter_pb2.Info()
info.name = 'User'
info.expression = 'Hell-o'

info.peeps.add().name = 'Roland'
info.peeps.add().name = 'Jake'
# info.peeps.add().name = 'Eddie'

print(stub.Greet(info).body)
