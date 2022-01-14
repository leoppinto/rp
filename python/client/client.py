import grpc

import product_pb2
import product_pb2_grpc

channel = grpc.insecure_channel('192.168.0.15:8300')
stub = product_pb2_grpc.ProductServiceStub(channel)
requestMsg = product_pb2.Product(name="Novo Produto")
response = stub.CreateProduct(requestMsg)
print(type(response))
print(response)