import grpc

# import the generated classes
import product_pb2
import product_pb2_grpc

# open a gRPC channel
channel = grpc.insecure_channel('192.168.0.15:8300')

# create a stub (client)
stub = product_pb2_grpc.ProductServiceStub(channel)

# create a valid request message
requestMsg = product_pb2.Product(name="Novo Produto")

# make the call
response = stub.CreateProduct(requestMsg)


print(type(response))
print(response)