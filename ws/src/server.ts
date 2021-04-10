import { createServer } from "http";
import { Server, Socket } from "socket.io";
import { Consumer, KafkaClient } from "kafka-node";

const httpServer = createServer();
const options = {
  cors: {
    origin: "*"
  }
}
const io = new Server(httpServer, options);

var client = new KafkaClient({kafkaHost: "kafka:9093"});
var consumer = new Consumer(client, [{ topic: 'websocket'}], {});
io.on("connection", (socket: Socket) => {
  console.log("Connected: " + socket.id);
  io.emit("message", "Connected (Server" + socket.id + ")");
  socket.on("message", (message) => {
    console.log("Message: " + message);

  })
  consumer.on('message', function(message) {
    console.log("Update Score: " + message.value);
    io.emit("update-score", message.value);
  })
});

httpServer.listen(8008);