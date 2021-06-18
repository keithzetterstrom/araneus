import asyncio

import items_pb2_grpc
import items_pb2
from utils import get_signature
import grpc


class ItemServicer(items_pb2_grpc.ItemServiceServicer):
    def SetSignature(self, request, context):
        s = get_signature(request)
        return items_pb2.ItemGRPC(Signature=s)


async def serve() -> None:
    server = grpc.aio.server()
    items_pb2_grpc.add_ItemServiceServicer_to_server(
        ItemServicer(), server)
    server.add_insecure_port('[::]:50051')
    await server.start()
    await server.wait_for_termination()


if __name__ == '__main__':
    asyncio.get_event_loop().run_until_complete(serve())
